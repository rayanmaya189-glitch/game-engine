package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/game_engine/wallet-service/internal/model"
	"github.com/game_engine/wallet-service/internal/repository"
)

func (s *WalletService) ReverseTransaction(ctx context.Context, txID, reason string) (*model.Transaction, error) {
	origTx, err := s.repo.GetTransactionByID(ctx, txID)
	if err != nil {
		return nil, err
	}
	if origTx == nil {
		return nil, errors.New("transaction not found")
	}
	if origTx.Status == TransactionStatusReversed {
		return nil, errors.New("transaction already reversed")
	}

	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, origTx.UserID, origTx.Currency, BalanceTypeReal)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	now := time.Now()
	reverseTx := &model.Transaction{
		UserID:           origTx.UserID,
		Type:             TransactionTypeReversal,
		Status:           TransactionStatusCompleted,
		Currency:         origTx.Currency,
		Amount:           origTx.Amount,
		NetAmount:        origTx.NetAmount,
		PaymentReference: origTx.PaymentReference,
		Description:      fmt.Sprintf("Reversal of %s: %s", txID, reason),
	}

	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		if err := txRepo.CreateTransaction(ctx, reverseTx); err != nil {
			return err
		}

		switch origTx.Type {
		case TransactionTypeDeposit:
			newAmount := wallet.Amount - origTx.NetAmount
			if newAmount < 0 {
				return errors.New("insufficient balance for reversal")
			}
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount, wallet.Version); err != nil {
				return err
			}
		case TransactionTypeWithdrawal:
			newAmount := wallet.Amount + origTx.NetAmount
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount, wallet.Version); err != nil {
				return err
			}
		case TransactionTypeWin:
			newAmount := wallet.Amount - origTx.NetAmount
			if newAmount < 0 {
				return errors.New("insufficient balance for reversal")
			}
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount, wallet.Version); err != nil {
				return err
			}
		}

		return txRepo.UpdateTransactionStatus(ctx, txID, TransactionStatusReversed, &now)
	})

	if err != nil {
		return nil, err
	}

	s.invalidateBalanceCache(ctx, origTx.UserID, origTx.Currency)

	s.publishEvent(ctx, EventTransactionReversed, map[string]interface{}{
		"original_transaction_id": txID,
		"reversal_transaction_id": reverseTx.TransactionID,
		"user_id":                 origTx.UserID,
		"amount":                  origTx.NetAmount,
		"reason":                  reason,
	})

	return reverseTx, nil
}

func (s *WalletService) GetTransactionHistory(ctx context.Context, userID string, txTypes []string, statuses []string, startDate, endDate time.Time, page, pageSize int) ([]*model.Transaction, int, error) {
	return s.repo.GetTransactionHistory(ctx, userID, txTypes, statuses, startDate, endDate, page, pageSize)
}

func (s *WalletService) GetPendingBets(ctx context.Context, userID, gameID string, page, pageSize int) ([]*model.Bet, int, error) {
	return s.repo.GetPendingBets(ctx, userID, gameID, page, pageSize)
}

func (s *WalletService) GetBonusBalance(ctx context.Context, userID, currency string) (int64, error) {
	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, userID, currency, BalanceTypeBonus)
	if err != nil {
		return 0, err
	}
	if wallet == nil {
		return 0, nil
	}
	return wallet.Amount, nil
}

func (s *WalletService) GetActiveBonuses(ctx context.Context, userID, currency string) ([]*model.BonusTransaction, error) {
	return s.repo.GetActiveBonusTransactions(ctx, userID, currency)
}
