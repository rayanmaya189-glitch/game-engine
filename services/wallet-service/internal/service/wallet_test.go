package service

import (
	"context"
	"testing"

	"github.com/game_engine/wallet-service/internal/config"
)

func testConfig() *config.Config {
	return &config.Config{
		Cache:   config.CacheConfig{BalanceTTL: 300, BetLockTTL: 30},
		Betting: config.BettingConfig{MinBetAmount: 100, MaxBetAmount: 100000},
		Deposit: config.DepositConfig{MinDepositAmount: 100, MaxDepositAmount: 1000000},
	}
}

func TestCreateDeposit_Validation(t *testing.T) {
	svc := &WalletService{cfg: testConfig()}

	tests := []struct {
		name    string
		amount  int64
		wantErr bool
	}{
		{"zero", 0, true},
		{"negative", -100, true},
		{"below min", 50, true},
		{"above max", 2000000, true},
		{"valid", 500, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.CreateDeposit(context.Background(), "u1", CurrencyUSD, "card", "stripe", "", tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlaceBet_Validation(t *testing.T) {
	svc := &WalletService{cfg: testConfig()}

	tests := []struct {
		name   string
		amount int64
	}{
		{"zero", 0},
		{"negative", -50},
		{"below min", 50},
		{"above max", 9999999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := svc.PlaceBet(context.Background(), "u1", "g1", "single", "sel", "2.0", CurrencyUSD, tt.amount)
			if err == nil {
				t.Error("expected error")
			}
		})
	}
}

func TestPlaceBet_InsufficientBalance(t *testing.T) {
	svc := &WalletService{cfg: testConfig()}
	bet, w, err := svc.PlaceBet(context.Background(), "u1", "g1", "single", "home", "2.0", CurrencyUSD, 100)
	if err == nil {
		t.Fatal("expected error")
	}
	if bet != nil || w != nil {
		t.Error("expected nil on error")
	}
}

func TestSettleBet_MissingBet(t *testing.T) {
	svc := &WalletService{cfg: testConfig()}
	bet, err := svc.SettleBet(context.Background(), "missing", BetSettlementTypeWon, 2000, "won")
	if err == nil {
		t.Fatal("expected error")
	}
	if bet != nil {
		t.Error("expected nil bet")
	}
}

func TestCancelBet_NotFound(t *testing.T) {
	svc := &WalletService{cfg: testConfig()}
	bet, err := svc.CancelBet(context.Background(), "nonexistent", "reason")
	if err == nil {
		t.Fatal("expected error")
	}
	if bet != nil {
		t.Error("expected nil bet")
	}
}

func TestConvertAmountToInt64(t *testing.T) {
	tests := []struct {
		name     string
		amount   *float64
		expected int64
	}{
		{"nil", nil, 0},
		{"zero", floatPtr(0.0), 0},
		{"10.50", floatPtr(10.50), 1050},
		{"100.99", floatPtr(100.99), 10099},
		{"0.01", floatPtr(0.01), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAmountToInt64(tt.amount); got != tt.expected {
				t.Errorf("got %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestFormatAmount(t *testing.T) {
	tests := []struct {
		amount   int64
		expected string
	}{
		{0, "0"}, {100, "100"}, {-50, "-50"}, {999999, "999999"},
	}
	for _, tt := range tests {
		if got := FormatAmount(tt.amount); got != tt.expected {
			t.Errorf("FormatAmount(%d) = %s, want %s", tt.amount, got, tt.expected)
		}
	}
}

func TestGetCurrencyFromProto(t *testing.T) {
	tests := []struct {
		input    int32
		expected string
	}{
		{0, "CURRENCY_0"}, {1, "CURRENCY_1"}, {99, "CURRENCY_99"},
	}
	for _, tt := range tests {
		if got := GetCurrencyFromProto(tt.input); got != tt.expected {
			t.Errorf("got %s, want %s", got, tt.expected)
		}
	}
}

func TestEventConstants(t *testing.T) {
	events := []string{
		EventDepositCreated, EventDepositCompleted,
		EventWithdrawalCreated, EventWithdrawalCompleted,
		EventBetPlaced, EventBetSettled, EventBetCancelled,
		EventBonusCredited, EventTransactionReversed,
	}
	for _, e := range events {
		if e == "" {
			t.Error("event constant should not be empty")
		}
	}
}

func TestBalanceTypeConstants(t *testing.T) {
	if BalanceTypeReal == "" || BalanceTypeBonus == "" || BalanceTypePromo == "" {
		t.Error("balance type constants should not be empty")
	}
}

func TestTransactionStatusConstants(t *testing.T) {
	statuses := []string{
		TransactionStatusPending, TransactionStatusProcessing,
		TransactionStatusCompleted, TransactionStatusFailed,
		TransactionStatusCancelled, TransactionStatusReversed,
	}
	for _, s := range statuses {
		if s == "" {
			t.Error("transaction status should not be empty")
		}
	}
}

func TestBetSettlementConstants(t *testing.T) {
	types := []string{
		BetSettlementTypePending, BetSettlementTypeWon,
		BetSettlementTypeLost, BetSettlementTypeCancelled,
	}
	for _, s := range types {
		if s == "" {
			t.Error("settlement type should not be empty")
		}
	}
}

func floatPtr(f float64) *float64 { return &f }
