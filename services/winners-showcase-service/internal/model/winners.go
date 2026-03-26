package model

import "time"

// WinType represents the type of win
type WinType string

const (
	WinTypeRegular     WinType = "regular"
	WinTypeBig         WinType = "big"
	WinTypeJackpot     WinType = "jackpot"
	WinTypeProgressive WinType = "progressive"
)

// Winner represents a winner record
type Winner struct {
	ID          int       `json:"id"`
	UserID      string    `json:"user_id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"` // Based on privacy settings
	WinAmount   float64   `json:"win_amount"`
	Currency    string    `json:"currency"`
	GameType    string    `json:"game_type"`
	GameName    string    `json:"game_name"`
	WinType     WinType   `json:"win_type"`
	Multiplier  float64   `json:"multiplier"`
	Timestamp   time.Time `json:"timestamp"`
}

// RecentWinnersResponse represents the API response for recent winners
type RecentWinnersResponse struct {
	Winners   []Winner  `json:"winners"`
	Total     int       `json:"total"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BigWinsResponse represents the API response for big wins
type BigWinsResponse struct {
	Wins      []Winner  `json:"wins"`
	Threshold float64   `json:"threshold"`
	Total     int       `json:"total"`
	UpdatedAt time.Time `json:"updated_at"`
}

// JackpotWinnersResponse represents the API response for jackpot winners
type JackpotWinnersResponse struct {
	Winners   []Winner  `json:"winners"`
	Threshold float64   `json:"threshold"`
	Total     int       `json:"total"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RecordWinRequest represents a request to record a new win
type RecordWinRequest struct {
	UserID     string  `json:"user_id" binding:"required"`
	Username   string  `json:"username" binding:"required"`
	WinAmount  float64 `json:"win_amount" binding:"required"`
	Currency   string  `json:"currency" binding:"required"`
	GameType   string  `json:"game_type" binding:"required"`
	GameName   string  `json:"game_name" binding:"required"`
	Multiplier float64 `json:"multiplier"`
}

// PrivacySettings represents player's privacy settings for winner display
type PrivacySettings struct {
	UserID            string `json:"user_id"`
	AnonymizeName     bool   `json:"anonymize_name"`
	ShowOnLeaderboard bool   `json:"show_on_leaderboard"`
	ShowOnJackpotList bool   `json:"show_on_jackpot_list"`
	OptOutOfShowcase  bool   `json:"opt_out_of_showcase"`
}

// UpdatePrivacyRequest represents a request to update privacy settings
type UpdatePrivacyRequest struct {
	AnonymizeName     bool `json:"anonymize_name"`
	ShowOnLeaderboard bool `json:"show_on_leaderboard"`
	ShowOnJackpotList bool `json:"show_on_jackpot_list"`
	OptOutOfShowcase  bool `json:"opt_out_of_showcase"`
}
