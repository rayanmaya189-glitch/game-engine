package three_card_poker

import (
	"errors"
	"time"

	"github.com/game_engine/card-games/internal/games/common"
)

// ThreeCardPokerGame represents a Three Card Poker game
type ThreeCardPokerGame struct {
	GameID         string
	PlayerID       string
	shoe           *common.Shoe
	playerHand     []common.Card
	dealerHand     []common.Card
	anteBet        int64
	playBet        int64
	pairPlusBet    int64
	Result         string
	Winner         string
	PlayerHandType string
	DealerHandType string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// NewThreeCardPokerGame creates a new Three Card Poker game
func NewThreeCardPokerGame(playerID string) *ThreeCardPokerGame {
	shoe := common.NewShoe(6)

	return &ThreeCardPokerGame{
		GameID:    generateGameID(),
		PlayerID:  playerID,
		shoe:      shoe,
		Result:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// PlaceAnte places the ante and optional pair plus bets
func (g *ThreeCardPokerGame) PlaceAnte(ante, pairPlus int64) error {
	if ante <= 0 {
		return errors.New("ante bet must be positive")
	}

	g.anteBet = ante
	g.pairPlusBet = pairPlus
	g.Result = "ante_placed"
	g.UpdatedAt = time.Now()

	return nil
}

// Deal deals 3 cards to player and dealer
func (g *ThreeCardPokerGame) Deal() error {
	if g.anteBet == 0 {
		return errors.New("ante not placed")
	}

	for i := 0; i < 3; i++ {
		card, err := g.shoe.Draw()
		if err != nil {
			return err
		}
		g.playerHand = append(g.playerHand, card)
	}

	for i := 0; i < 3; i++ {
		card, err := g.shoe.Draw()
		if err != nil {
			return err
		}
		g.dealerHand = append(g.dealerHand, card)
	}

	g.PlayerHandType = evaluateHand(g.playerHand)
	g.Result = "dealt"
	g.UpdatedAt = time.Now()

	return nil
}

// Play chooses to play (match ante) or fold
func (g *ThreeCardPokerGame) Play() error {
	if len(g.playerHand) != 3 || len(g.dealerHand) != 3 {
		return errors.New("cards not dealt")
	}

	g.playBet = g.anteBet
	g.Result = "playing"
	g.UpdatedAt = time.Now()

	return nil
}

// Fold folds the hand
func (g *ThreeCardPokerGame) Fold() (string, int64, error) {
	if len(g.playerHand) != 3 {
		return "", 0, errors.New("cards not dealt")
	}

	g.Result = "folded"
	g.UpdatedAt = time.Now()

	return "dealer", 0, nil
}
