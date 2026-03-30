package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/game_engine/user-service/internal/model"
)

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
