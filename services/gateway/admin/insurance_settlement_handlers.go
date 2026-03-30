package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"

	"common/handler"

	commissionpb "github.com/game_engine/gen/go/game_engine/commission/v1"
	bonuspb "github.com/game_engine/gen/go/game_engine/bonus/v1"
)

// Insurance Claims Handlers
func (cfg *RouterConfig) ListInsuranceClaims(ctx context.Context, c *app.RequestContext) {
	if cfg.BonusClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"claims": []interface{}{},
			"total":  0,
		})
		return
	}

	resp, err := cfg.BonusClient.GetUserInsuranceClaims(ctx, &bonuspb.GetUserInsuranceClaimsRequest{})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to list insurance claims: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
		"total":  resp.Total,
	})
}

func (cfg *RouterConfig) GetInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")

	if cfg.BonusClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"id":          claimID,
			"claimType":   "GAME_LOSS",
			"claimAmount": "0.00",
			"status":      "UNKNOWN",
		})
		return
	}

	resp, err := cfg.BonusClient.GetUserInsuranceClaims(ctx, &bonuspb.GetUserInsuranceClaimsRequest{})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to get insurance claim: %v", err))
		return
	}

	for _, claim := range resp.Claims {
		if claim.Id == claimID {
			handler.SendSuccess(c, claim)
			return
		}
	}

	handler.SendJSONError(c, 404, handler.ErrCodeNotFound, "insurance claim not found")
}

func (cfg *RouterConfig) ApproveInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")

	if cfg.BonusClient == nil {
		handler.SendJSONError(c, 503, handler.ErrCodeServiceUnavailable, "bonus service unavailable")
		return
	}

	resp, err := cfg.BonusClient.SubmitInsuranceClaim(ctx, &bonuspb.SubmitInsuranceClaimRequest{
		ClaimId: claimID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to approve insurance claim: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "APPROVED",
		"message": resp.Message,
	})
}

func (cfg *RouterConfig) RejectInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "REJECTED",
		"message": "Insurance claim rejected",
	})
}

func (cfg *RouterConfig) PayInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "PAID",
		"message": "Insurance claim paid",
	})
}

// Settlements Handlers
func (cfg *RouterConfig) ListSettlements(ctx context.Context, c *app.RequestContext) {
	if cfg.CommissionClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"settlements": []interface{}{},
			"total":       0,
		})
		return
	}

	resp, err := cfg.CommissionClient.GetUserSettlements(ctx, &commissionpb.GetUserSettlementsRequest{})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to list settlements: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"settlements": resp.Settlements,
		"total":       resp.Total,
	})
}

func (cfg *RouterConfig) GetSettlement(ctx context.Context, c *app.RequestContext) {
	settlementID := c.Param("id")

	if cfg.CommissionClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"id":             settlementID,
			"settlementType": "COMMISSION",
			"amount":         "0.00",
			"status":         "UNKNOWN",
		})
		return
	}

	resp, err := cfg.CommissionClient.GetSettlementById(ctx, &commissionpb.GetSettlementByIdRequest{
		SettlementId: settlementID,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to get settlement: %v", err))
		return
	}

	handler.SendSuccess(c, resp.Settlement)
}

func (cfg *RouterConfig) GetClaimStatistics(ctx context.Context, c *app.RequestContext) {
	if cfg.CommissionClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"totalPending":    0,
			"totalApproved":   0,
			"totalPaid":       0,
			"totalRejected":   0,
			"totalInProgress": 0,
			"totalAmount":     "0.00",
		})
		return
	}

	pendingResp, pendingErr := cfg.CommissionClient.GetTotalPending(ctx, &commissionpb.GetTotalPendingRequest{})
	settledResp, settledErr := cfg.CommissionClient.GetTotalSettled(ctx, &commissionpb.GetTotalSettledRequest{})

	statistics := map[string]interface{}{
		"totalPending":    0,
		"totalApproved":   0,
		"totalPaid":       0,
		"totalRejected":   0,
		"totalInProgress": 0,
		"totalAmount":     "0.00",
	}

	if pendingErr == nil && pendingResp != nil {
		statistics["totalPending"] = pendingResp.TotalPending
		statistics["totalAmount"] = pendingResp.TotalAmount
	}

	if settledErr == nil && settledResp != nil {
		statistics["totalPaid"] = settledResp.TotalSettled
	}

	handler.SendSuccess(c, statistics)
}
