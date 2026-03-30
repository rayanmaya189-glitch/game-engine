package room

import (
	"context"
	"fmt"
	"time"
)

// JoinTable adds a player to a table
func (m *Manager) JoinTable(ctx context.Context, tableID string, userID, username string, chips int64, password string) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	table, ok := m.tables[tableID]
	if !ok {
		return -1, fmt.Errorf("table not found: %s", tableID)
	}

	if table.Private && table.Password != password {
		return -1, fmt.Errorf("invalid password")
	}

	// Find an empty seat
	seatIndex := -1
	for i, seat := range table.Seats {
		if seat.UserID == "" {
			seatIndex = i
			break
		}
	}

	if seatIndex == -1 {
		return -1, fmt.Errorf("table is full")
	}

	// Check buy-in limits
	if chips < table.BuyInMin {
		return -1, fmt.Errorf("buy-in below minimum: %d", table.BuyInMin)
	}
	if table.BuyInMax > 0 && chips > table.BuyInMax {
		return -1, fmt.Errorf("buy-in above maximum: %d", table.BuyInMax)
	}

	// Assign seat
	seat := table.Seats[seatIndex]
	seat.UserID = userID
	seat.Username = username
	seat.Chips = chips
	seat.Ready = false
	seat.Connected = true
	table.Seats[seatIndex] = seat
	table.UpdatedAt = time.Now()

	// Save to Redis
	if err := m.saveTable(ctx, table); err != nil {
		return -1, err
	}

	return seatIndex, nil
}

// LeaveTable removes a player from a table
func (m *Manager) LeaveTable(ctx context.Context, tableID, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	table, ok := m.tables[tableID]
	if !ok {
		return fmt.Errorf("table not found: %s", tableID)
	}

	// Find the player's seat
	for i, seat := range table.Seats {
		if seat.UserID == userID {
			// Clear seat
			table.Seats[i] = &Seat{
				ID:        i,
				Connected: false,
			}
			break
		}
	}

	table.UpdatedAt = time.Now()

	// Save to Redis
	return m.saveTable(ctx, table)
}

// SetPlayerReady sets a player's ready status
func (m *Manager) SetPlayerReady(ctx context.Context, tableID, userID string, ready bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	table, ok := m.tables[tableID]
	if !ok {
		return fmt.Errorf("table not found: %s", tableID)
	}

	for _, seat := range table.Seats {
		if seat.UserID == userID {
			seat.Ready = ready
			break
		}
	}

	table.UpdatedAt = time.Now()

	// Check if all players are ready
	if m.allPlayersReady(table) {
		table.Status = TableStatusPlaying
	}

	return m.saveTable(ctx, table)
}

// allPlayersReady checks if all active players are ready
func (m *Manager) allPlayersReady(table *Table) bool {
	activePlayers := 0
	readyPlayers := 0

	for _, seat := range table.Seats {
		if seat.UserID != "" {
			activePlayers++
			if seat.Ready {
				readyPlayers++
			}
		}
	}

	return activePlayers >= table.MinPlayers && activePlayers == readyPlayers
}

// AddSpectator adds a spectator to a table
func (m *Manager) AddSpectator(ctx context.Context, tableID, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	table, ok := m.tables[tableID]
	if !ok {
		return fmt.Errorf("table not found: %tableID")
	}

	if len(table.Spectators) >= m.config.Multiplayer.MaxSpectators {
		return fmt.Errorf("max spectators reached")
	}

	table.Spectators[userID] = true
	table.UpdatedAt = time.Now()

	return m.saveTable(ctx, table)
}

// RemoveSpectator removes a spectator from a table
func (m *Manager) RemoveSpectator(ctx context.Context, tableID, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	table, ok := m.tables[tableID]
	if !ok {
		return fmt.Errorf("table not found: %s", tableID)
	}

	delete(table.Spectators, userID)
	table.UpdatedAt = time.Now()

	return m.saveTable(ctx, table)
}
