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
	rankings  []string
}

// MatchmakingPlayer represents a player in matchmaking
type MatchmakingPlayer struct {
	UserID         string
	Username       string
	SkillRating    int
	StakeLevel     int
	GameType       string
	WaitStart      time.Time
	PreferredTable string
}

// MatchmakingConfig represents matchmaking configuration
type MatchmakingConfig struct {
	MaxSkillDiff     int
	MaxWaitTime      time.Duration
	MinPlayers       int
	MaxPlayers       int
	AutoStartEnabled bool
}

// Matchmaker handles matchmaking between players
type Matchmaker struct {
	queues map[string]*MatchmakingQueue
	config *MatchmakingConfig
	mu     sync.RWMutex
}

// NewMatchmakingQueue creates a new matchmaking queue
func NewMatchmakingQueue() *MatchmakingQueue {
	return &MatchmakingQueue{
		players:   make(map[string]*MatchmakingPlayer),
		waitTimes: make(map[string]time.Time),
		rankings:  make([]string, 0),
	}
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

// AddPlayer adds a player to the matchmaking queue
func (mq *MatchmakingQueue) AddPlayer(ctx context.Context, player *MatchmakingPlayer) error {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	if _, exists := mq.players[player.UserID]; exists {
		return nil
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

	for _, p := range mq.players {
		if p.UserID == player.UserID {
			continue
		}

		if p.GameType != player.GameType || p.StakeLevel != player.StakeLevel {
			continue
		}

		if abs(p.SkillRating-player.SkillRating) > 200 {
			continue
		}

		candidates = append(candidates, p)
	}

	sort.Slice(candidates, func(i, j int) bool {
		return mq.waitTimes[candidates[i].UserID].Before(mq.waitTimes[candidates[j].UserID])
	})

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
