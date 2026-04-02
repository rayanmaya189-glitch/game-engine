package tournament

import (
	"context"
	"testing"
	"time"
)

type mockRedisClient struct{}

func defaultSettings() TournamentSettings {
	return TournamentSettings{
		AutoStart:         false,
		RebuyEnabled:      false,
		AddonEnabled:      false,
		Reentries:         0,
		LateRegistration:  300,
		StartingChips:     1000,
		BlindLevels:       generateDefaultBlindLevels(),
		PrizeDistribution: []int{50, 30, 20},
	}
}

func newTestManager() *Manager {
	return &Manager{
		tournaments: make(map[string]*Tournament),
	}
}

func TestCreateTournament(t *testing.T) {
	tests := []struct {
		name      string
		tName     string
		tType     TournamentType
		gameType  string
		entryFee  int64
		buyIn     int64
		minP      int
		maxP      int
		startTime time.Time
		settings  TournamentSettings
		wantErr   bool
	}{
		{
			name:      "scheduled tournament",
			tName:     "Weekly Poker",
			tType:     TournamentTypeScheduled,
			gameType:  "poker",
			entryFee:  100,
			buyIn:     100,
			minP:      2,
			maxP:      100,
			startTime: time.Now().Add(1 * time.Hour),
			settings:  defaultSettings(),
			wantErr:   false,
		},
		{
			name:      "sit and go with auto start",
			tName:     "Quick Game",
			tType:     TournamentTypeSitAndGo,
			gameType:  "poker",
			entryFee:  50,
			buyIn:     50,
			minP:      0,
			maxP:      0,
			startTime: time.Time{},
			settings: TournamentSettings{
				AutoStart:         true,
				StartingChips:     1500,
				BlindLevels:       generateDefaultBlindLevels(),
				PrizeDistribution: []int{70, 30},
			},
			wantErr: false,
		},
		{
			name:      "knockout tournament",
			tName:     "Bounty Hunter",
			tType:     TournamentTypeKnockout,
			gameType:  "poker",
			entryFee:  200,
			buyIn:     200,
			minP:      4,
			maxP:      64,
			startTime: time.Now().Add(2 * time.Hour),
			settings:  defaultSettings(),
			wantErr:   false,
		},
		{
			name:      "freeroll tournament",
			tName:     "Free Roll",
			tType:     TournamentTypeFreeroll,
			gameType:  "poker",
			entryFee:  500,
			buyIn:     0,
			minP:      2,
			maxP:      500,
			startTime: time.Now().Add(30 * time.Minute),
			settings:  defaultSettings(),
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newTestManager()
			ctx := context.Background()

			result, err := m.CreateTournament(ctx, tt.tName, tt.tType, tt.gameType,
				tt.entryFee, tt.buyIn, tt.minP, tt.maxP, tt.startTime, tt.settings)

			if (err != nil) != tt.wantErr {
				t.Fatalf("CreateTournament() error = %v, wantErr %v", err, tt.wantErr)
			}

			if result == nil {
				t.Fatal("expected tournament, got nil")
			}

			if result.Name != tt.tName {
				t.Errorf("name = %v, want %v", result.Name, tt.tName)
			}
			if result.Type != tt.tType {
				t.Errorf("type = %v, want %v", result.Type, tt.tType)
			}
			if result.Status != TournamentStatusPending {
				t.Errorf("status = %v, want %v", result.Status, TournamentStatusPending)
			}
			if tt.tType == TournamentTypeFreeroll && result.EntryFee != 0 {
				t.Errorf("freeroll entry fee should be 0, got %d", result.EntryFee)
			}
			if result.MaxPlayers == 0 {
				t.Error("max players should not be 0")
			}
		})
	}
}

func TestGetTournament(t *testing.T) {
	m := newTestManager()
	ctx := context.Background()

	tournament, _ := m.CreateTournament(ctx, "Test", TournamentTypeScheduled, "poker",
		100, 100, 2, 100, time.Now().Add(time.Hour), defaultSettings())

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{"existing tournament", tournament.ID, false},
		{"non-existent tournament", "invalid-id", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := m.GetTournament(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetTournament() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && result.ID != tt.id {
				t.Errorf("id = %v, want %v", result.ID, tt.id)
			}
		})
	}
}

