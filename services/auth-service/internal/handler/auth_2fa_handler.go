package handler

import (
	"context"
	"time"

	"github.com/game_engine/auth-service/internal/model"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Enable2FA handles 2FA enablement
func (h *AuthHandler) Enable2FA(ctx context.Context, req *Enable2FARequest) (*Enable2FAResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	user, err := h.authService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Verify password
	if !h.authService.CheckPassword(req.Password, user.PasswordHash) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}

	// Generate TOTP secret
	secret, qrCodeURL, err := h.authService.GenerateTOTPSecret(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate TOTP secret: %v", err)
	}

	// Generate backup codes
	backupCodes, err := h.authService.GenerateBackupCodes(10)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate backup codes: %v", err)
	}

	// Store secret temporarily (not enabled until verified)
	if err := h.authService.Store2FASecret(ctx, userID, secret); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store 2FA secret: %v", err)
	}

	return &Enable2FAResponse{
		Secret:      secret,
		QrCodeUrl:   qrCodeURL,
		BackupCodes: backupCodes,
	}, nil
}

// Verify2FA handles 2FA verification
func (h *AuthHandler) Verify2FA(ctx context.Context, req *Verify2FARequest) (*Verify2FAResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	user, err := h.authService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	var valid bool
	if req.IsBackupCode {
		valid = h.authService.ValidateBackupCode(ctx, userID, req.Code)
	} else {
		valid = h.authService.ValidateTOTP(user.TwoFactorSecret, req.Code)
	}

	if !valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid 2FA code")
	}

	// Enable 2FA
	if err := h.authService.Enable2FAForUser(ctx, userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to enable 2FA: %v", err)
	}

	// Generate full access token
	sessionID := uuid.New()
	accessToken, expiresAt, err := h.authService.GenerateAccessToken(user.ID, sessionID, []model.UserRole{model.RolePlayer})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	deviceID := "" // Would get from request
	if err := h.authService.StoreRefreshToken(ctx, user.ID, sessionID, refreshToken, deviceID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store refresh token: %v", err)
	}

	// Publish 2FA enabled event
	h.authService.PublishEvent("player.events.2fa_enabled", map[string]interface{}{
		"user_id":   user.ID.String(),
		"timestamp": time.Now().Unix(),
	})

	return &Verify2FAResponse{
		Success:      true,
		Message:      "2FA enabled successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    timestamppb.New(expiresAt),
	}, nil
}

// Disable2FA handles 2FA disable
func (h *AuthHandler) Disable2FA(ctx context.Context, req *Disable2FARequest) (*emptypb.Empty, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	user, err := h.authService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Verify password
	if !h.authService.CheckPassword(req.Password, user.PasswordHash) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}

	// Disable 2FA
	if err := h.authService.Disable2FAForUser(ctx, userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to disable 2FA: %v", err)
	}

	// Publish 2FA disabled event
	h.authService.PublishEvent("player.events.2fa_disabled", map[string]interface{}{
		"user_id":   user.ID.String(),
		"timestamp": time.Now().Unix(),
	})

	return &emptypb.Empty{}, nil
}
