package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/router"

	merchantpb "github.com/game_engine/gen/go/game_engine/merchant/v1"

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
	MerchantClient        *client.MerchantClient
	UserClient            *client.UserClient
	WalletClient          *client.WalletClient
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

	// Merchant routes with API Key authentication
	merchant := r.Group("/api/v1/merchant")
	merchant.Use(cfg.AuthMiddleware.APIKeyValidation())

	// Player data
	merchant.GET("/players", cfg.ListMerchantPlayers)
	merchant.GET("/players/:id", cfg.GetMerchantPlayer)

	// Reports
	merchant.GET("/reports/revenue", cfg.GetRevenueReports)
	merchant.GET("/reports/players", cfg.GetPlayerReports)
	merchant.GET("/reports/games", cfg.GetGameReports)

	// Config
	merchant.GET("/config", cfg.GetMerchantConfig)
	merchant.PUT("/config", cfg.UpdateMerchantConfig)

	// Webhooks
	merchant.POST("/webhooks/register", cfg.RegisterWebhook)
	merchant.GET("/webhooks", cfg.ListWebhooks)
	merchant.DELETE("/webhooks/:id", cfg.DeleteWebhook)

	// Sub-agents
	merchant.GET("/agents", cfg.ListSubAgents)
	merchant.GET("/agents/:id", cfg.GetSubAgent)
	merchant.POST("/agents", cfg.CreateSubAgent)
	merchant.PUT("/agents/:id", cfg.UpdateSubAgent)
	merchant.PUT("/agents/:id/status", cfg.UpdateSubAgentStatus)

	return r
}

func (cfg *RouterConfig) ListMerchantPlayers(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")
	search := c.Query("search")
	status := c.Query("status")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.ListPlayers(ctx, &merchantpb.ListPlayersRequest{
		MerchantId: merchantID,
		Page:       page,
		Limit:      limit,
		Search:     search,
		Status:     status,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"players":     resp.Players,
		"total":       resp.Total,
		"page":        resp.Page,
		"merchant_id": merchantID,
	})
}

func (cfg *RouterConfig) GetMerchantPlayer(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	playerID := c.Param("id")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.GetPlayer(ctx, &merchantpb.GetPlayerRequest{
		MerchantId: merchantID,
		PlayerId:   playerID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Player not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"player":      resp.Player,
		"merchant_id": merchantID,
	})
}

func (cfg *RouterConfig) GetRevenueReports(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	groupBy := c.DefaultQuery("group_by", "day")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.GetRevenueReport(ctx, &merchantpb.GetRevenueReportRequest{
		MerchantId: merchantID,
		StartDate:  startDate,
		EndDate:    endDate,
		GroupBy:    groupBy,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"revenue":    resp.Revenue,
		"currency":   resp.Currency,
		"start_date": startDate,
		"end_date":   endDate,
		"breakdown":  resp.Breakdown,
	})
}

func (cfg *RouterConfig) GetPlayerReports(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.GetPlayerReport(ctx, &merchantpb.GetPlayerReportRequest{
		MerchantId: merchantID,
		StartDate:  startDate,
		EndDate:    endDate,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"total_players":  resp.TotalPlayers,
		"active_players": resp.ActivePlayers,
		"new_players":    resp.NewPlayers,
		"merchant_id":    merchantID,
	})
}

