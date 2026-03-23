package room

import (
	"context"
	"sort"
	"sync"
	"time"
)

// MatchmakingQueue represents a matchmaking queue
type MatchmakingQueue struct {
	mu        sync.RWMutex
	players   map[string]*MatchmakingPlayer
	waitTimes map[string]time.Time
	rankings  []string // Sorted player IDs by skill
}

// MatchmakingPlayer represents a player in matchmaking
type MatchmakingPlayer struct {
	UserID         string
	Username       string
	SkillRating    int // ELO-style rating
	StakeLevel     int // 1=Micro, 2=Low, 3=Medium, 4=High, 5=VIP
	GameType       string
	WaitStart      time.Time
	PreferredTable string // Specific table they want to join
}

// NewMatchmakingQueue creates a new matchmaking queue
func NewMatchmakingQueue() *MatchmakingQueue {
	return &MatchmakingQueue{
		players:   make(map[string]*MatchmakingPlayer),
		waitTimes: make(map[string]time.Time),
		rankings:  make([]string, 0),
	}
}

// AddPlayer adds a player to the matchmaking queue
func (mq *MatchmakingQueue) AddPlayer(ctx context.Context, player *MatchmakingPlayer) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if _, exists := mq.players[player.UserID]; exists {
		return nil // Already in queue
	}

	mq.players[player.UserID] = player
	mq.waitTimes[player.UserID] = time.Now()
	mq.recalculateRankings()

	return nil
}

// RemovePlayer removes a player from the queue
func (mq *MatchmakingQueue) RemovePlayer(ctx context.Context, userID string) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	delete(mq.players, userID)
	delete(mq.waitTimes, userID)
	mq.recalculateRankings()
}

// FindMatch finds a suitable match for a player
func (mq *MatchmakingQueue) FindMatch(ctx context.Context, player *MatchmakingPlayer) []*MatchmakingPlayer {
	mq.mu.RLock()
	defer mq.mu.RUnlock()

	var candidates []*MatchmakingPlayer

	// Filter by game type and stake level
	for _, p := range mq.players {
		if p.UserID == player.UserID {
			continue
		}

		if p.GameType != player.GameType || p.StakeLevel != player.StakeLevel {
			continue
		}

		// Skill rating should be within range (max 200 points difference)
		if abs(p.SkillRating-player.SkillRating) > 200 {
			continue
		}

		candidates = append(candidates, p)
	}

	// Sort by wait time (oldest first)
	sort.Slice(candidates, func(i, j int) bool {
		return mq.waitTimes[candidates[i].UserID].Before(mq.waitTimes[candidates[j].UserID])
	})

	// Return top matches (up to 9 for multiplayer games)
	if len(candidates) > 9 {
		candidates = candidates[:9]
	}

	return candidates
}

// GetQueueInfo returns current queue information
func (mq *MatchmakingQueue) GetQueueInfo(gameType string, stakeLevel int) (int, time.Duration) {
	mq.mu.RLock()
	defer mq.mu.RUnlock()

	var count int
	var oldest time.Duration

	for _, p := range mq.players {
		if p.GameType == gameType && p.StakeLevel == stakeLevel {
			count++
			wait := time.Since(mq.waitTimes[p.UserID])
			if oldest == 0 || wait > oldest {
				oldest = wait
			}
		}
	}

	return count, oldest
}

// UpdateSkillRating updates a player's skill rating
func (mq *MatchmakingQueue) UpdateSkillRating(userID string, newRating int) {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if player, exists := mq.players[userID]; exists {
		player.SkillRating = newRating
		mq.recalculateRankings()
	}
}

// recalculateRankings recalculates the sorted rankings
func (mq *MatchmakingQueue) recalculateRankings() {
	mq.rankings = make([]string, 0, len(mq.players))
	for userID := range mq.players {
		mq.rankings = append(mq.rankings, userID)
	}

	sort.Slice(mq.rankings, func(i, j int) bool {
		return mq.players[mq.rankings[i]].SkillRating > mq.players[mq.rankings[j]].SkillRating
	})
}

// MatchmakingConfig represents matchmaking configuration
type MatchmakingConfig struct {
	MaxSkillDiff     int           // Maximum skill difference for matches
	MaxWaitTime      time.Duration // Maximum time to wait for a match
	MinPlayers       int           // Minimum players to start
	MaxPlayers       int           // Maximum players in a match
	AutoStartEnabled bool          // Auto-start when min players reached
}

// DefaultMatchmakingConfig returns default configuration
func DefaultMatchmakingConfig() *MatchmakingConfig {
	return &MatchmakingConfig{
		MaxSkillDiff:     200,
		MaxWaitTime:      60 * time.Second,
		MinPlayers:       2,
		MaxPlayers:       9,
		AutoStartEnabled: true,
	}
}

