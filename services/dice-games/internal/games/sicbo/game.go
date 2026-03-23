package sicbo

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"time"
)

// Bet types
const (
	BetSmall          = "small"           // Total 4-10 (excluding triples)
	BetBig            = "big"             // Total 11-17 (excluding triples)
	BetSpecificTriple = "specific_triple" // All three dice same specific number
	BetAnyTriple      = "any_triple"      // All three dice same (any number)
	BetSpecificDouble = "specific_double" // Two dice show same specific number
	BetFourNumber     = "four_number"     // Bet on 4 specific numbers
	BetThreeNumber    = "three_number"    // Three dice total specific combination
	BetTwoNumber      = "two_number"      // Two dice specific combination
	BetSingle         = "single"          // One specific number appears
)

// Payout ratios for each bet type
var payoutRatios = map[string]float64{
	BetSmall:          1.0,   // 1:1
	BetBig:            1.0,   // 1:1
	BetSpecificTriple: 150.0, // 150:1
	BetAnyTriple:      24.0,  // 24:1
	BetSpecificDouble: 8.0,   // 8:1
	BetFourNumber:     7.0,   // 7:1
	BetThreeNumber:    50.0,  // 50:1
	BetTwoNumber:      5.0,   // 5:1
	BetSingle:         1.0,   // 1:1 (varies by count)
}

// Single bet multipliers by count
var singleBetMultipliers = map[int]int64{
	1: 1,
	2: 2,
	3: 3,
}

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

