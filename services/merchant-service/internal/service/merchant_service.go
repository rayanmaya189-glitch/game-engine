package service

import (
	"context"
	"strconv"

	merchantpb "github.com/game-engine/gen/go/gameengine/merchant/v1"
	"github.com/game-engine/merchant-service/internal/config"
	"github.com/game-engine/merchant-service/internal/repository"
)

type MerchantService struct {
	repo *repository.MerchantRepository
	cfg  *config.Config
}

func NewMerchantService(repo *repository.MerchantRepository, cfg *config.Config) *MerchantService {
	return &MerchantService{repo: repo, cfg: cfg}
}

func (s *MerchantService) ListPlayers(ctx context.Context, req *merchantpb.ListPlayersRequest) (*merchantpb.ListPlayersResponse, error) {
	page := int(req.Page)
	if page < 1 {
		page = 1
	}
	if page > 10000 {
		page = 10000
	}
	limit := int(req.Limit)
	if limit < 1 || limit > 100 {
		limit = 20
	}

	players, total, err := s.repo.ListPlayers(ctx, req.MerchantId, page, limit, req.Search)
	if err != nil {
		return nil, err
	}

	pbPlayers := make([]*merchantpb.MerchantPlayer, len(players))
	for i, p := range players {
		pbPlayers[i] = &merchantpb.MerchantPlayer{
			PlayerId: p.PlayerID,
			Username: p.Username,
			Email:    p.Email,
			Status:   p.Status,
		}
	}
	return &merchantpb.ListPlayersResponse{Players: pbPlayers, Total: int32(total)}, nil
}

func (s *MerchantService) GetPlayer(ctx context.Context, req *merchantpb.GetPlayerRequest) (*merchantpb.GetPlayerResponse, error) {
	p, err := s.repo.GetPlayer(ctx, req.MerchantId, req.PlayerId)
	if err != nil {
		return nil, err
	}
	return &merchantpb.GetPlayerResponse{
		Player: &merchantpb.MerchantPlayer{
			PlayerId: p.PlayerID,
			Username: p.Username,
			Email:    p.Email,
			Status:   p.Status,
		},
	}, nil
}

func (s *MerchantService) GetRevenueReport(ctx context.Context, req *merchantpb.GetRevenueReportRequest) (*merchantpb.GetRevenueReportResponse, error) {
	report, err := s.repo.GetRevenueReport(ctx, req.MerchantId, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}
	return &merchantpb.GetRevenueReportResponse{
		TotalRevenue:     report.TotalRevenue,
		TotalDeposits:    report.TotalDeposits,
		TotalWithdrawals: report.TotalWithdrawals,
		TotalPlayers:     int32(report.TotalPlayers),
	}, nil
}

func (s *MerchantService) GetPlayerReport(ctx context.Context, req *merchantpb.GetPlayerReportRequest) (*merchantpb.GetPlayerReportResponse, error) {
	report, err := s.repo.GetPlayerReport(ctx, req.MerchantId, req.PlayerId, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}
	return &merchantpb.GetPlayerReportResponse{
		TotalBets:   report.TotalBets,
		TotalWins:   report.TotalWins,
		NetRevenue:  report.NetRevenue,
		GamesPlayed: int32(report.GamesPlayed),
	}, nil
}

func (s *MerchantService) GetGameReport(ctx context.Context, req *merchantpb.GetGameReportRequest) (*merchantpb.GetGameReportResponse, error) {
	report, err := s.repo.GetGameReport(ctx, req.MerchantId, req.GameId, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}
	return &merchantpb.GetGameReportResponse{
		TotalBets:    report.TotalBets,
		TotalWins:    report.TotalWins,
		TotalPlayers: int32(report.TotalPlayers),
		Plays:        int32(report.Plays),
	}, nil
}

func (s *MerchantService) GetConfig(ctx context.Context, req *merchantpb.GetConfigRequest) (*merchantpb.GetConfigResponse, error) {
	config, err := s.repo.GetConfig(ctx, req.MerchantId)
	if err != nil {
		return nil, err
	}
	return &merchantpb.GetConfigResponse{Config: config}, nil
}

func (s *MerchantService) UpdateConfig(ctx context.Context, req *merchantpb.UpdateConfigRequest) (*merchantpb.UpdateConfigResponse, error) {
	config := make(map[string]string)
	if req.CommissionRate > 0 {
		config["commission_rate"] = strconv.FormatFloat(req.CommissionRate, 'f', 2, 64)
	}
	err := s.repo.UpdateConfig(ctx, req.MerchantId, config)
	if err != nil {
		return nil, err
	}
	return &merchantpb.UpdateConfigResponse{Success: true}, nil
}

