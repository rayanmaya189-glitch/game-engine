package handler

import (
	"context"

	"github.com/game_engine/auth-service/internal/model"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// VerifyEmail handles email verification
func (h *AuthHandler) VerifyEmail(ctx context.Context, req *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	userID, err := h.authService.ValidateEmailVerificationToken(ctx, req.Code)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid or expired token")
	}

	if err := h.authService.MarkEmailVerified(ctx, *userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify email: %v", err)
	}

	// Generate tokens
	sessionID := uuid.New()
	accessToken, _, err := h.authService.GenerateAccessToken(*userID, sessionID, []model.UserRole{model.RolePlayer})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	return &VerifyEmailResponse{
		Success:      true,
		Message:      "Email verified successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// VerifyPhone handles phone verification
func (h *AuthHandler) VerifyPhone(ctx context.Context, req *VerifyPhoneRequest) (*VerifyPhoneResponse, error) {
	// Similar to email verification
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	if err := h.authService.MarkPhoneVerified(ctx, userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify phone: %v", err)
	}

	sessionID := uuid.New()
	accessToken, _, err := h.authService.GenerateAccessToken(userID, sessionID, []model.UserRole{model.RolePlayer})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	return &VerifyPhoneResponse{
		Success:      true,
		Message:      "Phone verified successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
