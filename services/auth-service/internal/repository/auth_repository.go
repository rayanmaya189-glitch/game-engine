package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/game-engine/auth-service/internal/model"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// AuthRepository handles database operations for authentication
type AuthRepository struct {
	db *sql.DB
}

// NewAuthRepository creates a new AuthRepository
func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// CreateUser creates a new user in the database
func (r *AuthRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (
			id, email, phone, password_hash, country, language, currency,
			status, email_verified, phone_verified, two_factor_enabled,
			marketing_consent, accept_terms, referral_code, referred_by,
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Email, user.Phone, user.PasswordHash,
		user.Country, user.Language, user.Currency, user.Status,
		user.EmailVerified, user.PhoneVerified, user.TwoFactorEnabled,
		user.MarketingConsent, user.AcceptTerms, user.ReferralCode,
		user.ReferredBy, user.CreatedAt, user.UpdatedAt,
	)
	return err
}

// GetUserByEmail retrieves a user by email
func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
		SELECT id, email, phone, password_hash, country, language, currency,
			status, email_verified, phone_verified, two_factor_enabled,
			two_factor_secret, marketing_consent, accept_terms, referral_code,
			referred_by, failed_login_count, locked_until, created_at, updated_at,
			last_login_at
		FROM users WHERE email = $1
	`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Phone, &user.PasswordHash,
		&user.Country, &user.Language, &user.Currency, &user.Status,
		&user.EmailVerified, &user.PhoneVerified, &user.TwoFactorEnabled,
		&user.TwoFactorSecret, &user.MarketingConsent, &user.AcceptTerms,
		&user.ReferralCode, &user.ReferredBy, &user.FailedLoginCount,
		&user.LockedUntil, &user.CreatedAt, &user.UpdatedAt, &user.LastLoginAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

// GetUserByPhone retrieves a user by phone
func (r *AuthRepository) GetUserByPhone(ctx context.Context, phone string) (*model.User, error) {
	query := `
		SELECT id, email, phone, password_hash, country, language, currency,
			status, email_verified, phone_verified, two_factor_enabled,
			two_factor_secret, marketing_consent, accept_terms, referral_code,
			referred_by, failed_login_count, locked_until, created_at, updated_at,
			last_login_at
		FROM users WHERE phone = $1
	`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, phone).Scan(
		&user.ID, &user.Email, &user.Phone, &user.PasswordHash,
		&user.Country, &user.Language, &user.Currency, &user.Status,
		&user.EmailVerified, &user.PhoneVerified, &user.TwoFactorEnabled,
		&user.TwoFactorSecret, &user.MarketingConsent, &user.AcceptTerms,
		&user.ReferralCode, &user.ReferredBy, &user.FailedLoginCount,
		&user.LockedUntil, &user.CreatedAt, &user.UpdatedAt, &user.LastLoginAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

// GetUserByID retrieves a user by ID
func (r *AuthRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	query := `
		SELECT id, email, phone, password_hash, country, language, currency,
			status, email_verified, phone_verified, two_factor_enabled,
			two_factor_secret, marketing_consent, accept_terms, referral_code,
			referred_by, failed_login_count, locked_until, created_at, updated_at,
			last_login_at
		FROM users WHERE id = $1
	`

	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.Phone, &user.PasswordHash,
		&user.Country, &user.Language, &user.Currency, &user.Status,
		&user.EmailVerified, &user.PhoneVerified, &user.TwoFactorEnabled,
		&user.TwoFactorSecret, &user.MarketingConsent, &user.AcceptTerms,
		&user.ReferralCode, &user.ReferredBy, &user.FailedLoginCount,
		&user.LockedUntil, &user.CreatedAt, &user.UpdatedAt, &user.LastLoginAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

// UpdateUser updates a user in the database
func (r *AuthRepository) UpdateUser(ctx context.Context, user *model.User) error {
	query := `
		UPDATE users SET
			email = $2, phone = $3, password_hash = $4, country = $5,
			language = $6, currency = $7, status = $8, email_verified = $9,
			phone_verified = $10, two_factor_enabled = $11, two_factor_secret = $12,
			marketing_consent = $13, accept_terms = $14, referral_code = $15,
			referred_by = $16, failed_login_count = $17, locked_until = $18,
			updated_at = $19, last_login_at = $20
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Email, user.Phone, user.PasswordHash, user.Country,
		user.Language, user.Currency, user.Status, user.EmailVerified,
		user.PhoneVerified, user.TwoFactorEnabled, user.TwoFactorSecret,
		user.MarketingConsent, user.AcceptTerms, user.ReferralCode,
		user.ReferredBy, user.FailedLoginCount, user.LockedUntil,
		user.UpdatedAt, user.LastLoginAt,
	)
	return err
}

