package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	gamepb "github.com/game_engine/common-service/proto/gen/go/game/v1"

	"common/handler"
)

// GetSlotGames handles getting slot games
func (cfg *RouterConfig) GetSlotGames(ctx context.Context, c *app.RequestContext) {
	cfg.listGamesByCategory(ctx, c, "slot")
}

// GetCardGames handles getting card games
func (cfg *RouterConfig) GetCardGames(ctx context.Context, c *app.RequestContext) {
	cfg.listGamesByCategory(ctx, c, "card")
}

// GetDiceGames handles getting dice games
func (cfg *RouterConfig) GetDiceGames(ctx context.Context, c *app.RequestContext) {
	cfg.listGamesByCategory(ctx, c, "dice")
}

func (cfg *RouterConfig) listGamesByCategory(ctx context.Context, c *app.RequestContext, category string) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.ListGames(ctx, &gamepb.ListGamesRequest{
		Category: category,
		Page:     c.DefaultQuery("page", "1"),
		Limit:    c.DefaultQuery("limit", "20"),
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": resp.Games,
		"total": resp.Total,
	})
}
