package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/game_engine/gateway/common/handler"
)

// Merchants Handlers
func (cfg *RouterConfig) ListMerchants(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"merchants": []interface{}{},
		"total":     0,
	})
}

func (cfg *RouterConfig) GetMerchant(ctx context.Context, c *app.RequestContext) {
	merchantID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":             merchantID,
		"name":           "Merchant Name",
		"email":          "merchant@example.com",
		"commissionRate": 10,
		"status":         "active",
	})
}

func (cfg *RouterConfig) CreateMerchant(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"id":      "new_merchant_id",
		"message": "Merchant created successfully",
	})
}

func (cfg *RouterConfig) UpdateMerchant(ctx context.Context, c *app.RequestContext) {
	merchantID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      merchantID,
		"message": "Merchant updated successfully",
	})
}

func (cfg *RouterConfig) UpdateMerchantStatus(ctx context.Context, c *app.RequestContext) {
	merchantID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      merchantID,
		"status":  "updated",
		"message": "Merchant status updated",
	})
}

func (cfg *RouterConfig) DeleteMerchant(ctx context.Context, c *app.RequestContext) {
	merchantID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      merchantID,
		"message": "Merchant deleted successfully",
	})
}

// Agents Handlers
func (cfg *RouterConfig) ListAgents(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"agents": []interface{}{},
		"total":  0,
	})
}

func (cfg *RouterConfig) GetAgent(ctx context.Context, c *app.RequestContext) {
	agentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":     agentID,
		"name":   "Agent Name",
		"email":  "agent@example.com",
		"tier":   "Gold",
		"status": "active",
	})
}

func (cfg *RouterConfig) CreateAgent(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"id":      "new_agent_id",
		"message": "Agent created successfully",
	})
}

func (cfg *RouterConfig) UpdateAgent(ctx context.Context, c *app.RequestContext) {
	agentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      agentID,
		"message": "Agent updated successfully",
	})
}

func (cfg *RouterConfig) UpdateAgentStatus(ctx context.Context, c *app.RequestContext) {
	agentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      agentID,
		"status":  "updated",
		"message": "Agent status updated",
	})
}

func (cfg *RouterConfig) DeleteAgent(ctx context.Context, c *app.RequestContext) {
	agentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      agentID,
		"message": "Agent deleted successfully",
	})
}