// CreateSession creates a new session in the database
func (r *AuthRepository) CreateSession(ctx context.Context, session *model.Session) error {
	query := `
		INSERT INTO sessions (
			id, user_id, device_type, os_type, browser_type, device_id,
			device_name, ip_address, user_agent, refresh_token, expires_at,
			created_at, last_used_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := r.db.ExecContext(ctx, query,
		session.ID, session.UserID,
		session.DeviceInfo.DeviceType, session.DeviceInfo.OSType,
		session.DeviceInfo.BrowserType, session.DeviceInfo.DeviceID,
		session.DeviceInfo.DeviceName, session.DeviceInfo.IPAddress,
		session.DeviceInfo.UserAgent, session.RefreshToken,
		session.ExpiresAt, session.CreatedAt, session.LastUsedAt,
	)
	return err
}

// GetSessionByID retrieves a session by ID
func (r *AuthRepository) GetSessionByID(ctx context.Context, id uuid.UUID) (*model.Session, error) {
	query := `
		SELECT id, user_id, device_type, os_type, browser_type, device_id,
			device_name, ip_address, user_agent, refresh_token, expires_at,
			created_at, last_used_at
		FROM sessions WHERE id = $1 AND expires_at > NOW()
	`

	session := &model.Session{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&session.ID, &session.UserID,
		&session.DeviceInfo.DeviceType, &session.DeviceInfo.OSType,
		&session.DeviceInfo.BrowserType, &session.DeviceInfo.DeviceID,
		&session.DeviceInfo.DeviceName, &session.DeviceInfo.IPAddress,
		&session.DeviceInfo.UserAgent, &session.RefreshToken,
		&session.ExpiresAt, &session.CreatedAt, &session.LastUsedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return session, err
}

// DeleteSession deletes a session
func (r *AuthRepository) DeleteSession(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM sessions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// DeleteUserSessions deletes all sessions for a user
func (r *AuthRepository) DeleteUserSessions(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM sessions WHERE user_id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

// RecordLoginAttempt records a login attempt
func (r *AuthRepository) RecordLoginAttempt(ctx context.Context, attempt *model.LoginAttempt) error {
	query := `
		INSERT INTO login_attempts (id, user_id, ip_address, attempted_at, success)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query,
		attempt.ID, attempt.UserID, attempt.IPAddress, attempt.AttemptedAt, attempt.Success,
	)
	return err
}

// GetLoginAttempts gets the number of failed login attempts in a time window
func (r *AuthRepository) GetLoginAttempts(ctx context.Context, userID uuid.UUID, since time.Time) (int, error) {
	query := `
		SELECT COUNT(*) FROM login_attempts
		WHERE user_id = $1 AND attempted_at > $2 AND success = false
	`

	var count int
	err := r.db.QueryRowContext(ctx, query, userID, since).Scan(&count)
	return count, err
}

