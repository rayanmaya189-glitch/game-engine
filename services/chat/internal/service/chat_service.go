package service

import (
	"context"
	"time"

	"github.com/game_engine/chat/internal/room"
)

// ChatService provides chat business logic
type ChatService struct {
	manager *room.Manager
}

// NewChatService creates a new chat service
func NewChatService(manager *room.Manager) (*ChatService, error) {
	return &ChatService{
		manager: manager,
	}, nil
}

// CreateRoom creates a new chat room
func (s *ChatService) CreateRoom(ctx context.Context, name string, roomType room.RoomType, maxUsers int, ownerID string, password string) (*room.ChatRoom, error) {
	return s.manager.CreateRoom(ctx, name, roomType, maxUsers, ownerID, password)
}

// GetRoom gets a room by ID
func (s *ChatService) GetRoom(ctx context.Context, roomID string) (*room.ChatRoom, error) {
	return s.manager.GetRoom(ctx, roomID)
}

// JoinRoom joins a chat room
func (s *ChatService) JoinRoom(ctx context.Context, roomID, userID string, password string) error {
	return s.manager.JoinRoom(ctx, roomID, userID, password)
}

// LeaveRoom leaves a chat room
func (s *ChatService) LeaveRoom(ctx context.Context, roomID, userID string) error {
	return s.manager.LeaveRoom(ctx, roomID, userID)
}

// SendMessage sends a message to a room
func (s *ChatService) SendMessage(ctx context.Context, roomID, userID, username, content, msgType string) (*room.Message, error) {
	return s.manager.SendMessage(ctx, roomID, userID, username, content, msgType)
}

// GetMessages gets messages from a room
func (s *ChatService) GetMessages(ctx context.Context, roomID string, limit int) ([]room.Message, error) {
	return s.manager.GetMessages(ctx, roomID, limit)
}

// ListRooms lists chat rooms
func (s *ChatService) ListRooms(ctx context.Context, roomType room.RoomType) ([]*room.ChatRoom, error) {
	return s.manager.ListRooms(ctx, roomType)
}

// CloseRoom closes a chat room
func (s *ChatService) CloseRoom(ctx context.Context, roomID string) error {
	return s.manager.CloseRoom(ctx, roomID)
}

// UpdateRoom updates room settings
func (s *ChatService) UpdateRoom(ctx context.Context, roomID string, name string, password string, maxUsers int) error {
	return s.manager.UpdateRoom(ctx, roomID, name, password, maxUsers)
}

// GetRoomMembers gets room members
func (s *ChatService) GetRoomMembers(ctx context.Context, roomID string) (map[string]bool, error) {
	return s.manager.GetRoomMembers(ctx, roomID)
}

// IsUserInRoom checks if user is in room
func (s *ChatService) IsUserInRoom(ctx context.Context, roomID, userID string) (bool, error) {
	return s.manager.IsUserInRoom(ctx, roomID, userID)
}

// SendEmote sends an emote message
func (s *ChatService) SendEmote(ctx context.Context, roomID, userID, username, emote string) (*room.Message, error) {
	return s.manager.SendMessage(ctx, roomID, userID, username, emote, "emote")
}

// GetMessageHistory gets message history with pagination
func (s *ChatService) GetMessageHistory(ctx context.Context, roomID string, before time.Time, limit int) ([]room.Message, error) {
	messages, err := s.manager.GetMessages(ctx, roomID, limit*2)
	if err != nil {
		return nil, err
	}

	// Filter messages before the given time
	result := make([]room.Message, 0)
	for _, msg := range messages {
		if msg.Timestamp.Before(before) {
			result = append(result, msg)
			if len(result) >= limit {
				break
			}
		}
	}

	return result, nil
}
