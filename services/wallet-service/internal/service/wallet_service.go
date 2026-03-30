package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/game_engine/wallet-service/internal/config"
	"github.com/game_engine/wallet-service/internal/repository"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

type WalletService struct {
	repo  *repository.WalletRepository
	redis *redis.Client
	nats  *nats.Conn
	cfg   *config.Config
}

func NewWalletService(repo *repository.WalletRepository, redis *redis.Client, natsConn *nats.Conn, cfg *config.Config) (*WalletService, error) {
	return &WalletService{
		repo:  repo,
		redis: redis,
		nats:  natsConn,
		cfg:   cfg,
	}, nil
}

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

func (s *WalletService) invalidateBalanceCache(ctx context.Context, userID, currency string) {
	s.redis.Del(ctx, fmt.Sprintf("balance:%s:%s:%s", userID, currency, BalanceTypeReal))
	s.redis.Del(ctx, fmt.Sprintf("balance:%s:%s:%s", userID, currency, BalanceTypeBonus))
	s.redis.Del(ctx, fmt.Sprintf("balance:%s:%s:%s", userID, currency, BalanceTypePromo))
	s.redis.Del(ctx, fmt.Sprintf("balances:%s", userID))
}

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

func ConvertAmountToInt64(amount *float64) int64 {
	if amount == nil {
		return 0
	}
	return int64(*amount * 100)
}

func FormatAmount(amount int64) string {
	return strconv.FormatInt(amount, 10)
}

func GetCurrencyFromProto(currency int32) string {
	return fmt.Sprintf("CURRENCY_%d", currency)
}
