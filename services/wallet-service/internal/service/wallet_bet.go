package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/game_engine/wallet-service/internal/model"
	"github.com/game_engine/wallet-service/internal/repository"
)

func (s *WalletService) PlaceBet(ctx context.Context, userID, gameID, betType, selection, odds, currency string, amount int64) (*model.Bet, *model.Wallet, error) {
	if amount <= 0 {
		return nil, nil, errors.New("bet amount must be positive")
	}
	if amount < s.cfg.Betting.MinBetAmount || amount > s.cfg.Betting.MaxBetAmount {
		return nil, nil, fmt.Errorf("bet amount must be between %d and %d", s.cfg.Betting.MinBetAmount, s.cfg.Betting.MaxBetAmount)
	}

	lockKey := fmt.Sprintf("bet_lock:%s:%s", userID, gameID)
	lockExists, err := s.redis.Exists(ctx, lockKey).Result()
	if err == nil && lockExists > 0 {
		return nil, nil, errors.New("pending bet exists for this game")
	}
	s.redis.Set(ctx, lockKey, "1", time.Duration(s.cfg.Cache.BetLockTTL)*time.Second)

	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, userID, currency, BalanceTypeReal)
	if err != nil {
		return nil, nil, err
	}
	if wallet == nil {
		return nil, nil, errors.New("wallet not found")
	}

	bonusWallet, _ := s.repo.GetWalletByUserIDAndType(ctx, userID, currency, BalanceTypeBonus)
	var availableBalance int64
	if bonusWallet != nil {
		availableBalance = wallet.Amount + bonusWallet.Amount - wallet.LockedAmount
	} else {
		availableBalance = wallet.Amount - wallet.LockedAmount
	}
	if availableBalance < amount {
		return nil, nil, errors.New("insufficient balance")
	}

	potentialWin := amount * 2

	bet := &model.Bet{
		UserID:         userID,
		GameID:         gameID,
		BetType:        betType,
		Selection:      selection,
		Odds:           odds,
		Stake:          amount,
		PotentialWin:   potentialWin,
		Status:         TransactionStatusPending,
		SettlementType: BetSettlementTypePending,
	}

	tx := &model.Transaction{
		UserID:      userID,
		Type:        TransactionTypeBet,
		Status:      TransactionStatusCompleted,
		Currency:    currency,
		Amount:      amount,
		NetAmount:   amount,
		GameID:      gameID,
		BetID:       bet.BetID,
		Description: "Bet placed",
	}

	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		if err := txRepo.CreateBet(ctx, bet); err != nil {
			return err
		}
		tx.BetID = bet.BetID

		if err := txRepo.CreateTransaction(ctx, tx); err != nil {
			return err
		}

		if amount <= wallet.Amount-wallet.LockedAmount {
			newLocked := wallet.LockedAmount + amount
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, wallet.Amount, newLocked, wallet.Version); err != nil {
				return err
			}
			wallet.LockedAmount = newLocked
		} else {
			realPart := wallet.Amount - wallet.LockedAmount
			bonusNeeded := amount - realPart

			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, wallet.Amount, wallet.Amount, wallet.Version); err != nil {
				return err
			}
			wallet.LockedAmount = wallet.Amount

			if bonusWallet != nil && bonusWallet.Amount > 0 {
				bonusToUse := bonusNeeded
				if bonusWallet.Amount < bonusNeeded {
					bonusToUse = bonusWallet.Amount
				}
				newLocked := bonusWallet.LockedAmount + bonusToUse
				if err := txRepo.UpdateWalletAmount(ctx, bonusWallet.ID, bonusWallet.Amount, newLocked, bonusWallet.Version); err != nil {
					return err
				}
				bonusWallet.LockedAmount = newLocked
			}
		}
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	s.invalidateBalanceCache(ctx, userID, currency)

	s.publishEvent(ctx, EventBetPlaced, map[string]interface{}{
		"bet_id":        bet.BetID,
		"user_id":       userID,
		"game_id":       gameID,
		"amount":        amount,
		"currency":      currency,
		"potential_win": potentialWin,
	})

	return bet, wallet, nil
}

