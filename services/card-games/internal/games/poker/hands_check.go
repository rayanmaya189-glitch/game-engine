package poker

import (
	"sort"

	"github.com/game_engine/card-games/internal/games/common"
)

func checkFourOfAKind(cards []common.Card) ([]common.Card, HandRank) {
	ranks := getRankCounts(cards)

	for rank, count := range ranks {
		if count == 4 {
			// Found four of a kind
			hand := getCardsByRank(cards, rank)
			// Add highest kicker
			kicker := getHighestKicker(cards, rank)
			hand = append(hand, kicker)
			return hand, FourOfAKind
		}
	}

	return nil, HighCard
}

func checkFullHouse(cards []common.Card) ([]common.Card, HandRank) {
	ranks := getRankCounts(cards)

	var threeOfAKindRank common.Rank
	var pairRank common.Rank

	for rank, count := range ranks {
		if count >= 3 {
			threeOfAKindRank = rank
		} else if count >= 2 {
			pairRank = rank
		}
	}

	if threeOfAKindRank > 0 && pairRank > 0 {
		hand := getCardsByRank(cards, threeOfAKindRank)
		hand = append(hand, getCardsByRank(cards, pairRank)...)
		return hand, FullHouse
	}

	return nil, HighCard
}

func checkFlush(cards []common.Card) ([]common.Card, HandRank) {
	suits := make(map[common.Suit]int)
	for _, card := range cards {
		suits[card.Suit]++
	}

	for suit, count := range suits {
		if count >= 5 {
			var flushCards []common.Card
			for _, card := range cards {
				if card.Suit == suit {
					flushCards = append(flushCards, card)
				}
			}
			// Sort by rank descending
			sort.Slice(flushCards, func(i, j int) bool {
				return flushCards[i].Rank > flushCards[j].Rank
			})
			return flushCards[:5], Flush
		}
	}

	return nil, HighCard
}

func checkStraight(cards []common.Card) []common.Card {
	ranks := getUniqueRanks(cards)
	if len(ranks) < 5 {
		return nil
	}

	// Sort ranks ascending
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i] < ranks[j]
	})

	// Check for straight (including wheel A-2-3-4-5)
	consecutive := 1
	straightStart := 0

	for i := 1; i < len(ranks); i++ {
		if ranks[i] == ranks[i-1]+1 {
			consecutive++
		} else if ranks[i] == ranks[i-1] {
			continue
		} else {
			consecutive = 1
			straightStart = i
		}

		if consecutive == 5 {
			// Found straight
			return getCardsByRanks(cards, ranks[straightStart:straightStart+5])
		}
	}

	// Check for wheel (A-2-3-4-5)
	if len(ranks) >= 5 {
		wheel := []common.Rank{common.Ace, common.Two, common.Three, common.Four, common.Five}
		matches := 0
		for _, r := range wheel {
			for _, cr := range ranks {
				if r == cr {
					matches++
					break
				}
			}
		}
		if matches == 5 {
			return getCardsByRanks(cards, wheel)
		}
	}

	return nil
}

func checkThreeOfAKind(cards []common.Card) ([]common.Card, HandRank) {
	ranks := getRankCounts(cards)

	var threeRank common.Rank
	for rank, count := range ranks {
		if count >= 3 {
			threeRank = rank
			break
		}
	}

	if threeRank > 0 {
		hand := getCardsByRank(cards, threeRank)
		// Add two highest kickers
		kickers := getHighestKickers(cards, threeRank, 2)
		hand = append(hand, kickers...)
		return hand, ThreeOfAKind
	}

	return nil, HighCard
}

func checkTwoPair(cards []common.Card) ([]common.Card, HandRank) {
	ranks := getRankCounts(cards)

	var pairs []common.Rank
	for rank, count := range ranks {
		if count >= 2 {
			pairs = append(pairs, rank)
		}
	}

	if len(pairs) >= 2 {
		// Sort pairs by rank descending
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i] > pairs[j]
		})

		hand := getCardsByRank(cards, pairs[0])
		hand = append(hand, getCardsByRank(cards, pairs[1])...)
		// Add highest kicker
		kicker := getHighestKicker(cards, pairs[0], pairs[1])
		hand = append(hand, kicker)

		return hand, TwoPair
	}

	return nil, HighCard
}

func checkPair(cards []common.Card) ([]common.Card, HandRank) {
	ranks := getRankCounts(cards)

	var pairRank common.Rank
	for rank, count := range ranks {
		if count >= 2 {
			pairRank = rank
			break
		}
	}

	if pairRank > 0 {
		hand := getCardsByRank(cards, pairRank)
		// Add three highest kickers
		kickers := getHighestKickers(cards, pairRank, 3)
		hand = append(hand, kickers...)
		return hand, Pair
	}

	return nil, HighCard
}

func getHighCard(cards []common.Card) []common.Card {
	sorted := make([]common.Card, len(cards))
	copy(sorted, cards)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Rank > sorted[j].Rank
	})

	return sorted[:5]
}
