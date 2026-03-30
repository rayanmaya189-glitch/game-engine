package sicbo

import (
	"errors"
	"fmt"
)

// PlaceBet places a bet for a player
func (g *Game) PlaceBet(betType, playerID string, amount int64) error {
	if amount <= 0 {
		return errors.New("bet amount must be positive")
	}

	if amount < g.MinBet || amount > g.MaxBet {
		return fmt.Errorf("bet amount must be between %d and %d", g.MinBet, g.MaxBet)
	}

	// Validate bet type
	if _, ok := payoutRatios[betType]; !ok {
		return errors.New("invalid bet type")
	}

	if g.Bets[betType] == nil {
		g.Bets[betType] = make(map[string]int64)
	}
	g.Bets[betType][playerID] += amount

	return nil
}

// PlaceBetWithNumber places a bet that requires a specific number (for single, double, etc.)
func (g *Game) PlaceBetWithNumber(betType, playerID string, number int, amount int64) error {
	if number < 1 || number > 6 {
		return errors.New("number must be between 1 and 6")
	}

	betKey := fmt.Sprintf("%s_%d", betType, number)
	return g.PlaceBet(betKey, playerID, amount)
}

// ClearBets clears all bets for a player
func (g *Game) ClearBets(playerID string) {
	for betType, playerBets := range g.Bets {
		delete(playerBets, playerID)
		if len(playerBets) == 0 {
			delete(g.Bets, betType)
		}
	}
}

// ClearAllBets clears all bets
func (g *Game) ClearAllBets() {
	g.Bets = make(map[string]map[string]int64)
}

// GetBet retrieves the bet amount for a specific player and bet type
func (g *Game) GetBet(betType, playerID string) int64 {
	if bets, ok := g.Bets[betType]; ok {
		return bets[playerID]
	}
	return 0
}

// GetTotalBets calculates the total bet amount for a player
func (g *Game) GetTotalBets(playerID string) int64 {
	total := int64(0)
	for _, playerBets := range g.Bets {
		total += playerBets[playerID]
	}
	return total
}

// GetBetInfo returns information about a specific bet type
func (g *Game) GetBetInfo(betType string) (string, string, error) {
	name, ok := betNames[betType]
	if !ok {
		return "", "", errors.New("invalid bet type")
	}

	ratio, ok := payoutRatios[betType]
	if !ok {
		return "", "", errors.New("bet type not found")
	}

	return name, fmt.Sprintf("%.0f:1", ratio), nil
}

// GetAvailableBets returns all available bet types
func (g *Game) GetAvailableBets() []string {
	return []string{
		BetSmall, BetBig, BetAnyTriple,
		BetFourNumber, BetThreeNumber, BetTwoNumber,
	}
}

// GetAvailableSingleBets returns single number bet keys
func (g *Game) GetAvailableSingleBets() []string {
	bets := []string{}
	for i := 1; i <= 6; i++ {
		bets = append(bets, fmt.Sprintf("single_%d", i))
	}
	return bets
}

// GetAvailableDoubleBets returns specific double bet keys
func (g *Game) GetAvailableDoubleBets() []string {
	bets := []string{}
	for i := 1; i <= 6; i++ {
		bets = append(bets, fmt.Sprintf("specific_double_%d", i))
	}
	return bets
}

// GetAvailableTripleBets returns specific triple bet keys
func (g *Game) GetAvailableTripleBets() []string {
	bets := []string{}
	for i := 1; i <= 6; i++ {
		bets = append(bets, fmt.Sprintf("specific_triple_%d", i))
	}
	return bets
}
