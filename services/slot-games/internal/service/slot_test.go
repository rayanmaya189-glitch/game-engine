package service

import (
	"testing"
)

func TestNewGameService(t *testing.T) {
	svc, err := NewGameService()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if svc == nil {
		t.Fatal("expected non-nil service")
	}
	games := svc.ListGames()
	if len(games) != 0 {
		t.Errorf("expected 0 games, got %d", len(games))
	}
}

func TestCreateGame(t *testing.T) {
	tests := []struct {
		name      string
		gameType  string
		id        string
		wantErr   bool
		errSubstr string
	}{
		{"classic game", "classic", "game-1", false, ""},
		{"video game", "video", "game-2", false, ""},
		{"invalid type", "poker", "game-3", true, "invalid game type"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _ := NewGameService()
			game, err := svc.CreateGame(tt.gameType, tt.id)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tt.errSubstr != "" && err.Error() != tt.errSubstr {
					t.Errorf("expected error %q, got %q", tt.errSubstr, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if game.ID != tt.id {
				t.Errorf("expected game ID %s, got %s", tt.id, game.ID)
			}

			got, err := svc.GetGame(tt.id)
			if err != nil {
				t.Fatalf("GetGame failed: %v", err)
			}
			if got.ID != tt.id {
				t.Errorf("GetGame returned wrong game: %s", got.ID)
			}
		})
	}
}

func TestSpin(t *testing.T) {
	tests := []struct {
		name     string
		gameType string
		lineBet  int64
		lines    int
		wantErr  bool
	}{
		{"classic valid spin", "classic", 1, 1, false},
		{"video valid spin", "video", 2, 10, false},
		{"bet below minimum", "classic", 0, 1, true},
		{"lines exceed max", "classic", 1, 999, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _ := NewGameService()
			svc.CreateGame(tt.gameType, "spin-test")

			state, err := svc.Spin("spin-test", tt.lineBet, tt.lines)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if state == nil {
				t.Fatal("expected non-nil state")
			}
			if !state.IsComplete {
				t.Error("expected game state to be complete after spin")
			}
			if state.Bet != tt.lineBet*int64(tt.lines) {
				t.Errorf("expected bet %d, got %d", tt.lineBet*int64(tt.lines), state.Bet)
			}
		})
	}
}

func TestSpinNotFound(t *testing.T) {
	svc, _ := NewGameService()
	_, err := svc.Spin("nonexistent", 1, 1)
	if err == nil {
		t.Fatal("expected error for missing game")
	}
}

func TestDeleteGame(t *testing.T) {
	svc, _ := NewGameService()
	svc.CreateGame("classic", "del-1")

	err := svc.DeleteGame("del-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = svc.GetGame("del-1")
	if err == nil {
		t.Error("expected error after deletion")
	}
}

func TestDeleteGameNotFound(t *testing.T) {
	svc, _ := NewGameService()
	err := svc.DeleteGame("nonexistent")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestListGames(t *testing.T) {
	svc, _ := NewGameService()
	svc.CreateGame("classic", "a")
	svc.CreateGame("video", "b")

	games := svc.ListGames()
	if len(games) != 2 {
		t.Errorf("expected 2 games, got %d", len(games))
	}
}

func TestProvablyFairSpin(t *testing.T) {
	svc, _ := NewGameService()
	svc.CreateGame("classic", "pf-1")

	state, err := svc.ProvablyFairSpin("pf-1", "server-seed", "client-seed", 1, 1, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !state.ProvablyFair {
		t.Error("expected provably fair flag to be true")
	}
}

func TestProvablyFairSpinNotFound(t *testing.T) {
	svc, _ := NewGameService()
	_, err := svc.ProvablyFairSpin("missing", "s", "c", 0, 1, 1)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestProvablyFairDeterministic(t *testing.T) {
	svc1, _ := NewGameService()
	svc1.CreateGame("classic", "det-1")
	svc2, _ := NewGameService()
	svc2.CreateGame("classic", "det-1")

	state1, _ := svc1.ProvablyFairSpin("det-1", "seed", "client", 0, 1, 1)
	state2, _ := svc2.ProvablyFairSpin("det-1", "seed", "client", 0, 1, 1)

	if state1.Win != state2.Win {
		t.Errorf("provably fair results differ: %d vs %d", state1.Win, state2.Win)
	}
}
