package service

import (
	"context"

	"github.com/game-engine/multiplayer/internal/room"
)

// MultiplayerService provides multiplayer game logic
type MultiplayerService struct {
	manager *room.Manager
}

// NewMultiplayerService creates a new multiplayer service
func NewMultiplayerService(manager *room.Manager) (*MultiplayerService, error) {
	return &MultiplayerService{
		manager: manager,
	}, nil
}

// CreateRoom creates a new game room
func (s *MultiplayerService) CreateRoom(ctx context.Context, name, gameType string, tableType room.TableType) (*room.Room, error) {
	return s.manager.CreateRoom(ctx, name, gameType, tableType)
}

// GetRoom gets a room by ID
func (s *MultiplayerService) GetRoom(ctx context.Context, roomID string) (*room.Room, error) {
	return s.manager.GetRoom(ctx, roomID)
}

// ListRooms lists all rooms
func (s *MultiplayerService) ListRooms(ctx context.Context) ([]*room.Room, error) {
	return s.manager.ListRooms(ctx)
}

// CreateTable creates a new table
func (s *MultiplayerService) CreateTable(ctx context.Context, roomID, name, gameType string, tableType room.TableType, minPlayers, maxPlayers int, buyInMin, buyInMax int64, isPrivate bool, password string) (*room.Table, error) {
	return s.manager.CreateTable(ctx, roomID, name, gameType, tableType, minPlayers, maxPlayers, buyInMin, buyInMax, isPrivate, password)
}

// GetTable gets a table by ID
func (s *MultiplayerService) GetTable(ctx context.Context, tableID string) (*room.Table, error) {
	return s.manager.GetTable(ctx, tableID)
}

// JoinTable adds a player to a table
func (s *MultiplayerService) JoinTable(ctx context.Context, tableID, userID, username string, chips int64, password string) (int, error) {
	return s.manager.JoinTable(ctx, tableID, userID, username, chips, password)
}

// LeaveTable removes a player from a table
func (s *MultiplayerService) LeaveTable(ctx context.Context, tableID, userID string) error {
	return s.manager.LeaveTable(ctx, tableID, userID)
}

// SetPlayerReady sets a player's ready status
func (s *MultiplayerService) SetPlayerReady(ctx context.Context, tableID, userID string, ready bool) error {
	return s.manager.SetPlayerReady(ctx, tableID, userID, ready)
}

// AddSpectator adds a spectator
func (s *MultiplayerService) AddSpectator(ctx context.Context, tableID, userID string) error {
	return s.manager.AddSpectator(ctx, tableID, userID)
}

// RemoveSpectator removes a spectator
func (s *MultiplayerService) RemoveSpectator(ctx context.Context, tableID, userID string) error {
	return s.manager.RemoveSpectator(ctx, tableID, userID)
}

// UpdateGameState updates game state
func (s *MultiplayerService) UpdateGameState(ctx context.Context, tableID string, state map[string]interface{}) error {
	return s.manager.UpdateGameState(ctx, tableID, state)
}

// SetCurrentTurn sets the current turn
func (s *MultiplayerService) SetCurrentTurn(ctx context.Context, tableID, userID string) error {
	return s.manager.SetCurrentTurn(ctx, tableID, userID)
}

// CloseTable closes a table
func (s *MultiplayerService) CloseTable(ctx context.Context, tableID string) error {
	return s.manager.CloseTable(ctx, tableID)
}

// ListTables lists tables
func (s *MultiplayerService) ListTables(ctx context.Context, roomID, gameType string, includePrivate bool) ([]*room.Table, error) {
	return s.manager.ListTables(ctx, roomID, gameType, includePrivate)
}
