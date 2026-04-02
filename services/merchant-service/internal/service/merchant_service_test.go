package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/game_engine/merchant-service/internal/model"
)

type mockMerchantRepo struct {
	players    map[string][]model.MerchantPlayer
	agents     map[string][]model.Agent
	webhooks   map[string][]model.Webhook
	configs    map[string]map[string]string
	agentIndex map[string]int
}

func newMockMerchantRepo() *mockMerchantRepo {
	return &mockMerchantRepo{
		players:    make(map[string][]model.MerchantPlayer),
		agents:     make(map[string][]model.Agent),
		webhooks:   make(map[string][]model.Webhook),
		configs:    make(map[string]map[string]string),
		agentIndex: make(map[string]int),
	}
}

func (m *mockMerchantRepo) ListPlayers(ctx context.Context, merchantID string, page, limit int, search string) ([]model.MerchantPlayer, int, error) {
	players := m.players[merchantID]
	return players, len(players), nil
}

func (m *mockMerchantRepo) GetPlayer(ctx context.Context, merchantID, playerID string) (*model.MerchantPlayer, error) {
	for _, p := range m.players[merchantID] {
		if p.PlayerID == playerID {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("player not found")
}

func (m *mockMerchantRepo) CreateAgent(ctx context.Context, merchantID, username, email string, sendInvitation bool) (string, error) {
	id := fmt.Sprintf("agent-%d", m.agentIndex[merchantID])
	m.agentIndex[merchantID]++
	m.agents[merchantID] = append(m.agents[merchantID], model.Agent{
		AgentID: id, Username: username, Email: email, Status: "active",
	})
	return id, nil
}

func (m *mockMerchantRepo) UpdateAgent(ctx context.Context, merchantID, agentID, username, email string) error {
	for i, a := range m.agents[merchantID] {
		if a.AgentID == agentID {
			m.agents[merchantID][i].Username = username
			m.agents[merchantID][i].Email = email
			return nil
		}
	}
	return fmt.Errorf("agent not found")
}

func (m *mockMerchantRepo) UpdateAgentStatus(ctx context.Context, merchantID, agentID, status string) error {
	for i, a := range m.agents[merchantID] {
		if a.AgentID == agentID {
			m.agents[merchantID][i].Status = status
			return nil
		}
	}
	return fmt.Errorf("agent not found")
}

func TestListPlayers(t *testing.T) {
	repo := newMockMerchantRepo()
	repo.players["m-1"] = []model.MerchantPlayer{
		{PlayerID: "p-1", Username: "player1", Status: "active"},
		{PlayerID: "p-2", Username: "player2", Status: "active"},
	}

	tests := []struct {
		name       string
		merchantID string
		page       int
		limit      int
		wantLen    int
	}{
		{"first page", "m-1", 1, 10, 2},
		{"empty merchant", "m-99", 1, 10, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			players, total, err := repo.ListPlayers(context.Background(), tt.merchantID, tt.page, tt.limit, "")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(players) != tt.wantLen {
				t.Errorf("got %d players, want %d", len(players), tt.wantLen)
			}
			if total != tt.wantLen {
				t.Errorf("total = %d, want %d", total, tt.wantLen)
			}
		})
	}
}

func TestGetPlayer(t *testing.T) {
	repo := newMockMerchantRepo()
	repo.players["m-1"] = []model.MerchantPlayer{
		{PlayerID: "p-1", Username: "player1", Status: "active"},
	}

	tests := []struct {
		name       string
		merchantID string
		playerID   string
		wantErr    bool
	}{
		{"existing player", "m-1", "p-1", false},
		{"non-existent player", "m-1", "p-99", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player, err := repo.GetPlayer(context.Background(), tt.merchantID, tt.playerID)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetPlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && player.PlayerID != tt.playerID {
				t.Errorf("playerID = %v, want %v", player.PlayerID, tt.playerID)
			}
		})
	}
}

func TestCreateAgent(t *testing.T) {
	repo := newMockMerchantRepo()

	tests := []struct {
		name     string
		username string
		email    string
	}{
		{"first agent", "agent1", "a1@test.com"},
		{"second agent", "agent2", "a2@test.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := repo.CreateAgent(context.Background(), "m-1", tt.username, tt.email, false)
			if err != nil {
				t.Fatalf("CreateAgent() error = %v", err)
			}
			if id == "" {
				t.Error("expected non-empty agent ID")
			}
		})
	}

	if len(repo.agents["m-1"]) != 2 {
		t.Errorf("expected 2 agents, got %d", len(repo.agents["m-1"]))
	}
}

func TestUpdateAgent(t *testing.T) {
	repo := newMockMerchantRepo()
	repo.agents["m-1"] = []model.Agent{
		{AgentID: "a-1", Username: "old", Email: "old@test.com", Status: "active"},
	}

	tests := []struct {
		name     string
		agentID  string
		username string
		email    string
		wantErr  bool
	}{
		{"update existing", "a-1", "new_name", "new@test.com", false},
		{"update non-existent", "a-99", "x", "x@test.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdateAgent(context.Background(), "m-1", tt.agentID, tt.username, tt.email)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UpdateAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateAgentStatus(t *testing.T) {
	repo := newMockMerchantRepo()
	repo.agents["m-1"] = []model.Agent{
		{AgentID: "a-1", Status: "active"},
	}

	tests := []struct {
		name    string
		agentID string
		status  string
		wantErr bool
	}{
		{"suspend agent", "a-1", "suspended", false},
		{"activate agent", "a-1", "active", false},
		{"non-existent agent", "a-99", "active", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdateAgentStatus(context.Background(), "m-1", tt.agentID, tt.status)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UpdateAgentStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
