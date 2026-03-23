package room

import (
	"context"
	"math"
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
	AccountAges        map[string]time.Duration // account creation time
	WinRates           map[string]float64       // by game type
	SuspiciousScore    float64
	LastUpdate         time.Time
}

// GameSnapshot represents a snapshot of game state
type GameSnapshot struct {
	Timestamp time.Time
	TableID   string
	GameType  string
	Players   []string
	Actions   map[string]string // playerID -> action
	Results   map[string]int    // playerID -> result (chips won/lost)
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

	// Record for each player
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

		// Update win rate
		if result, won := results[playerID]; won && result > 0 {
			cd.updateWinRate(stats, gameType, true)
		} else if _, lost := results[playerID]; lost {
			cd.updateWinRate(stats, gameType, false)
		}
	}

	// Record for table
	cd.gameHistory[tableID] = append(cd.gameHistory[tableID], snapshot)

	// Clean old data
	cd.cleanOldData()
}

// updateWinRate updates the win rate for a game type
func (cd *CollusionDetector) updateWinRate(stats *PlayerStats, gameType string, won bool) {
	current := stats.WinRates[gameType]
	games := float64(len(stats.GameHistory))

	if games == 0 {
		stats.WinRates[gameType] = 0
		return
	}

	// Simple moving average
	newWinRate := current + (1-current)/games
	if !won {
		newWinRate = current - current/games
	}

	stats.WinRates[gameType] = newWinRate
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

// DetectCollusion detects potential collusion between two players
func (cd *CollusionDetector) DetectCollusion(ctx context.Context, player1ID, player2ID string) (float64, []string) {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	stats1, ok1 := cd.playerStats[player1ID]
	stats2, ok2 := cd.playerStats[player2ID]

	if !ok1 || !ok2 {
		return 0, nil
	}

	var reasons []string
	score := 0.0

	// Check 1: Same IP address
	// Not checking - just for reference, would require IP data

	// Check 2: Same device fingerprint
	if len(stats1.DeviceFingerprints) > 0 && len(stats2.DeviceFingerprints) > 0 {
		for fp := range stats1.DeviceFingerprints {
			if stats2.DeviceFingerprints[fp] {
				score += 0.3
				reasons = append(reasons, "same device fingerprint")
				break
			}
		}
	}

	// Check 3: Suspected IP overlap (would need real IP tracking)
	// Simplified: check if they're playing at the same tables frequently
	score += cd.calculateTableCooccurrence(stats1, stats2)

	// Check 4: Unusual betting patterns
	score += cd.detectUnusualPatterns(stats1, stats2)

	// Cap score at 1.0
	if score > 1.0 {
		score = 1.0
	}

	return score, reasons
}

// calculateTableCooccurrence calculates how often two players play together
func (cd *CollusionDetector) calculateTableCooccurrence(stats1, stats2 *PlayerStats) float64 {
	if len(stats1.GameHistory) == 0 || len(stats2.GameHistory) == 0 {
		return 0
	}

	// Build set of tables for each player
	tables1 := make(map[string]int)
	tables2 := make(map[string]int)

	for _, game := range stats1.GameHistory {
		tables1[game.TableID]++
	}
	for _, game := range stats2.GameHistory {
		tables2[game.TableID]++
	}

	// Count common tables
	common := 0
	for table := range tables1 {
		if tables2[table] > 0 {
			common++
		}
	}

	// Calculate cooccurrence ratio
	games1 := float64(len(stats1.GameHistory))
	games2 := float64(len(stats2.GameHistory))
	minGames := math.Min(games1, games2)

	if minGames == 0 {
		return 0
	}

	cooccurrence := float64(common) / minGames

	// High cooccurrence is suspicious
	if cooccurrence > 0.5 {
		return 0.4
	} else if cooccurrence > 0.2 {
		return 0.2
	}

	return 0
}

// detectUnusualPatterns detects unusual betting patterns between two players
func (cd *CollusionDetector) detectUnusualPatterns(stats1, stats2 *PlayerStats) float64 {
	// Check if both players have very similar win rates (suspicious)
	score := 0.0

	for gameType, winRate1 := range stats1.WinRates {
		winRate2, exists := stats2.WinRates[gameType]
		if !exists {
			continue
		}

		// If both have very high or very similar win rates
		diff := math.Abs(winRate1 - winRate2)
		if diff < 0.1 && (winRate1 > 0.6 || winRate1 < 0.3) {
			score += 0.2
		}
	}

	return score
}

// AnalyzeTable analyzes a table for collusion patterns
func (cd *CollusionDetector) AnalyzeTable(ctx context.Context, tableID string) []CollusionAlert {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	history, ok := cd.gameHistory[tableID]
	if !ok || len(history) == 0 {
		return nil
	}

	// Get unique players from recent games
	playerSet := make(map[string]bool)
	for _, snapshot := range history[len(history)-10:] { // Last 10 games
		for _, player := range snapshot.Players {
			playerSet[player] = true
		}
	}

	players := make([]string, 0, len(playerSet))
	for p := range playerSet {
		players = append(players, p)
	}

	// Check all pairs
	var alerts []CollusionAlert
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			score, reasons := cd.DetectCollusion(ctx, players[i], players[j])
			if score >= cd.alertThreshold {
				alerts = append(alerts, CollusionAlert{
					Player1:   players[i],
					Player2:   players[j],
					Score:     score,
					Reasons:   reasons,
					TableID:   tableID,
					Timestamp: time.Now(),
				})
			}
		}
	}

	return alerts
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
		// Average with existing score
		stats.SuspiciousScore = (stats.SuspiciousScore + score) / 2
		stats.LastUpdate = time.Now()
	}
}

// cleanOldData removes data outside the window
func (cd *CollusionDetector) cleanOldData() {
	cutoff := time.Now().Add(-cd.windowSize)

	// Clean player stats
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

	// Clean table history
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

// GetPlayerStats returns statistics for a player
func (cd *CollusionDetector) GetPlayerStats(ctx context.Context, userID string) *PlayerStats {
	cd.mu.RLock()
	defer cd.mu.RUnlock()

	return cd.playerStats[userID]
}
