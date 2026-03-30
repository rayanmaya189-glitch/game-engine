package room

import (
	"context"
	"fmt"
	"time"
)

// UpdateGameState updates the game state for a table
func (m *Manager) UpdateGameState(ctx context.Context, tableID string, state map[string]interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	table, ok := m.tables[tableID]
	if !ok {
		return fmt.Errorf("table not found: %s", tableID)
	}

	table.GameState = state
	table.UpdatedAt = time.Now()

	return m.saveTable(ctx, table)
}

// SetCurrentTurn sets the current turn player
func (m *Manager) SetCurrentTurn(ctx context.Context, tableID, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	table, ok := m.tables[tableID]
	if !ok {
		return fmt.Errorf("table not found: %s", tableID)
	}

	table.CurrentTurn = userID
	table.UpdatedAt = time.Now()

	return m.saveTable(ctx, table)
}
