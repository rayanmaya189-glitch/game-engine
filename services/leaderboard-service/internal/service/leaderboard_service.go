package service

import (
	"context"
	"fmt"
	"log"
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
	// Anti-gaming validation
	if err := req.Validate(s.config.MinBetThreshold); err != nil {
		return err
	}

	// Use Redis for high-performance score updates
	return s.repo.RedisUpdateScore(ctx, req.UserID, model.LeaderboardTypeDaily, req.GameType, req.Score)
}

func (s *LeaderboardService) UpdatePlayerScoreWithPeriod(ctx context.Context, req model.UpdateScoreRequest, leaderboardType model.LeaderboardType) error {
	// Anti-gaming validation
	if err := req.Validate(s.config.MinBetThreshold); err != nil {
		return err
	}

	// Update all relevant leaderboard types
	leaderboardTypes := []model.LeaderboardType{
		model.LeaderboardTypeDaily,
		model.LeaderboardTypeWeekly,
		model.LeaderboardTypeMonthly,
		model.LeaderboardTypeAllTime,
	}

	// Add biggest win tracking if this was a win
	if req.IsWin && req.WinAmount > 0 {
		leaderboardTypes = append(leaderboardTypes, model.LeaderboardTypeBiggestWin)
	}

	// Add most active tracking
	leaderboardTypes = append(leaderboardTypes, model.LeaderboardTypeMostActive)

	// Update each leaderboard type in Redis
	for _, lbType := range leaderboardTypes {
		if err := s.repo.RedisUpdateScore(ctx, req.UserID, lbType, req.GameType, req.Score); err != nil {
			log.Printf("Failed to update score for leaderboard type %s: %v", lbType, err)
		}
	}

	// Also update database for persistence
	return s.repo.UpdatePlayerScore(ctx, req)
}

func (s *LeaderboardService) DistributePrizes(ctx context.Context, req model.PrizeDistributionRequest) (*model.PrizeDistribution, error) {
	// Get prize configuration for this leaderboard type
	prizeKey := string(req.LeaderboardType)
	if req.GameType != "" {
		prizeKey = fmt.Sprintf("%s_%s", req.LeaderboardType, req.GameType)
	}

	prizes, ok := s.config.Prizes[prizeKey]
	if !ok {
		return nil, fmt.Errorf("no prize configuration found for %s", prizeKey)
	}

	// Get top players for the leaderboard
	entries, err := s.repo.GetLeaderboard(ctx, req.LeaderboardType, req.GameType, len(prizes))
	if err != nil {
		return nil, err
	}

	// Calculate prize distribution
	distribution := &model.PrizeDistribution{
		LeaderboardType: req.LeaderboardType,
		GameType:        req.GameType,
		Period:          getPeriodString(req.LeaderboardType),
		Prizes:          []model.Prize{},
		TotalValue:      0,
		DistributedAt:   time.Now(),
	}

	for _, entry := range entries {
		for _, prizeConfig := range prizes {
			if entry.Rank >= prizeConfig.FromRank && entry.Rank <= prizeConfig.ToRank {
				prize := model.Prize{
					Rank:  entry.Rank,
					Type:  prizeConfig.Type,
					Value: prizeConfig.Value,
				}
				distribution.Prizes = append(distribution.Prizes, prize)
				distribution.TotalValue += prizeConfig.Value

				// In production, this would trigger wallet credit
				if !req.DryRun && s.config.PrizeAutoCredit {
					log.Printf("Would credit prize to user %s: %+v", entry.UserID, prize)
				}
				break
			}
		}
	}

	if req.DryRun {
		log.Printf("DRY RUN: Prize distribution calculated: %+v", distribution)
	}

	return distribution, nil
}

func (s *LeaderboardService) SyncLeaderboard(ctx context.Context, leaderboardType model.LeaderboardType, gameType string) error {
	return s.repo.RedisSyncFromDB(ctx, leaderboardType, gameType, s.config.TopPlayersCount)
}

func (s *LeaderboardService) ResetLeaderboard(ctx context.Context, leaderboardType model.LeaderboardType, gameType string) error {
	return s.repo.RedisResetLeaderboard(ctx, leaderboardType, gameType)
}

func getPeriodString(leaderboardType model.LeaderboardType) string {
	now := time.Now()
	switch leaderboardType {
	case model.LeaderboardTypeDaily:
		return now.Format("2006-01-02")
	case model.LeaderboardTypeWeekly:
		year, week := now.ISOWeek()
		return fmt.Sprintf("%d-W%02d", year, week)
	case model.LeaderboardTypeMonthly:
		return now.Format("2006-01")
	case model.LeaderboardTypeAllTime:
		return "all"
	default:
		return now.Format("2006-01-02")
	}
}
