package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ResetPassword handles password reset request
func (h *AuthHandler) ResetPassword(ctx context.Context, req *ResetPasswordRequest) (*emptypb.Empty, error) {
	user, err := h.authService.GetUserByIdentifier(ctx, req.Identifier)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find user: %v", err)
	}
	if user == nil {
		// Don't reveal if user exists
		return &emptypb.Empty{}, nil
	}

	token, err := h.authService.GeneratePasswordResetToken(ctx, user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate reset token: %v", err)
	}

	_ = token // Would be sent via email in production

	return &emptypb.Empty{}, nil
}

// ConfirmResetPassword handles password reset confirmation
func (h *AuthHandler) ConfirmResetPassword(ctx context.Context, req *ConfirmResetPasswordRequest) (*emptypb.Empty, error) {
	// Validate password
	if err := h.authService.ValidatePassword(req.NewPassword); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid password: %v", err)
	}

	// Check password match
	if req.NewPassword != req.ConfirmPassword {
		return nil, status.Errorf(codes.InvalidArgument, "passwords do not match")
	}

	// Validate token
	userID, err := h.authService.ValidatePasswordResetToken(ctx, req.Token)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid or expired token")
	}

	// Hash new password
	hash, err := h.authService.HashPassword(req.NewPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	// Update password
	if err := h.authService.UpdatePassword(ctx, *userID, hash); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update password: %v", err)
	}

	// Invalidate all sessions
	if err := h.authService.DeleteAllUserSessions(ctx, *userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to invalidate sessions: %v", err)
	}

	// Publish password reset event
	data, _ := json.Marshal(map[string]interface{}{
		"user_id":   userID.String(),
		"timestamp": time.Now().Unix(),
	})
	h.authService.PublishEvent("player.events.password_reset", data)

	return &emptypb.Empty{}, nil
}

// ChangePassword handles password change
func (h *AuthHandler) ChangePassword(ctx context.Context, req *ChangePasswordRequest) (*emptypb.Empty, error) {
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

	// Verify current password
	if !h.authService.CheckPassword(req.CurrentPassword, user.PasswordHash) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid current password")
	}

	// Validate new password
	if err := h.authService.ValidatePassword(req.NewPassword); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid new password: %v", err)
	}

	// Check password match
	if req.NewPassword != req.ConfirmPassword {
		return nil, status.Errorf(codes.InvalidArgument, "passwords do not match")
	}

	// Hash new password
	hash, err := h.authService.HashPassword(req.NewPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	// Update password
	if err := h.authService.UpdatePassword(ctx, userID, hash); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update password: %v", err)
	}

	// Invalidate all sessions
	if err := h.authService.DeleteAllUserSessions(ctx, userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to invalidate sessions: %v", err)
	}

	return &emptypb.Empty{}, nil
}
