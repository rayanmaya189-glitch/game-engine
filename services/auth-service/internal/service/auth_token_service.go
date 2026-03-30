package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/game_engine/auth-service/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// GenerateAccessToken generates a JWT access token
func (s *AuthService) GenerateAccessToken(userID, sessionID uuid.UUID, roles []model.UserRole) (string, time.Time, error) {
	expiresAt := time.Now().Add(s.cfg.JWT.AccessTokenExpiry)

	claims := jwt.MapClaims{
		"user_id":    userID.String(),
		"session_id": sessionID.String(),
		"roles":      roles,
		"type":       "access",
		"exp":        expiresAt.Unix(),
		"iat":        time.Now().Unix(),
		"iss":        s.cfg.JWT.Issuer,
		"aud":        s.cfg.JWT.Audience,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, expiresAt, nil
}

// GenerateRefreshToken generates an opaque refresh token
func (s *AuthService) GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// ValidateAccessToken validates a JWT access token
func (s *AuthService) ValidateAccessToken(tokenString string) (*model.TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &s.privateKey.PublicKey, nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	// Check expiration
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, ErrInvalidToken
	}
	if time.Now().Unix() > int64(exp) {
		return nil, ErrTokenExpired
	}

	// Parse claims
	userID, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return nil, ErrInvalidToken
	}

	sessionID, err := uuid.Parse(claims["session_id"].(string))
	if err != nil {
		return nil, ErrInvalidToken
	}

	rolesRaw, ok := claims["roles"].([]interface{})
	if !ok {
		return nil, ErrInvalidToken
	}

	roles := make([]model.UserRole, len(rolesRaw))
	for i, r := range rolesRaw {
		roles[i] = model.UserRole(r.(string))
	}

	return &model.TokenClaims{
		UserID:    userID,
		SessionID: sessionID,
		Roles:     roles,
		TokenType: claims["type"].(string),
	}, nil
}

// StoreRefreshToken stores refresh token in Redis
func (s *AuthService) StoreRefreshToken(ctx context.Context, userID, sessionID uuid.UUID, token string, deviceID string) error {
	data := model.RefreshTokenData{
		UserID:    userID.String(),
		SessionID: sessionID.String(),
		DeviceID:  deviceID,
		ExpiresAt: time.Now().Add(s.cfg.JWT.RefreshTokenExpiry).Unix(),
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return s.redis.Set(ctx, fmt.Sprintf("refresh:%s", token), dataJSON, s.cfg.JWT.RefreshTokenExpiry).Err()
}

// GetRefreshToken retrieves refresh token data from Redis
func (s *AuthService) GetRefreshToken(ctx context.Context, token string) (*model.RefreshTokenData, error) {
	dataJSON, err := s.redis.Get(ctx, fmt.Sprintf("refresh:%s", token)).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var data model.RefreshTokenData
	if err := json.Unmarshal(dataJSON, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

// DeleteRefreshToken deletes refresh token from Redis
func (s *AuthService) DeleteRefreshToken(ctx context.Context, token string) error {
	return s.redis.Del(ctx, fmt.Sprintf("refresh:%s", token)).Err()
}

// GeneratePartialToken generates a partial token for 2FA
func (s *AuthService) GeneratePartialToken(userID, sessionID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userID.String(),
		"session_id": sessionID.String(),
		"type":       "partial",
		"exp":        time.Now().Add(5 * time.Minute).Unix(),
		"iat":        time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(s.privateKey)
}
