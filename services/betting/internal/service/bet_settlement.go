package service

import (
	"errors"
	"fmt"
	"time"
)

// AcceptBet accepts a bet
func (s *BettingService) AcceptBet(bet *Bet) error {
	if bet.Status != BetStatusPlaced {
		return errors.New("bet is not in placed status")
	}

	bet.Status = BetStatusAccepted
	bet.UpdatedAt = time.Now()

	return nil
}

// ActivateBet activates a bet (when event starts)
func (s *BettingService) ActivateBet(bet *Bet) error {
	if bet.Status != BetStatusAccepted {
		return errors.New("bet is not in accepted status")
	}

	bet.Status = BetStatusActive
	bet.UpdatedAt = time.Now()

	return nil
}

// SettleBet settles a bet with results
func (s *BettingService) SettleBet(bet *Bet, results map[string]string) error {
	if bet.Status != BetStatusActive {
		return errors.New("bet is not in active status")
	}

	// Update selection results
	allWon := true
	allSettled := true

	for i := range bet.Selections {
		result, ok := results[bet.Selections[i].OutcomeID]
		if !ok {
			allSettled = false
			continue
		}

		bet.Selections[i].Result = result
		bet.Selections[i].Status = OutcomeWon

		t := time.Now()
		bet.Selections[i].SettledAt = &t

		switch result {
		case "won":
			// Selection won
		case "lost":
			bet.Selections[i].Status = OutcomeLost
			allWon = false
		case "void":
			bet.Selections[i].Status = OutcomeVoid
		case "push":
			bet.Selections[i].Status = OutcomePush
		}
	}

	// Determine bet result
	if !allSettled {
		return errors.New("not all selections have results")
	}

	switch bet.Type {
	case BetTypeSingle:
		if bet.Selections[0].Status == OutcomeWon {
			bet.Status = BetStatusSettled
		} else if bet.Selections[0].Status == OutcomeVoid {
			// Voided selection = stake returned
			bet.Status = BetStatusSettled
		} else {
			bet.Status = BetStatusSettled
		}

	case BetTypeAccumulator:
		if allWon {
			bet.Status = BetStatusSettled
		} else {
			bet.Status = BetStatusSettled // Lost
		}

	case BetTypeSystem:
		// System bets settle based on number of winning selections
		wins := 0
		for _, sel := range bet.Selections {
			if sel.Status == OutcomeWon {
				wins++
			}
		}
		bet.Status = BetStatusSettled
	}

	now := time.Now()
	bet.SettledAt = &now
	bet.UpdatedAt = now

	return nil
}

// CalculatePayout calculates the actual payout for a settled bet
func (s *BettingService) CalculatePayout(bet *Bet) int64 {
	if bet.Status != BetStatusSettled {
		return 0
	}

	switch bet.Type {
	case BetTypeSingle:
		if bet.Selections[0].Status == OutcomeWon {
			return bet.PotentialWin
		} else if bet.Selections[0].Status == OutcomeVoid {
			return bet.Stake // Return stake for void
		}
		return 0

	case BetTypeAccumulator:
		// Simplified: return full win if all won, 0 if any lost
		allWon := true
		for _, sel := range bet.Selections {
			if sel.Status != OutcomeWon {
				allWon = false
				break
			}
		}
		if allWon {
			return bet.PotentialWin
		}
		return 0

	case BetTypeSystem:
		// Simplified: would need detailed calculation
		wins := 0
		for _, sel := range bet.Selections {
			if sel.Status == OutcomeWon {
				wins++
			}
		}
		if wins >= 2 { // Minimum for any return
			return int64(float64(bet.Stake) * bet.Odds * 0.5) // Simplified
		}
		return 0

	default:
		return 0
	}
}

// VoidBet voids a bet with a reason
func (s *BettingService) VoidBet(bet *Bet, reason string) error {
	if bet.Status == BetStatusSettled || bet.Status == BetStatusPaid {
		return errors.New("cannot void a settled or paid bet")
	}

	bet.Status = BetStatusVoided
	bet.VoidReason = reason
	bet.UpdatedAt = time.Now()

	return nil
}

// ValidateBet validates a bet
func (s *BettingService) ValidateBet(bet *Bet) error {
	if bet.Stake < s.minBet {
		return fmt.Errorf("minimum bet is %d", s.minBet)
	}
	if bet.Stake > s.maxBet {
		return fmt.Errorf("maximum bet is %d", s.maxBet)
	}
	if bet.PotentialWin > s.maxPayout {
		return fmt.Errorf("maximum payout is %d", s.maxPayout)
	}
	if bet.Odds > s.maxOdds {
		return fmt.Errorf("maximum odds is %.2f", s.maxOdds)
	}
	return nil
}
