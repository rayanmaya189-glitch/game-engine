package router

import (
	"handler"
	"net/http"
)

func notImplemented(name string) interface{} {
	return func(c interface{}, ctx interface{}) {
		if rc, ok := ctx.(interface {
			JSON(int, interface{})
			GetString(string) string
		}); ok {
			rc.JSON(http.StatusNotImplemented, map[string]interface{}{
				"success": false,
				"error":   handler.NewErrorResponse(handler.ErrCodeServiceUnavailable, name+" not yet implemented"),
			})
		}
	}
}

// Player handler placeholders
func registerHandler(opts *RouterOption) interface{}       { return notImplemented("register") }
func loginHandler(opts *RouterOption) interface{}          { return notImplemented("login") }
func refreshTokenHandler(opts *RouterOption) interface{}   { return notImplemented("refresh token") }
func logoutHandler(opts *RouterOption) interface{}         { return notImplemented("logout") }
func getProfileHandler(opts *RouterOption) interface{}     { return notImplemented("get profile") }
func updateProfileHandler(opts *RouterOption) interface{}  { return notImplemented("update profile") }
func getBalanceHandler(opts *RouterOption) interface{}     { return notImplemented("get balance") }
func getTransactionsHandler(opts *RouterOption) interface{} { return notImplemented("get transactions") }
func depositHandler(opts *RouterOption) interface{}        { return notImplemented("deposit") }
func withdrawHandler(opts *RouterOption) interface{}       { return notImplemented("withdraw") }
func listGamesHandler(opts *RouterOption) interface{}      { return notImplemented("list games") }
func getGameHandler(opts *RouterOption) interface{}        { return notImplemented("get game") }
func playGameHandler(opts *RouterOption) interface{}       { return notImplemented("play game") }
func getCategoriesHandler(opts *RouterOption) interface{}  { return notImplemented("get categories") }
func getFeaturedGamesHandler(opts *RouterOption) interface{} { return notImplemented("get featured games") }
func getPopularGamesHandler(opts *RouterOption) interface{}  { return notImplemented("get popular games") }

// Admin handler placeholders
func listPlayersHandler(opts *RouterOption) interface{}      { return notImplemented("list players") }
func getPlayerHandler(opts *RouterOption) interface{}        { return notImplemented("get player") }
func updatePlayerStatusHandler(opts *RouterOption) interface{} { return notImplemented("update player status") }
func getPlayerStatsHandler(opts *RouterOption) interface{}   { return notImplemented("get player stats") }
func getKYCListHandler(opts *RouterOption) interface{}       { return notImplemented("get KYC list") }
func approveKYCHandler(opts *RouterOption) interface{}      { return notImplemented("approve KYC") }
func rejectKYCHandler(opts *RouterOption) interface{}       { return notImplemented("reject KYC") }
func listAdminGamesHandler(opts *RouterOption) interface{}  { return notImplemented("list admin games") }
func createGameHandler(opts *RouterOption) interface{}      { return notImplemented("create game") }
func updateGameHandler(opts *RouterOption) interface{}      { return notImplemented("update game") }
func getAllTransactionsHandler(opts *RouterOption) interface{} { return notImplemented("get all transactions") }
func adjustBalanceHandler(opts *RouterOption) interface{}   { return notImplemented("adjust balance") }

// Merchant handler placeholders
func listMerchantPlayersHandler(opts *RouterOption) interface{} { return notImplemented("list merchant players") }
func getMerchantPlayerHandler(opts *RouterOption) interface{}   { return notImplemented("get merchant player") }
func getRevenueReportsHandler(opts *RouterOption) interface{}   { return notImplemented("get revenue reports") }
func getPlayerReportsHandler(opts *RouterOption) interface{}    { return notImplemented("get player reports") }
func getMerchantConfigHandler(opts *RouterOption) interface{}   { return notImplemented("get merchant config") }
func updateMerchantConfigHandler(opts *RouterOption) interface{} { return notImplemented("update merchant config") }
func registerWebhookHandler(opts *RouterOption) interface{}   { return notImplemented("register webhook") }
func listWebhooksHandler(opts *RouterOption) interface{}      { return notImplemented("list webhooks") }

// Agent handler placeholders
func listAgentPlayersHandler(opts *RouterOption) interface{} { return notImplemented("list agent players") }
func getAgentPlayerHandler(opts *RouterOption) interface{}   { return notImplemented("get agent player") }
func getCommissionsHandler(opts *RouterOption) interface{}   { return notImplemented("get commissions") }
func getPendingCommissionsHandler(opts *RouterOption) interface{} { return notImplemented("get pending commissions") }
func trackClickHandler(opts *RouterOption) interface{}       { return notImplemented("track click") }
func redirectToRegistrationHandler(opts *RouterOption) interface{} { return notImplemented("redirect to registration") }
func getPerformanceReportsHandler(opts *RouterOption) interface{} { return notImplemented("get performance reports") }
