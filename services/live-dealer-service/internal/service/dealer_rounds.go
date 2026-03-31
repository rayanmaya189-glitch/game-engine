package service

import (
	"context"
	"fmt"
	"time"

	"github.com/game_engine/live-dealer-service/internal/model"
)

// --- Game Round Management ---

func (s *LiveDealerService) StartRound(ctx context.Context, tableID string) (*model.GameState, error) {
	table, err := s.repo.GetTable(ctx, tableID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if table.Status != "open" {
		return nil, fmt.Errorf("session is not open")
	}

	existing, _ := s.repo.GetActiveRound(ctx, tableID)
	if existing != nil {
		return nil, fmt.Errorf("round already in progress")
	}

	gameState := &model.GameState{
		TableID:     tableID,
		RoundID:     generateID(),
		Phase:       "betting",
		Cards:       []string{},
		DealerCards: []string{},
		Pot:         0,
		StartTime:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateGameState(ctx, gameState); err != nil {
		return nil, fmt.Errorf("failed to create round: %w", err)
	}

	return gameState, nil
}

func (s *LiveDealerService) PlaceBet(ctx context.Context, playerID, roundID, betType string, amount float64) (*model.Bet, error) {
	gameState, err := s.repo.GetGameState(ctx, roundID)
	if err != nil {
		return nil, fmt.Errorf("round not found: %w", err)
	}

	if gameState.Phase != "betting" {
		return nil, fmt.Errorf("betting phase is over")
	}

	table, err := s.repo.GetTable(ctx, gameState.TableID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if amount < table.MinBet || amount > table.MaxBet {
		return nil, fmt.Errorf("bet amount outside table limits")
	}

	players, err := s.repo.GetPlayersByTable(ctx, gameState.TableID)
	if err != nil {
		return nil, fmt.Errorf("failed to get players: %w", err)
	}

	var player *model.Player
	for _, p := range players {
		if p.PlayerID == playerID {
			player = p
			break
		}
	}
	if player == nil {
		return nil, fmt.Errorf("player not in session")
	}

	if amount > player.Chips {
		return nil, fmt.Errorf("insufficient chips")
	}

	bet := &model.Bet{
		BetID:     generateID(),
		PlayerID:  playerID,
		TableID:   gameState.TableID,
		RoundID:   roundID,
		BetType:   betType,
		BetAmount: amount,
		Odds:      1.0,
		Potential: amount,
		Result:    "pending",
		PlacedAt:  time.Now(),
	}

	if err := s.repo.CreateBet(ctx, bet); err != nil {
		return nil, fmt.Errorf("failed to place bet: %w", err)
	}

	player.Chips -= amount
	player.CurrentBet += amount
	if err := s.repo.UpdatePlayer(ctx, player); err != nil {
		return nil, fmt.Errorf("failed to update player: %w", err)
	}

	gameState.Pot += amount
	if err := s.repo.UpdateGameState(ctx, gameState); err != nil {
		return nil, fmt.Errorf("failed to update game state: %w", err)
	}

	return bet, nil
}

func (s *LiveDealerService) EndBetting(ctx context.Context, roundID string) error {
	gameState, err := s.repo.GetGameState(ctx, roundID)
	if err != nil {
		return fmt.Errorf("round not found: %w", err)
	}

	if gameState.Phase != "betting" {
		return fmt.Errorf("not in betting phase")
	}

	gameState.Phase = "playing"
	gameState.UpdatedAt = time.Now()

	return s.repo.UpdateGameState(ctx, gameState)
}

func (s *LiveDealerService) ResolveRound(ctx context.Context, roundID, winner string) error {
	gameState, err := s.repo.GetGameState(ctx, roundID)
	if err != nil {
		return fmt.Errorf("round not found: %w", err)
	}

	if gameState.Phase == "finished" {
		return fmt.Errorf("round already finished")
	}

	gameState.Winner = winner
	gameState.Phase = "resolving"
	gameState.EndTime = time.Now()
	gameState.UpdatedAt = time.Now()

	bets, err := s.repo.GetBetsByRound(ctx, roundID)
	if err != nil {
		return fmt.Errorf("failed to get bets: %w", err)
	}

	var totalPayout float64
	for _, bet := range bets {
		switch winner {
		case "player":
			bet.Result = "won"
			bet.Payout = bet.BetAmount * 2
			totalPayout += bet.Payout
		case "push":
			bet.Result = "void"
			bet.Payout = bet.BetAmount
		default:
			bet.Result = "lost"
			bet.Payout = 0
		}
		bet.ResultedAt = time.Now()

		if err := s.repo.UpdateBet(ctx, bet); err != nil {
			return fmt.Errorf("failed to update bet: %w", err)
		}

		if winner == "player" || winner == "push" {
			players, _ := s.repo.GetPlayersByTable(ctx, gameState.TableID)
			for _, p := range players {
				if p.PlayerID == bet.PlayerID {
					p.Chips += bet.Payout
					s.repo.UpdatePlayer(ctx, p)
					break
				}
			}
		}
	}

	gameState.Payout = totalPayout
	gameState.Phase = "finished"

	return s.repo.UpdateGameState(ctx, gameState)
}

func (s *LiveDealerService) GetRound(ctx context.Context, roundID string) (*model.GameState, error) {
	return s.repo.GetGameState(ctx, roundID)
}
