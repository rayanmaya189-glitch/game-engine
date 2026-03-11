package handler

import (
	"context"

	merchantpb "github.com/game-engine/gen/go/gameengine/merchant/v1"
	"github.com/game-engine/merchant-service/internal/service"
)

var _ merchantpb.MerchantServiceServer = (*MerchantHandler)(nil)

type MerchantHandler struct {
	service *service.MerchantService
}

func NewMerchantHandler(svc *service.MerchantService) *MerchantHandler {
	return &MerchantHandler{service: svc}
}

func (h *MerchantHandler) ListPlayers(ctx context.Context, req *merchantpb.ListPlayersRequest) (*merchantpb.ListPlayersResponse, error) {
	return h.service.ListPlayers(ctx, req)
}

func (h *MerchantHandler) GetPlayer(ctx context.Context, req *merchantpb.GetPlayerRequest) (*merchantpb.GetPlayerResponse, error) {
	return h.service.GetPlayer(ctx, req)
}

func (h *MerchantHandler) GetRevenueReport(ctx context.Context, req *merchantpb.GetRevenueReportRequest) (*merchantpb.GetRevenueReportResponse, error) {
	return h.service.GetRevenueReport(ctx, req)
}

func (h *MerchantHandler) GetPlayerReport(ctx context.Context, req *merchantpb.GetPlayerReportRequest) (*merchantpb.GetPlayerReportResponse, error) {
	return h.service.GetPlayerReport(ctx, req)
}

func (h *MerchantHandler) GetGameReport(ctx context.Context, req *merchantpb.GetGameReportRequest) (*merchantpb.GetGameReportResponse, error) {
	return h.service.GetGameReport(ctx, req)
}

func (h *MerchantHandler) GetConfig(ctx context.Context, req *merchantpb.GetConfigRequest) (*merchantpb.GetConfigResponse, error) {
	return h.service.GetConfig(ctx, req)
}

func (h *MerchantHandler) UpdateConfig(ctx context.Context, req *merchantpb.UpdateConfigRequest) (*merchantpb.UpdateConfigResponse, error) {
	return h.service.UpdateConfig(ctx, req)
}

func (h *MerchantHandler) RegisterWebhook(ctx context.Context, req *merchantpb.RegisterWebhookRequest) (*merchantpb.RegisterWebhookResponse, error) {
	return h.service.RegisterWebhook(ctx, req)
}

func (h *MerchantHandler) ListWebhooks(ctx context.Context, req *merchantpb.ListWebhooksRequest) (*merchantpb.ListWebhooksResponse, error) {
	return h.service.ListWebhooks(ctx, req)
}

func (h *MerchantHandler) DeleteWebhook(ctx context.Context, req *merchantpb.DeleteWebhookRequest) (*merchantpb.DeleteWebhookResponse, error) {
	return h.service.DeleteWebhook(ctx, req)
}

func (h *MerchantHandler) ListAgents(ctx context.Context, req *merchantpb.ListAgentsRequest) (*merchantpb.ListAgentsResponse, error) {
	return h.service.ListAgents(ctx, req)
}

func (h *MerchantHandler) GetAgent(ctx context.Context, req *merchantpb.GetAgentRequest) (*merchantpb.GetAgentResponse, error) {
	return h.service.GetAgent(ctx, req)
}

func (h *MerchantHandler) CreateAgent(ctx context.Context, req *merchantpb.CreateAgentRequest) (*merchantpb.CreateAgentResponse, error) {
	return h.service.CreateAgent(ctx, req)
}

func (h *MerchantHandler) UpdateAgent(ctx context.Context, req *merchantpb.UpdateAgentRequest) (*merchantpb.UpdateAgentResponse, error) {
	return h.service.UpdateAgent(ctx, req)
}

func (h *MerchantHandler) UpdateAgentStatus(ctx context.Context, req *merchantpb.UpdateAgentStatusRequest) (*merchantpb.UpdateAgentStatusResponse, error) {
	return h.service.UpdateAgentStatus(ctx, req)
}
