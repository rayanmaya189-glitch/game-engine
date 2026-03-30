package blackjack

import (
	"fmt"

	"github.com/game_engine/card-games/internal/games/common"
)

// Action represents a player's action in blackjack
type Action int

const (
	ActionHit Action = iota
	ActionStand
	ActionDouble
	ActionSplit
	ActionSurrender
)

// GameState represents the current state of a blackjack game
type GameState int

const (
	StateWaiting GameState = iota
	StatePlayerTurn
	StateDealerTurn
	StateGameOver
)

// Result represents the outcome of a blackjack hand
type Result int

const (
	ResultPending Result = iota
	ResultWin
	ResultLoss
	ResultPush
	ResultBlackjack
	ResultSurrender
)

// Player represents a player in the game
type Player struct {
	ID            string  `json:"id"`
	Hands         []*Hand `json:"hands"`
	CurrentHand   int     `json:"current_hand"`
	InsuranceBet  int64   `json:"insurance_bet"`
	CurrentAction Action  `json:"current_action"`
	Result        Result  `json:"result"`
}

// Hand represents a player's hand in blackjack
type Hand struct {
	Cards       []common.Card `json:"cards"`
	Bet         int64         `json:"bet"`
	IsDoubled   bool          `json:"is_doubled"`
	IsSplit     bool          `json:"is_split"`
	CanSplit    bool          `json:"can_split"`
	Surrendered bool          `json:"surrendered"`
	Result      Result        `json:"result"`
}

// Dealer represents the dealer
type Dealer struct {
	HoleCard  common.Card   `json:"hole_card"`
	Hand      []common.Card `json:"hand"`
	GameState GameState     `json:"game_state"`
}

// Game represents a blackjack game
type Game struct {
	ID              string
	Shoe            *common.Shoe
	Dealer          *Dealer
	Players         map[string]*Player
	Config          *Config
	GameState       GameState
	CurrentPlayerID string
}

// Config holds blackjack configuration
type Config struct {
	AllowSurrender       bool
	AllowLateSurrender   bool
	DealerStandsOnSoft17 bool
	MaxSplits            int
	BlackjackPayout      float64
}

// GameStateResponse represents the game state for clients
type GameStateResponse struct {
	GameID          string        `json:"game_id"`
	GameState       string        `json:"game_state"`
	DealerHand      []common.Card `json:"dealer_hand"`
	DealerTotal     int           `json:"dealer_total"`
	Players         []PlayerState `json:"players"`
	CurrentPlayerID string        `json:"current_player_id"`
}

// PlayerState represents a player's state
type PlayerState struct {
	ID          string      `json:"id"`
	Hands       []HandState `json:"hands"`
	CurrentHand int         `json:"current_hand"`
}

// HandState represents a hand's state
type HandState struct {
	Cards       []common.Card `json:"cards"`
	Bet         int64         `json:"bet"`
	Total       int           `json:"total"`
	IsDoubled   bool          `json:"is_doubled"`
	IsSplit     bool          `json:"is_split"`
	Surrendered bool          `json:"surrendered"`
	Result      string        `json:"result"`
}

// String returns a string representation of Action
func (a Action) String() string {
	switch a {
	case ActionHit:
		return "hit"
	case ActionStand:
		return "stand"
	case ActionDouble:
		return "double"
	case ActionSplit:
		return "split"
	case ActionSurrender:
		return "surrender"
	default:
		return "unknown"
	}
}

// String returns a string representation of Result
func (r Result) String() string {
	switch r {
	case ResultPending:
		return "pending"
	case ResultWin:
		return "win"
	case ResultLoss:
		return "loss"
	case ResultPush:
		return "push"
	case ResultBlackjack:
		return "blackjack"
	case ResultSurrender:
		return "surrender"
	default:
		return "unknown"
	}
}

func init() {
	// Register string converters
	_ = fmt.Sprintf("%s", ActionHit)
	_ = fmt.Sprintf("%s", ResultWin)
}
