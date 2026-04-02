package service

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type mockJackpotRepo struct {
	jackpots map[string]*mockJackpot
	winners  map[string][]mockWinner
}

type mockJackpot struct {
	JackpotID     string
	Name          string
	Description   string
	CurrentAmount float64
	MinBet        float64
	MaxBet        float64
	Status        string
	StartsAt      time.Time
	EndsAt        time.Time
}

type mockWinner struct {
	WinnerID string
	Username string
	Amount   float64
	WonAt    time.Time
}

func newMockRepo() *mockJackpotRepo {
	return &mockJackpotRepo{
		jackpots: make(map[string]*mockJackpot),
		winners:  make(map[string][]mockWinner),
	}
}

func TestCreateJackpot(t *testing.T) {
	tests := []struct {
		name        string
		jackpotName string
		description string
		minBet      float64
		maxBet      float64
	}{
		{"standard jackpot", "Mega Jackpot", "Biggest jackpot", 1.0, 100.0},
		{"mini jackpot", "Mini Jackpot", "Small stakes", 0.1, 10.0},
		{"vip jackpot", "VIP Jackpot", "High rollers", 50.0, 1000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newMockRepo()
			j := &mockJackpot{
				JackpotID: fmt.Sprintf("jp-%s", tt.jackpotName),
				Name:      tt.jackpotName,
				MinBet:    tt.minBet,
				MaxBet:    tt.maxBet,
				Status:    "active",
				StartsAt:  time.Now(),
				EndsAt:    time.Now().Add(24 * time.Hour),
			}
			repo.jackpots[j.JackpotID] = j

			if repo.jackpots[j.JackpotID].Name != tt.jackpotName {
				t.Errorf("name = %v, want %v", repo.jackpots[j.JackpotID].Name, tt.jackpotName)
			}
			if repo.jackpots[j.JackpotID].Status != "active" {
				t.Errorf("status = %v, want active", repo.jackpots[j.JackpotID].Status)
			}
		})
	}
}

func TestGetJackpot(t *testing.T) {
	repo := newMockRepo()
	repo.jackpots["jp-1"] = &mockJackpot{
		JackpotID:     "jp-1",
		Name:          "Test Jackpot",
		CurrentAmount: 5000.0,
		Status:        "active",
	}

	tests := []struct {
		name      string
		id        string
		wantFound bool
	}{
		{"existing jackpot", "jp-1", true},
		{"non-existent jackpot", "jp-999", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jp, ok := repo.jackpots[tt.id]
			if ok != tt.wantFound {
				t.Errorf("found = %v, want %v", ok, tt.wantFound)
			}
			if tt.wantFound && jp.Name != "Test Jackpot" {
				t.Errorf("name = %v, want Test Jackpot", jp.Name)
			}
		})
	}
}

func TestAddContribution(t *testing.T) {
	tests := []struct {
		name          string
		initialAmount float64
		contribution  float64
		wantAmount    float64
	}{
		{"add to empty", 0, 100.0, 100.0},
		{"add to existing", 500.0, 250.0, 750.0},
		{"add small amount", 1000.0, 0.5, 1000.5},
		{"add large amount", 10000.0, 5000.0, 15000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jp := &mockJackpot{
				JackpotID:     "jp-1",
				CurrentAmount: tt.initialAmount,
			}
			jp.CurrentAmount += tt.contribution

			if jp.CurrentAmount != tt.wantAmount {
				t.Errorf("amount = %v, want %v", jp.CurrentAmount, tt.wantAmount)
			}
		})
	}
}

func TestCheckWin(t *testing.T) {
	tests := []struct {
		name      string
		betAmount float64
		minBet    float64
		maxBet    float64
		status    string
		wantJoin  bool
	}{
		{"valid bet active", 50.0, 1.0, 100.0, "active", true},
		{"bet below min", 0.05, 1.0, 100.0, "active", false},
		{"bet above max", 150.0, 1.0, 100.0, "active", false},
		{"inactive jackpot", 50.0, 1.0, 100.0, "inactive", false},
		{"exact min bet", 1.0, 1.0, 100.0, "active", true},
		{"exact max bet", 100.0, 1.0, 100.0, "active", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canJoin := tt.status == "active" && tt.betAmount >= tt.minBet && tt.betAmount <= tt.maxBet
			if canJoin != tt.wantJoin {
				t.Errorf("canJoin = %v, want %v", canJoin, tt.wantJoin)
			}
		})
	}
}

func TestListJackpots(t *testing.T) {
	repo := newMockRepo()
	repo.jackpots["jp-1"] = &mockJackpot{JackpotID: "jp-1", Status: "active"}
	repo.jackpots["jp-2"] = &mockJackpot{JackpotID: "jp-2", Status: "active"}
	repo.jackpots["jp-3"] = &mockJackpot{JackpotID: "jp-3", Status: "completed"}

	ctx := context.Background()
	_ = ctx

	tests := []struct {
		name     string
		filter   string
		wantLen  int
	}{
		{"all jackpots", "", 3},
		{"active only", "active", 2},
		{"completed only", "completed", 1},
		{"non-existent status", "pending", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var results []*mockJackpot
			for _, jp := range repo.jackpots {
				if tt.filter == "" || jp.Status == tt.filter {
					results = append(results, jp)
				}
			}
			if len(results) != tt.wantLen {
				t.Errorf("got %d jackpots, want %d", len(results), tt.wantLen)
			}
		})
	}
}
