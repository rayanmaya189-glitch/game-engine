package three_card_poker

import (
	"errors"
	"time"
)

// Reveal reveals dealer's cards and settles
func (g *ThreeCardPokerGame) Reveal() (string, int64, error) {
	if len(g.playerHand) != 3 || len(g.dealerHand) != 3 {
		return "", 0, errors.New("cards not dealt")
	}

	g.DealerHandType = evaluateHand(g.dealerHand)

	dealerQualifies := compareHands(g.dealerHand, g.playerHand) >= 0

	var totalPayout int64
	win := false

	if !dealerQualifies {
		g.Winner = "player"
		g.Result = "resolved"
		totalPayout = g.anteBet * 2
		win = true
	} else {
		result := compareHands(g.dealerHand, g.playerHand)
		if result > 0 {
			g.Winner = "dealer"
			g.Result = "resolved"
		} else if result < 0 {
			g.Winner = "player"
			g.Result = "resolved"
			totalPayout = g.playBet * 2
			win = true
		} else {
			g.Winner = "push"
			g.Result = "resolved"
			totalPayout = g.anteBet + g.playBet
		}
	}

	if win && (g.PlayerHandType == Straight || g.PlayerHandType == ThreeOfAKind || g.PlayerHandType == StraightFlush) {
		anteBonus := float64(g.anteBet) * anteBonusPayouts[g.PlayerHandType]
		totalPayout += int64(anteBonus)
	}

	if g.pairPlusBet > 0 && g.PlayerHandType != HighCard {
		pairPlusPayout := float64(g.pairPlusBet) * pairPlusPayouts[g.PlayerHandType]
		totalPayout += int64(pairPlusPayout)
	}

	g.UpdatedAt = time.Now()
	return g.Winner, totalPayout, nil
}

// GetResult returns the game result
func (g *ThreeCardPokerGame) GetResult() map[string]interface{} {
	result := map[string]interface{}{
		"game_id":       g.GameID,
		"ante_bet":      g.anteBet,
		"play_bet":      g.playBet,
		"pair_plus_bet": g.pairPlusBet,
		"winner":        g.Winner,
		"result":        g.Result,
		"player_hand":   g.PlayerHandType,
	}

	if len(g.playerHand) > 0 {
		cards := make([]string, len(g.playerHand))
		for i, c := range g.playerHand {
			cards[i] = c.String()
		}
		result["player_cards"] = cards
	}

	if len(g.dealerHand) > 0 {
		result["dealer_hand"] = g.DealerHandType
		cards := make([]string, len(g.dealerHand))
		for i, c := range g.dealerHand {
			cards[i] = c.String()
		}
		result["dealer_cards"] = cards
	}

	return result
}
