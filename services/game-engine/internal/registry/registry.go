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

// ListByType returns game definitions filtered by type
func (r *GameRegistry) ListByType(gameType string) []*GameDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*GameDefinition
	for _, game := range r.games {
		if game.Type == gameType {
			result = append(result, game)
		}
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

// GetActiveGames returns all active games
func (r *GameRegistry) GetActiveGames() []*GameDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*GameDefinition
	for _, game := range r.games {
		if game.Status == "active" {
			result = append(result, game)
		}
	}

	return result
}

// DefaultGames returns default game definitions
func DefaultGames() []*GameDefinition {
	return []*GameDefinition{
		{
			ID:          "blackjack",
			Name:        "Blackjack",
			Type:        "card",
			Subtype:     "blackjack",
			Description: "Classic blackjack game",
			RTP:         99.5,
			HouseEdge:   0.5,
			MinBet:      1,
			MaxBet:      10000,
			MaxPayout:   50000,
			Status:      "active",
		},
		{
			ID:          "baccarat",
			Name:        "Baccarat",
			Type:        "card",
			Subtype:     "baccarat",
			Description: "Classic baccarat game",
			RTP:         98.94,
			HouseEdge:   1.06,
			MinBet:      1,
			MaxBet:      10000,
			MaxPayout:   50000,
			Status:      "active",
		},
		{
			ID:          "roulette",
			Name:        "European Roulette",
			Type:        "table",
			Subtype:     "roulette",
			Description: "European roulette with single zero",
			RTP:         97.3,
			HouseEdge:   2.7,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   35000,
			Status:      "active",
		},
		{
			ID:          "slots_classic",
			Name:        "Classic Slots",
			Type:        "slot",
			Subtype:     "classic",
			Description: "3-reel classic slot machine",
			RTP:         95.5,
			HouseEdge:   4.5,
			MinBet:      1,
			MaxBet:      100,
			MaxPayout:   1000,
			Status:      "active",
		},
		{
			ID:          "slots_video",
			Name:        "Video Slots",
			Type:        "slot",
			Subtype:     "video",
			Description: "5-reel video slot machine",
			RTP:         96.5,
			HouseEdge:   3.5,
			MinBet:      1,
			MaxBet:      500,
			MaxPayout:   50000,
			Status:      "active",
		},
		{
			ID:          "craps",
			Name:        "Craps",
			Type:        "dice",
			Subtype:     "craps",
			Description: "Classic craps table game",
			RTP:         98.5,
			HouseEdge:   1.5,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   25000,
			Status:      "active",
		},
		{
			ID:          "sicbo",
			Name:        "Sic Bo",
			Type:        "dice",
			Subtype:     "sicbo",
			Description: "Chinese dice game",
			RTP:         97.5,
			HouseEdge:   2.5,
			MinBet:      1,
			MaxBet:      5000,
			MaxPayout:   25000,
			Status:      "active",
		},
		{
			ID:          "poker",
			Name:        "Texas Hold'em",
			Type:        "card",
			Subtype:     "poker",
			Description: "Texas Hold'em poker",
			RTP:         96.0,
			HouseEdge:   4.0,
			MinBet:      1,
			MaxBet:      10000,
			MaxPayout:   100000,
			Status:      "active",
		},
	}
}

// InitializeDefaultGames registers default games
func (r *GameRegistry) InitializeDefaultGames() {
	for _, game := range DefaultGames() {
		if err := r.Register(game); err != nil {
			continue
		}
	}
}
