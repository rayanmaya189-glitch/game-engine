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

	// Claims Management
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

	// Merchants Management
	admin.GET("/merchants", cfg.ListMerchants)
	admin.GET("/merchants/:id", cfg.GetMerchant)
	admin.POST("/merchants", cfg.CreateMerchant)
	admin.PUT("/merchants/:id", cfg.UpdateMerchant)
	admin.PUT("/merchants/:id/status", cfg.UpdateMerchantStatus)
	admin.DELETE("/merchants/:id", cfg.DeleteMerchant)

	// Agents Management
	admin.GET("/agents", cfg.ListAgents)
	admin.GET("/agents/:id", cfg.GetAgent)
	admin.POST("/agents", cfg.CreateAgent)
	admin.PUT("/agents/:id", cfg.UpdateAgent)
	admin.PUT("/agents/:id/status", cfg.UpdateAgentStatus)
	admin.DELETE("/agents/:id", cfg.DeleteAgent)

	// Tournaments Management
	admin.GET("/tournaments", cfg.ListTournaments)
	admin.GET("/tournaments/:id", cfg.GetTournament)
	admin.POST("/tournaments", cfg.CreateTournament)
	admin.PUT("/tournaments/:id", cfg.UpdateTournament)
	admin.PUT("/tournaments/:id/status", cfg.UpdateTournamentStatus)
	admin.DELETE("/tournaments/:id", cfg.DeleteTournament)
	admin.GET("/tournaments/:id/leaderboard", cfg.GetTournamentLeaderboard)

	// Jackpots Management
	admin.GET("/jackpots", cfg.ListJackpots)
	admin.GET("/jackpots/:id", cfg.GetJackpot)
	admin.POST("/jackpots", cfg.CreateJackpot)
	admin.PUT("/jackpots/:id", cfg.UpdateJackpot)
	admin.PUT("/jackpots/:id/status", cfg.UpdateJackpotStatus)
	admin.DELETE("/jackpots/:id", cfg.DeleteJackpot)
	admin.GET("/jackpots/:id/hits", cfg.GetJackpotHits)

	// Bonuses Management
	admin.GET("/bonuses", cfg.ListBonuses)
	admin.GET("/bonuses/:id", cfg.GetBonus)
	admin.POST("/bonuses", cfg.CreateBonus)
	admin.PUT("/bonuses/:id", cfg.UpdateBonus)
	admin.PUT("/bonuses/:id/status", cfg.UpdateBonusStatus)
	admin.DELETE("/bonuses/:id", cfg.DeleteBonus)

	// Payments Management
	admin.GET("/payments", cfg.ListPayments)
	admin.GET("/payments/:id", cfg.GetPayment)
	admin.PUT("/payments/:id/approve", cfg.ApprovePayment)
	admin.PUT("/payments/:id/reject", cfg.RejectPayment)
	admin.PUT("/payments/:id/process", cfg.ProcessPayment)

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

// Commission Claims Handlers
func (cfg *RouterConfig) ListCommissionClaims(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"claims":     []interface{}{},
		"total":      0,
		"page":       1,
		"totalPages": 0,
	})
}

func (cfg *RouterConfig) GetCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":          claimID,
		"claimType":   "COMMISSION",
		"amount":      "100.00",
		"status":      "PENDING",
		"claimReason": "Commission claim request",
	})
}

func (cfg *RouterConfig) ApproveCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "APPROVED",
		"message": "Commission claim approved",
	})
}

func (cfg *RouterConfig) RejectCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "REJECTED",
		"message": "Commission claim rejected",
	})
}

func (cfg *RouterConfig) PayCommissionClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":            claimID,
		"status":        "PAID",
		"message":       "Commission claim paid",
		"transactionId": "txn_" + claimID,
	})
}

// Rebet Claims Handlers
func (cfg *RouterConfig) ListRebetClaims(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"claims": []interface{}{},
		"total":  0,
	})
}

func (cfg *RouterConfig) GetRebetClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":        claimID,
		"bonusCode": "BONUS123",
		"amount":    "50.00",
		"status":    "CLAIMABLE",
	})
}

func (cfg *RouterConfig) ApproveRebetClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "APPROVED",
		"message": "Rebet claim approved",
	})
}

