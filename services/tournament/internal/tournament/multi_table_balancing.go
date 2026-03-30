package tournament

import "context"

// balanceTables balances players across tables
func (m *MultiTableTournament) balanceTables(ctx context.Context) {
	var tablesToBalance []*MTTTable
	minPlayers := m.Settings.SeatsPerTable

	for _, table := range m.Tables {
		if table.IsFinal {
			continue
		}
		if len(table.Players) < minPlayers {
			tablesToBalance = append(tablesToBalance, table)
		}
	}

	// Move players from tables with excess to tables with deficit
	for _, table := range m.Tables {
		if table.IsFinal || len(table.Players) <= minPlayers {
			continue
		}

		for _, player := range table.Players {
			if player.IsEliminated {
				continue
			}

			// Find a table with fewer players
			for _, targetTable := range tablesToBalance {
				if len(targetTable.Players) < targetTable.SeatCount {
					// Move player
					delete(table.Players, player.UserID)
					player.TableID = targetTable.TableID
					targetTable.Players[player.UserID] = player
					break
				}
			}
		}
	}
}
