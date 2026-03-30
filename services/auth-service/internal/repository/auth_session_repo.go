package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/game_engine/auth-service/internal/model"
	"github.com/google/uuid"
)

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
