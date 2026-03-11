package model

import "time"

type MerchantPlayer struct {
	PlayerID string `json:"player_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

type Agent struct {
	AgentID  string `json:"agent_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}

type Webhook struct {
	WebhookID string `json:"webhook_id"`
	URL       string `json:"url"`
	Events    string `json:"events"`
	Status    string `json:"status"`
}

type RevenueReport struct {
	TotalRevenue     float64 `json:"total_revenue"`
	TotalDeposits    float64 `json:"total_deposits"`
	TotalWithdrawals float64 `json:"total_withdrawals"`
	TotalPlayers     int     `json:"total_players"`
}

type PlayerReport struct {
	TotalBets   float64 `json:"total_bets"`
	TotalWins   float64 `json:"total_wins"`
	NetRevenue  float64 `json:"net_revenue"`
	GamesPlayed int     `json:"games_played"`
}

type GameReport struct {
	TotalBets    float64 `json:"total_bets"`
	TotalWins    float64 `json:"total_wins"`
	TotalPlayers int     `json:"total_players"`
	Plays        int     `json:"plays"`
}

type Merchant struct {
	MerchantID     string    `json:"merchant_id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Status         string    `json:"status"`
	CommissionRate float64   `json:"commission_rate"`
	CreatedAt      time.Time `json:"created_at"`
}
