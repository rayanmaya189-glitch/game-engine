package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/router"

	authpb "github.com/game-engine/gen/go/game-engine/auth/v1"
	bonuspb "github.com/game-engine/gen/go/game-engine/bonus/v1"
	commissionpb "github.com/game-engine/gen/go/game-engine/commission/v1"
	gamepb "github.com/game-engine/gen/go/game-engine/game/v1"
	jackpotpb "github.com/game-engine/gen/go/game-engine/jackpot/v1"
	tournamentpb "github.com/game-engine/gen/go/game-engine/tournament/v1"
	userpb "github.com/game-engine/gen/go/game-engine/user/v1"
	walletpb "github.com/game-engine/gen/go/game-engine/wallet/v1"

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
	BonusClient           *client.BonusClient
	TournamentClient      *client.TournamentClient
	JackpotClient         *client.JackpotClient
	PaymentClient         *client.PaymentClient
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
			games.GET("/slots", cfg.GetSlotGames)
			games.GET("/cards", cfg.GetCardGames)
			games.GET("/dice", cfg.GetDiceGames)
		}

		// Tournaments
		tournaments := protected.Group("/tournaments")
		{
			tournaments.GET("", cfg.ListTournaments)
			tournaments.GET("/:id", cfg.GetTournament)
			tournaments.POST("/:id/join", cfg.JoinTournament)
			tournaments.GET("/:id/leaderboard", cfg.GetTournamentLeaderboard)
		}

		// Jackpots
		jackpots := protected.Group("/jackpots")
		{
			jackpots.GET("", cfg.ListJackpots)
			jackpots.GET("/:id", cfg.GetJackpot)
			jackpots.GET("/:id/winners", cfg.GetJackpotWinners)
		}

		// Bonuses
		bonuses := protected.Group("/bonuses")
		{
			bonuses.GET("", cfg.ListBonuses)
			bonuses.GET("/:id", cfg.GetBonus)
			bonuses.POST("/:id/claim", cfg.ClaimBonus)
			bonuses.GET("/my-bonuses", cfg.GetMyBonuses)
		}

		// Claims routes
		claims := protected.Group("/claims")
		{
			// Commission claims
			claims.POST("/commission", cfg.SubmitCommissionClaim)
			claims.GET("/commission", cfg.GetUserCommissionClaims)
			claims.GET("/commission/status/:status", cfg.GetCommissionClaimsByStatus)
			claims.POST("/commission/:id/claim", cfg.ClaimCommission)

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

	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"fullName"`
		Phone    string `json:"phone"`
		Currency string `json:"currency"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	// Call auth service via GRPC
	resp, err := cfg.AuthClient.Register(ctx, &authpb.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
		Phone:    req.Phone,
		Currency: req.Currency,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"user_id": resp.UserId,
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

	// Call auth service via GRPC
	resp, err := cfg.AuthClient.Login(ctx, &authpb.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Invalid credentials", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
		"expires_in":    resp.ExpiresIn,
		"token_type":    "Bearer",
	})
}

// RefreshToken handles token refresh
func (cfg *RouterConfig) RefreshToken(ctx context.Context, c *app.RequestContext) {
	if cfg.AuthClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Auth service unavailable", nil)
		return
	}

	var req struct {
		RefreshToken string `json:"refreshToken"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	resp, err := cfg.AuthClient.RefreshToken(ctx, &authpb.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})

	if err != nil {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Invalid refresh token", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
		"expires_in":    resp.ExpiresIn,
		"token_type":    "Bearer",
	})
}

// Logout handles user logout
func (cfg *RouterConfig) Logout(ctx context.Context, c *app.RequestContext) {
	if cfg.AuthClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Auth service unavailable", nil)
		return
	}

	// Get token from header
	token := c.GetHeader("Authorization")
	if token != "" {
		cfg.AuthClient.Logout(ctx, &authpb.LogoutRequest{
			Token: token,
		})
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

	if cfg.UserClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "User service unavailable", nil)
		return
	}

	resp, err := cfg.UserClient.GetProfile(ctx, &userpb.GetProfileRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "User not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"user_id":    resp.User.Id,
		"username":   resp.User.Username,
		"email":      resp.User.Email,
		"full_name":  resp.User.FullName,
		"phone":      resp.User.Phone,
		"status":     resp.User.Status,
		"kyc_status": resp.User.KycStatus,
		"created_at": resp.User.CreatedAt,
	})
}

// UpdateProfile handles updating user profile
func (cfg *RouterConfig) UpdateProfile(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.UserClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "User service unavailable", nil)
		return
	}

	var req struct {
		FullName string `json:"fullName"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	_, err := cfg.UserClient.UpdateProfile(ctx, &userpb.UpdateProfileRequest{
		UserId:   userID,
		FullName: req.FullName,
		Phone:    req.Phone,
		Email:    req.Email,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Profile updated successfully",
	})
}

