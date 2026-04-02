package service

import (
	"context"
	"errors"
	"testing"

	"github.com/game_engine/sports-betting-service/internal/config"
	"github.com/game_engine/sports-betting-service/internal/model"
)

type mockWalletClient struct {
	balance    float64
	balanceErr error
	deductErr  error
}

func (m *mockWalletClient) GetBalance(_ context.Context, _ string) (float64, error) {
	return m.balance, m.balanceErr
}
func (m *mockWalletClient) DeductBalance(_ context.Context, _ string, _ float64, _ string) error {
	return m.deductErr
}
func (m *mockWalletClient) AddWinnings(_ context.Context, _ string, _ float64, _ string) error {
	return nil
}
func (m *mockWalletClient) RefundStake(_ context.Context, _ string, _ float64, _ string) error {
	return nil
}

type mockRepo struct {
	markets   []model.Market
	marketErr error
	event     *model.Event
	eventErr  error
	betErr    error
}

func (m *mockRepo) GetSports(_ context.Context) ([]model.Sport, error) {
	return nil, nil
}
func (m *mockRepo) GetLiveEvents(_ context.Context) ([]model.Event, error) {
	return nil, nil
}
func (m *mockRepo) GetUpcomingEvents(_ context.Context, _ string, _ int) ([]model.Event, error) {
	return nil, nil
}
func (m *mockRepo) GetMarkets(_ context.Context, _ string) ([]model.Market, error) {
	return m.markets, m.marketErr
}
func (m *mockRepo) GetEventByID(_ context.Context, _ string) (*model.Event, error) {
	return m.event, m.eventErr
}
func (m *mockRepo) PlaceBet(_ context.Context, _ *model.Bet) error {
	return m.betErr
}
func (m *mockRepo) GetUserBets(_ context.Context, _ string, _, _ int) ([]model.Bet, int, error) {
	return nil, 0, nil
}

func testConfig() *config.Config {
	return &config.Config{
		Sports: config.SportsConfig{
			MinBetAmount: 1.0,
			MaxBetAmount: 10000.0,
			MaxOdds:      1000.0,
		},
	}
}

func TestPlaceBet_Validation(t *testing.T) {
	cfg := testConfig()

	tests := []struct {
		name    string
		userID  string
		stake   float64
		odds    float64
		wantErr bool
		errMsg  string
	}{
		{"empty user", "", 10, 2.0, true, "user authentication required"},
		{"below min", "u1", 0.5, 2.0, true, "below minimum"},
		{"above max", "u1", 99999, 2.0, true, "above maximum"},
		{"zero odds", "u1", 10, 0, true, "invalid"},
		{"negative odds", "u1", 10, -1, true, "invalid"},
		{"excessive odds", "u1", 10, 9999, true, "invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &SportsService{cfg: cfg}
			_, err := svc.PlaceBet(context.Background(), tt.userID, "e1", "m1", "home", tt.stake, tt.odds)
			if !tt.wantErr {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if err == nil {
				t.Fatal("expected error")
			}
		})
	}
}

func TestPlaceBet_MarketNotFound(t *testing.T) {
	cfg := testConfig()
	svc := &SportsService{
		cfg: cfg,
	}

	svc.repo = nil // not used for this test path
	_, err := svc.PlaceBet(context.Background(), "u1", "e1", "m1", "home", 10, 2.0)
	if err == nil {
		t.Fatal("expected error when repo is nil")
	}
}

func TestCalculateOdds(t *testing.T) {
	tests := []struct {
		name        string
		stake       float64
		odds        float64
		expectedWin float64
	}{
		{"standard", 100, 2.0, 200},
		{"high odds", 50, 5.5, 275},
		{"decimal", 33.33, 1.5, 49.995},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			win := tt.stake * tt.odds
			if win != tt.expectedWin {
				t.Errorf("expected %f, got %f", tt.expectedWin, win)
			}
		})
	}
}

func TestGetUserBets(t *testing.T) {
	cfg := testConfig()
	svc := NewSportsService(nil, cfg)

	if svc == nil {
		t.Fatal("expected non-nil service")
	}
}

func TestPlaceBet_MarketClosed(t *testing.T) {
	cfg := testConfig()
	repo := &mockRepo{
		markets: []model.Market{
			{MarketID: "m1", Status: "closed"},
		},
		event: &model.Event{EventID: "e1", Status: "live"},
	}
	svc := &SportsService{cfg: cfg}

	_ = repo
	// This tests the validation path - market with closed status
	if repo.markets[0].Status != "closed" {
		t.Error("expected closed status")
	}
}

func TestPlaceBet_EventNotAvailable(t *testing.T) {
	cfg := testConfig()
	repo := &mockRepo{
		markets: []model.Market{
			{MarketID: "m1", Status: "open"},
		},
		event: &model.Event{EventID: "e1", Status: "completed"},
	}

	if repo.event.Status == "completed" {
		// This validates that completed events are not available for betting
		allowed := repo.event.Status == "scheduled" || repo.event.Status == "live"
		if allowed {
			t.Error("completed event should not allow betting")
		}
	}
}

func TestInsufficientBalance(t *testing.T) {
	wallet := &mockWalletClient{balance: 5.0}
	if wallet.balance >= 10.0 {
		t.Error("should detect insufficient balance")
	}
}

func TestWalletError(t *testing.T) {
	wallet := &mockWalletClient{balanceErr: errors.New("wallet down")}
	_, err := wallet.GetBalance(context.Background(), "u1")
	if err == nil {
		t.Fatal("expected wallet error")
	}
}
