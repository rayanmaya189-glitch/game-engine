package model

import "time"

type Jackpot struct {
	JackpotID     string    `json:"jackpot_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	CurrentAmount float64   `json:"current_amount"`
	MinBet        float64   `json:"min_bet"`
	MaxBet        float64   `json:"max_bet"`
	Status        string    `json:"status"`
	StartsAt      time.Time `json:"starts_at"`
	EndsAt        time.Time `json:"ends_at"`
}

type Winner struct {
	WinnerID string    `json:"winner_id"`
	Username string    `json:"username"`
	Amount   float64   `json:"amount"`
	WonAt    time.Time `json:"won_at"`
}

type JackpotHistoryEntry struct {
	JackpotID   string    `json:"jackpot_id"`
	JackpotName string    `json:"jackpot_name"`
	Amount      float64   `json:"amount"`
	Result      string    `json:"result"`
	PlayedAt    time.Time `json:"played_at"`
}
