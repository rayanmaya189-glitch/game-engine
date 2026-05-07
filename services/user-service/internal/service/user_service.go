package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/game_engine/user-service/internal/config"
	"github.com/game_engine/user-service/internal/model"
	"github.com/game_engine/user-service/internal/repository"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

// UserService handles user business logic
type UserService struct {
	repo  *repository.UserRepository
	redis *redis.Client
	nats  *nats.Conn
	cfg   *config.Config
}

// NewUserService creates a new user service
func NewUserService(repo *repository.UserRepository, redis *redis.Client, nats *nats.Conn, cfg *config.Config) (*UserService, error) {
	return &UserService{
		repo:  repo,
		redis: redis,
		nats:  nats,
		cfg:   cfg,
	}, nil
}

// NATS Event types
const (
	EventProfileUpdated = "player.events.profile_updated"
	EventKYCSubmitted   = "player.events.kyc_submitted"
	EventKYCApproved    = "player.events.kyc_approved"
	EventKYCRejected    = "player.events.kyc_rejected"
	EventStatusChanged  = "player.events.status_changed"
)

// GetProfile retrieves a player profile by user ID
func (s *UserService) GetProfile(ctx context.Context, userID string) (*model.Profile, error) {
	// Try cache first
	cacheKey := fmt.Sprintf("profile:%s", userID)
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var profile model.Profile
		if json.Unmarshal([]byte(cached), &profile) == nil {
			return &profile, nil
		}
	}

	// Fetch from database
	profile, err := s.repo.GetProfileByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if profile == nil {
		return nil, errors.New("profile not found")
	}

	// Cache the result
	data, _ := json.Marshal(profile)
	s.redis.Set(ctx, cacheKey, data, time.Duration(s.cfg.Cache.ProfileTTL)*time.Second)

	return profile, nil
}

// UpdateProfile updates a player's profile
func (s *UserService) UpdateProfile(ctx context.Context, userID string, updates *model.Profile) (*model.Profile, error) {
	// Rate limiting check
	rateKey := fmt.Sprintf("rate:profile_update:%s", userID)
	allowed, err := s.redis.Eval(ctx, `
		local count = redis.call('INCR', KEYS[1])
		if count == 1 then
			redis.call('EXPIRE', KEYS[1], 60)
		end
		return count
	`, []string{rateKey}).Int()

	if err == nil && allowed > s.cfg.RateLimiting.ProfileUpdateMax {
		return nil, errors.New("rate limit exceeded for profile updates")
	}

	// Get existing profile
	existing, err := s.repo.GetProfileByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, errors.New("profile not found")
	}

	// Apply updates
	if updates.DisplayName != "" {
		existing.DisplayName = updates.DisplayName
	}
	if updates.AvatarURL != "" {
		existing.AvatarURL = updates.AvatarURL
	}
	if updates.Language != "" {
		existing.Language = updates.Language
	}
	if updates.Timezone != "" {
		existing.Timezone = updates.Timezone
	}
	if updates.FirstName != "" {
		existing.FirstName = updates.FirstName
	}
	if updates.LastName != "" {
		existing.LastName = updates.LastName
	}
	if updates.DateOfBirth != nil {
		existing.DateOfBirth = updates.DateOfBirth
	}
	if updates.Gender != "" {
		existing.Gender = updates.Gender
	}

	// Update in database
	err = s.repo.UpdateProfile(ctx, existing)
	if err != nil {
		return nil, err
	}

	// Invalidate cache
	s.redis.Del(ctx, fmt.Sprintf("profile:%s", userID))

	// Publish event
	s.publishEvent(ctx, EventProfileUpdated, map[string]string{
		"user_id": userID,
	})

	return existing, nil
}

// GetPlayerSettings retrieves player settings
func (s *UserService) GetPlayerSettings(ctx context.Context, userID string) (*model.PlayerSettings, error) {
	settings, err := s.repo.GetPlayerSettings(ctx, userID)
	if err != nil {
		return nil, err
	}

	if settings == nil {
		// Return default settings
		return &model.PlayerSettings{
			UserID:             userID,
			EmailNotifications: true,
			SMSNotifications:   false,
			PushNotifications:  true,
			ProfilePublic:      false,
			ShowOnlineStatus:   true,
			AutoPlay:           false,
			SoundVolume:        50,
			Theme:              "default",
		}, nil
	}

	return settings, nil
}

