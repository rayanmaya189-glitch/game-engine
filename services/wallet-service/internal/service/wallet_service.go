package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/game_engine/wallet-service/internal/config"
	"github.com/game_engine/wallet-service/internal/model"
	"github.com/game_engine/wallet-service/internal/repository"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

// WalletService handles wallet business logic
type WalletService struct {
	repo  *repository.WalletRepository
	redis *redis.Client
	nats  *nats.Conn
	cfg   *config.Config
}

// NewWalletService creates a new wallet service
func NewWalletService(repo *repository.WalletRepository, redis *redis.Client, natsConn *nats.Conn, cfg *config.Config) (*WalletService, error) {
	return &WalletService{
		repo:  repo,
		redis: redis,
		nats:  natsConn,
		cfg:   cfg,
	}, nil
}

// NATS Event types
const (
	EventDepositCreated      = "wallet.events.deposit_created"
	EventDepositCompleted    = "wallet.events.deposit_completed"
	EventWithdrawalCreated   = "wallet.events.withdrawal_created"
	EventWithdrawalCompleted = "wallet.events.withdrawal_completed"
	EventBetPlaced           = "wallet.events.bet_placed"
	EventBetSettled          = "wallet.events.bet_settled"
	EventBetCancelled        = "wallet.events.bet_cancelled"
	EventBonusCredited       = "wallet.events.bonus_credited"
	EventTransactionReversed = "wallet.events.transaction_reversed"
)

// Constants
const (
	CurrencyUSD = "CURRENCY_USD"
	CurrencyEUR = "CURRENCY_EUR"
	CurrencyGBP = "CURRENCY_GBP"

	BalanceTypeReal  = "BALANCE_TYPE_REAL"
	BalanceTypeBonus = "BALANCE_TYPE_BONUS"
	BalanceTypePromo = "BALANCE_TYPE_PROMOTIONAL"

	TransactionTypeDeposit    = "TRANSACTION_TYPE_DEPOSIT"
	TransactionTypeWithdrawal = "TRANSACTION_TYPE_WITHDRAWAL"
	TransactionTypeBet        = "TRANSACTION_TYPE_BET"
	TransactionTypeWin        = "TRANSACTION_TYPE_WIN"
	TransactionTypeBonus      = "TRANSACTION_TYPE_BONUS"
	TransactionTypeRefund     = "TRANSACTION_TYPE_REFUND"
	TransactionTypeReversal   = "TRANSACTION_TYPE_REVERSAL"
	TransactionTypeAdjustment = "TRANSACTION_TYPE_ADJUSTMENT"

	TransactionStatusPending    = "TRANSACTION_STATUS_PENDING"
	TransactionStatusProcessing = "TRANSACTION_STATUS_PROCESSING"
	TransactionStatusCompleted  = "TRANSACTION_STATUS_COMPLETED"
	TransactionStatusFailed     = "TRANSACTION_STATUS_FAILED"
	TransactionStatusCancelled  = "TRANSACTION_STATUS_CANCELLED"
	TransactionStatusReversed   = "TRANSACTION_STATUS_REVERSED"

	BetSettlementTypePending   = "BET_SETTLEMENT_TYPE_PENDING"
	BetSettlementTypeWon       = "BET_SETTLEMENT_TYPE_WON"
	BetSettlementTypeLost      = "BET_SETTLEMENT_TYPE_LOST"
	BetSettlementTypeCancelled = "BET_SETTLEMENT_TYPE_CANCELLED"

	BonusTypeWelcome   = "BONUS_TYPE_WELCOME"
	BonusTypeDeposit   = "BONUS_TYPE_DEPOSIT"
	BonusTypeNoDeposit = "BONUS_TYPE_NO_DEPOSIT"
	BonusTypeFreeSpins = "BONUS_TYPE_FREE_SPINS"
	BonusTypeCashback  = "BONUS_TYPE_CASHBACK"
	BonusTypeLoyalty   = "BONUS_TYPE_LOYALTY"

	BonusStatusActive  = "ACTIVE"
	BonusStatusUsed    = "USED"
	BonusStatusExpired = "EXPIRED"
)