// GetBalance handles getting wallet balance
func (cfg *RouterConfig) GetBalance(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.WalletClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Wallet service unavailable", nil)
		return
	}

	resp, err := cfg.WalletClient.GetBalance(ctx, &walletpb.GetBalanceRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"balance": map[string]interface{}{
			"main":     resp.MainBalance,
			"bonus":    resp.BonusBalance,
			"currency": resp.Currency,
		},
	})
}

// GetTransactions handles getting wallet transactions
func (cfg *RouterConfig) GetTransactions(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.WalletClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Wallet service unavailable", nil)
		return
	}

	// Parse query parameters
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")
	txnType := c.Query("type")
	status := c.Query("status")

	resp, err := cfg.WalletClient.GetTransactions(ctx, &walletpb.GetTransactionsRequest{
		UserId: userID,
		Page:   page,
		Limit:  limit,
		Type:   txnType,
		Status: status,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	transactions := make([]map[string]interface{}, len(resp.Transactions))
	for i, txn := range resp.Transactions {
		transactions[i] = map[string]interface{}{
			"id":          txn.Id,
			"type":        txn.Type,
			"amount":      txn.Amount,
			"status":      txn.Status,
			"description": txn.Description,
			"created_at":  txn.CreatedAt,
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"transactions": transactions,
		"total":        resp.Total,
		"page":         resp.Page,
		"limit":        resp.Limit,
	})
}

// Deposit handles deposit request
func (cfg *RouterConfig) Deposit(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	var req struct {
		Amount    float64 `json:"amount"`
		Method    string  `json:"method"`
		PaymentID string  `json:"paymentId"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.WalletClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Wallet service unavailable", nil)
		return
	}

	resp, err := cfg.WalletClient.Deposit(ctx, &walletpb.DepositRequest{
		UserId:    userID,
		Amount:    req.Amount,
		Method:    req.Method,
		PaymentId: req.PaymentID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Deposit successful",
		"transaction_id": resp.TransactionId,
		"new_balance":    resp.NewBalance,
	})
}

// Withdraw handles withdraw request
func (cfg *RouterConfig) Withdraw(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	var req struct {
		Amount float64 `json:"amount"`
		Method string  `json:"method"`
		BankID string  `json:"bankId"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.WalletClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Wallet service unavailable", nil)
		return
	}

	resp, err := cfg.WalletClient.Withdraw(ctx, &walletpb.WithdrawRequest{
		UserId: userID,
		Amount: req.Amount,
		Method: req.Method,
		BankId: req.BankID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Withdrawal request submitted",
		"transaction_id": resp.TransactionId,
	})
}

// ListGames handles listing games
func (cfg *RouterConfig) ListGames(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	// Parse query parameters
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")
	category := c.Query("category")
	provider := c.Query("provider")
	search := c.Query("search")

	resp, err := cfg.GameClient.ListGames(ctx, &gamepb.ListGamesRequest{
		Page:     page,
		Limit:    limit,
		Category: category,
		Provider: provider,
		Search:   search,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	games := make([]map[string]interface{}, len(resp.Games))
	for i, game := range resp.Games {
		games[i] = map[string]interface{}{
			"id":          game.Id,
			"name":        game.Name,
			"provider":    game.Provider,
			"category":    game.Category,
			"thumbnail":   game.Thumbnail,
			"rtp":         game.Rtp,
			"min_bet":     game.MinBet,
			"max_bet":     game.MaxBet,
			"volatility":  game.Volatility,
			"status":      game.Status,
			"is_featured": game.IsFeatured,
			"is_new":      game.IsNew,
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": games,
		"total": resp.Total,
		"page":  resp.Page,
	})
}

// GetGame handles getting game details
func (cfg *RouterConfig) GetGame(ctx context.Context, c *app.RequestContext) {
	gameID := c.Param("id")

	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.GetGame(ctx, &gamepb.GetGameRequest{
		GameId: gameID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Game not found", nil)
		return
	}

	game := resp.Game
	handler.SendSuccess(c, map[string]interface{}{
		"game": map[string]interface{}{
			"id":          game.Id,
			"name":        game.Name,
			"provider":    game.Provider,
			"category":    game.Category,
			"description": game.Description,
			"thumbnail":   game.Thumbnail,
			"images":      game.Images,
			"rtp":         game.Rtp,
			"min_bet":     game.MinBet,
			"max_bet":     game.MaxBet,
			"volatility":  game.Volatility,
			"features":    game.Features,
			"paylines":    game.Paylines,
			"reels":       game.Reels,
			"status":      game.Status,
		},
	})
}

// PlayGame handles game play URL generation
func (cfg *RouterConfig) PlayGame(ctx context.Context, c *app.RequestContext) {
	gameID := c.Param("id")
	userID := c.GetString("user_id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.GetGameURL(ctx, &gamepb.GetGameURLRequest{
		GameId: gameID,
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"game_id": gameID,
		"url":     resp.Url,
		"token":   resp.Token,
	})
}

// GetCategories handles getting game categories
func (cfg *RouterConfig) GetCategories(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.GetCategories(ctx, &gamepb.GetCategoriesRequest{})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	categories := make([]map[string]interface{}, len(resp.Categories))
	for i, cat := range resp.Categories {
		categories[i] = map[string]interface{}{
			"id":         cat.Id,
			"name":       cat.Name,
			"slug":       cat.Slug,
			"icon":       cat.Icon,
			"game_count": cat.GameCount,
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"categories": categories,
	})
}

// GetFeaturedGames handles getting featured games
func (cfg *RouterConfig) GetFeaturedGames(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.GetFeaturedGames(ctx, &gamepb.GetFeaturedGamesRequest{
		Limit: 10,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	games := make([]map[string]interface{}, len(resp.Games))
	for i, game := range resp.Games {
		games[i] = map[string]interface{}{
			"id":        game.Id,
			"name":      game.Name,
			"provider":  game.Provider,
			"category":  game.Category,
			"thumbnail": game.Thumbnail,
			"rtp":       game.Rtp,
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": games,
	})
}

// GetPopularGames handles getting popular games
func (cfg *RouterConfig) GetPopularGames(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.GetPopularGames(ctx, &gamepb.GetPopularGamesRequest{
		Limit: 20,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	games := make([]map[string]interface{}, len(resp.Games))
	for i, game := range resp.Games {
		games[i] = map[string]interface{}{
			"id":         game.Id,
			"name":       game.Name,
			"provider":   game.Provider,
			"category":   game.Category,
			"thumbnail":  game.Thumbnail,
			"play_count": game.PlayCount,
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": games,
	})
}

// GetSlotGames handles getting slot games
func (cfg *RouterConfig) GetSlotGames(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.ListGames(ctx, &gamepb.ListGamesRequest{
		Category: "slot",
		Page:     c.DefaultQuery("page", "1"),
		Limit:    c.DefaultQuery("limit", "20"),
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": resp.Games,
		"total": resp.Total,
	})
}

// GetCardGames handles getting card games
func (cfg *RouterConfig) GetCardGames(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.ListGames(ctx, &gamepb.ListGamesRequest{
		Category: "card",
		Page:     c.DefaultQuery("page", "1"),
		Limit:    c.DefaultQuery("limit", "20"),
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": resp.Games,
		"total": resp.Total,
	})
}

// GetDiceGames handles getting dice games
func (cfg *RouterConfig) GetDiceGames(ctx context.Context, c *app.RequestContext) {
	if cfg.GameClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Game service unavailable", nil)
		return
	}

	resp, err := cfg.GameClient.ListGames(ctx, &gamepb.ListGamesRequest{
		Category: "dice",
		Page:     c.DefaultQuery("page", "1"),
		Limit:    c.DefaultQuery("limit", "20"),
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"games": resp.Games,
		"total": resp.Total,
	})
}

// ListTournaments handles listing tournaments
func (cfg *RouterConfig) ListTournaments(ctx context.Context, c *app.RequestContext) {
	if cfg.TournamentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Tournament service unavailable", nil)
		return
	}

	resp, err := cfg.TournamentClient.ListTournaments(ctx, &tournamentpb.ListTournamentsRequest{
		Status: c.Query("status"),
		Page:   c.DefaultQuery("page", "1"),
		Limit:  c.DefaultQuery("limit", "20"),
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"tournaments": resp.Tournaments,
		"total":       resp.Total,
	})
}

// GetTournament handles getting tournament details
func (cfg *RouterConfig) GetTournament(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")

	if cfg.TournamentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Tournament service unavailable", nil)
		return
	}

	resp, err := cfg.TournamentClient.GetTournament(ctx, &tournamentpb.GetTournamentRequest{
		TournamentId: tournamentID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Tournament not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"tournament": resp.Tournament,
	})
}

// JoinTournament handles joining a tournament
func (cfg *RouterConfig) JoinTournament(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")
	userID := c.GetString("user_id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.TournamentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Tournament service unavailable", nil)
		return
	}

	resp, err := cfg.TournamentClient.JoinTournament(ctx, &tournamentpb.JoinTournamentRequest{
		TournamentId: tournamentID,
		UserId:       userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":       "Joined tournament successfully",
		"tournament_id": tournamentID,
		"position":      resp.Position,
	})
}

// GetTournamentLeaderboard handles getting tournament leaderboard
func (cfg *RouterConfig) GetTournamentLeaderboard(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")

	if cfg.TournamentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Tournament service unavailable", nil)
		return
	}

	resp, err := cfg.TournamentClient.GetLeaderboard(ctx, &tournamentpb.GetLeaderboardRequest{
		TournamentId: tournamentID,
		Limit:        50,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"leaderboard": resp.Entries,
	})
}

// ListJackpots handles listing jackpots
func (cfg *RouterConfig) ListJackpots(ctx context.Context, c *app.RequestContext) {
	if cfg.JackpotClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Jackpot service unavailable", nil)
		return
	}

	resp, err := cfg.JackpotClient.ListJackpots(ctx, &jackpotpb.ListJackpotsRequest{
		Status: c.Query("status"),
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"jackpots": resp.Jackpots,
	})
}

// GetJackpot handles getting jackpot details
func (cfg *RouterConfig) GetJackpot(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")

	if cfg.JackpotClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Jackpot service unavailable", nil)
		return
	}

	resp, err := cfg.JackpotClient.GetJackpot(ctx, &jackpotpb.GetJackpotRequest{
		JackpotId: jackpotID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Jackpot not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"jackpot": resp.Jackpot,
	})
}

