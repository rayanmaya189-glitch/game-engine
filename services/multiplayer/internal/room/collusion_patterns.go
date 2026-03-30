package room

import (
	"context"
	"math"
	"time"
)

// updateWinRate updates the win rate for a game type
func (cd *CollusionDetector) updateWinRate(stats *PlayerStats, gameType string, won bool) {
	current := stats.WinRates[gameType]
	games := float64(len(stats.GameHistory))

	if games == 0 {
		stats.WinRates[gameType] = 0
		return
	}

	newWinRate := current + (1-current)/games
	if !won {
		newWinRate = current - current/games
	}

	stats.WinRates[gameType] = newWinRate
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

	if len(stats1.DeviceFingerprints) > 0 && len(stats2.DeviceFingerprints) > 0 {
		for fp := range stats1.DeviceFingerprints {
			if stats2.DeviceFingerprints[fp] {
				score += 0.3
				reasons = append(reasons, "same device fingerprint")
				break
			}
		}
	}

	score += cd.calculateTableCooccurrence(stats1, stats2)
	score += cd.detectUnusualPatterns(stats1, stats2)

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

	tables1 := make(map[string]int)
	tables2 := make(map[string]int)

	for _, game := range stats1.GameHistory {
		tables1[game.TableID]++
	}
	for _, game := range stats2.GameHistory {
		tables2[game.TableID]++
	}

	common := 0
	for table := range tables1 {
		if tables2[table] > 0 {
			common++
		}
	}

	games1 := float64(len(stats1.GameHistory))
	games2 := float64(len(stats2.GameHistory))
	minGames := math.Min(games1, games2)

	if minGames == 0 {
		return 0
	}

	cooccurrence := float64(common) / minGames

	if cooccurrence > 0.5 {
		return 0.4
	} else if cooccurrence > 0.2 {
		return 0.2
	}

	return 0
}

// detectUnusualPatterns detects unusual betting patterns between two players
func (cd *CollusionDetector) detectUnusualPatterns(stats1, stats2 *PlayerStats) float64 {
	score := 0.0

	for gameType, winRate1 := range stats1.WinRates {
		winRate2, exists := stats2.WinRates[gameType]
		if !exists {
			continue
		}

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

	playerSet := make(map[string]bool)
	for _, snapshot := range history[len(history)-10:] {
		for _, player := range snapshot.Players {
			playerSet[player] = true
		}
	}

	players := make([]string, 0, len(playerSet))
	for p := range playerSet {
		players = append(players, p)
	}

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
