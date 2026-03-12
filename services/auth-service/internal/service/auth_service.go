package service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/game-engine/auth-service/internal/config"
	"github.com/game-engine/auth-service/internal/model"
	"github.com/game-engine/auth-service/internal/repository"
	"github.com/game-engine/auth-service/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/redis/go-redis/v9"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserExists         = errors.New("user already exists")
	ErrUser               = errors.New("NotFound          user not found")
	ErrAccountLocked      = errors.New("account is locked")
	ErrAccountSuspended   = errors.New("account is suspended")
	ErrAccountInactive    = errors.New("account is inactive")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrPasswordMismatch   = errors.New("passwords do not match")
	ErrTooManyAttempts    = errors.New("too many login attempts")
	Err2FARequired        = errors.New("2FA verification required")
	ErrInvalid2FA         = errors.New("invalid 2FA code")
	Err2FANotEnabled      = errors.New("2FA is not enabled")
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidPhone       = errors.New("invalid phone format")
	ErrTermsNotAccepted   = errors.New("terms must be accepted")
)

// AuthService handles authentication business logic
type AuthService struct {
	repo       *repository.AuthRepository
	redis      *redis.Client
	cfg        *config.Config
	privateKey *rsa.PrivateKey
	validate   *validator.Validate
}

// NewAuthService creates a new AuthService
func NewAuthService(repo *repository.AuthRepository, redisClient *redis.Client, cfg *config.Config) (*AuthService, error) {
	validate := validator.New()

	// Generate RSA key pair for JWT signing
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key: %w", err)
	}

	return &AuthService{
		repo:       repo,
		redis:      redisClient,
		cfg:        cfg,
		privateKey: privateKey,
		validate:   validate,
	}, nil
}

// ValidatePassword validates password complexity
func (s *AuthService) ValidatePassword(password string) error {
	if len(password) < s.cfg.Password.MinLength {
		return fmt.Errorf("password must be at least %d characters", s.cfg.Password.MinLength)
	}

	if s.cfg.Password.RequireUppercase && !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}

	if s.cfg.Password.RequireLowercase && !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}

	if s.cfg.Password.RequireNumber && !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New("password must contain at least one number")
	}

	if s.cfg.Password.RequireSpecial && !regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

// ValidateEmail validates email format
func (s *AuthService) ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

// ValidatePhone validates phone format
func (s *AuthService) ValidatePhone(phone string) error {
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	cleaned := strings.ReplaceAll(phone, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	if !phoneRegex.MatchString(cleaned) {
		return ErrInvalidPhone
	}
	return nil
}

// HashPassword creates a bcrypt hash of the password
func (s *AuthService) HashPassword(password string) (string, error) {
	return utils.HashPassword(password, s.cfg.Password.BcryptCost)
}

// CheckPassword compares a password with a hash
func (s *AuthService) CheckPassword(password, hash string) bool {
	return utils.CheckPassword(password, hash)
}

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

// ConnectNATS connects to NATS
func (s *AuthService) ConnectNATS() (*tls.Conn, error) {
	// Simplified NATS connection for now
	// In production, use nats.go library
	return nil, nil
}

// PublishEvent publishes an event to NATS
func (s *AuthService) PublishEvent(subject string, data interface{}) error {
	// Simplified - in production use NATS
	return nil
}

// CheckUserExists checks if a user exists
func (s *AuthService) CheckUserExists(ctx context.Context, identifier string) (bool, error) {
	// Try email first
	exists, err := s.repo.EmailExists(ctx, identifier)
	if err != nil {
		return false, err
	}
	if exists {
		return true, nil
	}
	return s.repo.PhoneExists(ctx, identifier)
}

// GetUserByIdentifier gets user by email or phone
func (s *AuthService) GetUserByIdentifier(ctx context.Context, identifier string) (*model.User, error) {
	// Try email first
	user, err := s.repo.GetUserByEmail(ctx, identifier)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	return s.repo.GetUserByPhone(ctx, identifier)
}

// GetUserByID gets user by ID
func (s *AuthService) GetUserByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}

// CreateUser creates a new user
func (s *AuthService) CreateUser(ctx context.Context, user *model.User) error {
	return s.repo.CreateUser(ctx, user)
}

// CreateSession creates a new session
func (s *AuthService) CreateSession(ctx context.Context, session *model.Session) error {
	return s.repo.CreateSession(ctx, session)
}

// DeleteSession deletes a session
func (s *AuthService) DeleteSession(ctx context.Context, sessionID uuid.UUID) error {
	return s.repo.DeleteSession(ctx, sessionID)
}

// DeleteAllUserSessions deletes all sessions for a user
func (s *AuthService) DeleteAllUserSessions(ctx context.Context, userID uuid.UUID) error {
	return s.repo.DeleteUserSessions(ctx, userID)
}

// GetUserIDFromSession gets user ID from session
func (s *AuthService) GetUserIDFromSession(ctx context.Context, sessionID uuid.UUID) (uuid.UUID, error) {
	session, err := s.repo.GetSessionByID(ctx, sessionID)
	if err != nil {
		return uuid.Nil, err
	}
	if session == nil {
		return uuid.Nil, ErrInvalidToken
	}
	return session.UserID, nil
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

// UpdateLastLogin updates last login time
func (s *AuthService) UpdateLastLogin(ctx context.Context, userID uuid.UUID) error {
	return s.repo.UpdateLastLogin(ctx, userID)
}

// MarkEmailVerified marks email as verified
func (s *AuthService) MarkEmailVerified(ctx context.Context, userID uuid.UUID) error {
	return s.repo.MarkEmailVerified(ctx, userID)
}

// MarkPhoneVerified marks phone as verified
func (s *AuthService) MarkPhoneVerified(ctx context.Context, userID uuid.UUID) error {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	user.PhoneVerified = true
	user.UpdatedAt = time.Now()
	return s.repo.UpdateUser(ctx, user)
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
