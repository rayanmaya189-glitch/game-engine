package room

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/game_engine/chat/internal/moderation"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// RoomType represents the type of chat room
type RoomType string

const (
	RoomTypePublic  RoomType = "public"
	RoomTypePrivate RoomType = "private"
	RoomTypeTable   RoomType = "table"
	RoomTypeGame    RoomType = "game"
)

// ChatRoom represents a chat room
type ChatRoom struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Type         RoomType        `json:"type"`
	Members      map[string]bool `json:"members"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	MaxUsers     int             `json:"max_users"`
	OwnerID      string          `json:"owner_id"`
	Password     string          `json:"password,omitempty"`
	MessageCount int             `json:"message_count"`
	ActiveUsers  int             `json:"active_users"`
}

// Message represents a chat message
type Message struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"` // text, emote, system
	Edited    bool      `json:"edited"`
	Deleted   bool      `json:"deleted"`
}

// RoomStats represents room statistics
type RoomStats struct {
	TotalMessages int       `json:"total_messages"`
	UniqueUsers   int       `json:"unique_users"`
	ActiveUsers   int       `json:"active_users"`
	PeakUsers     int       `json:"peak_users"`
	CreatedAt     time.Time `json:"created_at"`
	LastActivity  time.Time `json:"last_activity"`
}

// Manager handles chat room lifecycle
type Manager struct {
	mu          sync.RWMutex
	rooms       map[string]*ChatRoom
	config      *Config
	redisClient *redis.Client
	filter      *moderation.ProfanityFilter
	moderator   *moderation.Moderator
}

// NewManager creates a new chat room manager
func NewManager(config *Config, redisClient *redis.Client) (*Manager, error) {
	// Convert room.Config to moderation.FilterConfig
	filterConfig := &moderation.FilterConfig{
		ProfanityFilter: moderation.ProfanityConfig{
			Enabled:         config.ProfanityFilter.Enabled,
			ReplacementChar: config.ProfanityFilter.ReplacementChar,
			FilterLevel:     config.ProfanityFilter.FilterLevel,
		},
		Moderation: moderation.ModerationConfig{
			AutoMuteThreshold:   config.Moderation.AutoMuteThreshold,
			MuteDurationMinutes: config.Moderation.MuteDurationMinutes,
			BanDurationHours:    config.Moderation.BanDurationHours,
			RequiresModerator:   config.Moderation.RequiresModerator,
		},
	}

	m := &Manager{
		rooms:       make(map[string]*ChatRoom),
		config:      config,
		redisClient: redisClient,
		filter:      moderation.NewProfanityFilter(filterConfig),
		moderator:   moderation.NewModerator(filterConfig, redisClient),
	}

	// Load existing rooms from Redis
	if err := m.loadRooms(context.Background()); err != nil {
		return nil, err
	}

	return m, nil
}

// CreateRoom creates a new chat room
func (m *Manager) CreateRoom(ctx context.Context, name string, roomType RoomType, maxUsers int, ownerID string, password string) (*ChatRoom, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	room := &ChatRoom{
		ID:        uuid.New().String(),
		Name:      name,
		Type:      roomType,
		Members:   make(map[string]bool),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		MaxUsers:  maxUsers,
		OwnerID:   ownerID,
		Password:  password,
	}

	m.rooms[room.ID] = room

	// Save to Redis
	if err := m.saveRoom(ctx, room); err != nil {
		return nil, err
	}

	return room, nil
}

// GetRoom retrieves a chat room
func (m *Manager) GetRoom(ctx context.Context, roomID string) (*ChatRoom, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	room, ok := m.rooms[roomID]
	if !ok {
		// Try to load from Redis
		data, err := m.redisClient.Get(ctx, fmt.Sprintf("chat:room:%s", roomID)).Bytes()
		if err != nil {
			return nil, fmt.Errorf("room not found: %s", roomID)
		}

		var room ChatRoom
		if err := json.Unmarshal(data, &room); err != nil {
			return nil, err
		}
		return &room, nil
	}

	return room, nil
}

// JoinRoom adds a user to a chat room
func (m *Manager) JoinRoom(ctx context.Context, roomID, userID string, password string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	room, ok := m.rooms[roomID]
	if !ok {
		return fmt.Errorf("room not found: %s", roomID)
	}

	if room.Password != "" && room.Password != password {
		return fmt.Errorf("invalid password")
	}

	if room.MaxUsers > 0 && len(room.Members) >= room.MaxUsers {
		return fmt.Errorf("room is full")
	}

	room.Members[userID] = true
	room.UpdatedAt = time.Now()

	// Save to Redis
	return m.saveRoom(ctx, room)
}

// LeaveRoom removes a user from a chat room
func (m *Manager) LeaveRoom(ctx context.Context, roomID, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	room, ok := m.rooms[roomID]
	if !ok {
		return fmt.Errorf("room not found: %s", roomID)
	}

	delete(room.Members, userID)
	room.UpdatedAt = time.Now()

	// Save to Redis
	return m.saveRoom(ctx, room)
}

// ListRooms lists available chat rooms
func (m *Manager) ListRooms(ctx context.Context, roomType RoomType) ([]*ChatRoom, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result []*ChatRoom
	for _, room := range m.rooms {
		if roomType != "" && room.Type != roomType {
			continue
		}
		result = append(result, room)
	}

	return result, nil
}

// CloseRoom closes a chat room
func (m *Manager) CloseRoom(ctx context.Context, roomID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	room, ok := m.rooms[roomID]
	if !ok {
		return fmt.Errorf("room not found: %s", roomID)
	}

	room.UpdatedAt = time.Now()

	// Remove from Redis
	m.redisClient.Del(ctx, fmt.Sprintf("chat:room:%s", roomID))
	m.redisClient.Del(ctx, fmt.Sprintf("chat:messages:%s", roomID))

	delete(m.rooms, roomID)

	return nil
}

// UpdateRoom updates room settings
func (m *Manager) UpdateRoom(ctx context.Context, roomID string, name string, password string, maxUsers int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	room, ok := m.rooms[roomID]
	if !ok {
		return fmt.Errorf("room not found: %s", roomID)
	}

	if name != "" {
		room.Name = name
	}
	if password != "" {
		room.Password = password
	}
	if maxUsers > 0 {
		room.MaxUsers = maxUsers
	}
	room.UpdatedAt = time.Now()

	return m.saveRoom(ctx, room)
}

// saveRoom saves a room to Redis
func (m *Manager) saveRoom(ctx context.Context, room *ChatRoom) error {
	data, err := json.Marshal(room)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("chat:room:%s", room.ID)
	return m.redisClient.Set(ctx, key, data, 0).Err()
}

// loadRooms loads rooms from Redis
func (m *Manager) loadRooms(ctx context.Context) error {
	keys, err := m.redisClient.Keys(ctx, "chat:room:*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		data, err := m.redisClient.Get(ctx, key).Bytes()
		if err != nil {
			continue
		}

		var room ChatRoom
		if err := json.Unmarshal(data, &room); err != nil {
			continue
		}

		m.rooms[room.ID] = &room
	}

	return nil
}
