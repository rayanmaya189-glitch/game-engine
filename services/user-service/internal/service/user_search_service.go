package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/game_engine/user-service/internal/model"
)

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
