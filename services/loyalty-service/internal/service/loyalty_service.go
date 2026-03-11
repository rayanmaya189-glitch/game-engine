package service

import (
	"context"
	"fmt"

	"github.com/game-engine/loyalty-service/internal/config"
	"github.com/game-engine/loyalty-service/internal/model"
	"github.com/game-engine/loyalty-service/internal/repository"
	"github.com/google/uuid"
)

type LoyaltyService struct {
	repo *repository.LoyaltyRepository
	cfg  *config.Config
}

func NewLoyaltyService(repo *repository.LoyaltyRepository, cfg *config.Config) *LoyaltyService {
	return &LoyaltyService{repo: repo, cfg: cfg}
}

type GetMemberResponse struct {
	Member *model.LoyaltyMember
}

func (s *LoyaltyService) GetMember(ctx context.Context, userID string) (*GetMemberResponse, error) {
	member, err := s.repo.GetMember(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &GetMemberResponse{Member: member}, nil
}

type GetPointsHistoryResponse struct {
	Transactions []model.PointsTransaction
	Total        int
}

func (s *LoyaltyService) GetPointsHistory(ctx context.Context, userID string, limit, offset int) (*GetPointsHistoryResponse, error) {
	txs, total, err := s.repo.GetPointsHistory(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return &GetPointsHistoryResponse{Transactions: txs, Total: total}, nil
}

type GetTiersResponse struct {
	Tiers []model.Tier
}

func (s *LoyaltyService) GetTiers(ctx context.Context) (*GetTiersResponse, error) {
	tiers, err := s.repo.GetTiers(ctx)
	if err != nil {
		return nil, err
	}
	return &GetTiersResponse{Tiers: tiers}, nil
}

type GetRewardsResponse struct {
	Rewards []model.Reward
}

func (s *LoyaltyService) GetRewards(ctx context.Context) (*GetRewardsResponse, error) {
	rewards, err := s.repo.GetRewards(ctx)
	if err != nil {
		return nil, err
	}
	return &GetRewardsResponse{Rewards: rewards}, nil
}

type RedeemRewardResponse struct {
	Success         bool
	Message         string
	RemainingPoints int
}

func (s *LoyaltyService) RedeemReward(ctx context.Context, userID, rewardID string, pointsCost int) (*RedeemRewardResponse, error) {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	member, err := s.repo.GetMemberTx(ctx, tx, userID)
	if err != nil {
		return nil, fmt.Errorf("member not found for user %s: %w", userID, err)
	}

	if member.Points < pointsCost {
		return nil, fmt.Errorf("insufficient points: have %d, need %d", member.Points, pointsCost)
	}

	newPoints := member.Points - pointsCost
	err = s.repo.UpdateMemberPointsTx(ctx, tx, userID, newPoints, member.LifetimePoints, member.Tier)
	if err != nil {
		return nil, fmt.Errorf("failed to update points: %w", err)
	}

	// Record redemption transaction for audit trail
	redemptionTx := &model.PointsTransaction{
		TransactionID: uuid.New().String(),
		UserID:        userID,
		Amount:        -pointsCost,
		Type:          "debit",
		Source:        "redemption",
		ReferenceID:   rewardID,
		Description:   fmt.Sprintf("Reward redemption: %s", rewardID),
	}
	err = s.repo.AddPointsTransactionTx(ctx, tx, redemptionTx)
	if err != nil {
		return nil, fmt.Errorf("failed to record redemption transaction: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &RedeemRewardResponse{Success: true, Message: "Reward redeemed successfully", RemainingPoints: newPoints}, nil
}

type GetLeaderboardResponse struct {
	Members []model.LoyaltyMember
}

func (s *LoyaltyService) GetLeaderboard(ctx context.Context, limit int) (*GetLeaderboardResponse, error) {
	if limit < 1 || limit > 100 {
		limit = 20
	}

	members, err := s.repo.GetTopMembers(ctx, limit)
	if err != nil {
		return nil, err
	}
	return &GetLeaderboardResponse{Members: members}, nil
}

type AddPointsResponse struct {
	PointsEarned int
	TotalPoints  int
	NewTier      string
	TierUpgraded bool
}

func (s *LoyaltyService) AddPoints(ctx context.Context, userID, betID string, betAmount float64) (*AddPointsResponse, error) {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	member, err := s.repo.GetMemberTx(ctx, tx, userID)
	if err != nil {
		return nil, err
	}

	multiplier := s.cfg.Loyalty.PointsPerBet
	if vipMult, ok := s.cfg.Loyalty.PointsMultiplierVIP[member.Tier]; ok {
		multiplier *= vipMult
	}

	pointsEarned := int(float64(betAmount) * multiplier)
	newPoints := member.Points + pointsEarned
	newLifetime := member.LifetimePoints + pointsEarned

	newTier := s.calculateTier(newLifetime)
	err = s.repo.UpdateMemberPointsTx(ctx, tx, userID, newPoints, newLifetime, newTier)
	if err != nil {
		return nil, fmt.Errorf("failed to update points: %w", err)
	}

	// Record points transaction for audit trail
	pointsTx := &model.PointsTransaction{
		TransactionID: uuid.New().String(),
		UserID:        userID,
		Amount:        pointsEarned,
		Type:          "credit",
		Source:        "bet",
		ReferenceID:   betID,
		Description:   fmt.Sprintf("Points earned from bet: %.2f", betAmount),
	}
	err = s.repo.AddPointsTransactionTx(ctx, tx, pointsTx)
	if err != nil {
		return nil, fmt.Errorf("failed to record points transaction: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &AddPointsResponse{
		PointsEarned: pointsEarned,
		TotalPoints:  newPoints,
		NewTier:      newTier,
		TierUpgraded: newTier != member.Tier,
	}, nil
}

func (s *LoyaltyService) calculateTier(lifetimePoints int) string {
	thresholds := s.cfg.Loyalty.LevelThresholds
	tiers := []string{"bronze", "silver", "gold", "platinum", "diamond", "vip"}

	for i := len(thresholds) - 1; i >= 0; i-- {
		if lifetimePoints >= thresholds[i] {
			return tiers[i]
		}
	}
	return "bronze"
}
