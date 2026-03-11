package model

import "time"

type Player struct {
	PlayerID      string    `json:"player_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Status        string    `json:"status"`
	TotalDeposits float64   `json:"total_deposits"`
	TotalBets     float64   `json:"total_bets"`
	Balance       float64   `json:"balance"`
	DepositLimit  float64   `json:"deposit_limit,omitempty"`
	BetLimit      float64   `json:"bet_limit,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

type Dashboard struct {
	TotalPlayers      int     `json:"total_players"`
	ActivePlayers     int     `json:"active_players"`
	TotalCommission   float64 `json:"total_commission"`
	PendingCommission float64 `json:"pending_commission"`
}

type Commission struct {
	CommissionID string     `json:"commission_id"`
	Amount       float64    `json:"amount"`
	Type         string     `json:"type"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	ClaimedAt    *time.Time `json:"claimed_at,omitempty"`
}

type Agent struct {
	AgentID   string    `json:"agent_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	Tier      string    `json:"tier"`
	CreatedAt time.Time `json:"created_at"`
}
