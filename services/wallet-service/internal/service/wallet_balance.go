package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/game_engine/wallet-service/internal/model"
)

func (s *WalletService) GetBalance(ctx context.Context, userID, currency string, balanceType string) (*model.Wallet, error) {
	cacheKey := fmt.Sprintf("balance:%s:%s:%s", userID, currency, balanceType)
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var wallet model.Wallet
		if json.Unmarshal([]byte(cached), &wallet) == nil {
			return &wallet, nil
		}
	}

	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, userID, currency, balanceType)
	if err != nil {
		return nil, err
	}

	if wallet == nil {
		wallet = &model.Wallet{
			UserID:       userID,
			Currency:     currency,
			BalanceType:  balanceType,
			Amount:       0,
			LockedAmount: 0,
		}
		if err := s.repo.CreateWallet(ctx, wallet); err != nil {
			return nil, err
		}
	}

	data, _ := json.Marshal(wallet)
	s.redis.Set(ctx, cacheKey, data, time.Duration(s.cfg.Cache.BalanceTTL)*time.Second)

	return wallet, nil
}

func (s *WalletService) GetAllBalances(ctx context.Context, userID string) ([]*model.Wallet, error) {
	cacheKey := fmt.Sprintf("balances:%s", userID)
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var wallets []*model.Wallet
		if json.Unmarshal([]byte(cached), &wallets) == nil {
			return wallets, nil
		}
	}

	wallets, err := s.repo.GetWalletsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(wallets) == 0 {
		defaultWallets := []*model.Wallet{
			{UserID: userID, Currency: CurrencyUSD, BalanceType: BalanceTypeReal, Amount: 0, LockedAmount: 0},
			{UserID: userID, Currency: CurrencyUSD, BalanceType: BalanceTypeBonus, Amount: 0, LockedAmount: 0},
		}
		for _, w := range defaultWallets {
			if err := s.repo.CreateWallet(ctx, w); err != nil {
				return nil, err
			}
			wallets = append(wallets, w)
		}
	}

	data, _ := json.Marshal(wallets)
	s.redis.Set(ctx, cacheKey, data, time.Duration(s.cfg.Cache.BalanceTTL)*time.Second)

	return wallets, nil
}
