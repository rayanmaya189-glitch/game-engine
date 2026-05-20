package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/game_engine/gateway/common/handler"
)

func (cfg *RouterConfig) ListAdminGames(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"games": []interface{}{},
	})
}

func (cfg *RouterConfig) CreateGame(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"game_id": "new_game_id",
		"message": "Game created successfully",
	})
}

func (cfg *RouterConfig) UpdateGame(ctx context.Context, c *app.RequestContext) {
	gameID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"game_id": gameID,
		"message": "Game updated successfully",
	})
}

func (cfg *RouterConfig) GetAllTransactions(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"transactions": []interface{}{},
	})
}

func (cfg *RouterConfig) AdjustBalance(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Balance adjusted successfully",
		"transaction_id": "adj_123",
	})
}

func (cfg *RouterConfig) ReportsHandler(ctx context.Context, c *app.RequestContext) {
	path := string(c.Request.URI().Path())
	handler.SendSuccess(c, map[string]interface{}{
		"path":    path,
		"message": "Reports endpoint",
	})
}
