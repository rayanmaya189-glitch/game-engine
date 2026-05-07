package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	merchantpb "github.com/game_engine/common-service/proto/gen/go/merchant/v1"

	"common/handler"
)

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
