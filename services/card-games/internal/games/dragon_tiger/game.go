package dragon_tiger

import (
	"errors"
	"fmt"
	"time"

	"github.com/game_engine/card-games/internal/games/common"
)

// DragonTigerGame represents a Dragon Tiger card game
type DragonTigerGame struct {
	GameID     string
	PlayerID   string
	shoe       *common.Shoe
	dragonCard *common.Card
	tigerCard  *common.Card
	betType    string
	betAmount  int64
	payout     float64
	Result     string
	Winner     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Bet types
const (
	BetDragon     = "dragon"
	BetTiger      = "tiger"
	BetTie        = "tie"
	BetDragonOdd  = "dragon_odd"
	BetDragonEven = "dragon_even"
	BetTigerOdd   = "tiger_odd"
	BetTigerEven  = "tiger_even"
)

// Payouts
var payouts = map[string]float64{
	BetDragon:     1.0,
	BetTiger:      1.0,
	BetTie:        11.0,
	BetDragonOdd:  0.92,
	BetDragonEven: 0.92,
	BetTigerOdd:   0.92,
	BetTigerEven:  0.92,
}

// NewDragonTigerGame creates a new Dragon Tiger game
func NewDragonTigerGame(playerID string) *DragonTigerGame {
	shoe := common.NewShoe(6) // 6 deck shoe

	return &DragonTigerGame{
		GameID:    generateGameID(),
		PlayerID:  playerID,
		shoe:      shoe,
		Result:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// PlaceBet places a bet on Dragon, Tiger, or Tie
func (g *DragonTigerGame) PlaceBet(betType string, amount int64) error {
	if g.dragonCard != nil {
		return errors.New("cards already dealt")
	}

	validBets := []string{BetDragon, BetTiger, BetTie, BetDragonOdd, BetDragonEven, BetTigerOdd, BetTigerEven}
	valid := false
	for _, b := range validBets {
		if b == betType {
			valid = true
			break
		}
	}

	if !valid {
		return errors.New("invalid bet type")
	}

	g.betType = betType
	g.betAmount = amount
	g.UpdatedAt = time.Now()

	return nil
}

// Deal deals the cards for Dragon and Tiger
func (g *DragonTigerGame) Deal() error {
	if g.betType == "" {
		return errors.New("bet not placed")
	}

	// Deal one card to Dragon
	dragonCard, err := g.shoe.Draw()
	if err != nil {
		return err
	}
	g.dragonCard = &dragonCard

	// Deal one card to Tiger
	tigerCard, err := g.shoe.Draw()
	if err != nil {
		return err
	}
	g.tigerCard = &tigerCard

	g.Result = "dealt"
	g.UpdatedAt = time.Now()

	return nil
}

// Settle settles the bet and determines the winner
func (g *DragonTigerGame) Settle() (string, int64, error) {
	if g.dragonCard == nil || g.tigerCard == nil {
		return "", 0, errors.New("cards not dealt")
	}

	dragonValue := getCardValue(g.dragonCard)
	tigerValue := getCardValue(g.tigerCard)

	// Determine winner
	if dragonValue > tigerValue {
		g.Winner = BetDragon
	} else if tigerValue > dragonValue {
		g.Winner = BetTiger
	} else {
		g.Winner = BetTie
	}

	// Calculate payout
	win := false

	// Main bets
	if g.betType == g.Winner {
		win = true
	} else if g.betType == BetTie && g.Winner == BetTie {
		win = true
	}

	// Odd/Even bets
	if g.betType == BetDragonOdd && dragonValue%2 == 1 {
		win = true
	} else if g.betType == BetDragonEven && dragonValue%2 == 0 && dragonValue != 0 {
		win = true
	} else if g.betType == BetTigerOdd && tigerValue%2 == 1 {
		win = true
	} else if g.betType == BetTigerEven && tigerValue%2 == 0 && tigerValue != 0 {
		win = true
	}

	g.Result = "settled"
	g.UpdatedAt = time.Now()

	if win {
		payout := float64(g.betAmount) * (1 + payouts[g.betType])
		return g.Winner, int64(payout), nil
	}

	return g.Winner, 0, nil
}

// GetResult returns the game result
func (g *DragonTigerGame) GetResult() map[string]interface{} {
	result := map[string]interface{}{
		"game_id":    g.GameID,
		"bet_type":   g.betType,
		"bet_amount": g.betAmount,
		"winner":     g.Winner,
		"result":     g.Result,
	}

	if g.dragonCard != nil {
		result["dragon_card"] = g.dragonCard.String()
		result["dragon_value"] = getCardValue(g.dragonCard)
	}

	if g.tigerCard != nil {
		result["tiger_card"] = g.tigerCard.String()
		result["tiger_value"] = getCardValue(g.tigerCard)
	}

	return result
}

// GetCardValue returns the value of a card (A=1, 2-10=face value, J=11, Q=12, K=13)
func getCardValue(card *common.Card) int {
	switch card.Rank {
	case common.Ace:
		return 1
	case common.Two:
		return 2
	case common.Three:
		return 3
	case common.Four:
		return 4
	case common.Five:
		return 5
	case common.Six:
		return 6
	case common.Seven:
		return 7
	case common.Eight:
		return 8
	case common.Nine:
		return 9
	case common.Ten, common.Jack, common.Queen, common.King:
		return 10
	default:
		return 0
	}
}

func generateGameID() string {
	return fmt.Sprintf("dt_%d_%d", time.Now().Unix(), time.Now().Nanosecond()%1000)
}
