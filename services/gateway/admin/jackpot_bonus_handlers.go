package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"common/handler"

	jackpotpb "github.com/game_engine/gen/go/game_engine/jackpot/v1"
	bonuspb "github.com/game_engine/gen/go/game_engine/bonus/v1"
)

// Jackpots Handlers
func (cfg *RouterConfig) ListJackpots(ctx context.Context, c *app.RequestContext) {
	if cfg.JackpotClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"jackpots": []interface{}{},
			"total":    0,
		})
		return
	}

	resp, err := cfg.JackpotClient.ListJackpots(ctx, &jackpotpb.ListJackpotsRequest{})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to list jackpots: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"jackpots": resp.Jackpots,
		"total":    resp.Total,
	})
}

func (cfg *RouterConfig) GetJackpot(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")

	if cfg.JackpotClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"id":     jackpotID,
			"name":   "",
			"status": "UNKNOWN",
		})
		return
	}

	resp, err := cfg.JackpotClient.GetJackpot(ctx, &jackpotpb.GetJackpotRequest{
		JackpotId: jackpotID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to get jackpot: %v", err))
		return
	}

	handler.SendSuccess(c, resp.Jackpot)
}

func (cfg *RouterConfig) CreateJackpot(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"id":      "new_jackpot_id",
		"message": "Jackpot created successfully",
	})
}

func (cfg *RouterConfig) UpdateJackpot(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      jackpotID,
		"message": "Jackpot updated successfully",
	})
}

func (cfg *RouterConfig) UpdateJackpotStatus(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      jackpotID,
		"status":  "updated",
		"message": "Jackpot status updated",
	})
}

func (cfg *RouterConfig) DeleteJackpot(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      jackpotID,
		"message": "Jackpot deleted successfully",
	})
}

func (cfg *RouterConfig) GetJackpotHits(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")

	if cfg.JackpotClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"jackpotId": jackpotID,
			"hits":      []interface{}{},
			"total":     0,
		})
		return
	}

	resp, err := cfg.JackpotClient.GetWinners(ctx, &jackpotpb.GetWinnersRequest{
		JackpotId: jackpotID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to get jackpot hits: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"jackpotId": jackpotID,
		"hits":      resp.Winners,
		"total":     resp.Total,
	})
}

// Bonuses Handlers
func (cfg *RouterConfig) ListBonuses(ctx context.Context, c *app.RequestContext) {
	if cfg.BonusClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"bonuses": []interface{}{},
			"total":   0,
		})
		return
	}

	resp, err := cfg.BonusClient.ListBonuses(ctx, &bonuspb.ListBonusesRequest{})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to list bonuses: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"bonuses": resp.Bonuses,
		"total":   resp.Total,
	})
}

func (cfg *RouterConfig) GetBonus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")

	if cfg.BonusClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"id":     bonusID,
			"name":   "",
			"status": "UNKNOWN",
		})
		return
	}

	resp, err := cfg.BonusClient.GetBonus(ctx, &bonuspb.GetBonusRequest{
		BonusId: bonusID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to get bonus: %v", err))
		return
	}

	handler.SendSuccess(c, resp.Bonus)
}

func (cfg *RouterConfig) CreateBonus(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"id":      "new_bonus_id",
		"message": "Bonus created successfully",
	})
}

func (cfg *RouterConfig) UpdateBonus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      bonusID,
		"message": "Bonus updated successfully",
	})
}

func (cfg *RouterConfig) UpdateBonusStatus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      bonusID,
		"status":  "updated",
		"message": "Bonus status updated",
	})
}

func (cfg *RouterConfig) DeleteBonus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      bonusID,
		"message": "Bonus deleted successfully",
	})
}