func TestRegisterUser(t *testing.T) {
	m := newTestManager()
	ctx := context.Background()

	tournament, _ := m.CreateTournament(ctx, "Test", TournamentTypeScheduled, "poker",
		100, 100, 2, 3, time.Now().Add(time.Hour), defaultSettings())

	tests := []struct {
		name     string
		userID   string
		username string
		wantErr  bool
	}{
		{"first user", "user-1", "player1", false},
		{"second user", "user-2", "player2", false},
		{"duplicate registration", "user-1", "player1", true},
		{"invalid tournament", "user-3", "player3", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := m.RegisterUser(ctx, tournament.ID, tt.userID, tt.username)
			if (err != nil) != tt.wantErr {
				t.Fatalf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	updated, _ := m.GetTournament(ctx, tournament.ID)
	if updated.CurrentPlayers != 2 {
		t.Errorf("current players = %d, want 2", updated.CurrentPlayers)
	}
}

func TestStartTournament(t *testing.T) {
	tests := []struct {
		name      string
		minP      int
		registerN int
		wantErr   bool
	}{
		{"enough players", 2, 3, false},
		{"not enough players", 5, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newTestManager()
			ctx := context.Background()

			tournament, _ := m.CreateTournament(ctx, "Test", TournamentTypeScheduled, "poker",
				100, 100, tt.minP, 100, time.Now().Add(time.Hour), defaultSettings())

			for i := 0; i < tt.registerN; i++ {
				m.RegisterUser(ctx, tournament.ID, string(rune('a'+i)), "player")
			}

			err := m.StartTournament(ctx, tournament.ID)
			if (err != nil) != tt.wantErr {
				t.Fatalf("StartTournament() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				updated, _ := m.GetTournament(ctx, tournament.ID)
				if updated.Status != TournamentStatusRunning {
					t.Errorf("status = %v, want %v", updated.Status, TournamentStatusRunning)
				}
			}
		})
	}
}

func TestUpdatePlayerScore(t *testing.T) {
	m := newTestManager()
	ctx := context.Background()
	tournament, _ := m.CreateTournament(ctx, "Test", TournamentTypeScheduled, "poker",
		100, 100, 2, 10, time.Now().Add(time.Hour), defaultSettings())
	m.RegisterUser(ctx, tournament.ID, "user-1", "player1")
	m.RegisterUser(ctx, tournament.ID, "user-2", "player2")
	m.StartTournament(ctx, tournament.ID)
	tests := []struct {
		name      string
		userID    string
		delta     int
		wantErr   bool
		wantScore int
	}{
		{"positive score", "user-1", 500, false, 500},
		{"negative score", "user-1", -100, false, 400},
		{"non-existent user", "user-99", 100, true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := m.UpdatePlayerScore(ctx, tournament.ID, tt.userID, tt.delta)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UpdatePlayerScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestListTournaments(t *testing.T) {
	m := newTestManager()
	ctx := context.Background()
	m.CreateTournament(ctx, "T1", TournamentTypeScheduled, "poker", 100, 100, 2, 100, time.Now().Add(time.Hour), defaultSettings())
	m.CreateTournament(ctx, "T2", TournamentTypeSitAndGo, "poker", 50, 50, 2, 10, time.Time{}, defaultSettings())
	m.CreateTournament(ctx, "T3", TournamentTypeScheduled, "blackjack", 200, 200, 2, 50, time.Now().Add(2*time.Hour), defaultSettings())

	tests := []struct {
		name    string
		status  TournamentStatus
		tType   TournamentType
		wantLen int
	}{
		{"all tournaments", "", "", 3},
		{"scheduled only", "", TournamentTypeScheduled, 2},
		{"sit and go only", "", TournamentTypeSitAndGo, 1},
		{"pending status", TournamentStatusPending, "", 3},
		{"running status", TournamentStatusRunning, "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := m.ListTournaments(ctx, tt.status, tt.tType)
			if err != nil {
				t.Fatalf("ListTournaments() error = %v", err)
			}
			if len(result) != tt.wantLen {
				t.Errorf("got %d tournaments, want %d", len(result), tt.wantLen)
			}
		})
	}
}
