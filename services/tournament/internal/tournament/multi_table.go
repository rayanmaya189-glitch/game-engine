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

// balanceTables balances players across tables
func (m *MultiTableTournament) balanceTables(ctx context.Context) {
	var tablesToBalance []*MTTTable
	minPlayers := m.Settings.SeatsPerTable

	for _, table := range m.Tables {
		if table.IsFinal {
			continue
		}
		if len(table.Players) < minPlayers {
			tablesToBalance = append(tablesToBalance, table)
		}
	}

	// Move players from tables with excess to tables with deficit
	for _, table := range m.Tables {
		if table.IsFinal || len(table.Players) <= minPlayers {
			continue
		}

		for _, player := range table.Players {
			if player.IsEliminated {
				continue
			}

			// Find a table with fewer players
			for _, targetTable := range tablesToBalance {
				if len(targetTable.Players) < targetTable.SeatCount {
					// Move player
					delete(table.Players, player.UserID)
					player.TableID = targetTable.TableID
					targetTable.Players[player.UserID] = player
					break
				}
			}
		}
	}
}

// createFinalTable creates the final table when enough players remain
func (m *MultiTableTournament) createFinalTable(ctx context.Context) {
	m.FinalTable = &MTTTable{
		TableID:     "final_table",
		TableNumber: len(m.Tables) + 1,
		Players:     make(map[string]*MTTPlayer),
		SeatCount:   m.Settings.FinalTableSeats,
		Level:       0,
		IsFinal:     true,
	}

	// Move top players to final table
	var topPlayers []*MTTPlayer
	for _, table := range m.Tables {
		for _, player := range table.Players {
			if !player.IsEliminated {
				topPlayers = append(topPlayers, player)
			}
		}
	}

	// Sort by chips (highest first) - simplified
	// In production, use proper sorting
	if len(topPlayers) > m.Settings.FinalTableSeats {
		topPlayers = topPlayers[:m.Settings.FinalTableSeats]
	}

	for _, player := range topPlayers {
		m.FinalTable.Players[player.UserID] = player
		player.TableID = "final_table"
		player.Position = len(m.FinalTable.Players)
	}
}

// countActivePlayers counts players still in the tournament
func (m *MultiTableTournament) countActivePlayers() int {
	count := 0
	for _, table := range m.Tables {
		for _, player := range table.Players {
			if !player.IsEliminated {
				count++
			}
		}
	}
	if m.FinalTable != nil {
		for _, player := range m.FinalTable.Players {
			if !player.IsEliminated {
				count++
			}
		}
	}
	return count
}

// EliminatePlayer removes a player from the tournament
func (m *MultiTableTournament) EliminatePlayer(ctx context.Context, userID string, rank int, prize int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Find player
	var player *MTTPlayer
	for _, table := range m.Tables {
		if p, exists := table.Players[userID]; exists {
			player = p
			break
		}
	}
	if m.FinalTable != nil {
		if p, exists := m.FinalTable.Players[userID]; exists {
			player = p
		}
	}

	if player == nil {
		return fmt.Errorf("player not found")
	}

	player.IsEliminated = true
	player.Rank = rank
	player.PrizeMoney = prize

	// Add to results
	m.Results = append(m.Results, MTTResult{
		Rank:       rank,
		UserID:     userID,
		PrizeMoney: prize,
	})

	m.UpdatedAt = time.Now()
	return m.saveToRedis(ctx)
}

// completeTournament finishes the tournament and calculates prizes
func (m *MultiTableTournament) completeTournament(ctx context.Context) {
	m.Status = TournamentStatusCompleted
	m.EndTime = time.Now()

	// Calculate prize distribution
	prizeDistribution := m.calculatePrizeDistribution()

	// Assign final positions and prizes
	rank := 1
	for _, table := range m.Tables {
		for _, player := range table.Players {
			if !player.IsEliminated {
				player.Rank = rank
				if rank-1 < len(prizeDistribution) {
					player.PrizeMoney = prizeDistribution[rank-1]
				}
				m.Results = append(m.Results, MTTResult{
					Rank:       rank,
					UserID:     player.UserID,
					PrizeMoney: player.PrizeMoney,
				})
				rank++
			}
		}
	}

	// Final table results
	if m.FinalTable != nil {
		for _, player := range m.FinalTable.Players {
			if !player.IsEliminated {
				player.Rank = rank
				if rank-1 < len(prizeDistribution) {
					player.PrizeMoney = prizeDistribution[rank-1]
				}
				m.Results = append(m.Results, MTTResult{
					Rank:       rank,
					UserID:     player.UserID,
					PrizeMoney: player.PrizeMoney,
				})
				rank++
			}
		}
	}
}

