package main

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	commissionpb "github.com/game_engine/common-service/proto/gen/go/commission/v1"

	"github.com/game_engine/gateway/common/handler"
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

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	uID, _ := strconv.ParseInt(userID, 10, 64)
	gID, _ := strconv.ParseInt(req.GameID, 10, 64)
	bID, _ := strconv.ParseInt(req.BetID, 10, 64)

	resp, err := cfg.CommissionClient.SubmitInsuranceClaim(ctx, &commissionpb.SubmitInsuranceClaimRequest{
		UserId:            uID,
		GameId:            gID,
		BetId:             bID,
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
		"claim_id": resp.Claim.Id,
		"status":   resp.Claim.Status,
	})
}

// GetUserInsuranceClaims handles getting user's insurance claims
func (cfg *RouterConfig) GetUserInsuranceClaims(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	uID, _ := strconv.ParseInt(userID, 10, 64)

	resp, err := cfg.CommissionClient.GetUserInsuranceClaims(ctx, &commissionpb.GetUserInsuranceClaimsRequest{
		UserId: uID,
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

	uID, _ := strconv.ParseInt(userID, 10, 64)

	resp, err := cfg.CommissionClient.GetUserSettlements(ctx, &commissionpb.GetUserSettlementsRequest{
		UserId: uID,
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

	sID, _ := strconv.ParseInt(settlementID, 10, 64)

	resp, err := cfg.CommissionClient.GetSettlementById(ctx, &commissionpb.GetSettlementByIdRequest{
		Id: sID,
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

	uID, _ := strconv.ParseInt(userID, 10, 64)

	resp, err := cfg.CommissionClient.GetUserTotalPending(ctx, &commissionpb.GetUserTotalPendingRequest{
		UserId: uID,
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

	uID, _ := strconv.ParseInt(userID, 10, 64)

	resp, err := cfg.CommissionClient.GetUserTotalSettled(ctx, &commissionpb.GetUserTotalSettledRequest{
		UserId: uID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"totalSettled": resp.TotalSettled,
	})
}
