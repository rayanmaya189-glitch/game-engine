package repository

import (
	"context"
	"fmt"

	"github.com/game_engine/sports-betting-service/internal/config"
	"github.com/game_engine/sports-betting-service/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type SportsRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewSportsRepository(db *pgxpool.Pool, redis *redis.Client) *SportsRepository {
	return &SportsRepository{db: db, redis: redis}
}

func NewPostgresDB(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := cfg.ConnectionString()
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return pool, nil
}

func NewRedisClient(cfg config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return client, nil
}

func (r *SportsRepository) GetSports(ctx context.Context) ([]model.Sport, error) {
	rows, err := r.db.Query(ctx, `
		SELECT sport_id, name, icon, status, sort_order FROM sports WHERE status = 'active' ORDER BY sort_order
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sports []model.Sport
	for rows.Next() {
		var s model.Sport
		if err := rows.Scan(&s.SportID, &s.Name, &s.Icon, &s.Status, &s.SortOrder); err != nil {
			return nil, err
		}
		sports = append(sports, s)
	}
	return sports, nil
}

func (r *SportsRepository) GetMarkets(ctx context.Context, eventID string) ([]model.Market, error) {
	rows, err := r.db.Query(ctx, `
		SELECT market_id, event_id, name, status, home_odds, draw_odds, away_odds
		FROM sports_markets WHERE event_id = $1 AND status IN ('open', 'closed')
	`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markets []model.Market
	for rows.Next() {
		var m model.Market
		if err := rows.Scan(&m.MarketID, &m.EventID, &m.Name, &m.Status, &m.HomeOdds, &m.DrawOdds, &m.AwayOdds); err != nil {
			return nil, err
		}
		markets = append(markets, m)
	}
	return markets, nil
}

func (r *SportsRepository) PlaceBet(ctx context.Context, bet *model.Bet) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO sports_bets (bet_id, user_id, event_id, market_id, selection, stake, odds, potential_win, status, placed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
	`, bet.BetID, bet.UserID, bet.EventID, bet.MarketID, bet.Selection, bet.Stake, bet.Odds, bet.PotentialWin, bet.Status)
	return err
}

func (r *SportsRepository) GetUserBets(ctx context.Context, userID string, limit, offset int) ([]model.Bet, int, error) {
	var total int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM sports_bets WHERE user_id = $1`, userID).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count bets: %w", err)
	}

	rows, err := r.db.Query(ctx, `
		SELECT bet_id, user_id, event_id, market_id, selection, stake, odds, potential_win, status, placed_at, settled_at
		FROM sports_bets WHERE user_id = $1 ORDER BY placed_at DESC LIMIT $2 OFFSET $3
	`, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var bets []model.Bet
	for rows.Next() {
		var b model.Bet
		if err := rows.Scan(&b.BetID, &b.UserID, &b.EventID, &b.MarketID, &b.Selection, &b.Stake, &b.Odds, &b.PotentialWin, &b.Status, &b.PlacedAt, &b.SettledAt); err != nil {
			return nil, 0, err
		}
		bets = append(bets, b)
	}
	return bets, total, nil
}

func (r *SportsRepository) SettleBet(ctx context.Context, betID string, status string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE sports_bets SET status = $1, settled_at = NOW() WHERE bet_id = $2
	`, status, betID)
	return err
}

func (r *SportsRepository) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return r.db.Begin(ctx)
}

func (r *SportsRepository) PlaceBetTx(ctx context.Context, tx pgx.Tx, bet *model.Bet) error {
	_, err := tx.Exec(ctx, `
		INSERT INTO sports_bets (bet_id, user_id, event_id, market_id, selection, stake, odds, potential_win, status, placed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
	`, bet.BetID, bet.UserID, bet.EventID, bet.MarketID, bet.Selection, bet.Stake, bet.Odds, bet.PotentialWin, bet.Status)
	return err
}

