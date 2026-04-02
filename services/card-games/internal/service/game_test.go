package service

import (
	"testing"

	"github.com/game_engine/card-games/internal/config"
)

func newTestConfig() *config.Config {
	return &config.Config{
		Game: config.GameConfig{
			DefaultDeckCount: 6,
			MaxDeckCount:     8,
			Blackjack: config.BlackjackConfig{
				AllowSurrender:       false,
				AllowLateSurrender:   false,
				DealerStandsOnSoft17: true,
				MaxSplits:            3,
				BlackjackPayout:      1.5,
			},
			Baccarat: config.BaccaratConfig{
				Commission: 0.05,
				MaxPlayers: 7,
				ShoeSize:   8,
			},
			Poker: config.PokerConfig{
				MinPlayers: 2,
				MaxPlayers: 9,
			},
			AndarBahar: config.AndarBaharConfig{
				SideToDealFirst:   "andar",
				MaxCommunityCards: 25,
			},
			TeenPatti: config.TeenPattiConfig{
				Variant:    "classic",
				MinPlayers: 2,
				MaxPlayers: 6,
			},
		},
	}
}

func TestNewGameService(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *config.Config
		wantErr bool
	}{
		{
			name:    "valid config",
			cfg:     newTestConfig(),
			wantErr: false,
		},
		{
			name:    "nil config",
			cfg:     nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, err := NewGameService(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewGameService() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && svc == nil {
				t.Fatal("NewGameService() returned nil service")
			}
		})
	}
}

func TestCreateBlackjack(t *testing.T) {
	tests := []struct {
		name    string
		gameID  string
		player  string
		bet     int64
		wantErr bool
	}{
		{"valid game", "bj-001", "player1", 100, false},
		{"zero bet", "bj-002", "player1", 0, true},
	}

	cfg := newTestConfig()
	svc, err := NewGameService(cfg)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := svc.CreateBlackjack(tt.gameID, tt.player, tt.bet)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CreateBlackjack() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && game == nil {
				t.Fatal("CreateBlackjack() returned nil game")
			}
		})
	}
}

func TestCreateBaccarat(t *testing.T) {
	cfg := newTestConfig()
	svc, err := NewGameService(cfg)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	game, err := svc.CreateBaccarat("bac-001", nil)
	if err != nil {
		t.Fatalf("CreateBaccarat() error = %v", err)
	}
	if game == nil {
		t.Fatal("CreateBaccarat() returned nil")
	}
}

func TestCreatePoker(t *testing.T) {
	cfg := newTestConfig()
	svc, err := NewGameService(cfg)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	game, err := svc.CreatePoker("poker-001", "texas_holdem")
	if err != nil {
		t.Fatalf("CreatePoker() error = %v", err)
	}
	if game == nil {
		t.Fatal("CreatePoker() returned nil")
	}
}

func TestCreateAndarBahar(t *testing.T) {
	cfg := newTestConfig()
	svc, err := NewGameService(cfg)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	game, err := svc.CreateAndarBahar("ab-001")
	if err != nil {
		t.Fatalf("CreateAndarBahar() error = %v", err)
	}
	if game == nil {
		t.Fatal("CreateAndarBahar() returned nil")
	}
}

func TestCreateTeenPatti(t *testing.T) {
	cfg := newTestConfig()
	svc, err := NewGameService(cfg)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	game, err := svc.CreateTeenPatti("tp-001")
	if err != nil {
		t.Fatalf("CreateTeenPatti() error = %v", err)
	}
	if game == nil {
		t.Fatal("CreateTeenPatti() returned nil")
	}
}

func TestGetGame(t *testing.T) {
	cfg := newTestConfig()
	svc, err := NewGameService(cfg)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	tests := []struct {
		name   string
		gameID string
		found  bool
	}{
		{"existing game", "bj-get-001", true},
		{"non-existent game", "nonexistent", false},
	}

	_, _ = svc.CreateBlackjack("bj-get-001", "player1", 50)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := svc.GetGame(tt.gameID)
			if err != nil {
				t.Fatalf("GetGame() error = %v", err)
			}
			if tt.found && game == nil {
				t.Fatal("GetGame() returned nil for existing game")
			}
			if !tt.found && game != nil {
				t.Fatal("GetGame() returned non-nil for non-existent game")
			}
		})
	}
}

func TestRemoveGame(t *testing.T) {
	cfg := newTestConfig()
	svc, err := NewGameService(cfg)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	_, _ = svc.CreateBlackjack("bj-rm-001", "player1", 50)
	svc.RemoveGame("bj-rm-001")

	game, err := svc.GetGame("bj-rm-001")
	if err != nil {
		t.Fatalf("GetGame() error = %v", err)
	}
	if game != nil {
		t.Fatal("GetGame() returned non-nil after RemoveGame()")
	}
}
