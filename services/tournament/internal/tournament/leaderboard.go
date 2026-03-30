package tournament

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// LeaderboardEntry represents a player's position on the leaderboard
type LeaderboardEntry struct {
	UserID       string    `json:"user_id"`
	Username     string    `json:"username"`
	Score        int       `json:"score"`
	Rank         int       `json:"rank"`
	Knockouts    int       `json:"knockouts"`
	Eliminated   bool      `json:"eliminated"`
	EliminatedAt time.Time `json:"eliminated_at,omitempty"`
}

// Leaderboard manages real-time tournament leaderboards
type Leaderboard struct {
	redisClient *redis.Client
}

// NewLeaderboard creates a new leaderboard manager
func NewLeaderboard(redisClient *redis.Client) *Leaderboard {
	return &Leaderboard{
		redisClient: redisClient,
	}
}

// UpdateRegistration adds a participant to the leaderboard
func (l *Leaderboard) UpdateRegistration(ctx context.Context, tournamentID string, participant Participant) error {
	key := fmt.Sprintf("leaderboard:%s", tournamentID)

	entry := LeaderboardEntry{
		UserID:     participant.UserID,
		Username:   participant.Username,
		Score:      participant.Score,
		Rank:       0,
		Knockouts:  0,
		Eliminated: false,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	return l.redisClient.ZAdd(ctx, key, redis.Z{
		Score:  0,
		Member: data,
	}).Err()
}

// UpdateScore updates a player's score
func (l *Leaderboard) UpdateScore(ctx context.Context, tournamentID string, userID string, score int) error {
	key := fmt.Sprintf("leaderboard:%s", tournamentID)

	entries, err := l.redisClient.ZRangeWithScores(ctx, key, 0, -1).Result()
	if err != nil {
		return err
	}

	for _, z := range entries {
		var entry LeaderboardEntry
		if err := json.Unmarshal([]byte(z.Member.(string)), &entry); err != nil {
			continue
		}

		if entry.UserID == userID {
			entry.Score = score
			data, _ := json.Marshal(entry)

			l.redisClient.ZRem(ctx, key, z.Member)
			l.redisClient.ZAdd(ctx, key, redis.Z{
				Score:  float64(score),
				Member: data,
			})
			break
		}
	}

	return nil
}

// UpdateElimination marks a player as eliminated
func (l *Leaderboard) UpdateElimination(ctx context.Context, tournamentID string, userID string, rank int) error {
	key := fmt.Sprintf("leaderboard:%s", tournamentID)

	entries, err := l.redisClient.ZRangeWithScores(ctx, key, 0, -1).Result()
	if err != nil {
		return err
	}

	for _, z := range entries {
		var entry LeaderboardEntry
		if err := json.Unmarshal([]byte(z.Member.(string)), &entry); err != nil {
			continue
		}

		if entry.UserID == userID {
			entry.Eliminated = true
			entry.EliminatedAt = time.Now()
			entry.Rank = rank
			data, _ := json.Marshal(entry)

			l.redisClient.ZRem(ctx, key, z.Member)
			l.redisClient.ZAdd(ctx, key, redis.Z{
				Score:  z.Score,
				Member: data,
			})
			break
		}
	}

	return nil
}

// UpdateFinalResult updates final tournament results
func (l *Leaderboard) UpdateFinalResult(ctx context.Context, tournamentID string, result Result) error {
	resultKey := fmt.Sprintf("results:%s", tournamentID)

	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return l.redisClient.SAdd(ctx, resultKey, data).Err()
}

// GetLeaderboard retrieves the current leaderboard
func (l *Leaderboard) GetLeaderboard(ctx context.Context, tournamentID string, limit int) ([]LeaderboardEntry, error) {
	key := fmt.Sprintf("leaderboard:%s", tournamentID)

	if limit <= 0 {
		limit = 100
	}

	entries, err := l.redisClient.ZRevRangeWithScores(ctx, key, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}

	result := make([]LeaderboardEntry, 0, len(entries))
	rank := 1

	for _, z := range entries {
		var entry LeaderboardEntry
		if err := json.Unmarshal([]byte(z.Member.(string)), &entry); err != nil {
			continue
		}
		entry.Rank = rank
		result = append(result, entry)
		rank++
	}

	return result, nil
}

// GetPlayerRank gets a specific player's rank
func (l *Leaderboard) GetPlayerRank(ctx context.Context, tournamentID string, userID string) (*LeaderboardEntry, error) {
	key := fmt.Sprintf("leaderboard:%s", tournamentID)

	entries, err := l.redisClient.ZRevRangeWithScores(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	rank := 1
	for _, z := range entries {
		var entry LeaderboardEntry
		if err := json.Unmarshal([]byte(z.Member.(string)), &entry); err != nil {
			continue
		}

		if entry.UserID == userID {
			entry.Rank = rank
			return &entry, nil
		}
		rank++
	}

	return nil, fmt.Errorf("player not found")
}

// GetTopPlayers returns the top N players
func (l *Leaderboard) GetTopPlayers(ctx context.Context, tournamentID string, n int) ([]LeaderboardEntry, error) {
	key := fmt.Sprintf("leaderboard:%s", tournamentID)

	if n <= 0 {
		n = 10
	}

	entries, err := l.redisClient.ZRevRangeWithScores(ctx, key, 0, int64(n-1)).Result()
	if err != nil {
		return nil, err
	}

	result := make([]LeaderboardEntry, 0, len(entries))
	rank := 1

	for _, z := range entries {
		var entry LeaderboardEntry
		if err := json.Unmarshal([]byte(z.Member.(string)), &entry); err != nil {
			continue
		}
		entry.Rank = rank
		result = append(result, entry)
		rank++
	}

	return result, nil
}

// RemoveParticipant removes a participant from the leaderboard
func (l *Leaderboard) RemoveParticipant(ctx context.Context, tournamentID string, userID string) error {
	key := fmt.Sprintf("leaderboard:%s", tournamentID)

	entries, err := l.redisClient.ZRangeWithScores(ctx, key, 0, -1).Result()
	if err != nil {
		return err
	}

	for _, z := range entries {
		var entry LeaderboardEntry
		if err := json.Unmarshal([]byte(z.Member.(string)), &entry); err != nil {
			continue
		}

		if entry.UserID == userID {
			return l.redisClient.ZRem(ctx, key, z.Member).Err()
		}
	}

	return nil
}
