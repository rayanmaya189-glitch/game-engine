package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/game_engine/live-dealer-service/internal/model"
)

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
