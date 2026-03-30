package service

import (
	"errors"
	"fmt"
)

// GetLimits returns the current limit configuration
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

// DecimalToFractional converts decimal odds to fractional
func DecimalToFractional(decimal float64) (int, int) {
	frac := decimal - 1
	// Simplified - would need proper fraction reduction
	numerator := int(frac * 100)
	return numerator, 100
}

// DecimalToAmerican converts decimal odds to American format
func DecimalToAmerican(decimal float64) int {
	if decimal >= 2.0 {
		return int((decimal - 1) * 100)
	}
	return int(-100 / (decimal - 1))
}

// DecimalToHongKong converts decimal odds to Hong Kong format
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
