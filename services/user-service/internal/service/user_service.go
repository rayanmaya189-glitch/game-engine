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
