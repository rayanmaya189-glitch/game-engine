package service

import (
	"errors"
	"fmt"
)

// GetBet returns a bet by ID (would query database in real implementation)
func (s *BettingService) GetBet(betID string) (*Bet, error) {
	// In real implementation, would query database
	return nil, errors.New("not implemented")
}

// GetBetHistory returns bet history for a user
func (s *BettingService) GetBetHistory(userID string, limit, offset int) ([]*Bet, error) {
	// In real implementation, would query database
	return []*Bet{}, nil
}

// GetOpenBets returns all open (non-settled) bets for a user
func (s *BettingService) GetOpenBets(userID string) ([]*Bet, error) {
	// In real implementation, would query database
	return []*Bet{}, nil
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
