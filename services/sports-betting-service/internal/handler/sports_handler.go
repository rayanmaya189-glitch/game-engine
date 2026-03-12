package handler

import (
	"context"

	"github.com/game_engine/sports-betting-service/internal/service"
)

type SportsHandler struct {
	service *service.SportsService
}

func NewSportsHandler(svc *service.SportsService) *SportsHandler {
	return &SportsHandler{service: svc}
}

func (h *SportsHandler) GetSports(ctx context.Context) (*service.GetSportsResponse, error) {
	return h.service.GetSports(ctx)
}

func (h *SportsHandler) GetLiveEvents(ctx context.Context) (*service.GetLiveEventsResponse, error) {
	return h.service.GetLiveEvents(ctx)
}

func (h *SportsHandler) GetUpcomingEvents(ctx context.Context, sportID string, limit int) (*service.GetUpcomingEventsResponse, error) {
	return h.service.GetUpcomingEvents(ctx, sportID, limit)
}

func (h *SportsHandler) GetMarkets(ctx context.Context, eventID string) (*service.GetMarketsResponse, error) {
	return h.service.GetMarkets(ctx, eventID)
}

func (h *SportsHandler) PlaceBet(ctx context.Context, userID, eventID, marketID, selection string, stake, odds float64) (*service.PlaceBetResponse, error) {
	return h.service.PlaceBet(ctx, userID, eventID, marketID, selection, stake, odds)
}

func (h *SportsHandler) GetUserBets(ctx context.Context, userID string, page, limit int) (*service.GetUserBetsResponse, error) {
	return h.service.GetUserBets(ctx, userID, page, limit)
}
