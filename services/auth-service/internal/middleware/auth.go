package middleware

import (
	"context"
	"strings"

	"github.com/game_engine/auth-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthMiddleware provides authentication middleware
type AuthMiddleware struct {
	authService *service.AuthService
}

// NewAuthMiddleware creates a new AuthMiddleware
func NewAuthMiddleware(authService *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}
}

// AuthInterceptor creates a unary interceptor for authentication
func (m *AuthMiddleware) AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Skip auth for certain methods
		if shouldSkipAuth(info.FullMethod) {
			return handler(ctx, req)
		}

		// Get token from metadata
		token, err := extractToken(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "missing authentication token")
		}

		// Validate token
		claims, err := m.authService.ValidateAccessToken(token)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		// Add claims to context
		ctx = context.WithValue(ctx, "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "session_id", claims.SessionID)
		ctx = context.WithValue(ctx, "roles", claims.Roles)

		return handler(ctx, req)
	}
}

// extractToken extracts token from gRPC metadata
func extractToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "missing metadata")
	}

	authHeader := md.Get("authorization")
	if len(authHeader) == 0 {
		return "", status.Error(codes.Unauthenticated, "missing authorization header")
	}

	// Extract token from "Bearer <token>"
	parts := strings.SplitN(authHeader[0], " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", status.Error(codes.Unauthenticated, "invalid authorization header format")
	}

	return parts[1], nil
}

// shouldSkipAuth checks if the method should skip authentication
func shouldSkipAuth(method string) bool {
	// Public methods that don't require authentication
	publicMethods := map[string]bool{
		"/game_engine.auth.v1.AuthService/Register":             true,
		"/game_engine.auth.v1.AuthService/Login":                true,
		"/game_engine.auth.v1.AuthService/ResetPassword":        true,
		"/game_engine.auth.v1.AuthService/ConfirmResetPassword": true,
		"/game_engine.auth.v1.AuthService/ValidateToken":        true,
		"/game_engine.auth.v1.AuthService/VerifyEmail":          true,
		"/game_engine.auth.v1.AuthService/VerifyPhone":          true,
		"/game_engine.auth.v1.AuthService/Enable2FA":            true,
		"/game_engine.auth.v1.AuthService/Verify2FA":            true,
	}

	return publicMethods[method]
}

// GetUserID extracts user ID from context
func GetUserID(ctx context.Context) string {
	if userID, ok := ctx.Value("user_id").([]uint8); ok {
		return string(userID)
	}
	return ""
}

// GetSessionID extracts session ID from context
func GetSessionID(ctx context.Context) string {
	if sessionID, ok := ctx.Value("session_id").([]uint8); ok {
		return string(sessionID)
	}
	return ""
}

// GetRoles extracts roles from context
func GetRoles(ctx context.Context) []string {
	if roles, ok := ctx.Value("roles").([]string); ok {
		return roles
	}
	return nil
}

// RoleCheck checks if user has required role
func RoleCheck(ctx context.Context, requiredRole string) bool {
	roles := GetRoles(ctx)
	for _, role := range roles {
		if role == requiredRole {
			return true
		}
	}
	return false
}
