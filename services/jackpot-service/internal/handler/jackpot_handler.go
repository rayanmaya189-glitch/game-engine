package handler

import (
	"context"

	jackpotpb "github.com/game_engine/gen/go/game_engine/jackpot/v1"
	"github.com/game_engine/jackpot-service/internal/service"
)

var _ jackpotpb.JackpotServiceServer = (*JackpotHandler)(nil)

type JackpotHandler struct {
	service *service.JackpotService
}

func NewJackpotHandler(svc *service.JackpotService) *JackpotHandler {
	return &JackpotHandler{service: svc}
}

func (h *JackpotHandler) ListJackpots(ctx context.Context, req *jackpotpb.ListJackpotsRequest) (*jackpotpb.ListJackpotsResponse, error) {
	return h.service.ListJackpots(ctx, req)
}

func (h *JackpotHandler) GetJackpot(ctx context.Context, req *jackpotpb.GetJackpotRequest) (*jackpotpb.GetJackpotResponse, error) {
	return h.service.GetJackpot(ctx, req)
}

func (h *JackpotHandler) GetWinners(ctx context.Context, req *jackpotpb.GetWinnersRequest) (*jackpotpb.GetWinnersResponse, error) {
	return h.service.GetWinners(ctx, req)
}

func (h *JackpotHandler) JoinJackpot(ctx context.Context, req *jackpotpb.JoinJackpotRequest) (*jackpotpb.JoinJackpotResponse, error) {
	return h.service.JoinJackpot(ctx, req)
}

func (h *JackpotHandler) GetJackpotHistory(ctx context.Context, req *jackpotpb.GetJackpotHistoryRequest) (*jackpotpb.GetJackpotHistoryResponse, error) {
	return h.service.GetJackpotHistory(ctx, req)
}
