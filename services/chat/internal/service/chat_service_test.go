package service

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type mockChatRoom struct {
	ID       string
	Name     string
	Type     string
	Members  map[string]bool
	OwnerID  string
	Password string
	MaxUsers int
}

type mockMessage struct {
	ID        string
	RoomID    string
	UserID    string
	Username  string
	Content   string
	Timestamp time.Time
	Type      string
}

type mockChatStore struct {
	rooms    map[string]*mockChatRoom
	messages map[string][]mockMessage
}

func newMockChatStore() *mockChatStore {
	return &mockChatStore{
		rooms:    make(map[string]*mockChatRoom),
		messages: make(map[string][]mockMessage),
	}
}

func TestCreateRoom(t *testing.T) {
	tests := []struct {
		name     string
		roomName string
		roomType string
		maxUsers int
		ownerID  string
	}{
		{"public room", "General Chat", "public", 100, "owner-1"},
		{"private room", "VIP Lounge", "private", 20, "owner-2"},
		{"game room", "Poker Table 1", "game", 9, "owner-3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := newMockChatStore()
			room := &mockChatRoom{
				ID:       fmt.Sprintf("room-%s", tt.roomName),
				Name:     tt.roomName,
				Type:     tt.roomType,
				Members:  make(map[string]bool),
				OwnerID:  tt.ownerID,
				MaxUsers: tt.maxUsers,
			}
			store.rooms[room.ID] = room

			stored := store.rooms[room.ID]
			if stored.Name != tt.roomName {
				t.Errorf("name = %v, want %v", stored.Name, tt.roomName)
			}
			if stored.MaxUsers != tt.maxUsers {
				t.Errorf("maxUsers = %d, want %d", stored.MaxUsers, tt.maxUsers)
			}
		})
	}
}

func TestJoinRoom(t *testing.T) {
	store := newMockChatStore()
	store.rooms["room-1"] = &mockChatRoom{
		ID:       "room-1",
		Name:     "Test Room",
		Members:  make(map[string]bool),
		MaxUsers: 10,
	}
	store.rooms["room-pwd"] = &mockChatRoom{
		ID:       "room-pwd",
		Name:     "Protected",
		Members:  make(map[string]bool),
		MaxUsers: 10,
		Password: "secret",
	}
	store.rooms["room-full"] = &mockChatRoom{
		ID:       "room-full",
		Name:     "Full Room",
		Members:  map[string]bool{"u1": true, "u2": true},
		MaxUsers: 2,
	}

	tests := []struct {
		name     string
		roomID   string
		userID   string
		password string
		wantErr  bool
		errMsg   string
	}{
		{"join open room", "room-1", "user-1", "", false, ""},
		{"join protected with correct password", "room-pwd", "user-1", "secret", false, ""},
		{"join protected with wrong password", "room-pwd", "user-2", "wrong", true, "invalid password"},
		{"join full room", "room-full", "user-3", "", true, "room is full"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			room := store.rooms[tt.roomID]

			var err error
			if room.Password != "" && room.Password != tt.password {
				err = fmt.Errorf("invalid password")
			} else if room.MaxUsers > 0 && len(room.Members) >= room.MaxUsers {
				err = fmt.Errorf("room is full")
			} else {
				room.Members[tt.userID] = true
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("JoinRoom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLeaveRoom(t *testing.T) {
	store := newMockChatStore()
	store.rooms["room-1"] = &mockChatRoom{
		ID:      "room-1",
		Members: map[string]bool{"user-1": true, "user-2": true},
	}

	room := store.rooms["room-1"]
	delete(room.Members, "user-1")

	if _, exists := room.Members["user-1"]; exists {
		t.Error("user-1 should have been removed")
	}
	if len(room.Members) != 1 {
		t.Errorf("expected 1 member, got %d", len(room.Members))
	}
}

func TestSendMessage(t *testing.T) {
	store := newMockChatStore()
	store.rooms["room-1"] = &mockChatRoom{ID: "room-1", Members: map[string]bool{"user-1": true}}

	tests := []struct {
		name    string
		content string
		msgType string
	}{
		{"text message", "Hello world!", "text"},
		{"emote message", "waves hello", "emote"},
		{"system message", "User joined", "system"},
		{"empty content", "", "text"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := mockMessage{
				ID:        fmt.Sprintf("msg-%d", len(store.messages["room-1"])),
				RoomID:    "room-1",
				UserID:    "user-1",
				Username:  "player1",
				Content:   tt.content,
				Timestamp: time.Now(),
				Type:      tt.msgType,
			}
			store.messages["room-1"] = append(store.messages["room-1"], msg)

			msgs := store.messages["room-1"]
			last := msgs[len(msgs)-1]
			if last.Content != tt.content {
				t.Errorf("content = %v, want %v", last.Content, tt.content)
			}
			if last.Type != tt.msgType {
				t.Errorf("type = %v, want %v", last.Type, tt.msgType)
			}
		})
	}
}

func TestGetMessages(t *testing.T) {
	store := newMockChatStore()
	now := time.Now()
	for i := 0; i < 20; i++ {
		store.messages["room-1"] = append(store.messages["room-1"], mockMessage{
			ID:        fmt.Sprintf("msg-%d", i),
			RoomID:    "room-1",
			UserID:    "user-1",
			Content:   fmt.Sprintf("message %d", i),
			Timestamp: now.Add(time.Duration(i) * time.Second),
		})
	}

	tests := []struct {
		name    string
		roomID  string
		limit   int
		wantLen int
	}{
		{"get 5 messages", "room-1", 5, 5},
		{"get all messages", "room-1", 100, 20},
		{"empty room", "room-99", 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msgs := store.messages[tt.roomID]
			if len(msgs) > tt.limit {
				msgs = msgs[:tt.limit]
			}
			if len(msgs) != tt.wantLen {
				t.Errorf("got %d messages, want %d", len(msgs), tt.wantLen)
			}
		})
	}
}