// GetBalance retrieves player balance
func (s *WalletService) GetBalance(ctx context.Context, userID, currency string, balanceType string) (*model.Wallet, error) {
	// Try cache first
	cacheKey := fmt.Sprintf("balance:%s:%s:%s", userID, currency, balanceType)
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var wallet model.Wallet
		if json.Unmarshal([]byte(cached), &wallet) == nil {
			return &wallet, nil
		}
	}

	// Fetch from database
	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, userID, currency, balanceType)
	if err != nil {
		return nil, err
	}

	if wallet == nil {
		// Create wallet if not exists
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

	// Cache the result
	data, _ := json.Marshal(wallet)
	s.redis.Set(ctx, cacheKey, data, time.Duration(s.cfg.Cache.BalanceTTL)*time.Second)

	return wallet, nil
}

// GetAllBalances retrieves all currency balances for a user
func (s *WalletService) GetAllBalances(ctx context.Context, userID string) ([]*model.Wallet, error) {
	// Try cache first
	cacheKey := fmt.Sprintf("balances:%s", userID)
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var wallets []*model.Wallet
		if json.Unmarshal([]byte(cached), &wallets) == nil {
			return wallets, nil
		}
	}

	// Fetch from database
	wallets, err := s.repo.GetWalletsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// If no wallets exist, create default ones
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

	// Cache the result
	data, _ := json.Marshal(wallets)
	s.redis.Set(ctx, cacheKey, data, time.Duration(s.cfg.Cache.BalanceTTL)*time.Second)

	return wallets, nil
}

// CreateDeposit initiates a deposit
func (s *WalletService) CreateDeposit(ctx context.Context, userID, currency, paymentMethod, paymentProvider, bonusCode string, amount int64) (*model.Transaction, error) {
	// Validate amount
	if amount <= 0 {
		return nil, errors.New("deposit amount must be positive")
	}

	// Check min/max limits
	if amount < 10 || amount > 100000 {
		return nil, errors.New("deposit amount must be between 10 and 100000")
	}

	// Create pending transaction
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

	// Use transaction for atomicity
	err := s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		// Create transaction record
		if err := txRepo.CreateTransaction(ctx, tx); err != nil {
			return err
		}

		// Ensure wallet exists
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

	// Invalidate balance cache
	s.invalidateBalanceCache(ctx, userID, currency)

	// Generate payment reference
	tx.PaymentReference = fmt.Sprintf("DEP-%s-%d", uuid.New().String()[:8], time.Now().Unix())

	// Publish event
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

// ConfirmDeposit confirms a deposit completion
func (s *WalletService) ConfirmDeposit(ctx context.Context, txID, providerStatus string) (*model.Transaction, error) {
	// Get existing transaction
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

	// Use transaction for atomicity
	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		// Update transaction status
		if err := txRepo.UpdateTransactionStatus(ctx, txID, TransactionStatusCompleted, &now); err != nil {
			return err
		}

		// Get and update wallet
		wallet, err := txRepo.GetWalletByUserIDAndType(ctx, tx.UserID, tx.Currency, BalanceTypeReal)
		if err != nil {
			return err
		}
		if wallet == nil {
			return errors.New("wallet not found")
		}

		// Update balance
		newAmount := wallet.Amount + tx.NetAmount
		if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount, wallet.Version); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Refresh transaction
	tx, _ = s.repo.GetTransactionByID(ctx, txID)
	tx.ProcessedAt = &now

	// Invalidate balance cache
	s.invalidateBalanceCache(ctx, tx.UserID, tx.Currency)

	// Publish event
	s.publishEvent(ctx, EventDepositCompleted, map[string]interface{}{
		"transaction_id":  tx.TransactionID,
		"user_id":         tx.UserID,
		"amount":          tx.NetAmount,
		"currency":        tx.Currency,
		"provider_status": providerStatus,
	})

	return tx, nil
}

