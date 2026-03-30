package craps

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

// PlaceBet places a bet for a player
func (g *Game) PlaceBet(betType, playerID string, amount int64) error {
	if amount <= 0 {
		return errors.New("bet amount must be positive")
	}

	// Validate bet type
	if _, ok := payoutRatios[betType]; !ok {
		return errors.New("invalid bet type")
	}

	// Check if bet is allowed in current phase
	if !g.isBetAllowed(betType) {
		return errors.New("bet not allowed in current phase")
	}

	if g.Bets[betType] == nil {
		g.Bets[betType] = make(map[string]int64)
	}
	g.Bets[betType][playerID] += amount

	return nil
}

// isBetAllowed checks if a bet type is allowed in the current phase
func (g *Game) isBetAllowed(betType string) bool {
	switch g.Phase {
	case PhaseComeOut:
		// These bets can be placed during come-out
		switch betType {
		case BetPassLine, BetDontPass, BetField, BetAny7, BetAnyCraps,
			BetHorn2, BetHorn3, BetHorn11, BetHorn12:
			return true
		default:
			return false
		}
	case PhasePoint:
		// All bets allowed during point phase
		return true
	default:
		return false
	}
}

// GetPayout retrieves the payout for a specific player and bet type
func (g *Game) GetPayout(betType, playerID string) int64 {
	if payouts, ok := g.Payouts[betType]; ok {
		return payouts[playerID]
	}
	return 0
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

// GetTotalPayouts calculates the total payout for a player
func (g *Game) GetTotalPayouts(playerID string) int64 {
	total := int64(0)
	for _, playerPayouts := range g.Payouts {
		total += playerPayouts[playerID]
	}
	return total
}

// GetAvailableBets returns all bet types available in current phase
func (g *Game) GetAvailableBets() []string {
	if g.Phase == PhaseComeOut {
		return []string{
			BetPassLine, BetDontPass, BetField, BetAny7, BetAnyCraps,
			BetHorn2, BetHorn3, BetHorn11, BetHorn12,
		}
	}
	return []string{
		BetPassLine, BetDontPass, BetField, BetBig6, BetBig8, BetAny7, BetAnyCraps,
		BetHorn2, BetHorn3, BetHorn11, BetHorn12,
		BetHard4, BetHard6, BetHard8, BetHard10,
		BetPlaceWin4, BetPlaceWin5, BetPlaceWin6, BetPlaceWin8, BetPlaceWin9, BetPlaceWin10,
		BetPlaceLose4, BetPlaceLose5, BetPlaceLose6, BetPlaceLose8, BetPlaceLose9, BetPlaceLose10,
		BetBuy4, BetBuy5, BetBuy6, BetBuy8, BetBuy9, BetBuy10,
		BetLay4, BetLay5, BetLay6, BetLay8, BetLay9, BetLay10,
	}
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

// GetBetInfo returns information about a specific bet type
func (g *Game) GetBetInfo(betType string) (string, string, error) {
	name, ok := betNames[betType]
	if !ok {
		return "", "", errors.New("invalid bet type")
	}

	ratio := payoutRatios[betType]

	return name, fmt.Sprintf("%.2f:1", ratio), nil
}

// GetState returns the game state
func (g *Game) GetState() *GameState {
	return &GameState{
		GameID:         g.ID,
		Dice:           g.Dice,
		Sum:            g.Dice[0] + g.Dice[1],
		Point:          g.Point,
		Phase:          g.Phase,
		Bets:           g.Bets,
		Payouts:        g.Payouts,
		ProvablyFair:   g.ProvablyFair,
		ServerSeedHash: hashSeed(g.ServerSeed),
		ClientSeed:     g.ClientSeed,
		Nonce:          g.Nonce,
		CreatedAt:      g.CreatedAt,
		UpdatedAt:      g.UpdatedAt,
	}
}

// hashSeed creates a hash of the seed for display
func hashSeed(seed string) string {
	if seed == "" {
		return ""
	}
	h := sha256.Sum256([]byte(seed))
	return fmt.Sprintf("%x", h[:8])
}
