package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

// generateDeck creates a standard 52-card deck
func (s *HybridGameService) generateDeck() []string {
	suits := []string{"H", "D", "C", "S"}
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := make([]string, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, rank+suit)
		}
	}
	return deck
}

// shuffleDeck randomly shuffles a deck using crypto/rand
func shuffleDeck(deck []string) {
	for i := len(deck) - 1; i > 0; i-- {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		deck[i], deck[j.Int64()] = deck[j.Int64()], deck[i]
	}
}

// calculateBaccaratScore calculates baccarat hand score (0-9)
func (s *HybridGameService) calculateBaccaratScore(cards []string) int {
	score := 0
	for _, card := range cards {
		rank := card[:len(card)-1]
		switch rank {
		case "A":
			score += 1
		case "2", "3", "4", "5", "6", "7", "8", "9":
			val := 0
			fmt.Sscanf(rank, "%d", &val)
			score += val
		}
	}
	return score % 10
}

// calculateRoulettePayouts calculates roulette payouts for all bet types
func (s *HybridGameService) calculateRoulettePayouts(result int, bets map[string]float64) map[string]float64 {
	payouts := make(map[string]float64)
	for betType, amount := range bets {
		switch betType {
		case fmt.Sprintf("%d", result):
			payouts[betType] = amount * 36
		case "red", "black":
			if (result > 0 && result%2 == 1) == (betType == "red") {
				payouts[betType] = amount * 2
			}
		case "odd", "even":
			if result > 0 && (result%2 == 1) == (betType == "odd") {
				payouts[betType] = amount * 2
			}
		case "1-18", "19-36":
			if (result >= 1 && result <= 18) == (betType == "1-18") && result > 0 {
				payouts[betType] = amount * 2
			}
		}
	}
	return payouts
}

// calculateBaccaratPayouts calculates baccarat payouts
func (s *HybridGameService) calculateBaccaratPayouts(outcome string, bets map[string]float64) map[string]float64 {
	payouts := make(map[string]float64)
	for betType, amount := range bets {
		if betType == outcome {
			switch outcome {
			case "player_wins":
				payouts[betType] = amount * 2
			case "banker_wins":
				payouts[betType] = amount * 1.95
			case "tie":
				payouts[betType] = amount * 9
			}
		}
	}
	return payouts
}

// calculateDicePayouts calculates dice (Sic Bo) payouts
func (s *HybridGameService) calculateDicePayouts(dice []int, bets map[string]float64) map[string]float64 {
	payouts := make(map[string]float64)
	total := dice[0] + dice[1] + dice[2]

	for betType, amount := range bets {
		switch betType {
		case "big":
			if total >= 11 && total <= 17 {
				payouts[betType] = amount * 2
			}
		case "small":
			if total >= 4 && total <= 10 {
				payouts[betType] = amount * 2
			}
		case "odd":
			if total%2 == 1 {
				payouts[betType] = amount * 2
			}
		case "even":
			if total%2 == 0 {
				payouts[betType] = amount * 2
			}
		case "triple":
			if dice[0] == dice[1] && dice[1] == dice[2] {
				payouts[betType] = amount * 31
			}
		}
	}
	return payouts
}

// generateProof creates a provably fair proof hash
func (s *HybridGameService) generateProof(roundID string, data []string) string {
	h := make([]byte, 32)
	rand.Read(h)
	return hex.EncodeToString(h)
}

// generateRoundID creates a unique round ID
func generateRoundID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return "RND-" + hex.EncodeToString(b)[:12]
}
