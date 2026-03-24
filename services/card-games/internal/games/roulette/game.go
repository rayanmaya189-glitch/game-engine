package roulette

import (
	"errors"
	"math/rand"
)

// BetType represents types of roulette bets
type BetType string

const (
	// Inside bets
	BetTypeStraight BetType = "straight" // Single number
	BetTypeSplit    BetType = "split"    // Two adjacent numbers
	BetTypeStreet   BetType = "street"   // Three numbers in a row
	BetTypeCorner   BetType = "corner"   // Four numbers
	BetTypeLine     BetType = "line"     // Six numbers
	BetTypeTrio     BetType = "trio"     // 0,1,2 or 00,2,3
	BetTypeBasket   BetType = "basket"   // 0,00,1,2,3

	// Outside bets
	BetTypeRed     BetType = "red"      // Red numbers
	BetTypeBlack   BetType = "black"    // Black numbers
	BetTypeOdd     BetType = "odd"      // Odd numbers
	BetTypeEven    BetType = "even"     // Even numbers
	BetTypeLow     BetType = "low"      // 1-18
	BetTypeHigh    BetType = "high"     // 19-36
	BetTypeDozen1  BetType = "dozen_1"  // 1-12
	BetTypeDozen2  BetType = "dozen_2"  // 13-24
	BetTypeDozen3  BetType = "dozen_3"  // 25-36
	BetTypeColumn1 BetType = "column_1" // 1,4,7,...,34
	BetTypeColumn2 BetType = "column_2" // 2,5,8,...,35
	BetTypeColumn3 BetType = "column_3" // 3,6,9,...,36
)

