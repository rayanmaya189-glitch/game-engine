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
	StartTime time.Time `json:"start_time"`
	CreatedAt time.Time `json:"created_at"`
}

// Market represents a betting market
type Market struct {
	MarketID string  `json:"market_id"`
	EventID  string  `json:"event_id"`
	Name     string  `json:"name"`   // 1X2, over/under, handicap, etc.
	Status   string  `json:"status"` // open, closed, settled
	HomeOdds float64 `json:"home_odds"`
	DrawOdds float64 `json:"draw_odds"`
	AwayOdds float64 `json:"away_odds"`
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
	Status       string     `json:"status"` // pending, won, lost, cancelled
	PlacedAt     time.Time  `json:"placed_at"`
	SettledAt    *time.Time `json:"settled_at,omitempty"`
}
