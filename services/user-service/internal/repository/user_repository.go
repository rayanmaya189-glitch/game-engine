package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/game_engine/user-service/internal/model"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetProfileByUserID retrieves a profile by user ID
func (r *UserRepository) GetProfileByUserID(ctx context.Context, userID string) (*model.Profile, error) {
	query := `
		SELECT id, user_id, email, phone, username, display_name, first_name, last_name,
		       date_of_birth, gender, avatar_url, country, language, currency, timezone,
		       status, kyc_level, email_verified, phone_verified, two_factor_enabled,
		       created_at, updated_at, last_login_at
		FROM profiles
		WHERE user_id = $1
	`

	var p model.Profile
	var dob, lla sql.NullString

	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&p.ID, &p.UserID, &p.Email, &p.Phone, &p.Username, &p.DisplayName,
		&p.FirstName, &p.LastName, &dob, &p.Gender, &p.AvatarURL, &p.Country,
		&p.Language, &p.Currency, &p.Timezone, &p.Status, &p.KYCLevel,
		&p.EmailVerified, &p.PhoneVerified, &p.TwoFactorEnabled,
		&p.CreatedAt, &p.UpdatedAt, &lla,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if dob.Valid {
		t, _ := time.Parse(time.RFC3339, dob.String)
		p.DateOfBirth = &t
	}

	if lla.Valid {
		t, _ := time.Parse(time.RFC3339, lla.String)
		p.LastLoginAt = &t
	}

	return &p, nil
}

// UpdateProfile updates a player's profile
func (r *UserRepository) UpdateProfile(ctx context.Context, p *model.Profile) error {
	query := `
		UPDATE profiles
		SET display_name = $1, avatar_url = $2, language = $3, timezone = $4,
		    first_name = $5, last_name = $6, date_of_birth = $7, gender = $8,
		    updated_at = NOW()
		WHERE user_id = $9
	`

	_, err := r.db.ExecContext(ctx, query,
		p.DisplayName, p.AvatarURL, p.Language, p.Timezone,
		p.FirstName, p.LastName, p.DateOfBirth, p.Gender, p.UserID,
	)
	return err
}

