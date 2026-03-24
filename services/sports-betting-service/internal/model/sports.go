package model

import "time"

// Sport represents a sport type
type Sport struct {
	SportID   string `json:"sport_id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Status    string `json:"status"`
	SortOrder int    `json:"sort_order"`
}

// Event represents a sports event
type Event struct {
	EventID   string    `json:"event_id"`
	SportID   string    `json:"sport_id"`
	LeagueID  string    `json:"league_id"`
	HomeTeam  string    `json:"home_team"`
	AwayTeam  string    `json:"away_team"`
	HomeScore int       `json:"home_score"`
	AwayScore int       `json:"away_score"`
	Status    string    `json:"status"` // scheduled, live, completed, cancelled
	Period    string    `json:"period"` // 1st Half, 2nd Quarter, etc.
	Minute    int       `json:"minute"` // Current minute/quarter time
	StartTime time.Time `json:"start_time"`
	CreatedAt time.Time `json:"created_at"`
}

// Market represents a betting market
type Market struct {
	MarketID   string      `json:"market_id"`
	EventID    string      `json:"event_id"`
	Name       string      `json:"name"`        // 1X2, over/under, handicap, etc.
	MarketType string      `json:"market_type"` // moneyline, spread, total
	Status     string      `json:"status"`      // open, closed, suspended, settled
	HomeOdds   float64     `json:"home_odds"`   // Legacy compatibility
	DrawOdds   float64     `json:"draw_odds"`   // Legacy compatibility
	AwayOdds   float64     `json:"away_odds"`   // Legacy compatibility
	Selections []Selection `json:"selections"`
}

// Selection represents a bet selection within a market
type Selection struct {
	SelectionID string          `json:"selection_id"`
	Selection   string          `json:"selection"` // home, away, over, under, etc.
	Odds        float64         `json:"odds"`
	Result      SelectionResult `json:"result"`
}

// SelectionResult represents the result of a selection
type SelectionResult struct {
	Valid bool `json:"valid"`
	Win   bool `json:"win"`
}

// Bet represents a user's bet
type Bet struct {
	BetID        string     `json:"bet_id"`
	UserID       string     `json:"user_id"`
	EventID      string     `json:"event_id"`
	MarketID     string     `json:"market_id"`
	Selection    string     `json:"selection"` // home, draw, away
	Stake        float64    `json:"stake"`
	Odds         float64    `json:"odds"`
	PotentialWin float64    `json:"potential_win"`
	Status       string     `json:"status"` // pending, won, lost, cancelled, cash_out
	PlacedAt     time.Time  `json:"placed_at"`
	SettledAt    *time.Time `json:"settled_at,omitempty"`
}
