package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	bonuspb "github.com/game_engine/common-service/proto/gen/go/bonus/v1"

	"github.com/game_engine/gateway/common/handler"
)

// CreateRebetClaim handles creating a rebet claim
func (cfg *RouterConfig) CreateRebetClaim(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	var req struct {
		BonusID          string  `json:"bonusId"`
		BonusCode        string  `json:"bonusCode"`
		BonusAmount      float64 `json:"bonusAmount"`
		RebetRequirement float64 `json:"rebetRequirement"`
		GameID           string  `json:"gameId"`
		BetID            string  `json:"betId"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.CreateRebetClaim(ctx, &bonuspb.CreateRebetClaimRequest{
		UserId:           userID,
		BonusId:          req.BonusID,
		BonusCode:        req.BonusCode,
		BonusAmount:      req.BonusAmount,
		RebetRequirement: req.RebetRequirement,
		GameId:           req.GameID,
		BetId:            req.BetID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":           "Rebet claim created",
		"rebet_id":          resp.RebetId,
		"status":            resp.Status,
		"rebet_requirement": resp.RebetRequirement,
		"current_rebet":     resp.CurrentRebet,
	})
}

// GetUserRebetClaims handles getting user's rebet claims
func (cfg *RouterConfig) GetUserRebetClaims(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.GetUserRebetClaims(ctx, &bonuspb.GetUserRebetClaimsRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
	})
}

// GetClaimableRebets handles getting claimable rebet bonuses
func (cfg *RouterConfig) GetClaimableRebets(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.GetClaimableRebets(ctx, &bonuspb.GetClaimableRebetsRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
	})
}

// ClaimRebet handles claiming a rebet bonus
func (cfg *RouterConfig) ClaimRebet(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	rebetID := c.Param("id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.ClaimRebet(ctx, &bonuspb.ClaimRebetRequest{
		RebetId: rebetID,
		UserId:  userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Rebet bonus claimed",
		"rebet_id":       rebetID,
		"amount":         resp.Amount,
		"transaction_id": resp.TransactionId,
	})
}
