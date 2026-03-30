package roulette

import "errors"

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

// calculateTotalPayout calculates total potential payout
func (g *Game) calculateTotalPayout() int64 {
	var total int64
	for _, bet := range g.Bets {
		total += bet.Amount * int64(payouts[bet.BetType])
	}
	return total
}