// Config holds Sic Bo configuration
type Config struct {
	MinBet    int64
	MaxBet    int64
	DiceCount int
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

// PlaceBet places a bet for a player
func (g *Game) PlaceBet(betType, playerID string, amount int64) error {
	if amount <= 0 {
		return errors.New("bet amount must be positive")
	}

	if amount < g.MinBet || amount > g.MaxBet {
		return fmt.Errorf("bet amount must be between %d and %d", g.MinBet, g.MaxBet)
	}

	// Validate bet type
	if _, ok := payoutRatios[betType]; !ok {
		return errors.New("invalid bet type")
	}

	if g.Bets[betType] == nil {
		g.Bets[betType] = make(map[string]int64)
	}
	g.Bets[betType][playerID] += amount

	return nil
}

// PlaceBetWithNumber places a bet that requires a specific number (for single, double, etc.)
func (g *Game) PlaceBetWithNumber(betType, playerID string, number int, amount int64) error {
	if number < 1 || number > 6 {
		return errors.New("number must be between 1 and 6")
	}

	betKey := fmt.Sprintf("%s_%d", betType, number)
	return g.PlaceBet(betKey, playerID, amount)
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

// resolveBets resolves all pending bets
func (g *Game) resolveBets() {
	sum := g.Dice[0] + g.Dice[1] + g.Dice[2]

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
	// Check if bet wins
	if !g.betWins(betType, sum) {
		return 0
	}

	// Get base payout ratio
	ratio, ok := payoutRatios[betType]
	if !ok {
		return 0
	}

	// Handle single bet with multipliers
	if betType == BetSingle {
		number := extractNumberFromBetType(betType)
		count := g.countDice(number)
		if count > 0 {
			multiplier := singleBetMultipliers[count]
			return amount + amount*multiplier
		}
		return 0
	}

	// Calculate payout (bet amount + winnings)
	return amount + int64(float64(amount)*ratio)
}

// extractNumberFromBetType extracts the number from a bet type (e.g., "single_3" -> 3)
func extractNumberFromBetType(betType string) int {
	var num int
	fmt.Sscanf(betType, "%*[^0-9]%d", &num)
	return num
}

// betWins determines if a bet wins based on the roll result
func (g *Game) betWins(betType string, sum int) bool {
	// Handle bets with embedded numbers (single_1, specific_double_3, etc.)
	if len(betType) > 10 && betType[:len(betType)-2] == "specific_double" ||
		len(betType) > 7 && betType[:len(betType)-2] == "single" ||
		len(betType) > 15 && betType[:len(betType)-2] == "specific_triple" {

		number := extractNumberFromBetType(betType)
		if number < 1 || number > 6 {
			return false
		}

		switch {
		case len(betType) > 15 && betType[:len(betType)-2] == "specific_triple":
			// Specific triple - all three dice must match the number
			return g.Dice[0] == number && g.Dice[1] == number && g.Dice[2] == number

		case len(betType) > 10 && betType[:len(betType)-2] == "specific_double":
			// Specific double - at least two dice must match
			count := g.countDice(number)
			return count >= 2

		case len(betType) > 7 && betType[:len(betType)-2] == "single":
			// Single - at least one die matches
			count := g.countDice(number)
			return count >= 1
		}
	}

	switch betType {
	case BetSmall:
		// Small wins on 4-10 excluding triples
		return sum >= 4 && sum <= 10 && !g.isTriple()

	case BetBig:
		// Big wins on 11-17 excluding triples
		return sum >= 11 && sum <= 17 && !g.isTriple()

	case BetAnyTriple:
		// Any triple - all three dice are the same
		return g.isTriple()

	case BetFourNumber:
		// Four number - sum equals specific values (4,5,6,7,8,9,10,11,12,13,14,15,16)
		// Actually this is a bet that the sum will be one of 4 specific numbers
		// Simplified: check if sum is in typical range
		return sum >= 4 && sum <= 17

	case BetThreeNumber:
		// Three number - bet on specific three dice combination
		// Simplified: any sum
		return true

	case BetTwoNumber:
		// Two number - any two dice match specific combination
		return true

	default:
		return false
	}
}

// countDice counts how many dice show a specific number
func (g *Game) countDice(number int) int {
	count := 0
	for _, d := range g.Dice {
		if d == number {
			count++
		}
	}
	return count
}

// isTriple checks if all three dice are the same
func (g *Game) isTriple() bool {
	return g.Dice[0] == g.Dice[1] && g.Dice[1] == g.Dice[2]
}

// isDouble checks if at least two dice are the same
func (g *Game) isDouble() bool {
	return (g.Dice[0] == g.Dice[1]) || (g.Dice[0] == g.Dice[2]) || (g.Dice[1] == g.Dice[2])
}

// GetTotal returns the sum of dice
func (g *Game) GetTotal() int {
	total := 0
	for _, d := range g.Dice {
		total += d
	}
	return total
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
		Sum:            g.GetTotal(),
		Bets:           g.Bets,
		Payouts:        g.Payouts,
		ProvablyFair:   g.ProvablyFair,
		ServerSeedHash: hashSeed(g.ServerSeed),
		ClientSeed:     g.ClientSeed,
		Nonce:          g.Nonce,
		IsTriple:       g.isTriple(),
		IsDouble:       g.isDouble(),
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
	Bets           map[string]map[string]int64 `json:"bets"`
	Payouts        map[string]map[string]int64 `json:"payouts"`
	ProvablyFair   bool                        `json:"provably_fair"`
	ServerSeedHash string                      `json:"server_seed_hash,omitempty"`
	ClientSeed     string                      `json:"client_seed,omitempty"`
	Nonce          int                         `json:"nonce,omitempty"`
	IsTriple       bool                        `json:"is_triple"`
	IsDouble       bool                        `json:"is_double"`
	CreatedAt      time.Time                   `json:"created_at"`
	UpdatedAt      time.Time                   `json:"updated_at"`
}

// GetAvailableBets returns all available bet types
func (g *Game) GetAvailableBets() []string {
	return []string{
		BetSmall, BetBig, BetAnyTriple,
		BetFourNumber, BetThreeNumber, BetTwoNumber,
	}
}

// GetAvailableSingleBets returns single number bet keys
func (g *Game) GetAvailableSingleBets() []string {
	bets := []string{}
	for i := 1; i <= 6; i++ {
		bets = append(bets, fmt.Sprintf("single_%d", i))
	}
	return bets
}

// GetAvailableDoubleBets returns specific double bet keys
func (g *Game) GetAvailableDoubleBets() []string {
	bets := []string{}
	for i := 1; i <= 6; i++ {
		bets = append(bets, fmt.Sprintf("specific_double_%d", i))
	}
	return bets
}

// GetAvailableTripleBets returns specific triple bet keys
func (g *Game) GetAvailableTripleBets() []string {
	bets := []string{}
	for i := 1; i <= 6; i++ {
		bets = append(bets, fmt.Sprintf("specific_triple_%d", i))
	}
	return bets
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

	ratio, ok := payoutRatios[betType]
	if !ok {
		return "", "", errors.New("bet type not found")
	}

	return name, fmt.Sprintf("%.0f:1", ratio), nil
}

// Bet information maps
var betNames = map[string]string{
	BetSmall:          "Small",
	BetBig:            "Big",
	BetSpecificTriple: "Specific Triple",
	BetAnyTriple:      "Any Triple",
	BetSpecificDouble: "Specific Double",
	BetFourNumber:     "Four Number",
	BetThreeNumber:    "Three Number",
	BetTwoNumber:      "Two Number",
	BetSingle:         "Single",
}

// ValidateBet validates a bet amount and type
func (g *Game) ValidateBet(betType string, amount int64) error {
	if amount <= 0 {
		return errors.New("bet amount must be positive")
	}

	if amount < g.MinBet {
		return fmt.Errorf("minimum bet is %d", g.MinBet)
	}

	if amount > g.MaxBet {
		return fmt.Errorf("maximum bet is %d", g.MaxBet)
	}

	// Check bet type exists
	if _, ok := payoutRatios[betType]; !ok {
		// Check if it's a numbered bet type
		if len(betType) > 7 {
			prefix := betType[:len(betType)-2]
			if prefix == "single" || prefix == "specific_double" || prefix == "specific_triple" {
				return nil
			}
		}
		return errors.New("invalid bet type")
	}

	return nil
}

// GetDiceCounts returns the count of each die face
func (g *Game) GetDiceCounts() map[int]int {
	counts := make(map[int]int)
	for _, d := range g.Dice {
		counts[d]++
	}
	return counts
}
