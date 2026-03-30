package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/game_engine/leaderboard-service/internal/config"
	"github.com/game_engine/leaderboard-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type LeaderboardRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewLeaderboardRepository(db *pgxpool.Pool, redis *redis.Client) *LeaderboardRepository {
	return &LeaderboardRepository{db: db, redis: redis}
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

// GetLeaderboard retrieves leaderboard entries from Redis cache or DB
func (r *LeaderboardRepository) GetLeaderboard(ctx context.Context, leaderboardType model.LeaderboardType, gameType string, limit int) ([]model.LeaderboardEntry, error) {
	key := fmt.Sprintf("leaderboard:%s:%s", leaderboardType, gameType)

	data, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		var entries []model.LeaderboardEntry
		if json.Unmarshal([]byte(data), &entries) == nil {
			return entries, nil
		}
	}

	return r.getLeaderboardFromDB(ctx, leaderboardType, gameType, limit)
}

func (r *LeaderboardRepository) getLeaderboardFromDB(ctx context.Context, leaderboardType model.LeaderboardType, gameType string, limit int) ([]model.LeaderboardEntry, error) {
	period := getPeriodFromType(leaderboardType)

	query := `
		SELECT rank, user_id, username, score, wins, win_amount, game_type, updated_at
		FROM v_leaderboard_entries
		WHERE period = $1 AND ($2 = '' OR game_type = $2)
		ORDER BY rank LIMIT $3
	`

	rows, err := r.db.Query(ctx, query, period, gameType, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []model.LeaderboardEntry
	for rows.Next() {
		var e model.LeaderboardEntry
		if err := rows.Scan(&e.Rank, &e.UserID, &e.Username, &e.Score, &e.Wins, &e.WinAmount, &e.GameType, &e.UpdatedAt); err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

// GetPlayerRank retrieves a player's rank
func (r *LeaderboardRepository) GetPlayerRank(ctx context.Context, userID string, leaderboardType model.LeaderboardType, gameType string) (*model.PlayerRankResponse, error) {
	period := getPeriodFromType(leaderboardType)

	query := `
		SELECT rank, user_id, username, score
		FROM v_leaderboard_entries
		WHERE user_id = $1 AND period = $2 AND ($3 = '' OR game_type = $3)
	`

	var rank model.PlayerRankResponse
	err := r.db.QueryRow(ctx, query, userID, period, gameType).Scan(&rank.Rank, &rank.UserID, &rank.Username, &rank.Score)
	if err != nil {
		return nil, fmt.Errorf("player not found on leaderboard")
	}

	rank.LeaderboardType = leaderboardType
	rank.Period = period
	return &rank, nil
}

// UpdatePlayerScore updates a player's score
func (r *LeaderboardRepository) UpdatePlayerScore(ctx context.Context, req model.UpdateScoreRequest) error {
	period := getCurrentPeriod(model.LeaderboardType(req.GameType))

	query := `
		INSERT INTO player_scores (user_id, username, score, wins, win_amount, game_type, period, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		ON CONFLICT (user_id, game_type, period) 
		DO UPDATE SET 
			score = player_scores.score + $3,
			wins = player_scores.wins + $4,
			win_amount = player_scores.win_amount + $5,
			updated_at = NOW()
	`
	_, err := r.db.Exec(ctx, query, req.UserID, req.Username, req.Score, req.IsWin, req.WinAmount, req.GameType, period)
	return err
}

// CacheLeaderboard caches leaderboard in Redis
func (r *LeaderboardRepository) CacheLeaderboard(ctx context.Context, leaderboardType model.LeaderboardType, gameType string, entries []model.LeaderboardEntry) error {
	key := fmt.Sprintf("leaderboard:%s:%s", leaderboardType, gameType)
	data, err := json.Marshal(entries)
	if err != nil {
		return err
	}
	return r.redis.Set(ctx, key, data, 5*time.Minute).Err()
}

// GetPeriodLeaderboard gets leaderboard for a specific period
func (r *LeaderboardRepository) GetPeriodLeaderboard(ctx context.Context, period string, gameType string, limit int) ([]model.LeaderboardEntry, error) {
	query := `
		SELECT rank, user_id, username, score, wins, win_amount, game_type, updated_at
		FROM v_leaderboard_entries
		WHERE period = $1 AND ($2 = '' OR game_type = $2)
		ORDER BY rank LIMIT $3
	`

	rows, err := r.db.Query(ctx, query, period, gameType, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []model.LeaderboardEntry
	for rows.Next() {
		var e model.LeaderboardEntry
		if err := rows.Scan(&e.Rank, &e.UserID, &e.Username, &e.Score, &e.Wins, &e.WinAmount, &e.GameType, &e.UpdatedAt); err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

// GetPrizeConfigs retrieves prize configurations from database
func (r *LeaderboardRepository) GetPrizeConfigs(ctx context.Context, leaderboardType string, tournamentID string) ([]config.PrizeConfig, error) {
	query := `
		SELECT id, tournament_id, from_rank, to_rank, prize_type, value, currency, is_percentage
		FROM tournament_prize_configs
		WHERE leaderboard_type = $1 
		  AND ($2 = '' OR tournament_id = $2)
		  AND active = true
		ORDER BY from_rank
	`

	rows, err := r.db.Query(ctx, query, leaderboardType, tournamentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prizes []config.PrizeConfig
	for rows.Next() {
		var p config.PrizeConfig
		if err := rows.Scan(&p.ID, &p.TournamentID, &p.FromRank, &p.ToRank, &p.PrizeType, &p.Value, &p.Currency, &p.IsPercentage); err != nil {
			return nil, err
		}
		prizes = append(prizes, p)
	}
	return prizes, nil
}

// RecordPrizeDistribution records a prize distribution
func (r *LeaderboardRepository) RecordPrizeDistribution(ctx context.Context, req model.PrizeDistribution) error {
	query := `
		INSERT INTO leaderboard_prize_distributions 
		(leaderboard_type, game_type, period, user_id, rank, prize_type, value, currency, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.Exec(ctx, query,
		req.LeaderboardType, req.GameType, req.Period,
		req.UserID, req.Rank, req.PrizeType, req.Value, req.Currency, "distributed")
	return err
}

// GetPrizeDistributionHistory retrieves prize distribution history for a user
func (r *LeaderboardRepository) GetPrizeDistributionHistory(ctx context.Context, userID string, limit int) ([]model.PrizeDistribution, error) {
	query := `
		SELECT leaderboard_type, game_type, period, user_id, rank, prize_type, value, currency, distributed_at
		FROM leaderboard_prize_distributions
		WHERE user_id = $1
		ORDER BY distributed_at DESC
		LIMIT $2
	`

	rows, err := r.db.Query(ctx, query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var distributions []model.PrizeDistribution
	for rows.Next() {
		var d model.PrizeDistribution
		if err := rows.Scan(&d.LeaderboardType, &d.GameType, &d.Period, &d.UserID,
			&d.Rank, &d.PrizeType, &d.Value, &d.Currency, &d.DistributedAt); err != nil {
			return nil, err
		}
		distributions = append(distributions, d)
	}
	return distributions, nil
}
