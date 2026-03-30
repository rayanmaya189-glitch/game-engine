package moderation

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
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
	config      *FilterConfig
	redisClient *redis.Client
}

// NewModerator creates a new moderator
func NewModerator(config *FilterConfig, redisClient *redis.Client) *Moderator {
	return &Moderator{
		config:      config,
		redisClient: redisClient,
	}
}

// CheckRateLimit checks if user has exceeded rate limit
func (m *Moderator) CheckRateLimit(ctx context.Context, userID string) bool {
	key := fmt.Sprintf("ratelimit:%s", userID)

	count, err := m.redisClient.Get(ctx, key).Int()
	if err == redis.Nil {
		return true
	}
	if err != nil {
		return true // Fail open
	}

	return count < m.config.Moderation.AutoMuteThreshold
}

// RecordMessage records a message for rate limiting
func (m *Moderator) RecordMessage(ctx context.Context, userID string) error {
	key := fmt.Sprintf("ratelimit:%s", userID)

	pipe := m.redisClient.Pipeline()
	pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, time.Minute)
	_, err := pipe.Exec(ctx)

	return err
}

// MuteUser mutes a user
func (m *Moderator) MuteUser(ctx context.Context, userID string, durationMinutes int) error {
	if durationMinutes == 0 {
		durationMinutes = m.config.Moderation.MuteDurationMinutes
	}

	key := fmt.Sprintf("mute:%s", userID)
	return m.redisClient.Set(ctx, key, "1", time.Duration(durationMinutes)*time.Minute).Err()
}

// UnmuteUser unmutes a user
func (m *Moderator) UnmuteUser(ctx context.Context, userID string) error {
	key := fmt.Sprintf("mute:%s", userID)
	return m.redisClient.Del(ctx, key).Err()
}

// IsMuted checks if user is muted
func (m *Moderator) IsMuted(ctx context.Context, userID string) bool {
	key := fmt.Sprintf("mute:%s", userID)
	_, err := m.redisClient.Get(ctx, key).Result()
	return err == nil
}

// BanUser bans a user
func (m *Moderator) BanUser(ctx context.Context, userID string, durationHours int) error {
	if durationHours == 0 {
		durationHours = m.config.Moderation.BanDurationHours
	}

	key := fmt.Sprintf("ban:%s", userID)
	return m.redisClient.Set(ctx, key, "1", time.Duration(durationHours)*time.Hour).Err()
}

// UnbanUser unbans a user
func (m *Moderator) UnbanUser(ctx context.Context, userID string) error {
	key := fmt.Sprintf("ban:%s", userID)
	return m.redisClient.Del(ctx, key).Err()
}

// IsBanned checks if user is banned
func (m *Moderator) IsBanned(ctx context.Context, userID string) bool {
	key := fmt.Sprintf("ban:%s", userID)
	_, err := m.redisClient.Get(ctx, key).Result()
	return err == nil
}

// GetModerationConfig returns the moderation config
func (m *Moderator) GetModerationConfig() *ModerationConfig {
	return &m.config.Moderation
}
