package handler

import (
	"context"

	leaderv1 "github.com/game_engine/leaderboard-service/gen/go/leaderboard/v1"
	"github.com/game_engine/leaderboard-service/internal/model"
	"github.com/game_engine/leaderboard-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LeaderboardHandler struct {
	leaderv1.UnimplementedLeaderboardServiceServer
	service *service.LeaderboardService
}

func NewLeaderboardHandler(s *service.LeaderboardService) *LeaderboardHandler {
	return &LeaderboardHandler{service: s}
}

func (h *LeaderboardHandler) GetDailyLeaderboard(ctx context.Context, req *leaderv1.GetLeaderboardRequest) (*leaderv1.GetLeaderboardResponse, error) {
	resp, err := h.service.GetDailyLeaderboard(ctx, req.GetGameType())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get daily leaderboard: %v", err)
	}
	return leaderboardResponseToProto(resp), nil
}

func (h *LeaderboardHandler) GetWeeklyLeaderboard(ctx context.Context, req *leaderv1.GetLeaderboardRequest) (*leaderv1.GetLeaderboardResponse, error) {
	resp, err := h.service.GetWeeklyLeaderboard(ctx, req.GetGameType())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get weekly leaderboard: %v", err)
	}
	return leaderboardResponseToProto(resp), nil
}

func (h *LeaderboardHandler) GetMonthlyLeaderboard(ctx context.Context, req *leaderv1.GetLeaderboardRequest) (*leaderv1.GetLeaderboardResponse, error) {
	resp, err := h.service.GetMonthlyLeaderboard(ctx, req.GetGameType())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get monthly leaderboard: %v", err)
	}
	return leaderboardResponseToProto(resp), nil
}

func (h *LeaderboardHandler) GetAllTimeLeaderboard(ctx context.Context, req *leaderv1.GetLeaderboardRequest) (*leaderv1.GetLeaderboardResponse, error) {
	resp, err := h.service.GetAllTimeLeaderboard(ctx, req.GetGameType())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all-time leaderboard: %v", err)
	}
	return leaderboardResponseToProto(resp), nil
}

func (h *LeaderboardHandler) GetPlayerRank(ctx context.Context, req *leaderv1.GetPlayerRankRequest) (*leaderv1.GetPlayerRankResponse, error) {
	lbType := model.LeaderboardType(req.GetLeaderboardType())
	var resp *model.PlayerRankResponse
	var err error

	switch lbType {
	case model.LeaderboardTypeDaily:
		resp, err = h.service.GetPlayerDailyRank(ctx, req.GetUserId(), req.GetGameType())
	case model.LeaderboardTypeWeekly:
		resp, err = h.service.GetPlayerWeeklyRank(ctx, req.GetUserId(), req.GetGameType())
	case model.LeaderboardTypeMonthly:
		resp, err = h.service.GetPlayerMonthlyRank(ctx, req.GetUserId(), req.GetGameType())
	case model.LeaderboardTypeAllTime:
		resp, err = h.service.GetPlayerAllTimeRank(ctx, req.GetUserId(), req.GetGameType())
	default:
		resp, err = h.service.GetPlayerDailyRank(ctx, req.GetUserId(), req.GetGameType())
	}

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "player rank not found: %v", err)
	}

	return &leaderv1.GetPlayerRankResponse{
		UserId:   resp.UserID,
		Username: resp.Username,
		Rank:     int32(resp.Rank),
		Score:    resp.Score,
		Type:     string(resp.LeaderboardType),
		Period:   resp.Period,
	}, nil
}

func (h *LeaderboardHandler) UpdatePlayerScore(ctx context.Context, req *leaderv1.UpdatePlayerScoreRequest) (*leaderv1.UpdatePlayerScoreResponse, error) {
	scoreReq := model.UpdateScoreRequest{
		UserID:    req.GetUserId(),
		Username:  req.GetUsername(),
		Score:     req.GetScore(),
		GameType:  req.GetGameType(),
		IsWin:     req.GetIsWin(),
		WinAmount: req.GetWinAmount(),
		BetAmount: req.GetBetAmount(),
	}

	if err := h.service.UpdatePlayerScore(ctx, scoreReq); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update score: %v", err)
	}

	return &leaderv1.UpdatePlayerScoreResponse{
		Success: true,
		Message: "Score updated successfully",
	}, nil
}

func (h *LeaderboardHandler) DistributePrizes(ctx context.Context, req *leaderv1.DistributePrizesRequest) (*leaderv1.DistributePrizesResponse, error) {
	prizeReq := model.PrizeDistributionRequest{
		LeaderboardType: model.LeaderboardType(req.GetLeaderboardType()),
		GameType:        req.GetGameType(),
		TournamentID:    req.GetTournamentId(),
		DryRun:          req.GetDryRun(),
	}

	result, err := h.service.DistributePrizes(ctx, prizeReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to distribute prizes: %v", err)
	}

	prizes := make([]*leaderv1.Prize, len(result.Prizes))
	for i, p := range result.Prizes {
		prizes[i] = &leaderv1.Prize{
			Rank:     int32(p.Rank),
			Type:     p.Type,
			Value:    p.Value,
			Currency: p.Currency,
		}
	}

	return &leaderv1.DistributePrizesResponse{
		LeaderboardType: string(result.LeaderboardType),
		GameType:        result.GameType,
		Period:          result.Period,
		Prizes:          prizes,
		TotalValue:      result.TotalValue,
		DistributedAt:   &result.DistributedAt,
	}, nil
}

func (h *LeaderboardHandler) SyncLeaderboard(ctx context.Context, req *leaderv1.SyncLeaderboardRequest) (*leaderv1.SyncLeaderboardResponse, error) {
	err := h.service.SyncLeaderboard(ctx, model.LeaderboardType(req.GetLeaderboardType()), req.GetGameType())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to sync leaderboard: %v", err)
	}

	return &leaderv1.SyncLeaderboardResponse{
		Success: true,
		Message: "Leaderboard synced successfully",
	}, nil
}

func (h *LeaderboardHandler) ResetLeaderboard(ctx context.Context, req *leaderv1.ResetLeaderboardRequest) (*leaderv1.ResetLeaderboardResponse, error) {
	err := h.service.ResetLeaderboard(ctx, model.LeaderboardType(req.GetLeaderboardType()), req.GetGameType())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to reset leaderboard: %v", err)
	}

	return &leaderv1.ResetLeaderboardResponse{
		Success: true,
		Message: "Leaderboard reset successfully",
	}, nil
}

func leaderboardResponseToProto(resp *model.LeaderboardResponse) *leaderv1.GetLeaderboardResponse {
	entries := make([]*leaderv1.LeaderboardEntry, len(resp.Entries))
	for i, e := range resp.Entries {
		updatedAt := e.UpdatedAt
		entries[i] = &leaderv1.LeaderboardEntry{
			Rank:      int32(e.Rank),
			UserId:    e.UserID,
			Username:  e.Username,
			Score:     e.Score,
			Wins:      int32(e.Wins),
			WinAmount: e.WinAmount,
			GameType:  e.GameType,
			UpdatedAt: &updatedAt,
		}
	}

	updatedAt := resp.UpdatedAt
	return &leaderv1.GetLeaderboardResponse{
		Type:      string(resp.Type),
		Period:    resp.Period,
		Entries:   entries,
		Total:     int32(resp.Total),
		UpdatedAt: &updatedAt,
	}
}
