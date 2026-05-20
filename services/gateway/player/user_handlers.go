package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	userpb "github.com/game_engine/common-service/proto/gen/go/user/v1"

	"github.com/game_engine/gateway/common/handler"
)

// GetProfile handles getting user profile
func (cfg *RouterConfig) GetProfile(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.UserClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "User service unavailable", nil)
		return
	}

	resp, err := cfg.UserClient.GetProfile(ctx, &userpb.GetProfileRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 404, handler.ErrCodeNotFound, "User not found", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"user_id":    resp.Profile.UserId,
		"username":   resp.Profile.Username,
		"email":      resp.Profile.Email,
		"first_name": resp.Profile.FirstName,
		"last_name":  resp.Profile.LastName,
		"phone":      resp.Profile.Phone,
		"status":     resp.Profile.Status,
		"kyc_level":  resp.Profile.KycLevel,
		"created_at": resp.Profile.CreatedAt,
	})
}

// UpdateProfile handles updating user profile
func (cfg *RouterConfig) UpdateProfile(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.UserClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "User service unavailable", nil)
		return
	}

	var req struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	_, err := cfg.UserClient.UpdateProfile(ctx, &userpb.UpdateProfileRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Profile updated successfully",
	})
}
