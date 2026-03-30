package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"common/handler"
)

func (cfg *RouterConfig) ListPlayers(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"players": []interface{}{},
		"total":   0,
	})
}

func (cfg *RouterConfig) GetPlayer(ctx context.Context, c *app.RequestContext) {
	playerID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"player_id": playerID,
		"username":  "player1",
		"email":     "player@example.com",
		"status":    "active",
	})
}

func (cfg *RouterConfig) UpdatePlayerStatus(ctx context.Context, c *app.RequestContext) {
	playerID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"player_id": playerID,
		"message":   "Player status updated",
	})
}

func (cfg *RouterConfig) GetPlayerStats(ctx context.Context, c *app.RequestContext) {
	playerID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"player_id":         playerID,
		"total_deposits":    "1000.00",
		"total_withdrawals": "500.00",
		"total_bets":        "2000.00",
	})
}

func (cfg *RouterConfig) GetKYCList(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"kyc_requests": []interface{}{},
	})
}

func (cfg *RouterConfig) ApproveKYC(ctx context.Context, c *app.RequestContext) {
	kycID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"kyc_id":  kycID,
		"status":  "approved",
		"message": "KYC approved successfully",
	})
}

func (cfg *RouterConfig) RejectKYC(ctx context.Context, c *app.RequestContext) {
	kycID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"kyc_id":  kycID,
		"status":  "rejected",
		"message": "KYC rejected",
	})
}
