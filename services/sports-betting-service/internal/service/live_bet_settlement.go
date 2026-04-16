package service

import (
	"context"
	"fmt"
	"time"

	"github.com/game_engine/sports-betting-service/internal/model"
)

// UpdateEventStatus updates the status of a live event
func (s *LiveBettingService) UpdateEventStatus(ctx context.Context, update model.EventStatusUpdate) error {
	// Update event in repository
	err := s.repo.UpdateEventStatus(ctx, update.EventID, update.Status, update.HomeScore, update.AwayScore, update.Period, update.Minute)
	if err != nil {
		return err
	}

	// If event is finished, settle all pending bets
	if update.Status == "finished" || update.Status == "closed" {
		go s.settleEventBets(ctx, update.EventID)
	}

	return nil
}

// settleEventBets settles all bets for a finished event
func (s *LiveBettingService) settleEventBets(ctx context.Context, eventID string) {
	// Get all pending bets for this event
	bets, err := s.repo.GetPendingBetsForEvent(ctx, eventID)
	if err != nil {
		return // Log error in production
	}

	// Get event result
	event, err := s.repo.GetEventByID(ctx, eventID)
	if err != nil {
		return
	}

	// Settle each bet
	for _, bet := range bets {
		result := s.calculateBetResult(ctx, &bet, event)
		_ = s.repo.UpdateBetStatus(ctx, bet.BetID, result.Status, result.WinAmount)
	}
}

// calculateBetResult calculates the result of a bet
func (s *LiveBettingService) calculateBetResult(ctx context.Context, bet *model.Bet, event *model.Event) BetResult {
	result := BetResult{}

	// Get the market for this bet
	markets, _ := s.repo.GetMarkets(ctx, bet.EventID)
	var market *model.Market
	for _, m := range markets {
		if m.MarketID == bet.MarketID {
			market = &m
			break
		}
	}

	if market == nil {
		result.Status = model.BetStatusVoided
		return result
	}

	// Find the selection
	var selection *model.Selection
	for _, sel := range market.Selections {
		if sel.Selection == bet.Selection {
			selection = &sel
			break
		}
	}

	if selection == nil || !selection.Result.Valid {
		result.Status = model.BetStatusPending
		return result
	}

	// Determine win/loss
	if selection.Result.Win {
		result.Status = model.BetStatusWon
		result.WinAmount = bet.Stake * bet.Odds
	} else {
		result.Status = model.BetStatusLost
	}

	return result
}

type BetResult struct {
	Status    model.BetStatus
	WinAmount float64
}

// GetCashOutAmount calculates the cash out amount for a bet
func (s *LiveBettingService) GetCashOutAmount(ctx context.Context, betID string) (*model.CashOutCalculation, error) {
	if !s.cfg.CashOutEnabled {
		return &model.CashOutCalculation{
			BetID:            betID,
			Eligible:         false,
			IneligibleReason: "Cash out is not available",
		}, nil
	}

	// Get the bet
	bet, err := s.repo.GetBetByID(ctx, betID)
	if err != nil {
		return nil, fmt.Errorf("bet not found: %s", betID)
	}

	if bet == nil {
		return nil, fmt.Errorf("bet not found: %s", betID)
	}

	// Check if bet is eligible for cash out
	if bet.Status != string(model.BetStatusPending) {
		return &model.CashOutCalculation{
			BetID:            betID,
			Eligible:         false,
			IneligibleReason: fmt.Sprintf("Bet status is %s, not eligible for cash out", bet.Status),
		}, nil
	}

	// Get current odds for each selection
	currentOdds := bet.Odds
	if parlayBet, err := s.repo.GetParlayBetByID(ctx, betID); err == nil && parlayBet != nil {
		// Recalculate current odds based on remaining selections
		currentOdds = 1.0
		for _, sel := range parlayBet.Selections {
			if sel.Status == "pending" {
				event, _ := s.repo.GetEventByID(ctx, sel.EventID)
				if event != nil && event.Status == "live" {
					// Get current odds from live markets
					markets, _ := s.repo.GetLiveMarkets(ctx, sel.EventID)
					for _, m := range markets {
						if m.MarketID == sel.MarketID {
							for _, s := range m.Selections {
								if s.Selection == sel.Selection {
									currentOdds *= s.Odds
									break
								}
							}
						}
					}
				} else {
					currentOdds *= sel.Odds
				}
			}
		}
	}

	// Calculate cash out amount using cash out formula
	// Cash out = (Original Stake * Current Odds) - Original Stake
	// Applied with a house edge (e.g., 5%)
	houseEdge := 0.05
	cashOutAmount := (bet.Stake * currentOdds) * (1 - houseEdge)
	profit := cashOutAmount - bet.Stake

	return &model.CashOutCalculation{
		BetID:         betID,
		OriginalStake: bet.Stake,
		OriginalOdds:  bet.Odds,
		CurrentOdds:   currentOdds,
		CurrentValue:  cashOutAmount,
		PotentialWin:  bet.PotentialWin,
		Profit:        profit,
		Eligible:      true,
	}, nil
}

