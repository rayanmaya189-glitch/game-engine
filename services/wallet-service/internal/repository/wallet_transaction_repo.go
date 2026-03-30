package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/game_engine/wallet-service/internal/model"
	"github.com/google/uuid"
)

// CreateTransaction creates a new transaction
func (r *WalletRepository) CreateTransaction(ctx context.Context, tx *model.Transaction) error {
	if tx.TransactionID == "" {
		tx.TransactionID = uuid.New().String()
	}
	tx.CreatedAt = time.Now()

	query := `
		INSERT INTO transactions (
			transaction_id, user_id, type, status, currency, amount, bonus_amount, fee, net_amount,
			payment_method, payment_provider, payment_reference, game_id, bet_id, description, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`

	_, err := r.db.ExecContext(ctx, query,
		tx.TransactionID,
		tx.UserID,
		tx.Type,
		tx.Status,
		tx.Currency,
		tx.Amount,
		tx.BonusAmount,
		tx.Fee,
		tx.NetAmount,
		tx.PaymentMethod,
		tx.PaymentProvider,
		tx.PaymentReference,
		tx.GameID,
		tx.BetID,
		tx.Description,
		tx.CreatedAt,
	)

	return err
}

// UpdateTransactionStatus updates transaction status
func (r *WalletRepository) UpdateTransactionStatus(ctx context.Context, txID, status string, processedAt *time.Time) error {
	query := `
		UPDATE transactions
		SET status = $1, processed_at = $2
		WHERE transaction_id = $3
	`

	_, err := r.db.ExecContext(ctx, query, status, processedAt, txID)
	return err
}

// GetTransactionByID retrieves a transaction by ID
func (r *WalletRepository) GetTransactionByID(ctx context.Context, txID string) (*model.Transaction, error) {
	query := `
		SELECT transaction_id, user_id, type, status, currency, amount, bonus_amount, fee, net_amount,
			payment_method, payment_provider, payment_reference, game_id, bet_id, description, created_at, processed_at
		FROM transactions
		WHERE transaction_id = $1
	`

	tx := &model.Transaction{}
	err := r.db.QueryRowContext(ctx, query, txID).Scan(
		&tx.TransactionID,
		&tx.UserID,
		&tx.Type,
		&tx.Status,
		&tx.Currency,
		&tx.Amount,
		&tx.BonusAmount,
		&tx.Fee,
		&tx.NetAmount,
		&tx.PaymentMethod,
		&tx.PaymentProvider,
		&tx.PaymentReference,
		&tx.GameID,
		&tx.BetID,
		&tx.Description,
		&tx.CreatedAt,
		&tx.ProcessedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// GetTransactionHistory retrieves transaction history with filters
func (r *WalletRepository) GetTransactionHistory(ctx context.Context, userID string, txTypes []string, statuses []string, startDate, endDate time.Time, page, pageSize int) ([]*model.Transaction, int, error) {
	baseQuery := `
		FROM transactions
		WHERE user_id = $1
	`
	args := []interface{}{userID}
	argIndex := 2

	if len(txTypes) > 0 {
		baseQuery += fmt.Sprintf(" AND type = ANY($%d)", argIndex)
		args = append(args, txTypes)
		argIndex++
	}

	if len(statuses) > 0 {
		baseQuery += fmt.Sprintf(" AND status = ANY($%d)", argIndex)
		args = append(args, statuses)
		argIndex++
	}

	if !startDate.IsZero() {
		baseQuery += fmt.Sprintf(" AND created_at >= $%d", argIndex)
		args = append(args, startDate)
		argIndex++
	}

	if !endDate.IsZero() {
		baseQuery += fmt.Sprintf(" AND created_at <= $%d", argIndex)
		args = append(args, endDate)
		argIndex++
	}

	// Count total
	countQuery := "SELECT COUNT(*) " + baseQuery
	var total int
	err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * pageSize
	selectQuery := `
		SELECT transaction_id, user_id, type, status, currency, amount, bonus_amount, fee, net_amount,
			payment_method, payment_provider, payment_reference, game_id, bet_id, description, created_at, processed_at
		` + baseQuery + fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, pageSize, offset)

	rows, err := r.db.QueryContext(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var transactions []*model.Transaction
	for rows.Next() {
		tx := &model.Transaction{}
		err := rows.Scan(
			&tx.TransactionID,
			&tx.UserID,
			&tx.Type,
			&tx.Status,
			&tx.Currency,
			&tx.Amount,
			&tx.BonusAmount,
			&tx.Fee,
			&tx.NetAmount,
			&tx.PaymentMethod,
			&tx.PaymentProvider,
			&tx.PaymentReference,
			&tx.GameID,
			&tx.BetID,
			&tx.Description,
			&tx.CreatedAt,
			&tx.ProcessedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		transactions = append(transactions, tx)
	}

	return transactions, total, nil
}

// CreateBonusTransaction creates a bonus transaction
func (r *WalletRepository) CreateBonusTransaction(ctx context.Context, bt *model.BonusTransaction) error {
	if bt.ID == "" {
		bt.ID = uuid.New().String()
	}
	bt.CreatedAt = time.Now()

	query := `
		INSERT INTO bonus_transactions (
			id, user_id, transaction_id, bonus_type, currency, amount, wagering_multiplier,
			wagering_required, wagering_met, bonus_code, status, expires_at, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	_, err := r.db.ExecContext(ctx, query,
		bt.ID,
		bt.UserID,
		bt.TransactionID,
		bt.BonusType,
		bt.Currency,
		bt.Amount,
		bt.WageringMultiplier,
		bt.WageringRequired,
		bt.WageringMet,
		bt.BonusCode,
		bt.Status,
		bt.ExpiresAt,
		bt.CreatedAt,
	)

	return err
}

// GetActiveBonusTransactions retrieves active bonus transactions for a user
func (r *WalletRepository) GetActiveBonusTransactions(ctx context.Context, userID, currency string) ([]*model.BonusTransaction, error) {
	query := `
		SELECT id, user_id, transaction_id, bonus_type, currency, amount, wagering_multiplier,
			wagering_required, wagering_met, bonus_code, status, expires_at, created_at, used_at
		FROM bonus_transactions
		WHERE user_id = $1 AND currency = $2 AND status = 'ACTIVE' AND (expires_at IS NULL OR expires_at > NOW())
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID, currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bonuses []*model.BonusTransaction
	for rows.Next() {
		bt := &model.BonusTransaction{}
		err := rows.Scan(
			&bt.ID,
			&bt.UserID,
			&bt.TransactionID,
			&bt.BonusType,
			&bt.Currency,
			&bt.Amount,
			&bt.WageringMultiplier,
			&bt.WageringRequired,
			&bt.WageringMet,
			&bt.BonusCode,
			&bt.Status,
			&bt.ExpiresAt,
			&bt.CreatedAt,
			&bt.UsedAt,
		)
		if err != nil {
			return nil, err
		}
		bonuses = append(bonuses, bt)
	}

	return bonuses, nil
}

// UpdateBonusWagering updates bonus wagering progress
func (r *WalletRepository) UpdateBonusWagering(ctx context.Context, bonusID string, wageringMet int64) error {
	query := `
		UPDATE bonus_transactions
		SET wagering_met = $1
		WHERE id = $2
	`

	_, err := r.db.ExecContext(ctx, query, wageringMet, bonusID)
	return err
}
