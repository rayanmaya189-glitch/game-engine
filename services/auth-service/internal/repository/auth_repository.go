package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/game_engine/auth-service/internal/model"
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
	if err != nil {
		return nil, err
	}

	roles, err := r.GetUserRoles(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	user.Roles = roles

	return user, nil
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
	if err != nil {
		return nil, err
	}

	roles, err := r.GetUserRoles(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	user.Roles = roles

	return user, nil
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
	if err != nil {
		return nil, err
	}

	roles, err := r.GetUserRoles(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	user.Roles = roles

	return user, nil
}

// GetUserRoles retrieves roles for a user
func (r *AuthRepository) GetUserRoles(ctx context.Context, userID uuid.UUID) ([]model.UserRole, error) {
	query := `
		SELECT r.name FROM roles r
		JOIN user_roles ur ON r.id = ur.role_id
		WHERE ur.user_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []model.UserRole
	for rows.Next() {
		var roleName string
		if err := rows.Scan(&roleName); err != nil {
			return nil, err
		}
		roles = append(roles, model.UserRole(roleName))
	}

	// Default to player if no roles found
	if len(roles) == 0 {
		roles = append(roles, model.RolePlayer)
	}

	return roles, nil
}

// Seed basic roles, permissions and superadmin user
func (r *AuthRepository) Seed(ctx context.Context) error {
	query := `
-- Seed basic roles
INSERT INTO roles (name, description) VALUES 
('superadmin', 'System Super Administrator with full access'),
('admin', 'Administrator with management access'),
('support', 'Support staff with read and limited write access'),
('player', 'Regular player')
ON CONFLICT (name) DO NOTHING;

-- Seed basic permissions
INSERT INTO permissions (name, description) VALUES 
('all', 'Global permission for all actions'),
('user.read', 'Read user profiles'),
('user.write', 'Modify user profiles'),
('user.delete', 'Delete users'),
('game.manage', 'Manage games and settings'),
('finance.manage', 'Manage payments and wallets')
ON CONFLICT (name) DO NOTHING;

-- Assign 'all' permission to 'superadmin'
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p 
WHERE r.name = 'superadmin' AND p.name = 'all'
ON CONFLICT DO NOTHING;

-- Create superadmin user
INSERT INTO users (
    email, 
    password_hash, 
    status, 
    email_verified, 
    marketing_consent, 
    accept_terms,
    created_at,
    updated_at
) VALUES (
    'admin@game-engine.com', 
    '$2a$12$A6Ae/bdusqjinL0zx/8CCOR50/aMbfEf6uLGU2sJDYD2TvrLgo1ga', 
    'active', 
    true, 
    true, 
    true,
    NOW(),
    NOW()
) ON CONFLICT (email) DO NOTHING;

-- Assign superadmin role to superadmin user
INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id FROM users u, roles r
WHERE u.email = 'admin@game-engine.com' AND r.name = 'superadmin'
ON CONFLICT DO NOTHING;
	`
	_, err := r.db.ExecContext(ctx, query)
	return err
}

// Migrate executes all migration files in the migrations directory
func (r *AuthRepository) Migrate(ctx context.Context) error {
	// Simple migration runner that executes SQL files in order
	migrations := []string{
		"auth-service/migrations/001_init_schema.sql",
		"auth-service/migrations/002_add_roles_and_permissions.sql",
	}

	for _, file := range migrations {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		_, err = r.db.ExecContext(ctx, string(content))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", file, err)
		}
	}

	return nil
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

// MarkEmailVerified marks an email as verified
func (r *AuthRepository) MarkEmailVerified(ctx context.Context, userID uuid.UUID) error {
	query := `UPDATE users SET email_verified = true, updated_at = NOW() WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}
