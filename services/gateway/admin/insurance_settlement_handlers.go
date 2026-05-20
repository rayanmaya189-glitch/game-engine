package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/game_engine/gateway/common/handler"

	commissionpb "github.com/game_engine/common-service/proto/gen/go/commission/v1"
)

// Insurance Claims Handlers
func (cfg *RouterConfig) ListInsuranceClaims(ctx context.Context, c *app.RequestContext) {
	if cfg.CommissionClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"claims": []interface{}{},
			"total":  0,
		})
		return
	}

	status := c.Query("status")
	if status == "" {
		status = "PENDING"
	}

	resp, err := cfg.CommissionClient.GetInsuranceClaimsByStatus(ctx, &commissionpb.GetInsuranceClaimsByStatusRequest{
		Status: status,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to list insurance claims: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
		"total":  len(resp.Claims),
	})
}

func (cfg *RouterConfig) GetInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	id, _ := strconv.ParseInt(claimID, 10, 64)

	if cfg.CommissionClient == nil {
		handler.SendSuccess(c, map[string]interface{}{
			"id":          claimID,
			"claimType":   "GAME_LOSS",
			"claimAmount": "0.00",
			"status":      "UNKNOWN",
		})
		return
	}

	// Try pending first
	resp, err := cfg.CommissionClient.GetInsuranceClaimsByStatus(ctx, &commissionpb.GetInsuranceClaimsByStatusRequest{
		Status: "PENDING",
	})
	if err == nil {
		for _, claim := range resp.Claims {
			if claim.Id == id {
				handler.SendSuccess(c, claim)
				return
			}
		}
	}

	// Try other statuses
	for _, status := range []string{"APPROVED", "PAID", "REJECTED"} {
		r, e := cfg.CommissionClient.GetInsuranceClaimsByStatus(ctx, &commissionpb.GetInsuranceClaimsByStatusRequest{
			Status: status,
		})
		if e == nil {
			for _, claim := range r.Claims {
				if claim.Id == id {
					handler.SendSuccess(c, claim)
					return
				}
			}
		}
	}

	handler.SendJSONError(c, 404, handler.ErrCodeNotFound, "insurance claim not found")
}

func (cfg *RouterConfig) ApproveInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	id, _ := strconv.ParseInt(claimID, 10, 64)

	if cfg.CommissionClient == nil {
		handler.SendJSONError(c, 503, handler.ErrCodeServiceUnavailable, "commission service unavailable")
		return
	}

	adminNote := c.Query("admin_note")
	reviewedByStr := c.Query("reviewed_by")
	reviewedBy, _ := strconv.ParseInt(reviewedByStr, 10, 64)

	resp, err := cfg.CommissionClient.ApproveInsuranceClaim(ctx, &commissionpb.ApproveInsuranceClaimRequest{
		Id:         id,
		ReviewedBy: reviewedBy,
		AdminNote:  adminNote,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to approve insurance claim: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "APPROVED",
		"message": "Insurance claim approved",
		"claim":   resp.Claim,
	})
}

func (cfg *RouterConfig) RejectInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	id, _ := strconv.ParseInt(claimID, 10, 64)

	if cfg.CommissionClient == nil {
		handler.SendJSONError(c, 503, handler.ErrCodeServiceUnavailable, "commission service unavailable")
		return
	}

	adminNote := c.Query("admin_note")
	reviewedByStr := c.Query("reviewed_by")
	reviewedBy, _ := strconv.ParseInt(reviewedByStr, 10, 64)

	resp, err := cfg.CommissionClient.RejectInsuranceClaim(ctx, &commissionpb.RejectInsuranceClaimRequest{
		Id:         id,
		ReviewedBy: reviewedBy,
		AdminNote:  adminNote,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to reject insurance claim: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "REJECTED",
		"message": "Insurance claim rejected",
		"claim":   resp.Claim,
	})
}

func (cfg *RouterConfig) PayInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	id, _ := strconv.ParseInt(claimID, 10, 64)

	if cfg.CommissionClient == nil {
		handler.SendJSONError(c, 503, handler.ErrCodeServiceUnavailable, "commission service unavailable")
		return
	}

	resp, err := cfg.CommissionClient.PayInsuranceClaim(ctx, &commissionpb.PayInsuranceClaimRequest{
		Id: id,
	})
	if err != nil {
		handler.SendJSONError(c, 500, handler.ErrCodeServiceUnavailable, fmt.Sprintf("failed to pay insurance claim: %v", err))
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "PAID",
		"message": "Insurance claim paid",
		"claim":   resp.Claim,
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
		"total":       len(resp.Settlements),
	})
}

func (cfg *RouterConfig) GetSettlement(ctx context.Context, c *app.RequestContext) {
	settlementID := c.Param("id")
	id, _ := strconv.ParseInt(settlementID, 10, 64)

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
		Id: id,
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

	pendingResp, pendingErr := cfg.CommissionClient.GetUserTotalPending(ctx, &commissionpb.GetUserTotalPendingRequest{})
	settledResp, settledErr := cfg.CommissionClient.GetUserTotalSettled(ctx, &commissionpb.GetUserTotalSettledRequest{})

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
	}

	if settledErr == nil && settledResp != nil {
		statistics["totalPaid"] = settledResp.TotalSettled
	}

	handler.SendSuccess(c, statistics)
}
