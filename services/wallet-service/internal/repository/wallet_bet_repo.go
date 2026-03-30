package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/game_engine/wallet-service/internal/model"
	"github.com/google/uuid"
)

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
