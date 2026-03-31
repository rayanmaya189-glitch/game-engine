package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/game_engine/live-dealer-service/internal/config"
	"github.com/game_engine/live-dealer-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type LiveDealerRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewLiveDealerRepository(db *pgxpool.Pool, redis *redis.Client) *LiveDealerRepository {
	return &LiveDealerRepository{db: db, redis: redis}
}

func NewPostgresDB(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := cfg.ConnectionString()
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}
	if cfg.MaxConnections > 0 {
		poolConfig.MaxConns = int32(cfg.MaxConnections)
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

// --- Table persistence ---

func (r *LiveDealerRepository) CreateTable(ctx context.Context, table *model.Table) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO live_dealer_tables (table_id, game_type, dealer_id, dealer_name, status, min_bet, max_bet, current_seat, max_seats, stake_limit, deck_count, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`, table.TableID, table.GameType, table.DealerID, table.DealerName, table.Status,
		table.MinBet, table.MaxBet, table.CurrentSeat, table.MaxSeats, table.StakeLimit,
		table.DeckCount, table.CreatedAt, table.UpdatedAt)
	return err
}

func (r *LiveDealerRepository) GetTable(ctx context.Context, tableID string) (*model.Table, error) {
	var t model.Table
	err := r.db.QueryRow(ctx, `
		SELECT table_id, game_type, dealer_id, dealer_name, status, min_bet, max_bet, current_seat, max_seats, stake_limit, deck_count, created_at, updated_at
		FROM live_dealer_tables WHERE table_id = $1
	`, tableID).Scan(&t.TableID, &t.GameType, &t.DealerID, &t.DealerName, &t.Status,
		&t.MinBet, &t.MaxBet, &t.CurrentSeat, &t.MaxSeats, &t.StakeLimit,
		&t.DeckCount, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get table: %w", err)
	}
	return &t, nil
}

func (r *LiveDealerRepository) ListTables(ctx context.Context, gameType, status string) ([]*model.Table, error) {
	query := `SELECT table_id, game_type, dealer_id, dealer_name, status, min_bet, max_bet, current_seat, max_seats, stake_limit, deck_count, created_at, updated_at FROM live_dealer_tables WHERE 1=1`
	args := []interface{}{}
	argIdx := 1

	if gameType != "" {
		query += fmt.Sprintf(" AND game_type = $%d", argIdx)
		args = append(args, gameType)
		argIdx++
	}
	if status != "" {
		query += fmt.Sprintf(" AND status = $%d", argIdx)
		args = append(args, status)
		argIdx++
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list tables: %w", err)
	}
	defer rows.Close()

	var tables []*model.Table
	for rows.Next() {
		var t model.Table
		if err := rows.Scan(&t.TableID, &t.GameType, &t.DealerID, &t.DealerName, &t.Status,
			&t.MinBet, &t.MaxBet, &t.CurrentSeat, &t.MaxSeats, &t.StakeLimit,
			&t.DeckCount, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tables = append(tables, &t)
	}
	return tables, nil
}

func (r *LiveDealerRepository) UpdateTable(ctx context.Context, table *model.Table) error {
	_, err := r.db.Exec(ctx, `
		UPDATE live_dealer_tables SET game_type=$2, dealer_id=$3, dealer_name=$4, status=$5, min_bet=$6, max_bet=$7, current_seat=$8, max_seats=$9, stake_limit=$10, deck_count=$11, updated_at=$12
		WHERE table_id = $1
	`, table.TableID, table.GameType, table.DealerID, table.DealerName, table.Status,
		table.MinBet, table.MaxBet, table.CurrentSeat, table.MaxSeats, table.StakeLimit,
		table.DeckCount, time.Now())
	return err
}

// --- Player persistence ---

func (r *LiveDealerRepository) CreatePlayer(ctx context.Context, player *model.Player) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO live_dealer_players (player_id, table_id, seat_number, chips, current_bet, joined_at, last_action, hand_total, is_finished)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, player.PlayerID, player.TableID, player.SeatNumber, player.Chips, player.CurrentBet,
		player.JoinedAt, player.LastAction, player.HandTotal, player.IsFinished)
	return err
}

func (r *LiveDealerRepository) GetPlayersByTable(ctx context.Context, tableID string) ([]*model.Player, error) {
	rows, err := r.db.Query(ctx, `
		SELECT player_id, table_id, seat_number, chips, current_bet, joined_at, last_action, hand_total, is_finished
		FROM live_dealer_players WHERE table_id = $1
	`, tableID)
	if err != nil {
		return nil, fmt.Errorf("failed to get players: %w", err)
	}
	defer rows.Close()

	var players []*model.Player
	for rows.Next() {
		var p model.Player
		if err := rows.Scan(&p.PlayerID, &p.TableID, &p.SeatNumber, &p.Chips, &p.CurrentBet,
			&p.JoinedAt, &p.LastAction, &p.HandTotal, &p.IsFinished); err != nil {
			return nil, err
		}
		players = append(players, &p)
	}
	return players, nil
}

func (r *LiveDealerRepository) DeletePlayer(ctx context.Context, playerID string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM live_dealer_players WHERE player_id = $1`, playerID)
	return err
}

