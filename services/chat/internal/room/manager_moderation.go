package room

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

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
