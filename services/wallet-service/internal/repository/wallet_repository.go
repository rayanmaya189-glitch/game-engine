package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/game_engine/wallet-service/internal/model"
	"github.com/google/uuid"
)

// WalletRepository handles database operations for wallets
type WalletRepository struct {
	db *sql.DB
}

// NewWalletRepository creates a new wallet repository
func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

// GetWalletByUserIDAndType retrieves a wallet by user ID and balance type
func (r *WalletRepository) GetWalletByUserIDAndType(ctx context.Context, userID, currency, balanceType string) (*model.Wallet, error) {
	query := `
		SELECT id, user_id, currency, balance_type, amount, locked_amount, version, created_at, updated_at
		FROM wallets
		WHERE user_id = $1 AND currency = $2 AND balance_type = $3
		FOR UPDATE
	`

	wallet := &model.Wallet{}
	err := r.db.QueryRowContext(ctx, query, userID, currency, balanceType).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.Currency,
		&wallet.BalanceType,
		&wallet.Amount,
		&wallet.LockedAmount,
		&wallet.Version,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

// GetWalletsByUserID retrieves all wallets for a user
func (r *WalletRepository) GetWalletsByUserID(ctx context.Context, userID string) ([]*model.Wallet, error) {
	query := `
		SELECT id, user_id, currency, balance_type, amount, locked_amount, version, created_at, updated_at
		FROM wallets
		WHERE user_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []*model.Wallet
	for rows.Next() {
		wallet := &model.Wallet{}
		err := rows.Scan(
			&wallet.ID,
			&wallet.UserID,
			&wallet.Currency,
			&wallet.BalanceType,
			&wallet.Amount,
			&wallet.LockedAmount,
			&wallet.Version,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}

	return wallets, nil
}

// CreateWallet creates a new wallet
func (r *WalletRepository) CreateWallet(ctx context.Context, wallet *model.Wallet) error {
	if wallet.ID == "" {
		wallet.ID = uuid.New().String()
	}
	wallet.CreatedAt = time.Now()
	wallet.UpdatedAt = time.Now()
	wallet.Version = 1

	query := `
		INSERT INTO wallets (id, user_id, currency, balance_type, amount, locked_amount, version, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.db.ExecContext(ctx, query,
		wallet.ID,
		wallet.UserID,
		wallet.Currency,
		wallet.BalanceType,
		wallet.Amount,
		wallet.LockedAmount,
		wallet.Version,
		wallet.CreatedAt,
		wallet.UpdatedAt,
	)

	return err
}

// UpdateWalletAmount updates wallet amount with optimistic locking
func (r *WalletRepository) UpdateWalletAmount(ctx context.Context, walletID string, newAmount, newLocked int64, version int) error {
	query := `
		UPDATE wallets
		SET amount = $1, locked_amount = $2, version = version + 1, updated_at = NOW()
		WHERE id = $3 AND version = $4
	`

	result, err := r.db.ExecContext(ctx, query, newAmount, newLocked, walletID, version)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("optimistic lock failed for wallet %s", walletID)
	}

	return nil
}
