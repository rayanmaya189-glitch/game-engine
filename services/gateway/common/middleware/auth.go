package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type AuthConfig struct {
	JWTSecret         string
	JWTExpiration     time.Duration
	RefreshExpiration time.Duration
	RedisClient       *redis.Client
	TokenBlacklistTTL time.Duration
}

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Merchant string `json:"merchant,omitempty"`
	Agent    string `json:"agent,omitempty"`
	jwt.RegisteredClaims
}

type AuthMiddleware struct {
	config *AuthConfig
}

func NewAuthMiddleware(config *AuthConfig) *AuthMiddleware {
	return &AuthMiddleware{config: config}
}

// JWTValidation validates JWT tokens
func (m *AuthMiddleware) JWTValidation() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// Skip auth for public routes
		path := string(ctx.Request.URI().Path())
		if isPublicRoute(path) {
			ctx.Next(c)
			return
		}

		// Get token from Authorization header
		authHeader := string(ctx.Request.Header.Authorization())
		if authHeader == "" {
			ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"error": "missing authorization header",
				"code":  "UNAUTHORIZED",
			})
			ctx.Abort()
			return
		}

		// Extract token (Bearer <token>)
		tokenString := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}

		// Validate token
		claims, err := m.validateToken(tokenString)
		if err != nil {
			ctx.JSON(consts.StatusUnauthorized, map[string]interface{}{
				"error": "invalid or expired token",
				"code":  "TOKEN_INVALID",
			})
			ctx.Abort()
			return
		}

		// Set claims in context
		ctx.Set("user_id", claims.UserID)
		ctx.Set("username", claims.Username)
		ctx.Set("role", claims.Role)
		ctx.Set("claims", claims)

		ctx.Next(c)
	}
}

// validateToken validates the JWT token
func (m *AuthMiddleware) validateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.config.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Check if token is blacklisted
		if m.config.RedisClient != nil {
			blacklistKey := fmt.Sprintf("blacklist:token:%s", tokenString)
			exists, err := m.config.RedisClient.Exists(context.Background(), blacklistKey).Result()
			if err == nil && exists > 0 {
				return nil, fmt.Errorf("token is blacklisted")
			}
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// GenerateToken generates a new JWT token
func (m *AuthMiddleware) GenerateToken(userID, username, role string) (string, error) {
	claims := &Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.config.JWTExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "game_engine",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.config.JWTSecret))
}

// GenerateRefreshToken generates a new refresh token
func (m *AuthMiddleware) GenerateRefreshToken(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.config.RefreshExpiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "game_engine",
			Subject:   "refresh",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.config.JWTSecret))
}

// BlacklistToken adds a token to the blacklist
func (m *AuthMiddleware) BlacklistToken(tokenString string) error {
	if m.config.RedisClient != nil {
		blacklistKey := fmt.Sprintf("blacklist:token:%s", tokenString)
		return m.config.RedisClient.Set(context.Background(), blacklistKey, "1", m.config.TokenBlacklistTTL).Err()
	}
	return nil
}

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

// Helper to convert config to JSON for middleware
func (m *AuthMiddleware) GetConfigJSON() string {
	config := map[string]interface{}{
		"jwt_expiration":     m.config.JWTExpiration.String(),
		"refresh_expiration": m.config.RefreshExpiration.String(),
	}
	b, _ := json.Marshal(config)
	return string(b)
}
