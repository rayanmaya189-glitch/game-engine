package repository

import (
	"context"
	"fmt"

	"github.com/game_engine/live-dealer-service/internal/model"
)

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