// CreateWithdrawal initiates a withdrawal request
func (s *WalletService) CreateWithdrawal(ctx context.Context, userID, currency, withdrawalMethodID string, amount int64) (*model.Transaction, error) {
	// Validate amount
	if amount <= 0 {
		return nil, errors.New("withdrawal amount must be positive")
	}

	// Check min/max limits
	if amount < s.cfg.Withdrawal.MinWithdrawalAmount || amount > s.cfg.Withdrawal.MaxWithdrawalAmount {
		return nil, fmt.Errorf("withdrawal amount must be between %d and %d", s.cfg.Withdrawal.MinWithdrawalAmount, s.cfg.Withdrawal.MaxWithdrawalAmount)
	}

	// Check daily limit
	dailyTotal, err := s.repo.GetDailyWithdrawalTotal(ctx, userID, currency)
	if err != nil {
		return nil, err
	}
	if dailyTotal+amount > s.cfg.Withdrawal.DailyWithdrawalLimit {
		return nil, fmt.Errorf("daily withdrawal limit exceeded")
	}

	// Get wallet balance
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

	// Create pending withdrawal transaction
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

	// Use transaction for atomicity
	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		// Lock funds by creating transaction
		if err := txRepo.CreateTransaction(ctx, tx); err != nil {
			return err
		}

		// Lock the amount
		newLocked := wallet.LockedAmount + amount
		if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, wallet.Amount, newLocked, wallet.Version); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Invalidate balance cache
	s.invalidateBalanceCache(ctx, userID, currency)

	approvalRequired := s.cfg.Withdrawal.ApprovalRequired

	// Publish event
	s.publishEvent(ctx, EventWithdrawalCreated, map[string]interface{}{
		"transaction_id":    tx.TransactionID,
		"user_id":           userID,
		"amount":            amount,
		"currency":          currency,
		"approval_required": approvalRequired,
	})

	return tx, nil
}

// ConfirmWithdrawal confirms a withdrawal
func (s *WalletService) ConfirmWithdrawal(ctx context.Context, txID, providerReference string, status string) (*model.Transaction, error) {
	// Get existing transaction
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

	// Use transaction for atomicity
	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		newStatus := status
		newLocked := wallet.LockedAmount - tx.NetAmount

		if status == TransactionStatusCompleted {
			// Deduct from balance on completion
			newAmount := wallet.Amount - tx.NetAmount
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, newLocked, wallet.Version); err != nil {
				return err
			}
		} else if status == TransactionStatusCancelled {
			// Release locked amount on cancel
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, wallet.Amount, newLocked, wallet.Version); err != nil {
				return err
			}
		}

		// Update transaction status
		return txRepo.UpdateTransactionStatus(ctx, txID, newStatus, &now)
	})

	if err != nil {
		return nil, err
	}

	// Refresh transaction
	tx, _ = s.repo.GetTransactionByID(ctx, txID)
	tx.ProcessedAt = &now

	// Invalidate balance cache
	s.invalidateBalanceCache(ctx, tx.UserID, tx.Currency)

	// Publish event
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

