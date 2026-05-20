package main

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"

	authpb "github.com/game_engine/common-service/proto/gen/go/auth/v1"

	"github.com/game_engine/gateway/common/handler"
)

// Register handles user registration
func (cfg *RouterConfig) Register(ctx context.Context, c *app.RequestContext) {
	if cfg.AuthClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Auth service unavailable", nil)
		return
	}

	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"fullName"`
		Phone    string `json:"phone"`
		Currency string `json:"currency"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	resp, err := cfg.AuthClient.Register(ctx, &authpb.RegisterRequest{
		Identifier:      req.Email,
		Password:        req.Password,
		ConfirmPassword: req.Password,
		Currency:        req.Currency,
		AcceptTerms:     true,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"user_id": resp.UserId,
		"message": "Registration successful",
	})
}

// Login handles user login
func (cfg *RouterConfig) Login(ctx context.Context, c *app.RequestContext) {
	if cfg.AuthClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Auth service unavailable", nil)
		return
	}

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	resp, err := cfg.AuthClient.Login(ctx, &authpb.LoginRequest{
		Identifier: req.Username,
		Password:   req.Password,
	})

	if err != nil {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Invalid credentials", nil)
		return
	}

	expiresIn := int64(0)
	if resp.ExpiresAt != nil {
		expiresIn = resp.ExpiresAt.GetSeconds() - time.Now().Unix()
		if expiresIn < 0 {
			expiresIn = 0
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
		"expires_in":    expiresIn,
		"token_type":    "Bearer",
	})
}

// RefreshToken handles token refresh
func (cfg *RouterConfig) RefreshToken(ctx context.Context, c *app.RequestContext) {
	if cfg.AuthClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Auth service unavailable", nil)
		return
	}

	var req struct {
		RefreshToken string `json:"refreshToken"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	resp, err := cfg.AuthClient.RefreshToken(ctx, &authpb.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})

	if err != nil {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Invalid refresh token", nil)
		return
	}

	expiresIn := int64(0)
	if resp.ExpiresAt != nil {
		expiresIn = resp.ExpiresAt.GetSeconds() - time.Now().Unix()
		if expiresIn < 0 {
			expiresIn = 0
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
		"expires_in":    expiresIn,
		"token_type":    "Bearer",
	})
}

// Logout handles user logout
func (cfg *RouterConfig) Logout(ctx context.Context, c *app.RequestContext) {
	if cfg.AuthClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Auth service unavailable", nil)
		return
	}

	sessionID := c.GetString("session_id")
	if sessionID != "" {
		_, _ = cfg.AuthClient.Logout(ctx, &authpb.LogoutRequest{
			SessionId:   sessionID,
			AllSessions: false,
		})
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Logout successful",
	})
}
