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

	// Commission routes
	agent.GET("/commissions", cfg.GetCommissions)
	agent.GET("/commissions/pending", cfg.GetPendingCommissions)

	// Affiliate routes
	affiliate := r.Group("/api/v1/affiliate/tracking")
	affiliate.POST("/click", cfg.TrackClick)
	affiliate.GET("/:code", cfg.RedirectToRegistration)

	// Affiliate reports
	affiliateReports := r.Group("/api/v1/affiliate/reports")
	affiliateReports.GET("/performance", cfg.GetPerformanceReports)

	return r
}

func (cfg *RouterConfig) ListAgentPlayers(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	handler.SendSuccess(c, map[string]interface{}{
		"players":  []interface{}{},
		"total":    0,
		"agent_id": agentID,
	})
}

func (cfg *RouterConfig) GetAgentPlayer(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	playerID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"player_id": playerID,
		"agent_id":  agentID,
		"username":  "player1",
	})
}

func (cfg *RouterConfig) GetCommissions(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	handler.SendSuccess(c, map[string]interface{}{
		"agent_id":    agentID,
		"commissions": []interface{}{},
		"total":       "100.00",
	})
}

func (cfg *RouterConfig) GetPendingCommissions(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	handler.SendSuccess(c, map[string]interface{}{
		"agent_id": agentID,
		"pending":  "50.00",
	})
}

func (cfg *RouterConfig) TrackClick(ctx context.Context, c *app.RequestContext) {
	code := c.Query("code")
	handler.SendSuccess(c, map[string]interface{}{
		"code":    code,
		"message": "Click tracked",
	})
}

func (cfg *RouterConfig) RedirectToRegistration(ctx context.Context, c *app.RequestContext) {
	code := c.Param("code")
	// Redirect to registration page with affiliate code
	c.Redirect(302, []byte("https://example.com/register?ref="+code))
}

func (cfg *RouterConfig) GetPerformanceReports(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	handler.SendSuccess(c, map[string]interface{}{
		"agent_id":      agentID,
		"total_clicks":  1000,
		"total_signups": 50,
		"conversion":    5.0,
	})
}
