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
	AuthClient            *client.AuthClient
	UserClient            *client.UserClient
	WalletClient          *client.WalletClient
	GameClient            *client.GameClient
	AllowedIPs            []string
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

	// Admin routes with JWT + MFA + Role check
	admin := r.Group("/api/v1/admin")
	admin.Use(cfg.AuthMiddleware.JWTValidation())
	admin.Use(cfg.AuthMiddleware.MFACheck())
	admin.Use(cfg.AuthMiddleware.RoleCheck("admin"))

	// Admin IP whitelist
	if len(cfg.AllowedIPs) > 0 {
		admin.Use(cfg.AuthMiddleware.IPWhitelistCheck(cfg.AllowedIPs))
	}

	// Player management
	admin.GET("/players", cfg.ListPlayers)
	admin.GET("/players/:id", cfg.GetPlayer)
	admin.PUT("/players/:id/status", cfg.UpdatePlayerStatus)
	admin.GET("/players/:id/stats", cfg.GetPlayerStats)

	// KYC management
	admin.GET("/kyc", cfg.GetKYCList)
	admin.PUT("/kyc/:id/approve", cfg.ApproveKYC)
	admin.PUT("/kyc/:id/reject", cfg.RejectKYC)

	// Game management
	admin.GET("/games", cfg.ListAdminGames)
	admin.POST("/games", cfg.CreateGame)
	admin.PUT("/games/:id", cfg.UpdateGame)

	// Wallet management
	admin.GET("/wallet/transactions", cfg.GetAllTransactions)
	admin.POST("/wallet/adjust", cfg.AdjustBalance)

	// Reports (future)
	admin.GET("/reports/*path", cfg.ReportsHandler)

	return r
}

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

func (cfg *RouterConfig) ListAdminGames(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"games": []interface{}{},
	})
}

func (cfg *RouterConfig) CreateGame(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"game_id": "new_game_id",
		"message": "Game created successfully",
	})
}

func (cfg *RouterConfig) UpdateGame(ctx context.Context, c *app.RequestContext) {
	gameID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"game_id": gameID,
		"message": "Game updated successfully",
	})
}

func (cfg *RouterConfig) GetAllTransactions(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"transactions": []interface{}{},
	})
}

func (cfg *RouterConfig) AdjustBalance(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Balance adjusted successfully",
		"transaction_id": "adj_123",
	})
}

func (cfg *RouterConfig) ReportsHandler(ctx context.Context, c *app.RequestContext) {
	path := string(c.Request.URI().Path())
	handler.SendSuccess(c, map[string]interface{}{
		"path":    path,
		"message": "Reports endpoint",
	})
}
