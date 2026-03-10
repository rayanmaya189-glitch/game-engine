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
	CommissionClient      *client.CommissionClient
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

		// Claims routes
		claims := protected.Group("/claims")
		{
			// Commission claims
			claims.POST("/commission", cfg.SubmitCommissionClaim)
			claims.GET("/commission", cfg.GetUserCommissionClaims)
			claims.GET("/commission/status/:status", cfg.GetCommissionClaimsByStatus)
			claims.POST("/commission/:id/claim", cfg.ClaimRebet)

			// Rebet claims
			claims.POST("/rebet", cfg.CreateRebetClaim)
			claims.GET("/rebet", cfg.GetUserRebetClaims)
			claims.GET("/rebet/claimable", cfg.GetClaimableRebets)
			claims.POST("/rebet/:id/claim", cfg.ClaimRebet)

			// Insurance claims
			claims.POST("/insurance", cfg.SubmitInsuranceClaim)
			claims.GET("/insurance", cfg.GetUserInsuranceClaims)

			// Settlements
			claims.GET("/settlements", cfg.GetUserSettlements)
			claims.GET("/settlements/:id", cfg.GetSettlementById)

			// Totals
			claims.GET("/total-pending", cfg.GetUserTotalPending)
			claims.GET("/total-settled", cfg.GetUserTotalSettled)
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

// SubmitCommissionClaim handles submitting a commission claim
func (cfg *RouterConfig) SubmitCommissionClaim(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	var req struct {
		AffiliateID  string `json:"affiliateId"`
		CommissionID string `json:"commissionId"`
		Amount       string `json:"amount"`
		ClaimReason  string `json:"claimReason"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":  "Commission claim submitted",
		"claim_id": "claim_123",
		"status":   "PENDING",
	})
}

// GetUserCommissionClaims handles getting user's commission claims
func (cfg *RouterConfig) GetUserCommissionClaims(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": []interface{}{},
	})
}

// GetCommissionClaimsByStatus handles getting commission claims by status
func (cfg *RouterConfig) GetCommissionClaimsByStatus(ctx context.Context, c *app.RequestContext) {
	status := c.Param("status")
	handler.SendSuccess(c, map[string]interface{}{
		"claims": []interface{}{},
		"status": status,
	})
}

// CreateRebetClaim handles creating a rebet claim
func (cfg *RouterConfig) CreateRebetClaim(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	var req struct {
		BonusID          string `json:"bonusId"`
		BonusCode        string `json:"bonusCode"`
		BonusAmount      string `json:"bonusAmount"`
		RebetRequirement string `json:"rebetRequirement"`
		GameID           string `json:"gameId"`
		BetID            string `json:"betId"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":           "Rebet claim created",
		"rebet_id":          "rebet_123",
		"status":            "IN_PROGRESS",
		"rebet_requirement": req.RebetRequirement,
		"current_rebet":     "0",
	})
}

// GetUserRebetClaims handles getting user's rebet claims
func (cfg *RouterConfig) GetUserRebetClaims(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": []interface{}{},
	})
}

// GetClaimableRebets handles getting claimable rebet bonuses
func (cfg *RouterConfig) GetClaimableRebets(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": []interface{}{},
	})
}

// ClaimRebet handles claiming a rebet bonus
func (cfg *RouterConfig) ClaimRebet(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	rebetID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Rebet bonus claimed",
		"rebet_id":       rebetID,
		"amount":         "100.00",
		"transaction_id": "txn_rebet_123",
	})
}

// SubmitInsuranceClaim handles submitting an insurance claim
func (cfg *RouterConfig) SubmitInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	var req struct {
		GameID            string `json:"gameId"`
		BetID             string `json:"betId"`
		InsurancePolicyID string `json:"insurancePolicyId"`
		ClaimType         string `json:"claimType"`
		InsuredAmount     string `json:"insuredAmount"`
		LossAmount        string `json:"lossAmount"`
		ClaimReason       string `json:"claimReason"`
		EvidenceDetails   string `json:"evidenceDetails"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":  "Insurance claim submitted",
		"claim_id": "insurance_123",
		"status":   "PENDING",
	})
}

// GetUserInsuranceClaims handles getting user's insurance claims
func (cfg *RouterConfig) GetUserInsuranceClaims(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": []interface{}{},
	})
}

// GetUserSettlements handles getting user's settlements
func (cfg *RouterConfig) GetUserSettlements(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"settlements": []interface{}{},
	})
}

// GetSettlementById handles getting settlement by ID
func (cfg *RouterConfig) GetSettlementById(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	settlementID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"settlement_id": settlementID,
		"amount":        "100.00",
		"type":          "COMMISSION",
		"status":        "COMPLETED",
	})
}

// GetUserTotalPending handles getting user's total pending claims
func (cfg *RouterConfig) GetUserTotalPending(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"totalPending": "500.00",
	})
}

// GetUserTotalSettled handles getting user's total settled claims
func (cfg *RouterConfig) GetUserTotalSettled(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"totalSettled": "1500.00",
	})
}
