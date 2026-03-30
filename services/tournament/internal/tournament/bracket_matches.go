package tournament

// GetBubblePlayer returns the player closest to elimination (short stack)
func (b *TournamentBracket) GetBubblePlayer() *BracketPlayer {
	b.mu.RLock()
	defer b.mu.RUnlock()

	var minChips int = -1
	var bubble *BracketPlayer

	for _, p := range b.Players {
		if p.IsEliminated {
			continue
		}
		if minChips == -1 || p.Chips < minChips {
			minChips = p.Chips
			bubble = p
		}
	}

	return bubble
}

// GetTableCounts returns the number of players per table
func (b *TournamentBracket) GetTableCounts() map[int]int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	counts := make(map[int]int)
	for _, table := range b.Tables {
		count := 0
		for _, p := range table.Seats {
			if !p.IsEliminated {
				count++
			}
		}
		counts[table.TableID] = count
	}

	return counts
}

// IsBalanced checks if tables are balanced (within 1 player of each other)
func (b *TournamentBracket) IsBalanced() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	var counts []int
	for _, table := range b.Tables {
		c := 0
		for _, p := range table.Seats {
			if !p.IsEliminated {
				c++
			}
		}
		counts = append(counts, c)
	}

	if len(counts) == 0 {
		return true
	}

	min := counts[0]
	max := counts[0]
	for _, c := range counts {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}

	return max-min <= 1
}

// GetPlayerPosition returns a player's position in the bracket
func (b *TournamentBracket) GetPlayerPosition(userID string) (tableID, seatNumber int, chips int) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if player, ok := b.Players[userID]; ok {
		return player.TableID, player.SeatNumber, player.Chips
	}

	return -1, -1, 0
}

// GetNextTableAssignment returns the next available table for a new player
func (b *TournamentBracket) GetNextTableAssignment() (tableID, seatNumber int) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	minPlayers := -1
	var targetTable *TableGroup

	for _, table := range b.Tables {
		count := 0
		for _, p := range table.Seats {
			if !p.IsEliminated {
				count++
			}
		}

		if count < table.MaxSeats && (minPlayers == -1 || count < minPlayers) {
			minPlayers = count
			targetTable = table
		}
	}

	if targetTable != nil {
		return targetTable.TableID, len(targetTable.Seats)
	}

	return -1, -1
}
