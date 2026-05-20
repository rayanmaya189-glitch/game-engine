package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	merchantpb "github.com/game_engine/common-service/proto/gen/go/merchant/v1"

	"github.com/game_engine/gateway/common/handler"
)

func (cfg *RouterConfig) GetRevenueReports(ctx context.Context, c *app.RequestContext) {
	merchantID := c.GetString("merchant_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if cfg.MerchantClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Merchant service unavailable", nil)
		return
	}

	resp, err := cfg.MerchantClient.GetRevenueReport(ctx, &merchantpb.GetRevenueReportRequest{
		MerchantId: merchantID,
		StartDate:  startDate,
		EndDate:    endDate,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"total_revenue":    resp.TotalRevenue,
		"total_deposits":   resp.TotalDeposits,
		"total_withdrawals": resp.TotalWithdrawals,
		"total_players":    resp.TotalPlayers,
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
		"total_bets":   resp.TotalBets,
		"total_wins":   resp.TotalWins,
		"net_revenue":  resp.NetRevenue,
		"games_played": resp.GamesPlayed,
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
		"total_bets":    resp.TotalBets,
		"total_wins":    resp.TotalWins,
		"total_players": resp.TotalPlayers,
		"plays":         resp.Plays,
	})
}
