package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/game_engine/chat/internal/room"
)

// DMService handles direct messaging (1-to-1 chat)
type DMService struct {
	manager *room.Manager
}

// NewDMService creates a new DM service
func NewDMService(manager *room.Manager) *DMService {
	return &DMService{manager: manager}
}

// GetOrCreateDMRoom gets or creates a private DM room between two users
func (s *DMService) GetOrCreateDMRoom(ctx context.Context, userID1, userID2 string) (*room.ChatRoom, error) {
	if userID1 == userID2 {
		return nil, fmt.Errorf("cannot create DM with yourself")
	}

	// Check if DM room already exists
	rooms, err := s.manager.ListRooms(ctx, room.RoomTypePrivate)
	if err == nil {
		for _, r := range rooms {
			if s.isDMRoom(r, userID1, userID2) {
				return r, nil
			}
		}
	}

	// Create new DM room
	roomName := fmt.Sprintf("dm:%s:%s", min(userID1, userID2), max(userID1, userID2))
	dmRoom, err := s.manager.CreateRoom(ctx, roomName, room.RoomTypePrivate, 2, userID1, "")
	if err != nil {
		return nil, fmt.Errorf("failed to create DM room: %w", err)
	}

	// Auto-join both users
	_ = s.manager.JoinRoom(ctx, dmRoom.RoomID, userID1, "")
	_ = s.manager.JoinRoom(ctx, dmRoom.RoomID, userID2, "")

	return dmRoom, nil
}

// SendDM sends a direct message between two users
func (s *DMService) SendDM(ctx context.Context, fromUserID, toUserID, content string) (*room.Message, error) {
	dmRoom, err := s.GetOrCreateDMRoom(ctx, fromUserID, toUserID)
	if err != nil {
		return nil, err
	}

	msg, err := s.manager.SendMessage(ctx, dmRoom.RoomID, fromUserID, fromUserID, content, "text")
	if err != nil {
		return nil, fmt.Errorf("failed to send DM: %w", err)
	}

	return msg, nil
}

// GetDMHistory gets message history between two users
func (s *DMService) GetDMHistory(ctx context.Context, userID1, userID2 string, limit int) ([]room.Message, error) {
	dmRoom, err := s.GetOrCreateDMRoom(ctx, userID1, userID2)
	if err != nil {
		return nil, err
	}

	return s.manager.GetMessages(ctx, dmRoom.RoomID, limit)
}

// ListDMConversations lists all DM conversations for a user
func (s *DMService) ListDMConversations(ctx context.Context, userID string) ([]DMConversation, error) {
	rooms, err := s.manager.ListRooms(ctx, room.RoomTypePrivate)
	if err != nil {
		return nil, err
	}

	var conversations []DMConversation
	for _, r := range rooms {
		if s.isUserInDMRoom(r, userID) {
			otherUserID := s.getOtherUser(r, userID)
			messages, _ := s.manager.GetMessages(ctx, r.RoomID, 1)

			var lastMessage *room.Message
			if len(messages) > 0 {
				lastMessage = &messages[0]
			}

			conversations = append(conversations, DMConversation{
				RoomID:       r.RoomID,
				OtherUserID:  otherUserID,
				LastMessage:  lastMessage,
				UpdatedAt:    r.UpdatedAt,
			})
		}
	}

	return conversations, nil
}

// DMConversation represents a DM conversation summary
type DMConversation struct {
	RoomID      string        `json:"room_id"`
	OtherUserID string        `json:"other_user_id"`
	LastMessage *room.Message `json:"last_message,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

// --- Helper Functions ---

func (s *DMService) isDMRoom(r *room.ChatRoom, userID1, userID2 string) bool {
	if r.Type != room.RoomTypePrivate {
		return false
	}
	return strings.Contains(r.Name, userID1) && strings.Contains(r.Name, userID2)
}

func (s *DMService) isUserInDMRoom(r *room.ChatRoom, userID string) bool {
	return strings.Contains(r.Name, userID)
}

func (s *DMService) getOtherUser(r *room.ChatRoom, userID string) string {
	parts := strings.Split(r.Name, ":")
	if len(parts) != 3 {
		return ""
	}
	if parts[1] == userID {
		return parts[2]
	}
	return parts[1]
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}

func max(a, b string) string {
	if a > b {
		return a
	}
	return b
}