// PlaceBet locks funds for a bet
func (s *WalletService) PlaceBet(ctx context.Context, userID, gameID, betType, selection, odds, currency string, amount int64) (*model.Bet, *model.Wallet, error) {
	// Validate amount
	if amount <= 0 {
		return nil, nil, errors.New("bet amount must be positive")
	}

	// Check min/max bet
	if amount < s.cfg.Betting.MinBetAmount || amount > s.cfg.Betting.MaxBetAmount {
		return nil, nil, fmt.Errorf("bet amount must be between %d and %d", s.cfg.Betting.MinBetAmount, s.cfg.Betting.MaxBetAmount)
	}

	// Check bet lock in Redis (prevent double bets)
	lockKey := fmt.Sprintf("bet_lock:%s:%s", userID, gameID)
	lockExists, err := s.redis.Exists(ctx, lockKey).Result()
	if err == nil && lockExists > 0 {
		return nil, nil, errors.New("pending bet exists for this game")
	}

	// Set temporary lock
	s.redis.Set(ctx, lockKey, "1", time.Duration(s.cfg.Cache.BetLockTTL)*time.Second)

	// Get wallet
	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, userID, currency, BalanceTypeReal)
	if err != nil {
		return nil, nil, err
	}
	if wallet == nil {
		return nil, nil, errors.New("wallet not found")
	}

	// Check bonus wallet first
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

	// Calculate potential win (simplified)
	potentialWin := amount * 2 // Default 2x, should use actual odds

	// Create bet and transaction in a transaction
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

	// Use transaction for atomicity
	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		// Create bet record
		if err := txRepo.CreateBet(ctx, bet); err != nil {
			return err
		}
		tx.BetID = bet.BetID

		// Create transaction record
		if err := txRepo.CreateTransaction(ctx, tx); err != nil {
			return err
		}

		// Lock funds from real wallet first, then bonus
		if amount <= wallet.Amount-wallet.LockedAmount {
			// Lock from real wallet
			newLocked := wallet.LockedAmount + amount
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, wallet.Amount, newLocked, wallet.Version); err != nil {
				return err
			}
			wallet.LockedAmount = newLocked
		} else {
			// Need to use bonus funds
			realPart := wallet.Amount - wallet.LockedAmount
			bonusNeeded := amount - realPart

			// Lock all real balance
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, wallet.Amount, wallet.Amount, wallet.Version); err != nil {
				return err
			}
			wallet.LockedAmount = wallet.Amount

			// Lock bonus funds if available
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

	// Invalidate balance cache
	s.invalidateBalanceCache(ctx, userID, currency)

	// Publish event
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

// SettleBet processes bet result
func (s *WalletService) SettleBet(ctx context.Context, betID string, settlementType string, winAmount int64, result string) (*model.Bet, error) {
	// Get bet
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

	// Get wallet
	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, bet.UserID, bet.Currency, BalanceTypeReal)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	// Get bonus wallet if exists
	bonusWallet, _ := s.repo.GetWalletByUserIDAndType(ctx, bet.UserID, bet.Currency, BalanceTypeBonus)

	// Use transaction for atomicity
	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		if settlementType == BetSettlementTypeWon {
			// Credit winnings
			newAmount := wallet.Amount + winAmount
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount-bet.Stake, wallet.Version); err != nil {
				return err
			}

			// Create win transaction
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

			// Update bonus wagering if bonus was used
			if bonusWallet != nil && bonusWallet.LockedAmount > 0 {
				bonusUsed := bet.Stake
				if bonusWallet.Amount < bonusUsed {
					bonusUsed = bonusWallet.Amount
				}
				// Update wagering requirement
				bonuses, _ := s.repo.GetActiveBonusTransactions(ctx, bet.UserID, bet.Currency)
				for _, b := range bonuses {
					newWagering := b.WageringMet + bonusUsed
					if err := txRepo.UpdateBonusWagering(ctx, b.ID, newWagering); err != nil {
						return err
					}
				}
			}
		} else if settlementType == BetSettlementTypeLost {
			// Funds already deducted, just release lock
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, wallet.Amount, wallet.LockedAmount-bet.Stake, wallet.Version); err != nil {
				return err
			}

			// Release bonus lock if any
			if bonusWallet != nil && bonusWallet.LockedAmount > 0 {
				bonusUsed := bet.Stake
				if bonusWallet.Amount < bonusUsed {
					bonusUsed = bonusWallet.Amount
				}
				if err := txRepo.UpdateWalletAmount(ctx, bonusWallet.ID, bonusWallet.Amount, bonusWallet.LockedAmount-bonusUsed, bonusWallet.Version); err != nil {
					return err
				}
			}
		} else if settlementType == BetSettlementTypeCancelled {
			// Refund stake
			newAmount := wallet.Amount + bet.Stake
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount-bet.Stake, wallet.Version); err != nil {
				return err
			}

			// Create refund transaction
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

			// Release bonus lock if any
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

		// Update bet settlement
		return txRepo.UpdateBetSettlement(ctx, betID, settlementType, TransactionStatusCompleted, winAmount)
	})

	if err != nil {
		return nil, err
	}

	// Invalidate balance cache
	s.invalidateBalanceCache(ctx, bet.UserID, bet.Currency)

	// Remove bet lock
	lockKey := fmt.Sprintf("bet_lock:%s:%s", bet.UserID, bet.GameID)
	s.redis.Del(ctx, lockKey)

	// Publish event
	s.publishEvent(ctx, EventBetSettled, map[string]interface{}{
		"bet_id":          bet.BetID,
		"user_id":         bet.UserID,
		"game_id":         bet.GameID,
		"settlement_type": settlementType,
		"win_amount":      winAmount,
		"result":          result,
	})

	// Refresh bet
	bet, _ = s.repo.GetBetByID(ctx, betID)
	return bet, nil
}

