package service

import (
	"errors"
	"time"

	"github.com/game_engine/live-dealer-service/internal/model"
)

// Game State Management

func (s *DealerService) StartRound(tableID string) (*model.GameState, error) {
	table, ok := s.tables[tableID]
	if !ok {
		return nil, errors.New("table not found")
	}

	if table.Status != "open" {
		return nil, errors.New("table is not open")
	}

	for _, gs := range s.gameStates {
		if gs.TableID == tableID && gs.Phase != "finished" {
			return nil, errors.New("round already in progress")
		}
	}

	gameState := &model.GameState{
		TableID:   tableID,
		RoundID:   generateID(),
		Phase:     "betting",
		Cards:     []string{},
		Pot:       0,
		StartTime: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.gameStates[gameState.RoundID] = gameState
	return gameState, nil
}

func (s *DealerService) PlaceBet(playerID, roundID, betType string, amount float64) (*model.Bet, error) {
	player, ok := s.players[playerID]
	if !ok {
		return nil, errors.New("player not found")
	}

	gameState, ok := s.gameStates[roundID]
	if !ok {
		return nil, errors.New("round not found")
	}

	if gameState.Phase != "betting" {
		return nil, errors.New("betting phase is over")
	}

	table, ok := s.tables[gameState.TableID]
	if !ok {
		return nil, errors.New("table not found")
	}

	if amount < table.MinBet || amount > table.MaxBet {
		return nil, errors.New("bet amount outside table limits")
	}

	if amount > player.Chips {
		return nil, errors.New("insufficient chips")
	}

	bet := &model.Bet{
		BetID:     generateID(),
		PlayerID:  playerID,
		TableID:   player.TableID,
		RoundID:   roundID,
		BetType:   betType,
		BetAmount: amount,
		Odds:      1.0,
		Potential: amount,
		Result:    "pending",
		PlacedAt:  time.Now(),
	}

	s.bets[bet.BetID] = bet
	player.Chips -= amount
	player.CurrentBet += amount
	gameState.Pot += amount

	return bet, nil
}

func (s *DealerService) EndBetting(roundID string) error {
	gameState, ok := s.gameStates[roundID]
	if !ok {
		return errors.New("round not found")
	}

	if gameState.Phase != "betting" {
		return errors.New("not in betting phase")
	}

	gameState.Phase = "playing"
	gameState.UpdatedAt = time.Now()
	return nil
}

func (s *DealerService) ResolveRound(roundID string, winner string) error {
	gameState, ok := s.gameStates[roundID]
	if !ok {
		return errors.New("round not found")
	}

	if gameState.Phase == "finished" {
		return errors.New("round already finished")
	}

	gameState.Winner = winner
	gameState.Phase = "resolving"
	gameState.EndTime = time.Now()
	gameState.UpdatedAt = time.Now()

	var totalPayout float64
	for _, bet := range s.bets {
		if bet.RoundID != roundID {
			continue
		}

		if winner == "player" {
			bet.Result = "won"
			bet.Payout = bet.BetAmount * 2
			totalPayout += bet.Payout

			if player, ok := s.players[bet.PlayerID]; ok {
				player.Chips += bet.Payout
			}
		} else if winner == "push" {
			bet.Result = "void"
			bet.Payout = bet.BetAmount

			if player, ok := s.players[bet.PlayerID]; ok {
				player.Chips += bet.Payout
			}
		} else {
			bet.Result = "lost"
		}

		bet.ResultedAt = time.Now()
	}

	gameState.Payout = totalPayout
	gameState.Phase = "finished"

	return nil
}

func (s *DealerService) GetRound(roundID string) (*model.GameState, error) {
	gameState, ok := s.gameStates[roundID]
	if !ok {
		return nil, errors.New("round not found")
	}
	return gameState, nil
}

func (s *DealerService) GetPlayerBets(playerID string) []*model.Bet {
	var result []*model.Bet
	for _, bet := range s.bets {
		if bet.PlayerID == playerID {
			result = append(result, bet)
		}
	}
	return result
}