// GetJackpotWinners handles getting jackpot winners
func (cfg *RouterConfig) GetJackpotWinners(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")

	if cfg.JackpotClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Jackpot service unavailable", nil)
		return
	}

	resp, err := cfg.JackpotClient.GetWinners(ctx, &jackpotpb.GetWinnersRequest{
		JackpotId: jackpotID,
		Limit:     20,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"winners": resp.Winners,
	})
}

// ListBonuses handles listing available bonuses
func (cfg *RouterConfig) ListBonuses(ctx context.Context, c *app.RequestContext) {
	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.ListBonuses(ctx, &bonuspb.ListBonusesRequest{
		Status: "active",
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"bonuses": resp.Bonuses,
	})
}

// GetBonus handles getting bonus details
func (cfg *RouterConfig) GetBonus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.GetBonus(ctx, &bonuspb.GetBonusRequest{
		BonusId: bonusID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Bonus not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"bonus": resp.Bonus,
	})
}

// ClaimBonus handles claiming a bonus
func (cfg *RouterConfig) ClaimBonus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")
	userID := c.GetString("user_id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.ClaimBonus(ctx, &bonuspb.ClaimBonusRequest{
		BonusId: bonusID,
		UserId:  userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":      "Bonus claimed successfully",
		"bonus_amount": resp.BonusAmount,
		"expires_at":   resp.ExpiresAt,
	})
}

