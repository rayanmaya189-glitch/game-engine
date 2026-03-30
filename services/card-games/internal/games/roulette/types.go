package roulette

import "math/rand"

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
