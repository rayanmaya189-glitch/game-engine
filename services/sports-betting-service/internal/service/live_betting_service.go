package service

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/game_engine/sports-betting-service/internal/model"
	"github.com/game_engine/sports-betting-service/internal/repository"
)

// LiveBettingService handles live/in-play betting operations
type LiveBettingService struct {
	repo *repository.SportsRepository
	cfg  *SportsConfig
}

// NewLiveBettingService creates a new live betting service
func NewLiveBettingService(repo *repository.SportsRepository, cfg *SportsConfig) *LiveBettingService {
	return &LiveBettingService{repo: repo, cfg: cfg}
}

// SportsConfig holds sports betting configuration
type SportsConfig struct {
	MinBetAmount        float64 `mapstructure:"min_bet_amount"`
	MaxBetAmount        float64 `mapstructure:"max_bet_amount"`
	MaxOdds             float64 `mapstructure:"max_odds"`
	MaxParlaySelections int     `mapstructure:"max_parlay_selections"`
	CashOutEnabled      bool    `mapstructure:"cash_out_enabled"`
}

// GetLiveEvents returns all live events
func (s *LiveBettingService) GetLiveEvents(ctx context.Context) ([]model.LiveEvent, error) {
	return s.repo.GetLiveEvents(ctx)
}

// GetLiveEventByID returns a specific live event
func (s *LiveBettingService) GetLiveEventByID(ctx context.Context, eventID string) (*model.LiveEvent, error) {
	event, err := s.repo.GetLiveEventByID(ctx, eventID)
	if err != nil {
		return nil, err
	}
	if event == nil {
		return nil, fmt.Errorf("event not found: %s", eventID)
	}
	return event, nil
}

// GetLiveMarkets returns markets for a live event
func (s *LiveBettingService) GetLiveMarkets(ctx context.Context, eventID string) ([]model.Market, error) {
	return s.repo.GetLiveMarkets(ctx, eventID)
}

// UpdateLiveOdds updates odds for a live event (called by external feed)
func (s *LiveBettingService) UpdateLiveOdds(ctx context.Context, eventID string, markets []model.Market) error {
	// Get current event to track odds changes
	currentEvent, err := s.repo.GetLiveEventByID(ctx, eventID)
	if err != nil {
		return err
	}

	// Track odds changes
	var oddsChanges []model.OddsChange
	for _, newMarket := range markets {
		for _, currentMarket := range currentEvent.Markets {
			if currentMarket.MarketID == newMarket.MarketID {
				for _, newSel := range newMarket.Selections {
					for _, curSel := range currentMarket.Selections {
						if curSel.SelectionID == newSel.SelectionID && curSel.Odds != newSel.Odds {
							oddsChanges = append(oddsChanges, model.OddsChange{
								MarketID:  newMarket.MarketID,
								Selection: curSel.Selection,
								OldOdds:   curSel.Odds,
								NewOdds:   newSel.Odds,
								Timestamp: time.Now(),
							})
						}
					}
				}
			}
		}
	}

	// Update the event
	err = s.repo.UpdateLiveEvent(ctx, eventID, markets)
	if err != nil {
		return err
	}

	// Save odds history if there were changes
	if len(oddsChanges) > 0 {
		_ = s.repo.SaveOddsChanges(ctx, eventID, oddsChanges)
	}

	return nil
}

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
		result := s.calculateBetResult(bet, event)
		_ = s.repo.UpdateBetStatus(ctx, bet.BetID, result.Status, result.WinAmount)
	}
}

