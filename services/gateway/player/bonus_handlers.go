package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	bonuspb "github.com/game_engine/common-service/proto/gen/go/bonus/v1"

	"github.com/game_engine/gateway/common/handler"
)

// ListBonuses handles listing available bonuses
func (cfg *RouterConfig) ListBonuses(ctx context.Context, c *app.RequestContext) {
	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.ListBonuses(ctx, &bonuspb.ListBonusesRequest{
		Status: "active",
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"bonuses": resp.Bonuses,
	})
}

// GetBonus handles getting bonus details
func (cfg *RouterConfig) GetBonus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.GetBonus(ctx, &bonuspb.GetBonusRequest{
		BonusId: bonusID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Bonus not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"bonus": resp.Bonus,
	})
}

// ClaimBonus handles claiming a bonus
func (cfg *RouterConfig) ClaimBonus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")
	userID := c.GetString("user_id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.ClaimBonus(ctx, &bonuspb.ClaimBonusRequest{
		BonusId: bonusID,
		UserId:  userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":      "Bonus claimed successfully",
		"bonus_amount": resp.BonusAmount,
		"expires_at":   resp.ExpiresAt,
	})
}

// GetMyBonuses handles getting user's claimed bonuses
func (cfg *RouterConfig) GetMyBonuses(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.GetUserBonuses(ctx, &bonuspb.GetUserBonusesRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"bonuses": resp.Bonuses,
	})
}
