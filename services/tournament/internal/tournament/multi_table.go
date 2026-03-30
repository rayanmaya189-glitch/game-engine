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

// MultiTableTournament represents a multi-table tournament (MTT)
type MultiTableTournament struct {
	ID              string                    `json:"id"`
	Name            string                    `json:"name"`
	Type            TournamentType            `json:"type"`
	Status          TournamentStatus          `json:"status"`
	GameType        string                    `json:"game_type"`
	MaxPlayers      int                       `json:"max_players"`
	MinPlayers      int                       `json:"min_players"`
	CurrentPlayers  int                       `json:"current_players"`
	EntryFee        int64                     `json:"entry_fee"`
	BuyIn           int64                     `json:"buy_in"`
	PrizePool       int64                     `json:"prize_pool"`
	StartTime       time.Time                 `json:"start_time"`
	EndTime         time.Time                 `json:"end_time"`
	CreatedAt       time.Time                 `json:"created_at"`
	UpdatedAt       time.Time                 `json:"updated_at"`
	Settings        MultiTableSettings        `json:"settings"`
	Tables          map[string]*MTTTable      `json:"tables"`
	FinalTable      *MTTTable                 `json:"final_table,omitempty"`
	RegisteredUsers map[string]MTTParticipant `json:"registered_users"`
	Results         []MTTResult               `json:"results"`
	RedisClient     *redis.Client
	mu              sync.RWMutex
}

// MultiTableSettings contains settings specific to multi-table tournaments
type MultiTableSettings struct {
	Tables                int  `json:"tables"`                  // Number of tables in initial round
	SeatsPerTable         int  `json:"seats_per_table"`         // Max seats per table
	TableBalanceThreshold int  `json:"table_balance_threshold"` // When to balance tables
	FinalTableSeats       int  `json:"final_table_seats"`       // Final table seat count
	BreakDuration         int  `json:"break_duration"`          // Break between levels (seconds)
	PlayUntilTableCount   int  `json:"play_until_table_count"`  // Number of tables to reach before final table
	AutoBalance           bool `json:"auto_balance"`            // Automatically balance tables
	ShuffleSeats          bool `json:"shuffle_seats"`           // Shuffle seating
	StartingChips         int  `json:"starting_chips"`          // Starting chips for players
}

// MTTTable represents a table in a multi-table tournament
type MTTTable struct {
	TableID     string                `json:"table_id"`
	TableNumber int                   `json:"table_number"`
	Players     map[string]*MTTPlayer `json:"players"`
	SeatCount   int                   `json:"seat_count"`
	Level       int                   `json:"level"`
	IsFinal     bool                  `json:"is_final"`
}

// MTTPlayer represents a player in a multi-table tournament
type MTTPlayer struct {
	UserID       string `json:"user_id"`
	Chips        int    `json:"chips"`
	Rank         int    `json:"rank"`
	TableID      string `json:"table_id"`
	IsEliminated bool   `json:"is_eliminated"`
	Position     int    `json:"position"` // Final table position
	PrizeMoney   int64  `json:"prize_money"`
}

// MTTParticipant represents a registered participant
type MTTParticipant struct {
	UserID       string    `json:"user_id"`
	RegisteredAt time.Time `json:"registered_at"`
	CheckedIn    bool      `json:"checked_in"`
	TableID      string    `json:"table_id,omitempty"`
}

// MTTResult represents a final tournament result
type MTTResult struct {
	Rank       int    `json:"rank"`
	UserID     string `json:"user_id"`
	PrizeMoney int64  `json:"prize_money"`
	Points     int    `json:"points"`
}

