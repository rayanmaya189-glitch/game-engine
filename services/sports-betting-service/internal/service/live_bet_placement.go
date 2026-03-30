package service

import (
	"context"
	"fmt"
	"time"

	"github.com/game_engine/sports-betting-service/internal/model"
)

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
