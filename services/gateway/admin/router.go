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
	RBACMiddleware        *middleware.RBACMiddleware
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
	AllowedIPs            []string
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

	// Admin routes with JWT + MFA + RBAC permission check
	admin := r.Group("/api/v1/admin")
	admin.Use(cfg.AuthMiddleware.JWTValidation())
	admin.Use(cfg.AuthMiddleware.MFACheck())
	admin.Use(cfg.AuthMiddleware.RequireAnyAdminRole())
	admin.Use(cfg.RBACMiddleware.AutoRoutePermission())

	// Admin IP whitelist
	if len(cfg.AllowedIPs) > 0 {
		admin.Use(cfg.AuthMiddleware.IPWhitelistCheck(cfg.AllowedIPs))
	}

	// Player management (permission: players:view, players:edit)
	admin.GET("/players", cfg.ListPlayers)
	admin.GET("/players/:id", cfg.GetPlayer)
	admin.PUT("/players/:id/status", cfg.UpdatePlayerStatus)
	admin.GET("/players/:id/stats", cfg.GetPlayerStats)

	// KYC management (permission: kyc:view, kyc:approve)
	admin.GET("/kyc", cfg.GetKYCList)
	admin.PUT("/kyc/:id/approve", cfg.ApproveKYC)
	admin.PUT("/kyc/:id/reject", cfg.RejectKYC)

	// Game management (permission: games:view, games:create, games:edit)
	admin.GET("/games", cfg.ListAdminGames)
	admin.POST("/games", cfg.CreateGame)
	admin.PUT("/games/:id", cfg.UpdateGame)

	// Wallet management (permission: wallet:view, wallet:adjust)
	admin.GET("/wallet/transactions", cfg.GetAllTransactions)
	admin.POST("/wallet/adjust", cfg.AdjustBalance)

	// Claims Management (permission: commission:view, commission:approve)
	admin.GET("/claims/commission", cfg.ListCommissionClaims)
	admin.GET("/claims/commission/:id", cfg.GetCommissionClaim)
	admin.POST("/claims/commission/:id/approve", cfg.ApproveCommissionClaim)
	admin.POST("/claims/commission/:id/reject", cfg.RejectCommissionClaim)
	admin.POST("/claims/commission/:id/pay", cfg.PayCommissionClaim)

	admin.GET("/claims/rebet", cfg.ListRebetClaims)
	admin.GET("/claims/rebet/:id", cfg.GetRebetClaim)
	admin.POST("/claims/rebet/:id/approve", cfg.ApproveRebetClaim)
	admin.POST("/claims/rebet/:id/reject", cfg.RejectRebetClaim)

	admin.GET("/claims/insurance", cfg.ListInsuranceClaims)
	admin.GET("/claims/insurance/:id", cfg.GetInsuranceClaim)
	admin.POST("/claims/insurance/:id/approve", cfg.ApproveInsuranceClaim)
	admin.POST("/claims/insurance/:id/reject", cfg.RejectInsuranceClaim)
	admin.POST("/claims/insurance/:id/pay", cfg.PayInsuranceClaim)

	admin.GET("/claims/settlements", cfg.ListSettlements)
	admin.GET("/claims/settlements/:id", cfg.GetSettlement)
	admin.GET("/claims/statistics", cfg.GetClaimStatistics)

	// Merchants Management (permission: merchants:view, merchants:create, merchants:edit, merchants:delete)
	admin.GET("/merchants", cfg.ListMerchants)
	admin.GET("/merchants/:id", cfg.GetMerchant)
	admin.POST("/merchants", cfg.CreateMerchant)
	admin.PUT("/merchants/:id", cfg.UpdateMerchant)
	admin.PUT("/merchants/:id/status", cfg.UpdateMerchantStatus)
	admin.DELETE("/merchants/:id", cfg.DeleteMerchant)

	// Agents Management (permission: agents:view, agents:create, agents:edit, agents:delete)
	admin.GET("/agents", cfg.ListAgents)
	admin.GET("/agents/:id", cfg.GetAgent)
	admin.POST("/agents", cfg.CreateAgent)
	admin.PUT("/agents/:id", cfg.UpdateAgent)
	admin.PUT("/agents/:id/status", cfg.UpdateAgentStatus)
	admin.DELETE("/agents/:id", cfg.DeleteAgent)

	// Tournaments Management (permission: tournaments:view, tournaments:create, tournaments:edit)
	admin.GET("/tournaments", cfg.ListTournaments)
	admin.GET("/tournaments/:id", cfg.GetTournament)
	admin.POST("/tournaments", cfg.CreateTournament)
	admin.PUT("/tournaments/:id", cfg.UpdateTournament)
	admin.PUT("/tournaments/:id/status", cfg.UpdateTournamentStatus)
	admin.DELETE("/tournaments/:id", cfg.DeleteTournament)
	admin.GET("/tournaments/:id/leaderboard", cfg.GetTournamentLeaderboard)

	// Jackpots Management (permission: jackpots:view, jackpots:create, jackpots:edit)
	admin.GET("/jackpots", cfg.ListJackpots)
	admin.GET("/jackpots/:id", cfg.GetJackpot)
	admin.POST("/jackpots", cfg.CreateJackpot)
	admin.PUT("/jackpots/:id", cfg.UpdateJackpot)
	admin.PUT("/jackpots/:id/status", cfg.UpdateJackpotStatus)
	admin.DELETE("/jackpots/:id", cfg.DeleteJackpot)
	admin.GET("/jackpots/:id/hits", cfg.GetJackpotHits)

	// Bonuses Management (permission: bonuses:view, bonuses:create, bonuses:edit)
	admin.GET("/bonuses", cfg.ListBonuses)
	admin.GET("/bonuses/:id", cfg.GetBonus)
	admin.POST("/bonuses", cfg.CreateBonus)
	admin.PUT("/bonuses/:id", cfg.UpdateBonus)
	admin.PUT("/bonuses/:id/status", cfg.UpdateBonusStatus)
	admin.DELETE("/bonuses/:id", cfg.DeleteBonus)

	// Payments Management (permission: payments:view, payments:approve)
	admin.GET("/payments", cfg.ListPayments)
	admin.GET("/payments/:id", cfg.GetPayment)
	admin.PUT("/payments/:id/approve", cfg.ApprovePayment)
	admin.PUT("/payments/:id/reject", cfg.RejectPayment)
	admin.PUT("/payments/:id/process", cfg.ProcessPayment)

	// Reports (permission: reports:view)
	admin.GET("/reports/*path", cfg.ReportsHandler)

	// RBAC Management (permission: roles:view, admin_users:view)
	admin.GET("/permissions", cfg.ListPermissions)
	admin.GET("/roles", cfg.ListRoles)
	admin.POST("/roles", cfg.CreateRole)
	admin.PUT("/roles/:id", cfg.UpdateRole)
	admin.DELETE("/roles/:id", cfg.DeleteRole)
	admin.GET("/admin-users", cfg.ListAdminUsers)
	admin.POST("/admin-users", cfg.CreateAdminUser)
	admin.PUT("/admin-users/:id", cfg.UpdateAdminUser)
	admin.DELETE("/admin-users/:id", cfg.DeleteAdminUser)

	return r
}
