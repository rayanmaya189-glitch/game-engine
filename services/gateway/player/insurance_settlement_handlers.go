package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	bonuspb "github.com/game_engine/common-service/proto/gen/go/bonus/v1"
	commissionpb "github.com/game_engine/common-service/proto/gen/go/commission/v1"

	"common/handler"
)

// SubmitInsuranceClaim handles submitting an insurance claim
func (cfg *RouterConfig) SubmitInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	var req struct {
		GameID            string  `json:"gameId"`
		BetID             string  `json:"betId"`
		InsurancePolicyID string  `json:"insurancePolicyId"`
		ClaimType         string  `json:"claimType"`
		InsuredAmount     float64 `json:"insuredAmount"`
		LossAmount        float64 `json:"lossAmount"`
		ClaimReason       string  `json:"claimReason"`
		EvidenceDetails   string  `json:"evidenceDetails"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.SubmitInsuranceClaim(ctx, &bonuspb.SubmitInsuranceClaimRequest{
		UserId:            userID,
		GameId:            req.GameID,
		BetId:             req.BetID,
		InsurancePolicyId: req.InsurancePolicyID,
		ClaimType:         req.ClaimType,
		InsuredAmount:     req.InsuredAmount,
		LossAmount:        req.LossAmount,
		ClaimReason:       req.ClaimReason,
		EvidenceDetails:   req.EvidenceDetails,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":  "Insurance claim submitted",
		"claim_id": resp.ClaimId,
		"status":   resp.Status,
	})
}

// GetUserInsuranceClaims handles getting user's insurance claims
func (cfg *RouterConfig) GetUserInsuranceClaims(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.GetUserInsuranceClaims(ctx, &bonuspb.GetUserInsuranceClaimsRequest{
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

// GetUserSettlements handles getting user's settlements
func (cfg *RouterConfig) GetUserSettlements(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetUserSettlements(ctx, &commissionpb.GetUserSettlementsRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"settlements": resp.Settlements,
	})
}

// GetSettlementById handles getting settlement by ID
func (cfg *RouterConfig) GetSettlementById(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	settlementID := c.Param("id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetSettlementById(ctx, &commissionpb.GetSettlementByIdRequest{
		SettlementId: settlementID,
		UserId:       userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Settlement not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"settlement": resp.Settlement,
	})
}

// GetUserTotalPending handles getting user's total pending claims
func (cfg *RouterConfig) GetUserTotalPending(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetTotalPending(ctx, &commissionpb.GetTotalPendingRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"totalPending": resp.TotalPending,
	})
}

// GetUserTotalSettled handles getting user's total settled claims
func (cfg *RouterConfig) GetUserTotalSettled(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetTotalSettled(ctx, &commissionpb.GetTotalSettledRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"totalSettled": resp.TotalSettled,
	})
}
