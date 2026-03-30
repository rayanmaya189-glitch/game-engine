package room

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// CreateTable creates a new table in a room
func (m *Manager) CreateTable(ctx context.Context, roomID, name, gameType string, tableType TableType, minPlayers, maxPlayers int, buyInMin, buyInMax int64, isPrivate bool, password string) (*Table, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	room, ok := m.rooms[roomID]
	if !ok {
		return nil, fmt.Errorf("room not found: %s", roomID)
	}

	table := &Table{
		ID:         uuid.New().String(),
		RoomID:     roomID,
		Name:       name,
		GameType:   gameType,
		Type:       tableType,
		Status:     TableStatusWaiting,
		MinPlayers: minPlayers,
		MaxPlayers: maxPlayers,
		BuyInMin:   buyInMin,
		BuyInMax:   buyInMax,
		Seats:      make(map[int]*Seat),
		Spectators: make(map[string]bool),
		GameState:  make(map[string]interface{}),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Private:    isPrivate,
		Password:   password,
	}

	// Initialize seats
	for i := 0; i < maxPlayers; i++ {
		table.Seats[i] = &Seat{
			ID:        i,
			Connected: false,
		}
	}

	room.Tables[table.ID] = table
	m.tables[table.ID] = table
	room.UpdatedAt = time.Now()

	// Save to Redis
	if err := m.saveTable(ctx, table); err != nil {
		return nil, err
	}

	return table, nil
}

// GetTable retrieves a table by ID
func (m *Manager) GetTable(ctx context.Context, tableID string) (*Table, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	table, ok := m.tables[tableID]
	if !ok {
		// Try to load from Redis
		data, err := m.redisClient.Get(ctx, fmt.Sprintf("table:%s", tableID)).Bytes()
		if err != nil {
			return nil, fmt.Errorf("table not found: %s", tableID)
		}

		var t Table
		if err := json.Unmarshal(data, &t); err != nil {
			return nil, err
		}
		return &t, nil
	}

	return table, nil
}

// ListTables lists available tables
func (m *Manager) ListTables(ctx context.Context, roomID string, gameType string, includePrivate bool) ([]*Table, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result []*Table

	for _, table := range m.tables {
		if roomID != "" && table.RoomID != roomID {
			continue
		}
		if gameType != "" && table.GameType != gameType {
			continue
		}
		if !includePrivate && table.Private {
			continue
		}
		if table.Status != TableStatusClosed {
			result = append(result, table)
		}
	}

	return result, nil
}

// CloseTable closes a table
func (m *Manager) CloseTable(ctx context.Context, tableID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	table, ok := m.tables[tableID]
	if !ok {
		return fmt.Errorf("table not found: %s", tableID)
	}

	table.Status = TableStatusClosed
	table.UpdatedAt = time.Now()

	// Remove from room
	if room, ok := m.rooms[table.RoomID]; ok {
		delete(room.Tables, tableID)
		room.UpdatedAt = time.Now()
	}

	delete(m.tables, tableID)

	// Save to Redis
	return m.saveTable(ctx, table)
}

// saveTable saves a table to Redis
func (m *Manager) saveTable(ctx context.Context, table *Table) error {
	data, err := json.Marshal(table)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("table:%s", table.ID)
	return m.redisClient.Set(ctx, key, data, 0).Err()
}
