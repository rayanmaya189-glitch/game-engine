package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/game_engine/wallet-service/internal/model"
	"github.com/game_engine/wallet-service/internal/repository"
	"github.com/google/uuid"
)

func (s *WalletService) CreateDeposit(ctx context.Context, userID, currency, paymentMethod, paymentProvider, bonusCode string, amount int64) (*model.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("deposit amount must be positive")
	}
	if amount < 10 || amount > 100000 {
		return nil, errors.New("deposit amount must be between 10 and 100000")
	}

	tx := &model.Transaction{
		UserID:          userID,
		Type:            TransactionTypeDeposit,
		Status:          TransactionStatusPending,
		Currency:        currency,
		Amount:          amount,
		NetAmount:       amount,
		PaymentMethod:   paymentMethod,
		PaymentProvider: paymentProvider,
		Description:     "Deposit initiated",
	}

	err := s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		if err := txRepo.CreateTransaction(ctx, tx); err != nil {
			return err
		}

		wallet, err := txRepo.GetWalletByUserIDAndType(ctx, userID, currency, BalanceTypeReal)
		if err != nil {
			return err
		}
		if wallet == nil {
			wallet = &model.Wallet{
				UserID:      userID,
				Currency:    currency,
				BalanceType: BalanceTypeReal,
				Amount:      0,
			}
			if err := txRepo.CreateWallet(ctx, wallet); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	s.invalidateBalanceCache(ctx, userID, currency)
	tx.PaymentReference = fmt.Sprintf("DEP-%s-%d", uuid.New().String()[:8], time.Now().Unix())

	s.publishEvent(ctx, EventDepositCreated, map[string]interface{}{
		"transaction_id":   tx.TransactionID,
		"user_id":          userID,
		"amount":           amount,
		"currency":         currency,
		"payment_method":   paymentMethod,
		"payment_provider": paymentProvider,
	})

	return tx, nil
}

func (s *WalletService) ConfirmDeposit(ctx context.Context, txID, providerStatus string) (*model.Transaction, error) {
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

	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		if err := txRepo.UpdateTransactionStatus(ctx, txID, TransactionStatusCompleted, &now); err != nil {
			return err
		}

		wallet, err := txRepo.GetWalletByUserIDAndType(ctx, tx.UserID, tx.Currency, BalanceTypeReal)
		if err != nil {
			return err
		}
		if wallet == nil {
			return errors.New("wallet not found")
		}

		newAmount := wallet.Amount + tx.NetAmount
		return txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount, wallet.Version)
	})

	if err != nil {
		return nil, err
	}

	tx, _ = s.repo.GetTransactionByID(ctx, txID)
	tx.ProcessedAt = &now
	s.invalidateBalanceCache(ctx, tx.UserID, tx.Currency)

	s.publishEvent(ctx, EventDepositCompleted, map[string]interface{}{
		"transaction_id":  tx.TransactionID,
		"user_id":         tx.UserID,
		"amount":          tx.NetAmount,
		"currency":        tx.Currency,
		"provider_status": providerStatus,
	})

	return tx, nil
}
