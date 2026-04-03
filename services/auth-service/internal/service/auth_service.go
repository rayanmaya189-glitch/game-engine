package service

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/game_engine/auth-service/internal/config"
	"github.com/game_engine/auth-service/internal/model"
	"github.com/game_engine/auth-service/internal/repository"
	"github.com/game_engine/auth-service/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

const (
	EventUserRegistered = "auth.events.user_registered"
	EventUserLoggedIn   = "auth.events.user_logged_in"
	EventUserLoggedOut  = "auth.events.user_logged_out"
	EventTokenRefreshed = "auth.events.token_refreshed"
	EventPasswordReset  = "auth.events.password_reset"
	Event2FAEnabled     = "auth.events.2fa_enabled"
	Event2FADisabled    = "auth.events.2fa_disabled"
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
	nats       *nats.Conn
}

// NewAuthService creates a new AuthService
func NewAuthService(repo *repository.AuthRepository, redisClient *redis.Client, cfg *config.Config) (*AuthService, error) {
	validate := validator.New()

	// Load RSA private key from file
	privateKey, err := loadPrivateKey(cfg.JWT.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load RSA private key: %w", err)
	}

	return &AuthService{
		repo:       repo,
		redis:      redisClient,
		cfg:        cfg,
		privateKey: privateKey,
		validate:   validate,
	}, nil
}

// loadPrivateKey reads an RSA private key from a PEM file
func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block from private key file")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// Try PKCS8 format as fallback
		parsedKey, pkcs8Err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if pkcs8Err != nil {
			return nil, fmt.Errorf("failed to parse private key: %w", err)
		}
		var ok bool
		key, ok = parsedKey.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("private key is not RSA")
		}
	}

	return key, nil
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

// ConnectNATS connects to NATS
func (s *AuthService) ConnectNATS(url string) error {
	nc, err := nats.Connect(url)
	if err != nil {
		return fmt.Errorf("failed to connect to NATS: %w", err)
	}
	s.nats = nc
	return nil
}

// DisconnectNATS closes the NATS connection
func (s *AuthService) DisconnectNATS() {
	if s.nats != nil {
		s.nats.Close()
	}
}

// PublishEvent publishes an event to NATS
func (s *AuthService) PublishEvent(subject string, data []byte) error {
	if s.nats == nil {
		return nil // NATS not configured
	}
	return s.nats.Publish(subject, data)
}

// SubscribeToEvents subscribes to NATS events
func (s *AuthService) SubscribeToEvents(subject string, handler func(*nats.Msg)) (*nats.Subscription, error) {
	if s.nats == nil {
		return nil, fmt.Errorf("NATS not connected")
	}
	return s.nats.Subscribe(subject, handler)
}
