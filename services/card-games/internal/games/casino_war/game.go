package casino_war

import (
	"errors"
	"fmt"
	"time"

	"github.com/game_engine/card-games/internal/games/common"
)

// CasinoWarGame represents a Casino War card game
type CasinoWarGame struct {
	GameID    string
	PlayerID  string
	shoe      *common.Shoe
	playerCard  *common.Card
	dealerCard  *common.Card
	betAmount  int64
	payout     float64
	Result    string
	Winner    string
	Surrendered bool
	WarCount   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Bet types
const (
	BetWar = "war"
	BetSurrender = "surrender"
)

// Payouts
var payouts = map[string]float64{
	"win":       1.0,
	"war_win":   2.0,    // Win after going to war pays 2:1
	"war_push":  1.0,    // War tie - push (bet returned)
	"surrender": 0.5,    // Surrender pays 0.5:1
}

// NewCasinoWarGame creates a new Casino War game
func NewCasinoWarGame(playerID string) *CasinoWarGame {
	shoe := common.NewShoe(6)

	return &CasinoWarGame{
		GameID:    generateGameID(),
		PlayerID:  playerID,
		shoe:      shoe,
		Result:    "pending",
		WarCount:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// DealInitial deals the initial cards
func (g *CasinoWarGame) DealInitial(amount int64) error {
	if amount <= 0 {
		return errors.New("bet amount must be positive")
	}

	g.betAmount = amount

	// Deal one card to player
	playerCard, err := g.shoe.Draw()
	if err != nil {
		return err
	}
	g.playerCard = &playerCard

	// Deal one card to dealer
	dealerCard, err := g.shoe.Draw()
	if err != nil {
		return err
	}
	g.dealerCard = &dealerCard

	g.Result = "dealt"
	g.UpdatedAt = time.Now()

	return nil
}

// Compare compares the cards and determines the winner
func (g *CasinoWarGame) Compare() (string, int64, error) {
	if g.playerCard == nil || g.dealerCard == nil {
		return "", 0, errors.New("cards not dealt")
	}

	playerValue := getCardValue(g.playerCard)
	dealerValue := getCardValue(g.dealerCard)

	if playerValue > dealerValue {
		g.Winner = "player"
		g.Result = "resolved"
		// Regular win pays 1:1
		payout := float64(g.betAmount) * (1 + payouts["win"])
		g.UpdatedAt = time.Now()
		return g.Winner, int64(payout), nil
	} else if dealerValue > playerValue {
		g.Winner = "dealer"
		g.Result = "resolved"
		g.UpdatedAt = time.Now()
		// Player loses - no payout
		return g.Winner, 0, nil
	} else {
		// War!
		g.Winner = "war"
		g.Result = "war"
		g.UpdatedAt = time.Now()
		return "war", 0, nil
	}
}

// GoToWar player chooses to go to war
func (g *CasinoWarGame) GoToWar() (string, int64, error) {
	if g.Result != "war" {
		return "", 0, errors.New("not in war state")
	}

	g.WarCount++
	
	// In war, player must match the original bet
	// Burn 3 cards, then deal 1 card each
	for i := 0; i < 3; i++ {
		_, err := g.shoe.Draw()
		if err != nil {
			return "", 0, err
		}
	}

	// Deal new cards
	playerCard, err := g.shoe.Draw()
	if err != nil {
		return "", 0, err
	}
	g.playerCard = &playerCard

	dealerCard, err := g.shoe.Draw()
	if err != nil {
		return "", 0, err
	}
	g.dealerCard = &dealerCard

	// Compare again
	playerValue := getCardValue(g.playerCard)
	dealerValue := getCardValue(g.dealerCard)

	if playerValue > dealerValue {
		g.Winner = "player"
		g.Result = "resolved"
		// War win pays 2:1 (original bet + war bet)
		payout := float64(g.betAmount) * 2 * (1 + payouts["war_win"])
		g.UpdatedAt = time.Now()
		return g.Winner, int64(payout), nil
	} else if dealerValue > playerValue {
		g.Winner = "dealer"
		g.Result = "resolved"
		g.UpdatedAt = time.Now()
		return g.Winner, 0, nil
	} else {
		// War tie - it's a push, player gets bet back
		g.Winner = "push"
		g.Result = "resolved"
		payout := float64(g.betAmount) * 2 * (1 + payouts["war_push"])
		g.UpdatedAt = time.Now()
		return g.Winner, int64(payout), nil
	}
}

// Surrender player chooses to surrender
func (g *CasinoWarGame) Surrender() (string, int64, error) {
	if g.Result != "war" {
		return "", 0, errors.New("can only surrender during war")
	}

	g.Surrendered = true
	g.Winner = "dealer"
	g.Result = "resolved"
	g.UpdatedAt = time.Now()

	// Surrender pays 0.5:1
	payout := float64(g.betAmount) * (1 + payouts["surrender"])
	return g.Winner, int64(payout), nil
}

// GetResult returns the game result
func (g *CasinoWarGame) GetResult() map[string]interface{} {
	result := map[string]interface{}{
		"game_id":     g.GameID,
		"bet_amount":  g.betAmount,
		"winner":       g.Winner,
		"result":       g.Result,
		"war_count":   g.WarCount,
		"surrendered": g.Surrendered,
	}

	if g.playerCard != nil {
		result["player_card"] = g.playerCard.String()
		result["player_value"] = getCardValue(g.playerCard)
	}

	if g.dealerCard != nil {
		result["dealer_card"] = g.dealerCard.String()
		result["dealer_value"] = getCardValue(g.dealerCard)
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
	return fmt.Sprintf("cw_%d_%d", time.Now().Unix(), time.Now().Nanosecond()%1000)
}