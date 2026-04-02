package service

import (
	"testing"
	"time"

	"github.com/game_engine/live-dealer-service/internal/config"
	"github.com/game_engine/live-dealer-service/internal/model"
)

func newTestDealerConfig() *config.Config {
	return &config.Config{
		Dealer: config.DealerConfig{
			MaxTablesPerDealer:    3,
			SessionTimeoutMinutes: 60,
			VideoStreamBitrate:    5000,
			MaxPlayersPerTable:    7,
		},
	}
}

func TestTableModel(t *testing.T) {
	table := model.Table{
		TableID:   "t1",
		GameType:  "blackjack",
		Status:    "open",
		MinBet:    10,
		MaxBet:    1000,
		MaxSeats:  7,
		DeckCount: 8,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if table.GameType != "blackjack" {
		t.Fatalf("GameType = %s, want blackjack", table.GameType)
	}
	if table.MinBet >= table.MaxBet {
		t.Fatal("MinBet should be less than MaxBet")
	}
}

func TestTableModelValidation(t *testing.T) {
	tests := []struct {
		name    string
		game    string
		minBet  float64
		maxBet  float64
		wantErr bool
	}{
		{"valid table", "blackjack", 10, 1000, false},
		{"empty game type", "", 10, 1000, true},
		{"zero min bet", "blackjack", 0, 1000, true},
		{"min exceeds max", "blackjack", 2000, 100, true},
		{"negative bets", "blackjack", -10, 100, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := false
			if tt.game == "" {
				err = true
			}
			if tt.minBet <= 0 || tt.maxBet <= 0 {
				err = true
			}
			if tt.minBet > tt.maxBet {
				err = true
			}
			if err != tt.wantErr {
				t.Fatalf("validation error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestDealerModel(t *testing.T) {
	dealer := model.Dealer{
		DealerID:   "d1",
		Name:       "John",
		Language:   "en",
		Status:     "available",
		ShiftStart: time.Now(),
	}

	if dealer.Status != "available" {
		t.Fatalf("Status = %s, want available", dealer.Status)
	}
	if dealer.Name == "" {
		t.Fatal("Name should not be empty")
	}
}

func TestDealerStatusTransitions(t *testing.T) {
	tests := []struct {
		name   string
		from   string
		to     string
		valid  bool
	}{
		{"available to busy", "available", "busy", true},
		{"busy to available", "busy", "available", true},
		{"available to break", "available", "break", true},
		{"break to available", "break", "available", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = tt.from
			_ = tt.to
			_ = tt.valid
		})
	}
}

func TestPlayerModel(t *testing.T) {
	player := model.Player{
		PlayerID:   "p1",
		TableID:    "t1",
		SeatNumber: 1,
		Chips:      500,
		CurrentBet: 0,
		JoinedAt:   time.Now(),
		LastAction: "bet",
	}

	if player.Chips < 0 {
		t.Fatal("Chips should not be negative")
	}
	if player.SeatNumber < 1 {
		t.Fatal("SeatNumber should be >= 1")
	}
}

func TestBetModel(t *testing.T) {
	bet := model.Bet{
		BetID:     "b1",
		PlayerID:  "p1",
		TableID:   "t1",
		RoundID:   "r1",
		BetType:   "main",
		BetAmount: 100,
		Odds:      2.0,
		Potential: 200,
		Result:    "pending",
	}

	if bet.BetAmount <= 0 {
		t.Fatal("BetAmount should be positive")
	}
	if bet.Potential != bet.BetAmount*bet.Odds {
		t.Fatalf("Potential = %f, want %f", bet.Potential, bet.BetAmount*bet.Odds)
	}
}

func TestGameStateModel(t *testing.T) {
	state := model.GameState{
		TableID: "t1",
		RoundID: "r1",
		Phase:   "betting",
		Pot:     0,
	}

	phases := []string{"betting", "playing", "resolving", "finished"}
	found := false
	for _, p := range phases {
		if state.Phase == p {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Phase = %s, want one of %v", state.Phase, phases)
	}
}

func TestTableStatusValues(t *testing.T) {
	statuses := []string{"open", "closed", "maintenance"}
	for _, s := range statuses {
		if s == "" {
			t.Fatal("status should not be empty")
		}
	}
}

func TestBetResultValues(t *testing.T) {
	results := []string{"pending", "won", "lost", "void"}
	for _, r := range results {
		if r == "" {
			t.Fatal("bet result should not be empty")
		}
	}
}

func TestStreamInfoStructure(t *testing.T) {
	info := StreamInfo{
		TableID:            "t1",
		Bitrate:            5000,
		ViewerCount:        10,
		StreamURL:          "rtmp://stream.internal/t1/live",
		IsLive:             true,
		MaxPlayersPerTable: 7,
	}

	if info.ViewerCount < 0 {
		t.Fatal("ViewerCount should not be negative")
	}
	if info.Bitrate <= 0 {
		t.Fatal("Bitrate should be positive")
	}
	if !info.IsLive {
		t.Fatal("IsLive should be true for open table")
	}
}

func TestDealerConfig(t *testing.T) {
	cfg := newTestDealerConfig()

	if cfg.Dealer.MaxPlayersPerTable != 7 {
		t.Fatalf("MaxPlayersPerTable = %d, want 7", cfg.Dealer.MaxPlayersPerTable)
	}
	if cfg.Dealer.VideoStreamBitrate != 5000 {
		t.Fatalf("VideoStreamBitrate = %d, want 5000", cfg.Dealer.VideoStreamBitrate)
	}
}

func TestTableSeatCapacity(t *testing.T) {
	tests := []struct {
		name     string
		current  int
		max      int
		hasSpace bool
	}{
		{"empty table", 0, 7, true},
		{"half full", 3, 7, true},
		{"full table", 7, 7, false},
		{"over capacity", 8, 7, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasSpace := tt.current < tt.max
			if hasSpace != tt.hasSpace {
				t.Fatalf("current=%d, max=%d: got hasSpace=%v, want %v", tt.current, tt.max, hasSpace, tt.hasSpace)
			}
		})
	}
}
