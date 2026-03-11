package repository

import (
	"context"
	"fmt"

	"github.com/game-engine/jackpot-service/internal/config"
	"github.com/game-engine/jackpot-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type JackpotRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewJackpotRepository(db *pgxpool.Pool, redis *redis.Client) *JackpotRepository {
	return &JackpotRepository{db: db, redis: redis}
}

func NewPostgresDB(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := cfg.ConnectionString()
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
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

func (r *JackpotRepository) ListJackpots(ctx context.Context, status string) ([]model.Jackpot, error) {
	query := `SELECT jackpot_id, name, description, current_amount, min_bet, max_bet, status, starts_at, ends_at FROM jackpots WHERE 1=1`
	args := []interface{}{}
	if status != "" {
		query += " AND status = $1"
		args = append(args, status)
	}
	query += " ORDER BY created_at DESC"

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list jackpots: %w", err)
	}
	defer rows.Close()

	var jackpots []model.Jackpot
	for rows.Next() {
		var j model.Jackpot
		if err := rows.Scan(&j.JackpotID, &j.Name, &j.Description, &j.CurrentAmount, &j.MinBet, &j.MaxBet, &j.Status, &j.StartsAt, &j.EndsAt); err != nil {
			return nil, fmt.Errorf("failed to scan jackpot: %w", err)
		}
		jackpots = append(jackpots, j)
	}
	return jackpots, nil
}

func (r *JackpotRepository) GetJackpot(ctx context.Context, jackpotID string) (*model.Jackpot, error) {
	var j model.Jackpot
	err := r.db.QueryRow(ctx, `SELECT jackpot_id, name, description, current_amount, min_bet, max_bet, status, starts_at, ends_at FROM jackpots WHERE jackpot_id = $1`, jackpotID).Scan(&j.JackpotID, &j.Name, &j.Description, &j.CurrentAmount, &j.MinBet, &j.MaxBet, &j.Status, &j.StartsAt, &j.EndsAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get jackpot: %w", err)
	}
	return &j, nil
}

func (r *JackpotRepository) GetWinners(ctx context.Context, jackpotID string, limit int32) ([]model.Winner, error) {
	rows, err := r.db.Query(ctx, `SELECT winner_id, username, amount, won_at FROM jackpot_winners WHERE jackpot_id = $1 ORDER BY won_at DESC LIMIT $2`, jackpotID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get winners: %w", err)
	}
	defer rows.Close()

	var winners []model.Winner
	for rows.Next() {
		var w model.Winner
		if err := rows.Scan(&w.WinnerID, &w.Username, &w.Amount, &w.WonAt); err != nil {
			return nil, fmt.Errorf("failed to scan winner: %w", err)
		}
		winners = append(winners, w)
	}
	return winners, nil
}

func (r *JackpotRepository) JoinJackpot(ctx context.Context, jackpotID, userID string, betAmount float64) (bool, string, error) {
	var j model.Jackpot
	err := r.db.QueryRow(ctx, `SELECT jackpot_id, min_bet, max_bet, status FROM jackpots WHERE jackpot_id = $1`, jackpotID).Scan(&j.JackpotID, &j.MinBet, &j.MaxBet, &j.Status)
	if err != nil {
		return false, "Jackpot not found", nil
	}

	if j.Status != "active" {
		return false, "Jackpot is not active", nil
	}

	if betAmount < j.MinBet || betAmount > j.MaxBet {
		return false, "Bet amount is outside the allowed range", nil
	}

	// Add player to jackpot participants
	_, err = r.db.Exec(ctx, `INSERT INTO jackpot_participants (jackpot_id, user_id, bet_amount, joined_at) VALUES ($1, $2, $3, NOW()) ON CONFLICT DO NOTHING`, jackpotID, userID, betAmount)
	if err != nil {
		return false, "Failed to join jackpot", fmt.Errorf("failed to join jackpot: %w", err)
	}

	return true, "Successfully joined jackpot", nil
}

func (r *JackpotRepository) GetJackpotHistory(ctx context.Context, userID string, page, limit int32) ([]model.JackpotHistoryEntry, int, error) {
	offset := (page - 1) * limit

	var total int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM jackpot_participants WHERE user_id = $1`, userID).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count jackpot history: %w", err)
	}

	rows, err := r.db.Query(ctx, `SELECT jh.jackpot_id, j.name, jh.amount, jh.result, jh.played_at FROM jackpot_history jh JOIN jackpots j ON jh.jackpot_id = j.jackpot_id WHERE jh.user_id = $1 ORDER BY jh.played_at DESC LIMIT $2 OFFSET $3`, userID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get jackpot history: %w", err)
	}
	defer rows.Close()

	var entries []model.JackpotHistoryEntry
	for rows.Next() {
		var e model.JackpotHistoryEntry
		if err := rows.Scan(&e.JackpotID, &e.JackpotName, &e.Amount, &e.Result, &e.PlayedAt); err != nil {
			return nil, 0, fmt.Errorf("failed to scan history entry: %w", err)
		}
		entries = append(entries, e)
	}
	return entries, total, nil
}
