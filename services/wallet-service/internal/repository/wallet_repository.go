package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gameengine/wallet-service/internal/model"
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

// CreateBet creates a new bet
func (r *WalletRepository) CreateBet(ctx context.Context, bet *model.Bet) error {
	if bet.BetID == "" {
		bet.BetID = uuid.New().String()
	}
	bet.PlacedAt = time.Now()

	query := `
		INSERT INTO bets (
			bet_id, user_id, game_id, bet_type, selection, odds, stake, potential_win, actual_win,
			settlement_type, status, placed_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := r.db.ExecContext(ctx, query,
		bet.BetID,
		bet.UserID,
		bet.GameID,
		bet.BetType,
		bet.Selection,
		bet.Odds,
		bet.Stake,
		bet.PotentialWin,
		bet.ActualWin,
		bet.SettlementType,
		bet.Status,
		bet.PlacedAt,
	)

	return err
}

// GetBetByID retrieves a bet by ID
func (r *WalletRepository) GetBetByID(ctx context.Context, betID string) (*model.Bet, error) {
	query := `
		SELECT bet_id, user_id, game_id, bet_type, selection, odds, stake, potential_win, actual_win,
			settlement_type, status, placed_at, settled_at
		FROM bets
		WHERE bet_id = $1
	`

	bet := &model.Bet{}
	err := r.db.QueryRowContext(ctx, query, betID).Scan(
		&bet.BetID,
		&bet.UserID,
		&bet.GameID,
		&bet.BetType,
		&bet.Selection,
		&bet.Odds,
		&bet.Stake,
		&bet.PotentialWin,
		&bet.ActualWin,
		&bet.SettlementType,
		&bet.Status,
		&bet.PlacedAt,
		&bet.SettledAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return bet, nil
}

// UpdateBetSettlement updates bet settlement
func (r *WalletRepository) UpdateBetSettlement(ctx context.Context, betID string, settlementType, status string, actualWin int64) error {
	now := time.Now()
	query := `
		UPDATE bets
		SET settlement_type = $1, status = $2, actual_win = $3, settled_at = $4
		WHERE bet_id = $5
	`

	_, err := r.db.ExecContext(ctx, query, settlementType, status, actualWin, now, betID)
	return err
}

// GetPendingBets retrieves pending bets for a user
func (r *WalletRepository) GetPendingBets(ctx context.Context, userID, gameID string, page, pageSize int) ([]*model.Bet, int, error) {
	baseQuery := `
		FROM bets
		WHERE status = 'TRANSACTION_STATUS_PENDING'
	`
	args := []interface{}{}
	argIndex := 1

	if userID != "" {
		baseQuery += fmt.Sprintf(" AND user_id = $%d", argIndex)
		args = append(args, userID)
		argIndex++
	}

	if gameID != "" {
		baseQuery += fmt.Sprintf(" AND game_id = $%d", argIndex)
		args = append(args, gameID)
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
		SELECT bet_id, user_id, game_id, bet_type, selection, odds, stake, potential_win, actual_win,
			settlement_type, status, placed_at, settled_at
		` + baseQuery + fmt.Sprintf(" ORDER BY placed_at DESC LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, pageSize, offset)

	rows, err := r.db.QueryContext(ctx, selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var bets []*model.Bet
	for rows.Next() {
		bet := &model.Bet{}
		err := rows.Scan(
			&bet.BetID,
			&bet.UserID,
			&bet.GameID,
			&bet.BetType,
			&bet.Selection,
			&bet.Odds,
			&bet.Stake,
			&bet.PotentialWin,
			&bet.ActualWin,
			&bet.SettlementType,
			&bet.Status,
			&bet.PlacedAt,
			&bet.SettledAt,
		)
		if err != nil {
			return nil, 0, err
		}
		bets = append(bets, bet)
	}

	return bets, total, nil
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

// ExpireBonuses marks expired bonuses
func (r *WalletRepository) ExpireBonuses(ctx context.Context) (int64, error) {
	query := `
		UPDATE bonus_transactions
		SET status = 'EXPIRED'
		WHERE status = 'ACTIVE' AND expires_at < NOW()
	`

	result, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// GetDailyWithdrawalTotal gets total withdrawals for today
func (r *WalletRepository) GetDailyWithdrawalTotal(ctx context.Context, userID, currency string) (int64, error) {
	query := `
		SELECT COALESCE(SUM(net_amount), 0)
		FROM transactions
		WHERE user_id = $1 
			AND currency = $2 
			AND type = 'TRANSACTION_TYPE_WITHDRAWAL'
			AND status = 'TRANSACTION_STATUS_COMPLETED'
			AND DATE(processed_at) = CURRENT_DATE
	`

	var total int64
	err := r.db.QueryRowContext(ctx, query, userID, currency).Scan(&total)
	return total, err
}

// BeginTx begins a new transaction
func (r *WalletRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}

// WithTx executes a function within a transaction
func (r *WalletRepository) WithTx(ctx context.Context, fn func(repo *WalletRepositoryWithTx) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	repoWithTx := &WalletRepositoryWithTx{tx: tx}

	if err := fn(repoWithTx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}

// WalletRepositoryWithTx wraps WalletRepository for transaction support
type WalletRepositoryWithTx struct {
	tx *sql.Tx
}

func (r *WalletRepositoryWithTx) GetWalletByUserIDAndType(ctx context.Context, userID, currency, balanceType string) (*model.Wallet, error) {
	query := `
		SELECT id, user_id, currency, balance_type, amount, locked_amount, version, created_at, updated_at
		FROM wallets
		WHERE user_id = $1 AND currency = $2 AND balance_type = $3
		FOR UPDATE
	`

	wallet := &model.Wallet{}
	err := r.tx.QueryRowContext(ctx, query, userID, currency, balanceType).Scan(
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

func (r *WalletRepositoryWithTx) CreateWallet(ctx context.Context, wallet *model.Wallet) error {
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

	_, err := r.tx.ExecContext(ctx, query,
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

func (r *WalletRepositoryWithTx) UpdateWalletAmount(ctx context.Context, walletID string, newAmount, newLocked int64, version int) error {
	query := `
		UPDATE wallets
		SET amount = $1, locked_amount = $2, version = version + 1, updated_at = NOW()
		WHERE id = $3 AND version = $4
	`

	result, err := r.tx.ExecContext(ctx, query, newAmount, newLocked, walletID, version)
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

func (r *WalletRepositoryWithTx) CreateTransaction(ctx context.Context, tx *model.Transaction) error {
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

	_, err := r.tx.ExecContext(ctx, query,
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

func (r *WalletRepositoryWithTx) CreateBet(ctx context.Context, bet *model.Bet) error {
	if bet.BetID == "" {
		bet.BetID = uuid.New().String()
	}
	bet.PlacedAt = time.Now()

	query := `
		INSERT INTO bets (
			bet_id, user_id, game_id, bet_type, selection, odds, stake, potential_win, actual_win,
			settlement_type, status, placed_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`

	_, err := r.tx.ExecContext(ctx, query,
		bet.BetID,
		bet.UserID,
		bet.GameID,
		bet.BetType,
		bet.Selection,
		bet.Odds,
		bet.Stake,
		bet.PotentialWin,
		bet.ActualWin,
		bet.SettlementType,
		bet.Status,
		bet.PlacedAt,
	)

	return err
}

func (r *WalletRepositoryWithTx) UpdateBetSettlement(ctx context.Context, betID string, settlementType, status string, actualWin int64) error {
	now := time.Now()
	query := `
		UPDATE bets
		SET settlement_type = $1, status = $2, actual_win = $3, settled_at = $4
		WHERE bet_id = $5
	`

	_, err := r.tx.ExecContext(ctx, query, settlementType, status, actualWin, now, betID)
	return err
}

func (r *WalletRepositoryWithTx) CreateBonusTransaction(ctx context.Context, bt *model.BonusTransaction) error {
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

	_, err := r.tx.ExecContext(ctx, query,
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

// UpdateTransactionStatus updates transaction status
func (r *WalletRepositoryWithTx) UpdateTransactionStatus(ctx context.Context, txID, status string, processedAt *time.Time) error {
	query := `
		UPDATE transactions
		SET status = $1, processed_at = $2
		WHERE transaction_id = $3
	`

	_, err := r.tx.ExecContext(ctx, query, status, processedAt, txID)
	return err
}

// UpdateBonusWagering updates bonus wagering progress
func (r *WalletRepositoryWithTx) UpdateBonusWagering(ctx context.Context, bonusID string, wageringMet int64) error {
	query := `
		UPDATE bonus_transactions
		SET wagering_met = $1
		WHERE id = $2
	`

	_, err := r.tx.ExecContext(ctx, query, wageringMet, bonusID)
	return err
}
