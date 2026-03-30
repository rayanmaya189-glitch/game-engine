package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/game_engine/betting/internal/repository"
)

// GetBet returns a bet by ID from the repository
func (s *BettingService) GetBet(betID string) (*Bet, error) {
	if s.repo == nil {
		return nil, errors.New("repository not configured")
	}
	repoBet, err := s.repo.GetBetByID(context.Background(), betID)
	if err != nil {
		return nil, fmt.Errorf("bet not found: %w", err)
	}
	return repoToServiceBet(repoBet), nil
}

// GetBetHistory returns bet history for a user from the repository
func (s *BettingService) GetBetHistory(userID string, limit, offset int) ([]*Bet, error) {
	if s.repo == nil {
		return nil, errors.New("repository not configured")
	}
	repoBets, err := s.repo.GetBetHistory(context.Background(), userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get bet history: %w", err)
	}
	bets := make([]*Bet, len(repoBets))
	for i, rb := range repoBets {
		bets[i] = repoToServiceBet(rb)
	}
	return bets, nil
}

// GetOpenBets returns all open (non-settled) bets for a user from the repository
func (s *BettingService) GetOpenBets(userID string) ([]*Bet, error) {
	if s.repo == nil {
		return nil, errors.New("repository not configured")
	}
	repoBets, err := s.repo.GetOpenBets(context.Background(), userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get open bets: %w", err)
	}
	bets := make([]*Bet, len(repoBets))
	for i, rb := range repoBets {
		bets[i] = repoToServiceBet(rb)
	}
	return bets, nil
}

// repoToServiceBet converts a repository Bet to a service Bet
func repoToServiceBet(rb *repository.Bet) *Bet {
	selections := make([]Selection, len(rb.Selections))
	for i, s := range rb.Selections {
		selections[i] = Selection{
			ID:         s.ID,
			EventID:    s.EventID,
			OutcomeID:  s.OutcomeID,
			Odds:       s.Odds,
			OddsFormat: s.OddsFormat,
			Status:     s.Status,
			Result:     s.Result,
			SettledAt:  s.SettledAt,
		}
	}
	return &Bet{
		ID:           rb.ID,
		UserID:       rb.UserID,
		Type:         rb.Type,
		Stake:        rb.Stake,
		Odds:         rb.Odds,
		PotentialWin: rb.PotentialWin,
		Status:       rb.Status,
		Selections:   selections,
		SystemType:   rb.SystemType,
		Currency:     rb.Currency,
		CreatedAt:    rb.CreatedAt,
		UpdatedAt:    rb.UpdatedAt,
		SettledAt:    rb.SettledAt,
		VoidReason:   rb.VoidReason,
	}
}

// GetBetTypeInfo returns information about a bet type
func (s *BettingService) GetBetTypeInfo(betType string) (string, int, error) {
	switch betType {
	case BetTypeSingle:
		return "Single Bet", 1, nil
	case BetTypeAccumulator:
		return "Accumulator", 2, nil
	case BetTypeSystem:
		return "System Bet", 3, nil
	default:
		return "", 0, errors.New("unknown bet type")
	}
}

// GetSystemBetInfo returns information about a system bet type
func (s *BettingService) GetSystemBetInfo(systemType string) (string, int, error) {
	info := map[string][2]int{
		SystemPatent:     {4, 3},
		SystemYankee:     {11, 4},
		SystemCanadian:   {26, 5},
		SystemHeinz:      {57, 6},
		SystemSuperHeinz: {120, 7},
		SystemGoliath:    {247, 8},
	}

	val, ok := info[systemType]
	if !ok {
		return "", 0, errors.New("unknown system type")
	}
	return fmt.Sprintf("%d bets", val[0]), val[1], nil
}

// CalculateOddsFromDecimal calculates various format odds from decimal
func (s *BettingService) CalculateOddsFromDecimal(decimal float64) map[string]float64 {
	return map[string]float64{
		OddsDecimal:    decimal,
		OddsFractional: decimal - 1,
		OddsAmerican:   (decimal - 1) * 100,
		OddsHongKong:   decimal - 1,
	}
}