func (cfg *RouterConfig) GetGameReports(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.GetGameReport(ctx, &merchantpb.GetGameReportRequest{
		MerchantId: merchantID,
		StartDate:  startDate,
		EndDate:    endDate,
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

func (cfg *RouterConfig) GetMerchantConfig(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.GetConfig(ctx, &merchantpb.GetConfigRequest{
		MerchantId: merchantID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Config not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"config":      resp.Config,
		"merchant_id": merchantID,
	})
}

func (cfg *RouterConfig) UpdateMerchantConfig(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")

	var req struct {
		CommissionRate float64                `json:"commissionRate"`
		Theme          string                 `json:"theme"`
		Settings       map[string]interface{} `json:"settings"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	_, err := cfg.MerchantClient.UpdateConfig(ctx, &merchantpb.UpdateConfigRequest{
		MerchantId:     merchantID,
		CommissionRate: req.CommissionRate,
		Theme:          req.Theme,
		Settings:       req.Settings,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id": merchantID,
		"message":     "Configuration updated successfully",
	})
}

func (cfg *RouterConfig) RegisterWebhook(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")

	var req struct {
		URL    string   `json:"url"`
		Events []string `json:"events"`
		Secret string   `json:"secret"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.RegisterWebhook(ctx, &merchantpb.RegisterWebhookRequest{
		MerchantId: merchantID,
		Url:        req.URL,
		Events:     req.Events,
		Secret:     req.Secret,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id": merchantID,
		"webhook_id":  resp.WebhookId,
		"message":     "Webhook registered successfully",
	})
}

func (cfg *RouterConfig) ListWebhooks(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.ListWebhooks(ctx, &merchantpb.ListWebhooksRequest{
		MerchantId: merchantID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id": merchantID,
		"webhooks":    resp.Webhooks,
	})
}

func (cfg *RouterConfig) DeleteWebhook(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	webhookID := c.Param("id")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	_, err := cfg.MerchantClient.DeleteWebhook(ctx, &merchantpb.DeleteWebhookRequest{
		MerchantId: merchantID,
		WebhookId:  webhookID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id": merchantID,
		"webhook_id":  webhookID,
		"message":     "Webhook deleted successfully",
	})
}

func (cfg *RouterConfig) ListSubAgents(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.ListAgents(ctx, &merchantpb.ListAgentsRequest{
		MerchantId: merchantID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"agents":      resp.Agents,
		"total":       resp.Total,
		"merchant_id": merchantID,
	})
}

func (cfg *RouterConfig) GetSubAgent(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	agentID := c.Param("id")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.GetAgent(ctx, &merchantpb.GetAgentRequest{
		MerchantId: merchantID,
		AgentId:    agentID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "Agent not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"agent":       resp.Agent,
		"merchant_id": merchantID,
	})
}

func (cfg *RouterConfig) CreateSubAgent(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")

	var req struct {
		Username       string  `json:"username"`
		Email          string  `json:"email"`
		Password       string  `json:"password"`
		FullName       string  `json:"fullName"`
		Phone          string  `json:"phone"`
		CommissionRate float64 `json:"commissionRate"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.CreateAgent(ctx, &merchantpb.CreateAgentRequest{
		MerchantId:     merchantID,
		Username:       req.Username,
		Email:          req.Email,
		Password:       req.Password,
		FullName:       req.FullName,
		Phone:          req.Phone,
		CommissionRate: req.CommissionRate,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"agent_id":    resp.AgentId,
		"merchant_id": merchantID,
		"message":     "Agent created successfully",
	})
}

func (cfg *RouterConfig) UpdateSubAgent(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	agentID := c.Param("id")

	var req struct {
		Email          string  `json:"email"`
		FullName       string  `json:"fullName"`
		Phone          string  `json:"phone"`
		CommissionRate float64 `json:"commissionRate"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	_, err := cfg.MerchantClient.UpdateAgent(ctx, &merchantpb.UpdateAgentRequest{
		MerchantId:     merchantID,
		AgentId:        agentID,
		Email:          req.Email,
		FullName:       req.FullName,
		Phone:          req.Phone,
		CommissionRate: req.CommissionRate,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id": merchantID,
		"agent_id":    agentID,
		"message":     "Agent updated successfully",
	})
}

func (cfg *RouterConfig) UpdateSubAgentStatus(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	agentID := c.Param("id")

	var req struct {
		Status string `json:"status"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	_, err := cfg.MerchantClient.UpdateAgentStatus(ctx, &merchantpb.UpdateAgentStatusRequest{
		MerchantId: merchantID,
		AgentId:    agentID,
		Status:     req.Status,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"merchant_id": merchantID,
		"agent_id":    agentID,
		"status":      req.Status,
		"message":     "Agent status updated successfully",
	})
}
