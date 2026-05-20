package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	commissionpb "github.com/game_engine/common-service/proto/gen/go/commission/v1"

	"github.com/game_engine/gateway/common/handler"
)

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
