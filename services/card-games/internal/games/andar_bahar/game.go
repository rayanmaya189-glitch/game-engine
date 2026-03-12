package andar_bahar

import (
	"errors"

	"github.com/game_engine/card-games/internal/games/common"
)

// Game represents an Andar Bahar game
type Game struct {
	ID          string
	Shoe        *common.Shoe
	Joker       common.Card
	AndarCards  []common.Card
	BaharCards  []common.Card
	WinningSide string // "andar" or "bahar"
	Config      *Config
}

// Config holds Andar Bahar configuration
type Config struct {
	SideToDealFirst   string
	MaxCommunityCards int
}

// NewGame creates a new Andar Bahar game
func NewGame(id string, shoe *common.Shoe, config *Config) *Game {
	return &Game{
		ID:          id,
		Shoe:        shoe,
		AndarCards:  make([]common.Card, 0),
		BaharCards:  make([]common.Card, 0),
		WinningSide: "",
		Config:      config,
	}
}

// SetJoker sets the joker (community card)
func (g *Game) SetJoker(card common.Card) {
	g.Joker = card
}

// DealAndar deals a card to Andar
func (g *Game) DealAndar() (common.Card, error) {
	card, err := g.Shoe.Draw()
	if err != nil {
		return common.Card{}, err
	}
	g.AndarCards = append(g.AndarCards, card)

	// Check for winner
	if card.Rank == g.Joker.Rank {
		g.WinningSide = "andar"
	}

	return card, nil
}

// DealBahar deals a card to Bahar
func (g *Game) DealBahar() (common.Card, error) {
	card, err := g.Shoe.Draw()
	if err != nil {
		return common.Card{}, err
	}
	g.BaharCards = append(g.BaharCards, card)

	// Check for winner
	if card.Rank == g.Joker.Rank {
		g.WinningSide = "bahar"
	}

	return card, nil
}

// DealAlternating deals cards alternating between Andar and Bahar
func (g *Game) DealAlternating(startSide string) error {
	currentSide := startSide

	for {
		// Check for winner
		if g.WinningSide != "" {
			break
		}

		// Check max cards
		if len(g.AndarCards)+len(g.BaharCards) >= g.Config.MaxCommunityCards {
			return errors.New("max community cards reached")
		}

		if currentSide == "andar" {
			_, err := g.DealAndar()
			if err != nil {
				return err
			}
			currentSide = "bahar"
		} else {
			_, err := g.DealBahar()
			if err != nil {
				return err
			}
			currentSide = "andar"
		}
	}

	return nil
}

// GetPayout calculates the payout based on the bet
func (g *Game) GetPayout(betSide string, betAmount int64) (int64, error) {
	if g.WinningSide == "" {
		return 0, errors.New("game not finished")
	}

	if betSide == g.WinningSide {
		// Calculate payout based on position
		var position int
		if g.WinningSide == "andar" {
			position = len(g.AndarCards)
		} else {
			position = len(g.BaharCards)
		}

		// Payout table (simplified)
		payoutMultiplier := getPayoutMultiplier(position)
		return betAmount * payoutMultiplier, nil
	}

	return -betAmount, nil
}

func getPayoutMultiplier(position int) int64 {
	switch {
	case position <= 5:
		return 5
	case position <= 10:
		return 4
	case position <= 15:
		return 3
	case position <= 20:
		return 2
	default:
		return 1
	}
}

// GetState returns the current game state
func (g *Game) GetState() *GameState {
	return &GameState{
		GameID:      g.ID,
		Joker:       g.Joker,
		AndarCards:  g.AndarCards,
		BaharCards:  g.BaharCards,
		WinningSide: g.WinningSide,
	}
}

// GameState represents the game state
type GameState struct {
	GameID      string        `json:"game_id"`
	Joker       common.Card   `json:"joker"`
	AndarCards  []common.Card `json:"andar_cards"`
	BaharCards  []common.Card `json:"bahar_cards"`
	WinningSide string        `json:"winning_side"`
}
