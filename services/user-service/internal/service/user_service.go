package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gameengine/user-service/internal/config"
	"github.com/gameengine/user-service/internal/model"
	"github.com/gameengine/user-service/internal/repository"
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

// GetKYCStatus retrieves KYC status for a user
func (s *UserService) GetKYCStatus(ctx context.Context, userID string) (*model.KYCStatus, error) {
	// Try cache first
	cacheKey := fmt.Sprintf("kyc_status:%s", userID)
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var status model.KYCStatus
		if json.Unmarshal([]byte(cached), &status) == nil {
			return &status, nil
		}
	}

	// Fetch from database
	status, err := s.repo.GetKYCStatus(ctx, userID)
	if err != nil {
		return nil, err
	}

	if status == nil {
		// Return default KYC status
		return &model.KYCStatus{
			UserID: userID,
			Status: "VERIFICATION_STATUS_UNSPECIFIED",
			Level:  "KYC_LEVEL_NONE",
		}, nil
	}

	// Cache the result
	data, _ := json.Marshal(status)
	s.redis.Set(ctx, cacheKey, data, time.Duration(s.cfg.Cache.KYCStatusTTL)*time.Second)

	return status, nil
}

// SubmitKYC submits KYC documents
func (s *UserService) SubmitKYC(ctx context.Context, userID, docType, docNumber, docData string) error {
	// Check if user has a profile
	profile, err := s.repo.GetProfileByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if profile == nil {
		return errors.New("profile not found")
	}

	// Get current KYC status
	currentStatus, _ := s.repo.GetKYCStatus(ctx, userID)

	// Determine next KYC level
	newLevel := "KYC_LEVEL_BASIC"
	if currentStatus != nil {
		switch currentStatus.Level {
		case "KYC_LEVEL_BASIC":
			newLevel = "KYC_LEVEL_INTERMEDIATE"
		case "KYC_LEVEL_INTERMEDIATE":
			newLevel = "KYC_LEVEL_FULL"
		}
	}

	// Create KYC document
	doc := &model.KYCDocument{
		ID:             uuid.New().String(),
		UserID:         userID,
		DocumentType:   docType,
		DocumentNumber: docNumber,
		DocumentData:   docData,
		Status:         "VERIFICATION_STATUS_PENDING",
	}

	err = s.repo.CreateKYCDocument(ctx, doc)
	if err != nil {
		return err
	}

	// Create or update KYC status
	kycStatus := &model.KYCStatus{
		ID:     uuid.New().String(),
		UserID: userID,
		Status: "VERIFICATION_STATUS_PENDING",
		Level:  newLevel,
	}

	if currentStatus == nil {
		err = s.repo.CreateKYCStatus(ctx, kycStatus)
	} else {
		kycStatus.ID = currentStatus.ID
		err = s.repo.UpdateKYCStatus(ctx, kycStatus)
	}

	if err != nil {
		return err
	}

	// Invalidate KYC cache
	s.redis.Del(ctx, fmt.Sprintf("kyc_status:%s", userID))

	// Publish event
	s.publishEvent(ctx, EventKYCSubmitted, map[string]string{
		"user_id": userID,
		"level":   newLevel,
	})

	// KYC provider integration would go here
	// For now, we'll auto-approve basic level if configured
	if s.cfg.KYC.AutoApproveBasic && newLevel == "KYC_LEVEL_BASIC" {
		go s.processKYCApproval(userID, newLevel)
	}

	return nil
}

// processKYCApproval simulates KYC approval (in production, this would call the KYC provider)
func (s *UserService) processKYCApproval(userID, level string) {
	time.Sleep(100 * time.Millisecond)

	ctx := context.Background()

	// Update KYC status
	kycStatus := &model.KYCStatus{
		UserID: userID,
		Status: "VERIFICATION_STATUS_VERIFIED",
		Level:  level,
	}
	s.repo.UpdateKYCStatus(ctx, kycStatus)

	// Update profile KYC level
	s.repo.UpdatePlayerKYCLevel(ctx, userID, level)

	// Invalidate cache
	s.redis.Del(ctx, fmt.Sprintf("kyc_status:%s", userID))
	s.redis.Del(ctx, fmt.Sprintf("profile:%s", userID))

	// Publish event
	s.publishEvent(ctx, EventKYCApproved, map[string]string{
		"user_id": userID,
		"level":   level,
	})
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

// GetPlayerByAdmin retrieves a player by identifier (admin function)
func (s *UserService) GetPlayerByAdmin(ctx context.Context, identifier string) (*model.Profile, *model.KYCStatus, error) {
	profile, err := s.repo.GetProfileByIdentifier(ctx, identifier)
	if err != nil {
		return nil, nil, err
	}

	if profile == nil {
		return nil, nil, errors.New("player not found")
	}

	kycStatus, _ := s.repo.GetKYCStatus(ctx, profile.UserID)
	if kycStatus == nil {
		kycStatus = &model.KYCStatus{
			UserID: profile.UserID,
			Status: "VERIFICATION_STATUS_UNSPECIFIED",
			Level:  "KYC_LEVEL_NONE",
		}
	}

	return profile, kycStatus, nil
}

// ListPlayers lists players with filters and pagination (admin function)
func (s *UserService) ListPlayers(ctx context.Context, status, kycLevel, country, search string, page, pageSize int) ([]*model.Profile, int, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListPlayers(ctx, status, kycLevel, country, search, pageSize, offset)
}

// UpdatePlayerStatus updates player status (admin function)
func (s *UserService) UpdatePlayerStatus(ctx context.Context, userID, status, reason string) error {
	// Get existing status
	profile, err := s.repo.GetProfileByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if profile == nil {
		return errors.New("player not found")
	}

	oldStatus := profile.Status

	err = s.repo.UpdatePlayerStatus(ctx, userID, status)
	if err != nil {
		return err
	}

	// Invalidate cache
	s.redis.Del(ctx, fmt.Sprintf("profile:%s", userID))

	// Publish event
	s.publishEvent(ctx, EventStatusChanged, map[string]string{
		"user_id":    userID,
		"old_status": oldStatus,
		"new_status": status,
		"reason":     reason,
	})

	return nil
}

// UpdatePlayerKYCLevel updates player KYC level (admin function)
func (s *UserService) UpdatePlayerKYCLevel(ctx context.Context, userID, level string) error {
	err := s.repo.UpdatePlayerKYCLevel(ctx, userID, level)
	if err != nil {
		return err
	}

	// Update KYC status
	kycStatus := &model.KYCStatus{
		UserID: userID,
		Status: "VERIFICATION_STATUS_VERIFIED",
		Level:  level,
	}
	s.repo.UpdateKYCStatus(ctx, kycStatus)

	// Invalidate caches
	s.redis.Del(ctx, fmt.Sprintf("profile:%s", userID))
	s.redis.Del(ctx, fmt.Sprintf("kyc_status:%s", userID))

	// Publish event
	s.publishEvent(ctx, EventKYCApproved, map[string]string{
		"user_id": userID,
		"level":   level,
	})

	return nil
}

// GetPlayerStats retrieves player statistics (admin function)
func (s *UserService) GetPlayerStats(ctx context.Context, userID string) (*model.PlayerStats, error) {
	stats, err := s.repo.GetPlayerStats(ctx, userID)
	if err != nil {
		return nil, err
	}

	if stats == nil {
		// Return empty stats
		return &model.PlayerStats{
			UserID:           userID,
			TotalDeposits:    0,
			TotalWithdrawals: 0,
			TotalBets:        0,
			TotalWins:        0,
			TotalBonuses:     0,
		}, nil
	}

	return stats, nil
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
