package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/game_engine/leaderboard-service/internal/model"
	"github.com/redis/go-redis/v9"
)

// RedisSortedSetLeaderboard retrieves leaderboard using Redis sorted sets (high-performance)
func (r *LeaderboardRepository) RedisSortedSetLeaderboard(ctx context.Context, leaderboardType model.LeaderboardType, gameType string, limit int) ([]model.LeaderboardEntry, error) {
	key := getRedisKey(leaderboardType, gameType)

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

	err := r.redis.ZIncrBy(ctx, key, score, userID).Err()
	if err != nil {
		return err
	}

	if leaderboardType != model.LeaderboardTypeAllTime {
		expiry := getRedisExpiry(leaderboardType)
		r.redis.Expire(ctx, key, expiry)
	}

	return nil
}

// RedisGetPlayerRank gets player's rank from Redis sorted set
func (r *LeaderboardRepository) RedisGetPlayerRank(ctx context.Context, userID string, leaderboardType model.LeaderboardType, gameType string) (*model.PlayerRankResponse, error) {
	key := getRedisKey(leaderboardType, gameType)

	rank, err := r.redis.ZRevRank(ctx, key, userID).Result()
	if err != nil {
		return nil, fmt.Errorf("player not found on leaderboard")
	}

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

	rank, err := r.redis.ZRevRank(ctx, key, userID).Result()
	if err != nil {
		return nil, err
	}

	start := rank - int64(count/2)
	if start < 0 {
		start = 0
	}
	end := start + int64(count) - 1

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

	r.redis.Del(ctx, key)

	members := make([]redis.Z, len(entries))
	for i, e := range entries {
		members[i] = redis.Z{Score: e.Score, Member: e.UserID}
	}

	if len(members) > 0 {
		if err := r.redis.ZAdd(ctx, key, members...).Err(); err != nil {
			return err
		}
	}

	for _, e := range entries {
		r.redis.HSet(ctx, getUserHashKey(e.UserID), "username", e.Username)
		r.redis.HSet(ctx, getUserHashKey(e.UserID), "wins", fmt.Sprintf("%d", e.Wins))
		r.redis.HSet(ctx, getUserHashKey(e.UserID), "win_amount", fmt.Sprintf("%.2f", e.WinAmount))
	}

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
		return 25 * time.Hour
	case model.LeaderboardTypeWeekly:
		return 8 * 24 * time.Hour
	case model.LeaderboardTypeMonthly:
		return 32 * 24 * time.Hour
	default:
		return 0
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
