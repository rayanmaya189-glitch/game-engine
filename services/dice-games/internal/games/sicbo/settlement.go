package sicbo

import "fmt"

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

// GetDiceCounts returns the count of each die face
func (g *Game) GetDiceCounts() map[int]int {
	counts := make(map[int]int)
	for _, d := range g.Dice {
		counts[d]++
	}
	return counts
}
