package tournament

import (
	"context"
	"sync"
	"time"
)

// BlindLevelManager manages tournament blind levels
type BlindLevelManager struct {
	mu           sync.RWMutex
	timers       map[string]*BlindTimer // tournamentID -> timer
	currentLevel map[string]int         // tournamentID -> current level index
}

// BlindTimer represents a timer for blind level progression
type BlindTimer struct {
	TournamentID string
	LevelIndex   int
	Level        BlindLevel
	Remaining    time.Duration
	OnLevelUp    func(tournamentID string, newLevel BlindLevel)
	onLevelUpCh  chan BlindLevel
}

// NewBlindLevelManager creates a new blind level manager
func NewBlindLevelManager() *BlindLevelManager {
	return &BlindLevelManager{
		timers:       make(map[string]*BlindTimer),
		currentLevel: make(map[string]int),
	}
}

// StartTournament starts the blind level timer for a tournament
func (bm *BlindLevelManager) StartTournament(ctx context.Context, tournamentID string, levels []BlindLevel, onLevelUp func(string, BlindLevel)) error {
	if len(levels) == 0 {
		return nil // No levels to run
	}

	bm.mu.Lock()
	defer bm.mu.Unlock()

	// Check if already running
	if _, exists := bm.timers[tournamentID]; exists {
		return nil // Already running
	}

	// Create channel for level up events
	onLevelUpCh := make(chan BlindLevel, 10)

	timer := &BlindTimer{
		TournamentID: tournamentID,
		LevelIndex:   0,
		Level:        levels[0],
		Remaining:    time.Duration(levels[0].Duration) * time.Second,
		OnLevelUp:    onLevelUp,
		onLevelUpCh:  onLevelUpCh,
	}

	bm.timers[tournamentID] = timer
	bm.currentLevel[tournamentID] = 0

	// Start the timer goroutine
	go bm.runTimer(ctx, tournamentID, levels, onLevelUpCh)

	return nil
}

// runTimer runs the blind level timer
func (bm *BlindLevelManager) runTimer(ctx context.Context, tournamentID string, levels []BlindLevel, onLevelUpCh chan BlindLevel) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			bm.mu.Lock()
			timer, ok := bm.timers[tournamentID]
			if !ok {
				bm.mu.Unlock()
				return
			}

			timer.Remaining -= time.Second

			if timer.Remaining <= 0 {
				// Time for next level
				timer.LevelIndex++

				if timer.LevelIndex >= len(levels) {
					// Tournament should be ending - notify but don't go to next level
					delete(bm.timers, tournamentID)
					bm.mu.Unlock()
					return
				}

				// Set up next level
				timer.Level = levels[timer.LevelIndex]
				timer.Remaining = time.Duration(levels[timer.LevelIndex].Duration) * time.Second
				bm.currentLevel[tournamentID] = timer.LevelIndex

				// Notify
				select {
				case onLevelUpCh <- timer.Level:
				default:
				}

				if timer.OnLevelUp != nil {
					go timer.OnLevelUp(tournamentID, timer.Level)
				}
			}
			bm.mu.Unlock()
		}
	}
}

// GetCurrentLevel returns the current blind level for a tournament
func (bm *BlindLevelManager) GetCurrentLevel(tournamentID string) (int, BlindLevel) {
	bm.mu.RLock()
	defer bm.mu.RUnlock()

	levelIdx, ok := bm.currentLevel[tournamentID]
	if !ok {
		return 0, BlindLevel{}
	}

	timer, ok := bm.timers[tournamentID]
	if !ok {
		return levelIdx, BlindLevel{}
	}

	return levelIdx, timer.Level
}

// GetRemainingTime returns remaining time for current level
func (bm *BlindLevelManager) GetRemainingTime(tournamentID string) time.Duration {
	bm.mu.RLock()
	defer bm.mu.Unlock()

	if timer, ok := bm.timers[tournamentID]; ok {
		return timer.Remaining
	}

	return 0
}

// Pause pauses the blind timer
func (bm *BlindLevelManager) Pause(tournamentID string) {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	if _, ok := bm.timers[tournamentID]; ok {
		// In a real implementation, we'd stop the timer
		// For now, just remove it to pause
		delete(bm.timers, tournamentID)
	}
}

