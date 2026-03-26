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

	// Try to get from Redis first
	data, err := r.redis.Get(ctx, key).Result()
	if err == nil {
		var entries []model.LeaderboardEntry
		if json.Unmarshal([]byte(data), &entries) == nil {
			return entries, nil
		}
	}

	// Fall back to database
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
	// Get current period
	period := getCurrentPeriod(model.LeaderboardType(req.GameType))

	// Use upsert to insert or update
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

// RedisSortedSetLeaderboard retrieves leaderboard using Redis sorted sets (high-performance)
func (r *LeaderboardRepository) RedisSortedSetLeaderboard(ctx context.Context, leaderboardType model.LeaderboardType, gameType string, limit int) ([]model.LeaderboardEntry, error) {
	key := getRedisKey(leaderboardType, gameType)

	// Get top players from Redis sorted set (highest score first)
	results, err := r.redis.ZRevRangeWithScores(ctx, key, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	var entries []model.LeaderboardEntry
	for i, z := range results {
		userID := z.Member.(string)
		entry := model.LeaderboardEntry{
			Rank:      i + 1,
			UserID:    userID,
			Score:     z.Score,
			UpdatedAt: time.Now(),
		}
		// Try to get username from Redis hash
		if username, err := r.redis.HGet(ctx, getUserHashKey(userID), "username").Result(); err == nil {
			entry.Username = username
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// RedisUpdateScore updates player score in Redis sorted set
func (r *LeaderboardRepository) RedisUpdateScore(ctx context.Context, userID string, leaderboardType model.LeaderboardType, gameType string, score float64) error {
	key := getRedisKey(leaderboardType, gameType)

	// Increment score in sorted set
	err := r.redis.ZIncrBy(ctx, key, score, userID).Err()
	if err != nil {
		return err
	}

	// Set expiry for periodic leaderboards (daily/weekly/monthly)
	if leaderboardType != model.LeaderboardTypeAllTime {
		expiry := getRedisExpiry(leaderboardType)
		r.redis.Expire(ctx, key, expiry)
	}

	return nil
}

// RedisGetPlayerRank gets player's rank from Redis sorted set
func (r *LeaderboardRepository) RedisGetPlayerRank(ctx context.Context, userID string, leaderboardType model.LeaderboardType, gameType string) (*model.PlayerRankResponse, error) {
	key := getRedisKey(leaderboardType, gameType)

	// Get player's rank (0-indexed, so add 1)
	rank, err := r.redis.ZRevRank(ctx, key, userID).Result()
	if err != nil {
		return nil, fmt.Errorf("player not found on leaderboard")
	}

	// Get player's score
	score, err := r.redis.ZScore(ctx, key, userID).Result()
	if err != nil {
		return nil, err
	}

	return &model.PlayerRankResponse{
		UserID:          userID,
		Rank:            int(rank) + 1,
		Score:           score,
		LeaderboardType: leaderboardType,
		Period:          getPeriodFromType(leaderboardType),
	}, nil
}

// RedisGetTopPlayersAroundPlayer gets top players around a specific player
func (r *LeaderboardRepository) RedisGetTopPlayersAroundPlayer(ctx context.Context, userID string, leaderboardType model.LeaderboardType, gameType string, count int) ([]model.LeaderboardEntry, error) {
	key := getRedisKey(leaderboardType, gameType)

	// Get player's rank
	rank, err := r.redis.ZRevRank(ctx, key, userID).Result()
	if err != nil {
		return nil, err
	}

	// Calculate range around player
	start := rank - int64(count/2)
	if start < 0 {
		start = 0
	}
	end := start + int64(count) - 1

	// Get players around
	results, err := r.redis.ZRevRangeWithScores(ctx, key, start, end).Result()
	if err != nil {
		return nil, err
	}

	var entries []model.LeaderboardEntry
	for i, z := range results {
		entry := model.LeaderboardEntry{
			Rank:      int(start) + i + 1,
			UserID:    z.Member.(string),
			Score:     z.Score,
			UpdatedAt: time.Now(),
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// RedisSyncFromDB syncs leaderboard from PostgreSQL to Redis
func (r *LeaderboardRepository) RedisSyncFromDB(ctx context.Context, leaderboardType model.LeaderboardType, gameType string, limit int) error {
	entries, err := r.getLeaderboardFromDB(ctx, leaderboardType, gameType, limit)
	if err != nil {
		return err
	}

	key := getRedisKey(leaderboardType, gameType)

	// Clear existing data
	r.redis.Del(ctx, key)

	// Add all entries to Redis
	members := make([]redis.Z, len(entries))
	for i, e := range entries {
		members[i] = redis.Z{Score: e.Score, Member: e.UserID}
	}

	if len(members) > 0 {
		if err := r.redis.ZAdd(ctx, key, members...).Err(); err != nil {
			return err
		}
	}

	// Store usernames in hash
	for _, e := range entries {
		r.redis.HSet(ctx, getUserHashKey(e.UserID), "username", e.Username)
		r.redis.HSet(ctx, getUserHashKey(e.UserID), "wins", fmt.Sprintf("%d", e.Wins))
		r.redis.HSet(ctx, getUserHashKey(e.UserID), "win_amount", fmt.Sprintf("%.2f", e.WinAmount))
	}

	// Set expiry for periodic leaderboards
	if leaderboardType != model.LeaderboardTypeAllTime {
		expiry := getRedisExpiry(leaderboardType)
		r.redis.Expire(ctx, key, expiry)
	}

	return nil
}

// RedisResetLeaderboard resets (clears) a Redis leaderboard
func (r *LeaderboardRepository) RedisResetLeaderboard(ctx context.Context, leaderboardType model.LeaderboardType, gameType string) error {
	key := getRedisKey(leaderboardType, gameType)
	return r.redis.Del(ctx, key).Err()
}

func getRedisKey(leaderboardType model.LeaderboardType, gameType string) string {
	period := getPeriodFromType(leaderboardType)
	if gameType == "" || gameType == "all" {
		return fmt.Sprintf("leaderboard:%s:all", leaderboardType)
	}
	return fmt.Sprintf("leaderboard:%s:%s:%s", leaderboardType, gameType, period)
}

func getUserHashKey(userID string) string {
	return fmt.Sprintf("leaderboard:user:%s", userID)
}

func getRedisExpiry(leaderboardType model.LeaderboardType) time.Duration {
	switch leaderboardType {
	case model.LeaderboardTypeDaily:
		return 25 * time.Hour // Slightly more than 24h to allow for clock skew
	case model.LeaderboardTypeWeekly:
		return 8 * 24 * time.Hour // 8 days
	case model.LeaderboardTypeMonthly:
		return 32 * 24 * time.Hour // 32 days
	default:
		return 0 // No expiry for all-time
	}
}

func getPeriodFromType(leaderboardType model.LeaderboardType) string {
	now := time.Now()
	switch leaderboardType {
	case model.LeaderboardTypeDaily:
		return now.Format("2006-01-02")
	case model.LeaderboardTypeWeekly:
		year, week := now.ISOWeek()
		return fmt.Sprintf("%d-W%02d", year, week)
	case model.LeaderboardTypeMonthly:
		return now.Format("2006-01")
	case model.LeaderboardTypeAllTime:
		return "alltime"
	default:
		return now.Format("2006-01-02")
	}
}

func getCurrentPeriod(leaderboardType model.LeaderboardType) string {
	return getPeriodFromType(leaderboardType)
}
