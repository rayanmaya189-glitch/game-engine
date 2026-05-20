package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	commissionpb "github.com/game_engine/common-service/proto/gen/go/commission/v1"

	"github.com/game_engine/gateway/common/handler"
)

// SubmitCommissionClaim handles submitting a commission claim
func (cfg *RouterConfig) SubmitCommissionClaim(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	var req struct {
		AffiliateID  string  `json:"affiliateId"`
		CommissionID string  `json:"commissionId"`
		Amount       float64 `json:"amount"`
		ClaimReason  string  `json:"claimReason"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.SubmitClaim(ctx, &commissionpb.SubmitClaimRequest{
		UserId:       userID,
		AffiliateId:  req.AffiliateID,
		CommissionId: req.CommissionID,
		Amount:       req.Amount,
		ClaimReason:  req.ClaimReason,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":  "Commission claim submitted",
		"claim_id": resp.ClaimId,
		"status":   resp.Status,
	})
}

// GetUserCommissionClaims handles getting user's commission claims
func (cfg *RouterConfig) GetUserCommissionClaims(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetUserClaims(ctx, &commissionpb.GetUserClaimsRequest{
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

// GetCommissionClaimsByStatus handles getting commission claims by status
func (cfg *RouterConfig) GetCommissionClaimsByStatus(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	status := c.Param("status")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetClaimsByStatus(ctx, &commissionpb.GetClaimsByStatusRequest{
		UserId: userID,
		Status: status,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
		"status": status,
	})
}

// ClaimCommission handles claiming a commission
func (cfg *RouterConfig) ClaimCommission(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	claimID := c.Param("id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.ClaimCommission(ctx, &commissionpb.ClaimCommissionRequest{
		ClaimId: claimID,
		UserId:  userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Commission claimed",
		"claim_id":       claimID,
		"amount":         resp.Amount,
		"transaction_id": resp.TransactionId,
	})
}
