package room

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// TableType represents the type of table
type TableType string

const (
	TableTypePublic     TableType = "public"
	TableTypePrivate    TableType = "private"
	TableTypeTournament TableType = "tournament"
)

// TableStatus represents the status of a table
type TableStatus string

const (
	TableStatusWaiting TableStatus = "waiting"
	TableStatusPlaying TableStatus = "playing"
	TableStatusPaused  TableStatus = "paused"
	TableStatusClosed  TableStatus = "closed"
)

// Room represents a game room
type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	GameType  string    `json:"game_type"`
	TableType TableType `json:"table_type"`
	Tables    map[string]*Table
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Table represents a game table
type Table struct {
	ID           string                 `json:"id"`
	RoomID       string                 `json:"room_id"`
	Name         string                 `json:"name"`
	GameType     string                 `json:"game_type"`
	Type         TableType              `json:"type"`
	Status       TableStatus            `json:"status"`
	MinPlayers   int                    `json:"min_players"`
	MaxPlayers   int                    `json:"max_players"`
	BuyInMin     int64                  `json:"buy_in_min"`
	BuyInMax     int64                  `json:"buy_in_max"`
	Seats        map[int]*Seat          `json:"seats"`
	Spectators   map[string]bool        `json:"spectators"`
	CurrentTurn  string                 `json:"current_turn"`
	GameState    map[string]interface{} `json:"game_state"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	Private      bool                   `json:"private"`
	Password     string                 `json:"password,omitempty"`
	TournamentID string                 `json:"tournament_id,omitempty"`
}

// Seat represents a seat at a table
type Seat struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Chips     int64  `json:"chips"`
	Bet       int64  `json:"bet"`
	Ready     bool   `json:"ready"`
	Connected bool   `json:"connected"`
	Avatar    string `json:"avatar"`
}

// Manager handles room and table lifecycle
type Manager struct {
	mu          sync.RWMutex
	rooms       map[string]*Room
	tables      map[string]*Table // tableID -> Table
	config      *Config
	redisClient *redis.Client
}

// NewManager creates a new room manager
func NewManager(config *Config, redisClient *redis.Client) (*Manager, error) {
	m := &Manager{
		rooms:       make(map[string]*Room),
		tables:      make(map[string]*Table),
		config:      config,
		redisClient: redisClient,
	}

	// Load existing rooms and tables from Redis
	if err := m.loadRooms(context.Background()); err != nil {
		return nil, err
	}

	return m, nil
}

// CreateRoom creates a new game room
func (m *Manager) CreateRoom(ctx context.Context, name, gameType string, tableType TableType) (*Room, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	room := &Room{
		ID:        uuid.New().String(),
		Name:      name,
		GameType:  gameType,
		TableType: tableType,
		Tables:    make(map[string]*Table),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	m.rooms[room.ID] = room

	// Save to Redis
	if err := m.saveRoom(ctx, room); err != nil {
		return nil, err
	}

	return room, nil
}

// GetRoom gets a room by ID
func (m *Manager) GetRoom(ctx context.Context, roomID string) (*Room, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	room, ok := m.rooms[roomID]
	if !ok {
		return nil, fmt.Errorf("room not found: %s", roomID)
	}

	return room, nil
}

// ListRooms lists all rooms
func (m *Manager) ListRooms(ctx context.Context) ([]*Room, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]*Room, 0, len(m.rooms))
	for _, room := range m.rooms {
		result = append(result, room)
	}

	return result, nil
}

// saveRoom saves a room to Redis
func (m *Manager) saveRoom(ctx context.Context, room *Room) error {
	data, err := json.Marshal(room)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("room:%s", room.ID)
	return m.redisClient.Set(ctx, key, data, 0).Err()
}

// loadRooms loads rooms from Redis
func (m *Manager) loadRooms(ctx context.Context) error {
	keys, err := m.redisClient.Keys(ctx, "room:*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		data, err := m.redisClient.Get(ctx, key).Bytes()
		if err != nil {
			continue
		}

		var room Room
		if err := json.Unmarshal(data, &room); err != nil {
			continue
		}

		m.rooms[room.ID] = &room
	}

	// Load tables
	keys, err = m.redisClient.Keys(ctx, "table:*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		data, err := m.redisClient.Get(ctx, key).Bytes()
		if err != nil {
			continue
		}

		var table Table
		if err := json.Unmarshal(data, &table); err != nil {
			continue
		}

		m.tables[table.ID] = &table
		if room, ok := m.rooms[table.RoomID]; ok {
			room.Tables[table.ID] = &table
		}
	}

	return nil
}
