package main

import (
	"github.com/cloudwego/hertz/pkg/route"
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

func NewRouter(cfg *RouterConfig) *route.Router {
	r := route.New()

	r.Use(cfg.LoggerMiddleware.RequestID())
	r.Use(cfg.LoggerMiddleware.StructuredLogger())
	r.Use(cfg.LoggerMiddleware.PanicRecovery())
	r.Use(cfg.CORSMiddleware.CORS())

	r.GET("/health", handler.HandleHealthCheck)
	r.GET("/ready", handler.HandleReadinessCheck)

	r.Use(cfg.RateLimiterMiddleware.RateLimiter())

	// Auth routes (no JWT required)
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
		users := protected.Group("/users")
		{
			users.GET("/profile", cfg.GetProfile)
			users.PUT("/profile", cfg.UpdateProfile)
		}

		wallet := protected.Group("/wallet")
		{
			wallet.GET("/balance", cfg.GetBalance)
			wallet.GET("/transactions", cfg.GetTransactions)
			wallet.POST("/deposit", cfg.Deposit)
			wallet.POST("/withdraw", cfg.Withdraw)
		}

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

		tournaments := protected.Group("/tournaments")
		{
			tournaments.GET("", cfg.ListTournaments)
			tournaments.GET("/:id", cfg.GetTournament)
			tournaments.POST("/:id/join", cfg.JoinTournament)
			tournaments.GET("/:id/leaderboard", cfg.GetTournamentLeaderboard)
		}

		jackpots := protected.Group("/jackpots")
		{
			jackpots.GET("", cfg.ListJackpots)
			jackpots.GET("/:id", cfg.GetJackpot)
			jackpots.GET("/:id/winners", cfg.GetJackpotWinners)
		}

		bonuses := protected.Group("/bonuses")
		{
			bonuses.GET("", cfg.ListBonuses)
			bonuses.GET("/:id", cfg.GetBonus)
			bonuses.POST("/:id/claim", cfg.ClaimBonus)
			bonuses.GET("/my-bonuses", cfg.GetMyBonuses)
		}

		claims := protected.Group("/claims")
		{
			claims.POST("/commission", cfg.SubmitCommissionClaim)
			claims.GET("/commission", cfg.GetUserCommissionClaims)
			claims.GET("/commission/status/:status", cfg.GetCommissionClaimsByStatus)
			claims.POST("/commission/:id/claim", cfg.ClaimCommission)

			claims.POST("/rebet", cfg.CreateRebetClaim)
			claims.GET("/rebet", cfg.GetUserRebetClaims)
			claims.GET("/rebet/claimable", cfg.GetClaimableRebets)
			claims.POST("/rebet/:id/claim", cfg.ClaimRebet)

			claims.POST("/insurance", cfg.SubmitInsuranceClaim)
			claims.GET("/insurance", cfg.GetUserInsuranceClaims)

			claims.GET("/settlements", cfg.GetUserSettlements)
			claims.GET("/settlements/:id", cfg.GetSettlementById)

			claims.GET("/total-pending", cfg.GetUserTotalPending)
			claims.GET("/total-settled", cfg.GetUserTotalSettled)
		}
	}

	return r
}