func (cfg *RouterConfig) RejectRebetClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "REJECTED",
		"message": "Rebet claim rejected",
	})
}

// Insurance Claims Handlers
func (cfg *RouterConfig) ListInsuranceClaims(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"claims": []interface{}{},
		"total":  0,
	})
}

func (cfg *RouterConfig) GetInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":          claimID,
		"claimType":   "GAME_LOSS",
		"claimAmount": "200.00",
		"status":      "PENDING",
	})
}

func (cfg *RouterConfig) ApproveInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "APPROVED",
		"message": "Insurance claim approved",
	})
}

func (cfg *RouterConfig) RejectInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      claimID,
		"status":  "REJECTED",
		"message": "Insurance claim rejected",
	})
}

func (cfg *RouterConfig) PayInsuranceClaim(ctx context.Context, c *app.RequestContext) {
	claimID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":            claimID,
		"status":        "PAID",
		"message":       "Insurance claim paid",
		"transactionId": "ins_txn_" + claimID,
	})
}

// Settlements Handlers
func (cfg *RouterConfig) ListSettlements(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"settlements": []interface{}{},
		"total":       0,
	})
}

func (cfg *RouterConfig) GetSettlement(ctx context.Context, c *app.RequestContext) {
	settlementID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":             settlementID,
		"settlementType": "COMMISSION",
		"amount":         "100.00",
		"status":         "COMPLETED",
	})
}

func (cfg *RouterConfig) GetClaimStatistics(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"totalPending":    15,
		"totalApproved":   25,
		"totalPaid":       100,
		"totalRejected":   5,
		"totalInProgress": 10,
		"totalAmount":     "15000.00",
	})
}

// Merchants Handlers
func (cfg *RouterConfig) ListMerchants(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"merchants": []interface{}{},
		"total":     0,
	})
}

func (cfg *RouterConfig) GetMerchant(ctx context.Context, c *app.RequestContext) {
	merchantID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":             merchantID,
		"name":           "Merchant Name",
		"email":          "merchant@example.com",
		"commissionRate": 10,
		"status":         "active",
	})
}

func (cfg *RouterConfig) CreateMerchant(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"id":      "new_merchant_id",
		"message": "Merchant created successfully",
	})
}

func (cfg *RouterConfig) UpdateMerchant(ctx context.Context, c *app.RequestContext) {
	merchantID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      merchantID,
		"message": "Merchant updated successfully",
	})
}

func (cfg *RouterConfig) UpdateMerchantStatus(ctx context.Context, c *app.RequestContext) {
	merchantID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      merchantID,
		"status":  "updated",
		"message": "Merchant status updated",
	})
}

func (cfg *RouterConfig) DeleteMerchant(ctx context.Context, c *app.RequestContext) {
	merchantID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      merchantID,
		"message": "Merchant deleted successfully",
	})
}

// Agents Handlers
func (cfg *RouterConfig) ListAgents(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"agents": []interface{}{},
		"total":  0,
	})
}

func (cfg *RouterConfig) GetAgent(ctx context.Context, c *app.RequestContext) {
	agentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":     agentID,
		"name":   "Agent Name",
		"email":  "agent@example.com",
		"tier":   "Gold",
		"status": "active",
	})
}

func (cfg *RouterConfig) CreateAgent(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"id":      "new_agent_id",
		"message": "Agent created successfully",
	})
}

func (cfg *RouterConfig) UpdateAgent(ctx context.Context, c *app.RequestContext) {
	agentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      agentID,
		"message": "Agent updated successfully",
	})
}

func (cfg *RouterConfig) UpdateAgentStatus(ctx context.Context, c *app.RequestContext) {
	agentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      agentID,
		"status":  "updated",
		"message": "Agent status updated",
	})
}

func (cfg *RouterConfig) DeleteAgent(ctx context.Context, c *app.RequestContext) {
	agentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      agentID,
		"message": "Agent deleted successfully",
	})
}

// Tournaments Handlers
func (cfg *RouterConfig) ListTournaments(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"tournaments": []interface{}{},
		"total":       0,
	})
}

func (cfg *RouterConfig) GetTournament(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":        tournamentID,
		"name":      "Tournament Name",
		"prizePool": 5000,
		"game":      "Slots",
		"status":    "active",
	})
}