// GetMyBonuses handles getting user's claimed bonuses
func (cfg *RouterConfig) GetMyBonuses(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.GetUserBonuses(ctx, &bonuspb.GetUserBonusesRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"bonuses": resp.Bonuses,
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
		AffiliateID  string  `json:"affiliateId"`
		CommissionID string  `json:"commissionId"`
		Amount       float64 `json:"amount"`
		ClaimReason  string  `json:"claimReason"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.SubmitClaim(ctx, &commissionpb.SubmitClaimRequest{
		UserId:       userID,
		AffiliateId:  req.AffiliateID,
		CommissionId: req.CommissionID,
		Amount:       req.Amount,
		ClaimReason:  req.ClaimReason,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":  "Commission claim submitted",
		"claim_id": resp.ClaimId,
		"status":   resp.Status,
	})
}

// GetUserCommissionClaims handles getting user's commission claims
func (cfg *RouterConfig) GetUserCommissionClaims(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetUserClaims(ctx, &commissionpb.GetUserClaimsRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
	})
}

// GetCommissionClaimsByStatus handles getting commission claims by status
func (cfg *RouterConfig) GetCommissionClaimsByStatus(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	status := c.Param("status")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetClaimsByStatus(ctx, &commissionpb.GetClaimsByStatusRequest{
		UserId: userID,
		Status: status,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
		"status": status,
	})
}

// ClaimCommission handles claiming a commission
func (cfg *RouterConfig) ClaimCommission(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	claimID := c.Param("id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.ClaimCommission(ctx, &commissionpb.ClaimCommissionRequest{
		ClaimId: claimID,
		UserId:  userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Commission claimed",
		"claim_id":       claimID,
		"amount":         resp.Amount,
		"transaction_id": resp.TransactionId,
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
		BonusID          string  `json:"bonusId"`
		BonusCode        string  `json:"bonusCode"`
		BonusAmount      float64 `json:"bonusAmount"`
		RebetRequirement float64 `json:"rebetRequirement"`
		GameID           string  `json:"gameId"`
		BetID            string  `json:"betId"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.CreateRebetClaim(ctx, &bonuspb.CreateRebetClaimRequest{
		UserId:           userID,
		BonusId:          req.BonusID,
		BonusCode:        req.BonusCode,
		BonusAmount:      req.BonusAmount,
		RebetRequirement: req.RebetRequirement,
		GameId:           req.GameID,
		BetId:            req.BetID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":           "Rebet claim created",
		"rebet_id":          resp.RebetId,
		"status":            resp.Status,
		"rebet_requirement": resp.RebetRequirement,
		"current_rebet":     resp.CurrentRebet,
	})
}

// GetUserRebetClaims handles getting user's rebet claims
func (cfg *RouterConfig) GetUserRebetClaims(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.GetUserRebetClaims(ctx, &bonuspb.GetUserRebetClaimsRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
	})
}

