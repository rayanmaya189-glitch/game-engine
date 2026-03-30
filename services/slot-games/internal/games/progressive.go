package games

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// NewProgressiveGame creates a new Progressive Jackpot slot game
func NewProgressiveGame(id, gameType string) *ProgressiveGame {
	config := &ProgressiveConfig{
		Reels:      5,
		Rows:       3,
		MinBet:     1,
		MaxBet:     100,
		RTP:        94.5, // Lower base RTP due to jackpot contribution
		Volatility: "high",
		Jackpots: map[string]ProgressiveJackpot{
			"Mini": {
				Name:             "Mini",
				SeedAmount:       100,
				ContributionRate: 0.02, // 2%
				MinBet:           1,
				Trigger:          "symbol",
				Symbols:          []string{"Mini1", "Mini2", "Mini3"},
				Odds:             0.001, // 0.1%
			},
			"Minor": {
				Name:             "Minor",
				SeedAmount:       500,
				ContributionRate: 0.015, // 1.5%
				MinBet:           1,
				Trigger:          "symbol",
				Symbols:          []string{"Minor1", "Minor2", "Minor3"},
				Odds:             0.0005, // 0.05%
			},
			"Major": {
				Name:             "Major",
				SeedAmount:       5000,
				ContributionRate: 0.01, // 1%
				MinBet:           5,
				Trigger:          "symbol",
				Symbols:          []string{"Major1", "Major2", "Major3"},
				Odds:             0.0001, // 0.01%
			},
			"Grand": {
				Name:             "Grand",
				SeedAmount:       50000,
				ContributionRate: 0.005, // 0.5%
				MinBet:           10,
				Trigger:          "random",
				Odds:             0.00001, // 0.001%
			},
		},
	}

	symbols := []Symbol{
		{ID: "W", Name: "Wild", Value: 10, IsWild: true, Weight: 4},
		{ID: "S", Name: "Scatter", Value: 9, IsScatter: true, Weight: 6},
		{ID: "Mini", Name: "Mini Jackpot", Value: 8, Weight: 4, IsBonus: true},
		{ID: "Minor", Name: "Minor Jackpot", Value: 7, Weight: 3, IsBonus: true},
		{ID: "Major", Name: "Major Jackpot", Value: 6, Weight: 2, IsBonus: true},
		{ID: "A", Name: "Ace", Value: 5, Weight: 12},
		{ID: "K", Name: "King", Value: 4, Weight: 14},
		{ID: "Q", Name: "Queen", Value: 3, Weight: 16},
		{ID: "J", Name: "Jack", Value: 2, Weight: 18},
		{ID: "10", Name: "Ten", Value: 1, Weight: 20},
	}

	return &ProgressiveGame{
		ID:        id,
		GameType:  gameType,
		Config:    config,
		Symbols:   symbols,
		Positions: make([][]int, config.Reels),
	}
}

// SetBet sets the bet amount
func (g *ProgressiveGame) SetBet(lineBet int64) error {
	if lineBet < g.Config.MinBet {
		return fmt.Errorf("minimum bet is %d", g.Config.MinBet)
	}
	if lineBet > g.Config.MaxBet {
		return fmt.Errorf("maximum bet is %d", g.Config.MaxBet)
	}

	g.LineBet = lineBet
	g.Bet = lineBet
	return nil
}

// Spin performs a spin
func (g *ProgressiveGame) Spin() error {
	if g.Bet < g.Config.MinBet {
		return fmt.Errorf("bet not set")
	}

	g.Win = 0
	g.TotalWin = 0
	g.JackpotTriggered = false
	g.JackpotTier = ""
	g.JackpotWin = 0

	// Generate random grid
	g.generateGrid()

	// Check for regular wins
	g.evaluateWins()

	// Check for jackpot (simplified probability check)
	g.checkJackpot()

	g.IsComplete = true
	return nil
}

// generateGrid generates random symbols
func (g *ProgressiveGame) generateGrid() {
	g.Positions = make([][]int, g.Config.Reels)
	for reel := 0; reel < g.Config.Reels; reel++ {
		g.Positions[reel] = make([]int, g.Config.Rows)
		for row := 0; row < g.Config.Rows; row++ {
			symbolIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(g.Symbols))))
			g.Positions[reel][row] = int(symbolIndex.Int64())
		}
	}
}

// evaluateWins evaluates regular wins
func (g *ProgressiveGame) evaluateWins() {
	// Simplified win evaluation - check for 3+ matching symbols
	for symbolIdx, symbol := range g.Symbols {
		if symbol.IsWild || symbol.IsScatter || symbol.IsBonus {
			continue
		}

		count := 0
		for reel := 0; reel < g.Config.Reels; reel++ {
			for row := 0; row < g.Config.Rows; row++ {
				if g.Positions[reel][row] == symbolIdx {
					count++
				}
			}
		}

		if count >= 3 {
			g.Win += int64(symbol.Value) * g.LineBet * int64(count-2)
		}
	}

	g.TotalWin = g.Win
}

// checkJackpot checks if jackpot is triggered
func (g *ProgressiveGame) checkJackpot() {
	// First check symbol-based jackpots
	for tier, jackpot := range g.Config.Jackpots {
		if jackpot.Trigger != "symbol" {
			continue
		}

		// Check minimum bet requirement
		if g.Bet < jackpot.MinBet {
			continue
		}

		// Check if jackpot symbols are present
		symbolCount := 0
		for _, sym := range jackpot.Symbols {
			for reel := 0; reel < g.Config.Reels; reel++ {
				for row := 0; row < g.Config.Rows; row++ {
					if g.Symbols[g.Positions[reel][row]].ID == sym {
						symbolCount++
					}
				}
			}
		}

		// Need 3 jackpot symbols to trigger
		if symbolCount >= 3 {
			g.JackpotTriggered = true
			g.JackpotTier = tier
			g.JackpotWin = jackpot.SeedAmount
			return
		}
	}

	// Check random-triggered jackpots (Grand)
	grandJackpot := g.Config.Jackpots["Grand"]
	if g.Bet >= grandJackpot.MinBet {
		randNum, _ := rand.Int(rand.Reader, big.NewInt(1000000))
		threshold := int64(grandJackpot.Odds * 1000000)
		if randNum.Int64() < threshold {
			g.JackpotTriggered = true
			g.JackpotTier = "Grand"
			g.JackpotWin = grandJackpot.SeedAmount
		}
	}
}

// GetJackpotAmount returns the current jackpot amount for a tier
func (g *ProgressiveGame) GetJackpotAmount(tier string) (int64, error) {
	jackpot, ok := g.Config.Jackpots[tier]
	if !ok {
		return 0, fmt.Errorf("unknown jackpot tier: %s", tier)
	}
	return jackpot.SeedAmount, nil
}

// GetState returns the game state
func (g *ProgressiveGame) GetState() *ProgressiveGameState {
	symbols := make([][]string, len(g.Positions))
	for i, reel := range g.Positions {
		symbols[i] = make([]string, len(reel))
		for j, s := range reel {
			symbols[i][j] = g.Symbols[s].ID
		}
	}

	return &ProgressiveGameState{
		GameID:           g.ID,
		GameType:         g.GameType,
		Reels:            g.Config.Reels,
		Rows:             g.Config.Rows,
		Symbols:          symbols,
		Bet:              g.Bet,
		Win:              g.Win,
		TotalWin:         g.TotalWin,
		JackpotTriggered: g.JackpotTriggered,
		JackpotTier:      g.JackpotTier,
		JackpotWin:       g.JackpotWin,
		IsComplete:       g.IsComplete,
	}
}