// Resume resumes the blind timer
func (bm *BlindLevelManager) Resume(ctx context.Context, tournamentID string, levels []BlindLevel, onLevelUp func(string, BlindLevel)) error {
	bm.mu.Lock()
	currentIdx := bm.currentLevel[tournamentID]
	bm.mu.Unlock()

	if currentIdx == 0 && len(levels) > 0 {
		return bm.StartTournament(ctx, tournamentID, levels[currentIdx:], onLevelUp)
	}

	return nil
}

// Stop stops the blind timer for a tournament
func (bm *BlindLevelManager) Stop(tournamentID string) {
	bm.mu.Lock()
	defer bm.mu.Unlock()

	delete(bm.timers, tournamentID)
	delete(bm.currentLevel, tournamentID)
}

// GenerateDefaultBlindLevels generates standard tournament blind levels
func GenerateDefaultBlindLevels(startingBigBlind int, levelDuration int, maxLevels int) []BlindLevel {
	levels := make([]BlindLevel, 0, maxLevels)
	smallBlind := startingBigBlind / 2
	bigBlind := startingBigBlind

	for level := 1; level <= maxLevels; level++ {
		levels = append(levels, BlindLevel{
			Level:      level,
			SmallBlind: smallBlind,
			BigBlind:   bigBlind,
			Duration:   levelDuration,
		})

		// Increase blinds every level (or every 2 levels)
		smallBlind = smallBlind * 3 / 2
		bigBlind = bigBlind * 3 / 2

		// Cap at reasonable limits
		if smallBlind > 10000 {
			smallBlind = 10000
		}
		if bigBlind > 20000 {
			bigBlind = 20000
		}
	}

	return levels
}

// GenerateTurboBlindLevels generates turbo tournament blind levels
func GenerateTurboBlindLevels(startingBigBlind int) []BlindLevel {
	return []BlindLevel{
		{Level: 1, SmallBlind: 25, BigBlind: 50, Duration: 300},
		{Level: 2, SmallBlind: 50, BigBlind: 100, Duration: 300},
		{Level: 3, SmallBlind: 75, BigBlind: 150, Duration: 300},
		{Level: 4, SmallBlind: 100, BigBlind: 200, Duration: 300},
		{Level: 5, SmallBlind: 150, BigBlind: 300, Duration: 300},
		{Level: 6, SmallBlind: 200, BigBlind: 400, Duration: 300},
		{Level: 7, SmallBlind: 300, BigBlind: 600, Duration: 300},
		{Level: 8, SmallBlind: 400, BigBlind: 800, Duration: 300},
		{Level: 9, SmallBlind: 500, BigBlind: 1000, Duration: 300},
		{Level: 10, SmallBlind: 750, BigBlind: 1500, Duration: 300},
	}
}

// GenerateHyperTurboBlindLevels generates hyper turbo blind levels
func GenerateHyperTurboBlindLevels() []BlindLevel {
	return []BlindLevel{
		{Level: 1, SmallBlind: 50, BigBlind: 100, Duration: 180},
		{Level: 2, SmallBlind: 100, BigBlind: 200, Duration: 180},
		{Level: 3, SmallBlind: 200, BigBlind: 400, Duration: 180},
		{Level: 4, SmallBlind: 300, BigBlind: 600, Duration: 180},
		{Level: 5, SmallBlind: 500, BigBlind: 1000, Duration: 180},
		{Level: 6, SmallBlind: 1000, BigBlind: 2000, Duration: 180},
	}
}

// GetBlindSchedule returns a formatted blind schedule for display
func GetBlindSchedule(levels []BlindLevel) []map[string]interface{} {
	schedule := make([]map[string]interface{}, len(levels))

	for i, level := range levels {
		schedule[i] = map[string]interface{}{
			"level":       level.Level,
			"small_blind": level.SmallBlind,
			"big_blind":   level.BigBlind,
			"duration":    level.Duration,
			"ante":        level.BigBlind / 5, // Optional ante
		}
	}

	return schedule
}
