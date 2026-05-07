package service

import (
	"testing"

	"github.com/game_engine/game-engine/internal/game"
	"github.com/game_engine/game-engine/internal/rng"
)

func TestStateMachineTransitions(t *testing.T) {
	sm := game.NewStateMachine()

	tests := []struct {
		name string
		from string
		to   string
		want bool
	}{
		{"init to betting", game.PhaseInit, game.PhaseBetting, true},
		{"betting to playing", game.PhaseBetting, game.PhasePlaying, true},
		{"playing to settling", game.PhasePlaying, game.PhaseSettling, true},
		{"settling to complete", game.PhaseSettling, game.PhaseComplete, true},
		{"complete to init", game.PhaseComplete, game.PhaseInit, true},
		{"init to playing", game.PhaseInit, game.PhasePlaying, false},
		{"betting to settling", game.PhaseBetting, game.PhaseSettling, false},
		{"complete to playing", game.PhaseComplete, game.PhasePlaying, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sm.CanTransition(tt.from, tt.to)
			if got != tt.want {
				t.Fatalf("CanTransition(%s, %s) = %v, want %v", tt.from, tt.to, got, tt.want)
			}
		})
	}
}

func TestStateMachineTransition(t *testing.T) {
	sm := game.NewStateMachine()
	state := &game.GameState{Phase: game.PhaseInit}

	if err := sm.Transition(state, game.PhaseBetting); err != nil {
		t.Fatalf("Transition(init -> betting) error = %v", err)
	}
	if state.Phase != game.PhaseBetting {
		t.Fatalf("phase = %s, want %s", state.Phase, game.PhaseBetting)
	}

	if err := sm.Transition(state, game.PhaseInit); err == nil {
		t.Fatal("expected error for invalid transition playing -> init")
	}
}

func TestNewGame(t *testing.T) {
	tests := []struct {
		name     string
		gameType string
		playerID string
	}{
		{"slot game", "slot", "player1"},
		{"roulette game", "roulette", "player2"},
		{"dice game", "dice", "player3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := game.NewGame(tt.gameType, tt.playerID, nil)
			if g == nil {
				t.Fatal("NewGame() returned nil")
			}
			if g.GameType != tt.gameType {
				t.Fatalf("GameType = %s, want %s", g.GameType, tt.gameType)
			}
			if g.PlayerID != tt.playerID {
				t.Fatalf("PlayerID = %s, want %s", g.PlayerID, tt.playerID)
			}
			if g.Phase != game.PhaseBetting {
				t.Fatalf("Phase = %s, want %s", g.Phase, game.PhaseBetting)
			}
		})
	}
}

func TestPlaceBet(t *testing.T) {
	g := game.NewGame("slot", "player1", nil)

	tests := []struct {
		name    string
		amount  int64
		wantErr bool
	}{
		{"valid bet", 100, false},
		{"second bet", 50, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := g.PlaceBet(tt.amount)
			if (err != nil) != tt.wantErr {
				t.Fatalf("PlaceBet(%d) error = %v, wantErr %v", tt.amount, err, tt.wantErr)
			}
		})
	}

	if g.TotalBet != 150 {
		t.Fatalf("TotalBet = %d, want 150", g.TotalBet)
	}
}

func TestPlaceBetAfterBettingPhase(t *testing.T) {
	sm := game.NewStateMachine()
	g := game.NewGame("slot", "player1", nil)
	_ = sm.Transition(g, game.PhasePlaying)

	err := g.PlaceBet(100)
	if err == nil {
		t.Fatal("expected error placing bet after betting phase")
	}
}

func TestCalculateRake(t *testing.T) {
	tests := []struct {
		name   string
		cfg    *game.RakeConfig
		netWin int64
		want   int64
	}{
		{"no rake", &game.RakeConfig{Type: game.RakeTypeNone}, 1000, 0},
		{"nil config", nil, 1000, 0},
		{"fixed rake", &game.RakeConfig{Type: game.RakeTypeFixed, Fixed: 50}, 1000, 50},
		{"percentage rake", &game.RakeConfig{Type: game.RakeTypePercentage, Percent: 0.05}, 1000, 50},
		{"hybrid rake within bounds", &game.RakeConfig{Type: game.RakeTypeHybrid, Percent: 0.05, MinCap: 10, MaxCap: 100}, 1000, 50},
		{"hybrid rake min cap", &game.RakeConfig{Type: game.RakeTypeHybrid, Percent: 0.05, MinCap: 100, MaxCap: 500}, 100, 100},
		{"hybrid rake max cap", &game.RakeConfig{Type: game.RakeTypeHybrid, Percent: 0.05, MinCap: 10, MaxCap: 25}, 1000, 25},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := game.CalculateRake(tt.cfg, tt.netWin)
			if got != tt.want {
				t.Fatalf("CalculateRake() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestProvablyFairGeneration(t *testing.T) {
	pf, err := rng.NewProvablyFair()
	if err != nil {
		t.Fatalf("NewProvablyFair() error = %v", err)
	}

	pf.SetClientSeed("test-client-seed-1234567890abcdef")

	hash := pf.GetServerSeedHash()
	if hash == "" {
		t.Fatal("server seed hash should not be empty")
	}

	tests := []struct {
		name string
		max  int
	}{
		{"small range", 10},
		{"medium range", 100},
		{"large range", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := pf.GenerateInt(tt.max)
			if err != nil {
				t.Fatalf("GenerateInt(%d) error = %v", tt.max, err)
			}
			if val < 0 || val >= tt.max {
				t.Fatalf("GenerateInt(%d) = %d, want [0, %d)", tt.max, val, tt.max)
			}
		})
	}
}

func TestProvablyFairDice(t *testing.T) {
	pf, err := rng.NewProvablyFair()
	if err != nil {
		t.Fatalf("NewProvablyFair() error = %v", err)
	}
	pf.SetClientSeed("dice-test-seed")

	rolls, err := pf.GenerateDice(5)
	if err != nil {
		t.Fatalf("GenerateDice() error = %v", err)
	}
	if len(rolls) != 5 {
		t.Fatalf("got %d rolls, want 5", len(rolls))
	}
	for i, r := range rolls {
		if r < 1 || r > 6 {
			t.Fatalf("roll[%d] = %d, want [1,6]", i, r)
		}
	}
}
