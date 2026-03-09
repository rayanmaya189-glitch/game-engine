package common

import (
	"errors"
	"fmt"
)

// Suit represents the four suits in a deck
type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

// Rank represents card ranks from Ace to King
type Rank int

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Card represents a playing card
type Card struct {
	Suit Suit `json:"suit"`
	Rank Rank `json:"rank"`
}

// String returns a string representation of the card
func (c Card) String() string {
	rankStr := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suitStr := []string{"♥", "♦", "♣", "♠"}

	if c.Rank == Ten {
		return fmt.Sprintf("10%s", suitStr[c.Suit])
	}
	return fmt.Sprintf("%s%s", rankStr[c.Rank], suitStr[c.Suit])
}

// Value returns the point value of the card (for blackjack)
func (c Card) Value() int {
	if c.Rank == Ace {
		return 11 // Will be adjusted to 1 if needed
	}
	if c.Rank >= Ten {
		return 10
	}
	return int(c.Rank) + 1
}

// IsAce returns true if the card is an Ace
func (c Card) IsAce() bool {
	return c.Rank == Ace
}

// Deck represents a collection of cards
type Deck struct {
	Cards []Card `json:"cards"`
}

// NewDeck creates a new standard 52-card deck
func NewDeck() *Deck {
	cards := make([]Card, 52)
	index := 0
	for suit := Hearts; suit <= Spades; suit++ {
		for rank := Ace; rank <= King; rank++ {
			cards[index] = Card{Suit: suit, Rank: rank}
			index++
		}
	}
	return &Deck{Cards: cards}
}

// NewDeckWithJokers creates a deck with jokers
func NewDeckWithJokers(jokerCount int) *Deck {
	cards := make([]Card, 52+jokerCount)
	index := 0
	for suit := Hearts; suit <= Spades; suit++ {
		for rank := Ace; rank <= King; rank++ {
			cards[index] = Card{Suit: suit, Rank: rank}
			index++
		}
	}
	// Add jokers (represented with special rank)
	for i := 0; i < jokerCount; i++ {
		cards[52+i] = Card{Suit: Suit(i % 4), Rank: Rank(13 + i)}
	}
	return &Deck{Cards: cards}
}

// Shoe represents multiple decks combined
type Shoe struct {
	DeckCount int    `json:"deck_count"`
	Cards     []Card `json:"cards"`
	NextCard  int    `json:"next_card"`
}

// NewShoe creates a new shoe with the specified number of decks
func NewShoe(deckCount int) *Shoe {
	if deckCount < 1 {
		deckCount = 1
	}
	if deckCount > 8 {
		deckCount = 8
	}

	totalCards := 52 * deckCount
	cards := make([]Card, 0, totalCards)

	for i := 0; i < deckCount; i++ {
		deck := NewDeck()
		cards = append(cards, deck.Cards...)
	}

	return &Shoe{
		DeckCount: deckCount,
		Cards:     cards,
		NextCard:  0,
	}
}

// Remaining returns the number of cards remaining in the shoe
func (s *Shoe) Remaining() int {
	return len(s.Cards) - s.NextCard
}

// Draw draws the next card from the shoe
func (s *Shoe) Draw() (Card, error) {
	if s.NextCard >= len(s.Cards) {
		return Card{}, errors.New("shoe is empty")
	}

	card := s.Cards[s.NextCard]
	s.NextCard++
	return card, nil
}

// DrawMultiple draws multiple cards from the shoe
func (s *Shoe) DrawMultiple(count int) ([]Card, error) {
	cards := make([]Card, count)
	for i := 0; i < count; i++ {
		card, err := s.Draw()
		if err != nil {
			return nil, err
		}
		cards[i] = card
	}
	return cards, nil
}

// Reset resets the shoe with new shuffled cards
func (s *Shoe) Reset(deckCount int, shuffler Shuffler) error {
	*s = *NewShoe(deckCount)
	return shuffler.Shuffle(s.Cards)
}

// NeedsReshuffle returns true if the shoe needs to be reshuffled
func (s *Shoe) NeedsReshuffle(threshold int) bool {
	return s.Remaining() < threshold
}

// Hand represents a player's hand of cards
type Hand struct {
	Cards []Card `json:"cards"`
	Bet   int64  `json:"bet"`
}

// NewHand creates a new empty hand
func NewHand() *Hand {
	return &Hand{
		Cards: make([]Card, 0),
		Bet:   0,
	}
}

// Add adds a card to the hand
func (h *Hand) Add(card Card) {
	h.Cards = append(h.Cards, card)
}

// Clear removes all cards from the hand
func (h *Hand) Clear() {
	h.Cards = h.Cards[:0]
}

// Count returns the number of cards in the hand
func (h *Hand) Count() int {
	return len(h.Cards)
}

// Total returns the point total of the hand (for blackjack)
func (h *Hand) Total() int {
	total := 0
	aces := 0

	for _, card := range h.Cards {
		if card.IsAce() {
			aces++
			total += 11
		} else {
			total += card.Value()
		}
	}

	// Convert aces from 11 to 1 as needed
	for aces > 0 && total > 21 {
		total -= 10
		aces--
	}

	return total
}

// IsBusted returns true if the hand is busted (over 21)
func (h *Hand) IsBusted() bool {
	return h.Total() > 21
}

// IsBlackjack returns true if the hand is a natural blackjack (Ace + 10-value)
func (h *Hand) IsBlackjack() bool {
	if len(h.Cards) != 2 {
		return false
	}

	card1, card2 := h.Cards[0], h.Cards[1]

	// Check for Ace + 10-value
	hasAce := card1.IsAce() || card2.IsAce()
	hasTen := card1.Value() == 10 || card2.Value() == 10

	return hasAce && hasTen
}

// IsSoft returns true if the hand contains an Ace counted as 11
func (h *Hand) IsSoft() bool {
	total := 0
	aces := 0

	for _, card := range h.Cards {
		if card.IsAce() {
			aces++
			total += 11
		} else {
			total += card.Value()
		}
	}

	return aces > 0 && total <= 21
}

// CanSplit returns true if the hand can be split
func (h *Hand) CanSplit() bool {
	if len(h.Cards) != 2 {
		return false
	}
	return h.Cards[0].Rank == h.Cards[1].Rank
}

// CardToIndex converts a card to its index in a standard deck
func CardToIndex(card Card) int {
	return int(card.Suit)*13 + int(card.Rank)
}

// IndexToCard converts an index to a card
func IndexToCard(index int) Card {
	return Card{
		Suit: Suit(index / 13),
		Rank: Rank(index % 13),
	}
}
