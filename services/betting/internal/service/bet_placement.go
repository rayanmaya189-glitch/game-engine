package service

import (
	"errors"
	"fmt"
	"time"
)

// PlaceSingle places a single bet
func (s *BettingService) PlaceSingle(userID, betID string, stake float64, selection Selection) (*Bet, error) {
	if stake < float64(s.minBet) {
		return nil, fmt.Errorf("minimum bet is %d", s.minBet)
	}
	if stake > float64(s.maxBet) {
		return nil, fmt.Errorf("maximum bet is %d", s.maxBet)
	}

	// Convert odds to decimal
	decimalOdds := s.normalizeOdds(selection.Odds, selection.OddsFormat)
	potentialWin := int64(stake * decimalOdds)

	if potentialWin > s.maxPayout {
		return nil, fmt.Errorf("maximum payout is %d", s.maxPayout)
	}

	bet := &Bet{
		ID:           betID,
		UserID:       userID,
		Type:         BetTypeSingle,
		Stake:        int64(stake),
		Odds:         decimalOdds,
		PotentialWin: potentialWin,
		Status:       BetStatusPlaced,
		Selections:   []Selection{selection},
		Currency:     s.defaultCurrency,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return bet, nil
}

// PlaceAccumulator places an accumulator bet
func (s *BettingService) PlaceAccumulator(userID, betID string, stake float64, selections []Selection) (*Bet, error) {
	if len(selections) < 2 {
		return nil, errors.New("accumulator requires at least 2 selections")
	}

	if stake < float64(s.minBet) {
		return nil, fmt.Errorf("minimum bet is %d", s.minBet)
	}
	if stake > float64(s.maxBet) {
		return nil, fmt.Errorf("maximum bet is %d", s.maxBet)
	}

	// Calculate cumulative odds
	cumulativeOdds := 1.0
	for i := range selections {
		selections[i].Odds = s.normalizeOdds(selections[i].Odds, selections[i].OddsFormat)
		cumulativeOdds *= selections[i].Odds
	}

	if cumulativeOdds > s.maxOdds {
		return nil, fmt.Errorf("maximum cumulative odds is %.2f", s.maxOdds)
	}

	potentialWin := int64(stake * cumulativeOdds)

	if potentialWin > s.maxPayout {
		return nil, fmt.Errorf("maximum payout is %d", s.maxPayout)
	}

	bet := &Bet{
		ID:           betID,
		UserID:       userID,
		Type:         BetTypeAccumulator,
		Stake:        int64(stake),
		Odds:         cumulativeOdds,
		PotentialWin: potentialWin,
		Status:       BetStatusPlaced,
		Selections:   selections,
		Currency:     s.defaultCurrency,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return bet, nil
}

// PlaceSystem places a system bet
func (s *BettingService) PlaceSystem(userID, betID string, stake float64, selections []Selection, systemType string) (*Bet, error) {
	validSystems := map[string]int{
		SystemPatent:     3, // 4 bets (3 singles + 3 doubles + 1 treble)
		SystemYankee:     4, // 11 bets
		SystemCanadian:   5, // 26 bets
		SystemHeinz:      6, // 57 bets
		SystemSuperHeinz: 7, // 120 bets
		SystemGoliath:    8, // 247 bets
	}

	minSelections, ok := validSystems[systemType]
	if !ok {
		return nil, errors.New("invalid system type")
	}

	if len(selections) < minSelections {
		return nil, fmt.Errorf("system %s requires at least %d selections", systemType, minSelections)
	}

	// Calculate number of combinations based on system type
	numCombos := s.calculateSystemCombinations(len(selections), systemType)
	totalStake := int64(stake) * int64(numCombos)

	if totalStake > s.maxBet {
		return nil, fmt.Errorf("total stake exceeds maximum of %d", s.maxBet)
	}

	// Calculate potential win (simplified - assumes all selections win)
	cumulativeOdds := 1.0
	for i := range selections {
		selections[i].Odds = s.normalizeOdds(selections[i].Odds, selections[i].OddsFormat)
		cumulativeOdds *= selections[i].Odds
	}

	// Simplified: estimate potential win based on all combinations
	// In reality, this would need detailed calculation per combination type
	potentialWin := int64(float64(totalStake) * cumulativeOdds)

	if potentialWin > s.maxPayout {
		return nil, fmt.Errorf("maximum payout is %d", s.maxPayout)
	}

	bet := &Bet{
		ID:           betID,
		UserID:       userID,
		Type:         BetTypeSystem,
		Stake:        totalStake,
		Odds:         cumulativeOdds,
		PotentialWin: potentialWin,
		Status:       BetStatusPlaced,
		Selections:   selections,
		SystemType:   systemType,
		Currency:     s.defaultCurrency,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return bet, nil
}

// calculateSystemCombinations calculates number of bets in a system bet
func (s *BettingService) calculateSystemCombinations(numSelections int, systemType string) int {
	// Simplified combination calculation
	switch systemType {
	case SystemPatent:
		return 4 + (numSelections-3)*3
	case SystemYankee:
		return 11 + (numSelections-4)*6
	case SystemCanadian:
		return 26 + (numSelections-5)*10
	case SystemHeinz:
		return 57 + (numSelections-6)*15
	case SystemSuperHeinz:
		return 120 + (numSelections-7)*21
	case SystemGoliath:
		return 247 + (numSelections-8)*28
	default:
		return 1
	}
}

// normalizeOdds converts odds to decimal format
func (s *BettingService) normalizeOdds(odds float64, format string) float64 {
	switch format {
	case OddsDecimal:
		return odds
	case OddsFractional:
		// 3/2 = 1 + 3/2 = 2.5
		parts := splitFraction(odds)
		return 1.0 + float64(parts[0])/float64(parts[1])
	case OddsAmerican:
		if odds > 0 {
			return 1 + odds/100
		}
		return 1 + 100/(-odds)
	case OddsHongKong:
		return 1 + odds
	default:
		return odds
	}
}

// splitFraction splits fractional odds like 3/2 into numerator and denominator
func splitFraction(odds float64) [2]int {
	// Simplified - in real implementation would parse string format
	return [2]int{1, 1}
}
