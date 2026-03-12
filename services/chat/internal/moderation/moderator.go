package moderation

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/game-engine/chat/internal/room"
)

// ModerationAction represents a moderation action
type ModerationAction string

const (
	ActionMute   ModerationAction = "mute"
	ActionBan    ModerationAction = "ban"
	ActionUnmute ModerationAction = "unmute"
	ActionUnban  ModerationAction = "unban"
)

// Moderator handles moderation actions
type Moderator struct {
	config      *room.Config
	redisClient *redis.Client
}

// NewModerator creates a new moderator
func NewModerator(config *room.Config, redisClient *redis.Client) *Moderator {
	return &Moderator{
		config:      config,
		redisClient: redisClient,
	}
}

// CheckRateLimit checks if user has exceeded rate limit
func (m *Moderator) CheckRateLimit(ctx context.Context, userID string) bool {
	key := fmt.Sprintf("ratelimit:%s", userID)

	count, err := m.redisClient.Incr(ctx, key).Result()
	if err != nil {
		return true // Allow on error
	}

	if count == 1 {
		// Set expiry for the rate limit window
		m.redisClient.Expire(ctx, key, time.Minute)
	}

	return count <= int64(m.config.Chat.RateLimitPerMinute)
}

// MuteUser mutes a user
func (m *Moderator) MuteUser(ctx context.Context, userID string, duration time.Duration) error {
	key := fmt.Sprintf("mute:%s", userID)

	err := m.redisClient.Set(ctx, key, "1", duration).Err()
	if err != nil {
		return err
	}

	return nil
}

// UnmuteUser unmutes a user
func (m *Moderator) UnmuteUser(ctx context.Context, userID string) error {
	key := fmt.Sprintf("mute:%s", userID)

	return m.redisClient.Del(ctx, key).Err()
}

// IsMuted checks if a user is muted
func (m *Moderator) IsMuted(ctx context.Context, userID string) (bool, error) {
	key := fmt.Sprintf("mute:%s", userID)

	exists, err := m.redisClient.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

// BanUser bans a user
func (m *Moderator) BanUser(ctx context.Context, userID string, duration time.Duration) error {
	key := fmt.Sprintf("ban:%s", userID)

	err := m.redisClient.Set(ctx, key, "1", duration).Err()
	if err != nil {
		return err
	}

	return nil
}

// UnbanUser unbans a user
func (m *Moderator) UnbanUser(ctx context.Context, userID string) error {
	key := fmt.Sprintf("ban:%s", userID)

	return m.redisClient.Del(ctx, key).Err()
}

// IsBanned checks if a user is banned
func (m *Moderator) IsBanned(ctx context.Context, userID string) (bool, error) {
	key := fmt.Sprintf("ban:%s", userID)

	exists, err := m.redisClient.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}

// RecordWarning records a warning for a user
func (m *Moderator) RecordWarning(ctx context.Context, userID string) (int, error) {
	key := fmt.Sprintf("warnings:%s", userID)

	count, err := m.redisClient.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	// Set expiry for warnings (reset daily)
	m.redisClient.Expire(ctx, key, 24*time.Hour)

	// Auto-mute if threshold reached
	if int(count) >= m.config.Moderation.AutoMuteThreshold {
		muteDuration := time.Duration(m.config.Moderation.MuteDurationMinutes) * time.Minute
		_ = m.MuteUser(ctx, userID, muteDuration)
	}

	return int(count), nil
}

// GetWarningCount gets the warning count for a user
func (m *Moderator) GetWarningCount(ctx context.Context, userID string) (int, error) {
	key := fmt.Sprintf("warnings:%s", userID)

	count, err := m.redisClient.Get(ctx, key).Int()
	if err == redis.Nil {
		return 0, nil
	}

	return count, err
}

// ClearWarnings clears warnings for a user
func (m *Moderator) ClearWarnings(ctx context.Context, userID string) error {
	key := fmt.Sprintf("warnings:%s", userID)

	return m.redisClient.Del(ctx, key).Err()
}

// PerformAction performs a moderation action
func (m *Moderator) PerformAction(ctx context.Context, userID string, action ModerationAction, duration time.Duration) error {
	switch action {
	case ActionMute:
		return m.MuteUser(ctx, userID, duration)
	case ActionUnmute:
		return m.UnmuteUser(ctx, userID)
	case ActionBan:
		return m.BanUser(ctx, userID, duration)
	case ActionUnban:
		return m.UnbanUser(ctx, userID)
	default:
		return fmt.Errorf("unknown action: %s", action)
	}
}
