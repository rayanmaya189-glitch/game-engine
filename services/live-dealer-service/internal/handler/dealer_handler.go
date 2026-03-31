package handler

import (
	"context"
	"fmt"
	"io"

	"github.com/game_engine/live-dealer-service/internal/model"
	"github.com/game_engine/live-dealer-service/internal/service"
)

type LiveDealerHandler struct {
	svc *service.LiveDealerService
}

func NewLiveDealerHandler(svc *service.LiveDealerService) *LiveDealerHandler {
	return &LiveDealerHandler{svc: svc}
}

type CreateSessionRequest struct {
	GameType string  `json:"game_type"`
	DealerID string  `json:"dealer_id"`
	MinBet   float64 `json:"min_bet"`
	MaxBet   float64 `json:"max_bet"`
}

type CreateSessionResponse struct {
	Table *model.Table `json:"table"`
}

func (h *LiveDealerHandler) CreateSession(ctx context.Context, req *CreateSessionRequest) (*CreateSessionResponse, error) {
	table, err := h.svc.CreateSession(ctx, req.GameType, req.DealerID, req.MinBet, req.MaxBet)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	return &CreateSessionResponse{Table: table}, nil
}

type JoinSessionRequest struct {
	TableID  string  `json:"table_id"`
	PlayerID string  `json:"player_id"`
	Chips    float64 `json:"chips"`
}

type JoinSessionResponse struct {
	Player *model.Player `json:"player"`
}

func (h *LiveDealerHandler) JoinSession(ctx context.Context, req *JoinSessionRequest) (*JoinSessionResponse, error) {
	player, err := h.svc.JoinSession(ctx, req.TableID, req.PlayerID, req.Chips)
	if err != nil {
		return nil, fmt.Errorf("failed to join session: %w", err)
	}
	return &JoinSessionResponse{Player: player}, nil
}

type LeaveSessionRequest struct {
	TableID  string `json:"table_id"`
	PlayerID string `json:"player_id"`
}

type LeaveSessionResponse struct {
	Success bool `json:"success"`
}

func (h *LiveDealerHandler) LeaveSession(ctx context.Context, req *LeaveSessionRequest) (*LeaveSessionResponse, error) {
	if err := h.svc.LeaveSession(ctx, req.TableID, req.PlayerID); err != nil {
		return nil, fmt.Errorf("failed to leave session: %w", err)
	}
	return &LeaveSessionResponse{Success: true}, nil
}

type EndSessionRequest struct {
	TableID string `json:"table_id"`
}

type EndSessionResponse struct {
	Success bool `json:"success"`
}

func (h *LiveDealerHandler) EndSession(ctx context.Context, req *EndSessionRequest) (*EndSessionResponse, error) {
	if err := h.svc.EndSession(ctx, req.TableID); err != nil {
		return nil, fmt.Errorf("failed to end session: %w", err)
	}
	return &EndSessionResponse{Success: true}, nil
}

type GetSessionRequest struct {
	TableID string `json:"table_id"`
}

type GetSessionResponse struct {
	Table   *model.Table    `json:"table"`
	Players []*model.Player `json:"players"`
}

func (h *LiveDealerHandler) GetSession(ctx context.Context, req *GetSessionRequest) (*GetSessionResponse, error) {
	table, err := h.svc.GetSession(ctx, req.TableID)
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	players, err := h.svc.GetSessionPlayers(ctx, req.TableID)
	if err != nil {
		return nil, fmt.Errorf("failed to get players: %w", err)
	}

	return &GetSessionResponse{Table: table, Players: players}, nil
}

type ListSessionsRequest struct {
	GameType string `json:"game_type"`
	Status   string `json:"status"`
}

type ListSessionsResponse struct {
	Sessions []*model.Table `json:"sessions"`
}

func (h *LiveDealerHandler) ListSessions(ctx context.Context, req *ListSessionsRequest) (*ListSessionsResponse, error) {
	tables, err := h.svc.ListSessions(ctx, req.GameType, req.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to list sessions: %w", err)
	}
	return &ListSessionsResponse{Sessions: tables}, nil
}

type StreamVideoRequest struct {
	TableID  string `json:"table_id"`
	PlayerID string `json:"player_id"`
}

type StreamVideoResponse struct {
	StreamURL   string `json:"stream_url"`
	Bitrate     int    `json:"bitrate"`
	ViewerCount int    `json:"viewer_count"`
	IsLive      bool   `json:"is_live"`
}

func (h *LiveDealerHandler) StreamVideo(ctx context.Context, req *StreamVideoRequest) (*StreamVideoResponse, error) {
	info, err := h.svc.GetStreamInfo(ctx, req.TableID)
	if err != nil {
		return nil, fmt.Errorf("failed to get stream info: %w", err)
	}
	return &StreamVideoResponse{
		StreamURL:   info.StreamURL,
		Bitrate:     info.Bitrate,
		ViewerCount: info.ViewerCount,
		IsLive:      info.IsLive,
	}, nil
}

type StreamVideoBidirectionalServer interface {
	Send(*StreamVideoResponse) error
	Recv() (*StreamVideoRequest, error)
	Context() context.Context
}

func (h *LiveDealerHandler) StreamVideoBidirectional(stream StreamVideoBidirectionalServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("stream receive error: %w", err)
		}

		resp, err := h.StreamVideo(stream.Context(), req)
		if err != nil {
			return fmt.Errorf("stream processing error: %w", err)
		}

		if err := stream.Send(resp); err != nil {
			return fmt.Errorf("stream send error: %w", err)
		}
	}
}
