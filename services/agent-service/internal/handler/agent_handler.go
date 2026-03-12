package handler

import (
	"context"

	"github.com/game_engine/agent-service/internal/service"
	agentpb "github.com/game_engine/gen/go/game_engine/agent/v1"
)

// Ensure AgentHandler implements the server interface
var _ agentpb.AgentServiceServer = (*AgentHandler)(nil)

type AgentHandler struct {
	service *service.AgentService
}

func NewAgentHandler(svc *service.AgentService) *AgentHandler {
	return &AgentHandler{
		service: svc,
	}
}

func (h *AgentHandler) ListPlayers(ctx context.Context, req *agentpb.ListPlayersRequest) (*agentpb.ListPlayersResponse, error) {
	return h.service.ListPlayers(ctx, req)
}

func (h *AgentHandler) GetPlayer(ctx context.Context, req *agentpb.GetPlayerRequest) (*agentpb.GetPlayerResponse, error) {
	return h.service.GetPlayer(ctx, req)
}

func (h *AgentHandler) UpdatePlayerLimit(ctx context.Context, req *agentpb.UpdatePlayerLimitRequest) (*agentpb.UpdatePlayerLimitResponse, error) {
	return h.service.UpdatePlayerLimit(ctx, req)
}

func (h *AgentHandler) GetDashboard(ctx context.Context, req *agentpb.GetDashboardRequest) (*agentpb.GetDashboardResponse, error) {
	return h.service.GetDashboard(ctx, req)
}
