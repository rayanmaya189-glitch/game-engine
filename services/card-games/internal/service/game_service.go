package service

import (
	"sync"

	"github.com/game_engine/card-games/internal/config"
	"github.com/game_engine/card-games/internal/games/andar_bahar"
	"github.com/game_engine/card-games/internal/games/baccarat"
	"github.com/game_engine/card-games/internal/games/blackjack"
	"github.com/game_engine/card-games/internal/games/common"
	"github.com/game_engine/card-games/internal/games/poker"
	"github.com/game_engine/card-games/internal/games/teen_patti"
)

// GameType represents the type of card game
type GameType string

const (
	GameTypeBlackjack  GameType = "blackjack"
	GameTypeBaccarat   GameType = "baccarat"
	GameTypePoker      GameType = "poker"
	GameTypeAndarBahar GameType = "andar_bahar"
	GameTypeTeenPatti  GameType = "teen_patti"
)

// GameService manages card games
type GameService struct {
	mu    sync.RWMutex
	games map[string]interface{}
	cfg   *config.Config
}

// NewGameService creates a new game service
func NewGameService(cfg *config.Config) (*GameService, error) {
	return &GameService{
		games: make(map[string]interface{}),
		cfg:   cfg,
	}, nil
}

// CreateBlackjack creates a new blackjack game
func (s *GameService) CreateBlackjack(gameID, playerID string, bet int64) (*blackjack.Game, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Create new shoe if needed
	shoe := common.NewShoe(s.cfg.Game.DefaultDeckCount)
	shuffler := common.NewFisherYatesShuffler(&simpleRNG{})
	shuffler.Shuffle(shoe.Cards)

	game := blackjack.NewGame(gameID, shoe, &blackjack.Config{
		AllowSurrender:       s.cfg.Game.Blackjack.AllowSurrender,
		AllowLateSurrender:   s.cfg.Game.Blackjack.AllowLateSurrender,
		DealerStandsOnSoft17: s.cfg.Game.Blackjack.DealerStandsOnSoft17,
		MaxSplits:            s.cfg.Game.Blackjack.MaxSplits,
		BlackjackPayout:      s.cfg.Game.Blackjack.BlackjackPayout,
	})

	if err := game.AddPlayer(playerID, bet); err != nil {
		return nil, err
	}

	if err := game.Start(); err != nil {
		return nil, err
	}

	s.games[gameID] = game
	return game, nil
}

// CreateBaccarat creates a new baccarat game
func (s *GameService) CreateBaccarat(gameID string, bets map[baccarat.BetType]int64) (*baccarat.Game, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	shoe := common.NewShoe(s.cfg.Game.Baccarat.ShoeSize)
	shuffler := common.NewFisherYatesShuffler(&simpleRNG{})
	shuffler.Shuffle(shoe.Cards)

	game := baccarat.NewGame(gameID, shoe, &baccarat.Config{
		Commission: s.cfg.Game.Baccarat.Commission,
		MaxPlayers: s.cfg.Game.Baccarat.MaxPlayers,
		ShoeSize:   s.cfg.Game.Baccarat.ShoeSize,
	})

	if err := game.DealInitialCards(); err != nil {
		return nil, err
	}

	game.Evaluate()
	game.SettleBets(bets)

	s.games[gameID] = game
	return game, nil
}

// CreatePoker creates a new poker game
func (s *GameService) CreatePoker(gameID string, gameType poker.GameType) (*poker.Game, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	shoe := common.NewShoe(1)
	shuffler := common.NewFisherYatesShuffler(&simpleRNG{})
	shuffler.Shuffle(shoe.Cards)

	game := poker.NewGame(gameID, gameType, shoe, &poker.Config{
		MinPlayers:    s.cfg.Game.Poker.MinPlayers,
		MaxPlayers:    s.cfg.Game.Poker.MaxPlayers,
		SmallBlind:    10,
		BigBlind:      20,
		StartingChips: 1000,
	})

	s.games[gameID] = game
	return game, nil
}

// CreateAndarBahar creates a new Andar Bahar game
func (s *GameService) CreateAndarBahar(gameID string) (*andar_bahar.Game, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	shoe := common.NewShoe(1)
	shuffler := common.NewFisherYatesShuffler(&simpleRNG{})
	shuffler.Shuffle(shoe.Cards)

	game := andar_bahar.NewGame(gameID, shoe, &andar_bahar.Config{
		SideToDealFirst:   s.cfg.Game.AndarBahar.SideToDealFirst,
		MaxCommunityCards: s.cfg.Game.AndarBahar.MaxCommunityCards,
	})

	// Draw joker
	joker, err := shoe.Draw()
	if err != nil {
		return nil, err
	}
	game.SetJoker(joker)

	s.games[gameID] = game
	return game, nil
}

// CreateTeenPatti creates a new Teen Patti game
func (s *GameService) CreateTeenPatti(gameID string) (*teen_patti.Game, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	shoe := common.NewShoe(1)
	shuffler := common.NewFisherYatesShuffler(&simpleRNG{})
	shuffler.Shuffle(shoe.Cards)

	game := teen_patti.NewGame(gameID, shoe, &teen_patti.Config{
		Variant:    s.cfg.Game.TeenPatti.Variant,
		MinPlayers: s.cfg.Game.TeenPatti.MinPlayers,
		MaxPlayers: s.cfg.Game.TeenPatti.MaxPlayers,
	})

	s.games[gameID] = game
	return game, nil
}

// GetGame retrieves a game by ID
func (s *GameService) GetGame(gameID string) (interface{}, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	game, exists := s.games[gameID]
	if !exists {
		return nil, nil
	}

	return game, nil
}

// RemoveGame removes a game
func (s *GameService) RemoveGame(gameID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.games, gameID)
}

// simpleRNG is a simple RNG implementation for shuffling
type simpleRNG struct{}

func (r *simpleRNG) Intn(n int) (int, error) {
	// Simple deterministic RNG for now
	// In production, this would use the RNG service
	return 0, nil
}
