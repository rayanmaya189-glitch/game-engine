package service

import (
	"errors"
	"time"

	"github.com/game_engine/betting/internal/repository"
)

// Bet types
const (
	BetTypeSingle      = "single"
	BetTypeAccumulator = "accumulator"
	BetTypeSystem      = "system"
)

// System bet types
const (
	SystemPatent     = "patent"
	SystemYankee     = "yankee"
	SystemCanadian   = "canadian"
	SystemHeinz      = "heinz"
	SystemSuperHeinz = "super_heinz"
	SystemGoliath    = "goliath"
)

// Bet status
const (
	BetStatusPlaced    = "placed"
	BetStatusAccepted  = "accepted"
	BetStatusActive    = "active"
	BetStatusSettled   = "settled"
	BetStatusPaid      = "paid"
	BetStatusVoided    = "voided"
	BetStatusCancelled = "cancelled"
)

// Odds format
const (
	OddsDecimal    = "decimal"
	OddsFractional = "fractional"
	OddsAmerican   = "american"
	OddsHongKong   = "hongkong"
)

// Outcome status
const (
	OutcomePending = "pending"
	OutcomeWon     = "won"
	OutcomeLost    = "lost"
	OutcomeVoid    = "void"
	OutcomePush    = "push"
)

// BettingService manages betting operations
type BettingService struct {
	minBet          int64
	maxBet          int64
	maxPayout       int64
	maxOdds         float64
	defaultCurrency string
	settlement      SettlementService
	limits          map[string]*LimitConfig
	repo            repository.BettingRepository
}

// LimitConfig represents bet limits
type LimitConfig struct {
	MinBet    int64 `json:"min_bet"`
	MaxBet    int64 `json:"max_bet"`
	MaxPayout int64 `json:"max_payout"`
}

// Selection represents a single selection in a bet
type Selection struct {
	ID         string     `json:"id"`
	EventID    string     `json:"event_id"`
	OutcomeID  string     `json:"outcome_id"`
	Odds       float64    `json:"odds"`
	OddsFormat string     `json:"odds_format"`
	Status     string     `json:"status"`
	Result     string     `json:"result"`
	SettledAt  *time.Time `json:"settled_at,omitempty"`
}

// Bet represents a bet
type Bet struct {
	ID           string      `json:"id"`
	UserID       string      `json:"user_id"`
	Type         string      `json:"type"`
	Stake        int64       `json:"stake"`
	Odds         float64     `json:"odds"` // Cumulative odds for accumulator/system
	PotentialWin int64       `json:"potential_win"`
	Status       string      `json:"status"`
	Selections   []Selection `json:"selections"`
	SystemType   string      `json:"system_type,omitempty"`
	Currency     string      `json:"currency"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	SettledAt    *time.Time  `json:"settled_at,omitempty"`
	VoidReason   string      `json:"void_reason,omitempty"`
}

// SettlementService handles bet settlement
type SettlementService interface {
	Settle(bet *Bet, results map[string]string) error
}

// BettingServiceImpl implements BettingService
type BettingServiceImpl struct {
	*BettingService
}

// NewBettingService creates a new betting service with repository
func NewBettingService(repo repository.BettingRepository) (*BettingService, error) {
	return &BettingService{
		minBet:          1,
		maxBet:          10000,
		maxPayout:       100000,
		maxOdds:         1000.0,
		defaultCurrency: "USD",
		limits:          make(map[string]*LimitConfig),
		repo:            repo,
	}, nil
}

// NewBettingServiceWithConfig creates a betting service with custom config and repository
func NewBettingServiceWithConfig(minBet, maxBet int64, maxPayout int64, repo repository.BettingRepository) (*BettingService, error) {
	if minBet <= 0 || maxBet <= 0 || maxBet < minBet {
		return nil, errors.New("invalid bet limits")
	}

	return &BettingService{
		minBet:          minBet,
		maxBet:          maxBet,
		maxPayout:       maxPayout,
		maxOdds:         1000.0,
		defaultCurrency: "USD",
		settlement:      nil,
		limits:          make(map[string]*LimitConfig),
		repo:            repo,
	}, nil
}

// NewBettingServiceFromConfig creates a betting service from configuration values
func NewBettingServiceFromConfig(minBet, maxBet, maxPayout int64, maxOdds float64, defaultCurrency string, repo repository.BettingRepository) (*BettingService, error) {
	if minBet <= 0 || maxBet <= 0 || maxBet < minBet {
		return nil, errors.New("invalid bet limits")
	}
	if defaultCurrency == "" {
		defaultCurrency = "USD"
	}

	return &BettingService{
		minBet:          minBet,
		maxBet:          maxBet,
		maxPayout:       maxPayout,
		maxOdds:         maxOdds,
		defaultCurrency: defaultCurrency,
		settlement:      nil,
		limits:          make(map[string]*LimitConfig),
		repo:            repo,
	}, nil
}

// GetRepository returns the underlying repository
func (s *BettingService) GetRepository() repository.BettingRepository {
	return s.repo
}
