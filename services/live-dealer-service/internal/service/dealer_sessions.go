package service

import (
	"context"
	"fmt"
	"time"

	"github.com/game_engine/live-dealer-service/internal/model"
)

// --- Player Management ---

func (s *LiveDealerService) JoinSession(ctx context.Context, tableID, playerID string, chips float64) (*model.Player, error) {
	table, err := s.repo.GetTable(ctx, tableID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if table.Status != "open" {
		return nil, fmt.Errorf("session is not open")
	}

	if table.CurrentSeat >= table.MaxSeats {
		return nil, fmt.Errorf("session is full")
	}

	existingPlayers, err := s.repo.GetPlayersByTable(ctx, tableID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing players: %w", err)
	}
	for _, p := range existingPlayers {
		if p.PlayerID == playerID {
			return nil, fmt.Errorf("player already in session")
		}
	}

	player := &model.Player{
		PlayerID:   playerID,
		TableID:    tableID,
		SeatNumber: table.CurrentSeat + 1,
		Chips:      chips,
		JoinedAt:   time.Now(),
		IsFinished: false,
	}

	if err := s.repo.CreatePlayer(ctx, player); err != nil {
		return nil, fmt.Errorf("failed to add player: %w", err)
	}

	table.CurrentSeat++
	table.UpdatedAt = time.Now()
	if err := s.repo.UpdateTable(ctx, table); err != nil {
		return nil, fmt.Errorf("failed to update table: %w", err)
	}

	if err := s.repo.SetPlayerOnline(ctx, tableID, playerID); err != nil {
		return nil, fmt.Errorf("failed to mark player online: %w", err)
	}

	if err := s.repo.CacheSessionState(ctx, tableID, table); err != nil {
		return nil, fmt.Errorf("failed to update session cache: %w", err)
	}

	return player, nil
}

func (s *LiveDealerService) LeaveSession(ctx context.Context, tableID, playerID string) error {
	table, err := s.repo.GetTable(ctx, tableID)
	if err != nil {
		return fmt.Errorf("session not found: %w", err)
	}

	if err := s.repo.DeletePlayer(ctx, playerID); err != nil {
		return fmt.Errorf("failed to remove player: %w", err)
	}

	if table.CurrentSeat > 0 {
		table.CurrentSeat--
		table.UpdatedAt = time.Now()
		if err := s.repo.UpdateTable(ctx, table); err != nil {
			return fmt.Errorf("failed to update table: %w", err)
		}
	}

	if err := s.repo.RemovePlayerOnline(ctx, tableID, playerID); err != nil {
		return fmt.Errorf("failed to mark player offline: %w", err)
	}

	if err := s.repo.CacheSessionState(ctx, tableID, table); err != nil {
		return fmt.Errorf("failed to update session cache: %w", err)
	}

	return nil
}

func (s *LiveDealerService) GetSessionPlayers(ctx context.Context, tableID string) ([]*model.Player, error) {
	return s.repo.GetPlayersByTable(ctx, tableID)
}
