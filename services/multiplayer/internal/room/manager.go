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

// saveRoom saves a room to Redis
func (m *Manager) saveRoom(ctx context.Context, room *Room) error {
	data, err := json.Marshal(room)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("room:%s", room.ID)
	return m.redisClient.Set(ctx, key, data, 0).Err()
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
