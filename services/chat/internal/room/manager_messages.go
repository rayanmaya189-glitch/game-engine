package room

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

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
