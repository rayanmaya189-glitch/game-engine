package roulette

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
