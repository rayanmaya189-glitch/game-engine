package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/router"

	affiliatepb "github.com/game-engine/gen/go/game-engine/affiliate/v1"
	agentpb "github.com/game-engine/gen/go/game-engine/agent/v1"

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
	AgentClient           *client.AgentClient
	AffiliateClient       *client.AffiliateClient
	UserClient            *client.UserClient
	CommissionClient      *client.CommissionClient
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

	// Agent routes with API Key + JWT authentication
	agent := r.Group("/api/v1/agent")
	agent.Use(cfg.AuthMiddleware.JWTValidation())
	agent.Use(cfg.AuthMiddleware.APIKeyValidation())

	// Downline players
	agent.GET("/players", cfg.ListAgentPlayers)
	agent.GET("/players/:id", cfg.GetAgentPlayer)
	agent.PUT("/players/:id/limit", cfg.UpdatePlayerLimit)

	// Commission routes
	agent.GET("/commissions", cfg.GetCommissions)
	agent.GET("/commissions/pending", cfg.GetPendingCommissions)
	agent.GET("/commissions/history", cfg.GetCommissionHistory)
	agent.POST("/commissions/claim", cfg.ClaimCommission)

	// Dashboard
	agent.GET("/dashboard", cfg.GetDashboard)

	// Affiliate routes
	affiliate := r.Group("/api/v1/affiliate/tracking")
	affiliate.POST("/click", cfg.TrackClick)
	affiliate.GET("/:code", cfg.RedirectToRegistration)

	// Affiliate reports
	affiliateReports := r.Group("/api/v1/affiliate/reports")
	affiliateReports.GET("/performance", cfg.GetPerformanceReports)
	affiliateReports.GET("/clicks", cfg.GetClickReports)
	affiliateReports.GET("/conversions", cfg.GetConversionReports)

	return r
}

func (cfg *RouterConfig) ListAgentPlayers(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")
	search := c.Query("search")
	status := c.Query("status")

	if cfg.AgentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Agent service unavailable", nil)
		return
	}

	resp, err := cfg.AgentClient.ListPlayers(ctx, &agentpb.ListPlayersRequest{
		AgentId: agentID,
		Page:    page,
		Limit:   limit,
		Search:  search,
		Status:  status,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"players":  resp.Players,
		"total":    resp.Total,
		"page":     resp.Page,
		"agent_id": agentID,
	})
}

func (cfg *RouterConfig) GetAgentPlayer(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	playerID := c.Param("id")

	if cfg.AgentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Agent service unavailable", nil)
		return
	}

	resp, err := cfg.AgentClient.GetPlayer(ctx, &agentpb.GetPlayerRequest{
		AgentId:  agentID,
		PlayerId: playerID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Player not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"player":   resp.Player,
		"agent_id": agentID,
	})
}

func (cfg *RouterConfig) UpdatePlayerLimit(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	playerID := c.Param("id")

	var req struct {
		DepositLimit float64 `json:"depositLimit"`
		BetLimit     float64 `json:"betLimit"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.AgentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Agent service unavailable", nil)
		return
	}

	_, err := cfg.AgentClient.UpdatePlayerLimit(ctx, &agentpb.UpdatePlayerLimitRequest{
		AgentId:      agentID,
		PlayerId:     playerID,
		DepositLimit: req.DepositLimit,
		BetLimit:     req.BetLimit,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"agent_id":  agentID,
		"player_id": playerID,
		"message":   "Player limit updated successfully",
	})
}

func (cfg *RouterConfig) GetCommissions(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetAgentCommissions(ctx, &commissionpb.GetAgentCommissionsRequest{
		AgentId:   agentID,
		StartDate: startDate,
		EndDate:   endDate,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"commissions": resp.Commissions,
		"total":       resp.Total,
		"agent_id":    agentID,
	})
}

func (cfg *RouterConfig) GetPendingCommissions(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetPendingCommissions(ctx, &commissionpb.GetPendingCommissionsRequest{
		AgentId: agentID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"pending":  resp.Pending,
		"agent_id": agentID,
	})
}

func (cfg *RouterConfig) GetCommissionHistory(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.GetCommissionHistory(ctx, &commissionpb.GetCommissionHistoryRequest{
		AgentId: agentID,
		Page:    page,
		Limit:   limit,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"history":  resp.History,
		"total":    resp.Total,
		"agent_id": agentID,
	})
}

func (cfg *RouterConfig) ClaimCommission(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")

	var req struct {
		CommissionID string  `json:"commissionId"`
		Amount       float64 `json:"amount"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.CommissionClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Commission service unavailable", nil)
		return
	}

	resp, err := cfg.CommissionClient.ClaimCommission(ctx, &commissionpb.ClaimCommissionRequest{
		AgentId:      agentID,
		CommissionId: req.CommissionID,
		Amount:       req.Amount,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Commission claimed successfully",
		"transaction_id": resp.TransactionId,
		"amount":         resp.Amount,
	})
}