func (r *LiveDealerRepository) UpdatePlayer(ctx context.Context, player *model.Player) error {
	_, err := r.db.Exec(ctx, `
		UPDATE live_dealer_players SET chips=$2, current_bet=$3, last_action=$4, hand_total=$5, is_finished=$6
		WHERE player_id = $1
	`, player.PlayerID, player.Chips, player.CurrentBet, player.LastAction, player.HandTotal, player.IsFinished)
	return err
}

// --- Game state persistence ---

func (r *LiveDealerRepository) CreateGameState(ctx context.Context, gs *model.GameState) error {
	cardsJSON, _ := json.Marshal(gs.Cards)
	dealerCardsJSON, _ := json.Marshal(gs.DealerCards)
	_, err := r.db.Exec(ctx, `
		INSERT INTO live_dealer_game_states (table_id, round_id, phase, cards, dealer_cards, pot, dealer_total, winner, payout, start_time, end_time, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`, gs.TableID, gs.RoundID, gs.Phase, string(cardsJSON), string(dealerCardsJSON),
		gs.Pot, gs.DealerTotal, gs.Winner, gs.Payout, gs.StartTime, gs.EndTime, gs.UpdatedAt)
	return err
}

func (r *LiveDealerRepository) GetGameState(ctx context.Context, roundID string) (*model.GameState, error) {
	var gs model.GameState
	var cardsJSON, dealerCardsJSON string
	err := r.db.QueryRow(ctx, `
		SELECT table_id, round_id, phase, cards, dealer_cards, pot, dealer_total, winner, payout, start_time, end_time, updated_at
		FROM live_dealer_game_states WHERE round_id = $1
	`, roundID).Scan(&gs.TableID, &gs.RoundID, &gs.Phase, &cardsJSON, &dealerCardsJSON,
		&gs.Pot, &gs.DealerTotal, &gs.Winner, &gs.Payout, &gs.StartTime, &gs.EndTime, &gs.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get game state: %w", err)
	}
	json.Unmarshal([]byte(cardsJSON), &gs.Cards)
	json.Unmarshal([]byte(dealerCardsJSON), &gs.DealerCards)
	return &gs, nil
}

func (r *LiveDealerRepository) UpdateGameState(ctx context.Context, gs *model.GameState) error {
	cardsJSON, _ := json.Marshal(gs.Cards)
	dealerCardsJSON, _ := json.Marshal(gs.DealerCards)
	_, err := r.db.Exec(ctx, `
		UPDATE live_dealer_game_states SET phase=$3, cards=$4, dealer_cards=$5, pot=$6, dealer_total=$7, winner=$8, payout=$9, end_time=$10, updated_at=$11
		WHERE round_id = $1 AND table_id = $2
	`, gs.RoundID, gs.TableID, gs.Phase, string(cardsJSON), string(dealerCardsJSON),
		gs.Pot, gs.DealerTotal, gs.Winner, gs.Payout, gs.EndTime, time.Now())
	return err
}

func (r *LiveDealerRepository) GetActiveRound(ctx context.Context, tableID string) (*model.GameState, error) {
	var gs model.GameState
	var cardsJSON, dealerCardsJSON string
	err := r.db.QueryRow(ctx, `
		SELECT table_id, round_id, phase, cards, dealer_cards, pot, dealer_total, winner, payout, start_time, end_time, updated_at
		FROM live_dealer_game_states WHERE table_id = $1 AND phase != 'finished'
	`, tableID).Scan(&gs.TableID, &gs.RoundID, &gs.Phase, &cardsJSON, &dealerCardsJSON,
		&gs.Pot, &gs.DealerTotal, &gs.Winner, &gs.Payout, &gs.StartTime, &gs.EndTime, &gs.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("no active round: %w", err)
	}
	json.Unmarshal([]byte(cardsJSON), &gs.Cards)
	json.Unmarshal([]byte(dealerCardsJSON), &gs.DealerCards)
	return &gs, nil
}

// --- Bet persistence ---

