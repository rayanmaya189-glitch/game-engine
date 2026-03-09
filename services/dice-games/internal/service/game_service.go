package service

// GameService manages dice games
type GameService struct{}

// NewGameService creates a new game service
func NewGameService() (*GameService, error) {
	return &GameService{}, nil
}
