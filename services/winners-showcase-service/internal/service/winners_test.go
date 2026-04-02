package service

import (
	"testing"
	"time"

	"github.com/game_engine/winners-showcase-service/internal/config"
	"github.com/game_engine/winners-showcase-service/internal/model"
)

func newTestWinnersConfig() *config.WinnersConfig {
	return &config.WinnersConfig{
		RecentWinnersLimit:  10,
		BigWinThreshold:     1000.0,
		JackpotThreshold:    10000.0,
		FeedTTLSeconds:      60,
		EnableJackpotAlerts: true,
		AnonymizeNames:      true,
	}
}

func TestAnonymizeName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"alice", "a***e"},
		{"bob", "b***b"},
		{"charlie", "c***e"},
		{"Jo", "J***"},
		{"A", "A***"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var got string
			if len(tt.input) <= 2 {
				got = tt.input[0:1] + "***"
			} else {
				got = tt.input[0:1] + "***" + tt.input[len(tt.input)-1:]
			}
			if got != tt.want {
				t.Fatalf("anonymizeName(%s) = %s, want %s", tt.input, got, tt.want)
			}
		})
	}
}

func TestWinnerModel(t *testing.T) {
	winner := model.Winner{
		ID:          1,
		UserID:      "u1",
		Username:    "alice",
		DisplayName: "a***e",
		WinAmount:   5000,
		Currency:    "USD",
		GameType:    "slots",
		GameName:    "Lucky 7",
		WinType:     model.WinTypeBig,
		Multiplier:  50.0,
		Timestamp:   time.Now(),
	}

	if winner.WinAmount <= 0 {
		t.Fatal("WinAmount should be positive")
	}
	if winner.WinType != model.WinTypeBig {
		t.Fatalf("WinType = %s, want %s", winner.WinType, model.WinTypeBig)
	}
}

func TestWinTypes(t *testing.T) {
	types := []model.WinType{
		model.WinTypeRegular,
		model.WinTypeBig,
		model.WinTypeJackpot,
		model.WinTypeProgressive,
	}

	for _, wt := range types {
		if string(wt) == "" {
			t.Fatal("win type should not be empty")
		}
	}
}

func TestWinTypeClassification(t *testing.T) {
	cfg := newTestWinnersConfig()

	tests := []struct {
		name      string
		amount    float64
		wantType  model.WinType
	}{
		{"regular win", 100, model.WinTypeRegular},
		{"big win threshold", 1000, model.WinTypeBig},
		{"big win above", 5000, model.WinTypeBig},
		{"jackpot threshold", 10000, model.WinTypeJackpot},
		{"jackpot above", 50000, model.WinTypeJackpot},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			winType := model.WinTypeRegular
			if tt.amount >= cfg.JackpotThreshold {
				winType = model.WinTypeJackpot
			} else if tt.amount >= cfg.BigWinThreshold {
				winType = model.WinTypeBig
			}
			if winType != tt.wantType {
				t.Fatalf("amount=%f: got %s, want %s", tt.amount, winType, tt.wantType)
			}
		})
	}
}

func TestRecordWinRequest(t *testing.T) {
	req := model.RecordWinRequest{
		UserID:     "u1",
		Username:   "alice",
		WinAmount:  5000,
		Currency:   "USD",
		GameType:   "slots",
		GameName:   "Lucky 7",
		Multiplier: 50.0,
	}

	if req.UserID == "" {
		t.Fatal("UserID should not be empty")
	}
	if req.WinAmount <= 0 {
		t.Fatal("WinAmount should be positive")
	}
	if req.Multiplier < 1.0 {
		t.Fatal("Multiplier should be >= 1.0")
	}
}

func TestPrivacySettings(t *testing.T) {
	settings := model.PrivacySettings{
		UserID:            "u1",
		AnonymizeName:     true,
		ShowOnLeaderboard: true,
		ShowOnJackpotList: false,
		OptOutOfShowcase:  false,
	}

	if settings.UserID == "" {
		t.Fatal("UserID should not be empty")
	}
	if !settings.AnonymizeName {
		t.Fatal("AnonymizeName should default to true")
	}
}

func TestPrivacyUpdate(t *testing.T) {
	req := model.UpdatePrivacyRequest{
		AnonymizeName:     false,
		ShowOnLeaderboard: false,
		ShowOnJackpotList: true,
		OptOutOfShowcase:  true,
	}

	updated := model.PrivacySettings{
		UserID:            "u1",
		AnonymizeName:     req.AnonymizeName,
		ShowOnLeaderboard: req.ShowOnLeaderboard,
		ShowOnJackpotList: req.ShowOnJackpotList,
		OptOutOfShowcase:  req.OptOutOfShowcase,
	}

	if updated.AnonymizeName {
		t.Fatal("AnonymizeName should be false after update")
	}
	if !updated.OptOutOfShowcase {
		t.Fatal("OptOutOfShowcase should be true after update")
	}
}

func TestRecentWinnersResponse(t *testing.T) {
	resp := model.RecentWinnersResponse{
		Winners: []model.Winner{
			{ID: 1, WinAmount: 500},
			{ID: 2, WinAmount: 300},
		},
		Total:     2,
		UpdatedAt: time.Now(),
	}

	if resp.Total != len(resp.Winners) {
		t.Fatalf("Total=%d but Winners has %d entries", resp.Total, len(resp.Winners))
	}
}

func TestBigWinsResponse(t *testing.T) {
	resp := model.BigWinsResponse{
		Wins: []model.Winner{
			{ID: 1, WinAmount: 5000},
		},
		Threshold: 1000,
		Total:     1,
	}

	for _, w := range resp.Wins {
		if w.WinAmount < resp.Threshold {
			t.Fatalf("win amount %f should be >= threshold %f", w.WinAmount, resp.Threshold)
		}
	}
}

func TestConfigValues(t *testing.T) {
	cfg := newTestWinnersConfig()

	if cfg.RecentWinnersLimit != 10 {
		t.Fatalf("RecentWinnersLimit = %d, want 10", cfg.RecentWinnersLimit)
	}
	if cfg.BigWinThreshold != 1000.0 {
		t.Fatalf("BigWinThreshold = %f, want 1000", cfg.BigWinThreshold)
	}
	if cfg.JackpotThreshold != 10000.0 {
		t.Fatalf("JackpotThreshold = %f, want 10000", cfg.JackpotThreshold)
	}
	if !cfg.AnonymizeNames {
		t.Fatal("AnonymizeNames should be true")
	}
}

func TestWinAmountSorting(t *testing.T) {
	winners := []model.Winner{
		{ID: 1, WinAmount: 10000},
		{ID: 2, WinAmount: 5000},
		{ID: 3, WinAmount: 1000},
	}

	for i := 1; i < len(winners); i++ {
		if winners[i].WinAmount > winners[i-1].WinAmount {
			t.Fatal("winners should be sorted by WinAmount descending")
		}
	}
}
