package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/game_engine/live-dealer-service/internal/config"
	"github.com/game_engine/live-dealer-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type LiveDealerRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewLiveDealerRepository(db *pgxpool.Pool, redis *redis.Client) *LiveDealerRepository {
	return &LiveDealerRepository{db: db, redis: redis}
}

func NewPostgresDB(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := cfg.ConnectionString()
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}
	if cfg.MaxConnections > 0 {
		poolConfig.MaxConns = int32(cfg.MaxConnections)
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

// --- Table persistence ---

func (r *LiveDealerRepository) CreateTable(ctx context.Context, table *model.Table) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO live_dealer_tables (table_id, game_type, dealer_id, dealer_name, status, min_bet, max_bet, current_seat, max_seats, stake_limit, deck_count, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`, table.TableID, table.GameType, table.DealerID, table.DealerName, table.Status,
		table.MinBet, table.MaxBet, table.CurrentSeat, table.MaxSeats, table.StakeLimit,
		table.DeckCount, table.CreatedAt, table.UpdatedAt)
	return err
}

func (r *LiveDealerRepository) GetTable(ctx context.Context, tableID string) (*model.Table, error) {
	var t model.Table
	err := r.db.QueryRow(ctx, `
		SELECT table_id, game_type, dealer_id, dealer_name, status, min_bet, max_bet, current_seat, max_seats, stake_limit, deck_count, created_at, updated_at
		FROM live_dealer_tables WHERE table_id = $1
	`, tableID).Scan(&t.TableID, &t.GameType, &t.DealerID, &t.DealerName, &t.Status,
		&t.MinBet, &t.MaxBet, &t.CurrentSeat, &t.MaxSeats, &t.StakeLimit,
		&t.DeckCount, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get table: %w", err)
	}
	return &t, nil
}

func (r *LiveDealerRepository) ListTables(ctx context.Context, gameType, status string) ([]*model.Table, error) {
	query := `SELECT table_id, game_type, dealer_id, dealer_name, status, min_bet, max_bet, current_seat, max_seats, stake_limit, deck_count, created_at, updated_at FROM live_dealer_tables WHERE 1=1`
	args := []interface{}{}
	argIdx := 1

	if gameType != "" {
		query += fmt.Sprintf(" AND game_type = $%d", argIdx)
		args = append(args, gameType)
		argIdx++
	}
	if status != "" {
		query += fmt.Sprintf(" AND status = $%d", argIdx)
		args = append(args, status)
		argIdx++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	defer rows.Close()

	var tables []*model.Table
	for rows.Next() {
		var t model.Table
		if err := rows.Scan(&t.TableID, &t.GameType, &t.DealerID, &t.DealerName, &t.Status,
			&t.MinBet, &t.MaxBet, &t.CurrentSeat, &t.MaxSeats, &t.StakeLimit,
			&t.DeckCount, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tables = append(tables, &t)
	}
	return tables, nil
}

func (r *LiveDealerRepository) UpdateTable(ctx context.Context, table *model.Table) error {
	_, err := r.db.Exec(ctx, `
		UPDATE live_dealer_tables SET game_type=$2, dealer_id=$3, dealer_name=$4, status=$5, min_bet=$6, max_bet=$7, current_seat=$8, max_seats=$9, stake_limit=$10, deck_count=$11, updated_at=$12
		WHERE table_id = $1
	`, table.TableID, table.GameType, table.DealerID, table.DealerName, table.Status,
		table.MinBet, table.MaxBet, table.CurrentSeat, table.MaxSeats, table.StakeLimit,
		table.DeckCount, time.Now())
	return err
}

// --- Dealer persistence ---

func (r *LiveDealerRepository) CreateDealer(ctx context.Context, dealer *model.Dealer) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO live_dealer_dealers (dealer_id, name, avatar, language, status, table_id, shift_start, shift_end)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, dealer.DealerID, dealer.Name, dealer.Avatar, dealer.Language, dealer.Status,
		dealer.TableID, dealer.ShiftStart, dealer.ShiftEnd)
	return err
}

func (r *LiveDealerRepository) GetDealer(ctx context.Context, dealerID string) (*model.Dealer, error) {
	var d model.Dealer
	err := r.db.QueryRow(ctx, `
		SELECT dealer_id, name, avatar, language, status, table_id, shift_start, shift_end
		FROM live_dealer_dealers WHERE dealer_id = $1
	`, dealerID).Scan(&d.DealerID, &d.Name, &d.Avatar, &d.Language, &d.Status,
		&d.TableID, &d.ShiftStart, &d.ShiftEnd)
	if err != nil {
		return nil, fmt.Errorf("failed to get dealer: %w", err)
	}
	return &d, nil
}

func (r *LiveDealerRepository) UpdateDealer(ctx context.Context, dealer *model.Dealer) error {
	_, err := r.db.Exec(ctx, `
		UPDATE live_dealer_dealers SET name=$2, avatar=$3, language=$4, status=$5, table_id=$6, shift_start=$7, shift_end=$8
		WHERE dealer_id = $1
	`, dealer.DealerID, dealer.Name, dealer.Avatar, dealer.Language, dealer.Status,
		dealer.TableID, dealer.ShiftStart, dealer.ShiftEnd)
	return err
}

func (r *LiveDealerRepository) ListDealers(ctx context.Context, status string) ([]*model.Dealer, error) {
	query := `SELECT dealer_id, name, avatar, language, status, table_id, shift_start, shift_end FROM live_dealer_dealers WHERE 1=1`
	args := []interface{}{}
	if status != "" {
		query += " AND status = $1"
		args = append(args, status)
	}
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list dealers: %w", err)
	}
	defer rows.Close()

	var dealers []*model.Dealer
	for rows.Next() {
		var d model.Dealer
		if err := rows.Scan(&d.DealerID, &d.Name, &d.Avatar, &d.Language, &d.Status,
			&d.TableID, &d.ShiftStart, &d.ShiftEnd); err != nil {
			return nil, err
		}
		dealers = append(dealers, &d)
	}
	return dealers, nil
}
