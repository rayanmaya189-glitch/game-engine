package three_card_poker

import (
	"fmt"
	"time"

	"github.com/game_engine/card-games/internal/games/common"
)

// Hand types (from lowest to highest)
const (
	HighCard      = "high_card"
	Pair          = "pair"
	Flush         = "flush"
	Straight      = "straight"
	ThreeOfAKind  = "three_of_a_kind"
	StraightFlush = "straight_flush"
)

// Payouts for Pair Plus (based on original bet)
var pairPlusPayouts = map[string]float64{
	HighCard:      0,
	Pair:          1.0,
	Flush:         3.0,
	Straight:      4.0,
	ThreeOfAKind:  30.0,
	StraightFlush: 40.0,
}

// Ante bonuses
var anteBonusPayouts = map[string]float64{
	Straight:      1.0,
	ThreeOfAKind:  4.0,
	StraightFlush: 5.0,
}

// evaluateHand evaluates a 3-card hand and returns its type
func evaluateHand(hand []common.Card) string {
	if len(hand) != 3 {
		return HighCard
	}

	sorted := make([]common.Card, len(hand))
	copy(sorted, hand)

	isFlush := sorted[0].Suit == sorted[1].Suit && sorted[1].Suit == sorted[2].Suit

	isStraight := false
	if sorted[2].Rank-sorted[0].Rank == 2 {
		isStraight = true
	} else if sorted[2].Rank == common.King && sorted[1].Rank == common.Queen && sorted[0].Rank == common.Ace {
		isStraight = true
	}

	ranks := []int{int(sorted[0].Rank), int(sorted[1].Rank), int(sorted[2].Rank)}
	hasPair := (ranks[0] == ranks[1]) || (ranks[1] == ranks[2]) || (ranks[0] == ranks[2])
	hasThreeOfAKind := (ranks[0] == ranks[1]) && (ranks[1] == ranks[2])

	if isStraight && isFlush {
		return StraightFlush
	}
	if hasThreeOfAKind {
		return ThreeOfAKind
	}
	if isStraight {
		return Straight
	}
	if isFlush {
		return Flush
	}
	if hasPair {
		return Pair
	}

	return HighCard
}

// compareHands compares two hands, returns 1 if first wins, -1 if second wins, 0 for tie
func compareHands(hand1, hand2 []common.Card) int {
	type1 := evaluateHand(hand1)
	type2 := evaluateHand(hand2)

	typeRank := map[string]int{
		HighCard:      1,
		Pair:          2,
		Flush:         3,
		Straight:      4,
		ThreeOfAKind:  5,
		StraightFlush: 6,
	}

	if typeRank[type1] > typeRank[type2] {
		return 1
	}
	if typeRank[type1] < typeRank[type2] {
		return -1
	}

	maxRank1 := 0
	maxRank2 := 0
	for _, c := range hand1 {
		if int(c.Rank) > maxRank1 {
			maxRank1 = int(c.Rank)
		}
	}
	for _, c := range hand2 {
		if int(c.Rank) > maxRank2 {
			maxRank2 = int(c.Rank)
		}
	}

	if maxRank1 > maxRank2 {
		return 1
	}
	if maxRank1 < maxRank2 {
		return -1
	}

	return 0
}

func generateGameID() string {
	return fmt.Sprintf("tcp_%d_%d", time.Now().Unix(), time.Now().Nanosecond()%1000)
}
