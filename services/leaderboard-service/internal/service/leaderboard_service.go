package service

import (
	"context"
	"time"

	"github.com/game_engine/leaderboard-service/internal/config"
	"github.com/game_engine/leaderboard-service/internal/model"
	"github.com/game_engine/leaderboard-service/internal/repository"
)

type LeaderboardService struct {
	repo   *repository.LeaderboardRepository
	config *config.LeaderboardConfig
}

func NewLeaderboardService(repo *repository.LeaderboardRepository, cfg *config.LeaderboardConfig) *LeaderboardService {
	return &LeaderboardService{repo: repo, config: cfg}
}

func (s *LeaderboardService) GetDailyLeaderboard(ctx context.Context, gameType string) (*model.LeaderboardResponse, error) {
	return s.getLeaderboard(ctx, model.LeaderboardTypeDaily, gameType)
}

func (s *LeaderboardService) GetWeeklyLeaderboard(ctx context.Context, gameType string) (*model.LeaderboardResponse, error) {
	return s.getLeaderboard(ctx, model.LeaderboardTypeWeekly, gameType)
}

func (s *LeaderboardService) GetMonthlyLeaderboard(ctx context.Context, gameType string) (*model.LeaderboardResponse, error) {
	return s.getLeaderboard(ctx, model.LeaderboardTypeMonthly, gameType)
}

func (s *LeaderboardService) GetAllTimeLeaderboard(ctx context.Context, gameType string) (*model.LeaderboardResponse, error) {
	return s.getLeaderboard(ctx, model.LeaderboardTypeAllTime, gameType)
}

func (s *LeaderboardService) getLeaderboard(ctx context.Context, leaderboardType model.LeaderboardType, gameType string) (*model.LeaderboardResponse, error) {
	if gameType == "" {
		gameType = "all"
	}

	entries, err := s.repo.GetLeaderboard(ctx, leaderboardType, gameType, s.config.TopPlayersCount)
	if err != nil {
		return nil, err
	}

	return &model.LeaderboardResponse{
		Type:      leaderboardType,
		Period:    getPeriodString(leaderboardType),
		Entries:   entries,
		Total:     len(entries),
		UpdatedAt: time.Now(),
	}, nil
}

func (s *LeaderboardService) GetPlayerDailyRank(ctx context.Context, userID string, gameType string) (*model.PlayerRankResponse, error) {
	return s.getPlayerRank(ctx, userID, model.LeaderboardTypeDaily, gameType)
}

func (s *LeaderboardService) GetPlayerWeeklyRank(ctx context.Context, userID string, gameType string) (*model.PlayerRankResponse, error) {
	return s.getPlayerRank(ctx, userID, model.LeaderboardTypeWeekly, gameType)
}

func (s *LeaderboardService) GetPlayerMonthlyRank(ctx context.Context, userID string, gameType string) (*model.PlayerRankResponse, error) {
	return s.getPlayerRank(ctx, userID, model.LeaderboardTypeMonthly, gameType)
}

func (s *LeaderboardService) GetPlayerAllTimeRank(ctx context.Context, userID string, gameType string) (*model.PlayerRankResponse, error) {
	return s.getPlayerRank(ctx, userID, model.LeaderboardTypeAllTime, gameType)
}

func (s *LeaderboardService) getPlayerRank(ctx context.Context, userID string, leaderboardType model.LeaderboardType, gameType string) (*model.PlayerRankResponse, error) {
	if gameType == "" {
		gameType = "all"
	}
	return s.repo.GetPlayerRank(ctx, userID, leaderboardType, gameType)
}

func (s *LeaderboardService) UpdatePlayerScore(ctx context.Context, req model.UpdateScoreRequest) error {
	return s.repo.UpdatePlayerScore(ctx, req)
}

func getPeriodString(leaderboardType model.LeaderboardType) string {
	now := time.Now()
	switch leaderboardType {
	case model.LeaderboardTypeDaily:
		return now.Format("2006-01-02")
	case model.LeaderboardTypeWeekly:
		year, week := now.ISOWeek()
		return "2024-W01"
	case model.LeaderboardTypeMonthly:
		return now.Format("2006-01")
	case model.LeaderboardTypeAllTime:
		return "all"
	default:
		return now.Format("2006-01-02")
	}
}
