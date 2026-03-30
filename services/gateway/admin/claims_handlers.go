package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"common/handler"
)

// Commission Claims Handlers
func (cfg *RouterConfig) ListCommissionClaims(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"claims":     []interface{}{},
		"total":      0,
		"page":       1,
		"totalPages": 0,
	})
}

func (cfg *RouterConfig) GetCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":          claimID,
		"claimType":   "COMMISSION",
		"amount":      "100.00",
		"status":      "PENDING",
		"claimReason": "Commission claim request",
	})
}

func (cfg *RouterConfig) ApproveCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "APPROVED",
		"message": "Commission claim approved",
	})
}

func (cfg *RouterConfig) RejectCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "REJECTED",
		"message": "Commission claim rejected",
	})
}

func (cfg *RouterConfig) PayCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":            claimID,
		"status":        "PAID",
		"message":       "Commission claim paid",
		"transactionId": "txn_" + claimID,
	})
}

// Rebet Claims Handlers
func (cfg *RouterConfig) ListRebetClaims(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"claims": []interface{}{},
		"total":  0,
	})
}

func (cfg *RouterConfig) GetRebetClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":        claimID,
		"bonusCode": "BONUS123",
		"amount":    "50.00",
		"status":    "CLAIMABLE",
	})
}

func (cfg *RouterConfig) ApproveRebetClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "APPROVED",
		"message": "Rebet claim approved",
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

// Insurance Claims Handlers
func (cfg *RouterConfig) ListInsuranceClaims(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"claims": []interface{}{},
		"total":  0,
	})
}

func (cfg *RouterConfig) GetInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":          claimID,
		"claimType":   "GAME_LOSS",
		"claimAmount": "200.00",
		"status":      "PENDING",
	})
}

func (cfg *RouterConfig) ApproveInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "APPROVED",
		"message": "Insurance claim approved",
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
		"id":            claimID,
		"status":        "PAID",
		"message":       "Insurance claim paid",
		"transactionId": "ins_txn_" + claimID,
	})
}

// Settlements Handlers
func (cfg *RouterConfig) ListSettlements(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"settlements": []interface{}{},
		"total":       0,
	})
}

func (cfg *RouterConfig) GetSettlement(ctx context.Context, c *app.RequestContext) {
	settlementID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":             settlementID,
		"settlementType": "COMMISSION",
		"amount":         "100.00",
		"status":         "COMPLETED",
	})
}

func (cfg *RouterConfig) GetClaimStatistics(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"totalPending":    15,
		"totalApproved":   25,
		"totalPaid":       100,
		"totalRejected":   5,
		"totalInProgress": 10,
		"totalAmount":     "15000.00",
	})
}