// CreatePasswordResetToken creates a password reset token
func (r *AuthRepository) CreatePasswordResetToken(ctx context.Context, token *model.PasswordResetToken) error {
	query := `
		INSERT INTO password_reset_tokens (id, user_id, token, expires_at, used, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.ExecContext(ctx, query,
		token.ID, token.UserID, token.Token, token.ExpiresAt, token.Used, token.CreatedAt,
	)
	return err
}

// GetPasswordResetToken retrieves a password reset token
func (r *AuthRepository) GetPasswordResetToken(ctx context.Context, token string) (*model.PasswordResetToken, error) {
	query := `
		SELECT id, user_id, token, expires_at, used, created_at
		FROM password_reset_tokens WHERE token = $1 AND used = false
	`

	t := &model.PasswordResetToken{}
	err := r.db.QueryRowContext(ctx, query, token).Scan(
		&t.ID, &t.UserID, &t.Token, &t.ExpiresAt, &t.Used, &t.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return t, err
}

// MarkPasswordResetTokenUsed marks a password reset token as used
func (r *AuthRepository) MarkPasswordResetTokenUsed(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE password_reset_tokens SET used = true WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// CreateEmailVerificationToken creates an email verification token
func (r *AuthRepository) CreateEmailVerificationToken(ctx context.Context, token *model.EmailVerificationToken) error {
	query := `
		INSERT INTO email_verification_tokens (id, user_id, token, expires_at, verified, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.ExecContext(ctx, query,
		token.ID, token.UserID, token.Token, token.ExpiresAt, token.Verified, token.CreatedAt,
	)
	return err
}

// GetEmailVerificationToken retrieves an email verification token
func (r *AuthRepository) GetEmailVerificationToken(ctx context.Context, token string) (*model.EmailVerificationToken, error) {
	query := `
		SELECT id, user_id, token, expires_at, verified, created_at
		FROM email_verification_tokens WHERE token = $1 AND verified = false
	`

	t := &model.EmailVerificationToken{}
	err := r.db.QueryRowContext(ctx, query, token).Scan(
		&t.ID, &t.UserID, &t.Token, &t.ExpiresAt, &t.Verified, &t.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return t, err
}

// MarkEmailVerified marks an email as verified
func (r *AuthRepository) MarkEmailVerified(ctx context.Context, userID uuid.UUID) error {
	query := `UPDATE users SET email_verified = true, updated_at = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

// EmailExists checks if an email already exists
func (r *AuthRepository) EmailExists(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
	return exists, err
}

// PhoneExists checks if a phone already exists
func (r *AuthRepository) PhoneExists(ctx context.Context, phone string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE phone = $1)`
	var exists bool
	err := r.db.QueryRowContext(ctx, query, phone).Scan(&exists)
	return exists, err
}

// UpdateFailedLoginCount updates the failed login count for a user
func (r *AuthRepository) UpdateFailedLoginCount(ctx context.Context, userID uuid.UUID, count int, lockedUntil *time.Time) error {
	query := `UPDATE users SET failed_login_count = $2, locked_until = $3, updated_at = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userID, count, lockedUntil)
	return err
}

// UpdateLastLogin updates the last login time for a user
func (r *AuthRepository) UpdateLastLogin(ctx context.Context, userID uuid.UUID) error {
	query := `UPDATE users SET last_login_at = NOW(), failed_login_count = 0, locked_until = NULL, updated_at = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}

// GetReferralCode generates a unique referral code
func (r *AuthRepository) GetReferralCode(ctx context.Context) (string, error) {
	for i := 0; i < 10; i++ {
		code := fmt.Sprintf("GE%d", uuid.New().ID()%100000)
		query := `SELECT EXISTS(SELECT 1 FROM users WHERE referral_code = $1)`
		var exists bool
		err := r.db.QueryRowContext(ctx, query, code).Scan(&exists)
		if err != nil {
			return "", err
		}
		if !exists {
			return code, nil
		}
	}
	return "", fmt.Errorf("failed to generate unique referral code")
}
