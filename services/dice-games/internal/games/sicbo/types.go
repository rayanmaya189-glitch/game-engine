package sicbo

import (
	"errors"
	"fmt"
	"time"
)

// Bet types
const (
	BetSmall          = "small"           // Total 4-10 (excluding triples)
	BetBig            = "big"             // Total 11-17 (excluding triples)
	BetSpecificTriple = "specific_triple" // All three dice same specific number
	BetAnyTriple      = "any_triple"      // All three dice same (any number)
	BetSpecificDouble = "specific_double" // Two dice show same specific number
	BetFourNumber     = "four_number"     // Bet on 4 specific numbers
	BetThreeNumber    = "three_number"    // Three dice total specific combination
	BetTwoNumber      = "two_number"      // Two dice specific combination
	BetSingle         = "single"          // One specific number appears
)

// Payout ratios for each bet type
var payoutRatios = map[string]float64{
	BetSmall:          1.0,   // 1:1
	BetBig:            1.0,   // 1:1
	BetSpecificTriple: 150.0, // 150:1
	BetAnyTriple:      24.0,  // 24:1
	BetSpecificDouble: 8.0,   // 8:1
	BetFourNumber:     7.0,   // 7:1
	BetThreeNumber:    50.0,  // 50:1
	BetTwoNumber:      5.0,   // 5:1
	BetSingle:         1.0,   // 1:1 (varies by count)
}

// Single bet multipliers by count
var singleBetMultipliers = map[int]int64{
	1: 1,
	2: 2,
	3: 3,
}

// Bet information maps
var betNames = map[string]string{
	BetSmall:          "Small",
	BetBig:            "Big",
	BetSpecificTriple: "Specific Triple",
	BetAnyTriple:      "Any Triple",
	BetSpecificDouble: "Specific Double",
	BetFourNumber:     "Four Number",
	BetThreeNumber:    "Three Number",
	BetTwoNumber:      "Two Number",
	BetSingle:         "Single",
}

// Roller interface for dice rolling
type Roller interface {
	Roll(diceCount int) ([]int, error)
}

// Config holds Sic Bo configuration
type Config struct {
	MinBet    int64
	MaxBet    int64
	DiceCount int
}

// GameState represents the game state
type GameState struct {
	GameID         string                      `json:"game_id"`
	Dice           []int                       `json:"dice"`
	Sum            int                         `json:"sum"`
	Bets           map[string]map[string]int64 `json:"bets"`
	Payouts        map[string]map[string]int64 `json:"payouts"`
	ProvablyFair   bool                        `json:"provably_fair"`
	ServerSeedHash string                      `json:"server_seed_hash,omitempty"`
	ClientSeed     string                      `json:"client_seed,omitempty"`
	Nonce          int                         `json:"nonce,omitempty"`
	IsTriple       bool                        `json:"is_triple"`
	IsDouble       bool                        `json:"is_double"`
	CreatedAt      time.Time                   `json:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at"`
}

// ValidateBet validates a bet amount and type
func (g *Game) ValidateBet(betType string, amount int64) error {
	if amount <= 0 {
		return errors.New("bet amount must be positive")
	}

	if amount < g.MinBet {
		return fmt.Errorf("minimum bet is %d", g.MinBet)
	}

	if amount > g.MaxBet {
		return fmt.Errorf("maximum bet is %d", g.MaxBet)
	}

	// Check bet type exists
	if _, ok := payoutRatios[betType]; !ok {
		// Check if it's a numbered bet type
		if len(betType) > 7 {
			prefix := betType[:len(betType)-2]
			if prefix == "single" || prefix == "specific_double" || prefix == "specific_triple" {
				return nil
			}
		}
		return errors.New("invalid bet type")
	}

	return nil
}
