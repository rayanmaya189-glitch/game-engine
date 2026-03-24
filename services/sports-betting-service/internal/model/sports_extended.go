package model

import (
	"time"
)

// BetType represents the type of bet
type BetType string

const (
	BetTypeSingle      BetType = "single"
	BetTypeAccumulator BetType = "accumulator"
	BetTypeSystem      BetType = "system"
	BetTypeParlay      BetType = "parlay"
)

// BetStatus represents the status of a bet
type BetStatus string

const (
	BetStatusPending   BetStatus = "pending"
	BetStatusWon       BetStatus = "won"
	BetStatusLost      BetStatus = "lost"
	BetStatusCancelled BetStatus = "cancelled"
	BetStatusVoided    BetStatus = "voided"
	BetStatusCashOut   BetStatus = "cash_out"
)

// LiveEvent represents a live/in-play sporting event
type LiveEvent struct {
	EventID     string       `json:"event_id"`
	SportID     string       `json:"sport_id"`
	HomeTeam    string       `json:"home_team"`
	AwayTeam    string       `json:"away_team"`
	HomeScore   int          `json:"home_score"`
	AwayScore   int          `json:"away_score"`
	Period      string       `json:"period"` // 1st Half, 2nd Half, Quarter 1, etc.
	Minute      int          `json:"minute"` // Current minute in play
	Status      string       `json:"status"` // live, halftime, finished
	StartTime   time.Time    `json:"start_time"`
	Markets     []Market     `json:"markets"`
	OddsChanges []OddsChange `json:"odds_changes"` // Track odds history
	LastUpdate  time.Time    `json:"last_update"`
}

// OddsChange represents a change in odds over time
type OddsChange struct {
	MarketID  string    `json:"market_id"`
	Selection string    `json:"selection"`
	OldOdds   float64   `json:"old_odds"`
	NewOdds   float64   `json:"new_odds"`
	Timestamp time.Time `json:"timestamp"`
}

// ParlayBet represents a parlay (accumulator) bet
type ParlayBet struct {
	BetID        string            `json:"bet_id"`
	UserID       string            `json:"user_id"`
	Selections   []ParlaySelection `json:"selections"`
	Stake        float64           `json:"stake"`
	TotalOdds    float64           `json:"total_odds"`
	PotentialWin float64           `json:"potential_win"`
	Status       BetStatus         `json:"status"`
	PlacedAt     time.Time         `json:"placed_at"`
	SettledAt    *time.Time        `json:"settled_at,omitempty"`
}

// ParlaySelection represents a single selection in a parlay bet
type ParlaySelection struct {
	EventID   string  `json:"event_id"`
	EventName string  `json:"event_name"`
	MarketID  string  `json:"market_id"`
	Selection string  `json:"selection"` // Home, Draw, Over, etc.
	Odds      float64 `json:"odds"`
	Status    string  `json:"status"` // pending, won, lost, voided
	Result    *string `json:"result,omitempty"`
}

// CashOut represents a cash out request
type CashOut struct {
	CashOutID     string     `json:"cash_out_id"`
	BetID         string     `json:"bet_id"`
	UserID        string     `json:"user_id"`
	OriginalStake float64    `json:"original_stake"`
	OriginalOdds  float64    `json:"original_odds"`
	CurrentOdds   float64    `json:"current_odds"`
	CashOutAmount float64    `json:"cash_out_amount"`
	Profit        float64    `json:"profit"`
	Status        string     `json:"status"` // pending, completed, failed
	RequestedAt   time.Time  `json:"requested_at"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
}

// CashOutCalculation represents the calculation for cash out value
type CashOutCalculation struct {
	BetID            string  `json:"bet_id"`
	OriginalStake    float64 `json:"original_stake"`
	OriginalOdds     float64 `json:"original_odds"`
	CurrentOdds      float64 `json:"current_odds"`
	CurrentValue     float64 `json:"current_value"`
	PotentialWin     float64 `json:"potential_win"`
	Profit           float64 `json:"profit"`
	Eligible         bool    `json:"eligible"`
	IneligibleReason string  `json:"ineligible_reason,omitempty"`
}

// LiveOddsUpdate represents live odds update for WebSocket
type LiveOddsUpdate struct {
	EventID   string    `json:"event_id"`
	Markets   []Market  `json:"markets"`
	Timestamp time.Time `json:"timestamp"`
}

// BetBuilderSelection represents a selection in the bet builder
type BetBuilderSelection struct {
	EventID    string  `json:"event_id"`
	EventName  string  `json:"event_name"`
	MarketType string  `json:"market_type"` // moneyline, spread, total, etc.
	Selection  string  `json:"selection"`
	Odds       float64 `json:"odds"`
}

// BetBuilder calculates parlay odds
type BetBuilder struct {
	Selections   []BetBuilderSelection `json:"selections"`
	Stake        float64               `json:"stake"`
	TotalOdds    float64               `json:"total_odds"`
	PotentialWin float64               `json:"potential_wins"`
}

// SystemBet represents a system bet (e.g., Patent, Yankee, Canadian)
type SystemBet struct {
	BetID         string            `json:"bet_id"`
	UserID        string            `json:"user_id"`
	BetType       BetType           `json:"bet_type"`
	SystemType    string            `json:"system_type"` // Patent (3), Yankee (4), etc.
	Selections    []ParlaySelection `json:"selections"`
	Stake         float64           `json:"stake"`
	NumSelections int               `json:"num_selections"`
	NumWins       int               `json:"num_wins"` // For calculation
	TotalOdds     float64           `json:"total_odds"`
	PotentialWin  float64           `json:"potential_win"`
	Status        BetStatus         `json:"status"`
	PlacedAt      time.Time         `json:"placed_at"`
	SettledAt     *time.Time        `json:"settled_at,omitempty"`
}

// LiveBettingConfig holds configuration for live betting
type LiveBettingConfig struct {
	MinLiveOdds        float64 `json:"min_live_odds"`
	MaxLiveOdds        float64 `json:"max_live_odds"`
	AutoSuspend        bool    `json:"auto_suspend"`      // Auto-suspend when odds change too fast
	SuspendThreshold   float64 `json:"suspend_threshold"` // Threshold for auto-suspend
	CashOutEnabled     bool    `json:"cash_out_enabled"`
	MaxParlaySelection int     `json:"max_parlay_selection"`
}

// EventStatusUpdate represents an update to event status
type EventStatusUpdate struct {
	EventID   string    `json:"event_id"`
	Status    string    `json:"status"`
	HomeScore int       `json:"home_score,omitempty"`
	AwayScore int       `json:"away_score,omitempty"`
	Period    string    `json:"period,omitempty"`
	Minute    int       `json:"minute,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// MarketSuspension represents a suspended market
type MarketSuspension struct {
	MarketID    string     `json:"market_id"`
	EventID     string     `json:"event_id"`
	Reason      string     `json:"reason"`
	SuspendedAt time.Time  `json:"suspended_at"`
	ResumeAt    *time.Time `json:"resume_at,omitempty"`
}
