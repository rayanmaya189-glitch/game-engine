package poker

import (
	"errors"
	"fmt"
	"sort"

	"github.com/game-engine/card-games/internal/games/common"
)

// HandRank represents the rank of a poker hand
type HandRank int

const (
	HighCard HandRank = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

// GameType represents the type of poker game
type GameType int

const (
	TexasHoldem GameType = iota
	Omaha
)

// Street represents the betting street
type Street int

const (
	PreFlop Street = iota
	Flop
	Turn
	River
	Showdown
)

// Player represents a poker player
type Player struct {
	ID        string        `json:"id"`
	HoleCards []common.Card `json:"hole_cards"`
	Chips     int64         `json:"chips"`
	Bet       int64         `json:"bet"`
	HandRank  HandRank      `json:"hand_rank"`
	BestHand  []common.Card `json:"best_hand"`
	Folded    bool          `json:"folded"`
	AllIn     bool          `json:"all_in"`
}

// Game represents a poker game
type Game struct {
	ID            string
	GameType      GameType
	Shoe          *common.Shoe
	Players       map[string]*Player
	Community     []common.Card
	CurrentStreet Street
	Pot           int64
	SidePots      []int64
	Config        *Config
}

// Config holds poker configuration
type Config struct {
	MinPlayers    int
	MaxPlayers    int
	SmallBlind    int64
	BigBlind      int64
	StartingChips int64
}

// NewGame creates a new poker game
func NewGame(id string, gameType GameType, shoe *common.Shoe, config *Config) *Game {
	return &Game{
		ID:            id,
		GameType:      gameType,
		Shoe:          shoe,
		Players:       make(map[string]*Player),
		Community:     make([]common.Card, 0),
		CurrentStreet: PreFlop,
		Pot:           0,
		SidePots:      make([]int64, 0),
		Config:        config,
	}
}

// AddPlayer adds a player to the game
func (g *Game) AddPlayer(playerID string) error {
	if len(g.Players) >= g.Config.MaxPlayers {
		return errors.New("max players reached")
	}

	player := &Player{
		ID:        playerID,
		HoleCards: make([]common.Card, 0),
		Chips:     g.Config.StartingChips,
		Bet:       0,
		Folded:    false,
		AllIn:     false,
	}

	g.Players[playerID] = player
	return nil
}

// DealHoleCards deals hole cards to players
func (g *Game) DealHoleCards() error {
	// Deal 2 cards for Texas Hold'em, 4 for Omaha
	cardsPerPlayer := 2
	if g.GameType == Omaha {
		cardsPerPlayer = 4
	}

	for range g.Players {
		for i := 0; i < cardsPerPlayer; i++ {
			card, err := g.Shoe.Draw()
			if err != nil {
				return err
			}

			// Find a player who needs a card
			for _, player := range g.Players {
				if len(player.HoleCards) < cardsPerPlayer {
					player.HoleCards = append(player.HoleCards, card)
					break
				}
			}
		}
	}

	return nil
}

// DealCommunity deals community cards
func (g *Game) DealCommunity(count int) error {
	for i := 0; i < count; i++ {
		card, err := g.Shoe.Draw()
		if err != nil {
			return err
		}
		g.Community = append(g.Community, card)
	}

	// Update street
	switch g.CurrentStreet {
	case PreFlop:
		g.CurrentStreet = Flop
	case Flop:
		g.CurrentStreet = Turn
	case Turn:
		g.CurrentStreet = River
	case River:
		g.CurrentStreet = Showdown
	}

	return nil
}

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

// Helper functions

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

// GetState returns the current game state
func (g *Game) GetState() *GameState {
	state := &GameState{
		GameID:        g.ID,
		GameType:      string(g.GameType),
		Community:     g.Community,
		CurrentStreet: string(g.CurrentStreet),
		Pot:           g.Pot,
		Players:       make([]PlayerState, 0),
	}

	for _, player := range g.Players {
		state.Players = append(state.Players, PlayerState{
			ID:        player.ID,
			HoleCards: player.HoleCards,
			Chips:     player.Chips,
			Bet:       player.Bet,
			Folded:    player.Folded,
			AllIn:     player.AllIn,
			HandRank:  string(player.HandRank),
			BestHand:  player.BestHand,
		})
	}

	return state
}

// GameState represents the game state for clients
type GameState struct {
	GameID        string        `json:"game_id"`
	GameType      string        `json:"game_type"`
	Community     []common.Card `json:"community"`
	CurrentStreet string        `json:"current_street"`
	Pot           int64         `json:"pot"`
	Players       []PlayerState `json:"players"`
}

// PlayerState represents a player's state
type PlayerState struct {
	ID        string        `json:"id"`
	HoleCards []common.Card `json:"hole_cards"`
	Chips     int64         `json:"chips"`
	Bet       int64         `json:"bet"`
	Folded    bool          `json:"folded"`
	AllIn     bool          `json:"all_in"`
	HandRank  string        `json:"hand_rank"`
	BestHand  []common.Card `json:"best_hand"`
}

func init() {
	_ = fmt.Sprintf("%s", HighCard)
}
