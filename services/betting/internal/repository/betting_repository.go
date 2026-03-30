package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Selection represents a single selection in a bet (mirrors service.Selection)
type Selection struct {
	ID         string     `json:"id"`
	EventID    string     `json:"event_id"`
	OutcomeID  string     `json:"outcome_id"`
	Odds       float64    `json:"odds"`
	OddsFormat string     `json:"odds_format"`
	Status     string     `json:"status"`
	Result     string     `json:"result"`
	SettledAt  *time.Time `json:"settled_at,omitempty"`
}

// Bet represents a bet entity for persistence
type Bet struct {
	ID           string      `json:"id"`
	UserID       string      `json:"user_id"`
	Type         string      `json:"type"`
	Stake        int64       `json:"stake"`
	Odds         float64     `json:"odds"`
	PotentialWin int64       `json:"potential_win"`
	Status       string      `json:"status"`
	Selections   []Selection `json:"selections"`
	SystemType   string      `json:"system_type,omitempty"`
	Currency     string      `json:"currency"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	SettledAt    *time.Time  `json:"settled_at,omitempty"`
	VoidReason   string      `json:"void_reason,omitempty"`
}

// BettingRepository defines the interface for bet persistence
type BettingRepository interface {
	GetBetByID(ctx context.Context, betID string) (*Bet, error)
	GetBetHistory(ctx context.Context, userID string, limit, offset int) ([]*Bet, error)
	GetOpenBets(ctx context.Context, userID string) ([]*Bet, error)
	SaveBet(ctx context.Context, bet *Bet) error
	UpdateBet(ctx context.Context, bet *Bet) error
}

// PostgresBettingRepository implements BettingRepository using PostgreSQL
type PostgresBettingRepository struct {
	pool *pgxpool.Pool
}

// NewPostgresBettingRepository creates a new PostgreSQL betting repository
func NewPostgresBettingRepository(pool *pgxpool.Pool) *PostgresBettingRepository {
	return &PostgresBettingRepository{pool: pool}
}

func (r *PostgresBettingRepository) scanBet(row pgx.Row) (*Bet, error) {
	var b Bet
	var selectionsJSON []byte
	var settledAt *time.Time
	err := row.Scan(
		&b.ID, &b.UserID, &b.Type, &b.Stake, &b.Odds,
		&b.PotentialWin, &b.Status, &selectionsJSON,
		&b.SystemType, &b.Currency, &b.CreatedAt,
		&b.UpdatedAt, &settledAt, &b.VoidReason,
	)
	if err != nil {
		return nil, err
	}
	b.SettledAt = settledAt
	if err := json.Unmarshal(selectionsJSON, &b.Selections); err != nil {
		return nil, fmt.Errorf("failed to unmarshal selections: %w", err)
	}
	return &b, nil
}

func (r *PostgresBettingRepository) GetBetByID(ctx context.Context, betID string) (*Bet, error) {
	row := r.pool.QueryRow(ctx, `
		SELECT id, user_id, type, stake, odds, potential_win, status,
			selections, system_type, currency, created_at, updated_at, settled_at, void_reason
		FROM bets WHERE id = $1
	`, betID)
	bet, err := r.scanBet(row)
	if err != nil {
		return nil, fmt.Errorf("bet not found: %w", err)
	}
	return bet, nil
}

func (r *PostgresBettingRepository) GetBetHistory(ctx context.Context, userID string, limit, offset int) ([]*Bet, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, user_id, type, stake, odds, potential_win, status,
			selections, system_type, currency, created_at, updated_at, settled_at, void_reason
		FROM bets WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query bet history: %w", err)
	}
	defer rows.Close()

	var bets []*Bet
	for rows.Next() {
		bet, err := r.scanBet(rows)
		if err != nil {
			return nil, err
		}
		bets = append(bets, bet)
	}
	return bets, nil
}

func (r *PostgresBettingRepository) GetOpenBets(ctx context.Context, userID string) ([]*Bet, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, user_id, type, stake, odds, potential_win, status,
			selections, system_type, currency, created_at, updated_at, settled_at, void_reason
		FROM bets
		WHERE user_id = $1 AND status IN ('placed', 'accepted', 'active')
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query open bets: %w", err)
	}
	defer rows.Close()

	var bets []*Bet
	for rows.Next() {
		bet, err := r.scanBet(rows)
		if err != nil {
			return nil, err
		}
		bets = append(bets, bet)
	}
	return bets, nil
}

func (r *PostgresBettingRepository) SaveBet(ctx context.Context, bet *Bet) error {
	selectionsJSON, err := json.Marshal(bet.Selections)
	if err != nil {
		return fmt.Errorf("failed to marshal selections: %w", err)
	}

	_, err = r.pool.Exec(ctx, `
		INSERT INTO bets (id, user_id, type, stake, odds, potential_win, status,
			selections, system_type, currency, created_at, updated_at, settled_at, void_reason)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	`, bet.ID, bet.UserID, bet.Type, bet.Stake, bet.Odds,
		bet.PotentialWin, bet.Status, selectionsJSON,
		bet.SystemType, bet.Currency, bet.CreatedAt,
		bet.UpdatedAt, bet.SettledAt, bet.VoidReason)
	if err != nil {
		return fmt.Errorf("failed to save bet: %w", err)
	}
	return nil
}

func (r *PostgresBettingRepository) UpdateBet(ctx context.Context, bet *Bet) error {
	selectionsJSON, err := json.Marshal(bet.Selections)
	if err != nil {
		return fmt.Errorf("failed to marshal selections: %w", err)
	}

	result, err := r.pool.Exec(ctx, `
		UPDATE bets SET user_id = $2, type = $3, stake = $4, odds = $5, potential_win = $6,
			status = $7, selections = $8, system_type = $9, currency = $10,
			updated_at = $11, settled_at = $12, void_reason = $13
		WHERE id = $1
	`, bet.ID, bet.UserID, bet.Type, bet.Stake, bet.Odds,
		bet.PotentialWin, bet.Status, selectionsJSON,
		bet.SystemType, bet.Currency, bet.UpdatedAt,
		bet.SettledAt, bet.VoidReason)
	if err != nil {
		return fmt.Errorf("failed to update bet: %w", err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("bet not found: %s", bet.ID)
	}
	return nil
}
