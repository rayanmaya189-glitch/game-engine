package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/game_engine/auth-service/internal/model"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	// Get refresh token data
	data, err := h.authService.GetRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get refresh token: %v", err)
	}
	if data == nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid refresh token")
	}

	userID, err := uuid.Parse(data.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid user ID in token")
	}

	sessionID, err := uuid.Parse(data.SessionID)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid session ID in token")
	}

	// Delete old refresh token (token rotation)
	if err := h.authService.DeleteRefreshToken(ctx, req.RefreshToken); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete old token: %v", err)
	}

	// Generate new tokens
	accessToken, expiresAt, err := h.authService.GenerateAccessToken(userID, sessionID, []model.UserRole{model.RolePlayer})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	newRefreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	// Store new refresh token
	deviceID := getDeviceID(req.DeviceInfo)
	if err := h.authService.StoreRefreshToken(ctx, userID, sessionID, newRefreshToken, deviceID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store refresh token: %v", err)
	}

	return &RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresAt:    timestamppb.New(expiresAt),
		SessionId:    sessionID.String(),
	}, nil
}

// Logout handles user logout
func (h *AuthHandler) Logout(ctx context.Context, req *LogoutRequest) (*emptypb.Empty, error) {
	sessionID, err := uuid.Parse(req.SessionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid session ID")
	}

	if req.AllSessions {
		// Delete all user sessions
		userID, err := h.authService.GetUserIDFromSession(ctx, sessionID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get user ID: %v", err)
		}
		if err := h.authService.DeleteAllUserSessions(ctx, userID); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to delete sessions: %v", err)
		}
	} else {
		// Delete specific session
		if err := h.authService.DeleteSession(ctx, sessionID); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to delete session: %v", err)
		}
	}

	// Publish logout event
	data, _ := json.Marshal(map[string]interface{}{
		"session_id": sessionID.String(),
		"timestamp":  time.Now().Unix(),
	})
	h.authService.PublishEvent("player.events.logged_out", data)

	return &emptypb.Empty{}, nil
}

// ValidateToken handles token validation
func (h *AuthHandler) ValidateToken(ctx context.Context, req *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	claims, err := h.authService.ValidateAccessToken(req.Token)
	if err != nil {
		return &ValidateTokenResponse{
			Valid: false,
		}, nil
	}

	return &ValidateTokenResponse{
		Valid:     true,
		UserId:    claims.UserID.String(),
		SessionId: claims.SessionID.String(),
		TokenType: claims.TokenType,
		ExpiresAt: timestamppb.New(time.Now().Add(15 * time.Minute)),
		Roles:     convertRoles(claims.Roles),
	}, nil
}
