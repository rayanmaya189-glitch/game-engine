package registry

import (
	"context"
	"errors"
	"sync"
	"time"
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

// GameRegistryClient defines the interface for fetching games from external sources
type GameRegistryClient interface {
	ListGames(ctx context.Context) ([]*GameDefinition, error)
	GetGame(ctx context.Context, id string) (*GameDefinition, error)
}

// RegistryConfig holds configuration for the game registry
type RegistryConfig struct {
	// Fallback defaults when registry is unavailable
	DefaultMinBet    int64
	DefaultMaxBet    int64
	DefaultMaxPayout int64
	DefaultRTP       float64
	DefaultHouseEdge float64
}

// DefaultRegistryConfig returns the default registry configuration
func DefaultRegistryConfig() *RegistryConfig {
	return &RegistryConfig{
		DefaultMinBet:    1,
		DefaultMaxBet:    10000,
		DefaultMaxPayout: 50000,
		DefaultRTP:       96.0,
		DefaultHouseEdge: 4.0,
	}
}

// GameRegistry manages game definitions
type GameRegistry struct {
	games  map[string]*GameDefinition
	mu     sync.RWMutex
	client GameRegistryClient
	cfg    *RegistryConfig
}

// NewGameRegistry creates a new game registry
func NewGameRegistry() *GameRegistry {
	return &GameRegistry{
		games: make(map[string]*GameDefinition),
		cfg:   DefaultRegistryConfig(),
	}
}

// NewGameRegistryWithClient creates a game registry backed by a gRPC client
func NewGameRegistryWithClient(client GameRegistryClient, cfg *RegistryConfig) *GameRegistry {
	if cfg == nil {
		cfg = DefaultRegistryConfig()
	}
	return &GameRegistry{
		games:  make(map[string]*GameDefinition),
		client: client,
		cfg:    cfg,
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

// InitializeDefaultGames loads games from the registry client, falling back to empty registry
func (r *GameRegistry) InitializeDefaultGames() {
	if r.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		games, err := r.client.ListGames(ctx)
		if err == nil {
			for _, game := range games {
				if err := r.Register(game); err != nil {
					continue
				}
			}
			return
		}
	}
	// No client available or fetch failed - registry starts empty
}

// ReloadGames refreshes the game list from the registry client
func (r *GameRegistry) ReloadGames(ctx context.Context) error {
	if r.client == nil {
		return errors.New("no registry client configured")
	}

	games, err := r.client.ListGames(ctx)
	if err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.games = make(map[string]*GameDefinition)
	for _, game := range games {
		if game.ID != "" {
			r.games[game.ID] = game
		}
	}
	return nil
}

// GetClient returns the underlying registry client
func (r *GameRegistry) GetClient() GameRegistryClient {
	return r.client
}