func (s *MerchantService) RegisterWebhook(ctx context.Context, req *merchantpb.RegisterWebhookRequest) (*merchantpb.RegisterWebhookResponse, error) {
	webhookID, err := s.repo.RegisterWebhook(ctx, req.MerchantId, req.Url, req.Events)
	if err != nil {
		return nil, err
	}
	return &merchantpb.RegisterWebhookResponse{WebhookId: webhookID}, nil
}

func (s *MerchantService) ListWebhooks(ctx context.Context, req *merchantpb.ListWebhooksRequest) (*merchantpb.ListWebhooksResponse, error) {
	webhooks, err := s.repo.ListWebhooks(ctx, req.MerchantId)
	if err != nil {
		return nil, err
	}
	pbWebhooks := make([]*merchantpb.Webhook, len(webhooks))
	for i, w := range webhooks {
		pbWebhooks[i] = &merchantpb.Webhook{
			WebhookId: w.WebhookID,
			Url:       w.URL,
			Events:    w.Events,
			Status:    w.Status,
		}
	}
	return &merchantpb.ListWebhooksResponse{Webhooks: pbWebhooks}, nil
}

func (s *MerchantService) DeleteWebhook(ctx context.Context, req *merchantpb.DeleteWebhookRequest) (*merchantpb.DeleteWebhookResponse, error) {
	err := s.repo.DeleteWebhook(ctx, req.MerchantId, req.WebhookId)
	if err != nil {
		return nil, err
	}
	return &merchantpb.DeleteWebhookResponse{Success: true}, nil
}

func (s *MerchantService) ListAgents(ctx context.Context, req *merchantpb.ListAgentsRequest) (*merchantpb.ListAgentsResponse, error) {
	page := int(req.Page)
	if page < 1 {
		page = 1
	}
	if page > 10000 {
		page = 10000
	}
	limit := int(req.Limit)
	if limit < 1 || limit > 100 {
		limit = 20
	}

	agents, total, err := s.repo.ListAgents(ctx, req.MerchantId, page, limit)
	if err != nil {
		return nil, err
	}

	pbAgents := make([]*merchantpb.Agent, len(agents))
	for i, a := range agents {
		pbAgents[i] = &merchantpb.Agent{
			AgentId:  a.AgentID,
			Username: a.Username,
			Email:    a.Email,
			Status:   a.Status,
		}
	}
	return &merchantpb.ListAgentsResponse{Agents: pbAgents, Total: int32(total)}, nil
}

func (s *MerchantService) GetAgent(ctx context.Context, req *merchantpb.GetAgentRequest) (*merchantpb.GetAgentResponse, error) {
	a, err := s.repo.GetAgent(ctx, req.MerchantId, req.AgentId)
	if err != nil {
		return nil, err
	}
	return &merchantpb.GetAgentResponse{
		Agent: &merchantpb.Agent{
			AgentId:  a.AgentID,
			Username: a.Username,
			Email:    a.Email,
			Status:   a.Status,
		},
	}, nil
}

func (s *MerchantService) CreateAgent(ctx context.Context, req *merchantpb.CreateAgentRequest) (*merchantpb.CreateAgentResponse, error) {
	agentID, err := s.repo.CreateAgent(ctx, req.MerchantId, req.Username, req.Email, req.SendInvitation)
	if err != nil {
		return nil, err
	}
	return &merchantpb.CreateAgentResponse{AgentId: agentID, Success: true}, nil
}

func (s *MerchantService) UpdateAgent(ctx context.Context, req *merchantpb.UpdateAgentRequest) (*merchantpb.UpdateAgentResponse, error) {
	err := s.repo.UpdateAgent(ctx, req.MerchantId, req.AgentId, req.Username, req.Email)
	if err != nil {
		return nil, err
	}
	return &merchantpb.UpdateAgentResponse{Success: true}, nil
}

func (s *MerchantService) UpdateAgentStatus(ctx context.Context, req *merchantpb.UpdateAgentStatusRequest) (*merchantpb.UpdateAgentStatusResponse, error) {
	err := s.repo.UpdateAgentStatus(ctx, req.MerchantId, req.AgentId, req.Status)
	if err != nil {
		return nil, err
	}
	return &merchantpb.UpdateAgentStatusResponse{Success: true}, nil
}
