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
	filter      *ProfanityFilter
	moderator   *Moderator
}

// NewManager creates a new chat room manager
func NewManager(config *Config, redisClient *redis.Client) (*Manager, error) {
	m := &Manager{
		rooms:       make(map[string]*ChatRoom),
		config:      config,
		redisClient: redisClient,
		filter:      NewProfanityFilter(config),
		moderator:   NewModerator(config, redisClient),
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

// SendMessage sends a message to a chat room
func (m *Manager) SendMessage(ctx context.Context, roomID, userID, username, content, msgType string) (*Message, error) {
	// Apply profanity filter
	if m.config.ProfanityFilter.Enabled {
		content = m.filter.Filter(content)
	}

	// Check rate limit
	if !m.moderator.CheckRateLimit(ctx, userID) {
		return nil, fmt.Errorf("rate limit exceeded")
	}

	message := &Message{
		ID:        uuid.New().String(),
		RoomID:    roomID,
		UserID:    userID,
		Username:  username,
		Content:   content,
		Timestamp: time.Now(),
		Type:      msgType,
	}

	// Save message to Redis
	key := fmt.Sprintf("chat:messages:%s", roomID)
	data, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	// Add to list and trim to max messages
	m.redisClient.LPush(ctx, key, data)
	m.redisClient.LTrim(ctx, key, 0, int64(m.config.Chat.MaxMessagesPerRoom))

	// Set expiry for message cleanup
	m.redisClient.Expire(ctx, key, time.Duration(m.config.Chat.MessageRetentionDays)*24*time.Hour)

	return message, nil
}

// GetMessages retrieves messages from a chat room
func (m *Manager) GetMessages(ctx context.Context, roomID string, limit int) ([]Message, error) {
	if limit <= 0 {
		limit = 50
	}
	if limit > 100 {
		limit = 100
	}

	key := fmt.Sprintf("chat:messages:%s", roomID)
	messages, err := m.redisClient.LRange(ctx, key, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	result := make([]Message, 0, len(messages))
	for _, msg := range messages {
		var message Message
		if err := json.Unmarshal([]byte(msg), &message); err != nil {
			continue
		}
		result = append(result, message)
	}

	return result, nil
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

// GetRoomMembers gets all members of a room
func (m *Manager) GetRoomMembers(ctx context.Context, roomID string) (map[string]bool, error) {
	room, err := m.GetRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	return room.Members, nil
}

// IsUserInRoom checks if a user is in a room
func (m *Manager) IsUserInRoom(ctx context.Context, roomID, userID string) (bool, error) {
	room, err := m.GetRoom(ctx, roomID)
	if err != nil {
		return false, err
	}

	return room.Members[userID], nil
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

// GetRoomStats gets room statistics
func (m *Manager) GetRoomStats(ctx context.Context, roomID string) (*RoomStats, error) {
	room, err := m.GetRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	messages, err := m.GetMessages(ctx, roomID, 1000)
	if err != nil {
		return nil, err
	}

	uniqueUsers := make(map[string]bool)
	var lastActivity time.Time
	for _, msg := range messages {
		if !msg.Deleted {
			uniqueUsers[msg.UserID] = true
			if msg.Timestamp.After(lastActivity) {
				lastActivity = msg.Timestamp
			}
		}
	}

	return &RoomStats{
		TotalMessages: room.MessageCount,
		UniqueUsers:   len(uniqueUsers),
		ActiveUsers:   len(room.Members),
		PeakUsers:     room.ActiveUsers,
		CreatedAt:     room.CreatedAt,
		LastActivity:  lastActivity,
	}, nil
}

// SetTypingIndicator sets user typing status
func (m *Manager) SetTypingIndicator(ctx context.Context, roomID, userID string, isTyping bool) error {
	key := fmt.Sprintf("chat:typing:%s", roomID)
	if isTyping {
		return m.redisClient.SAdd(ctx, key, userID).Err()
	}
	return m.redisClient.SRem(ctx, key, userID).Err()
}

// GetTypingUsers gets users currently typing in a room
func (m *Manager) GetTypingUsers(ctx context.Context, roomID string) ([]string, error) {
	key := fmt.Sprintf("chat:typing:%s", roomID)
	return m.redisClient.SMembers(ctx, key).Result()
}

// EditMessage edits a message (within 5 minute limit)
func (m *Manager) EditMessage(ctx context.Context, messageID, newContent string) (*Message, error) {
	// Apply profanity filter
	if m.config.ProfanityFilter.Enabled {
		newContent = m.filter.Filter(newContent)
	}

	// Find and update message in Redis
	// This requires iterating through all rooms - simplified implementation
	rooms, err := m.ListRooms(ctx, "")
	if err != nil {
		return nil, err
	}

	for _, room := range rooms {
		key := fmt.Sprintf("chat:messages:%s", room.ID)
		messages, err := m.redisClient.LRange(ctx, key, 0, -1).Result()
		if err != nil {
			continue
		}

		for _, msgData := range messages {
			var msg Message
			if err := json.Unmarshal([]byte(msgData), &msg); err != nil {
				continue
			}
			if msg.ID == messageID {
				// Check time limit (5 minutes)
				if time.Since(msg.Timestamp) > 5*time.Minute {
					return nil, fmt.Errorf("cannot edit message after 5 minutes")
				}
				msg.Content = newContent
				msg.Edited = true
				updated, _ := json.Marshal(msg)
				// Update in Redis list - this is simplified
				m.redisClient.LSet(ctx, key, 0, string(updated))
				return &msg, nil
			}
		}
	}

	return nil, fmt.Errorf("message not found")
}

// DeleteMessage deletes a message
func (m *Manager) DeleteMessage(ctx context.Context, messageID string) error {
	// Similar to EditMessage - find and mark as deleted
	rooms, err := m.ListRooms(ctx, "")
	if err != nil {
		return err
	}

	for _, room := range rooms {
		key := fmt.Sprintf("chat:messages:%s", room.ID)
		messages, err := m.redisClient.LRange(ctx, key, 0, -1).Result()
		if err != nil {
			continue
		}

		for _, msgData := range messages {
			var msg Message
			if err := json.Unmarshal([]byte(msgData), &msg); err != nil {
				continue
			}
			if msg.ID == messageID {
				msg.Deleted = true
				msg.Content = "[deleted]"
				updated, _ := json.Marshal(msg)
				m.redisClient.LSet(ctx, key, 0, string(updated))
				return nil
			}
		}
	}

	return fmt.Errorf("message not found")
}
