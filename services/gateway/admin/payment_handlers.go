package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/game_engine/gateway/common/handler"
)

// Payments Handlers
func (cfg *RouterConfig) ListPayments(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"payments": []interface{}{},
		"total":    0,
	})
}

func (cfg *RouterConfig) GetPayment(ctx context.Context, c *app.RequestContext) {
	paymentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":     paymentID,
		"userId": "user_123",
		"amount": 100,
		"method": "Bank Transfer",
		"type":   "deposit",
		"status": "pending",
	})
}

func (cfg *RouterConfig) ApprovePayment(ctx context.Context, c *app.RequestContext) {
	paymentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      paymentID,
		"status":  "approved",
		"message": "Payment approved",
	})
}

func (cfg *RouterConfig) RejectPayment(ctx context.Context, c *app.RequestContext) {
	paymentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      paymentID,
		"status":  "rejected",
		"message": "Payment rejected",
	})
}

func (cfg *RouterConfig) ProcessPayment(ctx context.Context, c *app.RequestContext) {
	paymentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":            paymentID,
		"status":        "completed",
		"message":       "Payment processed",
		"transactionId": "txn_" + paymentID,
	})
}
