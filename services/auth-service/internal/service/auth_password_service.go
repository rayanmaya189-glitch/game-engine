package service

import (
	"context"
	"fmt"
	"time"

	"github.com/game_engine/auth-service/internal/model"
	"github.com/google/uuid"
)

// CheckLoginAttempts checks if user has exceeded login attempt limit
func (s *AuthService) CheckLoginAttempts(ctx context.Context, userID uuid.UUID) error {
	since := time.Now().Add(-s.cfg.RateLimiting.WindowDuration)
	count, err := s.repo.GetLoginAttempts(ctx, userID, since)
	if err != nil {
		return err
	}

	if count >= s.cfg.RateLimiting.MaxLoginAttempts {
		// Check if account is locked
		user, err := s.repo.GetUserByID(ctx, userID)
		if err != nil {
			return err
		}
		if user != nil && user.LockedUntil != nil && time.Now().Before(*user.LockedUntil) {
			return ErrAccountLocked
		}

		// Lock the account
		lockedUntil := time.Now().Add(s.cfg.RateLimiting.LockoutDuration)
		if err := s.repo.UpdateFailedLoginCount(ctx, userID, count, &lockedUntil); err != nil {
			return err
		}
		return ErrTooManyAttempts
	}
	return nil
}

// RecordLoginAttempt records a login attempt
func (s *AuthService) RecordLoginAttempt(ctx context.Context, userID uuid.UUID, ipAddress string, success bool) error {
	attempt := &model.LoginAttempt{
		ID:          uuid.New(),
		UserID:      userID,
		IPAddress:   ipAddress,
		AttemptedAt: time.Now(),
		Success:     success,
	}
	return s.repo.RecordLoginAttempt(ctx, attempt)
}

// GeneratePasswordResetToken generates a password reset token
func (s *AuthService) GeneratePasswordResetToken(ctx context.Context, userID uuid.UUID) (string, error) {
	token := uuid.New().String()
	resetToken := &model.PasswordResetToken{
		ID:        uuid.New(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(s.cfg.TokenExpiry.PasswordReset),
		Used:      false,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreatePasswordResetToken(ctx, resetToken); err != nil {
		return "", err
	}

	// Also store in Redis for fast lookup
	key := fmt.Sprintf("pwd_reset:%s", token)
	return token, s.redis.Set(ctx, key, userID.String(), s.cfg.TokenExpiry.PasswordReset).Err()
}

// ValidatePasswordResetToken validates a password reset token
func (s *AuthService) ValidatePasswordResetToken(ctx context.Context, token string) (*uuid.UUID, error) {
	// Check Redis first
	key := fmt.Sprintf("pwd_reset:%s", token)
	userIDStr, err := s.redis.Get(ctx, key).Result()
	if err == nil {
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return nil, ErrInvalidToken
		}
		return &userID, nil
	}

	// Fall back to database
	resetToken, err := s.repo.GetPasswordResetToken(ctx, token)
	if err != nil {
		return nil, err
	}
	if resetToken == nil {
		return nil, ErrInvalidToken
	}
	if time.Now().After(resetToken.ExpiresAt) {
		return nil, ErrTokenExpired
	}

	return &resetToken.UserID, nil
}

// GenerateEmailVerificationToken generates an email verification token
func (s *AuthService) GenerateEmailVerificationToken(ctx context.Context, userID uuid.UUID) (string, error) {
	token := uuid.New().String()
	verificationToken := &model.EmailVerificationToken{
		ID:        uuid.New(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(s.cfg.TokenExpiry.EmailVerification),
		Verified:  false,
		CreatedAt: time.Now(),
	}

	if err := s.repo.CreateEmailVerificationToken(ctx, verificationToken); err != nil {
		return "", err
	}

	return token, nil
}

// ValidateEmailVerificationToken validates an email verification token
func (s *AuthService) ValidateEmailVerificationToken(ctx context.Context, token string) (*uuid.UUID, error) {
	verificationToken, err := s.repo.GetEmailVerificationToken(ctx, token)
	if err != nil {
		return nil, err
	}
	if verificationToken == nil {
		return nil, ErrInvalidToken
	}
	if time.Now().After(verificationToken.ExpiresAt) {
		return nil, ErrTokenExpired
	}

	return &verificationToken.UserID, nil
}

// UpdatePassword updates user password
func (s *AuthService) UpdatePassword(ctx context.Context, userID uuid.UUID, hash string) error {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	user.PasswordHash = hash
	user.UpdatedAt = time.Now()
	return s.repo.UpdateUser(ctx, user)
}
