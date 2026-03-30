package service

import (
	"context"
	"time"

	"github.com/game_engine/winners-showcase-service/internal/config"
	"github.com/game_engine/winners-showcase-service/internal/model"
	"github.com/game_engine/winners-showcase-service/internal/repository"
)

type WinnersService struct {
	repo   *repository.WinnersRepository
	config *config.WinnersConfig
}

func NewWinnersService(repo *repository.WinnersRepository, cfg *config.WinnersConfig) *WinnersService {
	return &WinnersService{repo: repo, config: cfg}
}

func (s *WinnersService) GetRecentWinners(ctx context.Context) (*model.RecentWinnersResponse, error) {
	winners, err := s.repo.GetRecentWinners(ctx, s.config.RecentWinnersLimit)
	if err != nil {
		return nil, err
	}

	// Apply privacy settings
	for i := range winners {
		if s.config.AnonymizeNames {
			winners[i].DisplayName = anonymizeName(winners[i].Username)
		}
	}

	return &model.RecentWinnersResponse{
		Winners:   winners,
		Total:     len(winners),
		UpdatedAt: time.Now(),
	}, nil
}

func (s *WinnersService) GetBigWins(ctx context.Context) (*model.BigWinsResponse, error) {
	winners, err := s.repo.GetBigWins(ctx, s.config.BigWinThreshold, 50)
	if err != nil {
		return nil, err
	}

	for i := range winners {
		if s.config.AnonymizeNames {
			winners[i].DisplayName = anonymizeName(winners[i].Username)
		}
	}

	return &model.BigWinsResponse{
		Wins:      winners,
		Threshold: s.config.BigWinThreshold,
		Total:     len(winners),
		UpdatedAt: time.Now(),
	}, nil
}

func (s *WinnersService) GetJackpotWinners(ctx context.Context) (*model.JackpotWinnersResponse, error) {
	winners, err := s.repo.GetJackpotWinners(ctx, s.config.JackpotThreshold, 50)
	if err != nil {
		return nil, err
	}

	for i := range winners {
		if s.config.AnonymizeNames {
			winners[i].DisplayName = anonymizeName(winners[i].Username)
		}
	}

	return &model.JackpotWinnersResponse{
		Winners:   winners,
		Threshold: s.config.JackpotThreshold,
		Total:     len(winners),
		UpdatedAt: time.Now(),
	}, nil
}

func (s *WinnersService) RecordWin(ctx context.Context, req model.RecordWinRequest) error {
	return s.repo.RecordWin(ctx, req, s.config.BigWinThreshold, s.config.JackpotThreshold)
}

func (s *WinnersService) GetPrivacySettings(ctx context.Context, userID string) (*model.PrivacySettings, error) {
	return s.repo.GetPrivacySettings(ctx, userID)
}

func (s *WinnersService) UpdatePrivacySettings(ctx context.Context, userID string, req model.UpdatePrivacyRequest) error {
	return s.repo.UpdatePrivacySettings(ctx, userID, req)
}

func anonymizeName(username string) string {
	if len(username) <= 2 {
		return username[0:1] + "***"
	}
	return username[0:1] + "***" + username[len(username)-1:]
}