func (cfg *RouterConfig) CreateTournament(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"id":      "new_tournament_id",
		"message": "Tournament created successfully",
	})
}

func (cfg *RouterConfig) UpdateTournament(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      tournamentID,
		"message": "Tournament updated successfully",
	})
}

func (cfg *RouterConfig) UpdateTournamentStatus(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      tournamentID,
		"status":  "updated",
		"message": "Tournament status updated",
	})
}

func (cfg *RouterConfig) DeleteTournament(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      tournamentID,
		"message": "Tournament deleted successfully",
	})
}

func (cfg *RouterConfig) GetTournamentLeaderboard(ctx context.Context, c *app.RequestContext) {
	tournamentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"tournamentId": tournamentID,
		"leaderboard":  []interface{}{},
	})
}

// Jackpots Handlers
func (cfg *RouterConfig) ListJackpots(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"jackpots": []interface{}{},
		"total":    0,
	})
}

func (cfg *RouterConfig) GetJackpot(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":            jackpotID,
		"name":          "Jackpot Name",
		"currentAmount": 50000,
		"game":          "Mega Moolah",
		"status":        "active",
	})
}

func (cfg *RouterConfig) CreateJackpot(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"id":      "new_jackpot_id",
		"message": "Jackpot created successfully",
	})
}

func (cfg *RouterConfig) UpdateJackpot(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      jackpotID,
		"message": "Jackpot updated successfully",
	})
}

func (cfg *RouterConfig) UpdateJackpotStatus(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      jackpotID,
		"status":  "updated",
		"message": "Jackpot status updated",
	})
}

func (cfg *RouterConfig) DeleteJackpot(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      jackpotID,
		"message": "Jackpot deleted successfully",
	})
}

func (cfg *RouterConfig) GetJackpotHits(ctx context.Context, c *app.RequestContext) {
	jackpotID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"jackpotId": jackpotID,
		"hits":      []interface{}{},
		"total":     0,
	})
}

// Bonuses Handlers
func (cfg *RouterConfig) ListBonuses(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"bonuses": []interface{}{},
		"total":   0,
	})
}

func (cfg *RouterConfig) GetBonus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":       bonusID,
		"name":     "Bonus Name",
		"type":     "Deposit",
		"amount":   100,
		"maxBonus": 500,
		"wagerReq": 35,
		"status":   "active",
	})
}

func (cfg *RouterConfig) CreateBonus(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"id":      "new_bonus_id",
		"message": "Bonus created successfully",
	})
}

func (cfg *RouterConfig) UpdateBonus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      bonusID,
		"message": "Bonus updated successfully",
	})
}

func (cfg *RouterConfig) UpdateBonusStatus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      bonusID,
		"status":  "updated",
		"message": "Bonus status updated",
	})
}

func (cfg *RouterConfig) DeleteBonus(ctx context.Context, c *app.RequestContext) {
	bonusID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      bonusID,
		"message": "Bonus deleted successfully",
	})
}

// Payments Handlers
func (cfg *RouterConfig) ListPayments(ctx context.Context, c *app.RequestContext) {
	handler.SendSuccess(c, map[string]interface{}{
		"payments": []interface{}{},
		"total":    0,
	})
}

func (cfg *RouterConfig) GetPayment(ctx context.Context, c *app.RequestContext) {
	paymentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":     paymentID,
		"userId": "user_123",
		"amount": 100,
		"method": "Bank Transfer",
		"type":   "deposit",
		"status": "pending",
	})
}

func (cfg *RouterConfig) ApprovePayment(ctx context.Context, c *app.RequestContext) {
	paymentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      paymentID,
		"status":  "approved",
		"message": "Payment approved",
	})
}

func (cfg *RouterConfig) RejectPayment(ctx context.Context, c *app.RequestContext) {
	paymentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":      paymentID,
		"status":  "rejected",
		"message": "Payment rejected",
	})
}

func (cfg *RouterConfig) ProcessPayment(ctx context.Context, c *app.RequestContext) {
	paymentID := c.Param("id")
	handler.SendSuccess(c, map[string]interface{}{
		"id":            paymentID,
		"status":        "completed",
		"message":       "Payment processed",
		"transactionId": "txn_" + paymentID,
	})
}
