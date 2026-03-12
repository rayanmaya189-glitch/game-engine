package model

import (
	"time"
)

// Table represents a live dealer table
type Table struct {
	TableID     string    `json:"table_id"`
	GameType    string    `json:"game_type"` // blackjack, baccarat, roulette
	DealerName  string    `json:"dealer_name"`
	Status      string    `json:"status"` // open, closed, maintenance
	MinBet      float64   `json:"min_bet"`
	MaxBet      float64   `json:"max_bet"`
	CurrentSeat int       `json:"current_seat"`
	MaxSeats    int       `json:"max_seats"`
	StakeLimit  float64   `json:"stake_limit"`
	DealerID    string    `json:"dealer_id"`
	DeckCount   int       `json:"deck_count"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Player represents a player at a table
type Player struct {
	PlayerID   string    `json:"player_id"`
	TableID    string    `json:"table_id"`
	SeatNumber int       `json:"seat_number"`
	Chips      float64   `json:"chips"`
	CurrentBet float64   `json:"current_bet"`
	JoinedAt   time.Time `json:"joined_at"`
	LastAction string    `json:"last_action"` // bet, hit, stand, split, double
	HandTotal  int       `json:"hand_total"`
	IsFinished bool      `json:"is_finished"`
}

// GameState represents the current state of a game
type GameState struct {
	TableID     string    `json:"table_id"`
	RoundID     string    `json:"round_id"`
	Phase       string    `json:"phase"` // betting, playing, resolving, finished
	Cards       []string  `json:"cards"`
	DealerCards []string  `json:"dealer_cards"`
	Pot         float64   `json:"pot"`
	DealerTotal int       `json:"dealer_total"`
	Winner      string    `json:"winner"` // player, dealer, push
	Payout      float64   `json:"payout"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Dealer represents a live dealer
type Dealer struct {
	DealerID   string    `json:"dealer_id"`
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	Language   string    `json:"language"`
	Status     string    `json:"status"` // available, busy, break
	TableID    string    `json:"table_id"`
	ShiftStart time.Time `json:"shift_start"`
	ShiftEnd   time.Time `json:"shift_end"`
}

// Bet represents a player's bet
type Bet struct {
	BetID      string    `json:"bet_id"`
	PlayerID   string    `json:"player_id"`
	TableID    string    `json:"table_id"`
	RoundID    string    `json:"round_id"`
	BetType    string    `json:"bet_type"` // main, side, bonus
	BetAmount  float64   `json:"bet_amount"`
	Odds       float64   `json:"odds"`
	Potential  float64   `json:"potential"`
	Result     string    `json:"result"` // pending, won, lost, void
	Payout     float64   `json:"payout"`
	PlacedAt   time.Time `json:"placed_at"`
	ResultedAt time.Time `json:"resulted_at"`
}
