package model

import "time"

// LeaderboardType represents the type of leaderboard
type LeaderboardType string

const (
	LeaderboardTypeDaily   LeaderboardType = "daily"
	LeaderboardTypeWeekly  LeaderboardType = "weekly"
	LeaderboardTypeMonthly LeaderboardType = "monthly"
	LeaderboardTypeAllTime LeaderboardType = "alltime"
)

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
}
