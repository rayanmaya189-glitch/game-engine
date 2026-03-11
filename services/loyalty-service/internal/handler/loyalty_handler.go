package handler

import (
	"context"

	"github.com/game-engine/loyalty-service/internal/service"
)

type LoyaltyHandler struct {
	service *service.LoyaltyService
}

func NewLoyaltyHandler(svc *service.LoyaltyService) *LoyaltyHandler {
	return &LoyaltyHandler{service: svc}
}

func (h *LoyaltyHandler) GetMember(ctx context.Context, userID string) (*service.GetMemberResponse, error) {
	return h.service.GetMember(ctx, userID)
}

func (h *LoyaltyHandler) GetPointsHistory(ctx context.Context, userID string, limit, offset int) (*service.GetPointsHistoryResponse, error) {
	return h.service.GetPointsHistory(ctx, userID, limit, offset)
}

func (h *LoyaltyHandler) GetTiers(ctx context.Context) (*service.GetTiersResponse, error) {
	return h.service.GetTiers(ctx)
}

func (h *LoyaltyHandler) GetRewards(ctx context.Context) (*service.GetRewardsResponse, error) {
	return h.service.GetRewards(ctx)
}

func (h *LoyaltyHandler) RedeemReward(ctx context.Context, userID, rewardID string, pointsCost int) (*service.RedeemRewardResponse, error) {
	return h.service.RedeemReward(ctx, userID, rewardID, pointsCost)
}

func (h *LoyaltyHandler) GetLeaderboard(ctx context.Context, limit int) (*service.GetLeaderboardResponse, error) {
	return h.service.GetLeaderboard(ctx, limit)
}

func (h *LoyaltyHandler) AddPoints(ctx context.Context, userID, betID string, betAmount float64) (*service.AddPointsResponse, error) {
	return h.service.AddPoints(ctx, userID, betID, betAmount)
}
