package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/game_engine/betting/internal/repository"
)

type mockBettingRepo struct {
	bets    map[string]*repository.Bet
	betList []*repository.Bet
	err     error
}

func newMockBettingRepo() *mockBettingRepo {
	return &mockBettingRepo{bets: make(map[string]*repository.Bet)}
}

func (m *mockBettingRepo) GetBetByID(ctx context.Context, betID string) (*repository.Bet, error) {
	if m.err != nil {
		return nil, m.err
	}
	b, ok := m.bets[betID]
	if !ok {
		return nil, fmt.Errorf("bet not found: %s", betID)
	}
	return b, nil
}

func (m *mockBettingRepo) GetBetHistory(ctx context.Context, userID string, limit, offset int) ([]*repository.Bet, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.betList, nil
}

func (m *mockBettingRepo) GetOpenBets(ctx context.Context, userID string) ([]*repository.Bet, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.betList, nil
}

func (m *mockBettingRepo) SaveBet(ctx context.Context, bet *repository.Bet) error {
	m.bets[bet.ID] = bet
	return m.err
}

func (m *mockBettingRepo) UpdateBet(ctx context.Context, bet *repository.Bet) error {
	m.bets[bet.ID] = bet
	return m.err
}

func setupBettingServiceTest(t *testing.T) (*BettingService, *mockBettingRepo) {
	t.Helper()
	repo := newMockBettingRepo()
	svc, err := NewBettingService(repo)
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}
	return svc, repo
}

func TestPlaceSingle(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)
	sel := Selection{Odds: 2.0, OddsFormat: OddsDecimal}

	bet, err := svc.PlaceSingle("u1", "b1", 10.0, sel)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bet.Type != BetTypeSingle || bet.Stake != 10 || bet.PotentialWin != 20 {
		t.Errorf("unexpected bet: type=%s stake=%d win=%d", bet.Type, bet.Stake, bet.PotentialWin)
	}
}

func TestPlaceSingle_Validation(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)
	sel := Selection{Odds: 2.0, OddsFormat: OddsDecimal}

	tests := []struct {
		name   string
		amount float64
	}{
		{"below min", 0.5},
		{"above max", 999999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := svc.PlaceSingle("u1", "b1", tt.amount, sel); err == nil {
				t.Error("expected error")
			}
		})
	}
}

func TestPlaceAccumulator(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)
	sels := []Selection{{Odds: 2.0, OddsFormat: OddsDecimal}, {Odds: 1.5, OddsFormat: OddsDecimal}}

	bet, err := svc.PlaceAccumulator("u1", "b1", 10.0, sels)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bet.Type != BetTypeAccumulator {
		t.Errorf("expected accumulator, got %s", bet.Type)
	}
}

func TestPlaceAccumulator_TooFew(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)
	if _, err := svc.PlaceAccumulator("u1", "b1", 10.0, []Selection{{Odds: 2.0}}); err == nil {
		t.Fatal("expected error")
	}
}

func TestPlaceSystem(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)
	sels := []Selection{{Odds: 2.0}, {Odds: 2.0}, {Odds: 2.0}}

	if _, err := svc.PlaceSystem("u1", "b1", 10.0, sels, "invalid"); err == nil {
		t.Fatal("expected error for invalid system")
	}

	for _, s := range []Selection{{Odds: 2.0, OddsFormat: OddsDecimal}, {Odds: 2.0, OddsFormat: OddsDecimal}, {Odds: 2.0, OddsFormat: OddsDecimal}} {
		_ = s
	}
	bet, err := svc.PlaceSystem("u1", "b1", 1.0, sels, SystemPatent)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bet.SystemType != SystemPatent {
		t.Errorf("expected patent, got %s", bet.SystemType)
	}
}

func TestSettleBet_Win(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)
	bet := &Bet{Status: BetStatusActive, Selections: []Selection{{OutcomeID: "o1", Status: OutcomePending}}}

	if err := svc.SettleBet(bet, map[string]string{"o1": "won"}); err != nil {
		t.Fatal(err)
	}
	if bet.Status != BetStatusSettled {
		t.Errorf("expected settled, got %s", bet.Status)
	}
}

func TestSettleBet_Loss(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)
	bet := &Bet{Status: BetStatusActive, Selections: []Selection{{OutcomeID: "o1", Status: OutcomePending}}}

	if err := svc.SettleBet(bet, map[string]string{"o1": "lost"}); err != nil {
		t.Fatal(err)
	}
	if bet.Selections[0].Status != OutcomeLost {
		t.Errorf("expected lost, got %s", bet.Selections[0].Status)
	}
}

