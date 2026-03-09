package service

// BettingService manages betting operations
type BettingService struct{}

// NewBettingService creates a new betting service
func NewBettingService() (*BettingService, error) {
	return &BettingService{}, nil
}
