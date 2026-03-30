package games

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

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
