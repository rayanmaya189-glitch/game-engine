package handler

import (
	"context"

	"github.com/game-engine/multiplayer/internal/room"
	"github.com/game-engine/multiplayer/internal/service"
)

// MultiplayerHandler handles multiplayer-related requests
type MultiplayerHandler struct {
	service *service.MultiplayerService
}

// NewMultiplayerHandler creates a new multiplayer handler
func NewMultiplayerHandler(service *service.MultiplayerService) *MultiplayerHandler {
	return &MultiplayerHandler{
		service: service,
	}
}

// CreateRoomRequest represents a create room request
type CreateRoomRequest struct {
	Name     string         `json:"name"`
	GameType string         `json:"game_type"`
	Type     room.TableType `json:"type"`
}

// CreateTableRequest represents a create table request
type CreateTableRequest struct {
	RoomID     string         `json:"room_id"`
	Name       string         `json:"name"`
	GameType   string         `json:"game_type"`
	Type       room.TableType `json:"type"`
	MinPlayers int            `json:"min_players"`
	MaxPlayers int            `json:"max_players"`
	BuyInMin   int64          `json:"buy_in_min"`
	BuyInMax   int64          `json:"buy_in_max"`
	Private    bool           `json:"private"`
	Password   string         `json:"password"`
}

// JoinTableRequest represents a join table request
type JoinTableRequest struct {
	TableID  string `json:"table_id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Chips    int64  `json:"chips"`
	Password string `json:"password"`
}

// ReadyRequest represents a ready request
type ReadyRequest struct {
	TableID string `json:"table_id"`
	UserID  string `json:"user_id"`
	Ready   bool   `json:"ready"`
}

// RoomResponse represents a room response
type RoomResponse struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	GameType   string         `json:"game_type"`
	Type       room.TableType `json:"type"`
	TableCount int            `json:"table_count"`
}

// TableResponse represents a table response
type TableResponse struct {
	ID          string           `json:"id"`
	RoomID      string           `json:"room_id"`
	Name        string           `json:"name"`
	GameType    string           `json:"game_type"`
	Type        room.TableType   `json:"type"`
	Status      room.TableStatus `json:"status"`
	MinPlayers  int              `json:"min_players"`
	MaxPlayers  int              `json:"max_players"`
	PlayerCount int              `json:"player_count"`
	Spectators  int              `json:"spectators"`
	Private     bool             `json:"private"`
}

// CreateRoom creates a new room
func (h *MultiplayerHandler) CreateRoom(ctx context.Context, req *CreateRoomRequest) (*RoomResponse, error) {
	room, err := h.service.CreateRoom(ctx, req.Name, req.GameType, req.Type)
	if err != nil {
		return nil, err
	}

	return &RoomResponse{
		ID:         room.ID,
		Name:       room.Name,
		GameType:   room.GameType,
		Type:       room.TableType,
		TableCount: len(room.Tables),
	}, nil
}

// GetRoom gets a room
func (h *MultiplayerHandler) GetRoom(ctx context.Context, roomID string) (*RoomResponse, error) {
	room, err := h.service.GetRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	return &RoomResponse{
		ID:         room.ID,
		Name:       room.Name,
		GameType:   room.GameType,
		Type:       room.TableType,
		TableCount: len(room.Tables),
	}, nil
}

// ListRooms lists rooms
func (h *MultiplayerHandler) ListRooms(ctx context.Context) ([]*RoomResponse, error) {
	rooms, err := h.service.ListRooms(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*RoomResponse, 0, len(rooms))
	for _, r := range rooms {
		result = append(result, &RoomResponse{
			ID:         r.ID,
			Name:       r.Name,
			GameType:   r.GameType,
			Type:       r.TableType,
			TableCount: len(r.Tables),
		})
	}

	return result, nil
}

// CreateTable creates a new table
func (h *MultiplayerHandler) CreateTable(ctx context.Context, req *CreateTableRequest) (*TableResponse, error) {
	table, err := h.service.CreateTable(ctx, req.RoomID, req.Name, req.GameType, req.Type, req.MinPlayers, req.MaxPlayers, req.BuyInMin, req.BuyInMax, req.Private, req.Password)
	if err != nil {
		return nil, err
	}

	return h.tableToResponse(table), nil
}

// GetTable gets a table
func (h *MultiplayerHandler) GetTable(ctx context.Context, tableID string) (*TableResponse, error) {
	table, err := h.service.GetTable(ctx, tableID)
	if err != nil {
		return nil, err
	}

	return h.tableToResponse(table), nil
}

// JoinTable joins a table
func (h *MultiplayerHandler) JoinTable(ctx context.Context, req *JoinTableRequest) (int, error) {
	return h.service.JoinTable(ctx, req.TableID, req.UserID, req.Username, req.Chips, req.Password)
}

// LeaveTable leaves a table
func (h *MultiplayerHandler) LeaveTable(ctx context.Context, tableID, userID string) error {
	return h.service.LeaveTable(ctx, tableID, userID)
}

// SetPlayerReady sets player ready status
func (h *MultiplayerHandler) SetPlayerReady(ctx context.Context, req *ReadyRequest) error {
	return h.service.SetPlayerReady(ctx, req.TableID, req.UserID, req.Ready)
}

// AddSpectator adds a spectator
func (h *MultiplayerHandler) AddSpectator(ctx context.Context, tableID, userID string) error {
	return h.service.AddSpectator(ctx, tableID, userID)
}

// RemoveSpectator removes a spectator
func (h *MultiplayerHandler) RemoveSpectator(ctx context.Context, tableID, userID string) error {
	return h.service.RemoveSpectator(ctx, tableID, userID)
}

// ListTables lists tables
func (h *MultiplayerHandler) ListTables(ctx context.Context, roomID, gameType string, includePrivate bool) ([]*TableResponse, error) {
	tables, err := h.service.ListTables(ctx, roomID, gameType, includePrivate)
	if err != nil {
		return nil, err
	}

	result := make([]*TableResponse, 0, len(tables))
	for _, t := range tables {
		result = append(result, h.tableToResponse(t))
	}

	return result, nil
}

// CloseTable closes a table
func (h *MultiplayerHandler) CloseTable(ctx context.Context, tableID string) error {
	return h.service.CloseTable(ctx, tableID)
}

func (h *MultiplayerHandler) tableToResponse(table *room.Table) *TableResponse {
	playerCount := 0
	for _, seat := range table.Seats {
		if seat.UserID != "" {
			playerCount++
		}
	}

	return &TableResponse{
		ID:          table.ID,
		RoomID:      table.RoomID,
		Name:        table.Name,
		GameType:    table.GameType,
		Type:        table.Type,
		Status:      table.Status,
		MinPlayers:  table.MinPlayers,
		MaxPlayers:  table.MaxPlayers,
		PlayerCount: playerCount,
		Spectators:  len(table.Spectators),
		Private:     table.Private,
	}
}
