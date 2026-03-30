package router

import (
	"github.com/cloudwego/hertz/pkg/router"

	"handler"
	"middleware"
)

type RouterOption struct {
	AuthMiddleware        *middleware.AuthMiddleware
	LoggerMiddleware      *middleware.LoggerMiddleware
	RateLimiterMiddleware *middleware.RateLimiterMiddleware
	CORSMiddleware        *middleware.CORSMiddleware
	ValidatorMiddleware   *middleware.ValidatorMiddleware
	ErrorHandler          *handler.ErrorHandler
}

func NewRouter(opts *RouterOption) *router.Router {
	r := router.New()

	// Global middleware (applied to all routes)
	if opts.LoggerMiddleware != nil {
		r.Use(opts.LoggerMiddleware.RequestID())
		r.Use(opts.LoggerMiddleware.StructuredLogger())
		r.Use(opts.LoggerMiddleware.PanicRecovery())
	}

	if opts.CORSMiddleware != nil {
		r.Use(opts.CORSMiddleware.CORS())
	}

	// Health check endpoints (no auth required)
	r.GET("/health", handler.HandleHealthCheck)
	r.GET("/ready", handler.HandleReadinessCheck)

	return r
}

func ConfigurePlayerRoutes(r *router.Router, opts *RouterOption) {
	// Apply player-specific middleware
	r.Use(opts.RateLimiterMiddleware.RateLimiter())
	r.Use(opts.AuthMiddleware.JWTValidation())

	// Auth routes (some require auth, some don't)
	authGroup := r.Group("/api/v1/auth")
	{
		authGroup.POST("/register", registerHandler(opts))
		authGroup.POST("/login", loginHandler(opts))
		authGroup.POST("/refresh", refreshTokenHandler(opts))
		authGroup.POST("/logout", logoutHandler(opts))
	}

	// User routes
	userGroup := r.Group("/api/v1/users")
	{
		userGroup.GET("/profile", getProfileHandler(opts))
		userGroup.PUT("/profile", updateProfileHandler(opts))
	}

	// Wallet routes
	walletGroup := r.Group("/api/v1/wallet")
	{
		walletGroup.GET("/balance", getBalanceHandler(opts))
		walletGroup.GET("/transactions", getTransactionsHandler(opts))
		walletGroup.POST("/deposit", depositHandler(opts))
		walletGroup.POST("/withdraw", withdrawHandler(opts))
	}

	// Game routes
	gameGroup := r.Group("/api/v1/games")
	{
		gameGroup.GET("", listGamesHandler(opts))
		gameGroup.GET("/:id", getGameHandler(opts))
		gameGroup.GET("/:id/play", playGameHandler(opts))
		gameGroup.GET("/categories", getCategoriesHandler(opts))
		gameGroup.GET("/featured", getFeaturedGamesHandler(opts))
		gameGroup.GET("/popular", getPopularGamesHandler(opts))
	}
}

func ConfigureAdminRoutes(r *router.Router, opts *RouterOption) {
	// Apply admin-specific middleware
	r.Use(opts.RateLimiterMiddleware.RateLimiter())
	r.Use(opts.AuthMiddleware.JWTValidation())
	r.Use(opts.AuthMiddleware.MFACheck())
	r.Use(opts.AuthMiddleware.RoleCheck("admin"))

	// Admin player routes
	adminPlayersGroup := r.Group("/api/v1/admin/players")
	{
		adminPlayersGroup.GET("", listPlayersHandler(opts))
		adminPlayersGroup.GET("/:id", getPlayerHandler(opts))
		adminPlayersGroup.PUT("/:id/status", updatePlayerStatusHandler(opts))
		adminPlayersGroup.GET("/:id/stats", getPlayerStatsHandler(opts))
	}

	// Admin KYC routes
	adminKYCGroup := r.Group("/api/v1/admin/kyc")
	{
		adminKYCGroup.GET("", getKYCListHandler(opts))
		adminKYCGroup.PUT("/:id/approve", approveKYCHandler(opts))
		adminKYCGroup.PUT("/:id/reject", rejectKYCHandler(opts))
	}

	// Admin game routes
	adminGamesGroup := r.Group("/api/v1/admin/games")
	{
		adminGamesGroup.GET("", listAdminGamesHandler(opts))
		adminGamesGroup.POST("", createGameHandler(opts))
		adminGamesGroup.PUT("/:id", updateGameHandler(opts))
	}

	// Admin wallet routes
	adminWalletGroup := r.Group("/api/v1/admin/wallet")
	{
		adminWalletGroup.GET("/transactions", getAllTransactionsHandler(opts))
		adminWalletGroup.POST("/adjust", adjustBalanceHandler(opts))
	}
}

func ConfigureMerchantRoutes(r *router.Router, opts *RouterOption) {
	// Apply merchant-specific middleware
	r.Use(opts.RateLimiterMiddleware.RateLimiter())
	r.Use(opts.AuthMiddleware.APIKeyValidation())

	// Merchant player routes
	merchantPlayersGroup := r.Group("/api/v1/merchant/players")
	{
		merchantPlayersGroup.GET("", listMerchantPlayersHandler(opts))
		merchantPlayersGroup.GET("/:id", getMerchantPlayerHandler(opts))
	}

	// Merchant reports routes
	merchantReportsGroup := r.Group("/api/v1/merchant/reports")
	{
		merchantReportsGroup.GET("/revenue", getRevenueReportsHandler(opts))
		merchantReportsGroup.GET("/players", getPlayerReportsHandler(opts))
	}

	// Merchant config routes
	merchantConfigGroup := r.Group("/api/v1/merchant/config")
	{
		merchantConfigGroup.GET("", getMerchantConfigHandler(opts))
		merchantConfigGroup.PUT("", updateMerchantConfigHandler(opts))
	}

	// Merchant webhooks routes
	merchantWebhooksGroup := r.Group("/api/v1/merchant/webhooks")
	{
		merchantWebhooksGroup.POST("/register", registerWebhookHandler(opts))
		merchantWebhooksGroup.GET("", listWebhooksHandler(opts))
	}
}

func ConfigureAgentRoutes(r *router.Router, opts *RouterOption) {
	// Apply agent-specific middleware
	r.Use(opts.RateLimiterMiddleware.RateLimiter())
	r.Use(opts.AuthMiddleware.JWTValidation())
	r.Use(opts.AuthMiddleware.APIKeyValidation())

	// Agent player routes
	agentPlayersGroup := r.Group("/api/v1/agent/players")
	{
		agentPlayersGroup.GET("", listAgentPlayersHandler(opts))
		agentPlayersGroup.GET("/:id", getAgentPlayerHandler(opts))
	}

	// Agent commission routes
	agentCommissionsGroup := r.Group("/api/v1/agent/commissions")
	{
		agentCommissionsGroup.GET("", getCommissionsHandler(opts))
		agentCommissionsGroup.GET("/pending", getPendingCommissionsHandler(opts))
	}

	// Affiliate tracking routes
	affiliateTrackingGroup := r.Group("/api/v1/affiliate/tracking")
	{
		affiliateTrackingGroup.POST("/click", trackClickHandler(opts))
		affiliateTrackingGroup.GET("/:code", redirectToRegistrationHandler(opts))
	}

	// Affiliate reports routes
	affiliateReportsGroup := r.Group("/api/v1/affiliate/reports")
	{
		affiliateReportsGroup.GET("/performance", getPerformanceReportsHandler(opts))
	}
}