// GetBetByID returns a bet by ID
func (r *SportsRepository) GetBetByID(ctx context.Context, betID string) (*model.Bet, error) {
	var b model.Bet
	err := r.db.QueryRow(ctx, `
		SELECT bet_id, user_id, event_id, market_id, selection, stake, odds, potential_win, status, placed_at, settled_at
		FROM sports_bets WHERE bet_id = $1
	`, betID).Scan(&b.BetID, &b.UserID, &b.EventID, &b.MarketID, &b.Selection, &b.Stake, &b.Odds, &b.PotentialWin, &b.Status, &b.PlacedAt, &b.SettledAt)
	if err != nil {
		return nil, fmt.Errorf("bet not found: %w", err)
	}
	return &b, nil
}

// GetPendingBetsForEvent returns all pending bets for an event
func (r *SportsRepository) GetPendingBetsForEvent(ctx context.Context, eventID string) ([]model.Bet, error) {
	rows, err := r.db.Query(ctx, `
		SELECT bet_id, user_id, event_id, market_id, selection, stake, odds, potential_win, status, placed_at
		FROM sports_bets WHERE event_id = $1 AND status = 'pending'
	`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bets []model.Bet
	for rows.Next() {
		var b model.Bet
		if err := rows.Scan(&b.BetID, &b.UserID, &b.EventID, &b.MarketID, &b.Selection, &b.Stake, &b.Odds, &b.PotentialWin, &b.Status, &b.PlacedAt); err != nil {
			return nil, err
		}
		bets = append(bets, b)
	}
	return bets, nil
}

// UpdateBetStatus updates bet status and settled amount
func (r *SportsRepository) UpdateBetStatus(ctx context.Context, betID string, status model.BetStatus, winAmount float64) error {
	if status == model.BetStatusWon || status == model.BetStatusCashOut {
		_, err := r.db.Exec(ctx, `
			UPDATE sports_bets SET status = $1, settled_at = NOW() WHERE bet_id = $2
		`, status, betID)
		return err
	}
	_, err := r.db.Exec(ctx, `
		UPDATE sports_bets SET status = $1, settled_at = NOW() WHERE bet_id = $2
	`, status, betID)
	return err
}

// SaveCashOut saves a cash out record
func (r *SportsRepository) SaveCashOut(ctx context.Context, cashOut *model.CashOut) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO sports_cashouts (cash_out_id, bet_id, user_id, original_stake, original_odds, current_odds, cash_out_amount, profit, status, requested_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
	`, cashOut.CashOutID, cashOut.BetID, cashOut.UserID, cashOut.OriginalStake, cashOut.OriginalOdds, cashOut.CurrentOdds, cashOut.CashOutAmount, cashOut.Profit, cashOut.Status)
	return err
}

// UpdateCashOutStatus updates cash out status
func (r *SportsRepository) UpdateCashOutStatus(ctx context.Context, cashOutID, status string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE sports_cashouts SET status = $1, completed_at = NOW() WHERE cash_out_id = $2
	`, status, cashOutID)
	return err
}

// GetUserCashOuts returns all cash outs for a user
func (r *SportsRepository) GetUserCashOuts(ctx context.Context, userID string, limit, offset int) ([]model.CashOut, int, error) {
	var total int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM sports_cashouts WHERE user_id = $1`, userID).Scan(&total); err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(ctx, `
		SELECT cash_out_id, bet_id, user_id, original_stake, original_odds, current_odds, cash_out_amount, profit, status, requested_at, completed_at
		FROM sports_cashouts WHERE user_id = $1 ORDER BY requested_at DESC LIMIT $2 OFFSET $3
	`, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var cashOuts []model.CashOut
	for rows.Next() {
		var c model.CashOut
		if err := rows.Scan(&c.CashOutID, &c.BetID, &c.UserID, &c.OriginalStake, &c.OriginalOdds, &c.CurrentOdds, &c.CashOutAmount, &c.Profit, &c.Status, &c.RequestedAt, &c.CompletedAt); err != nil {
			return nil, 0, err
		}
		cashOuts = append(cashOuts, c)
	}
	return cashOuts, total, nil
}
