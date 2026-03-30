package poker

import (
	"sort"

	"github.com/game_engine/card-games/internal/games/common"
)

// EvaluateHands evaluates all players' hands
func (g *Game) EvaluateHands() {
	for _, player := range g.Players {
		if player.Folded {
			continue
		}

		var bestHand []common.Card
		var bestRank HandRank

		// Combine hole cards with community cards
		allCards := append(player.HoleCards, g.Community...)

		// Find the best 5-card hand
		if g.GameType == TexasHoldem {
			bestHand, bestRank = evaluate5CardDraw(allCards, 2)
		} else {
			// Omaha: must use exactly 2 hole cards and 3 community cards
			bestHand, bestRank = evaluateOmahaHand(player.HoleCards, g.Community)
		}

		player.BestHand = bestHand
		player.HandRank = bestRank
	}
}

// evaluate5CardDraw evaluates the best 5-card hand from available cards
func evaluate5CardDraw(availableCards []common.Card, holeCardsToUse int) ([]common.Card, HandRank) {
	if len(availableCards) < 5 {
		return nil, HighCard
	}

	// Find the best combination
	// Check for straight flush
	if hand := checkStraightFlush(availableCards); hand != nil {
		return hand, StraightFlush
	}

	// Check for four of a kind
	if hand, rank := checkFourOfAKind(availableCards); hand != nil {
		return hand, rank
	}

	// Check for full house
	if hand, rank := checkFullHouse(availableCards); hand != nil {
		return hand, rank
	}

	// Check for flush
	if hand, rank := checkFlush(availableCards); hand != nil {
		return hand, rank
	}

	// Check for straight
	if hand := checkStraight(availableCards); hand != nil {
		return hand, Straight
	}

	// Check for three of a kind
	if hand, rank := checkThreeOfAKind(availableCards); hand != nil {
		return hand, rank
	}

	// Check for two pair
	if hand, rank := checkTwoPair(availableCards); hand != nil {
		return hand, rank
	}

	// Check for pair
	if hand, rank := checkPair(availableCards); hand != nil {
		return hand, rank
	}

	// High card
	return getHighCard(availableCards), HighCard
}

// evaluateOmahaHand evaluates the best Omaha hand
func evaluateOmahaHand(holeCards []common.Card, community []common.Card) ([]common.Card, HandRank) {
	if len(holeCards) < 4 || len(community) < 3 {
		return nil, HighCard
	}

	var bestHand []common.Card
	var bestRank HandRank

	// Try all combinations of 2 hole cards and 3 community cards
	for i := 0; i < len(holeCards); i++ {
		for j := i + 1; j < len(holeCards); j++ {
			selected := []common.Card{holeCards[i], holeCards[j]}
			selected = append(selected, community...)

			hand, rank := evaluate5CardDraw(selected, 2)
			if rank > bestRank {
				bestHand = hand
				bestRank = rank
			} else if rank == bestRank && compareHands(hand, bestHand) > 0 {
				bestHand = hand
			}
		}
	}

	return bestHand, bestRank
}

func checkStraightFlush(cards []common.Card) []common.Card {
	// Group by suit
	suits := make(map[common.Suit][]common.Card)
	for _, card := range cards {
		suits[card.Suit] = append(suits[card.Suit], card)
	}

	// Check each suit for straight
	for _, suitCards := range suits {
		if len(suitCards) >= 5 {
			if hand := checkStraight(suitCards); hand != nil {
				// Check for royal flush
				ranks := getRanks(hand)
				if contains(ranks, common.Ace) && contains(ranks, common.King) &&
					contains(ranks, common.Queen) && contains(ranks, common.Jack) &&
					contains(ranks, common.Ten) {
					return hand
				}
				return hand
			}
		}
	}

	return nil
}
