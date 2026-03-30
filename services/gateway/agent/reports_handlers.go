package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	affiliatepb "github.com/game_engine/gen/go/game_engine/affiliate/v1"

	"common/handler"
)

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
