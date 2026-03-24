package three_card_poker

import (
	"errors"
	"fmt"
	"sort"
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
	Pair:          1.0,  // 1:1
	Flush:         3.0,  // 3:1
	Straight:      4.0,  // 4:1
	ThreeOfAKind:  30.0, // 30:1
	StraightFlush: 40.0, // 40:1 (with ante bonus)
}

// Ante bonuses
var anteBonusPayouts = map[string]float64{
	Straight:      1.0, // 1:1
	ThreeOfAKind:  4.0, // 4:1
	StraightFlush: 5.0, // 5:1
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

	// Deal 3 cards to player
	for i := 0; i < 3; i++ {
		card, err := g.shoe.Draw()
		if err != nil {
			return err
		}
		g.playerHand = append(g.playerHand, card)
	}

	// Deal 3 cards to dealer
	for i := 0; i < 3; i++ {
		card, err := g.shoe.Draw()
		if err != nil {
			return err
		}
		g.dealerHand = append(g.dealerHand, card)
	}

	// Evaluate player's hand
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

	// Player folds - loses ante and pair plus
	return "dealer", 0, nil
}

// Reveal reveals dealer's cards and settles
func (g *ThreeCardPokerGame) Reveal() (string, int64, error) {
	if len(g.playerHand) != 3 || len(g.dealerHand) != 3 {
		return "", 0, errors.New("cards not dealt")
	}

	// Evaluate dealer's hand
	g.DealerHandType = evaluateHand(g.dealerHand)

	// Check if dealer qualifies (Queen high or better)
	dealerQualifies := compareHands(g.dealerHand, g.playerHand) >= 0

	var totalPayout int64
	win := false

	if !dealerQualifies {
		// Dealer doesn't qualify - ante pays 1:1, play bet is push
		g.Winner = "player"
		g.Result = "resolved"
		totalPayout = g.anteBet * 2 // Ante returned 1:1
		win = true
	} else {
		// Compare hands
		result := compareHands(g.dealerHand, g.playerHand)
		if result > 0 {
			g.Winner = "dealer"
			g.Result = "resolved"
		} else if result < 0 {
			g.Winner = "player"
			g.Result = "resolved"
			// Play bet pays 1:1
			totalPayout = g.playBet * 2
			win = true
		} else {
			g.Winner = "push"
			g.Result = "resolved"
			// Both bets push
			totalPayout = g.anteBet + g.playBet
		}
	}

	// Add ante bonus if player has straight or better
	if win && (g.PlayerHandType == Straight || g.PlayerHandType == ThreeOfAKind || g.PlayerHandType == StraightFlush) {
		anteBonus := float64(g.anteBet) * anteBonusPayouts[g.PlayerHandType]
		totalPayout += int64(anteBonus)
	}

	// Add pair plus payout regardless of main game outcome
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

// evaluateHand evaluates a 3-card hand and returns its type
func evaluateHand(hand []common.Card) string {
	if len(hand) != 3 {
		return HighCard
	}

	// Sort by rank
	sorted := make([]common.Card, len(hand))
	copy(sorted, hand)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Rank < sorted[j].Rank
	})

	// Check flush
	isFlush := sorted[0].Suit == sorted[1].Suit && sorted[1].Suit == sorted[2].Suit

	// Check straight
	isStraight := false
	if sorted[2].Rank-sorted[0].Rank == 2 {
		isStraight = true
	} else if sorted[2].Rank == common.King && sorted[1].Rank == common.Queen && sorted[0].Rank == common.Ace {
		// A-Q-K is a straight
		isStraight = true
	}

	// Check pairs and three of a kind
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

	// Higher hand type wins
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

	// Same hand type - compare by highest card
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
