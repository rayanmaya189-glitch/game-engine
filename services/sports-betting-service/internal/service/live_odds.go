package service

import (
	"math"
)

// Round helper for float comparison
func round(x float64) float64 {
	return math.Round(x*100) / 100
}
