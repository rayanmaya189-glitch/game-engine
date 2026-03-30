package craps

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"time"
)

// RandomRoller uses crypto RNG to roll dice
type RandomRoller struct{}

// NewRandomRoller creates a new random roller
func NewRandomRoller() *RandomRoller {
	return &RandomRoller{}
}

// Roll rolls two dice using crypto RNG
func (r *RandomRoller) Roll() ([]int, error) {
	dice := make([]int, 2)
	for i := 0; i < 2; i++ {
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
func (p *ProvablyFairRoller) Roll() ([]int, error) {
	dice := make([]int, 2)
	for i := 0; i < 2; i++ {
		hmacData := []byte(p.ServerSeed + p.ClientSeed + fmt.Sprintf("%d%d", p.Nonce, i))
		hash := sha256.Sum256(hmacData)
		index := int(hash[0]) % 6
		dice[i] = index + 1
	}
	p.Nonce++
	return dice, nil
}

// NewGame creates a new Craps game
func NewGame(id string) *Game {
	return &Game{
		ID:           id,
		Point:        0,
		Dice:         make([]int, 2),
		Phase:        PhaseComeOut,
		Bets:         make(map[string]map[string]int64),
		Payouts:      make(map[string]map[string]int64),
		Roller:       NewRandomRoller(),
		ProvablyFair: false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// NewProvablyFairGame creates a new provably fair Craps game
func NewProvablyFairGame(id, serverSeed, clientSeed string) *Game {
	g := NewGame(id)
	g.Roller = NewProvablyFairRoller(serverSeed, clientSeed)
	g.ServerSeed = serverSeed
	g.ClientSeed = clientSeed
	g.ProvablyFair = true
	return g
}

// Roll rolls the dice and resolves the game
func (g *Game) Roll() error {
	dice, err := g.Roller.Roll()
	if err != nil {
		return err
	}

	g.Dice = dice
	g.Phase = PhaseRoll
	g.UpdatedAt = time.Now()

	// Resolve all bets
	g.resolveBets()

	// Determine next phase
	g.determineNextPhase()

	return nil
}

// RollWithDice rolls with specific dice values (for testing or provably fair)
func (g *Game) RollWithDice(dice []int) error {
	if len(dice) != 2 {
		return errors.New("dice must contain exactly 2 values")
	}
	for _, d := range dice {
		if d < 1 || d > 6 {
			return errors.New("dice values must be between 1 and 6")
		}
	}

	g.Dice = dice
	g.Phase = PhaseRoll
	g.UpdatedAt = time.Now()

	// Resolve all bets
	g.resolveBets()

	// Determine next phase
	g.determineNextPhase()

	return nil
}

// determineNextPhase determines the next game phase based on roll result
func (g *Game) determineNextPhase() {
	sum := g.Dice[0] + g.Dice[1]

	switch g.Phase {
	case PhaseComeOut:
		switch sum {
		case 7, 11: // Natural - Pass Line wins
			g.Phase = PhaseComeOut
			g.Point = 0
		case 2, 3, 12: // Craps - Don't Pass wins (except 12 pushes)
			if sum == 12 {
				// Push - reset the game
				g.Phase = PhaseComeOut
				g.Point = 0
			} else {
				g.Phase = PhaseComeOut
				g.Point = 0
			}
		default: // Point established
			g.Phase = PhasePoint
			g.Point = sum
		}
	case PhasePoint:
		if sum == 7 {
			// Seven out - Don't Pass wins
			g.Phase = PhaseComeOut
			g.Point = 0
		} else if sum == g.Point {
			// Point made - Pass Line wins
			g.Phase = PhaseComeOut
			g.Point = 0
		}
		// Otherwise stay in Point phase
	}
}
