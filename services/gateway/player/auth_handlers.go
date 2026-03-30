package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	authpb "github.com/game_engine/gen/go/game_engine/auth/v1"

	"common/handler"
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
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		FullName: req.FullName,
		Phone:    req.Phone,
		Currency: req.Currency,
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
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Invalid credentials", nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
		"expires_in":    resp.ExpiresIn,
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

	handler.SendSuccess(c, map[string]interface{}{
		"access_token":  resp.AccessToken,
		"refresh_token": resp.RefreshToken,
		"expires_in":    resp.ExpiresIn,
		"token_type":    "Bearer",
	})
}

// Logout handles user logout
func (cfg *RouterConfig) Logout(ctx context.Context, c *app.RequestContext) {
	if cfg.AuthClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Auth service unavailable", nil)
		return
	}

	token := c.GetHeader("Authorization")
	if token != "" {
		cfg.AuthClient.Logout(ctx, &authpb.LogoutRequest{
			Token: token,
		})
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message": "Logout successful",
	})
}
