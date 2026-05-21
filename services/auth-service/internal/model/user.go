package model

import (
	"time"

	"github.com/google/uuid"
)

// UserStatus represents the status of a user account
type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusLocked    UserStatus = "locked"
	UserStatusInactive  UserStatus = "inactive"
)

// User represents a user account in the system
type User struct {
	ID               uuid.UUID  `json:"id" db:"id"`
	Email            string     `json:"email" db:"email"`
	Phone            string     `json:"phone" db:"phone"`
	PasswordHash     string     `json:"-" db:"password_hash"`
	Country          string     `json:"country" db:"country"`
	Language         string     `json:"language" db:"language"`
	Currency         string     `json:"currency" db:"currency"`
	Status           UserStatus `json:"status" db:"status"`
	EmailVerified    bool       `json:"email_verified" db:"email_verified"`
	PhoneVerified    bool       `json:"phone_verified" db:"phone_verified"`
	TwoFactorEnabled bool       `json:"two_factor_enabled" db:"two_factor_enabled"`
	TwoFactorSecret  string     `json:"-" db:"two_factor_secret"`
	MarketingConsent bool       `json:"marketing_consent" db:"marketing_consent"`
	AcceptTerms      bool       `json:"accept_terms" db:"accept_terms"`
	ReferralCode     string     `json:"referral_code" db:"referral_code"`
	ReferredBy       *uuid.UUID `json:"referred_by,omitempty" db:"referred_by"`
	FailedLoginCount int        `json:"failed_login_count" db:"failed_login_count"`
	LockedUntil      *time.Time `json:"locked_until,omitempty" db:"locked_until"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at" db:"updated_at"`
	LastLoginAt      *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
	Roles            []UserRole `json:"roles" db:"-"`
}

// Role represents a user role
type Role struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions"`
}

// Permission represents a user permission
type Permission struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

// Session represents an active user session
type Session struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	UserID       uuid.UUID  `json:"user_id" db:"user_id"`
	DeviceInfo   DeviceInfo `json:"device_info" db:"device_info"`
	RefreshToken string     `json:"refresh_token" db:"refresh_token"`
	IPAddress    string     `json:"ip_address" db:"ip_address"`
	UserAgent    string     `json:"user_agent" db:"user_agent"`
	ExpiresAt    time.Time  `json:"expires_at" db:"expires_at"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	LastUsedAt   time.Time  `json:"last_used_at" db:"last_used_at"`
}

// DeviceInfo represents device information
type DeviceInfo struct {
	DeviceType  string `json:"device_type" db:"device_type"`
	OSType      string `json:"os_type" db:"os_type"`
	BrowserType string `json:"browser_type" db:"browser_type"`
	DeviceID    string `json:"device_id" db:"device_id"`
	DeviceName  string `json:"device_name" db:"device_name"`
	IPAddress   string `json:"ip_address" db:"ip_address"`
	UserAgent   string `json:"user_agent" db:"user_agent"`
	Country     string `json:"country" db:"country"`
	City        string `json:"city" db:"city"`
	Timezone    string `json:"timezone" db:"timezone"`
}

// LoginAttempt tracks failed login attempts
type LoginAttempt struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
	IPAddress   string    `json:"ip_address" db:"ip_address"`
	AttemptedAt time.Time `json:"attempted_at" db:"attempted_at"`
	Success     bool      `json:"success" db:"success"`
}

// PasswordResetToken represents a password reset token
type PasswordResetToken struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Token     string    `json:"token" db:"token"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	Used      bool      `json:"used" db:"used"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// EmailVerificationToken represents an email verification token
type EmailVerificationToken struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Token     string    `json:"token" db:"token"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	Verified  bool      `json:"verified" db:"verified"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// UserRole represents the role of a user
type UserRole string

const (
	RolePlayer     UserRole = "player"
	RoleAdmin      UserRole = "admin"
	RoleSuperAdmin UserRole = "superadmin"
	RoleSupport    UserRole = "support"
)

// TokenClaims represents JWT token claims
type TokenClaims struct {
	UserID    uuid.UUID  `json:"user_id"`
	SessionID uuid.UUID  `json:"session_id"`
	Roles     []UserRole `json:"roles"`
	TokenType string     `json:"token_type"`
}

// RefreshTokenData represents refresh token data stored in Redis
type RefreshTokenData struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	DeviceID  string `json:"device_id"`
	ExpiresAt int64  `json:"expires_at"`
}
