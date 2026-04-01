package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"

	"github.com/game_engine/referral-service/internal/model"
	"github.com/game_engine/referral-service/internal/repository"
)

type ReferralService struct {
	repo *repository.ReferralRepository
}

func NewReferralService(repo *repository.ReferralRepository) *ReferralService {
	return &ReferralService{repo: repo}
}

func (s *ReferralService) GenerateReferralCode(playerID string) (*model.ReferralCode, error) {
	if playerID == "" {
		return nil, errors.New("player_id is required")
	}

	existing, err := s.repo.GetReferralCodeByPlayer(playerID)
	if err == nil && existing != nil {
		return existing, nil
	}

	code := generateCode("REF", 8)

	rc := &model.ReferralCode{
		PlayerID:    playerID,
		Code:        code,
		ReferralURL: fmt.Sprintf("https://casino.example.com/register?ref=%s", code),
		CreatedAt:   time.Now(),
	}

	if err := s.repo.CreateReferralCode(rc); err != nil {
		return nil, err
	}

	return rc, nil
}

func (s *ReferralService) GetReferralCode(playerID string) (*model.ReferralCode, error) {
	if playerID == "" {
		return nil, errors.New("player_id is required")
	}
	return s.repo.GetReferralCodeByPlayer(playerID)
}

func (s *ReferralService) TrackReferral(code, refereeID, source string) (*model.Referral, error) {
	if code == "" {
		return nil, errors.New("referral code is required")
	}
	if refereeID == "" {
		return nil, errors.New("referee_id is required")
	}

	referrerID, err := s.repo.GetReferrerByCode(code)
	if err != nil {
		return nil, fmt.Errorf("invalid referral code: %w", err)
	}

	if referrerID == refereeID {
		return nil, errors.New("cannot refer yourself")
	}

	referral := &model.Referral{
		ID:           uuid.New().String(),
		ReferrerID:   referrerID,
		RefereeID:    refereeID,
		ReferralCode: code,
		Status:       model.ReferralStatusActive,
		Source:       source,
		CreatedAt:    time.Now(),
	}

	if err := s.repo.CreateReferral(referral); err != nil {
		return nil, err
	}

	_ = s.repo.TrackReferralSignup(code)

	return referral, nil
}

func (s *ReferralService) QualifyReferral(referralID string) error {
	return s.repo.MarkQualified(referralID)
}

func (s *ReferralService) RewardReferral(referralID string, amount float64, rewardType model.RewardType) error {
	if amount <= 0 {
		return errors.New("reward amount must be positive")
	}
	return s.repo.MarkRewarded(referralID, amount, rewardType)
}

func (s *ReferralService) ClaimReward(referralID string, playerID string) error {
	referral, err := s.repo.GetReferralByID(referralID)
	if err != nil {
		return err
	}

	if referral.ReferrerID != playerID {
		return errors.New("not authorized to claim this reward")
	}

	if referral.Status != model.ReferralStatusRewarded {
		return errors.New("reward not available for claiming")
	}

	if referral.RewardClaimed {
		return errors.New("reward already claimed")
	}

	return s.repo.ClaimReward(referralID)
}

func (s *ReferralService) GetReferrals(referrerID string, filter *model.ReferralFilter) (*model.ReferralList, error) {
	if referrerID == "" {
		return nil, errors.New("referrer_id is required")
	}
	if filter == nil {
		filter = &model.ReferralFilter{Page: 1, PageSize: 20}
	}
	return s.repo.GetReferralsByReferrer(referrerID, filter)
}

func (s *ReferralService) GetReferralStats(playerID string) (*model.ReferralStats, error) {
	if playerID == "" {
		return nil, errors.New("player_id is required")
	}
	return s.repo.GetReferralStats(playerID)
}

func (s *ReferralService) GetAvailableRewards() ([]*model.ReferralReward, error) {
	return s.repo.GetActiveRewards()
}

func (s *ReferralService) TrackClick(code string) error {
	return s.repo.TrackReferralClick(code)
}

func (s *ReferralService) CalculateMultiTierReward(referrerID string, amount float64, rate float64) (float64, error) {
	childReferrals, err := s.repo.GetChildReferrals(referrerID)
	if err != nil {
		return 0, err
	}

	totalMultiTier := 0.0
	for _, child := range childReferrals {
		totalMultiTier += child.RewardAmount * rate
	}

	return totalMultiTier, nil
}

func generateCode(prefix string, length int) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := range result {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		result[i] = chars[n.Int64()]
	}

	return prefix + string(result)
}
