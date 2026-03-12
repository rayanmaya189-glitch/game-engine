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
