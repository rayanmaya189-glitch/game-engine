package poker

import (
	"github.com/game_engine/card-games/internal/games/common"
)

// HandRank represents the rank of a poker hand
type HandRank int

const (
	HighCard HandRank = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

// GameType represents the type of poker game
type GameType int

const (
	TexasHoldem GameType = iota
	Omaha
)

// Street represents the betting street
type Street int

const (
	PreFlop Street = iota
	Flop
	Turn
	River
	Showdown
)

// Player represents a poker player
type Player struct {
	ID        string        `json:"id"`
	HoleCards []common.Card `json:"hole_cards"`
	Chips     int64         `json:"chips"`
	Bet       int64         `json:"bet"`
	HandRank  HandRank      `json:"hand_rank"`
	BestHand  []common.Card `json:"best_hand"`
	Folded    bool          `json:"folded"`
	AllIn     bool          `json:"all_in"`
}

// Game represents a poker game
type Game struct {
	ID            string
	GameType      GameType
	Shoe          *common.Shoe
	Players       map[string]*Player
	Community     []common.Card
	CurrentStreet Street
	Pot           int64
	SidePots      []int64
	Config        *Config
}

// Config holds poker configuration
type Config struct {
	MinPlayers    int
	MaxPlayers    int
	SmallBlind    int64
	BigBlind      int64
	StartingChips int64
}

// GameState represents the game state for clients
type GameState struct {
	GameID        string        `json:"game_id"`
	GameType      string        `json:"game_type"`
	Community     []common.Card `json:"community"`
	CurrentStreet string        `json:"current_street"`
	Pot           int64         `json:"pot"`
	Players       []PlayerState `json:"players"`
}

// PlayerState represents a player's state
type PlayerState struct {
	ID        string        `json:"id"`
	HoleCards []common.Card `json:"hole_cards"`
	Chips     int64         `json:"chips"`
	Bet       int64         `json:"bet"`
	Folded    bool          `json:"folded"`
	AllIn     bool          `json:"all_in"`
	HandRank  string        `json:"hand_rank"`
	BestHand  []common.Card `json:"best_hand"`
}
