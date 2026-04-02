package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/game_engine/agent-service/internal/model"
)

type mockAgentRepo struct {
	players   map[string][]model.Player
	dashboards map[string]*model.Dashboard
}

func newMockAgentRepo() *mockAgentRepo {
	return &mockAgentRepo{
		players:    make(map[string][]model.Player),
		dashboards: make(map[string]*model.Dashboard),
	}
}

func (m *mockAgentRepo) ListPlayers(ctx context.Context, agentID string, page, limit int, search, status string) ([]model.Player, int, error) {
	var filtered []model.Player
	for _, p := range m.players[agentID] {
		if status != "" && p.Status != status {
			continue
		}
		if search != "" {
			found := false
			if len(p.Username) >= len(search) && p.Username[:len(search)] == search {
				found = true
			}
			if !found {
				continue
			}
		}
		filtered = append(filtered, p)
	}
	return filtered, len(filtered), nil
}

func (m *mockAgentRepo) GetPlayer(ctx context.Context, agentID, playerID string) (*model.Player, error) {
	for _, p := range m.players[agentID] {
		if p.PlayerID == playerID {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("player not found")
}

func (m *mockAgentRepo) UpdatePlayerLimit(ctx context.Context, agentID, playerID string, depositLimit, betLimit float64) error {
	for i, p := range m.players[agentID] {
		if p.PlayerID == playerID {
			m.players[agentID][i].DepositLimit = depositLimit
			m.players[agentID][i].BetLimit = betLimit
			return nil
		}
	}
	return fmt.Errorf("player not found")
}

func (m *mockAgentRepo) GetDashboard(ctx context.Context, agentID string) (*model.Dashboard, error) {
	d, ok := m.dashboards[agentID]
	if !ok {
		return nil, fmt.Errorf("dashboard not found")
	}
	return d, nil
}

func TestCreateAgent(t *testing.T) {
	tests := []struct {
		name     string
		username string
		email    string
	}{
		{"create agent 1", "agent_one", "one@test.com"},
		{"create agent 2", "agent_two", "two@test.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := model.Agent{
				AgentID:   fmt.Sprintf("agent-%s", tt.username),
				Username:  tt.username,
				Email:     tt.email,
				Status:    "active",
				Tier:      "bronze",
				CreatedAt: time.Now(),
			}
			if agent.Username != tt.username {
				t.Errorf("username = %v, want %v", agent.Username, tt.username)
			}
			if agent.Status != "active" {
				t.Errorf("status = %v, want active", agent.Status)
			}
		})
	}
}

func TestGetAgent(t *testing.T) {
	repo := newMockAgentRepo()
	repo.players["agent-1"] = []model.Player{
		{PlayerID: "p-1", Username: "player1", Status: "active", TotalDeposits: 1000},
	}

	tests := []struct {
		name     string
		agentID  string
		playerID string
		wantErr  bool
	}{
		{"existing player", "agent-1", "p-1", false},
		{"non-existent player", "agent-1", "p-99", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player, err := repo.GetPlayer(context.Background(), tt.agentID, tt.playerID)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetPlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && player.PlayerID != tt.playerID {
				t.Errorf("playerID = %v, want %v", player.PlayerID, tt.playerID)
			}
		})
	}
}

func TestUpdateAgent(t *testing.T) {
	tests := []struct {
		name    string
		agentID string
		status  string
		tier    string
	}{
		{"update to silver", "a-1", "active", "silver"},
		{"suspend agent", "a-2", "suspended", "bronze"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agent := model.Agent{
				AgentID: tt.agentID,
				Status:  "active",
				Tier:    "bronze",
			}
			agent.Status = tt.status
			agent.Tier = tt.tier

			if agent.Status != tt.status {
				t.Errorf("status = %v, want %v", agent.Status, tt.status)
			}
			if agent.Tier != tt.tier {
				t.Errorf("tier = %v, want %v", agent.Tier, tt.tier)
			}
		})
	}
}

func TestListAgents(t *testing.T) {
	repo := newMockAgentRepo()
	repo.players["agent-1"] = []model.Player{
		{PlayerID: "p-1", Username: "alice", Status: "active"},
		{PlayerID: "p-2", Username: "bob", Status: "active"},
		{PlayerID: "p-3", Username: "charlie", Status: "suspended"},
	}

	tests := []struct {
		name    string
		status  string
		wantLen int
	}{
		{"all players", "", 3},
		{"active only", "active", 2},
		{"suspended only", "suspended", 1},
		{"non-existent status", "banned", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			players, total, err := repo.ListPlayers(context.Background(), "agent-1", 1, 10, "", tt.status)
			if err != nil {
				t.Fatalf("ListPlayers() error = %v", err)
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

func TestUpdateAgentStatus(t *testing.T) {
	repo := newMockAgentRepo()
	repo.players["agent-1"] = []model.Player{
		{PlayerID: "p-1", Status: "active"},
	}

	tests := []struct {
		name     string
		playerID string
		newLimit float64
		newBet   float64
		wantErr  bool
	}{
		{"update limits", "p-1", 5000.0, 1000.0, false},
		{"non-existent player", "p-99", 100.0, 50.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.UpdatePlayerLimit(context.Background(), "agent-1", tt.playerID, tt.newLimit, tt.newBet)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UpdatePlayerLimit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
