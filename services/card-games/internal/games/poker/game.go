package poker

import (
	"errors"

	"github.com/game_engine/card-games/internal/games/common"
)

// NewGame creates a new poker game
func NewGame(id string, gameType GameType, shoe *common.Shoe, config *Config) *Game {
	return &Game{
		ID:            id,
		GameType:      gameType,
		Shoe:          shoe,
		Players:       make(map[string]*Player),
		Community:     make([]common.Card, 0),
		CurrentStreet: PreFlop,
		Pot:           0,
		SidePots:      make([]int64, 0),
		Config:        config,
	}
}

// AddPlayer adds a player to the game
func (g *Game) AddPlayer(playerID string) error {
	if len(g.Players) >= g.Config.MaxPlayers {
		return errors.New("max players reached")
	}

	player := &Player{
		ID:        playerID,
		HoleCards: make([]common.Card, 0),
		Chips:     g.Config.StartingChips,
		Bet:       0,
		Folded:    false,
		AllIn:     false,
	}

	g.Players[playerID] = player
	return nil
}

// DealHoleCards deals hole cards to players
func (g *Game) DealHoleCards() error {
	// Deal 2 cards for Texas Hold'em, 4 for Omaha
	cardsPerPlayer := 2
	if g.GameType == Omaha {
		cardsPerPlayer = 4
	}

	for range g.Players {
		for i := 0; i < cardsPerPlayer; i++ {
			card, err := g.Shoe.Draw()
			if err != nil {
				return err
			}

			// Find a player who needs a card
			for _, player := range g.Players {
				if len(player.HoleCards) < cardsPerPlayer {
					player.HoleCards = append(player.HoleCards, card)
					break
				}
			}
		}
	}

	return nil
}

// DealCommunity deals community cards
func (g *Game) DealCommunity(count int) error {
	for i := 0; i < count; i++ {
		card, err := g.Shoe.Draw()
		if err != nil {
			return err
		}
		g.Community = append(g.Community, card)
	}

	// Update street
	switch g.CurrentStreet {
	case PreFlop:
		g.CurrentStreet = Flop
	case Flop:
		g.CurrentStreet = Turn
	case Turn:
		g.CurrentStreet = River
	case River:
		g.CurrentStreet = Showdown
	}

	return nil
}

// GetState returns the current game state
func (g *Game) GetState() *GameState {
	state := &GameState{
		GameID:        g.ID,
		GameType:      string(g.GameType),
		Community:     g.Community,
		CurrentStreet: string(g.CurrentStreet),
		Pot:           g.Pot,
		Players:       make([]PlayerState, 0),
	}

	for _, player := range g.Players {
		state.Players = append(state.Players, PlayerState{
			ID:        player.ID,
			HoleCards: player.HoleCards,
			Chips:     player.Chips,
			Bet:       player.Bet,
			Folded:    player.Folded,
			AllIn:     player.AllIn,
			HandRank:  string(player.HandRank),
			BestHand:  player.BestHand,
		})
	}

	return state
}
