package room

import (
	"context"
	"sync"
	"time"
)

// CollusionDetector detects potential collusion between players
type CollusionDetector struct {
	mu             sync.RWMutex
	playerStats    map[string]*PlayerStats   // userID -> stats
	gameHistory    map[string][]GameSnapshot // tableID -> history
	windowSize     time.Duration
	alertThreshold float64
}

// PlayerStats tracks player statistics for collusion detection
type PlayerStats struct {
	UserID             string
	GameHistory        []GameSnapshot
	IPAddresses        map[string]bool
	DeviceFingerprints map[string]bool
	AccountAges        map[string]time.Duration
	WinRates           map[string]float64
	SuspiciousScore    float64
	LastUpdate         time.Time
}

// GameSnapshot represents a snapshot of game state
type GameSnapshot struct {
	Timestamp time.Time
	TableID   string
	GameType  string
	Players   []string
	Actions   map[string]string
	Results   map[string]int
}

// CollusionAlert represents a potential collusion alert
type CollusionAlert struct {
	Player1   string
	Player2   string
	Score     float64
	Reasons   []string
	TableID   string
	Timestamp time.Time
}

// NewCollusionDetector creates a new collusion detector
func NewCollusionDetector(windowSize time.Duration, alertThreshold float64) *CollusionDetector {
	if windowSize == 0 {
		windowSize = 24 * time.Hour
	}
	if alertThreshold == 0 {
		alertThreshold = 0.7
	}

	return &CollusionDetector{
		playerStats:    make(map[string]*PlayerStats),
		gameHistory:    make(map[string][]GameSnapshot),
		windowSize:     windowSize,
		alertThreshold: alertThreshold,
	}
}

// RecordGame records a game result for analysis
func (cd *CollusionDetector) RecordGame(ctx context.Context, tableID, gameType string, players []string, actions map[string]string, results map[string]int) {
	cd.mu.Lock()
	defer cd.mu.Unlock()

	snapshot := GameSnapshot{
		Timestamp: time.Now(),
		TableID:   tableID,
		GameType:  gameType,
		Players:   players,
		Actions:   actions,
		Results:   results,
	}

	for _, playerID := range players {
		stats, ok := cd.playerStats[playerID]
		if !ok {
			stats = &PlayerStats{
				UserID:             playerID,
				IPAddresses:        make(map[string]bool),
				DeviceFingerprints: make(map[string]bool),
				AccountAges:        make(map[string]time.Duration),
				WinRates:           make(map[string]float64),
			}
			cd.playerStats[playerID] = stats
		}

		stats.GameHistory = append(stats.GameHistory, snapshot)
		stats.LastUpdate = time.Now()

		if result, won := results[playerID]; won && result > 0 {
			cd.updateWinRate(stats, gameType, true)
		} else if _, lost := results[playerID]; lost {
			cd.updateWinRate(stats, gameType, false)
		}
	}

	cd.gameHistory[tableID] = append(cd.gameHistory[tableID], snapshot)
	cd.cleanOldData()
}

// RecordIP records an IP address for a player
func (cd *CollusionDetector) RecordIP(ctx context.Context, userID, ipAddress string) {
	cd.mu.Lock()
	defer cd.mu.Unlock()

	stats, ok := cd.playerStats[userID]
	if !ok {
		stats = &PlayerStats{
			UserID:             userID,
			IPAddresses:        make(map[string]bool),
			DeviceFingerprints: make(map[string]bool),
			AccountAges:        make(map[string]time.Duration),
			WinRates:           make(map[string]float64),
		}
		cd.playerStats[userID] = stats
	}

	stats.IPAddresses[ipAddress] = true
	stats.LastUpdate = time.Now()
}

// RecordDevice records a device fingerprint for a player
func (cd *CollusionDetector) RecordDevice(ctx context.Context, userID, fingerprint string) {
	cd.mu.Lock()
	defer cd.mu.Unlock()

	stats, ok := cd.playerStats[userID]
	if !ok {
		stats = &PlayerStats{
			UserID:             userID,
			IPAddresses:        make(map[string]bool),
			DeviceFingerprints: make(map[string]bool),
			AccountAges:        make(map[string]time.Duration),
			WinRates:           make(map[string]float64),
		}
		cd.playerStats[userID] = stats
	}

	stats.DeviceFingerprints[fingerprint] = true
	stats.LastUpdate = time.Now()
}

// GetSuspiciousPlayers returns players with high suspicious scores
func (cd *CollusionDetector) GetSuspiciousPlayers(ctx context.Context, minScore float64) []string {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	var result []string
	for userID, stats := range cd.playerStats {
		if stats.SuspiciousScore >= minScore {
			result = append(result, userID)
		}
	}

	return result
}

// UpdateSuspiciousScore updates the suspicious score for a player
func (cd *CollusionDetector) UpdateSuspiciousScore(ctx context.Context, userID string, score float64) {
	cd.mu.Lock()
	defer cd.mu.Unlock()

	if stats, ok := cd.playerStats[userID]; ok {
		stats.SuspiciousScore = (stats.SuspiciousScore + score) / 2
		stats.LastUpdate = time.Now()
	}
}

// GetPlayerStats returns statistics for a player
func (cd *CollusionDetector) GetPlayerStats(ctx context.Context, userID string) *PlayerStats {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	return cd.playerStats[userID]
}

// cleanOldData removes data outside the window
func (cd *CollusionDetector) cleanOldData() {
	cutoff := time.Now().Add(-cd.windowSize)

	for userID, stats := range cd.playerStats {
		var recentHistory []GameSnapshot
		for _, snap := range stats.GameHistory {
			if snap.Timestamp.After(cutoff) {
				recentHistory = append(recentHistory, snap)
			}
		}
		if len(recentHistory) == 0 {
			delete(cd.playerStats, userID)
		} else {
			stats.GameHistory = recentHistory
		}
	}

	for tableID, history := range cd.gameHistory {
		var recentHistory []GameSnapshot
		for _, snap := range history {
			if snap.Timestamp.After(cutoff) {
				recentHistory = append(recentHistory, snap)
			}
		}
		if len(recentHistory) == 0 {
			delete(cd.gameHistory, tableID)
		} else {
			cd.gameHistory[tableID] = recentHistory
		}
	}
}
