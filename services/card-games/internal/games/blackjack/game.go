package blackjack

import (
	"errors"

	"github.com/game_engine/card-games/internal/games/common"
)

// NewGame creates a new blackjack game
func NewGame(id string, shoe *common.Shoe, config *Config) *Game {
	return &Game{
		ID:        id,
		Shoe:      shoe,
		Dealer:    &Dealer{},
		Players:   make(map[string]*Player),
		Config:    config,
		GameState: StateWaiting,
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
		}
	}

	// Set first player
	g.setNextPlayer()

	return nil
}

// PlayerAction performs an action for the current player
func (g *Game) PlayerAction(playerID string, action Action) error {
	if g.GameState != StatePlayerTurn {
		return errors.New("not player's turn")
	}

	if g.CurrentPlayerID != playerID {
		return errors.New("not current player")
	}

	player, exists := g.Players[playerID]
	if !exists {
		return errors.New("player not found")
	}

	hand := player.Hands[player.CurrentHand]
	if hand == nil {
		return errors.New("hand not found")
	}

	switch action {
	case ActionHit:
		return g.hit(player, hand)
	case ActionStand:
		return g.stand(player)
	case ActionDouble:
		return g.double(player, hand)
	case ActionSplit:
		return g.split(player, hand)
	case ActionSurrender:
		return g.surrender(player, hand)
	default:
		return errors.New("invalid action")
	}
}

// DealerPlay plays the dealer's hand
func (g *Game) DealerPlay() error {
	if g.GameState != StateDealerTurn {
		return errors.New("not dealer's turn")
	}

	// Reveal hole card
	g.Dealer.Hand = append(g.Dealer.Hand, g.Dealer.HoleCard)

	// Dealer draws according to rules
	for {
		total := calculateTotal(g.Dealer.Hand)
		soft := isSoft(g.Dealer.Hand)

		// Dealer rules
		if total > 21 {
			break // Busted
		}
		if total > 17 {
			break // Stand
		}
		if total == 17 && soft && !g.Config.DealerStandsOnSoft17 {
			// Hit on soft 17
		}
		if total < 17 {
			// Hit
		} else {
			break
		}

		card, err := g.Shoe.Draw()
		if err != nil {
			return err
		}
		g.Dealer.Hand = append(g.Dealer.Hand, card)
	}

	g.GameState = StateGameOver
	g.settle()

	return nil
}

// GetState returns the current game state
func (g *Game) GetState() *GameStateResponse {
	response := &GameStateResponse{
		GameID:          g.ID,
		GameState:       string(g.GameState),
		DealerHand:      g.Dealer.Hand,
		DealerTotal:     calculateTotal(g.Dealer.Hand),
		Players:         make([]PlayerState, 0),
		CurrentPlayerID: g.CurrentPlayerID,
	}

	for _, player := range g.Players {
		playerState := PlayerState{
			ID:          player.ID,
			Hands:       make([]HandState, len(player.Hands)),
			CurrentHand: player.CurrentHand,
		}

		for i, hand := range player.Hands {
			playerState.Hands[i] = HandState{
				Cards:       hand.Cards,
				Bet:         hand.Bet,
				Total:       calculateTotal(hand.Cards),
				IsDoubled:   hand.IsDoubled,
				IsSplit:     hand.IsSplit,
				Surrendered: hand.Surrendered,
				Result:      string(hand.Result),
			}
		}

		response.Players = append(response.Players, playerState)
	}

	return response
}
