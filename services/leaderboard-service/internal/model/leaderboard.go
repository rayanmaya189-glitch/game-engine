package model

import (
	"fmt"
	"time"
)

// LeaderboardType represents the type of leaderboard
type LeaderboardType string

const (
	LeaderboardTypeDaily        LeaderboardType = "daily"
	LeaderboardTypeWeekly       LeaderboardType = "weekly"
	LeaderboardTypeMonthly      LeaderboardType = "monthly"
	LeaderboardTypeAllTime      LeaderboardType = "alltime"
	LeaderboardTypeBiggestWin   LeaderboardType = "biggest_win"
	LeaderboardTypeMostActive   LeaderboardType = "most_active"
	LeaderboardTypeTournament   LeaderboardType = "tournament"
	LeaderboardTypeVIPPoints    LeaderboardType = "vip_points"
	LeaderboardTypeGameSpecific LeaderboardType = "game_specific"
)

// Additional Leaderboard Types as per Phase 6 plan:
// - Daily Winners (00:00-23:59 UTC) - Real-time
// - Weekly Winners (Monday 00:00 - Sunday 23:59 UTC) - Real-time
// - Monthly Winners (1st - Last day of month) - Real-time
// - All-Time Winners - Hourly updates
// - Biggest Win - Rolling 24 hours, Real-time
// - Most Active - Rolling 24 hours, Real-time
// - Game-Specific - Per game, configurable period
// - Tournament - During tournament, Real-time
// - VIP Points - Rolling 30 days, Real-time

// LeaderboardEntry represents a single entry on the leaderboard
type LeaderboardEntry struct {
	Rank      int       `json:"rank"`
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Score     float64   `json:"score"`
	Wins      int       `json:"wins"`
	WinAmount float64   `json:"win_amount"`
	GameType  string    `json:"game_type,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

// PlayerScore represents a player's score record
type PlayerScore struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Score     float64   `json:"score"`
	Wins      int       `json:"wins"`
	WinAmount float64   `json:"win_amount"`
	GameType  string    `json:"game_type"`
	Period    string    `json:"period"` // daily_YYYY-MM-DD, weekly_2024-W01, monthly_2024-01
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LeaderboardResponse represents the API response for leaderboard
type LeaderboardResponse struct {
	Type      LeaderboardType    `json:"type"`
	Period    string             `json:"period"`
	Entries   []LeaderboardEntry `json:"entries"`
	Total     int                `json:"total"`
	UpdatedAt time.Time          `json:"updated_at"`
}

// PlayerRankResponse represents the API response for player rank
type PlayerRankResponse struct {
	UserID          string          `json:"user_id"`
	Username        string          `json:"username"`
	Rank            int             `json:"rank"`
	Score           float64         `json:"score"`
	LeaderboardType LeaderboardType `json:"type"`
	Period          string          `json:"period"`
}

// UpdateScoreRequest represents a request to update player score
type UpdateScoreRequest struct {
	UserID    string  `json:"user_id" binding:"required"`
	Username  string  `json:"username" binding:"required"`
	Score     float64 `json:"score" binding:"required"`
	GameType  string  `json:"game_type" binding:"required"`
	IsWin     bool    `json:"is_win"`
	WinAmount float64 `json:"win_amount"`
	BetAmount float64 `json:"bet_amount"` // Added for anti-gaming validation
}

// Validate validates the score update request for anti-gaming rules
func (r *UpdateScoreRequest) Validate(minBetThreshold float64) error {
	// Anti-gaming rule: Minimum bet threshold must be met
	if minBetThreshold > 0 && r.BetAmount > 0 && r.BetAmount < minBetThreshold {
		return fmt.Errorf("bet amount %.2f is below minimum threshold %.2f", r.BetAmount, minBetThreshold)
	}

	// Anti-gaming rule: Maximum score per update should not be absurdly high
	if r.Score > 1000000 {
		return fmt.Errorf("score exceeds maximum allowed value")
	}

	// Validate bet amount is not negative
	if r.BetAmount < 0 {
		return fmt.Errorf("bet amount cannot be negative")
	}

	return nil
}

// PrizeDistribution represents prize distribution for a leaderboard period
type PrizeDistribution struct {
	ID              int             `json:"id"`
	LeaderboardType LeaderboardType `json:"leaderboard_type"`
	GameType        string          `json:"game_type"`
	Period          string          `json:"period"`
	UserID          string          `json:"user_id"`
	Rank            int             `json:"rank"`
	PrizeType       string          `json:"prize_type"`
	Value           float64         `json:"value"`
	Currency        string          `json:"currency"`
	Prizes          []Prize         `json:"prizes,omitempty"`
	TotalValue      float64         `json:"total_value"`
	DistributedAt   time.Time       `json:"distributed_at"`
	Status          string          `json:"status"`
}

// Prize represents a single prize
type Prize struct {
	Rank     int     `json:"rank"`
	Type     string  `json:"type"` // bonus, freespins, vip_points, merchandise
	Value    float64 `json:"value"`
	Currency string  `json:"currency,omitempty"`
}

// PrizeDistributionRequest represents a request to distribute prizes
type PrizeDistributionRequest struct {
	LeaderboardType LeaderboardType `json:"leaderboard_type" binding:"required"`
	GameType        string          `json:"game_type"`
	TournamentID    string          `json:"tournament_id"` // For tournament-specific prizes
	DryRun          bool            `json:"dry_run"`
}
