package rng

import "errors"

// GenerateCardDeck generates a shuffled card deck
// Returns cards as []int where 0-12 are hearts, 13-25 are diamonds, etc.
func (r *CryptoRNG) GenerateCardDeck(deckCount int) ([]int, error) {
	if deckCount < 1 || deckCount > 8 {
		return nil, errors.New("deck count must be between 1 and 8")
	}

	cards := make([]int, 52*deckCount)
	for d := 0; d < deckCount; d++ {
		for c := 0; c < 52; c++ {
			cards[d*52+c] = d*52 + c
		}
	}

	if err := r.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	}); err != nil {
		return nil, err
	}

	return cards, nil
}

// GenerateDiceRolls generates random dice rolls
// Returns dice values as []int where each die is 1-6
func (r *CryptoRNG) GenerateDiceRolls(diceCount int) ([]int, error) {
	if diceCount < 1 || diceCount > 6 {
		return nil, errors.New("dice count must be between 1 and 6")
	}

	rolls := make([]int, diceCount)
	for i := 0; i < diceCount; i++ {
		val, err := r.Intn(6)
		if err != nil {
			return nil, err
		}
		rolls[i] = val + 1
	}

	return rolls, nil
}

// GenerateSlotReels generates slot reel positions
// Returns positions for each reel as []int
func (r *CryptoRNG) GenerateSlotReels(reelCount, symbolCount int) ([]int, error) {
	if reelCount < 1 || reelCount > 12 {
		return nil, errors.New("reel count must be between 1 and 12")
	}
	if symbolCount < 3 {
		return nil, errors.New("symbol count must be at least 3")
	}

	positions := make([]int, reelCount)
	for i := 0; i < reelCount; i++ {
		val, err := r.Intn(symbolCount)
		if err != nil {
			return nil, err
		}
		positions[i] = val
	}

	return positions, nil
}
