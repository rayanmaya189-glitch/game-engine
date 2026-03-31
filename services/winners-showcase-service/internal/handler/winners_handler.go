package handler

import (
	"context"

	winnersv1 "github.com/game_engine/winners-showcase-service/gen/go/winners/v1"
	"github.com/game_engine/winners-showcase-service/internal/model"
	"github.com/game_engine/winners-showcase-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WinnersHandler struct {
	winnersv1.UnimplementedWinnersServiceServer
	service *service.WinnersService
}

func NewWinnersHandler(s *service.WinnersService) *WinnersHandler {
	return &WinnersHandler{service: s}
}

func (h *WinnersHandler) GetRecentWinners(ctx context.Context, req *winnersv1.GetRecentWinnersRequest) (*winnersv1.GetRecentWinnersResponse, error) {
	resp, err := h.service.GetRecentWinners(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get recent winners: %v", err)
	}

	winners := make([]*winnersv1.Winner, len(resp.Winners))
	for i, w := range resp.Winners {
		ts := w.Timestamp
		winners[i] = &winnersv1.Winner{
			Id:          int64(w.ID),
			UserId:      w.UserID,
			Username:    w.Username,
			DisplayName: w.DisplayName,
			WinAmount:   w.WinAmount,
			Currency:    w.Currency,
			GameType:    w.GameType,
			GameName:    w.GameName,
			WinType:     string(w.WinType),
			Multiplier:  w.Multiplier,
			Timestamp:   &ts,
		}
	}

	updatedAt := resp.UpdatedAt
	return &winnersv1.GetRecentWinnersResponse{
		Winners:   winners,
		Total:     int32(resp.Total),
		UpdatedAt: &updatedAt,
	}, nil
}

func (h *WinnersHandler) GetBigWins(ctx context.Context, req *winnersv1.GetBigWinsRequest) (*winnersv1.GetBigWinsResponse, error) {
	resp, err := h.service.GetBigWins(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get big wins: %v", err)
	}

	wins := make([]*winnersv1.Winner, len(resp.Wins))
	for i, w := range resp.Wins {
		ts := w.Timestamp
		wins[i] = &winnersv1.Winner{
			Id:          int64(w.ID),
			UserId:      w.UserID,
			Username:    w.Username,
			DisplayName: w.DisplayName,
			WinAmount:   w.WinAmount,
			Currency:    w.Currency,
			GameType:    w.GameType,
			GameName:    w.GameName,
			WinType:     string(w.WinType),
			Multiplier:  w.Multiplier,
			Timestamp:   &ts,
		}
	}

	updatedAt := resp.UpdatedAt
	return &winnersv1.GetBigWinsResponse{
		Wins:      wins,
		Threshold: resp.Threshold,
		Total:     int32(resp.Total),
		UpdatedAt: &updatedAt,
	}, nil
}

func (h *WinnersHandler) GetJackpotWinners(ctx context.Context, req *winnersv1.GetJackpotWinnersRequest) (*winnersv1.GetJackpotWinnersResponse, error) {
	resp, err := h.service.GetJackpotWinners(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get jackpot winners: %v", err)
	}

	winners := make([]*winnersv1.Winner, len(resp.Winners))
	for i, w := range resp.Winners {
		ts := w.Timestamp
		winners[i] = &winnersv1.Winner{
			Id:          int64(w.ID),
			UserId:      w.UserID,
			Username:    w.Username,
			DisplayName: w.DisplayName,
			WinAmount:   w.WinAmount,
			Currency:    w.Currency,
			GameType:    w.GameType,
			GameName:    w.GameName,
			WinType:     string(w.WinType),
			Multiplier:  w.Multiplier,
			Timestamp:   &ts,
		}
	}

	updatedAt := resp.UpdatedAt
	return &winnersv1.GetJackpotWinnersResponse{
		Winners:   winners,
		Threshold: resp.Threshold,
		Total:     int32(resp.Total),
		UpdatedAt: &updatedAt,
	}, nil
}

func (h *WinnersHandler) RecordWin(ctx context.Context, req *winnersv1.RecordWinRequest) (*winnersv1.RecordWinResponse, error) {
	recordReq := model.RecordWinRequest{
		UserID:     req.GetUserId(),
		Username:   req.GetUsername(),
		WinAmount:  req.GetWinAmount(),
		Currency:   req.GetCurrency(),
		GameType:   req.GetGameType(),
		GameName:   req.GetGameName(),
		Multiplier: req.GetMultiplier(),
	}

	if err := h.service.RecordWin(ctx, recordReq); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to record win: %v", err)
	}

	return &winnersv1.RecordWinResponse{
		Success: true,
		Message: "Win recorded successfully",
	}, nil
}

func (h *WinnersHandler) GetPrivacySettings(ctx context.Context, req *winnersv1.GetPrivacySettingsRequest) (*winnersv1.GetPrivacySettingsResponse, error) {
	resp, err := h.service.GetPrivacySettings(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get privacy settings: %v", err)
	}

	return &winnersv1.GetPrivacySettingsResponse{
		UserId:            resp.UserID,
		AnonymizeName:     resp.AnonymizeName,
		ShowOnLeaderboard: resp.ShowOnLeaderboard,
		ShowOnJackpotList: resp.ShowOnJackpotList,
		OptOutOfShowcase:  resp.OptOutOfShowcase,
	}, nil
}

func (h *WinnersHandler) UpdatePrivacySettings(ctx context.Context, req *winnersv1.UpdatePrivacySettingsRequest) (*winnersv1.UpdatePrivacySettingsResponse, error) {
	updateReq := model.UpdatePrivacyRequest{
		AnonymizeName:     req.GetAnonymizeName(),
		ShowOnLeaderboard: req.GetShowOnLeaderboard(),
		ShowOnJackpotList: req.GetShowOnJackpotList(),
		OptOutOfShowcase:  req.GetOptOutOfShowcase(),
	}

	if err := h.service.UpdatePrivacySettings(ctx, req.GetUserId(), updateReq); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update privacy settings: %v", err)
	}

	return &winnersv1.UpdatePrivacySettingsResponse{
		Success: true,
		Message: "Privacy settings updated",
	}, nil
}
