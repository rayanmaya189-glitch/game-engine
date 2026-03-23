package craps

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"time"
)

// Game phases
const (
	PhaseComeOut = "come_out"
	PhasePoint   = "point"
	PhaseRoll    = "roll"
)

// Bet types
const (
	BetPassLine    = "pass_line"
	BetDontPass    = "dont_pass"
	BetCome        = "come"
	BetDontCome    = "dont_come"
	BetPlaceWin4   = "place_win_4"
	BetPlaceWin5   = "place_win_5"
	BetPlaceWin6   = "place_win_6"
	BetPlaceWin8   = "place_win_8"
	BetPlaceWin9   = "place_win_9"
	BetPlaceWin10  = "place_win_10"
	BetPlaceLose4  = "place_lose_4"
	BetPlaceLose5  = "place_lose_5"
	BetPlaceLose6  = "place_lose_6"
	BetPlaceLose8  = "place_lose_8"
	BetPlaceLose9  = "place_lose_9"
	BetPlaceLose10 = "place_lose_10"
	BetField       = "field"
	BetBig6        = "big_6"
	BetBig8        = "big_8"
	BetAny7        = "any_7"
	BetAnyCraps    = "any_craps"
	BetHorn2       = "horn_2"
	BetHorn3       = "horn_3"
	BetHorn11      = "horn_11"
	BetHorn12      = "horn_12"
	BetHard4       = "hard_4"
	BetHard6       = "hard_6"
	BetHard8       = "hard_8"
	BetHard10      = "hard_10"
	BetBuy4        = "buy_4"
	BetBuy5        = "buy_5"
	BetBuy6        = "buy_6"
	BetBuy8        = "buy_8"
	BetBuy9        = "buy_9"
	BetBuy10       = "buy_10"
	BetLay4        = "lay_4"
	BetLay5        = "lay_5"
	BetLay6        = "lay_6"
	BetLay8        = "lay_8"
	BetLay9        = "lay_9"
	BetLay10       = "lay_10"
)

// Payout ratios for each bet type
var payoutRatios = map[string]float64{
	BetPassLine:    1.0,
	BetDontPass:    1.0,
	BetCome:        1.0,
	BetDontCome:    1.0,
	BetPlaceWin4:   9.0 / 5.0,  // 9:5
	BetPlaceWin5:   7.0 / 5.0,  // 7:5
	BetPlaceWin6:   7.0 / 6.0,  // 7:6
	BetPlaceWin8:   7.0 / 6.0,  // 7:6
	BetPlaceWin9:   7.0 / 5.0,  // 7:5
	BetPlaceWin10:  9.0 / 5.0,  // 9:5
	BetPlaceLose4:  5.0 / 11.0, // 5:11
	BetPlaceLose5:  5.0 / 8.0,  // 5:8
	BetPlaceLose6:  4.0 / 5.0,  // 4:5
	BetPlaceLose8:  4.0 / 5.0,  // 4:5
	BetPlaceLose9:  5.0 / 8.0,  // 5:8
	BetPlaceLose10: 5.0 / 11.0, // 5:11
	BetField:       1.0,        // 1:1 (2 and 12 pay 2:1)
	BetBig6:        1.0,
	BetBig8:        1.0,
	BetAny7:        4.0,
	BetAnyCraps:    7.0,
	BetHorn2:       27.0 / 4.0, // 27:4
	BetHorn3:       3.0,
	BetHorn11:      3.0,
	BetHorn12:      27.0 / 4.0, // 27:4
	BetHard4:       7.0,
	BetHard6:       7.0 / 6.0, // 7:6
	BetHard8:       7.0 / 6.0, // 7:6
	BetHard10:      7.0,
	BetBuy4:        2.0, // 5% commission
	BetBuy5:        2.0,
	BetBuy6:        2.0,
	BetBuy8:        2.0,
	BetBuy9:        2.0,
	BetBuy10:       2.0,
	BetLay4:        1.0 / 2.0, // 1:2
	BetLay5:        2.0 / 3.0, // 2:3
	BetLay6:        5.0 / 6.0, // 5:6
	BetLay8:        5.0 / 6.0, // 5:6
	BetLay9:        2.0 / 3.0, // 2:3
	BetLay10:       1.0 / 2.0, // 1:2
}

// CrapsNumbers identifies craps numbers (2, 3, 12)
var crapsNumbers = map[int]bool{2: true, 3: true, 12: true}

// PointNumbers identifies point numbers
var pointNumbers = map[int]bool{4: true, 5: true, 6: true, 8: true, 9: true, 10: true}

