package repository

import (
	"context"
	"fmt"

	"github.com/game_engine/sports-betting-service/internal/config"
	"github.com/game_engine/sports-betting-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type SportsRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewSportsRepository(db *pgxpool.Pool, redis *redis.Client) *SportsRepository {
	return &SportsRepository{db: db, redis: redis}
}

func NewPostgresDB(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := cfg.ConnectionString()
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return pool, nil
}

func NewRedisClient(cfg config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return client, nil
}

func (r *SportsRepository) GetSports(ctx context.Context) ([]model.Sport, error) {
	rows, err := r.db.Query(ctx, `
		SELECT sport_id, name, icon, status, sort_order FROM sports WHERE status = 'active' ORDER BY sort_order
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sports []model.Sport
	for rows.Next() {
		var s model.Sport
		if err := rows.Scan(&s.SportID, &s.Name, &s.Icon, &s.Status, &s.SortOrder); err != nil {
			return nil, err
		}
		sports = append(sports, s)
	}
	return sports, nil
}

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

func (r *SportsRepository) GetMarkets(ctx context.Context, eventID string) ([]model.Market, error) {
	rows, err := r.db.Query(ctx, `
		SELECT market_id, event_id, name, status, home_odds, draw_odds, away_odds
		FROM sports_markets WHERE event_id = $1 AND status IN ('open', 'closed')
	`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var markets []model.Market
	for rows.Next() {
		var m model.Market
		if err := rows.Scan(&m.MarketID, &m.EventID, &m.Name, &m.Status, &m.HomeOdds, &m.DrawOdds, &m.AwayOdds); err != nil {
			return nil, err
		}
		markets = append(markets, m)
	}
	return markets, nil
}

func (r *SportsRepository) PlaceBet(ctx context.Context, bet *model.Bet) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO sports_bets (bet_id, user_id, event_id, market_id, selection, stake, odds, potential_win, status, placed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
	`, bet.BetID, bet.UserID, bet.EventID, bet.MarketID, bet.Selection, bet.Stake, bet.Odds, bet.PotentialWin, bet.Status)
	return err
}

func (r *SportsRepository) GetUserBets(ctx context.Context, userID string, limit, offset int) ([]model.Bet, int, error) {
	var total int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM sports_bets WHERE user_id = $1`, userID).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count bets: %w", err)
	}

	rows, err := r.db.Query(ctx, `
		SELECT bet_id, user_id, event_id, market_id, selection, stake, odds, potential_win, status, placed_at, settled_at
		FROM sports_bets WHERE user_id = $1 ORDER BY placed_at DESC LIMIT $2 OFFSET $3
	`, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var bets []model.Bet
	for rows.Next() {
		var b model.Bet
		if err := rows.Scan(&b.BetID, &b.UserID, &b.EventID, &b.MarketID, &b.Selection, &b.Stake, &b.Odds, &b.PotentialWin, &b.Status, &b.PlacedAt, &b.SettledAt); err != nil {
			return nil, 0, err
		}
		bets = append(bets, b)
	}
	return bets, total, nil
}

func (r *SportsRepository) SettleBet(ctx context.Context, betID string, status string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE sports_bets SET status = $1, settled_at = NOW() WHERE bet_id = $2
	`, status, betID)
	return err
}

func (r *SportsRepository) BeginTx(ctx context.Context) (pgxpool.Tx, error) {
	return r.db.Begin(ctx)
}

func (r *SportsRepository) PlaceBetTx(ctx context.Context, tx *pgxpool.Tx, bet *model.Bet) error {
	_, err := tx.Exec(ctx, `
		INSERT INTO sports_bets (bet_id, user_id, event_id, market_id, selection, stake, odds, potential_win, status, placed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
	`, bet.BetID, bet.UserID, bet.EventID, bet.MarketID, bet.Selection, bet.Stake, bet.Odds, bet.PotentialWin, bet.Status)
	return err
}

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
		// Get selections for this market
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
	// Update markets in a transaction
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

		// Update selections
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

// UpdateEventStatus updates event status, score, period, and minute
func (r *SportsRepository) UpdateEventStatus(ctx context.Context, eventID, status string, homeScore, awayScore int, period string, minute int) error {
	_, err := r.db.Exec(ctx, `
		UPDATE sports_events SET status = $1, home_score = $2, away_score = $3, period = $4, minute = $5
		WHERE event_id = $6
	`, status, homeScore, awayScore, period, minute, eventID)
	return err
}

