package handler

import (
	"context"
	"time"

	"github.com/game-engine/chat/internal/room"
	"github.com/game-engine/chat/internal/service"
)

// ChatHandler handles chat-related requests
type ChatHandler struct {
	service *service.ChatService
}

// NewChatHandler creates a new chat handler
func NewChatHandler(service *service.ChatService) *ChatHandler {
	return &ChatHandler{
		service: service,
	}
}

// CreateRoomRequest represents a create room request
type CreateRoomRequest struct {
	Name     string        `json:"name"`
	Type     room.RoomType `json:"type"`
	MaxUsers int           `json:"max_users"`
	OwnerID  string        `json:"owner_id"`
	Password string        `json:"password"`
}

// JoinRoomRequest represents a join room request
type JoinRoomRequest struct {
	RoomID   string `json:"room_id"`
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

// SendMessageRequest represents a send message request
type SendMessageRequest struct {
	RoomID   string `json:"room_id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Content  string `json:"content"`
	Type     string `json:"type"`
}

// RoomResponse represents a room response
type RoomResponse struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Type        room.RoomType `json:"type"`
	MaxUsers    int           `json:"max_users"`
	MemberCount int           `json:"member_count"`
	OwnerID     string        `json:"owner_id"`
}

// MessageResponse represents a message response
type MessageResponse struct {
	ID        string    `json:"id"`
	RoomID    string    `json:"room_id"`
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	Type      string    `json:"type"`
}

// CreateRoom creates a new room
func (h *ChatHandler) CreateRoom(ctx context.Context, req *CreateRoomRequest) (*RoomResponse, error) {
	room, err := h.service.CreateRoom(ctx, req.Name, req.Type, req.MaxUsers, req.OwnerID, req.Password)
	if err != nil {
		return nil, err
	}

	return h.roomToResponse(room), nil
}

// GetRoom gets a room
func (h *ChatHandler) GetRoom(ctx context.Context, roomID string) (*RoomResponse, error) {
	room, err := h.service.GetRoom(ctx, roomID)
	if err != nil {
		return nil, err
	}

	return h.roomToResponse(room), nil
}

// JoinRoom joins a room
func (h *ChatHandler) JoinRoom(ctx context.Context, req *JoinRoomRequest) error {
	return h.service.JoinRoom(ctx, req.RoomID, req.UserID, req.Password)
}

// LeaveRoom leaves a room
func (h *ChatHandler) LeaveRoom(ctx context.Context, roomID, userID string) error {
	return h.service.LeaveRoom(ctx, roomID, userID)
}

// SendMessage sends a message
func (h *ChatHandler) SendMessage(ctx context.Context, req *SendMessageRequest) (*MessageResponse, error) {
	msgType := req.Type
	if msgType == "" {
		msgType = "text"
	}

	message, err := h.service.SendMessage(ctx, req.RoomID, req.UserID, req.Username, req.Content, msgType)
	if err != nil {
		return nil, err
	}

	return h.messageToResponse(message), nil
}

// GetMessages gets messages
func (h *ChatHandler) GetMessages(ctx context.Context, roomID string, limit int) ([]MessageResponse, error) {
	messages, err := h.service.GetMessages(ctx, roomID, limit)
	if err != nil {
		return nil, err
	}

	result := make([]MessageResponse, 0, len(messages))
	for _, msg := range messages {
		result = append(result, *h.messageToResponse(&msg))
	}

	return result, nil
}

// ListRooms lists rooms
func (h *ChatHandler) ListRooms(ctx context.Context, roomType room.RoomType) ([]*RoomResponse, error) {
	rooms, err := h.service.ListRooms(ctx, roomType)
	if err != nil {
		return nil, err
	}

	result := make([]*RoomResponse, 0, len(rooms))
	for _, r := range rooms {
		result = append(result, h.roomToResponse(r))
	}

	return result, nil
}

// CloseRoom closes a room
func (h *ChatHandler) CloseRoom(ctx context.Context, roomID string) error {
	return h.service.CloseRoom(ctx, roomID)
}

// SendEmote sends an emote
func (h *ChatHandler) SendEmote(ctx context.Context, roomID, userID, username, emote string) (*MessageResponse, error) {
	message, err := h.service.SendEmote(ctx, roomID, userID, username, emote)
	if err != nil {
		return nil, err
	}

	return h.messageToResponse(message), nil
}

// UpdateRoom updates room settings
func (h *ChatHandler) UpdateRoom(ctx context.Context, roomID, name, password string, maxUsers int) error {
	return h.service.UpdateRoom(ctx, roomID, name, password, maxUsers)
}

func (h *ChatHandler) roomToResponse(room *room.ChatRoom) *RoomResponse {
	return &RoomResponse{
		ID:          room.ID,
		Name:        room.Name,
		Type:        room.Type,
		MaxUsers:    room.MaxUsers,
		MemberCount: len(room.Members),
		OwnerID:     room.OwnerID,
	}
}

func (h *ChatHandler) messageToResponse(message *room.Message) *MessageResponse {
	return &MessageResponse{
		ID:        message.ID,
		RoomID:    message.RoomID,
		UserID:    message.UserID,
		Username:  message.Username,
		Content:   message.Content,
		Timestamp: message.Timestamp,
		Type:      message.Type,
	}
}