func (r *LiveDealerRepository) CreateBet(ctx context.Context, bet *model.Bet) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO live_dealer_bets (bet_id, player_id, table_id, round_id, bet_type, bet_amount, odds, potential, result, payout, placed_at, resulted_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`, bet.BetID, bet.PlayerID, bet.TableID, bet.RoundID, bet.BetType, bet.BetAmount,
		bet.Odds, bet.Potential, bet.Result, bet.Payout, bet.PlacedAt, bet.ResultedAt)
	return err
}

func (r *LiveDealerRepository) GetBetsByRound(ctx context.Context, roundID string) ([]*model.Bet, error) {
	rows, err := r.db.Query(ctx, `
		SELECT bet_id, player_id, table_id, round_id, bet_type, bet_amount, odds, potential, result, payout, placed_at, resulted_at
		FROM live_dealer_bets WHERE round_id = $1
	`, roundID)
	if err != nil {
		return nil, fmt.Errorf("failed to get bets: %w", err)
	}
	defer rows.Close()

	var bets []*model.Bet
	for rows.Next() {
		var b model.Bet
		if err := rows.Scan(&b.BetID, &b.PlayerID, &b.TableID, &b.RoundID, &b.BetType, &b.BetAmount,
			&b.Odds, &b.Potential, &b.Result, &b.Payout, &b.PlacedAt, &b.ResultedAt); err != nil {
			return nil, err
		}
		bets = append(bets, &b)
	}
	return bets, nil
}

func (r *LiveDealerRepository) UpdateBet(ctx context.Context, bet *model.Bet) error {
	_, err := r.db.Exec(ctx, `
		UPDATE live_dealer_bets SET result=$3, payout=$4, resulted_at=$5
		WHERE bet_id = $1 AND round_id = $2
	`, bet.BetID, bet.RoundID, bet.Result, bet.Payout, bet.ResultedAt)
	return err
}

// --- Redis session state ---

func (r *LiveDealerRepository) CacheSessionState(ctx context.Context, tableID string, state interface{}) error {
	key := fmt.Sprintf("live_dealer:session:%s", tableID)
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	return r.redis.Set(ctx, key, data, time.Duration(60)*time.Minute).Err()
}

func (r *LiveDealerRepository) GetCachedSessionState(ctx context.Context, tableID string, dest interface{}) error {
	key := fmt.Sprintf("live_dealer:session:%s", tableID)
	data, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

func (r *LiveDealerRepository) DeleteSessionState(ctx context.Context, tableID string) error {
	key := fmt.Sprintf("live_dealer:session:%s", tableID)
	return r.redis.Del(ctx, key).Err()
}

func (r *LiveDealerRepository) SetPlayerOnline(ctx context.Context, tableID, playerID string) error {
	key := fmt.Sprintf("live_dealer:online:%s", tableID)
	return r.redis.SAdd(ctx, key, playerID).Err()
}

func (r *LiveDealerRepository) RemovePlayerOnline(ctx context.Context, tableID, playerID string) error {
	key := fmt.Sprintf("live_dealer:online:%s", tableID)
	return r.redis.SRem(ctx, key, playerID).Err()
}

func (r *LiveDealerRepository) GetOnlinePlayers(ctx context.Context, tableID string) ([]string, error) {
	key := fmt.Sprintf("live_dealer:online:%s", tableID)
	return r.redis.SMembers(ctx, key).Result()
}

// --- Dealer persistence ---

func (r *LiveDealerRepository) CreateDealer(ctx context.Context, dealer *model.Dealer) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO live_dealer_dealers (dealer_id, name, avatar, language, status, table_id, shift_start, shift_end)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, dealer.DealerID, dealer.Name, dealer.Avatar, dealer.Language, dealer.Status,
		dealer.TableID, dealer.ShiftStart, dealer.ShiftEnd)
	return err
}

func (r *LiveDealerRepository) GetDealer(ctx context.Context, dealerID string) (*model.Dealer, error) {
	var d model.Dealer
	err := r.db.QueryRow(ctx, `
		SELECT dealer_id, name, avatar, language, status, table_id, shift_start, shift_end
		FROM live_dealer_dealers WHERE dealer_id = $1
	`, dealerID).Scan(&d.DealerID, &d.Name, &d.Avatar, &d.Language, &d.Status,
		&d.TableID, &d.ShiftStart, &d.ShiftEnd)
	if err != nil {
		return nil, fmt.Errorf("failed to get dealer: %w", err)
	}
	return &d, nil
}

func (r *LiveDealerRepository) UpdateDealer(ctx context.Context, dealer *model.Dealer) error {
	_, err := r.db.Exec(ctx, `
		UPDATE live_dealer_dealers SET name=$2, avatar=$3, language=$4, status=$5, table_id=$6, shift_start=$7, shift_end=$8
		WHERE dealer_id = $1
	`, dealer.DealerID, dealer.Name, dealer.Avatar, dealer.Language, dealer.Status,
		dealer.TableID, dealer.ShiftStart, dealer.ShiftEnd)
	return err
}

func (r *LiveDealerRepository) ListDealers(ctx context.Context, status string) ([]*model.Dealer, error) {
	query := `SELECT dealer_id, name, avatar, language, status, table_id, shift_start, shift_end FROM live_dealer_dealers WHERE 1=1`
	args := []interface{}{}
	if status != "" {
		query += " AND status = $1"
		args = append(args, status)
	}
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list dealers: %w", err)
	}
	defer rows.Close()

	var dealers []*model.Dealer
	for rows.Next() {
		var d model.Dealer
		if err := rows.Scan(&d.DealerID, &d.Name, &d.Avatar, &d.Language, &d.Status,
			&d.TableID, &d.ShiftStart, &d.ShiftEnd); err != nil {
			return nil, err
		}
		dealers = append(dealers, &d)
	}
	return dealers, nil
}
