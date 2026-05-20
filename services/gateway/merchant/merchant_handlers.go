package main

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	merchantpb "github.com/game_engine/common-service/proto/gen/go/merchant/v1"

	"github.com/game_engine/gateway/common/handler"
)

func (cfg *RouterConfig) ListMerchantPlayers(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "20")
	page, _ := strconv.ParseInt(pageStr, 10, 32)
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	search := c.Query("search")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.ListPlayers(ctx, &merchantpb.ListPlayersRequest{
		MerchantId: merchantID,
		Page:       int32(page),
		Limit:      int32(limit),
		Search:     search,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"players":     resp.Players,
		"total":       resp.Total,
		"page":        page,
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

func (cfg *RouterConfig) ListSubAgents(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "20")
	page, _ := strconv.ParseInt(pageStr, 10, 32)
	limit, _ := strconv.ParseInt(limitStr, 10, 32)

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.ListAgents(ctx, &merchantpb.ListAgentsRequest{
		MerchantId: merchantID,
		Page:       int32(page),
		Limit:      int32(limit),
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"agents":      resp.Agents,
		"total":       resp.Total,
		"page":        page,
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
		SendInvitation bool    `json:"sendInvitation"`
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
		SendInvitation: req.SendInvitation,
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
		Username string `json:"username"`
		Email    string `json:"email"`
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
		MerchantId: merchantID,
		AgentId:    agentID,
		Username:   req.Username,
		Email:      req.Email,
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
