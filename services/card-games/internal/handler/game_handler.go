package handler

import (
	"github.com/game_engine/card-games/internal/games/baccarat"
	"github.com/game_engine/card-games/internal/games/poker"
	"github.com/game_engine/card-games/internal/service"
)

// GameHandler handles game-related requests
type GameHandler struct {
	gameService *service.GameService
}

// NewGameHandler creates a new game handler
func NewGameHandler(gameService *service.GameService) *GameHandler {
	return &GameHandler{
		gameService: gameService,
	}
}

// CreateBlackjack creates a new blackjack game
func (h *GameHandler) CreateBlackjack(gameID, playerID string, bet int64) (interface{}, error) {
	return h.gameService.CreateBlackjack(gameID, playerID, bet)
}

// CreateBaccarat creates a new baccarat game
func (h *GameHandler) CreateBaccarat(gameID string, bets map[baccarat.BetType]int64) (interface{}, error) {
	return h.gameService.CreateBaccarat(gameID, bets)
}

// CreatePoker creates a new poker game
func (h *GameHandler) CreatePoker(gameID string, gameType poker.GameType) (interface{}, error) {
	return h.gameService.CreatePoker(gameID, gameType)
}

// CreateAndarBahar creates a new Andar Bahar game
func (h *GameHandler) CreateAndarBahar(gameID string) (interface{}, error) {
	return h.gameService.CreateAndarBahar(gameID)
}

// CreateTeenPatti creates a new Teen Patti game
func (h *GameHandler) CreateTeenPatti(gameID string) (interface{}, error) {
	return h.gameService.CreateTeenPatti(gameID)
}

// GetGame gets a game by ID
func (h *GameHandler) GetGame(gameID string) (interface{}, error) {
	return h.gameService.GetGame(gameID)
}
