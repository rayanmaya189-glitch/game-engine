package service

// GameService manages slot games
type GameService struct{}

// NewGameService creates a new game service
func NewGameService() (*GameService, error) {
	return &GameService{}, nil
}
