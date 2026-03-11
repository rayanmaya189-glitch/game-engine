package model

import "time"

// Member represents a loyalty program member (alias for LoyaltyMember)
type Member = LoyaltyMember

// LoyaltyMember represents a loyalty program member
type LoyaltyMember struct {
	UserID         string    `json:"user_id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Points         int       `json:"points"`
	LifetimePoints int       `json:"lifetime_points"`
	Tier           string    `json:"tier"`   // bronze, silver, gold, platinum, diamond, vip
	Status         string    `json:"status"` // active, suspended, expired
	JoinedAt       time.Time `json:"joined_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// PointsTransaction represents a points transaction record
type PointsTransaction struct {
	TransactionID string    `json:"transaction_id"`
	UserID        string    `json:"user_id"`
	Amount        int       `json:"amount"`
	Type          string    `json:"type"`   // credit, debit
	Source        string    `json:"source"` // bet, redemption, bonus, etc.
	ReferenceID   string    `json:"reference_id"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
}

// Tier represents a loyalty tier
type Tier struct {
	TierID           string  `json:"tier_id"`
	Name             string  `json:"name"`
	MinPoints        int     `json:"min_points"`
	MaxPoints        int     `json:"max_points"`
	PointsMultiplier float64 `json:"points_multiplier"`
	CashbackPercent  float64 `json:"cashback_percent"`
	Benefits         string  `json:"benefits"`
}

// Reward represents a redeemable reward
type Reward struct {
	RewardID    string    `json:"reward_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	PointsCost  int       `json:"points_cost"`
	Type        string    `json:"type"` // bonus, free_spin, merchandise, etc.
	Value       float64   `json:"value"`
	Status      string    `json:"status"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// Redemption represents a reward redemption
type Redemption struct {
	RedemptionID string    `json:"redemption_id"`
	UserID       string    `json:"user_id"`
	RewardID     string    `json:"reward_id"`
	PointsSpent  int       `json:"points_spent"`
	Status       string    `json:"status"`
	RedeemedAt   time.Time `json:"redeemed_at"`
}
