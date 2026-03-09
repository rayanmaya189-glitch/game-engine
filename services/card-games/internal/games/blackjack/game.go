package blackjack

import (
	"errors"
	"fmt"

	"github.com/gameengine/card-games/internal/games/common"
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
	ID            string       `json:"id"`
	Hands         []*Hand      `json:"hands"`
	CurrentHand   int          `json:"current_hand"`
	InsuranceBet  int64        `json:"insurance_bet"`
	CurrentAction Action       `json:"current_action"`
	Result        Result       `json:"result"`
}

// Hand represents a player's hand in blackjack
type Hand struct {
	Cards          []common.Card `json:"cards"`
	Bet            int64         `json:"bet"`
	IsDoubled      bool          `json:"is_doubled"`
	IsSplit        bool          `json:"is_split"`
	CanSplit       bool          `json:"can_split"`
	Surrendered    bool          `json:"surrendered"`
	Result         Result        `json:"result"`
}

// Dealer represents the dealer
type Dealer struct {
	HoleCard   common.Card `json:"hole_card"`
	Hand       []common.Card `json:"hand"`
	GameState  GameState    `json:"game_state"`
}

// Game represents a blackjack game
type Game struct {
	ID                string
	Shoe              *common.Shoe
	Dealer            *Dealer
	Players           map[string]*Player
	Config            *Config
	GameState         GameState
	CurrentPlayerID   string
}

// Config holds blackjack configuration
type Config struct {
	AllowSurrender       bool
	AllowLateSurrender  bool
	DealerStandsOnSoft17 bool
	MaxSplits           int
	BlackjackPayout     float64
}

// NewGame creates a new blackjack game
func NewGame(id string, shoe *common.Shoe, config *Config) *Game {
	return &Game{
		ID:              id,
		Shoe:            shoe,
		Dealer:          &Dealer{},
		Players:         make(map[string]*Player),
		Config:          config,
		GameState:       StateWaiting,
	}
}

// AddPlayer adds a player to the game
func (g *Game) AddPlayer(playerID string, bet int64) error {
	if _, exists := g.Players[playerID]; exists {
		return errors.New("player already exists")
	}

	// Draw initial cards
	card1, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	card2, err := g.Shoe.Draw()
	if err != nil {
		return err
	}

	player := &Player{
		ID: playerID,
		Hands: []*Hand{
			{
				Cards:     []common.Card{card1, card2},
				Bet:       bet,
				IsDoubled: false,
				IsSplit:   false,
				CanSplit:  card1.Rank == card2.Rank,
				Result:    ResultPending,
			},
		},
		CurrentHand: 0,
		Result:      ResultPending,
	}

	g.Players[playerID] = player
	return nil
}

// Start starts the game
func (g *Game) Start() error {
	// Dealer's hole card
	holeCard, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	g.Dealer.HoleCard = holeCard

	// Check for dealer blackjack
	dealerTotal := calculateTotal(g.Dealer.Hand)
	if dealerTotal == 21 {
		g.GameState = StateGameOver
		g.settleDealerBlackjack()
		return nil
	}

	// Check for player blackjacks
	g.GameState = StatePlayerTurn
	for _, player := range g.Players {
		for _, hand := range player.Hands {
			if isBlackjack(hand.Cards) {
				hand.Result = ResultBlackjack
			}
