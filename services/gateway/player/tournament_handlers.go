package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	tournamentpb "github.com/game_engine/common-service/proto/gen/go/tournament/v1"

	"github.com/game_engine/gateway/common/handler"
)

// ListTournaments handles listing tournaments
func (cfg *RouterConfig) ListTournaments(ctx context.Context, c *app.RequestContext) {
	if cfg.TournamentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Tournament service unavailable", nil)
		return
	}

	resp, err := cfg.TournamentClient.ListTournaments(ctx, &tournamentpb.ListTournamentsRequest{
		Status: c.Query("status"),
		Page:   c.DefaultQuery("page", "1"),
		Limit:  c.DefaultQuery("limit", "20"),
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"tournaments": resp.Tournaments,
		"total":       resp.Total,
	})
}

// GetTournament handles getting tournament details
func (cfg *RouterConfig) GetTournament(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")

	if cfg.TournamentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Tournament service unavailable", nil)
		return
	}

	resp, err := cfg.TournamentClient.GetTournament(ctx, &tournamentpb.GetTournamentRequest{
		TournamentId: tournamentID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Tournament not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"tournament": resp.Tournament,
	})
}

// JoinTournament handles joining a tournament
func (cfg *RouterConfig) JoinTournament(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")
	userID := c.GetString("user_id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.TournamentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Tournament service unavailable", nil)
		return
	}

	resp, err := cfg.TournamentClient.JoinTournament(ctx, &tournamentpb.JoinTournamentRequest{
		TournamentId: tournamentID,
		UserId:       userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":       "Joined tournament successfully",
		"tournament_id": tournamentID,
		"position":      resp.Position,
	})
}

// GetTournamentLeaderboard handles getting tournament leaderboard
func (cfg *RouterConfig) GetTournamentLeaderboard(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")

	if cfg.TournamentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Tournament service unavailable", nil)
		return
	}

	resp, err := cfg.TournamentClient.GetLeaderboard(ctx, &tournamentpb.GetLeaderboardRequest{
		TournamentId: tournamentID,
		Limit:        50,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"leaderboard": resp.Entries,
	})
}