// GetPendingBetsForEvent returns all pending bets for an event
func (r *SportsRepository) GetPendingBetsForEvent(ctx context.Context, eventID string) ([]model.Bet, error) {
	rows, err := r.db.Query(ctx, `
		SELECT bet_id, user_id, event_id, market_id, selection, stake, odds, potential_win, status, placed_at
		FROM sports_bets WHERE event_id = $1 AND status = 'pending'
	`, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bets []model.Bet
	for rows.Next() {
		var b model.Bet
		if err := rows.Scan(&b.BetID, &b.UserID, &b.EventID, &b.MarketID, &b.Selection, &b.Stake, &b.Odds, &b.PotentialWin, &b.Status, &b.PlacedAt); err != nil {
			return nil, err
		}
		bets = append(bets, b)
	}
	return bets, nil
}

// UpdateBetStatus updates bet status and settled amount
func (r *SportsRepository) UpdateBetStatus(ctx context.Context, betID string, status model.BetStatus, winAmount float64) error {
	if status == model.BetStatusWon || status == model.BetStatusCashOut {
		_, err := r.db.Exec(ctx, `
			UPDATE sports_bets SET status = $1, settled_at = NOW() WHERE bet_id = $2
		`, status, betID)
		return err
	}
	_, err := r.db.Exec(ctx, `
		UPDATE sports_bets SET status = $1, settled_at = NOW() WHERE bet_id = $2
	`, status, betID)
	return err
}

// GetBetByID returns a bet by ID
func (r *SportsRepository) GetBetByID(ctx context.Context, betID string) (*model.Bet, error) {
	var b model.Bet
	err := r.db.QueryRow(ctx, `
		SELECT bet_id, user_id, event_id, market_id, selection, stake, odds, potential_win, status, placed_at, settled_at
		FROM sports_bets WHERE bet_id = $1
	`, betID).Scan(&b.BetID, &b.UserID, &b.EventID, &b.MarketID, &b.Selection, &b.Stake, &b.Odds, &b.PotentialWin, &b.Status, &b.PlacedAt, &b.SettledAt)
	if err != nil {
		return nil, fmt.Errorf("bet not found: %w", err)
	}
	return &b, nil
}

// SaveParlayBet saves a parlay bet
func (r *SportsRepository) SaveParlayBet(ctx context.Context, bet *model.ParlayBet) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Save main bet record
	_, err = tx.Exec(ctx, `
		INSERT INTO sports_parlay_bets (bet_id, user_id, stake, total_odds, potential_win, status, placed_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`, bet.BetID, bet.UserID, bet.Stake, bet.TotalOdds, bet.PotentialWin, bet.Status)
	if err != nil {
		return err
	}

	// Save selections
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

	// Get selections
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

// SaveCashOut saves a cash out record
func (r *SportsRepository) SaveCashOut(ctx context.Context, cashOut *model.CashOut) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO sports_cashouts (cash_out_id, bet_id, user_id, original_stake, original_odds, current_odds, cash_out_amount, profit, status, requested_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
	`, cashOut.CashOutID, cashOut.BetID, cashOut.UserID, cashOut.OriginalStake, cashOut.OriginalOdds, cashOut.CurrentOdds, cashOut.CashOutAmount, cashOut.Profit, cashOut.Status)
	return err
}

// UpdateCashOutStatus updates cash out status
func (r *SportsRepository) UpdateCashOutStatus(ctx context.Context, cashOutID, status string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE sports_cashouts SET status = $1, completed_at = NOW() WHERE cash_out_id = $2
	`, status, cashOutID)
	return err
}

// GetUserCashOuts returns all cash outs for a user
func (r *SportsRepository) GetUserCashOuts(ctx context.Context, userID string, limit, offset int) ([]model.CashOut, int, error) {
	var total int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM sports_cashouts WHERE user_id = $1`, userID).Scan(&total); err != nil {
		return nil, 0, err
	}

	rows, err := r.db.Query(ctx, `
		SELECT cash_out_id, bet_id, user_id, original_stake, original_odds, current_odds, cash_out_amount, profit, status, requested_at, completed_at
		FROM sports_cashouts WHERE user_id = $1 ORDER BY requested_at DESC LIMIT $2 OFFSET $3
	`, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var cashOuts []model.CashOut
	for rows.Next() {
		var c model.CashOut
		if err := rows.Scan(&c.CashOutID, &c.BetID, &c.UserID, &c.OriginalStake, &c.OriginalOdds, &c.CurrentOdds, &c.CashOutAmount, &c.Profit, &c.Status, &c.RequestedAt, &c.CompletedAt); err != nil {
			return nil, 0, err
		}
		cashOuts = append(cashOuts, c)
	}
	return cashOuts, total, nil
}
