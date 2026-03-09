package common

import (
	"errors"
)

// Shuffler interface for card shuffling
type Shuffler interface {
	Shuffle(cards []Card) error
}

// FisherYatesShuffler implements Fisher-Yates shuffle algorithm
type FisherYatesShuffler struct {
	RNG RandomNumberGenerator
}

// RandomNumberGenerator interface for generating random numbers
type RandomNumberGenerator interface {
	Intn(n int) (int, error)
}

// NewFisherYatesShuffler creates a new Fisher-Yates shuffler
func NewFisherYatesShuffler(rng RandomNumberGenerator) *FisherYatesShuffler {
	return &FisherYatesShuffler{RNG: rng}
}

// Shuffle shuffles the cards using Fisher-Yates algorithm
func (s *FisherYatesShuffler) Shuffle(cards []Card) error {
	if s.RNG == nil {
		return errors.New("RNG is nil")
	}

	n := len(cards)
	if n <= 1 {
		return nil
	}

	for i := n - 1; i > 0; i-- {
		j, err := s.RNG.Intn(i + 1)
		if err != nil {
			return err
		}
		cards[i], cards[j] = cards[j], cards[i]
	}

	return nil
}

// Cut shuffles the deck by making a random cut
func (s *FisherYatesShuffler) Cut(cards []Card) error {
	if s.RNG == nil {
		return errors.New("RNG is nil")
	}

	n := len(cards)
	if n <= 1 {
		return nil
	}

	// Cut at a random position (between 1/4 and 3/4 of the deck)
	minCut := n / 4
	maxCut := 3 * n / 4

	cutPos, err := s.RNG.Intn(maxCut - minCut + 1)
	if err != nil {
		return err
	}
	cutPos += minCut

	// Perform the cut
	cards = append(cards[cutPos:], cards[:cutPos]...)

	return nil
}

// RiffleShuffler simulates a riffle shuffle
type RiffleShuffler struct {
	RNG RandomNumberGenerator
}

// NewRiffleShuffler creates a new riffle shuffler
func NewRiffleShuffler(rng RandomNumberGenerator) *RiffleShuffler {
	return &RiffleShuffler{RNG: rng}
}

// Shuffle shuffles the cards using a riffle shuffle simulation
func (s *RiffleShuffler) Shuffle(cards []Card) error {
	if s.RNG == nil {
		return errors.New("RNG is nil")
	}

	n := len(cards)
	if n <= 1 {
		return nil
	}

	// Split into two halves
	mid := n / 2
	left := cards[:mid]
	right := cards[mid:]

	result := make([]Card, 0, n)

	// Riffle together
	for len(left) > 0 || len(right) > 0 {
		// Randomly choose which half to take from
		chooseLeft, err := s.RNG.Intn(2)
		if err != nil {
			return err
		}

		if chooseLeft == 0 && len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			right = right[1:]
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		}
	}

	copy(cards, result)
	return nil
}

// SmartShuffler combines multiple shuffle methods for better randomization
type SmartShuffler struct {
	RNG RandomNumberGenerator
}

// NewSmartShuffler creates a new smart shuffler
func NewSmartShuffler(rng RandomNumberGenerator) *SmartShuffler {
	return &SmartShuffler{RNG: rng}
}

// Shuffle performs a thorough shuffle using multiple techniques
func (s *SmartShuffler) Shuffle(cards []Card) error {
	if s.RNG == nil {
		return errors.New("RNG is nil")
	}

	fisherYates := NewFisherYatesShuffler(s.RNG)
	riffle := NewRiffleShuffler(s.RNG)

	// Perform multiple riffle shuffles
	for i := 0; i < 3; i++ {
		if err := riffle.Shuffle(cards); err != nil {
			return err
		}
	}

	// Perform a cut
	if err := fisherYates.Cut(cards); err != nil {
		return err
	}

	// Final Fisher-Yates for perfect randomness
	if err := fisherYates.Shuffle(cards); err != nil {
		return err
	}

	return nil
}
