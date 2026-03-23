package service

import (
	"errors"
	"fmt"
	"time"
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
	minBet     int64
	maxBet     int64
	maxPayout  int64
	maxOdds    float64
	settlement SettlementService
	limits     map[string]*LimitConfig
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
	bets map[string]*Bet
}

// NewBettingService creates a new betting service
func NewBettingService() (*BettingService, error) {
	return &BettingService{
		minBet:    1,
		maxBet:    10000,
		maxPayout: 100000,
		maxOdds:   1000.0,
		limits:    make(map[string]*LimitConfig),
	}, nil
}

// NewBettingServiceWithConfig creates a betting service with custom config
func NewBettingServiceWithConfig(minBet, maxBet int64, maxPayout int64) (*BettingService, error) {
	if minBet <= 0 || maxBet <= 0 || maxBet < minBet {
		return nil, errors.New("invalid bet limits")
	}

	return &BettingService{
		minBet:     minBet,
		maxBet:     maxBet,
		maxPayout:  maxPayout,
		maxOdds:    1000.0,
		settlement: nil,
		limits:     make(map[string]*LimitConfig),
	}, nil
}

// PlaceSingle places a single bet
func (s *BettingService) PlaceSingle(userID, betID string, stake float64, selection Selection) (*Bet, error) {
	if stake < float64(s.minBet) {
		return nil, fmt.Errorf("minimum bet is %d", s.minBet)
	}
	if stake > float64(s.maxBet) {
		return nil, fmt.Errorf("maximum bet is %d", s.maxBet)
	}

	// Convert odds to decimal
	decimalOdds := s.normalizeOdds(selection.Odds, selection.OddsFormat)
	potentialWin := int64(stake * decimalOdds)

	if potentialWin > s.maxPayout {
		return nil, fmt.Errorf("maximum payout is %d", s.maxPayout)
	}

	bet := &Bet{
		ID:           betID,
		UserID:       userID,
		Type:         BetTypeSingle,
		Stake:        int64(stake),
		Odds:         decimalOdds,
		PotentialWin: potentialWin,
		Status:       BetStatusPlaced,
		Selections:   []Selection{selection},
		Currency:     "USD",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return bet, nil
}

// PlaceAccumulator places an accumulator bet
func (s *BettingService) PlaceAccumulator(userID, betID string, stake float64, selections []Selection) (*Bet, error) {
	if len(selections) < 2 {
		return nil, errors.New("accumulator requires at least 2 selections")
	}

	if stake < float64(s.minBet) {
		return nil, fmt.Errorf("minimum bet is %d", s.minBet)
	}
	if stake > float64(s.maxBet) {
		return nil, fmt.Errorf("maximum bet is %d", s.maxBet)
	}

	// Calculate cumulative odds
	cumulativeOdds := 1.0
	for i := range selections {
		selections[i].Odds = s.normalizeOdds(selections[i].Odds, selections[i].OddsFormat)
		cumulativeOdds *= selections[i].Odds
	}

	if cumulativeOdds > s.maxOdds {
		return nil, fmt.Errorf("maximum cumulative odds is %.2f", s.maxOdds)
	}

	potentialWin := int64(stake * cumulativeOdds)

	if potentialWin > s.maxPayout {
		return nil, fmt.Errorf("maximum payout is %d", s.maxPayout)
	}

	bet := &Bet{
		ID:           betID,
		UserID:       userID,
		Type:         BetTypeAccumulator,
		Stake:        int64(stake),
		Odds:         cumulativeOdds,
		PotentialWin: potentialWin,
		Status:       BetStatusPlaced,
		Selections:   selections,
		Currency:     "USD",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return bet, nil
}

// PlaceSystem places a system bet
func (s *BettingService) PlaceSystem(userID, betID string, stake float64, selections []Selection, systemType string) (*Bet, error) {
	validSystems := map[string]int{
		SystemPatent:     3, // 4 bets (3 singles + 3 doubles + 1 treble)
		SystemYankee:     4, // 11 bets
		SystemCanadian:   5, // 26 bets
		SystemHeinz:      6, // 57 bets
		SystemSuperHeinz: 7, // 120 bets
		SystemGoliath:    8, // 247 bets
	}

	minSelections, ok := validSystems[systemType]
	if !ok {
		return nil, errors.New("invalid system type")
	}

	if len(selections) < minSelections {
		return nil, fmt.Errorf("system %s requires at least %d selections", systemType, minSelections)
	}

	// Calculate number of combinations based on system type
	numCombos := s.calculateSystemCombinations(len(selections), systemType)
	totalStake := int64(stake) * int64(numCombos)

	if totalStake > s.maxBet {
		return nil, fmt.Errorf("total stake exceeds maximum of %d", s.maxBet)
	}

	// Calculate potential win (simplified - assumes all selections win)
	cumulativeOdds := 1.0
	for i := range selections {
		selections[i].Odds = s.normalizeOdds(selections[i].Odds, selections[i].OddsFormat)
		cumulativeOdds *= selections[i].Odds
	}

	// Simplified: estimate potential win based on all combinations
	// In reality, this would need detailed calculation per combination type
	potentialWin := int64(float64(totalStake) * cumulativeOdds)

	if potentialWin > s.maxPayout {
		return nil, fmt.Errorf("maximum payout is %d", s.maxPayout)
	}

	bet := &Bet{
		ID:           betID,
		UserID:       userID,
		Type:         BetTypeSystem,
		Stake:        totalStake,
		Odds:         cumulativeOdds,
		PotentialWin: potentialWin,
		Status:       BetStatusPlaced,
		Selections:   selections,
		SystemType:   systemType,
		Currency:     "USD",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return bet, nil
}

// calculateSystemCombinations calculates number of bets in a system bet
func (s *BettingService) calculateSystemCombinations(numSelections int, systemType string) int {
	// Simplified combination calculation
	switch systemType {
	case SystemPatent:
		return 4 + (numSelections-3)*3
	case SystemYankee:
		return 11 + (numSelections-4)*6
	case SystemCanadian:
		return 26 + (numSelections-5)*10
	case SystemHeinz:
		return 57 + (numSelections-6)*15
	case SystemSuperHeinz:
		return 120 + (numSelections-7)*21
	case SystemGoliath:
		return 247 + (numSelections-8)*28
	default:
		return 1
	}
}

// normalizeOdds converts odds to decimal format
func (s *BettingService) normalizeOdds(odds float64, format string) float64 {
	switch format {
	case OddsDecimal:
		return odds
	case OddsFractional:
		// 3/2 = 1 + 3/2 = 2.5
		parts := splitFraction(odds)
		return 1.0 + float64(parts[0])/float64(parts[1])
	case OddsAmerican:
		if odds > 0 {
			return 1 + odds/100
		}
		return 1 + 100/(-odds)
	case OddsHongKong:
		return 1 + odds
	default:
		return odds
	}
}

// splitFraction splits fractional odds like 3/2 into numerator and denominator
func splitFraction(odds float64) [2]int {
	// Simplified - in real implementation would parse string format
	return [2]int{1, 1}
}

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

// GetBet returns a bet by ID (would query database in real implementation)
func (s *BettingService) GetBet(betID string) (*Bet, error) {
	// In real implementation, would query database
	return nil, errors.New("not implemented")
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

// GetBetHistory returns bet history for a user
func (s *BettingService) GetBetHistory(userID string, limit, offset int) ([]*Bet, error) {
	// In real implementation, would query database
	return []*Bet{}, nil
}

// GetOpenBets returns all open (non-settled) bets for a user
func (s *BettingService) GetOpenBets(userID string) ([]*Bet, error) {
	// In real implementation, would query database
	return []*Bet{}, nil
}

// CalculateOddsFromDecimal calculates various format odds from decimal
func (s *BettingService) CalculateOddsFromDecimal(decimal float64) map[string]float64 {
	return map[string]float64{
		OddsDecimal:    decimal,
		OddsFractional: decimal - 1,
		OddsAmerican:   (decimal - 1) * 100,
		OddsHongKong:   decimal - 1,
	}
}

// Conversion helpers for odds format
func DecimalToFractional(decimal float64) (int, int) {
	frac := decimal - 1
	// Simplified - would need proper fraction reduction
	numerator := int(frac * 100)
	return numerator, 100
}

func DecimalToAmerican(decimal float64) int {
	if decimal >= 2.0 {
		return int((decimal - 1) * 100)
	}
	return int(-100 / (decimal - 1))
}

func DecimalToHongKong(decimal float64) float64 {
	return decimal - 1
}

// FormatOdds formats odds to specified format
func FormatOdds(decimal float64, format string) string {
	switch format {
	case OddsDecimal:
		return fmt.Sprintf("%.2f", decimal)
	case OddsFractional:
		num, den := DecimalToFractional(decimal)
		return fmt.Sprintf("%d/%d", num, den)
	case OddsAmerican:
		return fmt.Sprintf("%d", DecimalToAmerican(decimal))
	case OddsHongKong:
		return fmt.Sprintf("%.2f", DecimalToHongKong(decimal))
	default:
		return fmt.Sprintf("%.2f", decimal)
	}
}

// BetInfo returns information about a bet type
func (s *BettingService) GetBetTypeInfo(betType string) (string, int, error) {
	switch betType {
	case BetTypeSingle:
		return "Single Bet", 1, nil
	case BetTypeAccumulator:
		return "Accumulator", 2, nil
	case BetTypeSystem:
		return "System Bet", 3, nil
	default:
		return "", 0, errors.New("unknown bet type")
	}
}

// GetSystemBetInfo returns information about a system bet type
func (s *BettingService) GetSystemBetInfo(systemType string) (string, int, error) {
	info := map[string][2]int{
		SystemPatent:     {4, 3},
		SystemYankee:     {11, 4},
		SystemCanadian:   {26, 5},
		SystemHeinz:      {57, 6},
		SystemSuperHeinz: {120, 7},
		SystemGoliath:    {247, 8},
	}

	val, ok := info[systemType]
	if !ok {
		return "", 0, errors.New("unknown system type")
	}
	return fmt.Sprintf("%d bets", val[0]), val[1], nil
}

// LimitConfig returns the current limit configuration
func (s *BettingService) GetLimits() *LimitConfig {
	return &LimitConfig{
		MinBet:    s.minBet,
		MaxBet:    s.maxBet,
		MaxPayout: s.maxPayout,
	}
}

// SetLimits updates the bet limits
func (s *BettingService) SetLimits(minBet, maxBet, maxPayout int64) error {
	if minBet <= 0 || maxBet <= 0 || maxPayout <= 0 {
		return errors.New("all limit values must be positive")
	}
	if maxBet < minBet {
		return errors.New("max bet must be >= min bet")
	}

	s.minBet = minBet
	s.maxBet = maxBet
	s.maxPayout = maxPayout

	return nil
}
