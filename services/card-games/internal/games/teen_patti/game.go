package teen_patti

import (
	"sort"

	"github.com/game_engine/card-games/internal/games/common"
)

// HandRank represents the rank of a Teen Patti hand
type HandRank int

const (
	HighCardTeen HandRank = iota
	Pair
	Straight
	Flush
	ThreeOfAKind
	StraightFlush
)

// Game represents a Teen Patti game
type Game struct {
	ID        string
	Shoe      *common.Shoe
	Players   map[string][]common.Card
	Community []common.Card
	Config    *Config
}

// Config holds Teen Patti configuration
type Config struct {
	Variant    string // "classic", "999", "boat"
	MinPlayers int
	MaxPlayers int
}

// NewGame creates a new Teen Patti game
func NewGame(id string, shoe *common.Shoe, config *Config) *Game {
	return &Game{
		ID:        id,
		Shoe:      shoe,
		Players:   make(map[string][]common.Card),
		Community: make([]common.Card, 0),
		Config:    config,
	}
}

// AddPlayer adds a player
func (g *Game) AddPlayer(playerID string) error {
	if len(g.Players) >= g.Config.MaxPlayers {
		return nil // Error: max players
	}

	g.Players[playerID] = make([]common.Card, 0)
	return nil
}

// DealHoleCards deals 3 cards to each player
func (g *Game) DealHoleCards() error {
	for playerID := range g.Players {
		for i := 0; i < 3; i++ {
			card, err := g.Shoe.Draw()
			if err != nil {
				return err
			}
			g.Players[playerID] = append(g.Players[playerID], card)
		}
	}
	return nil
}

// EvaluateHands evaluates all players' hands
func (g *Game) EvaluateHands() map[string]HandRank {
	results := make(map[string]HandRank)

	for playerID, cards := range g.Players {
		rank := evaluateHand(cards)
		results[playerID] = rank
	}

	return results
}

func evaluateHand(cards []common.Card) HandRank {
	// Sort cards by rank descending
	sorted := make([]common.Card, len(cards))
	copy(sorted, cards)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Rank > sorted[j].Rank
	})

	// Check for straight flush
	if isStraightFlush(sorted) {
		return StraightFlush
	}

	// Check for three of a kind
	if isThreeOfAKind(sorted) {
		return ThreeOfAKind
	}

	// Check for flush
	if isFlush(sorted) {
		return Flush
	}

	// Check for straight
	if isStraight(sorted) {
		return Straight
	}

	// Check for pair
	if isPair(sorted) {
		return Pair
	}

	return HighCardTeen
}

func isThreeOfAKind(cards []common.Card) bool {
	if len(cards) < 3 {
		return false
	}
	return cards[0].Rank == cards[1].Rank && cards[1].Rank == cards[2].Rank
}

func isFlush(cards []common.Card) bool {
	if len(cards) < 3 {
		return false
	}
	return cards[0].Suit == cards[1].Suit && cards[1].Suit == cards[2].Suit
}

func isStraight(cards []common.Card) bool {
	if len(cards) < 3 {
		return false
	}

	// Check for regular straight
	if cards[0].Rank-cards[1].Rank == 1 && cards[1].Rank-cards[2].Rank == 1 {
		return true
	}

	// Check for A-2-3 (wheel)
	if cards[0].Rank == common.Ace && cards[1].Rank == common.Three && cards[2].Rank == common.Two {
		return true
	}

	// Check for Q-K-A
	if cards[0].Rank == common.Ace && cards[1].Rank == common.King && cards[2].Rank == common.Queen {
		return true
	}

	return false
}

func isStraightFlush(cards []common.Card) bool {
	return isFlush(cards) && isStraight(cards)
}

func isPair(cards []common.Card) bool {
	if len(cards) < 2 {
		return false
	}
	return cards[0].Rank == cards[1].Rank || cards[1].Rank == cards[2].Rank || cards[0].Rank == cards[2].Rank
}

// GetState returns the game state
func (g *Game) GetState() *GameState {
	state := &GameState{
		GameID:  g.ID,
		Players: make(map[string][]common.Card),
	}

	for playerID, cards := range g.Players {
		state.Players[playerID] = cards
	}

	return state
}

// GameState represents the game state
type GameState struct {
	GameID  string                   `json:"game_id"`
	Players map[string][]common.Card `json:"players"`
}
