package service

import (
	"testing"
	"time"

	"github.com/game_engine/leaderboard-service/internal/config"
	"github.com/game_engine/leaderboard-service/internal/model"
)

func newTestLeaderboardConfig() *config.LeaderboardConfig {
	return &config.LeaderboardConfig{
		TopPlayersCount:    10,
		CacheTTLSeconds:    300,
		ResetIntervalHours: 24,
		MinBetThreshold:    0.50,
		PrizeAutoCredit:    false,
	}
}

func TestUpdateScoreRequestValidation(t *testing.T) {
	tests := []struct {
		name    string
		req     model.UpdateScoreRequest
		minBet  float64
		wantErr bool
	}{
		{
			"valid request",
			model.UpdateScoreRequest{UserID: "u1", Score: 100, BetAmount: 10},
			0.50,
			false,
		},
		{
			"score exceeds max",
			model.UpdateScoreRequest{UserID: "u1", Score: 2000000, BetAmount: 10},
			0.50,
			true,
		},
		{
			"negative bet amount",
			model.UpdateScoreRequest{UserID: "u1", Score: 100, BetAmount: -5},
			0.50,
			true,
		},
		{
			"bet below threshold",
			model.UpdateScoreRequest{UserID: "u1", Score: 100, BetAmount: 0.10},
			0.50,
			true,
		},
		{
			"zero bet with threshold",
			model.UpdateScoreRequest{UserID: "u1", Score: 100, BetAmount: 0},
			0.50,
			false,
		},
		{
			"no threshold",
			model.UpdateScoreRequest{UserID: "u1", Score: 100, BetAmount: 0.01},
			0,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate(tt.minBet)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLeaderboardTypes(t *testing.T) {
	types := []model.LeaderboardType{
		model.LeaderboardTypeDaily,
		model.LeaderboardTypeWeekly,
		model.LeaderboardTypeMonthly,
		model.LeaderboardTypeAllTime,
		model.LeaderboardTypeBiggestWin,
		model.LeaderboardTypeMostActive,
	}

	for _, lt := range types {
		if string(lt) == "" {
			t.Fatal("leaderboard type should not be empty")
		}
	}
}

func TestLeaderboardResponseStructure(t *testing.T) {
	resp := model.LeaderboardResponse{
		Type:   model.LeaderboardTypeDaily,
		Period: "2024-01-15",
		Entries: []model.LeaderboardEntry{
			{Rank: 1, UserID: "u1", Username: "alice", Score: 500, UpdatedAt: time.Now()},
			{Rank: 2, UserID: "u2", Username: "bob", Score: 300, UpdatedAt: time.Now()},
		},
		Total:     2,
		UpdatedAt: time.Now(),
	}

	if resp.Type != model.LeaderboardTypeDaily {
		t.Fatalf("Type = %s, want daily", resp.Type)
	}
	if len(resp.Entries) != 2 {
		t.Fatalf("got %d entries, want 2", len(resp.Entries))
	}
	if resp.Entries[0].Rank != 1 {
		t.Fatalf("first entry rank = %d, want 1", resp.Entries[0].Rank)
	}
}

func TestPlayerRankResponseStructure(t *testing.T) {
	rank := model.PlayerRankResponse{
		UserID:          "u1",
		Username:        "alice",
		Rank:            5,
		Score:           200,
		LeaderboardType: model.LeaderboardTypeWeekly,
		Period:          "2024-W03",
	}

	if rank.Rank != 5 {
		t.Fatalf("Rank = %d, want 5", rank.Rank)
	}
	if rank.LeaderboardType != model.LeaderboardTypeWeekly {
		t.Fatalf("LeaderboardType = %s, want weekly", rank.LeaderboardType)
	}
}

func TestPrizeDistributionStructure(t *testing.T) {
	dist := model.PrizeDistribution{
		LeaderboardType: model.LeaderboardTypeDaily,
		GameType:        "slots",
		Period:          "2024-01-15",
		Prizes: []model.Prize{
			{Rank: 1, Type: "bonus", Value: 1000, Currency: "USD"},
			{Rank: 2, Type: "bonus", Value: 500, Currency: "USD"},
		},
		TotalValue: 1500,
	}

	if dist.TotalValue != 1500 {
		t.Fatalf("TotalValue = %f, want 1500", dist.TotalValue)
	}
	if len(dist.Prizes) != 2 {
		t.Fatalf("got %d prizes, want 2", len(dist.Prizes))
	}
}

func TestPrizeDistributionRequest(t *testing.T) {
	req := model.PrizeDistributionRequest{
		LeaderboardType: model.LeaderboardTypeDaily,
		GameType:        "slots",
		DryRun:          true,
	}

	if !req.DryRun {
		t.Fatal("DryRun should be true")
	}
	if req.LeaderboardType != model.LeaderboardTypeDaily {
		t.Fatalf("LeaderboardType = %s, want daily", req.LeaderboardType)
	}
}

func TestScoreRequestFields(t *testing.T) {
	req := model.UpdateScoreRequest{
		UserID:    "u1",
		Username:  "alice",
		Score:     500,
		GameType:  "blackjack",
		IsWin:     true,
		WinAmount: 250,
		BetAmount: 100,
	}

	if req.UserID != "u1" {
		t.Fatalf("UserID = %s, want u1", req.UserID)
	}
	if !req.IsWin {
		t.Fatal("IsWin should be true")
	}
}

func TestLeaderboardConfig(t *testing.T) {
	cfg := newTestLeaderboardConfig()

	if cfg.TopPlayersCount != 10 {
		t.Fatalf("TopPlayersCount = %d, want 10", cfg.TopPlayersCount)
	}
	if cfg.MinBetThreshold != 0.50 {
		t.Fatalf("MinBetThreshold = %f, want 0.50", cfg.MinBetThreshold)
	}
	if cfg.PrizeAutoCredit {
		t.Fatal("PrizeAutoCredit should be false")
	}
}

func TestLeaderboardEntrySorting(t *testing.T) {
	entries := []model.LeaderboardEntry{
		{Rank: 1, UserID: "u1", Score: 1000},
		{Rank: 2, UserID: "u2", Score: 800},
		{Rank: 3, UserID: "u3", Score: 600},
	}

	for i := 1; i < len(entries); i++ {
		if entries[i].Score > entries[i-1].Score {
			t.Fatalf("entries not sorted: entry %d score %f > entry %d score %f",
				i, entries[i].Score, i-1, entries[i-1].Score)
		}
	}
}

func TestPrizeStructure(t *testing.T) {
	prizes := []model.Prize{
		{Rank: 1, Type: "bonus", Value: 1000, Currency: "USD"},
		{Rank: 2, Type: "freespins", Value: 50},
		{Rank: 3, Type: "vip_points", Value: 100},
	}

	for _, p := range prizes {
		if p.Rank < 1 {
			t.Fatalf("prize rank %d should be >= 1", p.Rank)
		}
		if p.Type == "" {
			t.Fatal("prize type should not be empty")
		}
	}
}
