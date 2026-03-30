package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// isPublicRoute checks if the route is public (no auth required)
func isPublicRoute(path string) bool {
	publicRoutes := []string{
		"/health",
		"/ready",
		"/api/v1/auth/register",
		"/api/v1/auth/login",
		"/api/v1/auth/refresh",
		"/api/v1/auth/verify-email",
		"/api/v1/auth/reset-password",
		"/api/v1/affiliate/tracking",
	}

	for _, route := range publicRoutes {
		if path == route || path == route+"/" {
			return true
		}
	}
	return false
}

// APIKeyValidation validates API keys for merchant/agent gateways
func (m *AuthMiddleware) APIKeyValidation() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		apiKey := string(ctx.Request.Header.Get("X-API-Key"))
		if apiKey == "" {
			ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"error": "missing API key",
				"code":  "API_KEY_REQUIRED",
			})
			ctx.Abort()
			return
		}

		// Validate API key against database/redis
		// For now, we'll do a simple validation
		merchantID, valid := m.validateAPIKey(apiKey)
		if !valid {
			ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"error": "invalid API key",
				"code":  "API_KEY_INVALID",
			})
			ctx.Abort()
			return
		}

		ctx.Set("merchant_id", merchantID)
		ctx.Set("api_key", apiKey)
		ctx.Next(c)
	}
}

// validateAPIKey validates the API key
func (m *AuthMiddleware) validateAPIKey(apiKey string) (string, bool) {
	if m.config.RedisClient != nil {
		key := fmt.Sprintf("apikey:%s", apiKey)
		merchantID, err := m.config.RedisClient.Get(context.Background(), key).Result()
		if err == nil && merchantID != "" {
			return merchantID, true
		}
	}
	// Fallback: allow any key starting with "test_" for development
	if len(apiKey) > 5 && apiKey[:5] == "test_" {
		return "test-merchant", true
	}
	return "", false
}

// MFACheck validates MFA token
func (m *AuthMiddleware) MFACheck() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// Check if user has MFA enabled
		userID := ctx.GetString("user_id")
		if userID == "" {
			ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"error": "user not authenticated",
				"code":  "UNAUTHORIZED",
			})
			ctx.Abort()
			return
		}

		// Check MFA status in Redis
		if m.config.RedisClient != nil {
			mfaKey := fmt.Sprintf("mfa:enabled:%s", userID)
			mfaEnabled, err := m.config.RedisClient.Get(context.Background(), mfaKey).Bool()
			if err == nil && mfaEnabled {
				// Check if MFA code is provided
				mfaCode := string(ctx.Request.Header.Get("X-MFA-Code"))
				if mfaCode == "" {
					ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
						"error": "MFA code required",
						"code":  "MFA_REQUIRED",
					})
					ctx.Abort()
					return
				}

				// Validate MFA code
				valid := m.validateMFACode(userID, mfaCode)
				if !valid {
					ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
						"error": "invalid MFA code",
						"code":  "MFA_INVALID",
					})
					ctx.Abort()
					return
				}
			}
		}

		ctx.Next(c)
	}
}

// validateMFACode validates the MFA code
func (m *AuthMiddleware) validateMFACode(userID, code string) bool {
	if m.config.RedisClient != nil {
		// In production, validate against TOTP or backup codes
		// For now, accept any 6-digit code
		if len(code) == 6 {
			return true
		}
	}
	return false
}

// RoleCheck checks if user has required role
func (m *AuthMiddleware) RoleCheck(allowedRoles ...string) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		role := ctx.GetString("role")

		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				ctx.Next(c)
				return
			}
		}

		ctx.JSON(consts.StatusForbidden, map[string]interface{}{
			"error": "insufficient permissions",
			"code":  "FORBIDDEN",
		})
		ctx.Abort()
	}
}

// IPWhitelistCheck checks if client IP is whitelisted
func (m *AuthMiddleware) IPWhitelistCheck(allowedIPs []string) app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		clientIP := string(ctx.Request.Header.Get("X-Forwarded-For"))
		if clientIP == "" {
			clientIP = ctx.RemoteAddr().String()
		}

		// Check if IP is in whitelist
		for _, allowedIP := range allowedIPs {
			if clientIP == allowedIP {
				ctx.Next(c)
				return
			}
		}

		ctx.JSON(consts.StatusForbidden, map[string]interface{}{
			"error": "IP not whitelisted",
			"code":  "IP_NOT_ALLOWED",
		})
		ctx.Abort()
	}
}

// GetConfigJSON converts config to JSON for middleware
func (m *AuthMiddleware) GetConfigJSON() string {
	config := map[string]interface{}{
		"jwt_expiration":     m.config.JWTExpiration.String(),
		"refresh_expiration": m.config.RefreshExpiration.String(),
	}
	b, _ := json.Marshal(config)
	return string(b)
}
