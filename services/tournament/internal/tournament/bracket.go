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
	EliminatedAt int
}

// NewTournamentBracket creates a new tournament bracket
func NewTournamentBracket(tournamentID string, totalPlayers, tablesCount, seatsPerTable int) *TournamentBracket {
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

	for i := 0; i < tablesNeeded; i++ {
		bracket.Tables[i] = &TableGroup{
			TableID:  i,
			Seats:    make([]*BracketPlayer, 0, seatsPerTable),
			MinSeats: seatsPerTable / 2,
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
func (b *TournamentBracket) RebalanceTables() {
	b.mu.Lock()
	defer b.mu.Unlock()

	var activePlayers []*BracketPlayer
	for _, p := range b.Players {
		if !p.IsEliminated {
			activePlayers = append(activePlayers, p)
		}
	}

	if len(activePlayers) == 0 {
		return
	}

	sort.Slice(activePlayers, func(i, j int) bool {
		return activePlayers[i].Chips > activePlayers[j].Chips
	})

	for _, table := range b.Tables {
		table.Seats = make([]*BracketPlayer, 0)
	}

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

	for i, p := range players {
		p.Rank = i + 1
	}

	return players
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
