package limit

// Manager handles bet limits
type Manager struct {
	MinBet int64
	MaxBet int64
}

// NewManager creates a new limit manager
func NewManager(minBet, maxBet int64) *Manager {
	return &Manager{
		MinBet: minBet,
		MaxBet: maxBet,
	}
}

// Validate validates a bet amount
func (m *Manager) Validate(amount int64) bool {
	return amount >= m.MinBet && amount <= m.MaxBet
}