// Game represents a Craps game
type Game struct {
	ID           string
	Point        int
	Dice         []int
	Phase        string
	Bets         map[string]map[string]int64 // betType -> playerID -> amount
	Payouts      map[string]map[string]int64 // betType -> playerID -> payout
	Roller       Roller
	ProvablyFair bool
	ServerSeed   string
	ClientSeed   string
	Nonce        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Roller interface for dice rolling
type Roller interface {
	Roll() ([]int, error)
}

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

// PlaceBet places a bet for a player
func (g *Game) PlaceBet(betType, playerID string, amount int64) error {
	if amount <= 0 {
		return errors.New("bet amount must be positive")
	}

	// Validate bet type
	if _, ok := payoutRatios[betType]; !ok {
		return errors.New("invalid bet type")
	}

	// Check if bet is allowed in current phase
	if !g.isBetAllowed(betType) {
		return errors.New("bet not allowed in current phase")
	}

	if g.Bets[betType] == nil {
		g.Bets[betType] = make(map[string]int64)
	}
	g.Bets[betType][playerID] += amount

	return nil
}

// isBetAllowed checks if a bet type is allowed in the current phase
func (g *Game) isBetAllowed(betType string) bool {
	switch g.Phase {
	case PhaseComeOut:
		// These bets can be placed during come-out
		switch betType {
		case BetPassLine, BetDontPass, BetField, BetAny7, BetAnyCraps,
			BetHorn2, BetHorn3, BetHorn11, BetHorn12:
			return true
		default:
			return false
		}
	case PhasePoint:
		// All bets allowed during point phase
		return true
	default:
		return false
	}
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

// resolveBets resolves all pending bets
func (g *Game) resolveBets() {
	sum := g.Dice[0] + g.Dice[1]

	for betType, playerBets := range g.Bets {
		if g.Payouts[betType] == nil {
			g.Payouts[betType] = make(map[string]int64)
		}

		for playerID, amount := range playerBets {
			payout := g.calculatePayout(betType, amount, sum)
			g.Payouts[betType][playerID] = payout
		}
	}
}

// calculatePayout calculates the payout for a bet
func (g *Game) calculatePayout(betType string, amount int64, sum int) int64 {
	ratio, ok := payoutRatios[betType]
	if !ok {
		return 0
	}

	// Check if bet wins
	if !g.betWins(betType, sum) {
		return 0
	}

	// Calculate payout (bet amount + winnings)
	return amount + int64(float64(amount)*ratio)
}

// betWins determines if a bet wins based on the roll result
func (g *Game) betWins(betType string, sum int) bool {
	switch betType {
	case BetPassLine:
		if g.Phase == PhaseComeOut {
			// Win on 7 or 11
			return sum == 7 || sum == 11
		} else {
			// Win if point is made
			return sum == g.Point
		}

	case BetDontPass:
		if g.Phase == PhaseComeOut {
			// Lose on 7 or 11, win on 2 or 3, push on 12
			return sum == 2 || sum == 3
		} else {
			// Win if 7 is rolled before point
			return sum == 7
		}

	case BetField:
		// Win on 3, 4, 9, 10, 11; 2 and 12 pay double
		return sum >= 3 && sum <= 12 && sum != 7 && !crapsNumbers[sum]

	case BetBig6:
		// Win if 6 is rolled before 7
		if g.Phase == PhasePoint {
			return sum == 6
		}
		return false

	case BetBig8:
		// Win if 8 is rolled before 7
		if g.Phase == PhasePoint {
			return sum == 8
		}
		return false

	case BetAny7:
		return sum == 7

	case BetAnyCraps:
		return crapsNumbers[sum]

	case BetHorn2, BetHorn3, BetHorn11, BetHorn12:
		target := 0
		switch betType {
		case BetHorn2:
			target = 2
		case BetHorn3:
			target = 3
		case BetHorn11:
			target = 11
		case BetHorn12:
			target = 12
		}
		return sum == target

	case BetHard4:
		// Win if double 2 (4) is rolled before 7
		if g.Phase == PhasePoint {
			return g.Dice[0] == 2 && g.Dice[1] == 2
		}
		return false

	case BetHard6:
		// Win if double 3 (6) is rolled before 7
		if g.Phase == PhasePoint {
			return g.Dice[0] == 3 && g.Dice[1] == 3
		}
		return false

	case BetHard8:
		// Win if double 4 (8) is rolled before 7
		if g.Phase == PhasePoint {
			return g.Dice[0] == 4 && g.Dice[1] == 4
		}
		return false

	case BetHard10:
		// Win if double 5 (10) is rolled before 7
		if g.Phase == PhasePoint {
			return g.Dice[0] == 5 && g.Dice[1] == 5
		}
		return false

	case BetPlaceWin4:
		return g.Phase == PhasePoint && sum == 4

	case BetPlaceWin5:
		return g.Phase == PhasePoint && sum == 5

	case BetPlaceWin6:
		return g.Phase == PhasePoint && sum == 6

	case BetPlaceWin8:
		return g.Phase == PhasePoint && sum == 8

	case BetPlaceWin9:
		return g.Phase == PhasePoint && sum == 9

	case BetPlaceWin10:
		return g.Phase == PhasePoint && sum == 10

	case BetPlaceLose4:
		return g.Phase == PhasePoint && sum == 7

	case BetPlaceLose5:
		return g.Phase == PhasePoint && sum == 7

	case BetPlaceLose6:
		return g.Phase == PhasePoint && sum == 7

	case BetPlaceLose8:
		return g.Phase == PhasePoint && sum == 7

	case BetPlaceLose9:
		return g.Phase == PhasePoint && sum == 7

	case BetPlaceLose10:
		return g.Phase == PhasePoint && sum == 7

	case BetBuy4, BetBuy5, BetBuy6, BetBuy8, BetBuy9, BetBuy10:
		if g.Phase != PhasePoint {
			return false
		}
		target := 0
		switch betType {
		case BetBuy4:
			target = 4
		case BetBuy5:
			target = 5
		case BetBuy6:
			target = 6
		case BetBuy8:
			target = 8
		case BetBuy9:
			target = 9
		case BetBuy10:
			target = 10
		}
		return sum == target

	case BetLay4, BetLay5, BetLay6, BetLay8, BetLay9, BetLay10:
		if g.Phase != PhasePoint {
			return false
		}
		// Lay bets win if 7 is rolled before the target number
		return sum == 7

	default:
		return false
	}
}

// GetPayout retrieves the payout for a specific player and bet type
func (g *Game) GetPayout(betType, playerID string) int64 {
	if payouts, ok := g.Payouts[betType]; ok {
		return payouts[playerID]
	}
	return 0
}

// GetBet retrieves the bet amount for a specific player and bet type
func (g *Game) GetBet(betType, playerID string) int64 {
	if bets, ok := g.Bets[betType]; ok {
		return bets[playerID]
	}
	return 0
}

// GetTotalBets calculates the total bet amount for a player
func (g *Game) GetTotalBets(playerID string) int64 {
	total := int64(0)
	for _, playerBets := range g.Bets {
		total += playerBets[playerID]
	}
	return total
}

// GetTotalPayouts calculates the total payout for a player
func (g *Game) GetTotalPayouts(playerID string) int64 {
	total := int64(0)
	for _, playerPayouts := range g.Payouts {
		total += playerPayouts[playerID]
	}
	return total
}

// GetState returns the game state
func (g *Game) GetState() *GameState {
	return &GameState{
		GameID:         g.ID,
		Dice:           g.Dice,
		Sum:            g.Dice[0] + g.Dice[1],
		Point:          g.Point,
		Phase:          g.Phase,
		Bets:           g.Bets,
		Payouts:        g.Payouts,
		ProvablyFair:   g.ProvablyFair,
		ServerSeedHash: hashSeed(g.ServerSeed),
		ClientSeed:     g.ClientSeed,
		Nonce:          g.Nonce,
		CreatedAt:      g.CreatedAt,
		UpdatedAt:      g.UpdatedAt,
	}
}

// hashSeed creates a hash of the seed for display
func hashSeed(seed string) string {
	if seed == "" {
		return ""
	}
	h := sha256.Sum256([]byte(seed))
	return fmt.Sprintf("%x", h[:8])
}

// GameState represents the game state
type GameState struct {
	GameID         string                      `json:"game_id"`
	Dice           []int                       `json:"dice"`
	Sum            int                         `json:"sum"`
	Point          int                         `json:"point"`
	Phase          string                      `json:"phase"`
	Bets           map[string]map[string]int64 `json:"bets"`
	Payouts        map[string]map[string]int64 `json:"payouts"`
	ProvablyFair   bool                        `json:"provably_fair"`
	ServerSeedHash string                      `json:"server_seed_hash,omitempty"`
	ClientSeed     string                      `json:"client_seed,omitempty"`
	Nonce          int                         `json:"nonce,omitempty"`
	CreatedAt      time.Time                   `json:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at"`
}

// GetAvailableBets returns all bet types available in current phase
func (g *Game) GetAvailableBets() []string {
	if g.Phase == PhaseComeOut {
		return []string{
			BetPassLine, BetDontPass, BetField, BetAny7, BetAnyCraps,
			BetHorn2, BetHorn3, BetHorn11, BetHorn12,
		}
	}
	return []string{
		BetPassLine, BetDontPass, BetField, BetBig6, BetBig8, BetAny7, BetAnyCraps,
		BetHorn2, BetHorn3, BetHorn11, BetHorn12,
		BetHard4, BetHard6, BetHard8, BetHard10,
		BetPlaceWin4, BetPlaceWin5, BetPlaceWin6, BetPlaceWin8, BetPlaceWin9, BetPlaceWin10,
		BetPlaceLose4, BetPlaceLose5, BetPlaceLose6, BetPlaceLose8, BetPlaceLose9, BetPlaceLose10,
		BetBuy4, BetBuy5, BetBuy6, BetBuy8, BetBuy9, BetBuy10,
		BetLay4, BetLay5, BetLay6, BetLay8, BetLay9, BetLay10,
	}
}

// ClearBets clears all bets for a player
func (g *Game) ClearBets(playerID string) {
	for betType, playerBets := range g.Bets {
		delete(playerBets, playerID)
		if len(playerBets) == 0 {
			delete(g.Bets, betType)
		}
	}
}

// ClearAllBets clears all bets
func (g *Game) ClearAllBets() {
	g.Bets = make(map[string]map[string]int64)
}

// GetBetInfo returns information about a specific bet type
func (g *Game) GetBetInfo(betType string) (string, string, error) {
	name, ok := betNames[betType]
	if !ok {
		return "", "", errors.New("invalid bet type")
	}

	ratio := payoutRatios[betType]

	return name, fmt.Sprintf("%.2f:1", ratio), nil
}

// Bet information maps
var betNames = map[string]string{
	BetPassLine:    "Pass Line",
	BetDontPass:    "Don't Pass",
	BetCome:        "Come",
	BetDontCome:    "Don't Come",
	BetPlaceWin4:   "Place 4 Win",
	BetPlaceWin5:   "Place 5 Win",
	BetPlaceWin6:   "Place 6 Win",
	BetPlaceWin8:   "Place 8 Win",
	BetPlaceWin9:   "Place 9 Win",
	BetPlaceWin10:  "Place 10 Win",
	BetPlaceLose4:  "Place 4 Lose",
	BetPlaceLose5:  "Place 5 Lose",
	BetPlaceLose6:  "Place 6 Lose",
	BetPlaceLose8:  "Place 8 Lose",
	BetPlaceLose9:  "Place 9 Lose",
	BetPlaceLose10: "Place 10 Lose",
	BetField:       "Field",
	BetBig6:        "Big 6",
	BetBig8:        "Big 8",
	BetAny7:        "Any 7",
	BetAnyCraps:    "Any Craps",
	BetHorn2:       "Horn 2",
	BetHorn3:       "Horn 3",
	BetHorn11:      "Horn 11",
	BetHorn12:      "Horn 12",
	BetHard4:       "Hard 4",
	BetHard6:       "Hard 6",
	BetHard8:       "Hard 8",
	BetHard10:      "Hard 10",
	BetBuy4:        "Buy 4",
	BetBuy5:        "Buy 5",
	BetBuy6:        "Buy 6",
	BetBuy8:        "Buy 8",
	BetBuy9:        "Buy 9",
	BetBuy10:       "Buy 10",
	BetLay4:        "Lay 4",
	BetLay5:        "Lay 5",
	BetLay6:        "Lay 6",
	BetLay8:        "Lay 8",
	BetLay9:        "Lay 9",
	BetLay10:       "Lay 10",
}

var betDescriptions = map[string]string{
	BetPassLine: "Win on 7 or 11 (come-out) or match point",
	BetDontPass: "Win on 2, 3, or 7 before point",
	BetField:    "Win on 3,4,9,10,11 (2 and 12 pay double)",
	BetBig6:     "Win when 6 is rolled before 7",
	BetBig8:     "Win when 8 is rolled before 7",
	BetAny7:     "Win when any 7 is rolled",
	BetAnyCraps: "Win when 2, 3, or 12 is rolled",
	BetHard4:    "Win on double 2 (hard 4) before 7",
	BetHard6:    "Win on double 3 (hard 6) before 7",
	BetHard8:    "Win on double 4 (hard 8) before 7",
	BetHard10:   "Win on double 5 (hard 10) before 7",
}
