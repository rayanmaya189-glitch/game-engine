package dice

// Roller interface for dice rolling
type Roller interface {
	Roll(diceCount int) ([]int, error)
}

// RandomRoller uses RNG to roll dice
type RandomRoller struct{}

// NewRandomRoller creates a new random roller
func NewRandomRoller() *RandomRoller {
	return &RandomRoller{}
}

// Roll rolls dice using RNG
func (r *RandomRoller) Roll(diceCount int) ([]int, error) {
	// This would use the RNG service in production
	dice := make([]int, diceCount)
	for i := 0; i < diceCount; i++ {
		dice[i] = 1 // Placeholder - would use RNG
	}
	return dice, nil
}
