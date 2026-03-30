package blackjack

import (
	"errors"

	"github.com/game_engine/card-games/internal/games/common"
)

// Hit draws a card
func (g *Game) hit(player *Player, hand *Hand) error {
	card, err := g.Shoe.Draw()
	if err != nil {
		return err
	}

	hand.Cards = append(hand.Cards, card)

	// Check if busted
	total := calculateTotal(hand.Cards)
	if total > 21 {
		hand.Result = ResultLoss
		return g.nextHand()
	}

	// Check if can still hit
	return nil
}

// Stand moves to the next hand
func (g *Game) stand(player *Player) error {
	return g.nextHand()
}

// Double doubles the bet and draws one card
func (g *Game) double(player *Player, hand *Hand) error {
	if hand.IsDoubled {
		return errors.New("already doubled")
	}

	hand.Bet *= 2
	hand.IsDoubled = true

	card, err := g.Shoe.Draw()
	if err != nil {
		return err
	}

	hand.Cards = append(hand.Cards, card)

	// Check if busted
	total := calculateTotal(hand.Cards)
	if total > 21 {
		hand.Result = ResultLoss
	}

	return g.nextHand()
}

// Split splits the hand into two
func (g *Game) split(player *Player, hand *Hand) error {
	if !hand.CanSplit {
		return errors.New("cannot split")
	}

	if len(player.Hands) >= g.Config.MaxSplits {
		return errors.New("max splits reached")
	}

	// Get the two cards
	card1 := hand.Cards[0]
	card2 := hand.Cards[1]

	// Create two new hands
	newHand1 := &Hand{
		Cards:    []common.Card{card1},
		Bet:      hand.Bet,
		IsSplit:  true,
		CanSplit: card1.Rank == card1.Rank, // Can split Aces only once
		Result:   ResultPending,
	}

	newHand2 := &Hand{
		Cards:    []common.Card{card2},
		Bet:      hand.Bet,
		IsSplit:  true,
		CanSplit: card2.Rank == card2.Rank,
		Result:   ResultPending,
	}

	// Draw one card for each new hand
	draw1, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	newHand1.Cards = append(newHand1.Cards, draw1)

	draw2, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	newHand2.Cards = append(newHand2.Cards, draw2)

	// Replace the original hand
	player.Hands[player.CurrentHand] = newHand1
	player.Hands = append(player.Hands, newHand2)

	return nil
}

// Surrender forfeits half the bet
func (g *Game) surrender(player *Player, hand *Hand) error {
	if !g.Config.AllowSurrender {
		return errors.New("surrender not allowed")
	}

	hand.Surrendered = true
	hand.Result = ResultSurrender
	hand.Bet /= 2 // Return half

	return g.nextHand()
}

// nextHand moves to the next hand
func (g *Game) nextHand() error {
	player := g.Players[g.CurrentPlayerID]

	// Move to next hand
	player.CurrentHand++

	if player.CurrentHand >= len(player.Hands) {
		// All hands done for this player
		if !g.setNextPlayer() {
			// All players done, dealer's turn
			g.GameState = StateDealerTurn
		}
	}

	return nil
}

// setNextPlayer sets the next player
func (g *Game) setNextPlayer() bool {
	// Simple implementation - find next player with pending hands
	// In production, would maintain player order
	for _, player := range g.Players {
		if player.CurrentHand < len(player.Hands) {
			hand := player.Hands[player.CurrentHand]
			if hand.Result == ResultPending && !hand.Surrendered {
				g.CurrentPlayerID = player.ID
				return true
			}
		}
	}
	return false
}
