package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/game_engine/winners-showcase-service/internal/config"
	"github.com/game_engine/winners-showcase-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type WinnersRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewWinnersRepository(db *pgxpool.Pool, redis *redis.Client) *WinnersRepository {
	return &WinnersRepository{db: db, redis: redis}
}

func NewPostgresDB(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
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
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return client, nil
}

// GetRecentWinners retrieves recent winners from Redis cache or DB
func (r *WinnersRepository) GetRecentWinners(ctx context.Context, limit int) ([]model.Winner, error) {
	// Try Redis first
	key := "winners:recent"
	data, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		var winners []model.Winner
		if json.Unmarshal([]byte(data), &winners) == nil {
			return winners, nil
		}
	}

	// Fall back to database
	return r.getRecentWinnersFromDB(ctx, limit)
}

func (r *WinnersRepository) getRecentWinnersFromDB(ctx context.Context, limit int) ([]model.Winner, error) {
	query := `
		SELECT id, user_id, username, win_amount, currency, game_type, game_name, 
		       win_type, multiplier, timestamp
		FROM winners
		WHERE display_on_feed = true
		ORDER BY timestamp DESC
		LIMIT $1
	`

	rows, err := r.db.Query(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var winners []model.Winner
	for rows.Next() {
		var w model.Winner
		if err := rows.Scan(&w.ID, &w.UserID, &w.Username, &w.WinAmount, &w.Currency,
			&w.GameType, &w.GameName, &w.WinType, &w.Multiplier, &w.Timestamp); err != nil {
			return nil, err
		}
		w.DisplayName = w.Username
		winners = append(winners, w)
	}
	return winners, nil
}

// GetBigWins retrieves wins above threshold
func (r *WinnersRepository) GetBigWins(ctx context.Context, threshold float64, limit int) ([]model.Winner, error) {
	key := fmt.Sprintf("winners:big:%f", threshold)
	data, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		var winners []model.Winner
		if json.Unmarshal([]byte(data), &winners) == nil {
			return winners, nil
		}
	}

	return r.getBigWinsFromDB(ctx, threshold, limit)
}

func (r *WinnersRepository) getBigWinsFromDB(ctx context.Context, threshold float64, limit int) ([]model.Winner, error) {
	query := `
		SELECT id, user_id, username, win_amount, currency, game_type, game_name, 
		       win_type, multiplier, timestamp
		FROM winners
		WHERE win_amount >= $1 AND display_on_feed = true
		ORDER BY win_amount DESC
		LIMIT $2
	`

	rows, err := r.db.Query(ctx, query, threshold, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var winners []model.Winner
	for rows.Next() {
		var w model.Winner
		if err := rows.Scan(&w.ID, &w.UserID, &w.Username, &w.WinAmount, &w.Currency,
			&w.GameType, &w.GameName, &w.WinType, &w.Multiplier, &w.Timestamp); err != nil {
			return nil, err
		}
		w.DisplayName = w.Username
		winners = append(winners, w)
	}
	return winners, nil
}

// GetJackpotWinners retrieves jackpot wins
func (r *WinnersRepository) GetJackpotWinners(ctx context.Context, threshold float64, limit int) ([]model.Winner, error) {
	key := fmt.Sprintf("winners:jackpot:%f", threshold)
	data, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		var winners []model.Winner
		if json.Unmarshal([]byte(data), &winners) == nil {
			return winners, nil
		}
	}

	return r.getJackpotWinnersFromDB(ctx, threshold, limit)
}

func (r *WinnersRepository) getJackpotWinnersFromDB(ctx context.Context, threshold float64, limit int) ([]model.Winner, error) {
	query := `
		SELECT id, user_id, username, win_amount, currency, game_type, game_name, 
		       win_type, multiplier, timestamp
		FROM winners
		WHERE (win_type = 'jackpot' OR win_type = 'progressive') 
		  AND win_amount >= $1
		ORDER BY win_amount DESC
		LIMIT $2
	`

	rows, err := r.db.Query(ctx, query, threshold, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var winners []model.Winner
	for rows.Next() {
		var w model.Winner
		if err := rows.Scan(&w.ID, &w.UserID, &w.Username, &w.WinAmount, &w.Currency,
			&w.GameType, &w.GameName, &w.WinType, &w.Multiplier, &w.Timestamp); err != nil {
			return nil, err
		}
		w.DisplayName = w.Username
		winners = append(winners, w)
	}
	return winners, nil
}

// RecordWin records a new win and updates Redis cache
func (r *WinnersRepository) RecordWin(ctx context.Context, req model.RecordWinRequest, threshold float64, jackpotThreshold float64) error {
	// Determine win type
	winType := model.WinTypeRegular
	if req.WinAmount >= jackpotThreshold {
		winType = model.WinTypeJackpot
	} else if req.WinAmount >= threshold {
		winType = model.WinTypeBig
	}

	query := `
		INSERT INTO winners (user_id, username, win_amount, currency, game_type, game_name, 
		                    win_type, multiplier, timestamp, display_on_feed)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), true)
	`
	_, err := r.db.Exec(ctx, query, req.UserID, req.Username, req.WinAmount, req.Currency,
		req.GameType, req.GameName, winType, req.Multiplier)

	// Invalidate cache
	r.redis.Del(ctx, "winners:recent")

	return err
}

// GetPrivacySettings retrieves player's privacy settings
func (r *WinnersRepository) GetPrivacySettings(ctx context.Context, userID string) (*model.PrivacySettings, error) {
	query := `
		SELECT user_id, anonymize_name, show_on_leaderboard, show_on_jackpot_list, opt_out_of_showcase
		FROM winner_privacy_settings
		WHERE user_id = $1
	`

	var settings model.PrivacySettings
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&settings.UserID, &settings.AnonymizeName, &settings.ShowOnLeaderboard,
		&settings.ShowOnJackpotList, &settings.OptOutOfShowcase)

	if err != nil {
		// Return defaults if not found
		return &model.PrivacySettings{
			UserID:            userID,
			AnonymizeName:     true,
			ShowOnLeaderboard: true,
			ShowOnJackpotList: true,
			OptOutOfShowcase:  false,
		}, nil
	}
	return &settings, nil
}

// UpdatePrivacySettings updates player's privacy settings
func (r *WinnersRepository) UpdatePrivacySettings(ctx context.Context, userID string, req model.UpdatePrivacyRequest) error {
	query := `
		INSERT INTO winner_privacy_settings (user_id, anonymize_name, show_on_leaderboard, 
		                                    show_on_jackpot_list, opt_out_of_showcase)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id) DO UPDATE SET
			anonymize_name = $2,
			show_on_leaderboard = $3,
			show_on_jackpot_list = $4,
			opt_out_of_showcase = $5
	`
	_, err := r.db.Exec(ctx, query, userID, req.AnonymizeName, req.ShowOnLeaderboard,
		req.ShowOnJackpotList, req.OptOutOfShowcase)

	return err
}

// CacheWinners caches winners in Redis
func (r *WinnersRepository) CacheWinners(ctx context.Context, key string, winners []model.Winner, ttl time.Duration) error {
	data, err := json.Marshal(winners)
	if err != nil {
		return err
	}
	return r.redis.Set(ctx, key, data, ttl).Err()
}
