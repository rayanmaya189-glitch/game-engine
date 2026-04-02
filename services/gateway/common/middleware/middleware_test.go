package middleware

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"net"
	"testing"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v5"
)

func generateTestKeys(t *testing.T) (*rsa.PrivateKey, *rsa.PublicKey) {
	t.Helper()
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate RSA key: %v", err)
	}
	return privateKey, &privateKey.PublicKey
}

func createTestToken(t *testing.T, privateKey *rsa.PrivateKey, userID, username, role string, expiry time.Duration) string {
	t.Helper()
	claims := &Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "game_engine",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(privateKey)
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}
	return signed
}

type mockRequestContext struct {
	path        string
	authHeader  string
	method      string
	response    map[string]interface{}
	statusCode  int
	aborted     bool
	values      map[string]interface{}
	headers     map[string]string
}

func (m *mockRequestContext) Request() *protocol.Request  { return nil }
func (m *mockRequestContext) Response() *protocol.Response { return nil }

func TestJWTValidation_MissingAuth(t *testing.T) {
	priv, pub := generateTestKeys(t)
	_ = priv

	m := NewAuthMiddleware(&AuthConfig{
		PublicKey:     pub,
		PrivateKey:    priv,
		JWTExpiration: time.Hour,
	})

	handler := m.JWTValidation()
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
}

func TestGenerateToken(t *testing.T) {
	priv, _ := generateTestKeys(t)
	m := NewAuthMiddleware(&AuthConfig{
		PrivateKey:    priv,
		JWTExpiration: time.Hour,
	})

	token, err := m.GenerateToken("user-1", "testuser", "admin")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token == "" {
		t.Fatal("expected non-empty token")
	}
}

func TestGenerateRefreshToken(t *testing.T) {
	priv, _ := generateTestKeys(t)
	m := NewAuthMiddleware(&AuthConfig{
		PrivateKey:        priv,
		RefreshExpiration: time.Hour * 24,
	})

	token, err := m.GenerateRefreshToken("user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token == "" {
		t.Fatal("expected non-empty refresh token")
	}
}

func TestValidateToken_Valid(t *testing.T) {
	priv, pub := generateTestKeys(t)
	m := NewAuthMiddleware(&AuthConfig{
		PublicKey:     pub,
		PrivateKey:    priv,
		JWTExpiration: time.Hour,
	})

	tokenStr := createTestToken(t, priv, "u1", "alice", "admin", time.Hour)
	claims, err := m.validateToken(tokenStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if claims.UserID != "u1" {
		t.Errorf("expected u1, got %s", claims.UserID)
	}
	if claims.Username != "alice" {
		t.Errorf("expected alice, got %s", claims.Username)
	}
	if claims.Role != "admin" {
		t.Errorf("expected admin, got %s", claims.Role)
	}
}

func TestValidateToken_Expired(t *testing.T) {
	priv, pub := generateTestKeys(t)
	m := NewAuthMiddleware(&AuthConfig{
		PublicKey:     pub,
		PrivateKey:    priv,
		JWTExpiration: time.Hour,
	})

	tokenStr := createTestToken(t, priv, "u1", "alice", "admin", -time.Hour)
	_, err := m.validateToken(tokenStr)
	if err == nil {
		t.Fatal("expected error for expired token")
	}
}

func TestValidateToken_WrongKey(t *testing.T) {
	_, pub := generateTestKeys(t)
	otherPriv, _ := generateTestKeys(t)

	m := NewAuthMiddleware(&AuthConfig{
		PublicKey: pub,
	})

	tokenStr := createTestToken(t, otherPriv, "u1", "alice", "admin", time.Hour)
	_, err := m.validateToken(tokenStr)
	if err == nil {
		t.Fatal("expected error for token signed with different key")
	}
}

func TestRateLimiterMiddleware_Creation(t *testing.T) {
	m := NewRateLimiterMiddleware(&RateLimiterConfig{
		RequestsPerMinute: 60,
		BurstSize:         10,
		KeyPrefix:         "test",
	})
	if m == nil {
		t.Fatal("expected non-nil middleware")
	}
}

func TestRateLimiterMiddleware_Defaults(t *testing.T) {
	m := NewRateLimiterMiddleware(&RateLimiterConfig{})
	if m.config.RequestsPerMinute != 100 {
		t.Errorf("expected default 100, got %d", m.config.RequestsPerMinute)
	}
	if m.config.KeyPrefix != "ratelimit" {
		t.Errorf("expected default prefix, got %s", m.config.KeyPrefix)
	}
}

func TestRateLimiterHandler_Creation(t *testing.T) {
	m := NewRateLimiterMiddleware(&RateLimiterConfig{
		RequestsPerMinute: 100,
	})
	handler := m.RateLimiter()
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
}

func TestIPBasedRateLimiter_Creation(t *testing.T) {
	m := NewRateLimiterMiddleware(&RateLimiterConfig{})
	handler := m.IPBasedRateLimiter(50)
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
}

func TestCustomRateLimiter_Creation(t *testing.T) {
	m := NewRateLimiterMiddleware(&RateLimiterConfig{})
	handler := m.CustomRateLimiter(map[string]int{
		"/api/bet": 10,
	})
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
}

func TestMatchPath(t *testing.T) {
	tests := []struct {
		path    string
		pattern string
		want    bool
	}{
		{"/api/bet", "/api/bet", true},
		{"/api/bet/123", "/api/bet", false},
		{"/api/bet/123", "/api/bet/*", true},
		{"/api/user", "/api/bet/*", false},
		{"/health", "/health", true},
	}

	for _, tt := range tests {
		got := matchPath(tt.path, tt.pattern)
		if got != tt.want {
			t.Errorf("matchPath(%q, %q) = %v, want %v", tt.path, tt.pattern, got, tt.want)
		}
	}
}

func TestCORSConfig_Defaults(t *testing.T) {
	m := NewCORSMiddleware(nil)
	if m == nil {
		t.Fatal("expected non-nil middleware")
	}
	if len(m.config.AllowOrigins) != 1 || m.config.AllowOrigins[0] != "*" {
		t.Errorf("expected wildcard origin, got %v", m.config.AllowOrigins)
	}
}

func TestCORSMiddleware_Creation(t *testing.T) {
	m := NewCORSMiddleware(&CORSConfig{
		AllowOrigins: []string{"https://example.com"},
		AllowMethods: []string{"GET", "POST"},
	})
	if m == nil {
		t.Fatal("expected non-nil middleware")
	}
	handler := m.CORS()
	if handler == nil {
		t.Fatal("expected non-nil CORS handler")
	}
}

func TestCORSWithConfig(t *testing.T) {
	handler := CORSWithConfig(&CORSConfig{
		AllowOrigins: []string{"*"},
	})
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
}

func TestAdminCORS_Creation(t *testing.T) {
	m := NewCORSMiddleware(&CORSConfig{
		AllowOrigins: []string{"https://admin.example.com"},
	})
	handler := m.AdminCORS([]string{"https://admin.example.com"})
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
}

// suppress unused import warnings
var _ = net.IPv4
var _ = config.NewOptions
var _ = consts.StatusOK