// GetAddressByProfileID retrieves an address by profile ID
func (r *UserRepository) GetAddressByProfileID(ctx context.Context, profileID string) (*model.Address, error) {
	query := `
		SELECT id, profile_id, street, city, state, postal_code, country, created_at, updated_at
		FROM addresses
		WHERE profile_id = $1
	`

	var a model.Address
	err := r.db.QueryRowContext(ctx, query, profileID).Scan(
		&a.ID, &a.ProfileID, &a.Street, &a.City, &a.State, &a.PostalCode, &a.Country,
		&a.CreatedAt, &a.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &a, nil
}

// UpdateOrCreateAddress updates or creates an address
func (r *UserRepository) UpdateOrCreateAddress(ctx context.Context, profileID string, addr *model.Address) error {
	query := `
		INSERT INTO addresses (profile_id, street, city, state, postal_code, country, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
		ON CONFLICT (profile_id) DO UPDATE SET
			street = EXCLUDED.street,
			city = EXCLUDED.city,
			state = EXCLUDED.state,
			postal_code = EXCLUDED.postal_code,
			country = EXCLUDED.country,
			updated_at = NOW()
	`

	_, err := r.db.ExecContext(ctx, query,
		profileID, addr.Street, addr.City, addr.State, addr.PostalCode, addr.Country,
	)
	return err
}

// GetKYCStatus retrieves KYC status for a user
func (r *UserRepository) GetKYCStatus(ctx context.Context, userID string) (*model.KYCStatus, error) {
	query := `
		SELECT id, user_id, status, level, submitted_at, reviewed_at, rejection_reason, created_at, updated_at
		FROM kyc_status
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT 1
	`

	var k model.KYCStatus
	var sa, ra sql.NullString

	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&k.ID, &k.UserID, &k.Status, &k.Level, &sa, &ra, &k.RejectionReason,
		&k.CreatedAt, &k.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if sa.Valid {
		t, _ := time.Parse(time.RFC3339, sa.String)
		k.SubmittedAt = &t
	}

	if ra.Valid {
		t, _ := time.Parse(time.RFC3339, ra.String)
		k.ReviewedAt = &t
	}

	return &k, nil
}

// CreateKYCStatus creates a new KYC status record
func (r *UserRepository) CreateKYCStatus(ctx context.Context, k *model.KYCStatus) error {
	query := `
		INSERT INTO kyc_status (id, user_id, status, level, submitted_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW(), NOW())
	`

	_, err := r.db.ExecContext(ctx, query, k.ID, k.UserID, k.Status, k.Level)
	return err
}

// UpdateKYCStatus updates KYC status
func (r *UserRepository) UpdateKYCStatus(ctx context.Context, k *model.KYCStatus) error {
	query := `
		UPDATE kyc_status
		SET status = $1, level = $2, reviewed_at = NOW(), rejection_reason = $3, updated_at = NOW()
		WHERE user_id = $4
	`

	_, err := r.db.ExecContext(ctx, query, k.Status, k.Level, k.RejectionReason, k.UserID)
	return err
}

// CreateKYCDocument creates a new KYC document
func (r *UserRepository) CreateKYCDocument(ctx context.Context, d *model.KYCDocument) error {
	query := `
		INSERT INTO kyc_documents (id, user_id, document_type, document_number, document_data, status, submitted_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW(), NOW())
	`

	_, err := r.db.ExecContext(ctx, query,
		d.ID, d.UserID, d.DocumentType, d.DocumentNumber, d.DocumentData, d.Status,
	)
	return err
}

// GetPlayerSettings retrieves player settings
func (r *UserRepository) GetPlayerSettings(ctx context.Context, userID string) (*model.PlayerSettings, error) {
	query := `
		SELECT id, user_id, email_notifications, sms_notifications, push_notifications,
		       profile_public, show_online_status, auto_play, sound_volume, theme,
		       created_at, updated_at
		FROM player_settings
		WHERE user_id = $1
	`

	var ps model.PlayerSettings
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&ps.ID, &ps.UserID, &ps.EmailNotifications, &ps.SMSNotifications, &ps.PushNotifications,
		&ps.ProfilePublic, &ps.ShowOnlineStatus, &ps.AutoPlay, &ps.SoundVolume, &ps.Theme,
		&ps.CreatedAt, &ps.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &ps, nil
}

// UpdatePlayerSettings updates player settings
func (r *UserRepository) UpdatePlayerSettings(ctx context.Context, ps *model.PlayerSettings) error {
	query := `
		INSERT INTO player_settings (id, user_id, email_notifications, sms_notifications, push_notifications,
		                            profile_public, show_online_status, auto_play, sound_volume, theme,
		                            created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())
		ON CONFLICT (user_id) DO UPDATE SET
			email_notifications = EXCLUDED.email_notifications,
			sms_notifications = EXCLUDED.sms_notifications,
			push_notifications = EXCLUDED.push_notifications,
			profile_public = EXCLUDED.profile_public,
			show_online_status = EXCLUDED.show_online_status,
			auto_play = EXCLUDED.auto_play,
			sound_volume = EXCLUDED.sound_volume,
			theme = EXCLUDED.theme,
			updated_at = NOW()
	`

	_, err := r.db.ExecContext(ctx, query,
		ps.ID, ps.UserID, ps.EmailNotifications, ps.SMSNotifications, ps.PushNotifications,
		ps.ProfilePublic, ps.ShowOnlineStatus, ps.AutoPlay, ps.SoundVolume, ps.Theme,
	)
	return err
}

// GetPlayerLimits retrieves player limits
func (r *UserRepository) GetPlayerLimits(ctx context.Context, userID string) (*model.PlayerLimits, error) {
	query := `
		SELECT id, user_id, daily_limit, weekly_limit, monthly_limit, daily_loss_limit,
		       self_exclusion, exclusion_end_date, created_at, updated_at
		FROM player_limits
		WHERE user_id = $1
	`

	var pl model.PlayerLimits
	var eed sql.NullString

	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&pl.ID, &pl.UserID, &pl.DailyLimit, &pl.WeeklyLimit, &pl.MonthlyLimit, &pl.DailyLossLimit,
		&pl.SelfExclusion, &eed, &pl.CreatedAt, &pl.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if eed.Valid {
		t, _ := time.Parse(time.RFC3339, eed.String)
		pl.ExclusionEndDate = &t
	}

	return &pl, nil
}

// UpdatePlayerLimits updates player limits
func (r *UserRepository) UpdatePlayerLimits(ctx context.Context, pl *model.PlayerLimits) error {
	query := `
		INSERT INTO player_limits (id, user_id, daily_limit, weekly_limit, monthly_limit, daily_loss_limit,
		                          self_exclusion, exclusion_end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		ON CONFLICT (user_id) DO UPDATE SET
			daily_limit = EXCLUDED.daily_limit,
			weekly_limit = EXCLUDED.weekly_limit,
			monthly_limit = EXCLUDED.monthly_limit,
			daily_loss_limit = EXCLUDED.daily_loss_limit,
			self_exclusion = EXCLUDED.self_exclusion,
			exclusion_end_date = EXCLUDED.exclusion_end_date,
			updated_at = NOW()
	`

	_, err := r.db.ExecContext(ctx, query,
		pl.ID, pl.UserID, pl.DailyLimit, pl.WeeklyLimit, pl.MonthlyLimit, pl.DailyLossLimit,
		pl.SelfExclusion, pl.ExclusionEndDate,
	)
	return err
}

// GetPlayerStats retrieves player statistics
func (r *UserRepository) GetPlayerStats(ctx context.Context, userID string) (*model.PlayerStats, error) {
	query := `
		SELECT id, user_id, total_deposits, total_withdrawals, total_bets, total_wins, total_bonuses,
		       deposit_count, withdrawal_count, bet_count, win_count, created_at, updated_at
		FROM player_stats
		WHERE user_id = $1
	`

	var ps model.PlayerStats
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&ps.ID, &ps.UserID, &ps.TotalDeposits, &ps.TotalWithdrawals, &ps.TotalBets, &ps.TotalWins,
		&ps.TotalBonuses, &ps.DepositCount, &ps.WithdrawalCount, &ps.BetCount, &ps.WinCount,
		&ps.CreatedAt, &ps.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &ps, nil
}

// UpdatePlayerStatus updates player status
func (r *UserRepository) UpdatePlayerStatus(ctx context.Context, userID, status string) error {
	query := `
		UPDATE profiles
		SET status = $1, updated_at = NOW()
		WHERE user_id = $2
	`

	_, err := r.db.ExecContext(ctx, query, status, userID)
	return err
}

// UpdatePlayerKYCLevel updates player KYC level
func (r *UserRepository) UpdatePlayerKYCLevel(ctx context.Context, userID, level string) error {
	query := `
		UPDATE profiles
		SET kyc_level = $1, updated_at = NOW()
		WHERE user_id = $2
	`

	_, err := r.db.ExecContext(ctx, query, level, userID)
	return err
}

// ListPlayers lists players with filters and pagination
func (r *UserRepository) ListPlayers(ctx context.Context, status, kycLevel, country, search string, limit, offset int) ([]*model.Profile, int, error) {
	baseQuery := `
		FROM profiles
		WHERE 1=1
	`
	args := []interface{}{}
	argIndex := 1

	if status != "" {
		baseQuery += fmt.Sprintf(" AND status = $%d", argIndex)
		args = append(args, status)
		argIndex++
	}

	if kycLevel != "" {
		baseQuery += fmt.Sprintf(" AND kyc_level = $%d", argIndex)
		args = append(args, kycLevel)
		argIndex++
	}

	if country != "" {
		baseQuery += fmt.Sprintf(" AND country = $%d", argIndex)
		args = append(args, country)
		argIndex++
	}

	if search != "" {
		baseQuery += fmt.Sprintf(" AND (email ILIKE $%d OR username ILIKE $%d OR display_name ILIKE $%d)",
			argIndex, argIndex, argIndex)
		args = append(args, "%"+search+"%")
		argIndex++
	}

	// Get total count
	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Get paginated results
	selectQuery := `
		SELECT id, user_id, email, phone, username, display_name, first_name, last_name,
		       date_of_birth, gender, avatar_url, country, language, currency, timezone,
		       status, kyc_level, email_verified, phone_verified, two_factor_enabled,
		       created_at, updated_at, last_login_at
	` + baseQuery + fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, limit, offset)

	rows, err := r.db.QueryContext(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var profiles []*model.Profile
	for rows.Next() {
		var p model.Profile
		var dob, lla sql.NullString

		err := rows.Scan(
			&p.ID, &p.UserID, &p.Email, &p.Phone, &p.Username, &p.DisplayName,
			&p.FirstName, &p.LastName, &dob, &p.Gender, &p.AvatarURL, &p.Country,
			&p.Language, &p.Currency, &p.Timezone, &p.Status, &p.KYCLevel,
			&p.EmailVerified, &p.PhoneVerified, &p.TwoFactorEnabled,
			&p.CreatedAt, &p.UpdatedAt, &lla,
		)
		if err != nil {
			return nil, 0, err
		}

		if dob.Valid {
			t, _ := time.Parse(time.RFC3339, dob.String)
			p.DateOfBirth = &t
		}

		if lla.Valid {
			t, _ := time.Parse(time.RFC3339, lla.String)
			p.LastLoginAt = &t
		}

		profiles = append(profiles, &p)
	}

	return profiles, total, nil
}

// GetProfileByIdentifier retrieves a profile by email, phone, or username (for admin lookup)
func (r *UserRepository) GetProfileByIdentifier(ctx context.Context, identifier string) (*model.Profile, error) {
	query := `
		SELECT id, user_id, email, phone, username, display_name, first_name, last_name,
		       date_of_birth, gender, avatar_url, country, language, currency, timezone,
		       status, kyc_level, email_verified, phone_verified, two_factor_enabled,
		       created_at, updated_at, last_login_at
		FROM profiles
		WHERE user_id = $1 OR email = $1 OR phone = $1 OR username = $1
	`

	var p model.Profile
	var dob, lla sql.NullString

	err := r.db.QueryRowContext(ctx, query, identifier).Scan(
		&p.ID, &p.UserID, &p.Email, &p.Phone, &p.Username, &p.DisplayName,
		&p.FirstName, &p.LastName, &dob, &p.Gender, &p.AvatarURL, &p.Country,
		&p.Language, &p.Currency, &p.Timezone, &p.Status, &p.KYCLevel,
		&p.EmailVerified, &p.PhoneVerified, &p.TwoFactorEnabled,
		&p.CreatedAt, &p.UpdatedAt, &lla,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if dob.Valid {
		t, _ := time.Parse(time.RFC3339, dob.String)
		p.DateOfBirth = &t
	}

	if lla.Valid {
		t, _ := time.Parse(time.RFC3339, lla.String)
		p.LastLoginAt = &t
	}

	return &p, nil
}