func TestSettleBet_Errors(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)

	notActive := &Bet{Status: BetStatusPlaced, Selections: []Selection{{OutcomeID: "o1"}}}
	if err := svc.SettleBet(notActive, map[string]string{"o1": "won"}); err == nil {
		t.Error("expected error for non-active")
	}

	incomplete := &Bet{Status: BetStatusActive, Selections: []Selection{{OutcomeID: "o1"}, {OutcomeID: "o2"}}}
	if err := svc.SettleBet(incomplete, map[string]string{"o1": "won"}); err == nil {
		t.Error("expected error for incomplete")
	}
}

func TestGetBet(t *testing.T) {
	svc, repo := setupBettingServiceTest(t)
	repo.bets["b1"] = &repository.Bet{ID: "b1", UserID: "u1", Stake: 50}

	bet, err := svc.GetBet("b1")
	if err != nil {
		t.Fatal(err)
	}
	if bet.Stake != 50 {
		t.Errorf("expected 50, got %d", bet.Stake)
	}
}

func TestGetBet_Errors(t *testing.T) {
	nilRepo := &BettingService{}
	if _, err := nilRepo.GetBet("b1"); err == nil {
		t.Error("expected error for nil repo")
	}

	svc, repo := setupBettingServiceTest(t)
	repo.err = fmt.Errorf("db error")
	if _, err := svc.GetBet("b1"); err == nil {
		t.Error("expected error")
	}
}

func TestGetBetHistory(t *testing.T) {
	svc, repo := setupBettingServiceTest(t)
	repo.betList = []*repository.Bet{{ID: "b1"}, {ID: "b2"}}

	bets, err := svc.GetBetHistory("u1", 10, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(bets) != 2 {
		t.Errorf("expected 2, got %d", len(bets))
	}
}

func TestAcceptAndActivateBet(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)

	bet := &Bet{Status: BetStatusPlaced}
	if err := svc.AcceptBet(bet); err != nil || bet.Status != BetStatusAccepted {
		t.Error("accept failed")
	}
	if err := svc.ActivateBet(bet); err != nil || bet.Status != BetStatusActive {
		t.Error("activate failed")
	}
}

func TestVoidBet(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)

	bet := &Bet{Status: BetStatusPlaced}
	if err := svc.VoidBet(bet, "cancelled"); err != nil {
		t.Fatal(err)
	}
	if bet.Status != BetStatusVoided || bet.VoidReason != "cancelled" {
		t.Error("void failed")
	}

	if err := svc.VoidBet(&Bet{Status: BetStatusSettled}, ""); err == nil {
		t.Error("expected error for settled")
	}
}

func TestCalculatePayout(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)

	win := &Bet{Status: BetStatusSettled, PotentialWin: 200, Type: BetTypeSingle, Selections: []Selection{{Status: OutcomeWon}}}
	if got := svc.CalculatePayout(win); got != 200 {
		t.Errorf("expected 200, got %d", got)
	}

	_void := &Bet{Status: BetStatusSettled, Stake: 100, Type: BetTypeSingle, Selections: []Selection{{Status: OutcomeVoid}}}
	if got := svc.CalculatePayout(_void); got != 100 {
		t.Errorf("expected 100, got %d", got)
	}

	notSettled := &Bet{Status: BetStatusActive}
	if got := svc.CalculatePayout(notSettled); got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestGetBetTypeInfo(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)

	for _, bType := range []string{BetTypeSingle, BetTypeAccumulator, BetTypeSystem} {
		if _, _, err := svc.GetBetTypeInfo(bType); err != nil {
			t.Errorf("unexpected error for %s: %v", bType, err)
		}
	}
	if _, _, err := svc.GetBetTypeInfo("unknown"); err == nil {
		t.Error("expected error for unknown type")
	}
}

func TestGetSystemBetInfo(t *testing.T) {
	svc, _ := setupBettingServiceTest(t)

	for _, s := range []string{SystemPatent, SystemYankee, SystemCanadian, SystemHeinz, SystemSuperHeinz, SystemGoliath} {
		if _, _, err := svc.GetSystemBetInfo(s); err != nil {
			t.Errorf("unexpected error for %s: %v", s, err)
		}
	}
	if _, _, err := svc.GetSystemBetInfo("invalid"); err == nil {
		t.Error("expected error")
	}
}

func TestNewBettingServiceConfig(t *testing.T) {
	if _, err := NewBettingServiceWithConfig(0, 100, 1000, nil); err == nil {
		t.Error("expected error for zero min")
	}
	if _, err := NewBettingServiceWithConfig(100, 50, 1000, nil); err == nil {
		t.Error("expected error for max < min")
	}
}
