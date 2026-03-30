package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/game_engine/wallet-service/internal/model"
	"github.com/game_engine/wallet-service/internal/repository"
)

func (s *WalletService) CreateWithdrawal(ctx context.Context, userID, currency, withdrawalMethodID string, amount int64) (*model.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("withdrawal amount must be positive")
	}
	if amount < s.cfg.Withdrawal.MinWithdrawalAmount || amount > s.cfg.Withdrawal.MaxWithdrawalAmount {
		return nil, fmt.Errorf("withdrawal amount must be between %d and %d", s.cfg.Withdrawal.MinWithdrawalAmount, s.cfg.Withdrawal.MaxWithdrawalAmount)
	}

	dailyTotal, err := s.repo.GetDailyWithdrawalTotal(ctx, userID, currency)
	if err != nil {
		return nil, err
	}
	if dailyTotal+amount > s.cfg.Withdrawal.DailyWithdrawalLimit {
		return nil, fmt.Errorf("daily withdrawal limit exceeded")
	}

	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, userID, currency, BalanceTypeReal)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	available := wallet.Amount - wallet.LockedAmount
	if available < amount {
		return nil, errors.New("insufficient balance")
	}

	tx := &model.Transaction{
		UserID:        userID,
		Type:          TransactionTypeWithdrawal,
		Status:        TransactionStatusPending,
		Currency:      currency,
		Amount:        amount,
		NetAmount:     amount,
		PaymentMethod: withdrawalMethodID,
		Description:   "Withdrawal requested",
	}

	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		if err := txRepo.CreateTransaction(ctx, tx); err != nil {
			return err
		}
		newLocked := wallet.LockedAmount + amount
		return txRepo.UpdateWalletAmount(ctx, wallet.ID, wallet.Amount, newLocked, wallet.Version)
	})

	if err != nil {
		return nil, err
	}

	s.invalidateBalanceCache(ctx, userID, currency)

	s.publishEvent(ctx, EventWithdrawalCreated, map[string]interface{}{
		"transaction_id":    tx.TransactionID,
		"user_id":           userID,
		"amount":            amount,
		"currency":          currency,
		"approval_required": s.cfg.Withdrawal.ApprovalRequired,
	})

	return tx, nil
}

func (s *WalletService) ConfirmWithdrawal(ctx context.Context, txID, providerReference string, status string) (*model.Transaction, error) {
	tx, err := s.repo.GetTransactionByID(ctx, txID)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, errors.New("transaction not found")
	}
	if tx.Status != TransactionStatusPending {
		return nil, errors.New("transaction is not pending")
	}

	now := time.Now()
	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, tx.UserID, tx.Currency, BalanceTypeReal)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		newStatus := status
		newLocked := wallet.LockedAmount - tx.NetAmount

		if status == TransactionStatusCompleted {
			newAmount := wallet.Amount - tx.NetAmount
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, newLocked, wallet.Version); err != nil {
				return err
			}
		} else if status == TransactionStatusCancelled {
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, wallet.Amount, newLocked, wallet.Version); err != nil {
				return err
			}
		}

		return txRepo.UpdateTransactionStatus(ctx, txID, newStatus, &now)
	})

	if err != nil {
		return nil, err
	}

	tx, _ = s.repo.GetTransactionByID(ctx, txID)
	tx.ProcessedAt = &now
	s.invalidateBalanceCache(ctx, tx.UserID, tx.Currency)

	s.publishEvent(ctx, EventWithdrawalCompleted, map[string]interface{}{
		"transaction_id":     tx.TransactionID,
		"user_id":            tx.UserID,
		"amount":             tx.NetAmount,
		"currency":           tx.Currency,
		"status":             status,
		"provider_reference": providerReference,
	})

	return tx, nil
}
