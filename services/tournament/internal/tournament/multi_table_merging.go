package tournament

import (
	"context"
	"fmt"
	"time"
)

// createFinalTable creates the final table when enough players remain
func (m *MultiTableTournament) createFinalTable(ctx context.Context) {
	m.FinalTable = &MTTTable{
		TableID:     "final_table",
		TableNumber: len(m.Tables) + 1,
		Players:     make(map[string]*MTTPlayer),
		SeatCount:   m.Settings.FinalTableSeats,
		Level:       0,
		IsFinal:     true,
	}

	// Move top players to final table
	var topPlayers []*MTTPlayer
	for _, table := range m.Tables {
		for _, player := range table.Players {
			if !player.IsEliminated {
				topPlayers = append(topPlayers, player)
			}
		}
	}

	// Sort by chips (highest first) - simplified
	// In production, use proper sorting
	if len(topPlayers) > m.Settings.FinalTableSeats {
		topPlayers = topPlayers[:m.Settings.FinalTableSeats]
	}

	for _, player := range topPlayers {
		m.FinalTable.Players[player.UserID] = player
		player.TableID = "final_table"
		player.Position = len(m.FinalTable.Players)
	}
}

// countActivePlayers counts players still in the tournament
func (m *MultiTableTournament) countActivePlayers() int {
	count := 0
	for _, table := range m.Tables {
		for _, player := range table.Players {
			if !player.IsEliminated {
				count++
			}
		}
	}
	if m.FinalTable != nil {
		for _, player := range m.FinalTable.Players {
			if !player.IsEliminated {
				count++
			}
		}
	}
	return count
}

// EliminatePlayer removes a player from the tournament
func (m *MultiTableTournament) EliminatePlayer(ctx context.Context, userID string, rank int, prize int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Find player
	var player *MTTPlayer
	for _, table := range m.Tables {
		if p, exists := table.Players[userID]; exists {
			player = p
			break
		}
	}
	if m.FinalTable != nil {
		if p, exists := m.FinalTable.Players[userID]; exists {
			player = p
		}
	}

	if player == nil {
		return fmt.Errorf("player not found")
	}

	player.IsEliminated = true
	player.Rank = rank
	player.PrizeMoney = prize

	// Add to results
	m.Results = append(m.Results, MTTResult{
		Rank:       rank,
		UserID:     userID,
		PrizeMoney: prize,
	})

	m.UpdatedAt = time.Now()
	return m.saveToRedis(ctx)
}

// completeTournament finishes the tournament and calculates prizes
func (m *MultiTableTournament) completeTournament(ctx context.Context) {
	m.Status = TournamentStatusCompleted
	m.EndTime = time.Now()

	// Calculate prize distribution
	prizeDistribution := m.calculatePrizeDistribution()

	// Assign final positions and prizes
	rank := 1
	for _, table := range m.Tables {
		for _, player := range table.Players {
			if !player.IsEliminated {
				player.Rank = rank
				if rank-1 < len(prizeDistribution) {
					player.PrizeMoney = prizeDistribution[rank-1]
				}
				m.Results = append(m.Results, MTTResult{
					Rank:       rank,
					UserID:     player.UserID,
					PrizeMoney: player.PrizeMoney,
				})
				rank++
			}
		}
	}

	// Final table results
	if m.FinalTable != nil {
		for _, player := range m.FinalTable.Players {
			if !player.IsEliminated {
				player.Rank = rank
				if rank-1 < len(prizeDistribution) {
					player.PrizeMoney = prizeDistribution[rank-1]
				}
				m.Results = append(m.Results, MTTResult{
					Rank:       rank,
					UserID:     player.UserID,
					PrizeMoney: player.PrizeMoney,
				})
				rank++
			}
		}
	}
}

// calculatePrizeDistribution calculates prize amounts based on positions
func (m *MultiTableTournament) calculatePrizeDistribution() []int64 {
	// Standard distribution: top 15-20% of players get paid
	playerCount := m.countActivePlayers()
	if playerCount == 0 {
		return []int64{}
	}

	paidPlaces := playerCount / 5 // 20% get paid
	if paidPlaces < 1 {
		paidPlaces = 1
	}
	if paidPlaces > 10 {
		paidPlaces = 10
	}

	distribution := make([]int64, paidPlaces)

	// Tiered distribution (more to first place)
	for i := 0; i < paidPlaces; i++ {
		// Simple exponential distribution
		percentage := float64(paidPlaces - i)
		distribution[i] = int64(float64(m.PrizePool) * (percentage / float64(paidPlaces*paidPlaces/2)))
	}

	return distribution
}
