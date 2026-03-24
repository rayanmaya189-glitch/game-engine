package tournament

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// TournamentType represents the type of tournament
type TournamentType string

const (
	TournamentTypeSitAndGo  TournamentType = "sit_and_go"
	TournamentTypeScheduled TournamentType = "scheduled"
	TournamentTypeKnockout  TournamentType = "knockout"
	TournamentTypeFreeroll  TournamentType = "freeroll"
)

// TournamentStatus represents the status of a tournament
type TournamentStatus string

const (
	TournamentStatusPending     TournamentStatus = "pending"
	TournamentStatusRegistering TournamentStatus = "registering"
	TournamentStatusRunning     TournamentStatus = "running"
	TournamentStatusCompleted   TournamentStatus = "completed"
	TournamentStatusCancelled   TournamentStatus = "cancelled"
)

// Tournament represents a tournament in the system
type Tournament struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	Type            TournamentType         `json:"type"`
	Status          TournamentStatus       `json:"status"`
	GameType        string                 `json:"game_type"`
	MinPlayers      int                    `json:"min_players"`
	MaxPlayers      int                    `json:"max_players"`
	CurrentPlayers  int                    `json:"current_players"`
	EntryFee        int64                  `json:"entry_fee"`
	BuyIn           int64                  `json:"buy_in"`
	PrizePool       int64                  `json:"prize_pool"`
	StartTime       time.Time              `json:"start_time"`
	EndTime         time.Time              `json:"end_time"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	Settings        TournamentSettings     `json:"settings"`
	RegisteredUsers map[string]Participant `json:"registered_users"`
	Results         []Result               `json:"results"`
	Bracket         *Bracket               `json:"bracket,omitempty"`
}

// TournamentSettings contains tournament-specific settings
type TournamentSettings struct {
	AutoStart         bool         `json:"auto_start"`
	RebuyEnabled      bool         `json:"rebuy_enabled"`
	AddonEnabled      bool         `json:"addon_enabled"`
	Reentries         int          `json:"reentries"`
	LateRegistration  int          `json:"late_registration"` // seconds
	StartingChips     int          `json:"starting_chips"`
	BlindLevels       []BlindLevel `json:"blind_levels"`
	PrizeDistribution []int        `json:"prize_distribution"`
}

// BlindLevel represents a blind level in a tournament
type BlindLevel struct {
	Level      int `json:"level"`
	SmallBlind int `json:"small_blind"`
	BigBlind   int `json:"big_blind"`
	Duration   int `json:"duration"` // seconds
}

// Participant represents a player in a tournament
type Participant struct {
	UserID       string    `json:"user_id"`
	Username     string    `json:"username"`
	Chips        int       `json:"chips"`
	Position     int       `json:"position"`
	Status       string    `json:"status"` // active, eliminated, registered
	Score        int       `json:"score"`
	Rank         int       `json:"rank"`
	Knockouts    int       `json:"knockouts"`
	PrizeWon     int64     `json:"prize_won"`
	RegisteredAt time.Time `json:"registered_at"`
}

// Result represents a tournament result
type Result struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Rank      int    `json:"rank"`
	Prize     int64  `json:"prize"`
	Knockouts int    `json:"knockouts"`
	Score     int    `json:"score"`
}

// Bracket represents a knockout tournament bracket
type Bracket struct {
	Rounds []Round `json:"rounds"`
}

// Round represents a single round in a bracket
type Round struct {
	Number   int     `json:"number"`
	Matches  []Match `json:"matches"`
	Complete bool    `json:"complete"`
}

// Match represents a single match in the bracket
type Match struct {
	ID       string  `json:"id"`
	Player1  *string `json:"player1"`
	Player2  *string `json:"player2"`
	Winner   *string `json:"winner"`
	Score1   int     `json:"score1"`
	Score2   int     `json:"score2"`
	Complete bool    `json:"complete"`
}

// Manager handles tournament lifecycle
type Manager struct {
	mu          sync.RWMutex
	tournaments map[string]*Tournament
	RedisClient *redis.Client
	Config      *Config
	Leaderboard *Leaderboard
	PrizePool   *PrizePool
	Scheduler   *Scheduler
}

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

// RegisterUser registers a user for a tournament
func (m *Manager) RegisterUser(ctx context.Context, tournamentID string, userID, username string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	tournament, ok := m.tournaments[tournamentID]
	if !ok {
		return fmt.Errorf("tournament not found: %s", tournamentID)
	}

	if tournament.Status != TournamentStatusPending && tournament.Status != TournamentStatusRegistering {
		return fmt.Errorf("tournament is not accepting registrations")
	}

	if tournament.CurrentPlayers >= tournament.MaxPlayers {
		return fmt.Errorf("tournament is full")
	}

	if _, exists := tournament.RegisteredUsers[userID]; exists {
		return fmt.Errorf("user already registered")
	}

	participant := Participant{
		UserID:       userID,
		Username:     username,
		Chips:        tournament.Settings.StartingChips,
		Status:       "registered",
		RegisteredAt: time.Now(),
	}

	tournament.RegisteredUsers[userID] = participant
	tournament.CurrentPlayers++
	tournament.PrizePool = int64(tournament.CurrentPlayers) * tournament.EntryFee
	tournament.UpdatedAt = time.Now()

	// Update status to registering if we have at least min players
	if tournament.CurrentPlayers >= tournament.MinPlayers && tournament.Status == TournamentStatusPending {
		tournament.Status = TournamentStatusRegistering
	}

	// Save to Redis
	if err := m.saveTournament(ctx, tournament); err != nil {
		return err
	}

	// Update leaderboard
	m.Leaderboard.UpdateRegistration(ctx, tournamentID, participant)

	// Auto-start if enabled
	if tournament.Settings.AutoStart && tournament.Status == TournamentStatusRegistering {
		if tournament.Type == TournamentTypeSitAndGo {
			go m.startTournament(ctx, tournamentID)
		}
	}

	return nil
}

// UnregisterUser unregisters a user from a tournament
func (m *Manager) UnregisterUser(ctx context.Context, tournamentID string, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	tournament, ok := m.tournaments[tournamentID]
	if !ok {
		return fmt.Errorf("tournament not found: %s", tournamentID)
	}

	if tournament.Status == TournamentStatusRunning {
		return fmt.Errorf("cannot unregister while tournament is running")
	}

	participant, exists := tournament.RegisteredUsers[userID]
	if !exists {
		return fmt.Errorf("user not registered")
	}

	delete(tournament.RegisteredUsers, userID)
	tournament.CurrentPlayers--
	tournament.PrizePool = int64(tournament.CurrentPlayers) * tournament.EntryFee
	tournament.UpdatedAt = time.Now()

	// Revert status if below minimum
	if tournament.CurrentPlayers < tournament.MinPlayers && tournament.Status == TournamentStatusRegistering {
		tournament.Status = TournamentStatusPending
	}

	// Save to Redis
	if err := m.saveTournament(ctx, tournament); err != nil {
		return err
	}

	// Remove from leaderboard
	m.Leaderboard.RemoveParticipant(ctx, tournamentID, userID)
	_ = participant // Use participant for refund logic

	return nil
}

// StartTournament starts a tournament
func (m *Manager) StartTournament(ctx context.Context, tournamentID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.startTournament(ctx, tournamentID)
}

func (m *Manager) startTournament(ctx context.Context, tournamentID string) error {
	tournament, ok := m.tournaments[tournamentID]
	if !ok {
		return fmt.Errorf("tournament not found: %s", tournamentID)
	}

	if tournament.CurrentPlayers < tournament.MinPlayers {
		return fmt.Errorf("not enough players to start tournament")
	}

	tournament.Status = TournamentStatusRunning
	if tournament.StartTime.IsZero() {
		tournament.StartTime = time.Now()
	}
	tournament.UpdatedAt = time.Now()

	// Generate bracket for knockout tournaments
	if tournament.Type == TournamentTypeKnockout {
		tournament.Bracket = m.generateBracket(tournament)
	}

	// Initialize leaderboard
	for _, participant := range tournament.RegisteredUsers {
		m.Leaderboard.UpdateScore(ctx, tournamentID, participant.UserID, participant.Score)
	}

	// Save to Redis
	if err := m.saveTournament(context.Background(), tournament); err != nil {
		return err
	}

	return nil
}

// EndTournament ends a tournament and calculates results
func (m *Manager) EndTournament(ctx context.Context, tournamentID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	tournament, ok := m.tournaments[tournamentID]
	if !ok {
		return fmt.Errorf("tournament not found: %s", tournamentID)
	}

	if tournament.Status != TournamentStatusRunning {
		return fmt.Errorf("tournament is not running")
	}

	tournament.Status = TournamentStatusCompleted
	tournament.EndTime = time.Now()
	tournament.UpdatedAt = time.Now()

	// Calculate prizes
	results := m.PrizePool.CalculatePrizes(tournament)
	tournament.Results = results

	// Update leaderboard with final results
	for _, result := range results {
		m.Leaderboard.UpdateFinalResult(ctx, tournamentID, result)
	}

	// Save to Redis
	if err := m.saveTournament(ctx, tournament); err != nil {
		return err
	}

	return nil
}

// UpdatePlayerScore updates a player's score in a tournament
func (m *Manager) UpdatePlayerScore(ctx context.Context, tournamentID string, userID string, scoreDelta int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	tournament, ok := m.tournaments[tournamentID]
	if !ok {
		return fmt.Errorf("tournament not found: %s", tournamentID)
	}

	participant, exists := tournament.RegisteredUsers[userID]
	if !exists {
		return fmt.Errorf("player not found in tournament")
	}

	participant.Score += scoreDelta
	tournament.RegisteredUsers[userID] = participant
	tournament.UpdatedAt = time.Now()

	// Update leaderboard
	m.Leaderboard.UpdateScore(ctx, tournamentID, userID, participant.Score)

	// Save to Redis
	return m.saveTournament(ctx, tournament)
}

// EliminatePlayer eliminates a player from the tournament
func (m *Manager) EliminatePlayer(ctx context.Context, tournamentID string, userID string, rank int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	tournament, ok := m.tournaments[tournamentID]
	if !ok {
		return fmt.Errorf("tournament not found: %s", tournamentID)
	}

	participant, exists := tournament.RegisteredUsers[userID]
	if !exists {
		return fmt.Errorf("player not found in tournament")
	}

	participant.Status = "eliminated"
	participant.Rank = rank
	tournament.RegisteredUsers[userID] = participant
	tournament.UpdatedAt = time.Now()

	// Update leaderboard
	m.Leaderboard.UpdateElimination(ctx, tournamentID, userID, rank)

	// Check if tournament should end
	activePlayers := 0
	for _, p := range tournament.RegisteredUsers {
		if p.Status == "registered" || p.Status == "active" {
			activePlayers++
		}
	}

	if activePlayers <= 1 {
		tournament.Status = TournamentStatusCompleted
		tournament.EndTime = time.Now()

		// Calculate prizes
		results := m.PrizePool.CalculatePrizes(tournament)
		tournament.Results = results
	}

	// Save to Redis
	return m.saveTournament(ctx, tournament)
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

// generateBracket generates a bracket for knockout tournaments
func (m *Manager) generateBracket(tournament *Tournament) *Bracket {
	participants := make([]string, 0, len(tournament.RegisteredUsers))
	for userID := range tournament.RegisteredUsers {
		participants = append(participants, userID)
	}

	// Simple shuffle for seeding
	// In production, use proper randomization
	numRounds := 0
	for len(participants) > 1 {
		numRounds++
		participants = participants[:len(participants)/2]
	}

	bracket := &Bracket{
		Rounds: make([]Round, numRounds),
	}

	for i := 0; i < numRounds; i++ {
		numMatches := len(tournament.RegisteredUsers) / (1 << (i + 1))
		matches := make([]Match, numMatches)
		for j := 0; j < numMatches; j++ {
			matches[j] = Match{
				ID:       uuid.New().String(),
				Player1:  nil,
				Player2:  nil,
				Complete: false,
			}
		}
		bracket.Rounds[i] = Round{
			Number:   i + 1,
			Matches:  matches,
			Complete: false,
		}
	}

	return bracket
}

// generateDefaultBlindLevels generates default blind structure
func generateDefaultBlindLevels() []BlindLevel {
	levels := []BlindLevel{
		{Level: 1, SmallBlind: 10, BigBlind: 20, Duration: 900},
		{Level: 2, SmallBlind: 15, BigBlind: 30, Duration: 900},
		{Level: 3, SmallBlind: 25, BigBlind: 50, Duration: 900},
		{Level: 4, SmallBlind: 50, BigBlind: 100, Duration: 900},
		{Level: 5, SmallBlind: 75, BigBlind: 150, Duration: 900},
		{Level: 6, SmallBlind: 100, BigBlind: 200, Duration: 900},
		{Level: 7, SmallBlind: 150, BigBlind: 300, Duration: 900},
		{Level: 8, SmallBlind: 200, BigBlind: 400, Duration: 900},
		{Level: 9, SmallBlind: 300, BigBlind: 600, Duration: 900},
		{Level: 10, SmallBlind: 400, BigBlind: 800, Duration: 900},
	}
	return levels
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
