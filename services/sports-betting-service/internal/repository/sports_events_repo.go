package repository

import (
	"context"
	"fmt"

	"github.com/game_engine/sports-betting-service/internal/model"
)

func (r *SportsRepository) GetLiveEvents(ctx context.Context) ([]model.Event, error) {
	rows, err := r.db.Query(ctx, `
		SELECT event_id, sport_id, league_id, home_team, away_team, start_time, status, home_score, away_score, created_at
		FROM sports_events WHERE status = 'live' ORDER BY start_time
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var e model.Event
		if err := rows.Scan(&e.EventID, &e.SportID, &e.LeagueID, &e.HomeTeam, &e.AwayTeam, &e.StartTime, &e.Status, &e.HomeScore, &e.AwayScore, &e.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func (r *SportsRepository) GetUpcomingEvents(ctx context.Context, sportID string, limit int) ([]model.Event, error) {
	rows, err := r.db.Query(ctx, `
		SELECT event_id, sport_id, league_id, home_team, away_team, start_time, status, home_score, away_score, created_at
		FROM sports_events WHERE sport_id = $1 AND status = 'scheduled' AND start_time > NOW()
		ORDER BY start_time LIMIT $2
	`, sportID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []model.Event
	for rows.Next() {
		var e model.Event
		if err := rows.Scan(&e.EventID, &e.SportID, &e.LeagueID, &e.HomeTeam, &e.AwayTeam, &e.StartTime, &e.Status, &e.HomeScore, &e.AwayScore, &e.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func (r *SportsRepository) GetEventByID(ctx context.Context, eventID string) (*model.Event, error) {
	var e model.Event
	err := r.db.QueryRow(ctx, `
		SELECT event_id, sport_id, league_id, home_team, away_team, start_time, status, home_score, away_score, created_at
		FROM sports_events WHERE event_id = $1
	`, eventID).Scan(&e.EventID, &e.SportID, &e.LeagueID, &e.HomeTeam, &e.AwayTeam, &e.StartTime, &e.Status, &e.HomeScore, &e.AwayScore, &e.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("event not found: %w", err)
	}
	return &e, nil
}

// GetLiveEventByID returns a live event by ID
func (r *SportsRepository) GetLiveEventByID(ctx context.Context, eventID string) (*model.Event, error) {
	var e model.Event
	err := r.db.QueryRow(ctx, `
		SELECT event_id, sport_id, league_id, home_team, away_team, start_time, status, home_score, away_score, created_at, period, minute
		FROM sports_events WHERE event_id = $1
	`, eventID).Scan(&e.EventID, &e.SportID, &e.LeagueID, &e.HomeTeam, &e.AwayTeam, &e.StartTime, &e.Status, &e.HomeScore, &e.AwayScore, &e.CreatedAt, &e.Period, &e.Minute)
	if err != nil {
		return nil, fmt.Errorf("event not found: %w", err)
	}
	return &e, nil
}

// UpdateEventStatus updates event status, score, period, and minute
func (r *SportsRepository) UpdateEventStatus(ctx context.Context, eventID, status string, homeScore, awayScore int, period string, minute int) error {
	_, err := r.db.Exec(ctx, `
		UPDATE sports_events SET status = $1, home_score = $2, away_score = $3, period = $4, minute = $5
		WHERE event_id = $6
	`, status, homeScore, awayScore, period, minute, eventID)
	return err
}

// SaveParlayBet saves a parlay bet
func (r *SportsRepository) SaveParlayBet(ctx context.Context, bet *model.ParlayBet) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
		INSERT INTO sports_parlay_bets (bet_id, user_id, stake, total_odds, potential_win, status, placed_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`, bet.BetID, bet.UserID, bet.Stake, bet.TotalOdds, bet.PotentialWin, bet.Status)
	if err != nil {
		return err
	}

	for _, sel := range bet.Selections {
		_, err = tx.Exec(ctx, `
			INSERT INTO sports_parlay_selections (bet_id, event_id, event_name, market_id, selection, odds, status)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, bet.BetID, sel.EventID, sel.EventName, sel.MarketID, sel.Selection, sel.Odds, sel.Status)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

// GetParlayBetByID returns a parlay bet by ID
func (r *SportsRepository) GetParlayBetByID(ctx context.Context, betID string) (*model.ParlayBet, error) {
	var b model.ParlayBet
	err := r.db.QueryRow(ctx, `
		SELECT bet_id, user_id, stake, total_odds, potential_win, status, placed_at
		FROM sports_parlay_bets WHERE bet_id = $1
	`, betID).Scan(&b.BetID, &b.UserID, &b.Stake, &b.TotalOdds, &b.PotentialWin, &b.Status, &b.PlacedAt)
	if err != nil {
		return nil, fmt.Errorf("parlay bet not found: %w", err)
	}

	rows, err := r.db.Query(ctx, `
		SELECT event_id, event_name, market_id, selection, odds, status
		FROM sports_parlay_selections WHERE bet_id = $1
	`, betID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sel model.ParlaySelection
		if err := rows.Scan(&sel.EventID, &sel.EventName, &sel.MarketID, &sel.Selection, &sel.Odds, &sel.Status); err != nil {
			return nil, err
		}
		b.Selections = append(b.Selections, sel)
	}

	return &b, nil
}

// GetUserParlayBets returns all parlay bets for a user
func (r *SportsRepository) GetUserParlayBets(ctx context.Context, userID string, limit, offset int) ([]model.ParlayBet, int, error) {
	var total int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM sports_parlay_bets WHERE user_id = $1`, userID).Scan(&total); err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(ctx, `
		SELECT bet_id, user_id, stake, total_odds, potential_win, status, placed_at
		FROM sports_parlay_bets WHERE user_id = $1 ORDER BY placed_at DESC LIMIT $2 OFFSET $3
	`, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var bets []model.ParlayBet
	for rows.Next() {
		var b model.ParlayBet
		if err := rows.Scan(&b.BetID, &b.UserID, &b.Stake, &b.TotalOdds, &b.PotentialWin, &b.Status, &b.PlacedAt); err != nil {
			return nil, 0, err
		}
		bets = append(bets, b)
	}
	return bets, total, nil
}
