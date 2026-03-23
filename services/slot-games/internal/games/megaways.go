package games

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// MegawaysConfig represents Megaways slot game configuration
type MegawaysConfig struct {
	Reels                int     `json:"reels"`
	MinRows              int     `json:"min_rows"` // Minimum symbols per reel
	MaxRows              int     `json:"max_rows"` // Maximum symbols per reel
	MinBet               int64   `json:"min_bet"`
	MaxBet               int64   `json:"max_bet"`
	MinLineBet           int64   `json:"min_line_bet"`
	MaxLineBet           int64   `json:"max_line_bet"`
	RTP                  float64 `json:"rtp"`                    // Return to player
	Volatility           string  `json:"volatility"`             // low, medium, high
	MaxMegaways          int     `json:"max_megaways"`           // Maximum ways to win (default 117649)
	CascadeEnabled       bool    `json:"cascade_enabled"`        // Avalanche/cascade feature
	CascadeMultiplier    float64 `json:"cascade_multiplier"`     // Multiplier increase per cascade
	MaxCascadeMultiplier float64 `json:"max_cascade_multiplier"` // Maximum cascade multiplier
}

// MegawaysGame represents a Megaways slot game
type MegawaysGame struct {
	ID                string
	Config            *MegawaysConfig
	Symbols           []Symbol
	ReelSymbols       [][]int // [reel][row] - variable height
	RowCounts         []int   // Number of rows for each reel
	Bet               int64
	LineBet           int64
	Ways              int // Current number of megaways
	Win               int64
	CascadeLevel      int
	CascadeMultiplier float64
	WinLines          []WinLine
	TotalWin          int64
	IsComplete        bool
	ProvablyFair      bool
	ServerSeed        string
	ClientSeed        string
	Nonce             int
	CascadeWins       []int64 // Wins from each cascade level
}

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

// MegawaysGameState represents the game state
type MegawaysGameState struct {
	GameID            string     `json:"game_id"`
	Reels             int        `json:"reels"`
	RowCounts         []int      `json:"row_counts"`
	Symbols           [][]string `json:"symbols"`
	Bet               int64      `json:"bet"`
	LineBet           int64      `json:"line_bet"`
	Ways              int        `json:"ways"`
	Win               int64      `json:"win"`
	TotalWin          int64      `json:"total_win"`
	WinLines          []WinLine  `json:"win_lines"`
	CascadeLevel      int        `json:"cascade_level"`
	CascadeMultiplier float64    `json:"cascade_multiplier"`
	CascadeWins       []int64    `json:"cascade_wins"`
	IsComplete        bool       `json:"is_complete"`
	ProvablyFair      bool       `json:"provably_fair"`
}

// ClusterGame represents a Cluster Pays slot game
type ClusterGame struct {
	ID           string
	Config       *ClusterConfig
	Symbols      []Symbol
	Grid         [][]int // [row][col]
	Bet          int64
	Win          int64
	WinClusters  []Cluster
	TotalWin     int64
	IsComplete   bool
	ProvablyFair bool
	ServerSeed   string
	ClientSeed   string
	Nonce        int
}

// ClusterConfig represents Cluster game configuration
type ClusterConfig struct {
	Rows       int     `json:"rows"`
	Cols       int     `json:"cols"`
	MinBet     int64   `json:"min_bet"`
	MaxBet     int64   `json:"max_bet"`
	RTP        float64 `json:"rtp"`
	Volatility string  `json:"volatility"`
	MinCluster int     `json:"min_cluster"` // Minimum symbols for a cluster win
	Cascade    bool    `json:"cascade"`
}

// Cluster represents a winning cluster
type Cluster struct {
	Symbol    string     `json:"symbol"`
	Count     int        `json:"count"`
	Payout    int64      `json:"payout"`
	Positions []Position `json:"positions"`
}