// NewMultiTableTournament creates a new multi-table tournament
func NewMultiTableTournament(name string, gameType string, maxPlayers int, buyIn int64, settings MultiTableSettings) *MultiTableTournament {
	return &MultiTableTournament{
		ID:              uuid.New().String(),
		Name:            name,
		Type:            TournamentTypeScheduled,
		Status:          TournamentStatusPending,
		GameType:        gameType,
		MaxPlayers:      maxPlayers,
		MinPlayers:      settings.Tables * 2,
		CurrentPlayers:  0,
		BuyIn:           buyIn,
		Settings:        settings,
		Tables:          make(map[string]*MTTTable),
		RegisteredUsers: make(map[string]MTTParticipant),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// InitializeTables creates the initial tournament tables
func (m *MultiTableTournament) InitializeTables(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.Status != TournamentStatusPending {
		return fmt.Errorf("tournament is not in pending state")
	}

	// Create initial tables
	for i := 1; i <= m.Settings.Tables; i++ {
		tableID := fmt.Sprintf("table_%d", i)
		m.Tables[tableID] = &MTTTable{
			TableID:     tableID,
			TableNumber: i,
			Players:     make(map[string]*MTTPlayer),
			SeatCount:   m.Settings.SeatsPerTable,
			Level:       1,
			IsFinal:     false,
		}
	}

	m.Status = TournamentStatusRegistering
	m.UpdatedAt = time.Now()

	// Store in Redis for persistence
	return m.saveToRedis(ctx)
}

// RegisterPlayer adds a player to the tournament
func (m *MultiTableTournament) RegisterPlayer(ctx context.Context, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.Status != TournamentStatusRegistering {
		return fmt.Errorf("registration is not open")
	}

	if m.CurrentPlayers >= m.MaxPlayers {
		return fmt.Errorf("tournament is full")
	}

	if _, exists := m.RegisteredUsers[userID]; exists {
		return fmt.Errorf("player already registered")
	}

	// Add to registered users
	m.RegisteredUsers[userID] = MTTParticipant{
		UserID:       userID,
		RegisteredAt: time.Now(),
		CheckedIn:    false,
	}
	m.CurrentPlayers++
	m.UpdatedAt = time.Now()

	return m.saveToRedis(ctx)
}

// CheckIn seats a player at a table after they check in
func (m *MultiTableTournament) CheckIn(ctx context.Context, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.Status != TournamentStatusRegistering {
		return fmt.Errorf("check-in is not open")
	}

	participant, exists := m.RegisteredUsers[userID]
	if !exists {
		return fmt.Errorf("player not registered")
	}

	if participant.CheckedIn {
		return fmt.Errorf("player already checked in")
	}

	// Find a table with available seats
	var tableID string
	for id, table := range m.Tables {
		if len(table.Players) < table.SeatCount {
			tableID = id
			break
		}
	}

	if tableID == "" {
		return fmt.Errorf("no available seats")
	}
	// Create player
	player := &MTTPlayer{
		UserID:       userID,
		Chips:        m.Settings.StartingChips,
		Rank:         m.CurrentPlayers,
		TableID:      tableID,
		IsEliminated: false,
	}

	// Add to table
	m.Tables[tableID].Players[userID] = player
	participant.CheckedIn = true
	participant.TableID = tableID
	m.RegisteredUsers[userID] = participant

	m.UpdatedAt = time.Now()
	return m.saveToRedis(ctx)
}

// StartTournament begins the tournament
func (m *MultiTableTournament) StartTournament(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.Status != TournamentStatusRegistering {
		return fmt.Errorf("tournament cannot be started")
	}

	if m.CurrentPlayers < m.MinPlayers {
		return fmt.Errorf("not enough players registered")
	}

	m.Status = TournamentStatusRunning
	m.StartTime = time.Now()
	m.UpdatedAt = time.Now()

	// Start blind level timer
	go m.runBlindLevels(ctx)

	return m.saveToRedis(ctx)
}

// runBlindLevels manages tournament blind levels
func (m *MultiTableTournament) runBlindLevels(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(m.Settings.BreakDuration) * time.Second)
	defer ticker.Stop()

	currentLevel := 1
	for {
		select {
		case <-ticker.C:
			m.mu.Lock()
			if m.Status != TournamentStatusRunning {
				m.mu.Unlock()
				return
			}

			// Progress to next blind level
			currentLevel++
			for _, table := range m.Tables {
				table.Level = currentLevel
			}

			// Check for table merge conditions
			if m.Settings.AutoBalance {
				m.balanceTables(ctx)
			}

			// Check if we should reach final table
			if len(m.Tables) <= m.Settings.PlayUntilTableCount && m.FinalTable == nil {
				m.createFinalTable(ctx)
			}

			// Check for tournament completion
			remainingPlayers := m.countActivePlayers()
			if remainingPlayers <= m.Settings.FinalTableSeats {
				m.completeTournament(ctx)
			}

			m.UpdatedAt = time.Now()
			m.saveToRedis(ctx)
			m.mu.Unlock()

		case <-ctx.Done():
			return
		}
	}
}

// saveToRedis persists tournament state
func (m *MultiTableTournament) saveToRedis(ctx context.Context) error {
	if m.RedisClient == nil {
		return nil // No Redis configured
	}

	key := fmt.Sprintf("tournament:mtt:%s", m.ID)
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return m.RedisClient.Set(ctx, key, data, 24*time.Hour).Err()
}
