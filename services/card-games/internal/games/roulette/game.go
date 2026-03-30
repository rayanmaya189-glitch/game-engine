package roulette

import (
	"errors"
	"math/rand"
)

// NewWheel creates a new European roulette wheel
func NewWheel() *Wheel {
	// European roulette: numbers 0-36 (single zero)
	return &Wheel{
		Numbers: []int{0, 32, 15, 19, 4, 21, 2, 25, 17, 34, 6, 27, 13, 36, 11, 30, 8, 23, 10, 5, 24, 16, 33, 1, 20, 14, 31, 9, 22, 18, 29, 7, 28, 12, 35, 3, 26},
	}
}

// NewGame creates a new roulette game
func NewGame(id, tableID string) *Game {
	return &Game{
		ID:        id,
		TableID:   tableID,
		State:     GameStateWaiting,
		Wheel:     NewWheel(),
		Bets:      make([]PlayerBet, 0),
		MinBet:    1,
		MaxBet:    10000,
		MaxPayout: 50000,
		HouseEdge: 2.7, // European roulette
		rng:       rand.New(rand.NewSource(0)),
	}
}

// SetRNG sets the random number generator (for provably fair)
func (g *Game) SetRNG(rng *rand.Rand) {
	g.rng = rng
}

// StartBetting opens betting for a new round
func (g *Game) StartBetting() {
	g.State = GameStateBetting
	g.Bets = make([]PlayerBet, 0)
	g.ResultNumber = -1
}

// CloseBetting closes betting and starts spinning
func (g *Game) CloseBetting() error {
	if g.State != GameStateBetting {
		return errors.New("betting is not open")
	}
	g.State = GameStateSpinning
	return nil
}

// Spin spins the wheel and determines the result
func (g *Game) Spin() (int, error) {
	if g.State != GameStateSpinning {
		return -1, errors.New("cannot spin at this time")
	}

	// Random position on wheel
	stopPos := g.rng.Intn(37)
	result := g.Wheel.Numbers[stopPos]

	g.ResultNumber = result
	g.ResultColor = g.getColor(result)
	g.ResultParity = g.getParity(result)
	g.ResultRange = g.getRange(result)
	g.State = GameStateResult

	return result, nil
}

// GetResult returns the result of the last spin
func (g *Game) GetResult() map[string]interface{} {
	if g.State != GameStateResult {
		return nil
	}

	return map[string]interface{}{
		"number": g.ResultNumber,
		"color":  g.ResultColor,
		"parity": g.ResultParity,
		"range":  g.ResultRange,
		"dozen":  g.getDozen(g.ResultNumber),
		"column": g.getColumn(g.ResultNumber),
	}
}

func (g *Game) getColor(n int) string {
	if n == 0 {
		return "green"
	}
	for _, r := range redNumbers {
		if r == n {
			return "red"
		}
	}
	return "black"
}

func (g *Game) getParity(n int) string {
	if n == 0 {
		return "neither"
	}
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}

func (g *Game) getRange(n int) string {
	if n == 0 {
		return "neither"
	}
	if n <= 18 {
		return "low"
	}
	return "high"
}

func (g *Game) getDozen(n int) string {
	if n == 0 {
		return "zero"
	}
	if n <= 12 {
		return "dozen_1"
	}
	if n <= 24 {
		return "dozen_2"
	}
	return "dozen_3"
}

func (g *Game) getColumn(n int) string {
	if n == 0 {
		return "zero"
	}
	col := (n - 1) % 3
	return map[int]string{0: "column_1", 1: "column_2", 2: "column_3"}[col]
}
