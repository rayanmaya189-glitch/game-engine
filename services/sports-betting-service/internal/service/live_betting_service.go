package service

import (
	"context"
	"fmt"
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
func (s *LiveBettingService) GetLiveEvents(ctx context.Context) ([]model.Event, error) {
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
