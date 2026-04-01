package handler

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/game_engine/referral-service/internal/model"
	"github.com/game_engine/referral-service/internal/service"
)

type ReferralHandler struct {
	svc *service.ReferralService
}

func NewReferralHandler(svc *service.ReferralService) *ReferralHandler {
	return &ReferralHandler{svc: svc}
}

func (h *ReferralHandler) RegisterServices(server *grpc.Server) {}

func (h *ReferralHandler) GetReferralCode(ctx context.Context, playerID string) (*model.ReferralCode, error) {
	if playerID == "" {
		return nil, status.Error(codes.InvalidArgument, "player_id is required")
	}

	code, err := h.svc.GetReferralCode(playerID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "referral code not found: %v", err)
	}

	return code, nil
}

func (h *ReferralHandler) GenerateReferralCode(ctx context.Context, playerID string) (*model.ReferralCode, error) {
	if playerID == "" {
		return nil, status.Error(codes.InvalidArgument, "player_id is required")
	}

	code, err := h.svc.GenerateReferralCode(playerID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate code: %v", err)
	}

	return code, nil
}

func (h *ReferralHandler) TrackReferral(ctx context.Context, req []byte) (*model.Referral, error) {
	var input struct {
		Code      string `json:"code"`
		RefereeID string `json:"referee_id"`
		Source    string `json:"source"`
	}

	if err := json.Unmarshal(req, &input); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}

	referral, err := h.svc.TrackReferral(input.Code, input.RefereeID, input.Source)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to track referral: %v", err)
	}

	return referral, nil
}

func (h *ReferralHandler) GetReferralHistory(ctx context.Context, playerID string, page, pageSize int) (*model.ReferralList, error) {
	if playerID == "" {
		return nil, status.Error(codes.InvalidArgument, "player_id is required")
	}

	filter := &model.ReferralFilter{
		ReferrerID: playerID,
		Page:       page,
		PageSize:   pageSize,
	}

	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PageSize < 1 {
		filter.PageSize = 20
	}

	return h.svc.GetReferrals(playerID, filter)
}

func (h *ReferralHandler) GetReferralStats(ctx context.Context, playerID string) (*model.ReferralStats, error) {
	if playerID == "" {
		return nil, status.Error(codes.InvalidArgument, "player_id is required")
	}

	stats, err := h.svc.GetReferralStats(playerID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get stats: %v", err)
	}

	return stats, nil
}

func (h *ReferralHandler) GetAvailableRewards(ctx context.Context) ([]*model.ReferralReward, error) {
	rewards, err := h.svc.GetAvailableRewards()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get rewards: %v", err)
	}
	return rewards, nil
}

func (h *ReferralHandler) ClaimReward(ctx context.Context, referralID, playerID string) error {
	if referralID == "" || playerID == "" {
		return status.Error(codes.InvalidArgument, "referral_id and player_id are required")
	}

	if err := h.svc.ClaimReward(referralID, playerID); err != nil {
		return status.Errorf(codes.Internal, "failed to claim reward: %v", err)
	}

	return nil
}

func (h *ReferralHandler) QualifyReferral(ctx context.Context, referralID string) error {
	if referralID == "" {
		return status.Error(codes.InvalidArgument, "referral_id is required")
	}

	if err := h.svc.QualifyReferral(referralID); err != nil {
		return status.Errorf(codes.Internal, "failed to qualify referral: %v", err)
	}

	return nil
}

func (h *ReferralHandler) RewardReferral(ctx context.Context, referralID string, amount float64, rewardType string) error {
	if referralID == "" {
		return status.Error(codes.InvalidArgument, "referral_id is required")
	}

	if err := h.svc.RewardReferral(referralID, amount, model.RewardType(rewardType)); err != nil {
		return status.Errorf(codes.Internal, "failed to reward referral: %v", err)
	}

	return nil
}

func (h *ReferralHandler) TrackClick(ctx context.Context, code string) error {
	if code == "" {
		return status.Error(codes.InvalidArgument, "code is required")
	}

	if err := h.svc.TrackClick(code); err != nil {
		return status.Errorf(codes.Internal, "failed to track click: %v", err)
	}

	return nil
}
