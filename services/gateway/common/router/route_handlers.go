package router

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/game_engine/gateway/common/handler"
)

// HealthCheckHandler returns service health status
func HealthCheckHandler() app.HandlerFunc {
	return func(c interface{}, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, map[string]interface{}{
			"status":    "healthy",
			"service":   "gateway",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
	}
}

// ReadinessHandler returns service readiness status
func ReadinessHandler() app.HandlerFunc {
	return func(c interface{}, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, map[string]interface{}{
			"status":    "ready",
			"service":   "gateway",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		})
	}
}

func notImplemented(name string) app.HandlerFunc {
	return func(c interface{}, ctx *app.RequestContext) {
		requestID := ctx.GetString("request_id")
		ctx.JSON(consts.StatusNotImplemented, map[string]interface{}{
			"success": false,
			"error": handler.NewErrorResponse(
				handler.ErrCodeServiceUnavailable,
				name+" not yet implemented",
			).WithRequestID(requestID),
		})
	}
}

// Auth handlers
func registerHandler(opts *RouterOption) app.HandlerFunc  { return notImplemented("register") }
func loginHandler(opts *RouterOption) app.HandlerFunc     { return notImplemented("login") }
func refreshTokenHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("refresh token")
}
func logoutHandler(opts *RouterOption) app.HandlerFunc { return notImplemented("logout") }

// User handlers
func getProfileHandler(opts *RouterOption) app.HandlerFunc    { return notImplemented("get profile") }
func updateProfileHandler(opts *RouterOption) app.HandlerFunc { return notImplemented("update profile") }

// Wallet handlers
func getBalanceHandler(opts *RouterOption) app.HandlerFunc      { return notImplemented("get balance") }
func getTransactionsHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get transactions")
}
func depositHandler(opts *RouterOption) app.HandlerFunc  { return notImplemented("deposit") }
func withdrawHandler(opts *RouterOption) app.HandlerFunc { return notImplemented("withdraw") }

// Game handlers
func listGamesHandler(opts *RouterOption) app.HandlerFunc        { return notImplemented("list games") }
func getGameHandler(opts *RouterOption) app.HandlerFunc          { return notImplemented("get game") }
func playGameHandler(opts *RouterOption) app.HandlerFunc         { return notImplemented("play game") }
func getCategoriesHandler(opts *RouterOption) app.HandlerFunc    { return notImplemented("get categories") }
func getFeaturedGamesHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get featured games")
}
func getPopularGamesHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get popular games")
}

// Admin handlers
func listPlayersHandler(opts *RouterOption) app.HandlerFunc    { return notImplemented("list players") }
func getPlayerHandler(opts *RouterOption) app.HandlerFunc      { return notImplemented("get player") }
func updatePlayerStatusHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("update player status")
}
func getPlayerStatsHandler(opts *RouterOption) app.HandlerFunc { return notImplemented("get player stats") }
func getKYCListHandler(opts *RouterOption) app.HandlerFunc     { return notImplemented("get KYC list") }
func approveKYCHandler(opts *RouterOption) app.HandlerFunc     { return notImplemented("approve KYC") }
func rejectKYCHandler(opts *RouterOption) app.HandlerFunc      { return notImplemented("reject KYC") }
func listAdminGamesHandler(opts *RouterOption) app.HandlerFunc { return notImplemented("list admin games") }
func createGameHandler(opts *RouterOption) app.HandlerFunc     { return notImplemented("create game") }
func updateGameHandler(opts *RouterOption) app.HandlerFunc     { return notImplemented("update game") }
func getAllTransactionsHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get all transactions")
}
func adjustBalanceHandler(opts *RouterOption) app.HandlerFunc { return notImplemented("adjust balance") }

// Merchant handlers
func listMerchantPlayersHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("list merchant players")
}
func getMerchantPlayerHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get merchant player")
}
func getRevenueReportsHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get revenue reports")
}
func getPlayerReportsHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get player reports")
}
func getMerchantConfigHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get merchant config")
}
func updateMerchantConfigHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("update merchant config")
}
func registerWebhookHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("register webhook")
}
func listWebhooksHandler(opts *RouterOption) app.HandlerFunc { return notImplemented("list webhooks") }

// Agent handlers
func listAgentPlayersHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("list agent players")
}
func getAgentPlayerHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get agent player")
}
func getCommissionsHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get commissions")
}
func getPendingCommissionsHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get pending commissions")
}
func trackClickHandler(opts *RouterOption) app.HandlerFunc { return notImplemented("track click") }
func redirectToRegistrationHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("redirect to registration")
}
func getPerformanceReportsHandler(opts *RouterOption) app.HandlerFunc {
	return notImplemented("get performance reports")
}
