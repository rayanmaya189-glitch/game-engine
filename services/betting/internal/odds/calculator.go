package odds

// Calculator handles odds calculations
type Calculator struct{}

// NewCalculator creates a new odds calculator
func NewCalculator() *Calculator {
	return &Calculator{}
}

// CalculateDecimal calculates decimal odds payout
func (c *Calculator) CalculateDecimal(odds float64, stake int64) int64 {
	return int64(float64(stake) * odds)
}

// CalculateFractional calculates fractional odds payout
func (c *Calculator) CalculateFractional(numerator, denominator int64, stake int64) int64 {
	return stake + stake*numerator/denominator
}

// CalculateAmerican calculates American odds payout
func (c *Calculator) CalculateAmerican(odds int64, stake int64) int64 {
	if odds > 0 {
		return stake + stake*odds/100
	}
	return stake + stake*100/(-odds)
}