// Position represents a grid position
type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// NewClusterGame creates a new Cluster Pays slot game
func NewClusterGame(id string) *ClusterGame {
	config := &ClusterConfig{
		Rows:       6,
		Cols:       6,
		MinBet:     1,
		MaxBet:     500,
		RTP:        96.5,
		Volatility: "high",
		MinCluster: 5,
		Cascade:    true,
	}

	symbols := []Symbol{
		{ID: "W", Name: "Wild", Value: 10, IsWild: true, Weight: 3},
		{ID: "S", Name: "Scatter", Value: 9, IsScatter: true, Weight: 6},
		{ID: "L", Name: "Lightning", Value: 8, Weight: 8},
		{ID: "H", Name: "Heart", Value: 7, Weight: 10},
		{ID: "D", Name: "Diamond", Value: 6, Weight: 12},
		{ID: "S", Name: "Spade", Value: 5, Weight: 14},
		{ID: "C", Name: "Club", Value: 4, Weight: 16},
		{ID: "B", Name: "Blue", Value: 3, Weight: 18},
		{ID: "G", Name: "Green", Value: 2, Weight: 20},
		{ID: "R", Name: "Red", Value: 1, Weight: 22},
	}

	return &ClusterGame{
		ID:          id,
		Config:      config,
		Symbols:     symbols,
		Grid:        make([][]int, config.Rows),
		WinClusters: make([]Cluster, 0),
	}
}

// SetBet sets the bet amount
func (g *ClusterGame) SetBet(bet int64) error {
	if bet < g.Config.MinBet {
		return fmt.Errorf("minimum bet is %d", g.Config.MinBet)
	}
	if bet > g.Config.MaxBet {
		return fmt.Errorf("maximum bet is %d", g.Config.MaxBet)
	}
	g.Bet = bet
	return nil
}

// Spin performs a spin
func (g *ClusterGame) Spin() error {
	if g.Bet < g.Config.MinBet {
		return fmt.Errorf("bet not set")
	}

	g.Win = 0
	g.WinClusters = nil
	g.TotalWin = 0

	// Generate random grid
	g.generateGrid()

	// Evaluate clusters
	g.evaluateClusters()

	g.IsComplete = true
	return nil
}

// generateGrid generates random symbols for the grid
func (g *ClusterGame) generateGrid() {
	g.Grid = make([][]int, g.Config.Rows)
	for row := 0; row < g.Config.Rows; row++ {
		g.Grid[row] = make([]int, g.Config.Cols)
		for col := 0; col < g.Config.Cols; col++ {
			symbolIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(g.Symbols))))
			g.Grid[row][col] = int(symbolIndex.Int64())
		}
	}
}

// evaluateClusters finds and scores winning clusters
func (g *ClusterGame) evaluateClusters() {
	visited := make(map[string]bool)

	for row := 0; row < g.Config.Rows; row++ {
		for col := 0; col < g.Config.Cols; col++ {
			key := fmt.Sprintf("%d,%d", row, col)
			if visited[key] {
				continue
			}

			symbol := g.Grid[row][col]
			if g.Symbols[symbol].IsWild || g.Symbols[symbol].IsScatter {
				continue
			}

			// Find connected symbols
			cluster := g.findCluster(row, col, symbol, visited)
			if len(cluster) >= g.Config.MinCluster {
				payout := g.calculateClusterPayout(symbol, len(cluster))
				g.Win += payout
				g.WinClusters = append(g.WinClusters, Cluster{
					Symbol:    g.Symbols[symbol].ID,
					Count:     len(cluster),
					Payout:    payout,
					Positions: cluster,
				})
			}
		}
	}

	g.TotalWin = g.Win * g.Bet
}

// findCluster finds all connected positions with the same symbol
func (g *ClusterGame) findCluster(row, col, symbol int, visited map[string]bool) []Position {
	var cluster []Position
	queue := []Position{{Row: row, Col: col}}
	visited[fmt.Sprintf("%d,%d", row, col)] = true

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		cluster = append(cluster, pos)

		// Check neighbors (4-directional)
		neighbors := []Position{
			{pos.Row - 1, pos.Col},
			{pos.Row + 1, pos.Col},
			{pos.Row, pos.Col - 1},
			{pos.Row, pos.Col + 1},
		}

		for _, n := range neighbors {
			if n.Row < 0 || n.Row >= g.Config.Rows || n.Col < 0 || n.Col >= g.Config.Cols {
				continue
			}

			key := fmt.Sprintf("%d,%d", n.Row, n.Col)
			if visited[key] {
				continue
			}

			if g.Grid[n.Row][n.Col] == symbol || g.Symbols[g.Grid[n.Row][n.Col]].IsWild {
				visited[key] = true
				queue = append(queue, n)
			}
		}
	}

	return cluster
}

