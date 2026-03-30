package tournament

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/redis/go-redis/v9"
)

// GetGlobalLeaderboard returns the global leaderboard across all tournaments
func (l *Leaderboard) GetGlobalLeaderboard(ctx context.Context, limit int) ([]LeaderboardEntry, error) {
	key := "global:leaderboard"

	if limit <= 0 {
		limit = 100
	}

	entries, err := l.redisClient.ZRevRangeWithScores(ctx, key, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	result := make([]LeaderboardEntry, 0, len(entries))

	for _, z := range entries {
		var entry LeaderboardEntry
		if err := json.Unmarshal([]byte(z.Member.(string)), &entry); err != nil {
			continue
		}
		result = append(result, entry)
	}

	return result, nil
}

// UpdateGlobalScore updates a player's global score
func (l *Leaderboard) UpdateGlobalScore(ctx context.Context, userID string, username string, scoreDelta int) error {
	key := "global:leaderboard"

	currentScore, err := l.redisClient.ZScore(ctx, key, userID).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	newScore := currentScore + float64(scoreDelta)

	entry := LeaderboardEntry{
		UserID:   userID,
		Username: username,
		Score:    int(newScore),
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	return l.redisClient.ZAdd(ctx, key, redis.Z{
		Score:  newScore,
		Member: data,
	}).Err()
}

// ClearLeaderboard clears a tournament leaderboard
func (l *Leaderboard) ClearLeaderboard(ctx context.Context, tournamentID string) error {
	key := fmt.Sprintf("leaderboard:%s", tournamentID)
	return l.redisClient.Del(ctx, key).Err()
}

// GetLeaderboardWithSort allows custom sorting of leaderboard entries
func (l *Leaderboard) GetLeaderboardWithSort(ctx context.Context, tournamentID string, sortBy string, limit int) ([]LeaderboardEntry, error) {
	entries, err := l.GetLeaderboard(ctx, tournamentID, limit)
	if err != nil {
		return nil, err
	}

	switch sortBy {
	case "score":
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Score > entries[j].Score
		})
	case "knockouts":
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Knockouts > entries[j].Knockouts
		})
	case "rank":
		// Already sorted by rank
	}

	for i := range entries {
		entries[i].Rank = i + 1
	}

	return entries, nil
}