// Roulette number layout (European single-zero)
var (
	redNumbers   = []int{1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36}
	blackNumbers = []int{2, 4, 6, 8, 10, 11, 13, 15, 17, 20, 22, 24, 26, 28, 29, 31, 33, 35}

	// Bet type to numbers mapping
	betNumbers = map[BetType][]int{
		// Dozens
		BetTypeDozen1: {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		BetTypeDozen2: {13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24},
		BetTypeDozen3: {25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36},
		// Columns
		BetTypeColumn1: {1, 4, 7, 10, 13, 16, 19, 22, 25, 28, 31, 34},
		BetTypeColumn2: {2, 5, 8, 11, 14, 17, 20, 23, 26, 29, 32, 35},
		BetTypeColumn3: {3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36},
		// Low/High
		BetTypeLow:  {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18},
		BetTypeHigh: {19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36},
		// Odd/Even
		BetTypeOdd:  {1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35},
		BetTypeEven: {2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36},
		// Red/Black
		BetTypeRed:   redNumbers,
		BetTypeBlack: blackNumbers,
	}
)

// Payout ratios for each bet type
var payouts = map[BetType]int{
	BetTypeStraight: 35,
	BetTypeSplit:    17,
	BetTypeStreet:   11,
	BetTypeCorner:   8,
	BetTypeLine:     5,
	BetTypeTrio:     11,
	BetTypeBasket:   6,
	BetTypeRed:      1,
	BetTypeBlack:    1,
	BetTypeOdd:      1,
	BetTypeEven:     1,
	BetTypeLow:      1,
	BetTypeHigh:     1,
	BetTypeDozen1:   2,
	BetTypeDozen2:   2,
	BetTypeDozen3:   2,
	BetTypeColumn1:  2,
	BetTypeColumn2:  2,
	BetTypeColumn3:  2,
}

// GameState represents the state of a roulette game
type GameState string

const (
	GameStateWaiting  GameState = "waiting"  // Waiting for bets
	GameStateBetting  GameState = "betting"  // Betting open
	GameStateSpinning GameState = "spinning" // Wheel spinning
	GameStateResult   GameState = "result"   // Result announced
)

// PlayerBet represents a player's bet
type PlayerBet struct {
	PlayerID string
	BetType  BetType
	Numbers  []int // For straight, split, street, corner, line, trio, basket
	Amount   int64
}

// Game represents a roulette game
type Game struct {
	ID           string
	TableID      string
	State        GameState
	Wheel        *Wheel
	Bets         []PlayerBet
	ResultNumber int
	ResultColor  string
	ResultParity string
	ResultRange  string
	MinBet       int64
	MaxBet       int64
	MaxPayout    int64
	HouseEdge    float64
	rng          *rand.Rand
}

// Wheel represents the roulette wheel
type Wheel struct {
	Numbers []int
}

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

// PlaceBet places a bet
func (g *Game) PlaceBet(playerID string, betType BetType, numbers []int, amount int64) error {
	if g.State != GameStateBetting {
		return errors.New("betting is not open")
	}

	if amount < g.MinBet {
		return errors.New("bet amount is below minimum")
	}

	if amount > g.MaxBet {
		return errors.New("bet amount exceeds maximum")
	}

	// Validate bet
	if err := g.validateBet(betType, numbers); err != nil {
		return err
	}

	// Check total payout
	totalPayout := g.calculateTotalPayout()
	if totalPayout+int64(payouts[betType])*amount > g.MaxPayout {
		return errors.New("payout would exceed maximum")
	}

	bet := PlayerBet{
		PlayerID: playerID,
		BetType:  betType,
		Numbers:  numbers,
		Amount:   amount,
	}
	g.Bets = append(g.Bets, bet)

	return nil
}

// validateBet validates a bet
func (g *Game) validateBet(betType BetType, numbers []int) error {
	switch betType {
	case BetTypeStraight:
		if len(numbers) != 1 {
			return errors.New("straight bet requires 1 number")
		}
		if numbers[0] < 0 || numbers[0] > 36 {
			return errors.New("invalid number for straight bet")
		}

	case BetTypeSplit:
		if len(numbers) != 2 {
			return errors.New("split bet requires 2 numbers")
		}
		if !g.isAdjacent(numbers[0], numbers[1]) {
			return errors.New("numbers are not adjacent")
		}

	case BetTypeStreet:
		if len(numbers) != 3 {
			return errors.New("street bet requires 3 numbers")
		}
		if !g.isStreet(numbers) {
			return errors.New("invalid street")
		}

	case BetTypeCorner:
		if len(numbers) != 4 {
			return errors.New("corner bet requires 4 numbers")
		}
		if !g.isCorner(numbers) {
			return errors.New("invalid corner")
		}

	case BetTypeLine:
		if len(numbers) != 6 {
			return errors.New("line bet requires 6 numbers")
		}
		if !g.isLine(numbers) {
			return errors.New("invalid line")
		}

	case BetTypeTrio:
		if len(numbers) != 3 {
			return errors.New("trio bet requires 3 numbers")
		}
		validTrios := [][]int{{0, 1, 2}, {0, 2, 3}, {00, 2, 3}}
		valid := false
		for _, t := range validTrios {
			if sliceEqual(numbers, t) {
				valid = true
				break
			}
		}
		if !valid {
			return errors.New("invalid trio")
		}

	case BetTypeBasket:
		if len(numbers) != 5 {
			return errors.New("basket bet requires 5 numbers")
		}
		validBaskets := [][]int{{0, 1, 2, 3, 4}, {0, 00, 1, 2, 3}}
		valid := false
		for _, b := range validBaskets {
			if sliceEqual(numbers, b) {
				valid = true
				break
			}
		}
		if !valid {
			return errors.New("invalid basket")
		}

	case BetTypeRed, BetTypeBlack, BetTypeOdd, BetTypeEven, BetTypeLow, BetTypeHigh:
		// No numbers needed for these outside bets

	case BetTypeDozen1, BetTypeDozen2, BetTypeDozen3:
		// No numbers needed

	case BetTypeColumn1, BetTypeColumn2, BetTypeColumn3:
		// No numbers needed

	default:
		return errors.New("unknown bet type")
	}

	return nil
}

// isAdjacent checks if two numbers are adjacent on the table
func (g *Game) isAdjacent(a, b int) bool {
	if a == 0 || b == 0 {
		// 0 adjacent to 1, 2, 3
		adjacents := map[int][]int{0: {1, 2, 3}, 1: {0, 2, 4}, 2: {0, 1, 3, 5}, 3: {0, 2, 6}}
		for _, adj := range adjacents[a] {
			if b == adj {
				return true
			}
		}
		return false
	}

	rowA := (a - 1) / 3
	colA := (a - 1) % 3
	rowB := (b - 1) / 3
	colB := (b - 1) % 3

	// Same row and adjacent columns
	if rowA == rowB && abs(colA-colB) == 1 {
		return true
	}

	// Adjacent rows (for split at end of row)
	if abs(rowA-rowB) == 1 && ((colA == 0 && colB == 2) || (colA == 2 && colB == 0)) {
		return true
	}

	return false
}

// isStreet checks if numbers form a valid street (3 in a row)
func (g *Game) isStreet(numbers []int) bool {
	if len(numbers) != 3 {
		return false
	}
	// Sort numbers
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	// Valid streets: 1-3, 4-6, ..., 34-36
	first := numbers[0]
	if first < 1 || first > 34 {
		return false
	}
	if first%3 == 1 && numbers[1] == first+1 && numbers[2] == first+2 {
		return true
	}

	// Special case for 0, 1, 2 and 0, 2, 3
	if (numbers[0] == 0 && numbers[1] == 1 && numbers[2] == 2) ||
		(numbers[0] == 0 && numbers[1] == 2 && numbers[2] == 3) {
		return true
	}

	return false
}

// isCorner checks if numbers form a valid corner (4 numbers)
func (g *Game) isCorner(numbers []int) bool {
	if len(numbers) != 4 {
		return false
	}

	// Sort
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	// Valid corners are at column boundaries (1-2, 2-3)
	// Check all 4 numbers are in adjacent columns of adjacent rows
	valid := false
	for i := 1; i <= 11; i++ {
		// Corner at intersection of row i and i+1
		corner := []int{i, i + 1, i + 3, i + 4}
		if sliceEqual(numbers, corner) {
			valid = true
			break
		}
	}

	// Also check corners with 0
	if contains(numbers, 0) {
		valid = true // 0 can be part of corners
	}

	return valid
}

// isLine checks if numbers form a valid line (6 numbers, 2 rows)
func (g *Game) isLine(numbers []int) bool {
	if len(numbers) != 6 {
		return false
	}

	// Sort
	for i := 0; i < 6; i++ {
		for j := i + 1; j < 6; j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	// Valid lines: 1-6, 4-9, ..., 31-36
	for start := 1; start <= 31; start += 3 {
		line := []int{start, start + 1, start + 2, start + 3, start + 4, start + 5}
		if sliceEqual(numbers, line) {
			return true
		}
	}

	return false
}

// getColor returns the color of a number
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

// getParity returns odd/even
func (g *Game) getParity(n int) string {
	if n == 0 {
		return "neither"
	}
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}

// getRange returns low/high
func (g *Game) getRange(n int) string {
	if n == 0 {
		return "neither"
	}
	if n <= 18 {
		return "low"
	}
	return "high"
}

// calculateTotalPayout calculates total potential payout
func (g *Game) calculateTotalPayout() int64 {
	var total int64
	for _, bet := range g.Bets {
		total += bet.Amount * int64(payouts[bet.BetType])
	}
	return total
}

// Settle calculates winnings for all bets
func (g *Game) Settle() map[string]int64 {
	results := make(map[string]int64)

	for _, bet := range g.Bets {
		win := g.calculateWinnings(bet)
		if win > 0 {
			results[bet.PlayerID] += win
		}
	}

	return results
}

// calculateWinnings calculates winnings for a single bet
func (g *Game) calculateWinnings(bet PlayerBet) int64 {
	if g.ResultNumber < 0 {
		return 0
	}

	win := false
	resultNum := g.ResultNumber

	switch bet.BetType {
	case BetTypeStraight:
		win = contains(bet.Numbers, resultNum)

	case BetTypeSplit:
		win = contains(bet.Numbers, resultNum)

	case BetTypeStreet:
		win = contains(bet.Numbers, resultNum)

	case BetTypeCorner:
		win = contains(bet.Numbers, resultNum)

	case BetTypeLine:
		win = contains(bet.Numbers, resultNum)

	case BetTypeTrio:
		win = contains(bet.Numbers, resultNum)

	case BetTypeBasket:
		win = contains(bet.Numbers, resultNum)

	case BetTypeRed:
		win = g.ResultColor == "red"

	case BetTypeBlack:
		win = g.ResultColor == "black"

	case BetTypeOdd:
		win = g.ResultParity == "odd"

	case BetTypeEven:
		win = g.ResultParity == "even"

	case BetTypeLow:
		win = g.ResultRange == "low"

	case BetTypeHigh:
		win = g.ResultRange == "high"

	case BetTypeDozen1:
		win = resultNum >= 1 && resultNum <= 12

	case BetTypeDozen2:
		win = resultNum >= 13 && resultNum <= 24

	case BetTypeDozen3:
		win = resultNum >= 25 && resultNum <= 36

	case BetTypeColumn1:
		win = resultNum > 0 && (resultNum-1)%3 == 0

	case BetTypeColumn2:
		win = resultNum > 0 && (resultNum-2)%3 == 0

	case BetTypeColumn3:
		win = resultNum > 0 && resultNum%3 == 0
	}

	if win {
		// Return total payout including original bet
		return bet.Amount * int64(payouts[bet.BetType]+1)
	}

	return 0
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

// Helper functions
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func contains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
