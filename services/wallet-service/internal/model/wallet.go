package model

import (
	"time"

	commonv1 "github.com/game_engine/common-service/proto/gen/go/common/v1"
	walletsv1 "github.com/game_engine/common-service/proto/gen/go/wallet/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Wallet represents a player's wallet (main or bonus)
type Wallet struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Currency     string    `json:"currency"`
	BalanceType  string    `json:"balance_type"` // REAL, BONUS, PROMOTIONAL
	Amount       int64     `json:"amount"`
	LockedAmount int64     `json:"locked_amount"`
	Version      int       `json:"version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Transaction represents a financial transaction
type Transaction struct {
	TransactionID    string     `json:"transaction_id"`
	UserID           string     `json:"user_id"`
	Type             string     `json:"type"`   // DEPOSIT, WITHDRAWAL, BET, WIN, BONUS, etc.
	Status           string     `json:"status"` // PENDING, COMPLETED, FAILED, etc.
	Currency         string     `json:"currency"`
	Amount           int64      `json:"amount"`
	BonusAmount      int64      `json:"bonus_amount"`
	Fee              int64      `json:"fee"`
	NetAmount        int64      `json:"net_amount"`
	PaymentMethod    string     `json:"payment_method"`
	PaymentProvider  string     `json:"payment_provider"`
	PaymentReference string     `json:"payment_reference"`
	GameID           string     `json:"game_id"`
	BetID            string     `json:"bet_id"`
	Description      string     `json:"description"`
	CreatedAt        time.Time  `json:"created_at"`
	ProcessedAt      *time.Time `json:"processed_at"`
}

// Bet represents a bet placed by a player
type Bet struct {
	BetID          string     `json:"bet_id"`
	UserID         string     `json:"user_id"`
	Currency       string     `json:"currency"`
	GameID         string     `json:"game_id"`
	BetType        string     `json:"bet_type"`
	Selection      string     `json:"selection"`
	Odds           string     `json:"odds"`
	Stake          int64      `json:"stake"`
	PotentialWin   int64      `json:"potential_win"`
	ActualWin      int64      `json:"actual_win"`
	SettlementType string     `json:"settlement_type"` // PENDING, WON, LOST, CANCELLED
	Status         string     `json:"status"`          // PENDING, COMPLETED, CANCELLED
	PlacedAt       time.Time  `json:"placed_at"`
	SettledAt      *time.Time `json:"settled_at"`
}

// BonusTransaction represents bonus credits with wagering requirements
type BonusTransaction struct {
	ID                 string     `json:"id"`
	UserID             string     `json:"user_id"`
	TransactionID      string     `json:"transaction_id"`
	BonusType          string     `json:"bonus_type"` // WELCOME, DEPOSIT, NO_DEPOSIT, etc.
	Currency           string     `json:"currency"`
	Amount             int64      `json:"amount"`
	WageringMultiplier int        `json:"wagering_multiplier"`
	WageringRequired   int64      `json:"wagering_required"`
	WageringMet        int64      `json:"wagering_met"`
	BonusCode          string     `json:"bonus_code"`
	Status             string     `json:"status"` // ACTIVE, USED, EXPIRED
	ExpiresAt          *time.Time `json:"expires_at"`
	CreatedAt          time.Time  `json:"created_at"`
	UsedAt             *time.Time `json:"used_at"`
}

// ToProto converts Wallet to protobuf message
func (w *Wallet) ToProto() *commonv1.Money {
	return &commonv1.Money{
		Amount:   w.Amount,
		Currency: commonv1.Currency(commonv1.Currency_value[w.Currency]),
	}
}

// ToBalanceProto converts Wallet to BalanceEntry protobuf
func (w *Wallet) ToBalanceProto() *commonv1.BalanceEntry {
	return &commonv1.BalanceEntry{
		BalanceType:     commonv1.BalanceType(commonv1.BalanceType_value[w.BalanceType]),
		Amount:          w.ToProto(),
		LockedAmount:    &commonv1.Money{Amount: w.LockedAmount, Currency: commonv1.Currency(commonv1.Currency_value[w.Currency])},
		AvailableAmount: &commonv1.Money{Amount: w.Amount - w.LockedAmount, Currency: commonv1.Currency(commonv1.Currency_value[w.Currency])},
	}
}

// ToTransactionProto converts Transaction to protobuf message
func (t *Transaction) ToTransactionProto() *walletsv1.Transaction {
	protoTx := &walletsv1.Transaction{
		TransactionId:    t.TransactionID,
		UserId:           t.UserID,
		Type:             commonv1.TransactionType(commonv1.TransactionType_value[t.Type]),
		Status:           commonv1.TransactionStatus(commonv1.TransactionStatus_value[t.Status]),
		PaymentMethod:    commonv1.PaymentMethod(commonv1.PaymentMethod_value[t.PaymentMethod]),
		PaymentProvider:  t.PaymentProvider,
		PaymentReference: t.PaymentReference,
		GameId:           t.GameID,
		BetId:            t.BetID,
		Description:      t.Description,
		Amount: &commonv1.TransactionAmount{
			Requested: &commonv1.Money{Amount: t.Amount, Currency: commonv1.Currency(commonv1.Currency_value[t.Currency])},
			Approved:  &commonv1.Money{Amount: t.NetAmount, Currency: commonv1.Currency(commonv1.Currency_value[t.Currency])},
			Bonus:     &commonv1.Money{Amount: t.BonusAmount, Currency: commonv1.Currency(commonv1.Currency_value[t.Currency])},
			Fee:       &commonv1.Money{Amount: t.Fee, Currency: commonv1.Currency(commonv1.Currency_value[t.Currency])},
			Net:       &commonv1.Money{Amount: t.NetAmount, Currency: commonv1.Currency(commonv1.Currency_value[t.Currency])},
		},
		CreatedAt:   timestampToProto(t.CreatedAt),
		ProcessedAt: timestampToProtoPtr(t.ProcessedAt),
	}
	return protoTx
}

// ToBetProto converts Bet to protobuf message
func (b *Bet) ToBetProto() *walletsv1.Bet {
	return &walletsv1.Bet{
		BetId:          b.BetID,
		UserId:         b.UserID,
		GameId:         b.GameID,
		BetType:        b.BetType,
		Selection:      b.Selection,
		Odds:           b.Odds,
		Stake:          &commonv1.Money{Amount: b.Stake},
		PotentialWin:   &commonv1.Money{Amount: b.PotentialWin},
		ActualWin:      &commonv1.Money{Amount: b.ActualWin},
		SettlementType: walletsv1.BetSettlementType(walletsv1.BetSettlementType_value[b.SettlementType]),
		Status:         commonv1.TransactionStatus(commonv1.TransactionStatus_value[b.Status]),
		PlacedAt:       timestampToProto(b.PlacedAt),
	}
}

// ToTransactionProto converts Bet to protobuf Transaction message
func (b *Bet) ToTransactionProto() *walletsv1.Transaction {
	return &walletsv1.Transaction{
		TransactionId: b.BetID,
		UserId:        b.UserID,
		Type:          commonv1.TransactionType_TRANSACTION_TYPE_BET,
		Status:        commonv1.TransactionStatus(commonv1.TransactionStatus_value[b.Status]),
		Amount: &commonv1.TransactionAmount{
			Requested: &commonv1.Money{Amount: b.Stake},
			Approved:  &commonv1.Money{Amount: b.ActualWin},
		},
		GameId:    b.GameID,
		BetId:     b.BetID,
		CreatedAt: timestampToProto(b.PlacedAt),
	}
}

// ToBonusProto converts BonusTransaction to protobuf message
func (bt *BonusTransaction) ToBonusProto() *walletsv1.Transaction {
	return &walletsv1.Transaction{
		TransactionId: bt.TransactionID,
		UserId:        bt.UserID,
		Type:          commonv1.TransactionType_TRANSACTION_TYPE_BONUS,
		Status:        commonv1.TransactionStatus(commonv1.TransactionStatus_value[bt.Status]),
		Amount: &commonv1.TransactionAmount{
			Requested: &commonv1.Money{Amount: bt.Amount, Currency: commonv1.Currency(commonv1.Currency_value[bt.Currency])},
			Bonus:     &commonv1.Money{Amount: bt.Amount, Currency: commonv1.Currency(commonv1.Currency_value[bt.Currency])},
		},
		Description: bt.BonusCode,
	}
}

// Helper function to convert time.Time to protobuf timestamp
func timestampToProto(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

// Helper function to convert *time.Time to protobuf timestamp
func timestampToProtoPtr(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestampToProto(*t)
}

// GetWalletByType returns the wallet for a specific balance type
func GetWalletByType(wallets []*Wallet, balanceType string) *Wallet {
	for _, w := range wallets {
		if w.BalanceType == balanceType {
			return w
		}
	}
	return nil
}

// GetRealWallet returns the real (main) wallet
func GetRealWallet(wallets []*Wallet) *Wallet {
	return GetWalletByType(wallets, "BALANCE_TYPE_REAL")
}

// GetBonusWallet returns the bonus wallet
func GetBonusWallet(wallets []*Wallet) *Wallet {
	return GetWalletByType(wallets, "BALANCE_TYPE_BONUS")
}

// CalculateAvailableBalance calculates available balance (real - locked + bonus)
func CalculateAvailableBalance(real, bonus, locked int64) int64 {
	available := real - locked
	if available < 0 {
		available = 0
	}
	return available + bonus
}
