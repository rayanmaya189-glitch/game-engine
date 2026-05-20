package main

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/game_engine/gateway/common/client"
	"github.com/game_engine/gateway/common/handler"
	"github.com/game_engine/gateway/common/middleware"
)

type RouterConfig struct {
	AuthMiddleware        *middleware.AuthMiddleware
	LoggerMiddleware      *middleware.LoggerMiddleware
	RateLimiterMiddleware *middleware.RateLimiterMiddleware
	CORSMiddleware        *middleware.CORSMiddleware
	ValidatorMiddleware   *middleware.ValidatorMiddleware
	ErrorHandler          *handler.ErrorHandler
	MerchantClient        *client.MerchantClient
	UserClient            *client.UserClient
	WalletClient          *client.WalletClient
	CommissionClient      *client.CommissionClient
}

func NewRouter(cfg *RouterConfig) *route.Engine {
	r := route.New()

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
	merchant.GET("/reports/games", cfg.GetGameReports)

	// Config
	merchant.GET("/config", cfg.GetMerchantConfig)
	merchant.PUT("/config", cfg.UpdateMerchantConfig)

	// Webhooks
	merchant.POST("/webhooks/register", cfg.RegisterWebhook)
	merchant.GET("/webhooks", cfg.ListWebhooks)
	merchant.DELETE("/webhooks/:id", cfg.DeleteWebhook)

	// Sub-agents
	merchant.GET("/agents", cfg.ListSubAgents)
	merchant.GET("/agents/:id", cfg.GetSubAgent)
	merchant.POST("/agents", cfg.CreateSubAgent)
	merchant.PUT("/agents/:id", cfg.UpdateSubAgent)
	merchant.PUT("/agents/:id/status", cfg.UpdateSubAgentStatus)

	return r
}
