package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/game_engine/user-service/internal/model"
)

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
