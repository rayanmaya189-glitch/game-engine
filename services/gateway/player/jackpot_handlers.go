package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	jackpotpb "github.com/game_engine/common-service/proto/gen/go/jackpot/v1"

	"common/handler"
)

// ListJackpots handles listing jackpots
func (cfg *RouterConfig) ListJackpots(ctx context.Context, c *app.RequestContext) {
	if cfg.JackpotClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Jackpot service unavailable", nil)
		return
	}

	resp, err := cfg.JackpotClient.ListJackpots(ctx, &jackpotpb.ListJackpotsRequest{
		Status: c.Query("status"),
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"jackpots": resp.Jackpots,
	})
}

// GetJackpot handles getting jackpot details
func (cfg *RouterConfig) GetJackpot(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")

	if cfg.JackpotClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Jackpot service unavailable", nil)
		return
	}

	resp, err := cfg.JackpotClient.GetJackpot(ctx, &jackpotpb.GetJackpotRequest{
		JackpotId: jackpotID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Jackpot not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"jackpot": resp.Jackpot,
	})
}

// GetJackpotWinners handles getting jackpot winners
func (cfg *RouterConfig) GetJackpotWinners(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")

	if cfg.JackpotClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Jackpot service unavailable", nil)
		return
	}

	resp, err := cfg.JackpotClient.GetWinners(ctx, &jackpotpb.GetWinnersRequest{
		JackpotId: jackpotID,
		Limit:     20,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"winners": resp.Winners,
	})
}