// CancelBet cancels a pending bet
func (s *WalletService) CancelBet(ctx context.Context, betID, reason string) (*model.Bet, error) {
	// Get bet
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

	// Get wallet
	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, bet.UserID, bet.Currency, BalanceTypeReal)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	// Use transaction for atomicity
	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		// Refund stake
		newAmount := wallet.Amount + bet.Stake
		newLocked := wallet.LockedAmount - bet.Stake
		if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, newLocked, wallet.Version); err != nil {
			return err
		}

		// Create refund transaction
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

		// Update bet status
		return txRepo.UpdateBetSettlement(ctx, betID, BetSettlementTypeCancelled, TransactionStatusCompleted, 0)
	})

	if err != nil {
		return nil, err
	}

	// Invalidate balance cache
	s.invalidateBalanceCache(ctx, bet.UserID, bet.Currency)

	// Remove bet lock
	lockKey := fmt.Sprintf("bet_lock:%s:%s", bet.UserID, bet.GameID)
	s.redis.Del(ctx, lockKey)

	// Publish event
	s.publishEvent(ctx, EventBetCancelled, map[string]interface{}{
		"bet_id":  bet.BetID,
		"user_id": bet.UserID,
		"reason":  reason,
	})

	// Refresh bet
	bet, _ = s.repo.GetBetByID(ctx, betID)
	return bet, nil
}

// CreateBonusCredit adds bonus funds
func (s *WalletService) CreateBonusCredit(ctx context.Context, userID, currency, bonusType, bonusCode string, amount int64, expiresAt *time.Time) (*model.BonusTransaction, error) {
	// Validate amount
	if amount <= 0 {
		return nil, errors.New("bonus amount must be positive")
	}

	// Check max bonus
	if amount > s.cfg.Bonus.MaxBonusAmount {
		return nil, fmt.Errorf("bonus amount cannot exceed %d", s.cfg.Bonus.MaxBonusAmount)
	}

	// Create bonus transaction
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

	// Calculate expiry
	if bt.ExpiresAt == nil {
		expiryDays := s.cfg.Bonus.BonusExpiryDays
		expires := time.Now().AddDate(0, 0, expiryDays)
		bt.ExpiresAt = &expires
	}

	// Create transaction and update wallet
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
		// Create transaction record
		if err := txRepo.CreateTransaction(ctx, tx); err != nil {
			return err
		}
		bt.TransactionID = tx.TransactionID

		// Create bonus transaction record
		if err := txRepo.CreateBonusTransaction(ctx, bt); err != nil {
			return err
		}

		// Get or create bonus wallet
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

		// Credit bonus to bonus wallet
		newAmount := bonusWallet.Amount + amount
		return txRepo.UpdateWalletAmount(ctx, bonusWallet.ID, newAmount, bonusWallet.LockedAmount, bonusWallet.Version)
	})

	if err != nil {
		return nil, err
	}

	// Invalidate balance cache
	s.invalidateBalanceCache(ctx, userID, currency)

	// Publish event
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