// calculatePrizeDistribution calculates prize amounts based on positions
func (m *MultiTableTournament) calculatePrizeDistribution() []int64 {
	// Standard distribution: top 15-20% of players get paid
	playerCount := m.countActivePlayers()
	if playerCount == 0 {
		return []int64{}
	}

	paidPlaces := playerCount / 5 // 20% get paid
	if paidPlaces < 1 {
		paidPlaces = 1
	}
	if paidPlaces > 10 {
		paidPlaces = 10
	}

	distribution := make([]int64, paidPlaces)

	// Tiered distribution (more to first place)
	for i := 0; i < paidPlaces; i++ {
		// Simple exponential distribution
		percentage := float64(paidPlaces - i)
		distribution[i] = int64(float64(m.PrizePool) * (percentage / float64(paidPlaces*paidPlaces/2)))
	}

	return distribution
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

// GetTableStates returns current state of all tables
func (m *MultiTableTournament) GetTableStates() map[string]interface{} {
	result := make(map[string]interface{})

	for tableID, table := range m.Tables {
		playerList := make([]string, 0, len(table.Players))
		for userID, player := range table.Players {
			if !player.IsEliminated {
				playerList = append(playerList, fmt.Sprintf("%s:%d", userID, player.Chips))
			}
		}
		result[tableID] = map[string]interface{}{
			"table_number": table.TableNumber,
			"players":      playerList,
			"level":        table.Level,
		}
	}

	return result
}

// GetLeaderboard returns current tournament standings
func (m *MultiTableTournament) GetLeaderboard() []map[string]interface{} {
	type playerScore struct {
		userID string
		chips  int
		rank   int
	}

	var players []playerScore
	for _, table := range m.Tables {
		for userID, player := range table.Players {
			players = append(players, playerScore{
				userID: userID,
				chips:  player.Chips,
				rank:   player.Rank,
			})
		}
	}

	// Sort by chips descending (simplified)
	// In production, use proper sorting

	result := make([]map[string]interface{}, 0, len(players))
	for _, p := range players {
		result = append(result, map[string]interface{}{
			"user_id": p.userID,
			"chips":   p.chips,
			"rank":    p.rank,
		})
	}

	return result
}

// GetPlayerInfo returns information about a specific player
func (m *MultiTableTournament) GetPlayerInfo(userID string) (map[string]interface{}, error) {
	for _, table := range m.Tables {
		if player, exists := table.Players[userID]; exists {
			return map[string]interface{}{
				"user_id":    player.UserID,
				"chips":      player.Chips,
				"rank":       player.Rank,
				"table_id":   player.TableID,
				"eliminated": player.IsEliminated,
				"prize":      player.PrizeMoney,
			}, nil
		}
	}
	if m.FinalTable != nil {
		if player, exists := m.FinalTable.Players[userID]; exists {
			return map[string]interface{}{
				"user_id":    player.UserID,
				"chips":      player.Chips,
				"rank":       player.Rank,
				"table_id":   player.TableID,
				"position":   player.Position,
				"eliminated": player.IsEliminated,
				"prize":      player.PrizeMoney,
			}, nil
		}
	}
	return nil, fmt.Errorf("player not found")
}

// AddChips adds chips to a player (for rebuys/addons)
func (m *MultiTableTournament) AddChips(ctx context.Context, userID string, amount int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, table := range m.Tables {
		if player, exists := table.Players[userID]; exists {
			player.Chips += amount
			m.UpdatedAt = time.Now()
			return m.saveToRedis(ctx)
		}
	}

	if m.FinalTable != nil {
		if player, exists := m.FinalTable.Players[userID]; exists {
			player.Chips += amount
			m.UpdatedAt = time.Now()
			return m.saveToRedis(ctx)
		}
	}

	return fmt.Errorf("player not found")
}

// UpdateChips updates a player's chip count
func (m *MultiTableTournament) UpdateChips(ctx context.Context, userID string, newChips int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if newChips <= 0 {
		// Player eliminated - find their rank
		rank := m.countActivePlayers()
		prize := int64(0)
		for _, r := range m.Results {
			if r.Rank == rank {
				prize = r.PrizeMoney
				break
			}
		}
		return m.EliminatePlayer(ctx, userID, rank, prize)
	}

	for _, table := range m.Tables {
		if player, exists := table.Players[userID]; exists {
			player.Chips = newChips
			m.UpdatedAt = time.Now()
			return m.saveToRedis(ctx)
		}
	}

	return fmt.Errorf("player not found")
}
