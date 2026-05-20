package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/game_engine/gateway/common/handler"

	bonuspb "github.com/game_engine/common-service/proto/gen/go/bonus/v1"
	commissionpb "github.com/game_engine/common-service/proto/gen/go/commission/v1"
)

// Commission Claims Handlers
func (cfg *RouterConfig) ListCommissionClaims(ctx context.Context, c *app.RequestContext) {
	if cfg.CommissionClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"claims":     []interface{}{},
			"total":      0,
			"page":       1,
			"totalPages": 0,
		})
		return
	}

	resp, err := cfg.CommissionClient.GetClaimsByStatus(ctx, &commissionpb.GetClaimsByStatusRequest{
		Status: "PENDING",
		Page:   1,
		Limit:  50,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to list commission claims: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims":     resp.Claims,
		"total":      resp.Total,
		"page":       resp.Page,
		"totalPages": resp.TotalPages,
	})
}

func (cfg *RouterConfig) GetCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")

	if cfg.CommissionClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"id":          claimID,
			"claimType":   "COMMISSION",
			"amount":      "0.00",
			"status":      "UNKNOWN",
			"claimReason": "",
		})
		return
	}

	resp, err := cfg.CommissionClient.GetUserClaims(ctx, &commissionpb.GetUserClaimsRequest{
		ClaimId: claimID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to get commission claim: %v", err))
		return
	}

	if len(resp.Claims) == 0 {
		handler.SendJSONError(c, 404, handler.ErrCodeNotFound, "commission claim not found")
		return
	}

	handler.SendSuccess(c, resp.Claims[0])
}

func (cfg *RouterConfig) ApproveCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")

	if cfg.CommissionClient == nil {
		handler.SendJSONError(c, 503, handler.ErrCodeServiceUnavailable, "commission service unavailable")
		return
	}

	resp, err := cfg.CommissionClient.ClaimCommission(ctx, &commissionpb.ClaimCommissionRequest{
		ClaimId: claimID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to approve commission claim: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "APPROVED",
		"message": resp.Message,
	})
}

func (cfg *RouterConfig) RejectCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")

	if cfg.CommissionClient == nil {
		handler.SendJSONError(c, 503, handler.ErrCodeServiceUnavailable, "commission service unavailable")
		return
	}

	resp, err := cfg.CommissionClient.ClaimCommission(ctx, &commissionpb.ClaimCommissionRequest{
		ClaimId: claimID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to reject commission claim: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "REJECTED",
		"message": resp.Message,
	})
}

func (cfg *RouterConfig) PayCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")

	if cfg.CommissionClient == nil {
		handler.SendJSONError(c, 503, handler.ErrCodeServiceUnavailable, "commission service unavailable")
		return
	}

	resp, err := cfg.CommissionClient.ClaimCommission(ctx, &commissionpb.ClaimCommissionRequest{
		ClaimId: claimID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to pay commission claim: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "PAID",
		"message": resp.Message,
	})
}

// Rebet Claims Handlers
func (cfg *RouterConfig) ListRebetClaims(ctx context.Context, c *app.RequestContext) {
	if cfg.BonusClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"claims": []interface{}{},
			"total":  0,
		})
		return
	}

	resp, err := cfg.BonusClient.GetUserRebetClaims(ctx, &bonuspb.GetUserRebetClaimsRequest{})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to list rebet claims: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
		"total":  len(resp.Claims),
	})
}

func (cfg *RouterConfig) GetRebetClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")

	if cfg.BonusClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"id":     claimID,
			"amount": "0.00",
			"status": "UNKNOWN",
		})
		return
	}

	resp, err := cfg.BonusClient.GetClaimableRebets(ctx, &bonuspb.GetClaimableRebetsRequest{})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to get rebet claim: %v", err))
		return
	}

	for _, claim := range resp.Claims {
		if claim.RebetId == claimID {
			handler.SendSuccess(c, claim)
			return
		}
	}

	handler.SendJSONError(c, 404, handler.ErrCodeNotFound, "rebet claim not found")
}

func (cfg *RouterConfig) ApproveRebetClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")

	if cfg.BonusClient == nil {
		handler.SendJSONError(c, 503, handler.ErrCodeServiceUnavailable, "bonus service unavailable")
		return
	}

	resp, err := cfg.BonusClient.ClaimRebet(ctx, &bonuspb.ClaimRebetRequest{
		RebetId: claimID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to approve rebet claim: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "APPROVED",
		"message": "Rebet claim approved",
		"success": resp.Success,
	})
}

func (cfg *RouterConfig) RejectRebetClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "REJECTED",
		"message": "Rebet claim rejected",
	})
}