// calculateBetResult calculates the result of a bet
func (s *LiveBettingService) calculateBetResult(bet *model.Bet, event *model.Event) BetResult {
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

// PlaceParlayBet places a parlay (accumulator) bet
func (s *LiveBettingService) PlaceParlayBet(ctx context.Context, userID string, selections []model.ParlaySelection, stake float64) (*PlaceBetResponse, error) {
	// Validate selections
	if len(selections) < 2 {
		return nil, fmt.Errorf("parlay must have at least 2 selections")
	}
	if len(selections) > s.cfg.MaxParlaySelections {
		return nil, fmt.Errorf("parlay cannot have more than %d selections", s.cfg.MaxParlaySelections)
	}

	// Validate stake
	if stake < s.cfg.MinBetAmount {
		return nil, fmt.Errorf("stake %.2f is below minimum %.2f", stake, s.cfg.MinBetAmount)
	}
	if stake > s.cfg.MaxBetAmount {
		return nil, fmt.Errorf("stake %.2f is above maximum %.2f", stake, s.cfg.MaxBetAmount)
	}

	// Validate each selection and calculate total odds
	var validatedSelections []model.ParlaySelection
	totalOdds := 1.0

	for i, sel := range selections {
		// Validate event exists and is available
		event, err := s.repo.GetEventByID(ctx, sel.EventID)
		if err != nil || event == nil {
			return nil, fmt.Errorf("event not found: %s", sel.EventID)
		}

		if event.Status != "scheduled" && event.Status != "live" {
			return nil, fmt.Errorf("event %s is not available for betting", sel.EventID)
		}

		// Validate market exists
		markets, err := s.repo.GetMarkets(ctx, sel.EventID)
		if err != nil {
			return nil, fmt.Errorf("failed to get markets for event: %s", sel.EventID)
		}

		var marketFound bool
		for _, m := range markets {
			if m.MarketID == sel.MarketID {
				marketFound = true
				// Verify selection exists and odds match
				for _, s := range m.Selections {
					if s.Selection == sel.Selection {
						if s.Odds != sel.Odds {
							// Update odds to current
							sel.Odds = s.Odds
						}
						break
					}
				}
				break
			}
		}

		if !marketFound {
			return nil, fmt.Errorf("market not found: %s for event %s", sel.MarketID, sel.EventID)
		}

		sel.EventName = event.HomeTeam + " vs " + event.AwayTeam
		sel.Status = "pending"
		validatedSelections = append(validatedSelections, sel)
		totalOdds *= sel.Odds

		// Update selection index
		selections[i] = sel
	}

	// Validate total odds
	if totalOdds > s.cfg.MaxOdds {
		return nil, fmt.Errorf("total odds %.2f exceeds maximum %.2f", totalOdds, s.cfg.MaxOdds)
	}

	// Calculate potential win
	potentialWin := stake * totalOdds
	if potentialWin > s.cfg.MaxBetAmount*100 {
		return nil, fmt.Errorf("potential win %.2f exceeds maximum limit", potentialWin)
	}

	// Create the parlay bet
	bet := &model.ParlayBet{
		BetID:        generateBetID(),
		UserID:       userID,
		Selections:   validatedSelections,
		Stake:        stake,
		TotalOdds:    totalOdds,
		PotentialWin: potentialWin,
		Status:       model.BetStatusPending,
		PlacedAt:     time.Now(),
	}

	// Save bet to database
	err := s.repo.SaveParlayBet(ctx, bet)
	if err != nil {
		return nil, fmt.Errorf("failed to save bet: %w", err)
	}

	return &PlaceBetResponse{
		Success:      true,
		Message:      "Parlay bet placed successfully",
		BetID:        bet.BetID,
		PotentialWin: potentialWin,
	}, nil
}

// GetCashOutAmount calculates the cash out amount for a bet
func (s *LiveBettingService) GetCashOutAmount(ctx context.Context, betID string) (*CashOutCalculation, error) {
	if !s.cfg.CashOutEnabled {
		return &CashOutCalculation{
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
	if bet.Status != model.BetStatusPending {
		return &CashOutCalculation{
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

	return &CashOutCalculation{
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
func (s *LiveBettingService) RequestCashOut(ctx context.Context, userID, betID string) (*CashOut, error) {
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

// BuildParlay calculates the potential win for a parlay
func (s *LiveBettingService) BuildParlay(ctx context.Context, selections []model.BetBuilderSelection, stake float64) (*model.BetBuilder, error) {
	if len(selections) < 2 {
		return nil, fmt.Errorf("parlay must have at least 2 selections")
	}
	if len(selections) > s.cfg.MaxParlaySelections {
		return nil, fmt.Errorf("parlay cannot have more than %d selections", s.cfg.MaxParlaySelections)
	}

	totalOdds := 1.0
	var validatedSelections []model.BetBuilderSelection

	for _, sel := range selections {
		// Validate event is available
		event, err := s.repo.GetEventByID(ctx, sel.EventID)
		if err != nil || event == nil {
			return nil, fmt.Errorf("event not found: %s", sel.EventID)
		}

		if event.Status != "scheduled" && event.Status != "live" {
			return nil, fmt.Errorf("event %s is not available for betting", sel.EventID)
		}

		// Get current odds
		markets, err := s.repo.GetMarkets(ctx, sel.EventID)
		if err != nil {
			return nil, fmt.Errorf("failed to get markets for event: %s", sel.EventID)
		}

		var currentOdds float64
		for _, m := range markets {
			if m.MarketID == sel.MarketType || m.MarketType == sel.MarketType {
				for _, s := range m.Selections {
					if s.Selection == sel.Selection {
						currentOdds = s.Odds
						break
					}
				}
			}
		}

		if currentOdds == 0 {
			currentOdds = sel.Odds // Use provided odds if not found
		}

		sel.Odds = currentOdds
		sel.EventName = event.HomeTeam + " vs " + event.AwayTeam
		validatedSelections = append(validatedSelections, sel)
		totalOdds *= currentOdds
	}

	potentialWin := stake * totalOdds

	return &model.BetBuilder{
		Selections:   validatedSelections,
		Stake:        stake,
		TotalOdds:    totalOdds,
		PotentialWin: potentialWin,
	}, nil
}

// GetUserParlayBets returns all parlay bets for a user
func (s *LiveBettingService) GetUserParlayBets(ctx context.Context, userID string, page, limit int) ([]model.ParlayBet, int, error) {
	if limit < 1 || limit > 100 {
		limit = 20
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	return s.repo.GetUserParlayBets(ctx, userID, limit, offset)
}

// GetUserCashOuts returns all cash outs for a user
func (s *LiveBettingService) GetUserCashOuts(ctx context.Context, userID string, page, limit int) ([]model.CashOut, int, error) {
	if limit < 1 || limit > 100 {
		limit = 20
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	return s.repo.GetUserCashOuts(ctx, userID, limit, offset)
}

// Round helper for float comparison
func round(x float64) float64 {
	return math.Round(x*100) / 100
}
