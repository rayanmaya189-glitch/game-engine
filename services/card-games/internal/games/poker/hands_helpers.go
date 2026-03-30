package poker

import "github.com/game_engine/card-games/internal/games/common"

func getRanks(cards []common.Card) []common.Rank {
	ranks := make([]common.Rank, len(cards))
	for i, card := range cards {
		ranks[i] = card.Rank
	}
	return ranks
}

func getRankCounts(cards []common.Card) map[common.Rank]int {
	counts := make(map[common.Rank]int)
	for _, card := range cards {
		counts[card.Rank]++
	}
	return counts
}

func getUniqueRanks(cards []common.Card) []common.Rank {
	seen := make(map[common.Rank]bool)
	var ranks []common.Rank
	for _, card := range cards {
		if !seen[card.Rank] {
			seen[card.Rank] = true
			ranks = append(ranks, card.Rank)
		}
	}
	return ranks
}

func getCardsByRank(cards []common.Card, rank common.Rank) []common.Card {
	var result []common.Card
	for _, card := range cards {
		if card.Rank == rank {
			result = append(result, card)
		}
	}
	return result
}

func getCardsByRanks(cards []common.Card, ranks []common.Rank) []common.Card {
	var result []common.Card
	for _, rank := range ranks {
		for _, card := range cards {
			if card.Rank == rank && !containsCard(result, card) {
				result = append(result, card)
				break
			}
		}
	}
	return result
}

func getHighestKicker(cards []common.Card, exclude ...common.Rank) common.Card {
	excluded := make(map[common.Rank]bool)
	for _, r := range exclude {
		excluded[r] = true
	}

	var highest common.Card
	highest.Rank = -1

	for _, card := range cards {
		if !excluded[card.Rank] && card.Rank > highest.Rank {
			highest = card
		}
	}

	return highest
}

func getHighestKickers(cards []common.Card, exclude common.Rank, count int) []common.Card {
	var result []common.Card
	excluded := make(map[common.Rank]bool)
	excluded[exclude] = true

	for len(result) < count {
		kicker := getHighestKicker(cards, exclude)
		if kicker.Rank < 0 {
			break
		}
		result = append(result, kicker)
		excluded[kicker.Rank] = true
	}

	return result
}

func contains(ranks []common.Rank, rank common.Rank) bool {
	for _, r := range ranks {
		if r == rank {
			return true
		}
	}
	return false
}

func containsCard(cards []common.Card, card common.Card) bool {
	for _, c := range cards {
		if c.Suit == card.Suit && c.Rank == card.Rank {
			return true
		}
	}
	return false
}

func compareHands(hand1, hand2 []common.Card) int {
	if len(hand1) != len(hand2) {
		if len(hand1) > len(hand2) {
			return 1
		}
		return -1
	}

	for i := range hand1 {
		if hand1[i].Rank > hand2[i].Rank {
			return 1
		} else if hand1[i].Rank < hand2[i].Rank {
			return -1
		}
	}

	return 0
}
