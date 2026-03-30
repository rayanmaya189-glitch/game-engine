package blackjack

import "github.com/game_engine/card-games/internal/games/common"

// calculateTotal calculates the total of a hand
func calculateTotal(cards []common.Card) int {
	total := 0
	aces := 0

	for _, card := range cards {
		if card.IsAce() {
			aces++
			total += 11
		} else {
			total += card.Value()
		}
	}

	for aces > 0 && total > 21 {
		total -= 10
		aces--
	}

	return total
}

// isSoft returns true if the hand is soft (Ace counted as 11)
func isSoft(cards []common.Card) bool {
	total := 0
	aces := 0

	for _, card := range cards {
		if card.IsAce() {
			aces++
			total += 11
		} else {
			total += card.Value()
		}
	}

	return aces > 0 && total <= 21
}

// isBlackjack returns true if the hand is a natural blackjack
func isBlackjack(cards []common.Card) bool {
	if len(cards) != 2 {
		return false
	}

	card1, card2 := cards[0], cards[1]
	hasAce := card1.IsAce() || card2.IsAce()
	hasTen := card1.Value() == 10 || card2.Value() == 10

	return hasAce && hasTen
}

// FormatHand formats a hand for display
func FormatHand(cards []common.Card) string {
	result := ""
	for i, card := range cards {
		if i > 0 {
			result += ", "
		}
		result += card.String()
	}
	return result
}
