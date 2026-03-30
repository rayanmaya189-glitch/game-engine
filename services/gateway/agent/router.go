package main

import (
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
	AgentClient           *client.AgentClient
	AffiliateClient       *client.AffiliateClient
	UserClient            *client.UserClient
	CommissionClient      *client.CommissionClient
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

	// Agent routes with API Key + JWT authentication
	agent := r.Group("/api/v1/agent")
	agent.Use(cfg.AuthMiddleware.JWTValidation())
	agent.Use(cfg.AuthMiddleware.APIKeyValidation())

	// Downline players
	agent.GET("/players", cfg.ListAgentPlayers)
	agent.GET("/players/:id", cfg.GetAgentPlayer)
	agent.PUT("/players/:id/limit", cfg.UpdatePlayerLimit)

	// Commission routes
	agent.GET("/commissions", cfg.GetCommissions)
	agent.GET("/commissions/pending", cfg.GetPendingCommissions)
	agent.GET("/commissions/history", cfg.GetCommissionHistory)
	agent.POST("/commissions/claim", cfg.ClaimCommission)

	// Dashboard
	agent.GET("/dashboard", cfg.GetDashboard)

	// Affiliate routes
	affiliate := r.Group("/api/v1/affiliate/tracking")
	affiliate.POST("/click", cfg.TrackClick)
	affiliate.GET("/:code", cfg.RedirectToRegistration)

	// Affiliate reports
	affiliateReports := r.Group("/api/v1/affiliate/reports")
	affiliateReports.GET("/performance", cfg.GetPerformanceReports)
	affiliateReports.GET("/clicks", cfg.GetClickReports)
	affiliateReports.GET("/conversions", cfg.GetConversionReports)

	return r
}