// Matchmaker handles matchmaking between players
type Matchmaker struct {
	queues map[string]*MatchmakingQueue // Key: gameType_stakeLevel
	config *MatchmakingConfig
	mu     sync.RWMutex
}

// NewMatchmaker creates a new matchmaker
func NewMatchmaker(config *MatchmakingConfig) *Matchmaker {
	if config == nil {
		config = DefaultMatchmakingConfig()
	}

	return &Matchmaker{
		queues: make(map[string]*MatchmakingQueue),
		config: config,
	}
}

// getQueueKey generates queue key from game type and stake level
func (m *Matchmaker) getQueueKey(gameType string, stakeLevel int) string {
	return gameType + "_" + string(rune('0'+stakeLevel))
}

// JoinQueue adds a player to the matchmaking queue
func (m *Matchmaker) JoinQueue(ctx context.Context, userID, username, gameType string, skillRating, stakeLevel int) error {
	key := m.getQueueKey(gameType, stakeLevel)

	m.mu.Lock()
	defer m.mu.Unlock()

	queue, ok := m.queues[key]
	if !ok {
		queue = NewMatchmakingQueue()
		m.queues[key] = queue
	}

	player := &MatchmakingPlayer{
		UserID:      userID,
		Username:    username,
		SkillRating: skillRating,
		StakeLevel:  stakeLevel,
		GameType:    gameType,
		WaitStart:   time.Now(),
	}

	return queue.AddPlayer(ctx, player)
}

// LeaveQueue removes a player from the matchmaking queue
func (m *Matchmaker) LeaveQueue(ctx context.Context, userID, gameType string, stakeLevel int) {
	key := m.getQueueKey(gameType, stakeLevel)

	m.mu.Lock()
	defer m.mu.Unlock()

	if queue, ok := m.queues[key]; ok {
		queue.RemovePlayer(ctx, userID)
	}
}

// FindMatches finds matching players for a given player
func (m *Matchmaker) FindMatches(ctx context.Context, userID, username, gameType string, skillRating, stakeLevel int) []string {
	key := m.getQueueKey(gameType, stakeLevel)

	m.mu.RLock()
	queue, ok := m.queues[key]
	m.mu.RUnlock()

	if !ok {
		return nil
	}

	player := &MatchmakingPlayer{
		UserID:      userID,
		Username:    username,
		SkillRating: skillRating,
		StakeLevel:  stakeLevel,
		GameType:    gameType,
		WaitStart:   time.Now(),
	}

	matches := queue.FindMatch(ctx, player)

	result := make([]string, len(matches))
	for i, m := range matches {
		result[i] = m.UserID
	}

	return result
}

// AutoMatch attempts to automatically form a game
func (m *Matchmaker) AutoMatch(ctx context.Context, gameType string, stakeLevel int) []string {
	key := m.getQueueKey(gameType, stakeLevel)

	m.mu.RLock()
	queue, ok := m.queues[key]
	m.mu.RUnlock()

	if !ok {
		return nil
	}

	// Get all waiting players
	queue.mu.RLock()
	var players []*MatchmakingPlayer
	for _, p := range queue.players {
		players = append(players, p)
	}
	queue.mu.RUnlock()

	if len(players) < m.config.MinPlayers {
		return nil
	}

	// Try to form a group within skill range
	var matched []string
	skillSum := 0

	for _, p := range players {
		if len(matched) >= m.config.MaxPlayers {
			break
		}

		if len(matched) == 0 {
			matched = append(matched, p.UserID)
			skillSum = p.SkillRating
		} else {
			avgSkill := skillSum / len(matched)
			if abs(p.SkillRating-avgSkill) <= m.config.MaxSkillDiff {
				matched = append(matched, p.UserID)
				skillSum += p.SkillRating
			}
		}
	}

	// Check if we have enough players
	if len(matched) < m.config.MinPlayers {
		// Check wait time - if exceeded, start with fewer players
		firstPlayer := players[0]
		if time.Since(firstPlayer.WaitStart) > m.config.MaxWaitTime && len(matched) >= 2 {
			return matched
		}
		return nil
	}

	// Remove matched players from queue
	for _, userID := range matched {
		queue.RemovePlayer(ctx, userID)
	}

	return matched
}

// GetQueueStats returns matchmaking statistics
func (m *Matchmaker) GetQueueStats() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()

	stats := make(map[string]interface{})

	for key, queue := range m.queues {
		queue.mu.RLock()
		count := len(queue.players)
		queue.mu.RUnlock()
		stats[key] = map[string]int{"waiting": count}
	}

	return stats
}

// abs returns absolute value of integer
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
