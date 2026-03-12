package dice

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
)

// Roller interface for dice rolling
type Roller interface {
	Roll(diceCount int) ([]int, error)
}

// RandomRoller uses crypto RNG to roll dice
type RandomRoller struct{}

// NewRandomRoller creates a new random roller
func NewRandomRoller() *RandomRoller {
	return &RandomRoller{}
}

// Roll rolls dice using crypto RNG
func (r *RandomRoller) Roll(diceCount int) ([]int, error) {
	if diceCount <= 0 || diceCount > 10 {
		return nil, ErrInvalidDiceCount
	}

	dice := make([]int, diceCount)
	for i := 0; i < diceCount; i++ {
		// Generate random number between 1 and 6 (inclusive)
		n, err := rand.Int(rand.Reader, big.NewInt(6))
		if err != nil {
			return nil, err
		}
		dice[i] = int(n.Int64()) + 1
	}
	return dice, nil
}

// RollSum rolls dice and returns the sum
func (r *RandomRoller) RollSum(diceCount int) (int, error) {
	dice, err := r.Roll(diceCount)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, d := range dice {
		sum += d
	}
	return sum, nil
}

// ProvablyFairRoller uses seed for provably fair dice rolling
type ProvablyFairRoller struct {
	ServerSeed string
	ClientSeed string
	Nonce      int
}

// NewProvablyFairRoller creates a new provably fair roller
func NewProvablyFairRoller(serverSeed, clientSeed string) *ProvablyFairRoller {
	return &ProvablyFairRoller{
		ServerSeed: serverSeed,
		ClientSeed: clientSeed,
		Nonce:      0,
	}
}

// Roll rolls dice using HMAC-based provably fair algorithm
func (p *ProvablyFairRoller) Roll(diceCount int) ([]int, error) {
	if diceCount <= 0 || diceCount > 10 {
		return nil, ErrInvalidDiceCount
	}

	dice := make([]int, diceCount)
	for i := 0; i < diceCount; i++ {
		hmacData := []byte(p.ServerSeed + p.ClientSeed + fmt.Sprintf("%d%d", p.Nonce, i))
		hash := sha256Hash(hmacData)
		index := hash % 6
		dice[i] = index + 1
	}
	p.Nonce++
	return dice, nil
}

// RollSum rolls dice and returns the sum
func (p *ProvablyFairRoller) RollSum(diceCount int) (int, error) {
	dice, err := p.Roll(diceCount)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, d := range dice {
		sum += d
	}
	return sum, nil
}

func sha256Hash(data []byte) int {
	hash := sha256.Sum256(data)
	result := 0
	for i, b := range hash {
		result = (result * 31) + int(b) + i
	}
	return result % 6
}

var (
	ErrInvalidDiceCount = errors.New("invalid dice count: must be between 1 and 10")
)
