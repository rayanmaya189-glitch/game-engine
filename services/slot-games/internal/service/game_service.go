package service

import (
	"errors"

	"github.com/game-engine/slot-games/internal/games"
)

// GameService manages slot games
type GameService struct {
	games map[string]*games.Game
}

// NewGameService creates a new game service
func NewGameService() (*GameService, error) {
	return &GameService{
		games: make(map[string]*games.Game),
	}, nil
}

// CreateGame creates a new slot game
func (s *GameService) CreateGame(gameType, id string) (*games.Game, error) {
	var game *games.Game

	switch gameType {
	case "classic":
		game = games.NewClassicSlotGame(id)
	case "video":
		game = games.NewVideoSlotGame(id)
	default:
		return nil, errors.New("invalid game type")
	}

	s.games[id] = game
	return game, nil
}

// GetGame gets a game by ID
func (s *GameService) GetGame(id string) (*games.Game, error) {
	game, ok := s.games[id]
	if !ok {
		return nil, errors.New("game not found")
	}
	return game, nil
}

// Spin performs a spin on a game
func (s *GameService) Spin(id string, lineBet int64, lines int) (*games.GameState, error) {
	game, ok := s.games[id]
	if !ok {
		return nil, errors.New("game not found")
	}

	if err := game.SetBet(lineBet, lines); err != nil {
		return nil, err
	}

	if err := game.Spin(); err != nil {
		return nil, err
	}

	return game.GetState(), nil
}

// ProvablyFairSpin performs a provably fair spin
func (s *GameService) ProvablyFairSpin(id, serverSeed, clientSeed string, nonce int, lineBet int64, lines int) (*games.GameState, error) {
	game, ok := s.games[id]
	if !ok {
		return nil, errors.New("game not found")
	}

	if err := game.SetBet(lineBet, lines); err != nil {
		return nil, err
	}

	if err := game.ProvablyFairSpin(serverSeed, clientSeed, nonce); err != nil {
		return nil, err
	}

	return game.GetState(), nil
}

// DeleteGame deletes a game
func (s *GameService) DeleteGame(id string) error {
	if _, ok := s.games[id]; !ok {
		return errors.New("game not found")
	}
	delete(s.games, id)
	return nil
}

// ListGames lists all active games
func (s *GameService) ListGames() []string {
	ids := make([]string, 0, len(s.games))
	for id := range s.games {
		ids = append(ids, id)
	}
	return ids
}
