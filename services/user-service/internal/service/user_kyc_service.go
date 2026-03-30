package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/game_engine/user-service/internal/model"
	"github.com/google/uuid"
)

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
