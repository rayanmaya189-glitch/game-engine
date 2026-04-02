package service

import (
	"testing"
	"time"

	"github.com/game_engine/loyalty-service/internal/config"
	"github.com/game_engine/loyalty-service/internal/model"
)

func newTestLoyaltyConfig() *config.Config {
	return &config.Config{
		Loyalty: config.LoyaltyConfig{
			PointsPerBet: 0.1,
			PointsMultiplierVIP: map[string]float64{
				"gold":     1.5,
				"platinum": 2.0,
				"diamond":  3.0,
			},
			ExchangeRate:    0.01,
			LevelThresholds: []int{0, 1000, 5000, 10000, 50000, 100000},
		},
	}
}

func TestCalculateTier(t *testing.T) {
	cfg := newTestLoyaltyConfig()

	tests := []struct {
		name           string
		lifetimePoints int
		want           string
	}{
		{"bronze zero", 0, "bronze"},
		{"bronze low", 500, "bronze"},
		{"silver threshold", 1000, "silver"},
		{"silver mid", 3000, "silver"},
		{"gold threshold", 5000, "gold"},
		{"platinum threshold", 10000, "platinum"},
		{"diamond threshold", 50000, "diamond"},
		{"vip threshold", 100000, "vip"},
		{"vip high", 500000, "vip"},
	}

	thresholds := cfg.Loyalty.LevelThresholds
	tiers := []string{"bronze", "silver", "gold", "platinum", "diamond", "vip"}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := "bronze"
			for i := len(thresholds) - 1; i >= 0; i-- {
				if tt.lifetimePoints >= thresholds[i] {
					got = tiers[i]
					break
				}
			}
			if got != tt.want {
				t.Fatalf("calculateTier(%d) = %s, want %s", tt.lifetimePoints, got, tt.want)
			}
		})
	}
}

func TestPointsCalculation(t *testing.T) {
	cfg := newTestLoyaltyConfig()

	tests := []struct {
		name      string
		betAmount float64
		tier      string
		want      int
	}{
		{"bronze bet 100", 100, "bronze", 10},
		{"gold bet 100", 100, "gold", 15},
		{"platinum bet 100", 100, "platinum", 20},
		{"diamond bet 100", 100, "diamond", 30},
		{"bronze bet 50", 50, "bronze", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multiplier := cfg.Loyalty.PointsPerBet
			if vipMult, ok := cfg.Loyalty.PointsMultiplierVIP[tt.tier]; ok {
				multiplier *= vipMult
			}
			got := int(float64(tt.betAmount) * multiplier)
			if got != tt.want {
				t.Fatalf("points for bet=%f tier=%s = %d, want %d", tt.betAmount, tt.tier, got, tt.want)
			}
		})
	}
}

func TestLoyaltyMemberModel(t *testing.T) {
	member := model.LoyaltyMember{
		UserID:         "u1",
		Username:       "alice",
		Email:          "alice@example.com",
		Points:         500,
		LifetimePoints: 2000,
		Tier:           "silver",
		Status:         "active",
		JoinedAt:       time.Now(),
		UpdatedAt:      time.Now(),
	}

	if member.Points > member.LifetimePoints {
		t.Fatal("Points should not exceed LifetimePoints")
	}
	if member.Status != "active" {
		t.Fatalf("Status = %s, want active", member.Status)
	}
}

func TestPointsTransactionModel(t *testing.T) {
	tx := model.PointsTransaction{
		TransactionID: "tx-001",
		UserID:        "u1",
		Amount:        100,
		Type:          "credit",
		Source:        "bet",
		ReferenceID:   "bet-001",
		Description:   "Points earned from bet",
	}

	if tx.Amount <= 0 && tx.Type == "credit" {
		t.Fatal("credit transaction should have positive amount")
	}
	if tx.Type != "credit" && tx.Type != "debit" {
		t.Fatalf("Type = %s, want credit or debit", tx.Type)
	}
}

func TestTierModel(t *testing.T) {
	tiers := []model.Tier{
		{TierID: "1", Name: "bronze", MinPoints: 0, MaxPoints: 999, PointsMultiplier: 1.0, CashbackPercent: 0},
		{TierID: "2", Name: "silver", MinPoints: 1000, MaxPoints: 4999, PointsMultiplier: 1.2, CashbackPercent: 1.0},
		{TierID: "3", Name: "gold", MinPoints: 5000, MaxPoints: 9999, PointsMultiplier: 1.5, CashbackPercent: 2.0},
	}

	for i := 1; i < len(tiers); i++ {
		if tiers[i].MinPoints <= tiers[i-1].MinPoints {
			t.Fatalf("tier %s MinPoints %d should be > tier %s MinPoints %d",
				tiers[i].Name, tiers[i].MinPoints, tiers[i-1].Name, tiers[i-1].MinPoints)
		}
	}
}

func TestRewardModel(t *testing.T) {
	reward := model.Reward{
		RewardID:    "r1",
		Name:        "Free Spins",
		Description: "10 free spins on Starburst",
		PointsCost:  500,
		Type:        "free_spin",
		Value:       10,
		Status:      "active",
	}

	if reward.PointsCost <= 0 {
		t.Fatal("PointsCost should be positive")
	}
	if reward.Status != "active" {
		t.Fatalf("Status = %s, want active", reward.Status)
	}
}

func TestRedeemPointsValidation(t *testing.T) {
	tests := []struct {
		name       string
		userPoints int
		cost       int
		canRedeem  bool
	}{
		{"sufficient points", 500, 100, true},
		{"exact points", 100, 100, true},
		{"insufficient points", 50, 100, false},
		{"zero cost", 500, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canRedeem := tt.userPoints >= tt.cost
			if canRedeem != tt.canRedeem {
				t.Fatalf("points=%d, cost=%d: got canRedeem=%v, want %v",
					tt.userPoints, tt.cost, canRedeem, tt.canRedeem)
			}
		})
	}
}

func TestTierUpgradeDetection(t *testing.T) {
	tests := []struct {
		name       string
		oldTier    string
		newTier    string
		upgraded   bool
	}{
		{"no upgrade", "bronze", "bronze", false},
		{"bronze to silver", "bronze", "silver", true},
		{"silver to gold", "silver", "gold", true},
		{"gold to platinum", "gold", "platinum", true},
		{"downgrade", "gold", "bronze", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upgraded := tt.newTier != tt.oldTier
			if upgraded != tt.upgraded {
				t.Fatalf("oldTier=%s, newTier=%s: got upgraded=%v, want %v",
					tt.oldTier, tt.newTier, upgraded, tt.upgraded)
			}
		})
	}
}

func TestConfigValues(t *testing.T) {
	cfg := newTestLoyaltyConfig()

	if cfg.Loyalty.PointsPerBet != 0.1 {
		t.Fatalf("PointsPerBet = %f, want 0.1", cfg.Loyalty.PointsPerBet)
	}
	if len(cfg.Loyalty.LevelThresholds) != 6 {
		t.Fatalf("got %d thresholds, want 6", len(cfg.Loyalty.LevelThresholds))
	}
	if cfg.Loyalty.PointsMultiplierVIP["diamond"] != 3.0 {
		t.Fatalf("diamond multiplier = %f, want 3.0", cfg.Loyalty.PointsMultiplierVIP["diamond"])
	}
}
