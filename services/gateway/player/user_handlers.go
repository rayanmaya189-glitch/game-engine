package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	userpb "github.com/game_engine/gen/go/game_engine/user/v1"

	"common/handler"
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
		"user_id":    resp.User.Id,
		"username":   resp.User.Username,
		"email":      resp.User.Email,
		"full_name":  resp.User.FullName,
		"phone":      resp.User.Phone,
		"status":     resp.User.Status,
		"kyc_status": resp.User.KycStatus,
		"created_at": resp.User.CreatedAt,
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
		FullName string `json:"fullName"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	_, err := cfg.UserClient.UpdateProfile(ctx, &userpb.UpdateProfileRequest{
		UserId:   userID,
		FullName: req.FullName,
		Phone:    req.Phone,
		Email:    req.Email,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Profile updated successfully",
	})
}
