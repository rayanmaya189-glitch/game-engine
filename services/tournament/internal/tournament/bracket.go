package tournament

import (
	"sort"
	"sync"
)

// TournamentBracket represents a tournament bracket structure for MTT
type TournamentBracket struct {
	TournamentID string
	TotalPlayers int
	Tables       []*TableGroup
	Players      map[string]*BracketPlayer
	mu           sync.RWMutex
}

// TableGroup represents a group of players at a table
type TableGroup struct {
	TableID  int
	Seats    []*BracketPlayer
	MinSeats int
	MaxSeats int
}

// BracketPlayer represents a player in the bracket
type BracketPlayer struct {
	UserID       string
	Username     string
	Chips        int
	Rank         int
	TableID      int
	SeatNumber   int
	IsEliminated bool
	EliminatedAt int // level when eliminated
}

// NewTournamentBracket creates a new tournament bracket
func NewTournamentBracket(tournamentID string, totalPlayers, tablesCount, seatsPerTable int) *TournamentBracket {
	// Calculate number of tables needed
	tablesNeeded := (totalPlayers + seatsPerTable - 1) / seatsPerTable

	if tablesCount > 0 && tablesCount < tablesNeeded {
		tablesNeeded = tablesCount
	}

	bracket := &TournamentBracket{
		TournamentID: tournamentID,
		TotalPlayers: totalPlayers,
		Tables:       make([]*TableGroup, tablesNeeded),
		Players:      make(map[string]*BracketPlayer),
	}

	// Initialize tables
	for i := 0; i < tablesNeeded; i++ {
		bracket.Tables[i] = &TableGroup{
			TableID:  i,
			Seats:    make([]*BracketPlayer, 0, seatsPerTable),
			MinSeats: seatsPerTable / 2, // Minimum for balanced play
			MaxSeats: seatsPerTable,
		}
	}

	return bracket
}

// AddPlayer adds a player to the bracket
func (b *TournamentBracket) AddPlayer(userID, username string, chips int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	player := &BracketPlayer{
		UserID:     userID,
		Username:   username,
		Chips:      chips,
		TableID:    -1,
		SeatNumber: -1,
	}

	b.Players[userID] = player

	// Assign to first available seat
	b.assignSeat(player)
}

// assignSeat assigns a player to an available seat
func (b *TournamentBracket) assignSeat(player *BracketPlayer) {
	for _, table := range b.Tables {
		if len(table.Seats) < table.MaxSeats {
			player.TableID = table.TableID
			player.SeatNumber = len(table.Seats)
			table.Seats = append(table.Seats, player)
			return
		}
	}
}

// UpdateChips updates a player's chip count
func (b *TournamentBracket) UpdateChips(userID string, delta int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if player, ok := b.Players[userID]; ok {
		player.Chips += delta
		if player.Chips <= 0 {
			player.IsEliminated = true
		}
	}
}

// RebalanceTables rebalances players across tables
// Called when players are eliminated to keep table sizes balanced
func (b *TournamentBracket) RebalanceTables() {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Get all non-eliminated players sorted by chips
	var activePlayers []*BracketPlayer
	for _, p := range b.Players {
		if !p.IsEliminated {
			activePlayers = append(activePlayers, p)
		}
	}

	if len(activePlayers) == 0 {
		return
	}

	// Sort by chips (descending)
	sort.Slice(activePlayers, func(i, j int) bool {
		return activePlayers[i].Chips > activePlayers[j].Chips
	})

	// Clear all tables
	for _, table := range b.Tables {
		table.Seats = make([]*BracketPlayer, 0)
	}

	// Redistribute players evenly across tables
	for i, player := range activePlayers {
		tableIdx := i % len(b.Tables)
		player.TableID = tableIdx
		player.SeatNumber = len(b.Tables[tableIdx].Seats)
		b.Tables[tableIdx].Seats = append(b.Tables[tableIdx].Seats, player)
	}
}

// GetTablePlayers returns players at a specific table
func (b *TournamentBracket) GetTablePlayers(tableID int) []*BracketPlayer {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if tableID < 0 || tableID >= len(b.Tables) {
		return nil
	}

	result := make([]*BracketPlayer, len(b.Tables[tableID].Seats))
	copy(result, b.Tables[tableID].Seats)

	return result
}

// GetLeaderboard returns players sorted by chips
func (b *TournamentBracket) GetLeaderboard() []*BracketPlayer {
	b.mu.RLock()
	defer b.mu.RUnlock()

	var players []*BracketPlayer
	for _, p := range b.Players {
		if !p.IsEliminated {
			players = append(players, p)
		}
	}

	sort.Slice(players, func(i, j int) bool {
		if players[i].Chips != players[j].Chips {
			return players[i].Chips > players[j].Chips
		}
		return players[i].UserID < players[j].UserID
	})

	// Update ranks
	for i, p := range players {
		p.Rank = i + 1
	}

	return players
}

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

// GetPlayerCount returns total and active player counts
func (b *TournamentBracket) GetPlayerCount() (total, active, eliminated int) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	total = len(b.Players)
	for _, p := range b.Players {
		if p.IsEliminated {
			eliminated++
		} else {
			active++
		}
	}

	return
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
