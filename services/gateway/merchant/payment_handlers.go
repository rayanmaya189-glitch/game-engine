package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	merchantpb "github.com/game_engine/common-service/proto/gen/go/merchant/v1"

	"common/handler"
)

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