func (s *WalletService) SettleBet(ctx context.Context, betID string, settlementType string, winAmount int64, result string) (*model.Bet, error) {
	bet, err := s.repo.GetBetByID(ctx, betID)
	if err != nil {
		return nil, err
	}
	if bet == nil {
		return nil, errors.New("bet not found")
	}
	if bet.SettlementType != BetSettlementTypePending {
		return nil, errors.New("bet already settled")
	}

	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, bet.UserID, bet.Currency, BalanceTypeReal)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	bonusWallet, _ := s.repo.GetWalletByUserIDAndType(ctx, bet.UserID, bet.Currency, BalanceTypeBonus)

	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		switch settlementType {
		case BetSettlementTypeWon:
			newAmount := wallet.Amount + winAmount
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount-bet.Stake, wallet.Version); err != nil {
				return err
			}
			winTx := &model.Transaction{
				UserID:      bet.UserID,
				Type:        TransactionTypeWin,
				Status:      TransactionStatusCompleted,
				Currency:    bet.Currency,
				Amount:      winAmount,
				NetAmount:   winAmount,
				GameID:      bet.GameID,
				BetID:       bet.BetID,
				Description: fmt.Sprintf("Bet won: %s", result),
			}
			if err := txRepo.CreateTransaction(ctx, winTx); err != nil {
				return err
			}
			if bonusWallet != nil && bonusWallet.LockedAmount > 0 {
				bonusUsed := bet.Stake
				if bonusWallet.Amount < bonusUsed {
					bonusUsed = bonusWallet.Amount
				}
				bonuses, _ := s.repo.GetActiveBonusTransactions(ctx, bet.UserID, bet.Currency)
				for _, b := range bonuses {
					newWagering := b.WageringMet + bonusUsed
					if err := txRepo.UpdateBonusWagering(ctx, b.ID, newWagering); err != nil {
						return err
					}
				}
			}

		case BetSettlementTypeLost:
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, wallet.Amount, wallet.LockedAmount-bet.Stake, wallet.Version); err != nil {
				return err
			}
			if bonusWallet != nil && bonusWallet.LockedAmount > 0 {
				bonusUsed := bet.Stake
				if bonusWallet.Amount < bonusUsed {
					bonusUsed = bonusWallet.Amount
				}
				if err := txRepo.UpdateWalletAmount(ctx, bonusWallet.ID, bonusWallet.Amount, bonusWallet.LockedAmount-bonusUsed, bonusWallet.Version); err != nil {
					return err
				}
			}

		case BetSettlementTypeCancelled:
			newAmount := wallet.Amount + bet.Stake
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount-bet.Stake, wallet.Version); err != nil {
				return err
			}
			refundTx := &model.Transaction{
				UserID:      bet.UserID,
				Type:        TransactionTypeRefund,
				Status:      TransactionStatusCompleted,
				Currency:    bet.Currency,
				Amount:      bet.Stake,
				NetAmount:   bet.Stake,
				GameID:      bet.GameID,
				BetID:       bet.BetID,
				Description: "Bet cancelled - refund",
			}
			if err := txRepo.CreateTransaction(ctx, refundTx); err != nil {
				return err
			}
			if bonusWallet != nil && bonusWallet.LockedAmount > 0 {
				bonusUsed := bet.Stake
				if bonusWallet.Amount < bonusUsed {
					bonusUsed = bonusWallet.Amount
				}
				if err := txRepo.UpdateWalletAmount(ctx, bonusWallet.ID, bonusWallet.Amount, bonusWallet.LockedAmount-bonusUsed, bonusWallet.Version); err != nil {
					return err
				}
			}
		}
		return txRepo.UpdateBetSettlement(ctx, betID, settlementType, TransactionStatusCompleted, winAmount)
	})

	if err != nil {
		return nil, err
	}

	s.invalidateBalanceCache(ctx, bet.UserID, bet.Currency)
	lockKey := fmt.Sprintf("bet_lock:%s:%s", bet.UserID, bet.GameID)
	s.redis.Del(ctx, lockKey)

	s.publishEvent(ctx, EventBetSettled, map[string]interface{}{
		"bet_id":          bet.BetID,
		"user_id":         bet.UserID,
		"game_id":         bet.GameID,
		"settlement_type": settlementType,
		"win_amount":      winAmount,
		"result":          result,
	})

	bet, _ = s.repo.GetBetByID(ctx, betID)
	return bet, nil
}