// RequestCashOut requests a cash out for a bet
func (s *LiveBettingService) RequestCashOut(ctx context.Context, userID, betID string) (*model.CashOut, error) {
	if !s.cfg.CashOutEnabled {
		return nil, fmt.Errorf("cash out is not available")
	}

	// Get cash out calculation
	calc, err := s.GetCashOutAmount(ctx, betID)
	if err != nil {
		return nil, err
	}

	if !calc.Eligible {
		return nil, fmt.Errorf("bet is not eligible for cash out: %s", calc.IneligibleReason)
	}

	// Verify bet belongs to user
	bet, err := s.repo.GetBetByID(ctx, betID)
	if err != nil || bet.UserID != userID {
		return nil, fmt.Errorf("bet not found or does not belong to user")
	}

	// Create cash out record
	cashOut := &model.CashOut{
		CashOutID:     fmt.Sprintf("co_%d", time.Now().UnixNano()),
		BetID:         betID,
		UserID:        userID,
		OriginalStake: calc.OriginalStake,
		OriginalOdds:  calc.OriginalOdds,
		CurrentOdds:   calc.CurrentOdds,
		CashOutAmount: calc.CurrentValue,
		Profit:        calc.Profit,
		Status:        "pending",
		RequestedAt:   time.Now(),
	}

	// Save cash out request
	err = s.repo.SaveCashOut(ctx, cashOut)
	if err != nil {
		return nil, fmt.Errorf("failed to save cash out: %w", err)
	}

	// Process cash out (in production, this would be a transaction)
	err = s.processCashOut(ctx, cashOut)
	if err != nil {
		cashOut.Status = "failed"
		_ = s.repo.UpdateCashOutStatus(ctx, cashOut.CashOutID, "failed")
		return nil, fmt.Errorf("cash out processing failed: %w", err)
	}

	cashOut.Status = "completed"
	now := time.Now()
	cashOut.CompletedAt = &now
	_ = s.repo.UpdateCashOutStatus(ctx, cashOut.CashOutID, "completed")

	// Update bet status
	_ = s.repo.UpdateBetStatus(ctx, betID, model.BetStatusCashOut, calc.CurrentValue)

	return cashOut, nil
}

// processCashOut processes the cash out (refunds to wallet)
func (s *LiveBettingService) processCashOut(ctx context.Context, cashOut *model.CashOut) error {
	// In production, this would call wallet service via gRPC
	// walletClient := wallet.NewWalletClient(conn)
	// _, err := walletClient.CreditBalance(ctx, &wallet.CreditBalanceRequest{
	//     UserId: cashOut.UserID,
	//     Amount: cashOut.CashOutAmount,
	//     Type:   "cash_out",
	//     RefId:  cashOut.CashOutID,
	// })

	// For now, just return success
	return nil
}
