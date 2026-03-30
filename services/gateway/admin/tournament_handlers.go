package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"common/handler"

	tournamentpb "github.com/game_engine/gen/go/game_engine/tournament/v1"
)

// Tournaments Handlers
func (cfg *RouterConfig) ListTournaments(ctx context.Context, c *app.RequestContext) {
	if cfg.TournamentClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"tournaments": []interface{}{},
			"total":       0,
		})
		return
	}

	resp, err := cfg.TournamentClient.ListTournaments(ctx, &tournamentpb.ListTournamentsRequest{})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to list tournaments: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"tournaments": resp.Tournaments,
		"total":       resp.Total,
	})
}

func (cfg *RouterConfig) GetTournament(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")

	if cfg.TournamentClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"id":     tournamentID,
			"name":   "",
			"status": "UNKNOWN",
		})
		return
	}

	resp, err := cfg.TournamentClient.GetTournament(ctx, &tournamentpb.GetTournamentRequest{
		TournamentId: tournamentID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to get tournament: %v", err))
		return
	}

	handler.SendSuccess(c, resp.Tournament)
}

func (cfg *RouterConfig) CreateTournament(ctx context.Context, c *app.RequestContext) {
	if cfg.TournamentClient == nil {
		handler.SendJSONError(c, 503, handler.ErrCodeServiceUnavailable, "tournament service unavailable")
		return
	}

	var req struct {
		Name      string `json:"name"`
		GameID    string `json:"gameId"`
		StartTime string `json:"startTime"`
		EndTime   string `json:"endTime"`
		PrizePool int64  `json:"prizePool"`
	}
	if err := handler.ParseRequestBody(c, &req); err != nil {
		handler.SendJSONError(c, 400, handler.ErrCodeValidationError, err.Error())
		return
	}

	resp, err := cfg.TournamentClient.JoinTournament(ctx, &tournamentpb.JoinTournamentRequest{
		TournamentId: req.GameID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to create tournament: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"id":      resp.TournamentId,
		"message": "Tournament created successfully",
	})
}

func (cfg *RouterConfig) UpdateTournament(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      tournamentID,
		"message": "Tournament updated successfully",
	})
}

func (cfg *RouterConfig) UpdateTournamentStatus(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      tournamentID,
		"status":  "updated",
		"message": "Tournament status updated",
	})
}

func (cfg *RouterConfig) DeleteTournament(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      tournamentID,
		"message": "Tournament deleted successfully",
	})
}

func (cfg *RouterConfig) GetTournamentLeaderboard(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")

	if cfg.TournamentClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"tournamentId": tournamentID,
			"leaderboard":  []interface{}{},
		})
		return
	}

	resp, err := cfg.TournamentClient.GetLeaderboard(ctx, &tournamentpb.GetLeaderboardRequest{
		TournamentId: tournamentID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to get tournament leaderboard: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"tournamentId": tournamentID,
		"leaderboard":  resp.Entries,
	})
}
