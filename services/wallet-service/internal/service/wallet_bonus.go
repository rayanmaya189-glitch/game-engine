package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/game_engine/wallet-service/internal/model"
	"github.com/game_engine/wallet-service/internal/repository"
)

func (s *WalletService) CreateBonusCredit(ctx context.Context, userID, currency, bonusType, bonusCode string, amount int64, expiresAt *time.Time) (*model.BonusTransaction, error) {
	if amount <= 0 {
		return nil, errors.New("bonus amount must be positive")
	}
	if amount > s.cfg.Bonus.MaxBonusAmount {
		return nil, fmt.Errorf("bonus amount cannot exceed %d", s.cfg.Bonus.MaxBonusAmount)
	}

	bt := &model.BonusTransaction{
		UserID:             userID,
		BonusType:          bonusType,
		Currency:           currency,
		Amount:             amount,
		WageringMultiplier: s.cfg.Bonus.DefaultWageringMultiplier,
		WageringRequired:   amount * int64(s.cfg.Bonus.DefaultWageringMultiplier),
		WageringMet:        0,
		BonusCode:          bonusCode,
		Status:             BonusStatusActive,
		ExpiresAt:          expiresAt,
	}

	if bt.ExpiresAt == nil {
		expiryDays := s.cfg.Bonus.BonusExpiryDays
		expires := time.Now().AddDate(0, 0, expiryDays)
		bt.ExpiresAt = &expires
	}

	tx := &model.Transaction{
		UserID:      userID,
		Type:        TransactionTypeBonus,
		Status:      TransactionStatusCompleted,
		Currency:    currency,
		Amount:      amount,
		BonusAmount: amount,
		NetAmount:   amount,
		Description: fmt.Sprintf("Bonus credit: %s", bonusCode),
	}

	err := s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		if err := txRepo.CreateTransaction(ctx, tx); err != nil {
			return err
		}
		bt.TransactionID = tx.TransactionID

		if err := txRepo.CreateBonusTransaction(ctx, bt); err != nil {
			return err
		}

		bonusWallet, err := txRepo.GetWalletByUserIDAndType(ctx, userID, currency, BalanceTypeBonus)
		if err != nil {
			return err
		}
		if bonusWallet == nil {
			bonusWallet = &model.Wallet{
				UserID:      userID,
				Currency:    currency,
				BalanceType: BalanceTypeBonus,
				Amount:      0,
			}
			if err := txRepo.CreateWallet(ctx, bonusWallet); err != nil {
				return err
			}
		}

		newAmount := bonusWallet.Amount + amount
		return txRepo.UpdateWalletAmount(ctx, bonusWallet.ID, newAmount, bonusWallet.LockedAmount, bonusWallet.Version)
	})

	if err != nil {
		return nil, err
	}

	s.invalidateBalanceCache(ctx, userID, currency)

	s.publishEvent(ctx, EventBonusCredited, map[string]interface{}{
		"transaction_id":    bt.TransactionID,
		"user_id":           userID,
		"amount":            amount,
		"currency":          currency,
		"bonus_type":        bonusType,
		"bonus_code":        bonusCode,
		"wagering_required": bt.WageringRequired,
	})

	return bt, nil
}