func (cfg *RouterConfig) GetDashboard(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")

	if cfg.AgentClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Agent service unavailable", nil)
		return
	}

	resp, err := cfg.AgentClient.GetDashboard(ctx, &agentpb.GetDashboardRequest{
		AgentId: agentID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"total_players":      resp.TotalPlayers,
		"active_players":     resp.ActivePlayers,
		"total_commission":   resp.TotalCommission,
		"pending_commission": resp.PendingCommission,
		"agent_id":           agentID,
	})
}

func (cfg *RouterConfig) TrackClick(ctx context.Context, c *app.RequestContext) {
	code := c.Query("code")
	ipAddress := c.ClientIP()
	userAgent := string(c.Request.Header.UserAgent())

	if cfg.AffiliateClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Affiliate service unavailable", nil)
		return
	}

	resp, err := cfg.AffiliateClient.TrackClick(ctx, &affiliatepb.TrackClickRequest{
		AffiliateCode: code,
		IpAddress:     ipAddress,
		UserAgent:     userAgent,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"code":     code,
		"message":  "Click tracked successfully",
		"click_id": resp.ClickId,
	})
}

func (cfg *RouterConfig) RedirectToRegistration(ctx context.Context, c *app.RequestContext) {
	code := c.Param("code")
	// Redirect to registration page with affiliate code
	c.Redirect(302, []byte("https://example.com/register?ref="+code))
}

func (cfg *RouterConfig) GetPerformanceReports(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if cfg.AffiliateClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Affiliate service unavailable", nil)
		return
	}

	resp, err := cfg.AffiliateClient.GetPerformanceReport(ctx, &affiliatepb.GetPerformanceReportRequest{
		AffiliateId: agentID,
		StartDate:   startDate,
		EndDate:     endDate,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"total_clicks":  resp.TotalClicks,
		"total_signups": resp.TotalSignups,
		"conversion":    resp.Conversion,
		"revenue":       resp.Revenue,
		"affiliate_id":  agentID,
	})
}

func (cfg *RouterConfig) GetClickReports(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if cfg.AffiliateClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Affiliate service unavailable", nil)
		return
	}

	resp, err := cfg.AffiliateClient.GetClickReports(ctx, &affiliatepb.GetClickReportsRequest{
		AffiliateId: agentID,
		StartDate:   startDate,
		EndDate:     endDate,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"clicks":       resp.Clicks,
		"total":        resp.Total,
		"affiliate_id": agentID,
	})
}

func (cfg *RouterConfig) GetConversionReports(ctx context.Context, c *app.RequestContext) {
	agentID := c.GetString("user_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if cfg.AffiliateClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Affiliate service unavailable", nil)
		return
	}

	resp, err := cfg.AffiliateClient.GetConversionReports(ctx, &affiliatepb.GetConversionReportsRequest{
		AffiliateId: agentID,
		StartDate:   startDate,
		EndDate:     endDate,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"conversions":  resp.Conversions,
		"total":        resp.Total,
		"affiliate_id": agentID,
	})
}
