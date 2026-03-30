package games

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// NewMegawaysGame creates a new Megaways slot game
func NewMegawaysGame(id string) *MegawaysGame {
	config := &MegawaysConfig{
		Reels:                6,
		MinRows:              2,
		MaxRows:              7,
		MinBet:               1,
		MaxBet:               500,
		MinLineBet:           1,
		MaxLineBet:           10,
		RTP:                  96.0,
		Volatility:           "high",
		MaxMegaways:          117649,
		CascadeEnabled:       true,
		CascadeMultiplier:    1.0,
		MaxCascadeMultiplier: 5.0,
	}

	symbols := []Symbol{
		{ID: "W", Name: "Wild", Value: 10, IsWild: true, Weight: 4},
		{ID: "S", Name: "Scatter", Value: 9, IsScatter: true, Weight: 8},
		{ID: "A", Name: "Ace", Value: 8, Weight: 12},
		{ID: "K", Name: "King", Value: 7, Weight: 14},
		{ID: "Q", Name: "Queen", Value: 6, Weight: 16},
		{ID: "J", Name: "Jack", Value: 5, Weight: 18},
		{ID: "10", Name: "Ten", Value: 4, Weight: 20},
		{ID: "9", Name: "Nine", Value: 3, Weight: 22},
		{ID: "Gem1", Name: "Purple Gem", Value: 2, Weight: 10},
		{ID: "Gem2", Name: "Red Gem", Value: 1, Weight: 12},
	}

	return &MegawaysGame{
		ID:                id,
		Config:            config,
		Symbols:           symbols,
		ReelSymbols:       make([][]int, config.Reels),
		RowCounts:         make([]int, config.Reels),
		CascadeMultiplier: 1.0,
		ProvablyFair:      false,
	}
}

// SetBet sets the bet amount
func (g *MegawaysGame) SetBet(lineBet int64, ways int) error {
	if lineBet < g.Config.MinLineBet || lineBet > g.Config.MaxLineBet {
		return fmt.Errorf("line bet must be between %d and %d", g.Config.MinLineBet, g.Config.MaxLineBet)
	}

	g.LineBet = lineBet
	g.Bet = lineBet // Total bet = line bet (Megaways doesn't use traditional paylines)

	return nil
}

// Spin performs a spin with random symbols
func (g *MegawaysGame) Spin() error {
	if g.Bet < g.Config.MinBet {
		return fmt.Errorf("minimum bet is %d", g.Config.MinBet)
	}

	// Initialize cascade
	g.CascadeLevel = 0
	g.CascadeMultiplier = 1.0
	g.WinLines = nil
	g.CascadeWins = nil
	g.TotalWin = 0

	// First spin
	g.spinOnce()

	// Handle cascades if enabled
	if g.Config.CascadeEnabled {
		for g.CascadeLevel < 10 { // Max 10 cascades
			if !g.hasWinningCombinations() {
				break
			}
			g.CascadeLevel++
			g.CascadeMultiplier = 1.0 + float64(g.CascadeLevel)*g.Config.CascadeMultiplier
			if g.CascadeMultiplier > g.Config.MaxCascadeMultiplier {
				g.CascadeMultiplier = g.Config.MaxCascadeMultiplier
			}
			g.spinOnce()
		}
	}

	g.IsComplete = true
	return nil
}

// spinOnce generates symbols for one spin/cascade
func (g *MegawaysGame) spinOnce() {
	// Determine row count for each reel
	for i := 0; i < g.Config.Reels; i++ {
		rowCount, _ := rand.Int(rand.Reader, big.NewInt(int64(g.Config.MaxRows-g.Config.MinRows+1)))
		g.RowCounts[i] = int(rowCount.Int64()) + g.Config.MinRows
	}

	// Calculate total ways
	g.calculateWays()

	// Generate random symbols for each reel
	g.ReelSymbols = make([][]int, g.Config.Reels)
	for reel := 0; reel < g.Config.Reels; reel++ {
		g.ReelSymbols[reel] = make([]int, g.RowCounts[reel])
		for row := 0; row < g.RowCounts[reel]; row++ {
			symbolIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(g.Symbols))))
			g.ReelSymbols[reel][row] = int(symbolIndex.Int64())
		}
	}

	// Evaluate and add win
	g.evaluate()
}

