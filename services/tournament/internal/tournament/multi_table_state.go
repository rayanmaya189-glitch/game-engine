package tournament

import (
	"context"
	"fmt"
	"time"
)

// GetTableStates returns current state of all tables
func (m *MultiTableTournament) GetTableStates() map[string]interface{} {
	result := make(map[string]interface{})

	for tableID, table := range m.Tables {
		playerList := make([]string, 0, len(table.Players))
		for userID, player := range table.Players {
			if !player.IsEliminated {
				playerList = append(playerList, fmt.Sprintf("%s:%d", userID, player.Chips))
			}
		}
		result[tableID] = map[string]interface{}{
			"table_number": table.TableNumber,
			"players":      playerList,
			"level":        table.Level,
		}
	}

	return result
}

// GetLeaderboard returns current tournament standings
func (m *MultiTableTournament) GetLeaderboard() []map[string]interface{} {
	type playerScore struct {
		userID string
		chips  int
		rank   int
	}

	var players []playerScore
	for _, table := range m.Tables {
		for userID, player := range table.Players {
			players = append(players, playerScore{
				userID: userID,
				chips:  player.Chips,
				rank:   player.Rank,
			})
		}
	}

	// Sort by chips descending (simplified)
	// In production, use proper sorting

	result := make([]map[string]interface{}, 0, len(players))
	for _, p := range players {
		result = append(result, map[string]interface{}{
			"user_id": p.userID,
			"chips":   p.chips,
			"rank":    p.rank,
		})
	}

	return result
}

// GetPlayerInfo returns information about a specific player
func (m *MultiTableTournament) GetPlayerInfo(userID string) (map[string]interface{}, error) {
	for _, table := range m.Tables {
		if player, exists := table.Players[userID]; exists {
			return map[string]interface{}{
				"user_id":    player.UserID,
				"chips":      player.Chips,
				"rank":       player.Rank,
				"table_id":   player.TableID,
				"eliminated": player.IsEliminated,
				"prize":      player.PrizeMoney,
			}, nil
		}
	}
	if m.FinalTable != nil {
		if player, exists := m.FinalTable.Players[userID]; exists {
			return map[string]interface{}{
				"user_id":    player.UserID,
				"chips":      player.Chips,
				"rank":       player.Rank,
				"table_id":   player.TableID,
				"position":   player.Position,
				"eliminated": player.IsEliminated,
				"prize":      player.PrizeMoney,
			}, nil
		}
	}
	return nil, fmt.Errorf("player not found")
}

// AddChips adds chips to a player (for rebuys/addons)
func (m *MultiTableTournament) AddChips(ctx context.Context, userID string, amount int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, table := range m.Tables {
		if player, exists := table.Players[userID]; exists {
			player.Chips += amount
			m.UpdatedAt = time.Now()
			return m.saveToRedis(ctx)
		}
	}

	if m.FinalTable != nil {
		if player, exists := m.FinalTable.Players[userID]; exists {
			player.Chips += amount
			m.UpdatedAt = time.Now()
			return m.saveToRedis(ctx)
		}
	}

	return fmt.Errorf("player not found")
}

// UpdateChips updates a player's chip count
func (m *MultiTableTournament) UpdateChips(ctx context.Context, userID string, newChips int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if newChips <= 0 {
		// Player eliminated - find their rank
		rank := m.countActivePlayers()
		prize := int64(0)
		for _, r := range m.Results {
			if r.Rank == rank {
				prize = r.PrizeMoney
				break
			}
		}
		return m.EliminatePlayer(ctx, userID, rank, prize)
	}

	for _, table := range m.Tables {
		if player, exists := table.Players[userID]; exists {
			player.Chips = newChips
			m.UpdatedAt = time.Now()
			return m.saveToRedis(ctx)
		}
	}

	return fmt.Errorf("player not found")
}
