package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	agentpb "github.com/game_engine/gen/go/game_engine/agent/v1"

	"common/handler"
)

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
