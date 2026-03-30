package registry

import (
	"errors"
	"sync"
)

// GameDefinition represents a game definition
type GameDefinition struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	Subtype     string                 `json:"subtype"`
	Description string                 `json:"description"`
	Thumbnail   string                 `json:"thumbnail"`
	Config      map[string]interface{} `json:"config"`
	RTP         float64                `json:"rtp"`
	HouseEdge   float64                `json:"house_edge"`
	MinBet      int64                  `json:"min_bet"`
	MaxBet      int64                  `json:"max_bet"`
	MaxPayout   int64                  `json:"max_payout"`
	Status      string                 `json:"status"`
}

// GameRegistry manages game definitions
type GameRegistry struct {
	games map[string]*GameDefinition
	mu    sync.RWMutex
}

// NewGameRegistry creates a new game registry
func NewGameRegistry() *GameRegistry {
	return &GameRegistry{
		games: make(map[string]*GameDefinition),
	}
}

// Register registers a game definition
func (r *GameRegistry) Register(game *GameDefinition) error {
	if game.ID == "" {
		return errors.New("game ID is required")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.games[game.ID] = game
	return nil
}

// Get returns a game definition by ID
func (r *GameRegistry) Get(id string) (*GameDefinition, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	game, ok := r.games[id]
	if !ok {
		return nil, errors.New("game not found")
	}

	return game, nil
}

// List returns all game definitions
func (r *GameRegistry) List() []*GameDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*GameDefinition, 0, len(r.games))
	for _, game := range r.games {
		result = append(result, game)
	}

	return result
}

// Update updates a game definition
func (r *GameRegistry) Update(game *GameDefinition) error {
	if game.ID == "" {
		return errors.New("game ID is required")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.games[game.ID]; !ok {
		return errors.New("game not found")
	}

	r.games[game.ID] = game
	return nil
}

// Delete removes a game definition
func (r *GameRegistry) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.games[id]; !ok {
		return errors.New("game not found")
	}

	delete(r.games, id)
	return nil
}

// InitializeDefaultGames registers default games
func (r *GameRegistry) InitializeDefaultGames() {
	for _, game := range DefaultGames() {
		if err := r.Register(game); err != nil {
			continue
		}
	}
}
