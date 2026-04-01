package model

import (
	"time"
)

type ReferralStatus string

const (
	ReferralStatusPending   ReferralStatus = "PENDING"
	ReferralStatusActive    ReferralStatus = "ACTIVE"
	ReferralStatusQualified ReferralStatus = "QUALIFIED"
	ReferralStatusRewarded  ReferralStatus = "REWARDED"
	ReferralStatusExpired   ReferralStatus = "EXPIRED"
)

type RewardType string

const (
	RewardTypeDepositBonus RewardType = "DEPOSIT_BONUS"
	RewardTypeFreeSpins    RewardType = "FREE_SPINS"
	RewardTypeVIPPoints    RewardType = "VIP_POINTS"
	RewardTypeCash         RewardType = "CASH"
)

type Referral struct {
	ID           string         `json:"id"`
	ReferrerID   string         `json:"referrer_id"`
	RefereeID    string         `json:"referee_id"`
	ReferralCode string         `json:"referral_code"`
	Status       ReferralStatus `json:"status"`
	RewardAmount float64        `json:"reward_amount"`
	RewardType   RewardType     `json:"reward_type"`
	RewardClaimed bool          `json:"reward_claimed"`
	Source       string         `json:"source"`
	CampaignID   string         `json:"campaign_id"`
	CreatedAt    time.Time      `json:"created_at"`
	QualifiedAt  *time.Time     `json:"qualified_at"`
	RewardedAt   *time.Time     `json:"rewarded_at"`
	ClaimedAt    *time.Time     `json:"claimed_at"`
}

type ReferralReward struct {
	ID             int64     `json:"id"`
	RewardType     RewardType `json:"reward_type"`
	ReferrerBonus  float64   `json:"referrer_bonus"`
	RefereeBonus   float64   `json:"referee_bonus"`
	MinDeposit     float64   `json:"min_deposit"`
	MinBet         float64   `json:"min_bet"`
	Active         bool      `json:"active"`
	CreatedAt      time.Time `json:"created_at"`
}

type ReferralStats struct {
	TotalReferrals     int64   `json:"total_referrals"`
	ActiveReferrals    int64   `json:"active_referrals"`
	QualifiedReferrals int64   `json:"qualified_referrals"`
	TotalRewardsEarned float64 `json:"total_rewards_earned"`
	TotalRewardsPaid   float64 `json:"total_rewards_paid"`
	PendingRewards     float64 `json:"pending_rewards"`
}

type ReferralCode struct {
	PlayerID    string    `json:"player_id"`
	Code        string    `json:"code"`
	ReferralURL string    `json:"referral_url"`
	CreatedAt   time.Time `json:"created_at"`
	ClickCount  int64     `json:"click_count"`
	SignupCount int64     `json:"signup_count"`
}

type ReferralFilter struct {
	ReferrerID string
	RefereeID  string
	Status     string
	Page       int
	PageSize   int
}

type ReferralList struct {
	Referrals []*Referral `json:"referrals"`
	Total     int64       `json:"total"`
	Page      int         `json:"page"`
	PageSize  int         `json:"page_size"`
}
