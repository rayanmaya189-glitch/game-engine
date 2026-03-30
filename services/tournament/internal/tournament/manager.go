package tournament

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// NewManager creates a new tournament manager
func NewManager(config *Config, redisClient *redis.Client) (*Manager, error) {
	m := &Manager{
		tournaments: make(map[string]*Tournament),
		RedisClient: redisClient,
		Config:      config,
		Leaderboard: NewLeaderboard(redisClient),
		PrizePool:   NewPrizePool(redisClient),
		Scheduler:   NewScheduler(redisClient),
	}

	// Load existing tournaments from Redis
	if err := m.loadTournaments(context.Background()); err != nil {
		return nil, err
	}

	return m, nil
}

// CreateTournament creates a new tournament
func (m *Manager) CreateTournament(ctx context.Context, name string, tType TournamentType, gameType string, entryFee int64, buyIn int64, minPlayers, maxPlayers int, startTime time.Time, settings TournamentSettings) (*Tournament, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Apply type-specific defaults
	switch tType {
	case TournamentTypeSitAndGo:
		if settings.AutoStart && startTime.IsZero() {
			startTime = time.Now().Add(10 * time.Second)
		}
		if minPlayers == 0 {
			minPlayers = 2
		}
		if maxPlayers == 0 {
			maxPlayers = 10
		}
	case TournamentTypeScheduled:
		if minPlayers == 0 {
			minPlayers = 2
		}
		if maxPlayers == 0 {
			maxPlayers = 1000
		}
	case TournamentTypeKnockout:
		if minPlayers == 0 {
			minPlayers = 4
		}
		if maxPlayers == 0 {
			maxPlayers = 100
		}
	case TournamentTypeFreeroll:
		entryFee = 0
		if minPlayers == 0 {
			minPlayers = 2
		}
	}

	tournament := &Tournament{
		ID:              uuid.New().String(),
		Name:            name,
		Type:            tType,
		Status:          TournamentStatusPending,
		GameType:        gameType,
		MinPlayers:      minPlayers,
		MaxPlayers:      maxPlayers,
		CurrentPlayers:  0,
		EntryFee:        entryFee,
		BuyIn:           buyIn,
		PrizePool:       0,
		StartTime:       startTime,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		Settings:        settings,
		RegisteredUsers: make(map[string]Participant),
		Results:         make([]Result, 0),
	}

	// Set default blind levels if not provided
	if len(settings.BlindLevels) == 0 {
		tournament.Settings.BlindLevels = generateDefaultBlindLevels()
	}

	// Set default prize distribution if not provided
	if len(settings.PrizeDistribution) == 0 {
		tournament.Settings.PrizeDistribution = []int{50, 30, 20}
	}

	m.tournaments[tournament.ID] = tournament

	// Save to Redis
	if err := m.saveTournament(ctx, tournament); err != nil {
		return nil, err
	}

	// Schedule the tournament if it's a scheduled type
	if tType == TournamentTypeScheduled && !startTime.IsZero() {
		m.Scheduler.ScheduleTournament(ctx, tournament.ID, startTime)
	}

	return tournament, nil
}

// GetTournament retrieves a tournament by ID
func (m *Manager) GetTournament(ctx context.Context, id string) (*Tournament, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	tournament, ok := m.tournaments[id]
	if !ok {
		// Try to load from Redis
		data, err := m.RedisClient.Get(ctx, fmt.Sprintf("tournament:%s", id)).Bytes()
		if err != nil {
			return nil, fmt.Errorf("tournament not found: %s", id)
		}

		var t Tournament
		if err := json.Unmarshal(data, &t); err != nil {
			return nil, err
		}
		return &t, nil
	}

	return tournament, nil
}

// ListTournaments returns all tournaments with optional filters
func (m *Manager) ListTournaments(ctx context.Context, status TournamentStatus, tType TournamentType) ([]*Tournament, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var result []*Tournament
	for _, t := range m.tournaments {
		if status != "" && t.Status != status {
			continue
		}
		if tType != "" && t.Type != tType {
			continue
		}
		result = append(result, t)
	}

	return result, nil
}

// saveTournament saves a tournament to Redis
func (m *Manager) saveTournament(ctx context.Context, tournament *Tournament) error {
	data, err := json.Marshal(tournament)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("tournament:%s", tournament.ID)
	return m.RedisClient.Set(ctx, key, data, 0).Err()
}

// loadTournaments loads tournaments from Redis
func (m *Manager) loadTournaments(ctx context.Context) error {
	keys, err := m.RedisClient.Keys(ctx, "tournament:*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		data, err := m.RedisClient.Get(ctx, key).Bytes()
		if err != nil {
			continue
		}

		var tournament Tournament
		if err := json.Unmarshal(data, &tournament); err != nil {
			continue
		}

		m.tournaments[tournament.ID] = &tournament
	}

	return nil
}
