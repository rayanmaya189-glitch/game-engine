package repository

import (
	"context"
	"database/sql"
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