// GetClaimableRebets handles getting claimable rebet bonuses
func (cfg *RouterConfig) GetClaimableRebets(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.GetClaimableRebets(ctx, &bonuspb.GetClaimableRebetsRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
	})
}

// ClaimRebet handles claiming a rebet bonus
func (cfg *RouterConfig) ClaimRebet(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	rebetID := c.Param("id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.ClaimRebet(ctx, &bonuspb.ClaimRebetRequest{
		RebetId: rebetID,
		UserId:  userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Rebet bonus claimed",
		"rebet_id":       rebetID,
		"amount":         resp.Amount,
		"transaction_id": resp.TransactionId,
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
		GameID            string  `json:"gameId"`
		BetID             string  `json:"betId"`
		InsurancePolicyID string  `json:"insurancePolicyId"`
		ClaimType         string  `json:"claimType"`
		InsuredAmount     float64 `json:"insuredAmount"`
		LossAmount        float64 `json:"lossAmount"`
		ClaimReason       string  `json:"claimReason"`
		EvidenceDetails   string  `json:"evidenceDetails"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.SubmitInsuranceClaim(ctx, &bonuspb.SubmitInsuranceClaimRequest{
		UserId:            userID,
		GameId:            req.GameID,
		BetId:             req.BetID,
		InsurancePolicyId: req.InsurancePolicyID,
		ClaimType:         req.ClaimType,
		InsuredAmount:     req.InsuredAmount,
		LossAmount:        req.LossAmount,
		ClaimReason:       req.ClaimReason,
		EvidenceDetails:   req.EvidenceDetails,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":  "Insurance claim submitted",
		"claim_id": resp.ClaimId,
		"status":   resp.Status,
	})
}

// GetUserInsuranceClaims handles getting user's insurance claims
func (cfg *RouterConfig) GetUserInsuranceClaims(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.BonusClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Bonus service unavailable", nil)
		return
	}

	resp, err := cfg.BonusClient.GetUserInsuranceClaims(ctx, &bonuspb.GetUserInsuranceClaimsRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"claims": resp.Claims,
	})
}

// GetUserSettlements handles getting user's settlements
func (cfg *RouterConfig) GetUserSettlements(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetUserSettlements(ctx, &commissionpb.GetUserSettlementsRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"settlements": resp.Settlements,
	})
}

// GetSettlementById handles getting settlement by ID
func (cfg *RouterConfig) GetSettlementById(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	settlementID := c.Param("id")

	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetSettlementById(ctx, &commissionpb.GetSettlementByIdRequest{
		SettlementId: settlementID,
		UserId:       userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Settlement not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"settlement": resp.Settlement,
	})
}

// GetUserTotalPending handles getting user's total pending claims
func (cfg *RouterConfig) GetUserTotalPending(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetTotalPending(ctx, &commissionpb.GetTotalPendingRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"totalPending": resp.TotalPending,
	})
}

// GetUserTotalSettled handles getting user's total settled claims
func (cfg *RouterConfig) GetUserTotalSettled(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetTotalSettled(ctx, &commissionpb.GetTotalSettledRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"totalSettled": resp.TotalSettled,
	})
}
