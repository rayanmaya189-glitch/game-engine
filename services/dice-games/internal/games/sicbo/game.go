package sicbo

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"time"
)

// Game represents a Sic Bo game
type Game struct {
	ID           string
	Dice         []int
	Bets         map[string]map[string]int64 // betType -> playerID -> amount
	Payouts      map[string]map[string]int64 // betType -> playerID -> payout
	Roller       Roller
	ProvablyFair bool
	ServerSeed   string
	ClientSeed   string
	Nonce        int
	MinBet       int64
	MaxBet       int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
		return nil, errors.New("invalid dice count")
	}

	dice := make([]int, diceCount)
	for i := 0; i < diceCount; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(6))
		if err != nil {
			return nil, err
		}
		dice[i] = int(n.Int64()) + 1
	}
	return dice, nil
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
		return nil, errors.New("invalid dice count")
	}

	dice := make([]int, diceCount)
	for i := 0; i < diceCount; i++ {
		hmacData := []byte(p.ServerSeed + p.ClientSeed + fmt.Sprintf("%d%d", p.Nonce, i))
		hash := sha256.Sum256(hmacData)
		index := int(hash[0]) % 6
		dice[i] = index + 1
	}
	p.Nonce++
	return dice, nil
}

// NewGame creates a new Sic Bo game
func NewGame(id string, config *Config) *Game {
	if config == nil {
		config = &Config{
			MinBet:    1,
			MaxBet:    1000,
			DiceCount: 3,
		}
	}

	return &Game{
		ID:           id,
		Dice:         make([]int, config.DiceCount),
		Bets:         make(map[string]map[string]int64),
		Payouts:      make(map[string]map[string]int64),
		Roller:       NewRandomRoller(),
		ProvablyFair: false,
		MinBet:       config.MinBet,
		MaxBet:       config.MaxBet,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// NewProvablyFairGame creates a new provably fair Sic Bo game
func NewProvablyFairGame(id, serverSeed, clientSeed string, config *Config) *Game {
	g := NewGame(id, config)
	g.Roller = NewProvablyFairRoller(serverSeed, clientSeed)
	g.ServerSeed = serverSeed
	g.ClientSeed = clientSeed
	g.ProvablyFair = true
	return g
}

// Roll rolls the dice and resolves the game
func (g *Game) Roll() error {
	diceCount := len(g.Dice)
	if diceCount == 0 {
		diceCount = 3
	}

	dice, err := g.Roller.Roll(diceCount)
	if err != nil {
		return err
	}

	g.Dice = dice
	g.UpdatedAt = time.Now()

	// Resolve all bets
	g.resolveBets()

	return nil
}

// RollWithDice rolls with specific dice values (for testing or provably fair)
func (g *Game) RollWithDice(dice []int) error {
	if len(dice) != 3 {
		return errors.New("dice must contain exactly 3 values")
	}
	for _, d := range dice {
		if d < 1 || d > 6 {
			return errors.New("dice values must be between 1 and 6")
		}
	}

	g.Dice = dice
	g.UpdatedAt = time.Now()

	// Resolve all bets
	g.resolveBets()

	return nil
}

// hashSeed creates a hash of the seed for display
func hashSeed(seed string) string {
	if seed == "" {
		return ""
	}
	h := sha256.Sum256([]byte(seed))
	return fmt.Sprintf("%x", h[:8])
}