// calculateWays calculates the total number of megaways
func (g *MegawaysGame) calculateWays() {
	ways := 1
	for _, rows := range g.RowCounts {
		ways *= rows
	}
	g.Ways = ways
}

// hasWinningCombinations checks if there are winning combinations
func (g *MegawaysGame) hasWinningCombinations() bool {
	// Check for any wins in current ReelSymbols
	for reel := 0; reel < g.Config.Reels-1; reel++ {
		for row := 0; row < len(g.ReelSymbols[reel]); row++ {
			symbol := g.ReelSymbols[reel][row]
			if symbol > 1 && !g.Symbols[symbol].IsWild { // Skip wilds and scatter
				// Check if this could form a win
				return true
			}
		}
	}
	return false
}

// evaluate evaluates the spin result
func (g *MegawaysGame) evaluate() {
	g.Win = 0

	// Check all possible ways to win
	// For megaways, we check each possible starting position
	g.checkMegawaysWins()

	// Add to total with cascade multiplier
	if g.Win > 0 {
		cascadeWin := int64(float64(g.Win) * g.CascadeMultiplier)
		g.TotalWin += cascadeWin
		g.CascadeWins = append(g.CascadeWins, cascadeWin)
	}
}

// checkMegawaysWins checks for winning combinations in megaways format
func (g *MegawaysGame) checkMegawaysWins() {
	// Simplified: Check left-to-right wins across reels
	// Find minimum symbol matches across adjacent reels

	minMatches := 2 // Minimum for a win

	// For simplicity, check matching symbols across leftmost reels
	for startReel := 0; startReel < g.Config.Reels-minMatches+1; startReel++ {
		// Check each row in the first reel
		for row := 0; row < g.RowCounts[startReel]; row++ {
			firstSymbol := g.ReelSymbols[startReel][row]

			// Try to find matches in subsequent reels
			matches := 1
			for reel := startReel + 1; reel < g.Config.Reels; reel++ {
				matchFound := false
				for r := 0; r < g.RowCounts[reel]; r++ {
					symbol := g.ReelSymbols[reel][r]
					if symbol == firstSymbol || g.Symbols[symbol].IsWild {
						matches++
						matchFound = true
						break
					}
				}
				if !matchFound {
					break
				}
			}

			if matches >= minMatches {
				// Calculate win
				symbolValue := g.Symbols[firstSymbol].Value
				win := int64(symbolValue) * g.LineBet

				// Multiplier for more matches
				for i := minMatches; i < matches; i++ {
					win *= 2
				}

				g.Win += win
				g.WinLines = append(g.WinLines, WinLine{
					PaylineID: startReel*100 + row,
					Symbol:    g.Symbols[firstSymbol].ID,
					Count:     matches,
					Payout:    win,
				})
			}
		}
	}
}

// GetState returns the game state
func (g *MegawaysGame) GetState() *MegawaysGameState {
	symbols := make([][]string, len(g.ReelSymbols))
	for i, reel := range g.ReelSymbols {
		symbols[i] = make([]string, len(reel))
		for j, s := range reel {
			symbols[i][j] = g.Symbols[s].ID
		}
	}

	return &MegawaysGameState{
		GameID:            g.ID,
		Reels:             g.Config.Reels,
		RowCounts:         g.RowCounts,
		Symbols:           symbols,
		Bet:               g.Bet,
		LineBet:           g.LineBet,
		Ways:              g.Ways,
		Win:               g.Win,
		TotalWin:          g.TotalWin,
		WinLines:          g.WinLines,
		CascadeLevel:      g.CascadeLevel,
		CascadeMultiplier: g.CascadeMultiplier,
		CascadeWins:       g.CascadeWins,
		IsComplete:        g.IsComplete,
		ProvablyFair:      g.ProvablyFair,
	}
}
