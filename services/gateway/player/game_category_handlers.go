package main

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	commonpb "github.com/game_engine/common-service/proto/gen/go/common/v1"
	gamepb "github.com/game_engine/common-service/proto/gen/go/game/v1"

	"github.com/game_engine/gateway/common/handler"
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

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "20")
	page, _ := strconv.ParseInt(pageStr, 10, 32)
	limit, _ := strconv.ParseInt(limitStr, 10, 32)

	resp, err := cfg.GameClient.ListGames(ctx, &gamepb.ListGamesRequest{
		CategoryId: category,
		Pagination: &commonpb.PaginationRequest{
			Page:     int32(page),
			PageSize: int32(limit),
		},
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	total := int32(0)
	if resp.Pagination != nil {
		total = resp.Pagination.TotalItems
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": resp.Games,
		"total": total,
	})
}
