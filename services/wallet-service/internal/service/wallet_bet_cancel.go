package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/game_engine/wallet-service/internal/model"
	"github.com/game_engine/wallet-service/internal/repository"
)

func (s *WalletService) CancelBet(ctx context.Context, betID, reason string) (*model.Bet, error) {
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

	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		newAmount := wallet.Amount + bet.Stake
		newLocked := wallet.LockedAmount - bet.Stake
		if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, newLocked, wallet.Version); err != nil {
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
			Description: fmt.Sprintf("Bet cancelled: %s", reason),
		}
		if err := txRepo.CreateTransaction(ctx, refundTx); err != nil {
			return err
		}

		return txRepo.UpdateBetSettlement(ctx, betID, BetSettlementTypeCancelled, TransactionStatusCompleted, 0)
	})

	if err != nil {
		return nil, err
	}

	s.invalidateBalanceCache(ctx, bet.UserID, bet.Currency)
	lockKey := fmt.Sprintf("bet_lock:%s:%s", bet.UserID, bet.GameID)
	s.redis.Del(ctx, lockKey)

	s.publishEvent(ctx, EventBetCancelled, map[string]interface{}{
		"bet_id":  bet.BetID,
		"user_id": bet.UserID,
		"reason":  reason,
	})

	bet, _ = s.repo.GetBetByID(ctx, betID)
	return bet, nil
}
