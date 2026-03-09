package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/router"

	"common/client"
	"common/handler"
	"common/middleware"
)

type RouterConfig struct {
	AuthMiddleware        *middleware.AuthMiddleware
	LoggerMiddleware      *middleware.LoggerMiddleware
	RateLimiterMiddleware *middleware.RateLimiterMiddleware
	CORSMiddleware        *middleware.CORSMiddleware
	ValidatorMiddleware   *middleware.ValidatorMiddleware
	ErrorHandler          *handler.ErrorHandler
	UserClient            *client.UserClient
	WalletClient          *client.WalletClient
}

func NewRouter(cfg *RouterConfig) *router.Router {
	r := router.New()

	r.Use(cfg.LoggerMiddleware.RequestID())
	r.Use(cfg.LoggerMiddleware.StructuredLogger())
	r.Use(cfg.LoggerMiddleware.PanicRecovery())
	r.Use(cfg.CORSMiddleware.CORS())

	r.GET("/health", handler.HandleHealthCheck)
	r.GET("/ready", handler.HandleReadinessCheck)

	r.Use(cfg.RateLimiterMiddleware.RateLimiter())

	// Merchant routes with API Key authentication
	merchant := r.Group("/api/v1/merchant")
	merchant.Use(cfg.AuthMiddleware.APIKeyValidation())

	// Player data
	merchant.GET("/players", cfg.ListMerchantPlayers)
	merchant.GET("/players/:id", cfg.GetMerchantPlayer)

	// Reports
	merchant.GET("/reports/revenue", cfg.GetRevenueReports)
	merchant.GET("/reports/players", cfg.GetPlayerReports)

	// Config
	merchant.GET("/config", cfg.GetMerchantConfig)
	merchant.PUT("/config", cfg.UpdateMerchantConfig)

	// Webhooks
	merchant.POST("/webhooks/register", cfg.RegisterWebhook)
	merchant.GET("/webhooks", cfg.ListWebhooks)

	return r
}

func (cfg *RouterConfig) ListMerchantPlayers(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	handler.SendSuccess(c, map[string]interface{}{
		"players":     []interface{}{},
		"total":       0,
		"merchant_id": merchantID,
	})
}

func (cfg *RouterConfig) GetMerchantPlayer(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	playerID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"player_id":   playerID,
		"merchant_id": merchantID,
		"username":    "player1",
	})
}

func (cfg *RouterConfig) GetRevenueReports(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id": merchantID,
		"revenue":     "10000.00",
		"currency":    "USD",
	})
}

func (cfg *RouterConfig) GetPlayerReports(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id":    merchantID,
		"total_players":  100,
		"active_players": 50,
	})
}

func (cfg *RouterConfig) GetMerchantConfig(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id":     merchantID,
		"commission_rate": 10.0,
		"status":          "active",
	})
}

func (cfg *RouterConfig) UpdateMerchantConfig(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id": merchantID,
		"message":     "Configuration updated",
	})
}

func (cfg *RouterConfig) RegisterWebhook(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id": merchantID,
		"webhook_id":  "wh_123",
		"message":     "Webhook registered",
	})
}

func (cfg *RouterConfig) ListWebhooks(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id": merchantID,
		"webhooks":    []interface{}{},
	})
}
