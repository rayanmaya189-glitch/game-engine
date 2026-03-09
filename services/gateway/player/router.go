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
}

func NewRouter(cfg *RouterConfig) *router.Router {
	r := router.New()

	// Global middleware (applied to all routes)
	r.Use(cfg.LoggerMiddleware.RequestID())
	r.Use(cfg.LoggerMiddleware.StructuredLogger())
	r.Use(cfg.LoggerMiddleware.PanicRecovery())
	r.Use(cfg.CORSMiddleware.CORS())

	// Health check endpoints (no auth required)
	r.GET("/health", handler.HandleHealthCheck)
	r.GET("/ready", handler.HandleReadinessCheck)

	// Apply rate limiting to all routes
	r.Use(cfg.RateLimiterMiddleware.RateLimiter())

	// Auth routes (no JWT required for register/login/refresh)
	authGroup := r.Group("/api/v1/auth")
	{
		authGroup.POST("/register", cfg.Register)
		authGroup.POST("/login", cfg.Login)
		authGroup.POST("/refresh", cfg.RefreshToken)
		authGroup.POST("/logout", cfg.Logout)
	}

	// Protected routes (JWT required)
	protected := r.Group("/api/v1")
	protected.Use(cfg.AuthMiddleware.JWTValidation())
	{
		// User routes
		users := protected.Group("/users")
		{
			users.GET("/profile", cfg.GetProfile)
			users.PUT("/profile", cfg.UpdateProfile)
		}

		// Wallet routes
		wallet := protected.Group("/wallet")
		{
			wallet.GET("/balance", cfg.GetBalance)
			wallet.GET("/transactions", cfg.GetTransactions)
			wallet.POST("/deposit", cfg.Deposit)
			wallet.POST("/withdraw", cfg.Withdraw)
		}

		// Game routes
		games := protected.Group("/games")
		{
			games.GET("", cfg.ListGames)
			games.GET("/:id", cfg.GetGame)
			games.GET("/:id/play", cfg.PlayGame)
			games.GET("/categories", cfg.GetCategories)
			games.GET("/featured", cfg.GetFeaturedGames)
			games.GET("/popular", cfg.GetPopularGames)
		}
	}

	return r
}

// Register handles user registration
func (cfg *RouterConfig) Register(ctx context.Context, c *app.RequestContext) {
	if cfg.AuthClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Auth service unavailable", nil)
		return
	}

	// Parse request body
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	// Call auth service
	// resp, err := cfg.AuthClient.Register(ctx, &authpb.RegisterRequest{...})

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Registration successful",
	})
}

// Login handles user login
func (cfg *RouterConfig) Login(ctx context.Context, c *app.RequestContext) {
	if cfg.AuthClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Auth service unavailable", nil)
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	// Call auth service
	// resp, err := cfg.AuthClient.Login(ctx, &authpb.LoginRequest{...})

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Login successful",
	})
}

// RefreshToken handles token refresh
func (cfg *RouterConfig) RefreshToken(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"message": "Token refreshed",
	})
}

// Logout handles user logout
func (cfg *RouterConfig) Logout(ctx context.Context, c *app.RequestContext) {
	// Get token from context and blacklist it
	userID := c.GetString("user_id")
	if userID != "" && cfg.AuthMiddleware != nil {
		// Blacklist the token
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Logout successful",
	})
}

// GetProfile handles getting user profile
func (cfg *RouterConfig) GetProfile(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"user_id":  userID,
		"username": c.GetString("username"),
		"email":    "user@example.com",
		"status":   "active",
	})
}

// UpdateProfile handles updating user profile
func (cfg *RouterConfig) UpdateProfile(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Profile updated",
	})
}

// GetBalance handles getting wallet balance
func (cfg *RouterConfig) GetBalance(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"balance": map[string]interface{}{
			"main":     "1000.00",
			"bonus":    "100.00",
			"currency": "USD",
		},
	})
}

// GetTransactions handles getting wallet transactions
func (cfg *RouterConfig) GetTransactions(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"transactions": []interface{}{},
	})
}

// Deposit handles deposit request
func (cfg *RouterConfig) Deposit(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Deposit successful",
		"transaction_id": "txn_123",
	})
}

// Withdraw handles withdraw request
func (cfg *RouterConfig) Withdraw(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Withdrawal request submitted",
		"transaction_id": "txn_456",
	})
}

// ListGames handles listing games
func (cfg *RouterConfig) ListGames(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"games": []interface{}{},
	})
}

// GetGame handles getting game details
func (cfg *RouterConfig) GetGame(ctx context.Context, c *app.RequestContext) {
	gameID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"game_id":  gameID,
		"name":     "Sample Game",
		"provider": "Sample Provider",
		"category": "slot",
		"status":   "active",
	})
}

// PlayGame handles game play URL generation
func (cfg *RouterConfig) PlayGame(ctx context.Context, c *app.RequestContext) {
	gameID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"game_id": gameID,
		"url":     "https://game.example.com/play/" + gameID,
	})
}

// GetCategories handles getting game categories
func (cfg *RouterConfig) GetCategories(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"categories": []string{"slot", "live_casino", "table_games", "arcade"},
	})
}

// GetFeaturedGames handles getting featured games
func (cfg *RouterConfig) GetFeaturedGames(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"games": []interface{}{},
	})
}

// GetPopularGames handles getting popular games
func (cfg *RouterConfig) GetPopularGames(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"games": []interface{}{},
	})
}
