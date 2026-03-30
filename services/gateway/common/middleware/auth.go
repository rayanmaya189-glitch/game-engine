package middleware

import (
	"context"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type AuthConfig struct {
	PublicKey         *rsa.PublicKey
	PrivateKey        *rsa.PrivateKey
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
				"code":  "E1004",
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
				"code":  "E2003",
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

// validateToken validates the JWT token using RS256
func (m *AuthMiddleware) validateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.config.PublicKey, nil
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

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(m.config.PrivateKey)
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

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(m.config.PrivateKey)
}

// BlacklistToken adds a token to the blacklist
func (m *AuthMiddleware) BlacklistToken(tokenString string) error {
	if m.config.RedisClient != nil {
		blacklistKey := fmt.Sprintf("blacklist:token:%s", tokenString)
		return m.config.RedisClient.Set(context.Background(), blacklistKey, "1", m.config.TokenBlacklistTTL).Err()
	}
	return nil
}