// UpdatePlayerSettings updates player settings
func (s *UserService) UpdatePlayerSettings(ctx context.Context, userID string, updates *model.PlayerSettings) (*model.PlayerSettings, error) {
	// Get existing or create new
	settings, err := s.repo.GetPlayerSettings(ctx, userID)
	if err != nil {
		return nil, err
	}

	if settings == nil {
		settings = &model.PlayerSettings{
			ID:     uuid.New().String(),
			UserID: userID,
		}
	}

	// Apply updates
	if updates != nil {
		settings.EmailNotifications = updates.EmailNotifications
		settings.SMSNotifications = updates.SMSNotifications
		settings.PushNotifications = updates.PushNotifications
		settings.ProfilePublic = updates.ProfilePublic
		settings.ShowOnlineStatus = updates.ShowOnlineStatus
		settings.AutoPlay = updates.AutoPlay
		settings.SoundVolume = updates.SoundVolume
		settings.Theme = updates.Theme
	}

	// Validate settings
	if settings.SoundVolume < 0 || settings.SoundVolume > 100 {
		return nil, errors.New("sound volume must be between 0 and 100")
	}

	err = s.repo.UpdatePlayerSettings(ctx, settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

// SetPlayerLimits sets player limits (deposit, bet, loss)
func (s *UserService) SetPlayerLimits(ctx context.Context, userID string, limitType string, amount float64, period string) (*model.PlayerLimits, error) {
	// Get existing limits or create new
	limits, err := s.repo.GetPlayerLimits(ctx, userID)
	if err != nil {
		return nil, err
	}

	if limits == nil {
		limits = &model.PlayerLimits{
			ID:     uuid.New().String(),
			UserID: userID,
		}
	}

	// Apply the specific limit based on type
	switch limitType {
	case "DEPOSIT":
		limits.DailyLimit = int(amount)
	case "BET":
		// Bet limits stored separately
		limits.DailyLimit = int(amount) // Store as daily bet limit
	case "LOSS":
		limits.DailyLossLimit = int(amount)
	}

	err = s.repo.UpdatePlayerLimits(ctx, limits)
	if err != nil {
		return nil, err
	}

	// Invalidate cache
	s.redis.Del(ctx, fmt.Sprintf("limits:%s", userID))

	// Publish event
	s.publishEvent(ctx, "player.events.limits_updated", map[string]string{
		"user_id": userID,
		"type":    limitType,
	})

	return limits, nil
}

// GetPlayerLimits retrieves player limits
func (s *UserService) GetPlayerLimits(ctx context.Context, userID string) (*model.PlayerLimits, error) {
	// Try cache first
	cacheKey := fmt.Sprintf("limits:%s", userID)
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var limits model.PlayerLimits
		if json.Unmarshal([]byte(cached), &limits) == nil {
			return &limits, nil
		}
	}

	limits, err := s.repo.GetPlayerLimits(ctx, userID)
	if err != nil {
		return nil, err
	}

	if limits == nil {
		// Return default limits
		return &model.PlayerLimits{
			UserID:         userID,
			DailyLimit:     10000, // Default daily deposit limit
			WeeklyLimit:    50000,
			MonthlyLimit:   200000,
			DailyLossLimit: 5000,
		}, nil
	}

	// Cache the result
	data, _ := json.Marshal(limits)
	s.redis.Set(ctx, cacheKey, data, time.Duration(s.cfg.Cache.ProfileTTL)*time.Second)

	return limits, nil
}

// SelfExclude excludes a player from gambling
func (s *UserService) SelfExclude(ctx context.Context, userID string, duration string, reason string) error {
	limits, err := s.repo.GetPlayerLimits(ctx, userID)
	if err != nil {
		return err
	}

	if limits == nil {
		limits = &model.PlayerLimits{
			ID:     uuid.New().String(),
			UserID: userID,
		}
	}

	limits.SelfExclusion = true

	// Calculate exclusion end date based on duration
	now := time.Now()
	switch duration {
	case "EXCLUSION_24H":
		endDate := now.Add(24 * time.Hour)
		limits.ExclusionEndDate = &endDate
	case "EXCLUSION_7D":
		endDate := now.Add(7 * 24 * time.Hour)
		limits.ExclusionEndDate = &endDate
	case "EXCLUSION_30D":
		endDate := now.Add(30 * 24 * time.Hour)
		limits.ExclusionEndDate = &endDate
	case "EXCLUSION_90D":
		endDate := now.Add(90 * 24 * time.Hour)
		limits.ExclusionEndDate = &endDate
	case "EXCLUSION_6M":
		endDate := now.Add(180 * 24 * time.Hour)
		limits.ExclusionEndDate = &endDate
	case "EXCLUSION_1Y":
		endDate := now.Add(365 * 24 * time.Hour)
		limits.ExclusionEndDate = &endDate
	case "EXCLUSION_PERMANENT":
		// Permanent exclusion - no end date
		limits.ExclusionEndDate = nil
	}

	err = s.repo.UpdatePlayerLimits(ctx, limits)
	if err != nil {
		return err
	}

	// Update player status to suspended
	err = s.repo.UpdatePlayerStatus(ctx, userID, "SUSPENDED")
	if err != nil {
		return err
	}

	// Invalidate caches
	s.redis.Del(ctx, fmt.Sprintf("limits:%s", userID))
	s.redis.Del(ctx, fmt.Sprintf("profile:%s", userID))

	// Publish event
	s.publishEvent(ctx, "player.events.self_excluded", map[string]string{
		"user_id": userID,
		"duration": duration,
		"reason":   reason,
	})

	return nil
}

// CheckPlayerExclusion checks if a player is excluded
func (s *UserService) CheckPlayerExclusion(ctx context.Context, userID string) (bool, *time.Time, error) {
	limits, err := s.GetPlayerLimits(ctx, userID)
	if err != nil {
		return false, nil, err
	}

	if !limits.SelfExclusion {
		return false, nil, nil
	}

	// Check if exclusion has ended
	if limits.ExclusionEndDate != nil && limits.ExclusionEndDate.Before(time.Now()) {
		// Exclusion period has ended, clear it
		limits.SelfExclusion = false
		limits.ExclusionEndDate = nil
		err = s.repo.UpdatePlayerLimits(ctx, limits)
		if err != nil {
			return false, nil, err
		}
		// Restore status
		s.repo.UpdatePlayerStatus(ctx, userID, "ACTIVE")
		return false, nil, nil
	}

	return true, limits.ExclusionEndDate, nil
}

// publishEvent publishes a NATS event
func (s *UserService) publishEvent(ctx context.Context, eventType string, data map[string]string) {
	if s.nats == nil {
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	s.nats.Publish(eventType, jsonData)
}
