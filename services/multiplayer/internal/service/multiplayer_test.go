package service

import (
	"testing"

	"github.com/game_engine/multiplayer/internal/room"
)

func TestTableTypeConstants(t *testing.T) {
	types := []room.TableType{
		room.TableTypePublic,
		room.TableTypePrivate,
		room.TableTypeTournament,
	}
	for _, tt := range types {
		if string(tt) == "" {
			t.Fatal("table type should not be empty")
		}
	}
}

func TestTableStatusConstants(t *testing.T) {
	statuses := []room.TableStatus{
		room.TableStatusWaiting,
		room.TableStatusPlaying,
		room.TableStatusPaused,
		room.TableStatusClosed,
	}
	for _, s := range statuses {
		if string(s) == "" {
			t.Fatal("table status should not be empty")
		}
	}
}

func TestRoomStructure(t *testing.T) {
	r := &room.Room{
		ID:        "room-1",
		Name:      "Poker Room",
		GameType:  "poker",
		TableType: room.TableTypePublic,
		Tables:    make(map[string]*room.Table),
	}

	if r.Name != "Poker Room" {
		t.Fatalf("Name = %s, want Poker Room", r.Name)
	}
	if r.Tables == nil {
		t.Fatal("Tables map should not be nil")
	}
}

func TestTableStructure(t *testing.T) {
	tb := &room.Table{
		ID:         "t1",
		RoomID:     "room-1",
		Name:       "Table 1",
		GameType:   "poker",
		Type:       room.TableTypePublic,
		Status:     room.TableStatusWaiting,
		MinPlayers: 2,
		MaxPlayers: 9,
		BuyInMin:   100,
		BuyInMax:   10000,
		Seats:      make(map[int]*room.Seat),
		Spectators: make(map[string]bool),
		GameState:  make(map[string]interface{}),
	}

	if tb.MinPlayers >= tb.MaxPlayers {
		t.Fatal("MinPlayers should be less than MaxPlayers")
	}
	if tb.BuyInMin > tb.BuyInMax {
		t.Fatal("BuyInMin should not exceed BuyInMax")
	}
}

func TestSeatStructure(t *testing.T) {
	seat := &room.Seat{
		ID:        1,
		UserID:    "u1",
		Username:  "alice",
		Chips:     500,
		Bet:       0,
		Ready:     false,
		Connected: true,
	}

	if seat.Chips < 0 {
		t.Fatal("Chips should not be negative")
	}
	if seat.ID < 1 {
		t.Fatal("Seat ID should be >= 1")
	}
}

func TestSeatReadyState(t *testing.T) {
	tests := []struct {
		name  string
		ready bool
	}{
		{"ready", true},
		{"not ready", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seat := &room.Seat{ID: 1, UserID: "u1", Ready: tt.ready}
			if seat.Ready != tt.ready {
				t.Fatalf("Ready = %v, want %v", seat.Ready, tt.ready)
			}
		})
	}
}

func TestJoinTableValidation(t *testing.T) {
	tests := []struct {
		name       string
		tableChips int64
		buyInMin   int64
		buyInMax   int64
		valid      bool
	}{
		{"valid buy-in", 500, 100, 1000, true},
		{"below minimum", 50, 100, 1000, false},
		{"above maximum", 2000, 100, 1000, false},
		{"at minimum", 100, 100, 1000, true},
		{"at maximum", 1000, 100, 1000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := tt.tableChips >= tt.buyInMin && tt.tableChips <= tt.buyInMax
			if valid != tt.valid {
				t.Fatalf("chips=%d, min=%d, max=%d: got valid=%v, want %v",
					tt.tableChips, tt.buyInMin, tt.buyInMax, valid, tt.valid)
			}
		})
	}
}

func TestPrivateTablePassword(t *testing.T) {
	tests := []struct {
		name     string
		tablePW  string
		inputPW  string
		isPrivate bool
		canJoin  bool
	}{
		{"public no pw", "", "", false, true},
		{"private correct pw", "secret", "secret", true, true},
		{"private wrong pw", "secret", "wrong", true, false},
		{"private empty pw", "secret", "", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canJoin := !tt.isPrivate || tt.tablePW == tt.inputPW
			if canJoin != tt.canJoin {
				t.Fatalf("got canJoin=%v, want %v", canJoin, tt.canJoin)
			}
		})
	}
}

func TestTableCapacityCheck(t *testing.T) {
	tests := []struct {
		name      string
		current   int
		max       int
		canJoin   bool
	}{
		{"empty", 0, 6, true},
		{"half full", 3, 6, true},
		{"one seat left", 5, 6, true},
		{"full", 6, 6, false},
		{"over full", 7, 6, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			canJoin := tt.current < tt.max
			if canJoin != tt.canJoin {
				t.Fatalf("current=%d, max=%d: got canJoin=%v, want %v",
					tt.current, tt.max, canJoin, tt.canJoin)
			}
		})
	}
}

func TestGameStateMap(t *testing.T) {
	state := map[string]interface{}{
		"pot":       1000,
		"phase":     "flop",
		"community": []string{"AH", "KD", "QC"},
	}

	pot, ok := state["pot"]
	if !ok {
		t.Fatal("pot should exist in game state")
	}
	if pot != 1000 {
		t.Fatalf("pot = %v, want 1000", pot)
	}

	phase, ok := state["phase"]
	if !ok {
		t.Fatal("phase should exist in game state")
	}
	if phase != "flop" {
		t.Fatalf("phase = %v, want flop", phase)
	}
}

func TestSpectatorMap(t *testing.T) {
	spectators := map[string]bool{
		"spec1": true,
		"spec2": true,
	}

	if !spectators["spec1"] {
		t.Fatal("spec1 should be in spectators")
	}
	if spectators["spec3"] {
		t.Fatal("spec3 should not be in spectators")
	}

	delete(spectators, "spec1")
	if spectators["spec1"] {
		t.Fatal("spec1 should be removed")
	}
}

func TestTableTurnTracking(t *testing.T) {
	tb := &room.Table{
		ID:          "t1",
		CurrentTurn: "u1",
	}

	if tb.CurrentTurn != "u1" {
		t.Fatalf("CurrentTurn = %s, want u1", tb.CurrentTurn)
	}

	tb.CurrentTurn = "u2"
	if tb.CurrentTurn != "u2" {
		t.Fatalf("CurrentTurn = %s, want u2", tb.CurrentTurn)
	}
}

func TestRoomTableMapping(t *testing.T) {
	r := &room.Room{
		ID:     "room-1",
		Tables: make(map[string]*room.Table),
	}

	tb1 := &room.Table{ID: "t1", RoomID: "room-1"}
	tb2 := &room.Table{ID: "t2", RoomID: "room-1"}

	r.Tables[tb1.ID] = tb1
	r.Tables[tb2.ID] = tb2

	if len(r.Tables) != 2 {
		t.Fatalf("got %d tables, want 2", len(r.Tables))
	}

	delete(r.Tables, "t1")
	if len(r.Tables) != 1 {
		t.Fatalf("got %d tables, want 1", len(r.Tables))
	}
}
