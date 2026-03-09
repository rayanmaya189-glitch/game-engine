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
