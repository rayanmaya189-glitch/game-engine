package repository

import (
	"context"
	"fmt"

	"github.com/game_engine/sports-betting-service/internal/model"
)

func (r *SportsRepository) CacheOdds(ctx context.Context, eventID string, markets []model.Market) error {
	key := fmt.Sprintf("sports:odds:%s", eventID)
	for _, m := range markets {
		marketKey := fmt.Sprintf("%s:%s:%f:%f:%f", m.MarketID, m.Name, m.HomeOdds, m.DrawOdds, m.AwayOdds)
		if err := r.redis.SAdd(ctx, key, marketKey).Err(); err != nil {
			return err
		}
	}
	return r.redis.Expire(ctx, key, 300).Err()
}

// GetLiveMarkets returns markets for a live event with full selections
func (r *SportsRepository) GetLiveMarkets(ctx context.Context, eventID string) ([]model.Market, error) {
	rows, err := r.db.Query(ctx, `
		SELECT market_id, event_id, name, market_type, status
		FROM sports_markets WHERE event_id = $1 AND status IN ('open', 'suspended')
	`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markets []model.Market
	for rows.Next() {
		var m model.Market
		if err := rows.Scan(&m.MarketID, &m.EventID, &m.Name, &m.MarketType, &m.Status); err != nil {
			return nil, err
		}
		selRows, err := r.db.Query(ctx, `
			SELECT selection_id, selection, odds FROM sports_market_selections WHERE market_id = $1
		`, m.MarketID)
		if err != nil {
			return nil, err
		}
		for selRows.Next() {
			var sel model.Selection
			if err := selRows.Scan(&sel.SelectionID, &sel.Selection, &sel.Odds); err != nil {
				selRows.Close()
				return nil, err
			}
			m.Selections = append(m.Selections, sel)
		}
		selRows.Close()
		markets = append(markets, m)
	}
	return markets, nil
}

// UpdateLiveEvent updates a live event with new odds
func (r *SportsRepository) UpdateLiveEvent(ctx context.Context, eventID string, markets []model.Market) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for _, m := range markets {
		_, err := tx.Exec(ctx, `
			UPDATE sports_markets SET status = $1 WHERE market_id = $2 AND event_id = $3
		`, m.Status, m.MarketID, eventID)
		if err != nil {
			return err
		}

		for _, sel := range m.Selections {
			_, err := tx.Exec(ctx, `
				UPDATE sports_market_selections SET odds = $1 WHERE selection_id = $2
			`, sel.Odds, sel.SelectionID)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// SaveOddsChanges saves odds change history
func (r *SportsRepository) SaveOddsChanges(ctx context.Context, eventID string, changes []model.OddsChange) error {
	for _, c := range changes {
		_, err := r.db.Exec(ctx, `
			INSERT INTO sports_odds_history (event_id, market_id, selection, old_odds, new_odds, changed_at)
			VALUES ($1, $2, $3, $4, $5, NOW())
		`, eventID, c.MarketID, c.Selection, c.OldOdds, c.NewOdds)
		if err != nil {
			return err
		}
	}
	return nil
}
