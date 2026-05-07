package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	gamepb "github.com/game_engine/common-service/proto/gen/go/game/v1"

	"common/handler"
)

// ListGames handles listing games
func (cfg *RouterConfig) ListGames(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")
	category := c.Query("category")
	provider := c.Query("provider")
	search := c.Query("search")

	resp, err := cfg.GameClient.ListGames(ctx, &gamepb.ListGamesRequest{
		Page:     page,
		Limit:    limit,
		Category: category,
		Provider: provider,
		Search:   search,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	games := make([]map[string]interface{}, len(resp.Games))
	for i, game := range resp.Games {
		games[i] = map[string]interface{}{
			"id":          game.Id,
			"name":        game.Name,
			"provider":    game.Provider,
			"category":    game.Category,
			"thumbnail":   game.Thumbnail,
			"rtp":         game.Rtp,
			"min_bet":     game.MinBet,
			"max_bet":     game.MaxBet,
			"volatility":  game.Volatility,
			"status":      game.Status,
			"is_featured": game.IsFeatured,
			"is_new":      game.IsNew,
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": games,
		"total": resp.Total,
		"page":  resp.Page,
	})
}

// GetGame handles getting game details
func (cfg *RouterConfig) GetGame(ctx context.Context, c *app.RequestContext) {
	gameID := c.Param("id")

	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.GetGame(ctx, &gamepb.GetGameRequest{
		GameId: gameID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Game not found", nil)
		return
	}

	game := resp.Game
	handler.SendSuccess(c, map[string]interface{}{
		"game": map[string]interface{}{
			"id":          game.Id,
			"name":        game.Name,
			"provider":    game.Provider,
			"category":    game.Category,
			"description": game.Description,
			"thumbnail":   game.Thumbnail,
			"images":      game.Images,
			"rtp":         game.Rtp,
			"min_bet":     game.MinBet,
			"max_bet":     game.MaxBet,
			"volatility":  game.Volatility,
			"features":    game.Features,
			"paylines":    game.Paylines,
			"reels":       game.Reels,
			"status":      game.Status,
		},
	})
}

// PlayGame handles game play URL generation
func (cfg *RouterConfig) PlayGame(ctx context.Context, c *app.RequestContext) {
	gameID := c.Param("id")
	userID := c.GetString("user_id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.GetGameURL(ctx, &gamepb.GetGameURLRequest{
		GameId: gameID,
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"game_id": gameID,
		"url":     resp.Url,
		"token":   resp.Token,
	})
}

// GetCategories handles getting game categories
func (cfg *RouterConfig) GetCategories(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.GetCategories(ctx, &gamepb.GetCategoriesRequest{})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	categories := make([]map[string]interface{}, len(resp.Categories))
	for i, cat := range resp.Categories {
		categories[i] = map[string]interface{}{
			"id":         cat.Id,
			"name":       cat.Name,
			"slug":       cat.Slug,
			"icon":       cat.Icon,
			"game_count": cat.GameCount,
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"categories": categories,
	})
}

// GetFeaturedGames handles getting featured games
func (cfg *RouterConfig) GetFeaturedGames(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.GetFeaturedGames(ctx, &gamepb.GetFeaturedGamesRequest{
		Limit: 10,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	games := make([]map[string]interface{}, len(resp.Games))
	for i, game := range resp.Games {
		games[i] = map[string]interface{}{
			"id":        game.Id,
			"name":      game.Name,
			"provider":  game.Provider,
			"category":  game.Category,
			"thumbnail": game.Thumbnail,
			"rtp":       game.Rtp,
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": games,
	})
}

// GetPopularGames handles getting popular games
func (cfg *RouterConfig) GetPopularGames(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.GetPopularGames(ctx, &gamepb.GetPopularGamesRequest{
		Limit: 20,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	games := make([]map[string]interface{}, len(resp.Games))
	for i, game := range resp.Games {
		games[i] = map[string]interface{}{
			"id":         game.Id,
			"name":       game.Name,
			"provider":   game.Provider,
			"category":   game.Category,
			"thumbnail":  game.Thumbnail,
			"play_count": game.PlayCount,
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": games,
	})
}
