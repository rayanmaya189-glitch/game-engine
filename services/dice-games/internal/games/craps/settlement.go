package craps

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