// calculateClusterPayout calculates payout for a cluster
func (g *ClusterGame) calculateClusterPayout(symbolIndex, count int) int64 {
	if count < g.Config.MinCluster {
		return 0
	}

	baseValue := int64(g.Symbols[symbolIndex].Value)

	// Progressive multiplier based on cluster size
	multiplier := int64(1)
	for i := g.Config.MinCluster; i < count; i++ {
		multiplier *= 2
	}

	return baseValue * g.Bet * multiplier
}

// GetState returns the game state
func (g *ClusterGame) GetState() *ClusterGameState {
	symbolGrid := make([][]string, g.Config.Rows)
	for row := 0; row < g.Config.Rows; row++ {
		symbolGrid[row] = make([]string, g.Config.Cols)
		for col := 0; col < g.Config.Cols; col++ {
			symbolGrid[row][col] = g.Symbols[g.Grid[row][col]].ID
		}
	}

	return &ClusterGameState{
		GameID:     g.ID,
		Rows:       g.Config.Rows,
		Cols:       g.Config.Cols,
		Symbols:    symbolGrid,
		Bet:        g.Bet,
		Win:        g.Win,
		TotalWin:   g.TotalWin,
		Clusters:   g.WinClusters,
		IsComplete: g.IsComplete,
	}
}

// ClusterGameState represents the game state
type ClusterGameState struct {
	GameID     string     `json:"game_id"`
	Rows       int        `json:"rows"`
	Cols       int        `json:"cols"`
	Symbols    [][]string `json:"symbols"`
	Bet        int64      `json:"bet"`
	Win        int64      `json:"win"`
	TotalWin   int64      `json:"total_win"`
	Clusters   []Cluster  `json:"clusters"`
	IsComplete bool       `json:"is_complete"`
}

// ProgressiveGame represents a Progressive Jackpot slot game
type ProgressiveGame struct {
	ID               string
	GameType         string
	Config           *ProgressiveConfig
	Symbols          []Symbol
	Positions        [][]int
	Bet              int64
	LineBet          int64
	Win              int64
	TotalWin         int64
	JackpotTriggered bool
	JackpotTier      string
	JackpotWin       int64
	IsComplete       bool
}

// ProgressiveConfig represents Progressive game configuration
type ProgressiveConfig struct {
	Reels      int                           `json:"reels"`
	Rows       int                           `json:"rows"`
	MinBet     int64                         `json:"min_bet"`
	MaxBet     int64                         `json:"max_bet"`
	RTP        float64                       `json:"rtp"`
	Volatility string                        `json:"volatility"`
	Jackpots   map[string]ProgressiveJackpot `json:"jackpots"`
}

// ProgressiveJackpot represents a progressive jackpot tier
type ProgressiveJackpot struct {
	Name             string   `json:"name"`
	SeedAmount       int64    `json:"seed_amount"`
	ContributionRate float64  `json:"contribution_rate"` // Percentage of bet
	MinBet           int64    `json:"min_bet"`
	Trigger          string   `json:"trigger"` // "symbol", "random", "feature"
	Symbols          []string `json:"symbols,omitempty"`
	Odds             float64  `json:"odds"` // Probability of winning
}

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

// ProgressiveGameState represents the game state
type ProgressiveGameState struct {
	GameID           string     `json:"game_id"`
	GameType         string     `json:"game_type"`
	Reels            int        `json:"reels"`
	Rows             int        `json:"rows"`
	Symbols          [][]string `json:"symbols"`
	Bet              int64      `json:"bet"`
	Win              int64      `json:"win"`
	TotalWin         int64      `json:"total_win"`
	JackpotTriggered bool       `json:"jackpot_triggered"`
	JackpotTier      string     `json:"jackpot_tier,omitempty"`
	JackpotWin       int64      `json:"jackpot_win,omitempty"`
	IsComplete       bool       `json:"is_complete"`
}

// Seed hash for provably fair
func megawaysSeedHash(seed string) string {
	if seed == "" {
		return ""
	}
	h := sha256.Sum256([]byte(seed))
	return fmt.Sprintf("%x", h[:8])
}