// ReverseTransaction reverses a transaction (admin operation)
func (s *WalletService) ReverseTransaction(ctx context.Context, txID, reason string) (*model.Transaction, error) {
	// Get original transaction
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

	// Get wallet
	wallet, err := s.repo.GetWalletByUserIDAndType(ctx, origTx.UserID, origTx.Currency, BalanceTypeReal)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	// Create reversal transaction
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

	// Use transaction for atomicity
	err = s.repo.WithTx(ctx, func(txRepo *repository.WalletRepositoryWithTx) error {
		// Create reversal transaction
		if err := txRepo.CreateTransaction(ctx, reverseTx); err != nil {
			return err
		}

		// Reverse the amount based on original transaction type
		switch origTx.Type {
		case TransactionTypeDeposit:
			// Reverse deposit - deduct from balance
			newAmount := wallet.Amount - origTx.NetAmount
			if newAmount < 0 {
				return errors.New("insufficient balance for reversal")
			}
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount, wallet.Version); err != nil {
				return err
			}

		case TransactionTypeWithdrawal:
			// Reverse withdrawal - add back to balance
			newAmount := wallet.Amount + origTx.NetAmount
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount, wallet.Version); err != nil {
				return err
			}

		case TransactionTypeWin:
			// Reverse win - deduct from balance
			newAmount := wallet.Amount - origTx.NetAmount
			if newAmount < 0 {
				return errors.New("insufficient balance for reversal")
			}
			if err := txRepo.UpdateWalletAmount(ctx, wallet.ID, newAmount, wallet.LockedAmount, wallet.Version); err != nil {
				return err
			}
		}

		// Mark original transaction as reversed
		return txRepo.UpdateTransactionStatus(ctx, txID, TransactionStatusReversed, &now)
	})

	if err != nil {
		return nil, err
	}

	// Invalidate balance cache
	s.invalidateBalanceCache(ctx, origTx.UserID, origTx.Currency)

	// Publish event
	s.publishEvent(ctx, EventTransactionReversed, map[string]interface{}{
		"original_transaction_id": txID,
		"reversal_transaction_id": reverseTx.TransactionID,
		"user_id":                 origTx.UserID,
		"amount":                  origTx.NetAmount,
		"reason":                  reason,
	})

	return reverseTx, nil
}

// GetTransactionHistory retrieves transaction history
func (s *WalletService) GetTransactionHistory(ctx context.Context, userID string, txTypes []string, statuses []string, startDate, endDate time.Time, page, pageSize int) ([]*model.Transaction, int, error) {
	return s.repo.GetTransactionHistory(ctx, userID, txTypes, statuses, startDate, endDate, page, pageSize)
}

// GetPendingBets retrieves pending bets
func (s *WalletService) GetPendingBets(ctx context.Context, userID, gameID string, page, pageSize int) ([]*model.Bet, int, error) {
	return s.repo.GetPendingBets(ctx, userID, gameID, page, pageSize)
}

// GetBonusBalance gets total bonus balance
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

// GetActiveBonuses gets active bonuses for a user
func (s *WalletService) GetActiveBonuses(ctx context.Context, userID, currency string) ([]*model.BonusTransaction, error) {
	return s.repo.GetActiveBonusTransactions(ctx, userID, currency)
}

// invalidateBalanceCache clears balance cache for a user
func (s *WalletService) invalidateBalanceCache(ctx context.Context, userID, currency string) {
	// Clear specific balance cache
	s.redis.Del(ctx, fmt.Sprintf("balance:%s:%s:%s", userID, currency, BalanceTypeReal))
	s.redis.Del(ctx, fmt.Sprintf("balance:%s:%s:%s", userID, currency, BalanceTypeBonus))
	s.redis.Del(ctx, fmt.Sprintf("balance:%s:%s:%s", userID, currency, BalanceTypePromo))

	// Clear all balances cache
	s.redis.Del(ctx, fmt.Sprintf("balances:%s", userID))
}

// publishEvent publishes a NATS event
func (s *WalletService) publishEvent(ctx context.Context, eventType string, data map[string]interface{}) {
	if s.nats == nil {
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	s.nats.Publish(eventType, jsonData)
}

// ConvertAmountToInt64 converts protobuf amount to int64
func ConvertAmountToInt64(amount *float64) int64 {
	if amount == nil {
		return 0
	}
	return int64(*amount * 100) // Convert to cents
}

// FormatAmount formats int64 amount to display string
func FormatAmount(amount int64) string {
	return strconv.FormatInt(amount, 10)
}

// GetCurrencyFromProto converts proto currency to string
func GetCurrencyFromProto(currency int32) string {
	return fmt.Sprintf("CURRENCY_%d", currency)
}
