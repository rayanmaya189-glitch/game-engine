package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// GenerateTOTPSecret generates a TOTP secret for 2FA
func (s *AuthService) GenerateTOTPSecret(userID uuid.UUID) (string, string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      s.cfg.TOTP.Issuer,
		AccountName: userID.String(),
		Algorithm:   otp.AlgorithmSHA1,
		Digits:      otp.DigitsSix,
		Period:      uint(s.cfg.TOTP.Period),
	})
	if err != nil {
		return "", "", err
	}

	return key.Secret(), key.URL(), nil
}

// ValidateTOTP validates a TOTP code
func (s *AuthService) ValidateTOTP(secret, code string) bool {
	return totp.Validate(code, secret)
}

// GenerateBackupCodes generates backup codes for 2FA
func (s *AuthService) GenerateBackupCodes(count int) ([]string, error) {
	codes := make([]string, count)
	for i := 0; i < count; i++ {
		// Generate 8-character alphanumeric codes
		code := make([]byte, 8)
		for j := 0; j < 8; j++ {
			n, _ := rand.Int(rand.Reader, big.NewInt(36))
			code[j] = "abcdefghijklmnopqrstuvwxyz0123456789"[n.Int64()]
		}
		codes[i] = strings.ToUpper(string(code))
	}
	return codes, nil
}

// Store2FASecret stores 2FA secret temporarily
func (s *AuthService) Store2FASecret(ctx context.Context, userID uuid.UUID, secret string) error {
	key := fmt.Sprintf("2fa_secret:%s", userID.String())
	return s.redis.Set(ctx, key, secret, 10*time.Minute).Err()
}

// ValidateBackupCode validates a backup code
func (s *AuthService) ValidateBackupCode(ctx context.Context, userID uuid.UUID, code string) bool {
	key := fmt.Sprintf("2fa_backup:%s:%s", userID.String(), code)
	result, err := s.redis.Get(ctx, key).Result()
	if err != nil || result != "valid" {
		return false
	}
	// Mark code as used
	s.redis.Del(ctx, key)
	return true
}

// Enable2FAForUser enables 2FA for a user
func (s *AuthService) Enable2FAForUser(ctx context.Context, userID uuid.UUID) error {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	user.TwoFactorEnabled = true
	user.UpdatedAt = time.Now()
	return s.repo.UpdateUser(ctx, user)
}

// Disable2FAForUser disables 2FA for a user
func (s *AuthService) Disable2FAForUser(ctx context.Context, userID uuid.UUID) error {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	user.TwoFactorEnabled = false
	user.TwoFactorSecret = ""
	user.UpdatedAt = time.Now()
	return s.repo.UpdateUser(ctx, user)
}
