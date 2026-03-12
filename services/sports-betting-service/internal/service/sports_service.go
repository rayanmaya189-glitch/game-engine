package service

import (
	"context"
	"fmt"
	"time"

	"github.com/game_engine/sports-betting-service/internal/config"
	"github.com/game_engine/sports-betting-service/internal/model"
	"github.com/game_engine/sports-betting-service/internal/repository"
)

type SportsService struct {
	repo *repository.SportsRepository
	cfg  *config.Config
}

func NewSportsService(repo *repository.SportsRepository, cfg *config.Config) *SportsService {
	return &SportsService{repo: repo, cfg: cfg}
}

type GetSportsResponse struct {
	Sports []model.Sport
}

func (s *SportsService) GetSports(ctx context.Context) (*GetSportsResponse, error) {
	sports, err := s.repo.GetSports(ctx)
	if err != nil {
		return nil, err
	}
	return &GetSportsResponse{Sports: sports}, nil
}

type GetLiveEventsResponse struct {
	Events []model.Event
}

func (s *SportsService) GetLiveEvents(ctx context.Context) (*GetLiveEventsResponse, error) {
	events, err := s.repo.GetLiveEvents(ctx)
	if err != nil {
		return nil, err
	}
	return &GetLiveEventsResponse{Events: events}, nil
}

type GetUpcomingEventsResponse struct {
	Events []model.Event
}

func (s *SportsService) GetUpcomingEvents(ctx context.Context, sportID string, limit int) (*GetUpcomingEventsResponse, error) {
	if limit < 1 || limit > 50 {
		limit = 20
	}

	events, err := s.repo.GetUpcomingEvents(ctx, sportID, limit)
	if err != nil {
		return nil, err
	}
	return &GetUpcomingEventsResponse{Events: events}, nil
}

type GetMarketsResponse struct {
	Markets []model.Market
}

func (s *SportsService) GetMarkets(ctx context.Context, eventID string) (*GetMarketsResponse, error) {
	markets, err := s.repo.GetMarkets(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return &GetMarketsResponse{Markets: markets}, nil
}

type PlaceBetResponse struct {
	Success      bool
	Message      string
	BetID        string
	PotentialWin float64
}

func (s *SportsService) PlaceBet(ctx context.Context, userID, eventID, marketID, selection string, stake, odds float64) (*PlaceBetResponse, error) {
	// Validate user authentication
	if userID == "" {
		return nil, fmt.Errorf("user authentication required")
	}

	// Validate bet amount
	if stake < s.cfg.Sports.MinBetAmount {
		return nil, fmt.Errorf("stake %.2f is below minimum %.2f", stake, s.cfg.Sports.MinBetAmount)
	}
	if stake > s.cfg.Sports.MaxBetAmount {
		return nil, fmt.Errorf("stake %.2f is above maximum %.2f", stake, s.cfg.Sports.MaxBetAmount)
	}

	// Validate odds
	if odds <= 0 || odds > s.cfg.Sports.MaxOdds {
		return nil, fmt.Errorf("odds %.2f is invalid (must be between 0 and %.2f)", odds, s.cfg.Sports.MaxOdds)
	}

	// Validate market exists and is open
	markets, err := s.repo.GetMarkets(ctx, eventID)
	if err != nil {
		return nil, fmt.Errorf("failed to get markets: %w", err)
	}
	if len(markets) == 0 {
		return nil, fmt.Errorf("event not found for event_id %s", eventID)
	}

	// Find the specific market
	var marketFound bool
	for _, m := range markets {
		if m.MarketID == marketID {
			marketFound = true
			if m.Status != "open" {
				return nil, fmt.Errorf("market %s is closed", marketID)
			}
			break
		}
	}
	if !marketFound {
		return nil, fmt.Errorf("market not found: %s", marketID)
	}

	// Validate event status
	event, err := s.repo.GetEventByID(ctx, eventID)
	if err != nil {
		return nil, fmt.Errorf("failed to get event: %w", err)
	}
	if event.Status != "scheduled" && event.Status != "live" {
		return nil, fmt.Errorf("event is not available for betting (status: %s)", event.Status)
	}

	potentialWin := stake * odds
	if potentialWin > s.cfg.Sports.MaxBetAmount*100 {
		return nil, fmt.Errorf("potential win %.2f exceeds maximum limit %.2f", potentialWin, s.cfg.Sports.MaxBetAmount*100)
	}

	// TODO: Integrate with wallet service via gRPC call to check and deduct balance
	// walletClient := wallet.NewWalletClient(conn)
	// balance, err := walletClient.GetBalance(ctx, &wallet.GetBalanceRequest{UserId: userID})
	// if err != nil || balance.Amount < stake {
	//     return nil, fmt.Errorf("insufficient funds")
	// }

	// Begin transaction for atomic bet placement
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Place bet within transaction
	bet := &model.Bet{
		BetID:        generateBetID(),
		UserID:       userID,
		EventID:      eventID,
		MarketID:     marketID,
		Selection:    selection,
		Stake:        stake,
		Odds:         odds,
		PotentialWin: potentialWin,
		Status:       "pending",
	}

	err = s.repo.PlaceBetTx(ctx, tx, bet)
	if err != nil {
		return nil, fmt.Errorf("failed to place bet: %w", err)
	}

	// Deduct from wallet within the same transaction
	// TODO: Integrate with actual wallet service
	// walletDeducted, err := s.repo.DeductWalletBalanceTx(ctx, tx, userID, stake)
	// if err != nil {
	//     return nil, fmt.Errorf("failed to deduct from wallet: %w", err)
	// }

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &PlaceBetResponse{
		Success:      true,
		Message:      "Bet placed successfully",
		BetID:        bet.BetID,
		PotentialWin: potentialWin,
	}, nil
}

type GetUserBetsResponse struct {
	Bets  []model.Bet
	Total int
}

func (s *SportsService) GetUserBets(ctx context.Context, userID string, page, limit int) (*GetUserBetsResponse, error) {
	if limit < 1 || limit > 100 {
		limit = 20
	}
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	bets, total, err := s.repo.GetUserBets(ctx, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	return &GetUserBetsResponse{Bets: bets, Total: total}, nil
}

// generateBetID generates a unique bet ID using timestamp and random component
func generateBetID() string {
	return fmt.Sprintf("bet_%d_%d", time.Now().UnixNano(), time.Now().Nanosecond()%1000)
}
