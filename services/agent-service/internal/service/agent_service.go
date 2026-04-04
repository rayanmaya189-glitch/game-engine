package service

import (
	"context"

	"github.com/game_engine/agent-service/internal/config"
	"github.com/game_engine/agent-service/internal/repository"
	agentpb "github.com/game_engine/agent-service/pkg/game_engine/agent/v1"
)

type AgentService struct {
	repo *repository.AgentRepository
	cfg  *config.Config
}

func NewAgentService(repo *repository.AgentRepository, cfg *config.Config) *AgentService {
	return &AgentService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *AgentService) ListPlayers(ctx context.Context, req *agentpb.ListPlayersRequest) (*agentpb.ListPlayersResponse, error) {
	page := int(req.Page)
	if page < 1 {
		page = 1
	}
	limit := int(req.Limit)
	if limit < 1 || limit > 100 {
		limit = 20
	}

	players, total, err := s.repo.ListPlayers(ctx, req.AgentId, page, limit, req.Search, req.Status)
	if err != nil {
		return nil, err
	}

	pbPlayers := make([]*agentpb.Player, len(players))
	for i, p := range players {
		pbPlayers[i] = &agentpb.Player{
			PlayerId:      p.PlayerID,
			Username:      p.Username,
			Email:         p.Email,
			Status:        p.Status,
			TotalDeposits: p.TotalDeposits,
			TotalBets:     p.TotalBets,
			Balance:       p.Balance,
			CreatedAt:     p.CreatedAt.Unix(),
		}
	}

	return &agentpb.ListPlayersResponse{
		Players: pbPlayers,
		Total:   int32(total),
		Page:    int32(page),
	}, nil
}

func (s *AgentService) GetPlayer(ctx context.Context, req *agentpb.GetPlayerRequest) (*agentpb.GetPlayerResponse, error) {
	player, err := s.repo.GetPlayer(ctx, req.AgentId, req.PlayerId)
	if err != nil {
		return nil, err
	}

	return &agentpb.GetPlayerResponse{
		Player: &agentpb.Player{
			PlayerId:      player.PlayerID,
			Username:      player.Username,
			Email:         player.Email,
			Status:        player.Status,
			TotalDeposits: player.TotalDeposits,
			TotalBets:     player.TotalBets,
			Balance:       player.Balance,
			CreatedAt:     player.CreatedAt.Unix(),
		},
	}, nil
}

func (s *AgentService) UpdatePlayerLimit(ctx context.Context, req *agentpb.UpdatePlayerLimitRequest) (*agentpb.UpdatePlayerLimitResponse, error) {
	err := s.repo.UpdatePlayerLimit(ctx, req.AgentId, req.PlayerId, req.DepositLimit, req.BetLimit)
	if err != nil {
		return &agentpb.UpdatePlayerLimitResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &agentpb.UpdatePlayerLimitResponse{
		Success: true,
		Message: "Player limits updated successfully",
	}, nil
}

func (s *AgentService) GetDashboard(ctx context.Context, req *agentpb.GetDashboardRequest) (*agentpb.GetDashboardResponse, error) {
	dashboard, err := s.repo.GetDashboard(ctx, req.AgentId)
	if err != nil {
		return nil, err
	}

	return &agentpb.GetDashboardResponse{
		TotalPlayers:      int32(dashboard.TotalPlayers),
		ActivePlayers:     int32(dashboard.ActivePlayers),
		TotalCommission:   dashboard.TotalCommission,
		PendingCommission: dashboard.PendingCommission,
	}, nil
}
