package tournament

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// ScheduledTournament represents a scheduled tournament
type ScheduledTournament struct {
	TournamentID string    `json:"tournament_id"`
	ScheduledAt  time.Time `json:"scheduled_at"`
	Started      bool      `json:"started"`
	Cancelled    bool      `json:"cancelled"`
}

// Scheduler handles tournament scheduling
type Scheduler struct {
	mu            sync.RWMutex
	redisClient   *redis.Client
	scheduledJobs map[string]*time.Timer
}

// NewScheduler creates a new tournament scheduler
func NewScheduler(redisClient *redis.Client) *Scheduler {
	return &Scheduler{
		redisClient:   redisClient,
		scheduledJobs: make(map[string]*time.Timer),
	}
}

// ScheduleTournament schedules a tournament to start at a specific time
func (s *Scheduler) ScheduleTournament(ctx context.Context, tournamentID string, startTime time.Time) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delay := startTime.Sub(time.Now())
	if delay <= 0 {
		return fmt.Errorf("start time must be in the future")
	}

	scheduled := ScheduledTournament{
		TournamentID: tournamentID,
		ScheduledAt:  startTime,
		Started:      false,
		Cancelled:    false,
	}

	// Save to Redis
	key := fmt.Sprintf("scheduled:%s", tournamentID)
	data, err := json.Marshal(scheduled)
	if err != nil {
		return err
	}

	if err := s.redisClient.Set(ctx, key, data, delay+time.Hour).Err(); err != nil {
		return err
	}

	// Schedule the job
	timer := time.AfterFunc(delay, func() {
		s.startTournament(ctx, tournamentID)
	})

	s.scheduledJobs[tournamentID] = timer

	return nil
}

// CancelScheduledTournament cancels a scheduled tournament
func (s *Scheduler) CancelScheduledTournament(ctx context.Context, tournamentID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Cancel timer if exists
	if timer, ok := s.scheduledJobs[tournamentID]; ok {
		timer.Stop()
		delete(s.scheduledJobs, tournamentID)
	}

	// Update in Redis
	key := fmt.Sprintf("scheduled:%s", tournamentID)
	data, err := s.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	var scheduled ScheduledTournament
	if err := json.Unmarshal(data, &scheduled); err != nil {
		return err
	}

	scheduled.Cancelled = true
	data, _ = json.Marshal(scheduled)

	return s.redisClient.Set(ctx, key, data, 0).Err()
}

// GetScheduledTournaments gets all scheduled tournaments
func (s *Scheduler) GetScheduledTournaments(ctx context.Context) ([]ScheduledTournament, error) {
	keys, err := s.redisClient.Keys(ctx, "scheduled:*").Result()
	if err != nil {
		return nil, err
	}

	scheduled := make([]ScheduledTournament, 0, len(keys))
	for _, key := range keys {
		data, err := s.redisClient.Get(ctx, key).Bytes()
		if err != nil {
			continue
		}

		var st ScheduledTournament
		if err := json.Unmarshal(data, &st); err != nil {
			continue
		}

		if !st.Cancelled && !st.Started {
			scheduled = append(scheduled, st)
		}
	}

	return scheduled, nil
}

// GetNextTournament gets the next scheduled tournament
func (s *Scheduler) GetNextTournament(ctx context.Context) (*ScheduledTournament, error) {
	scheduled, err := s.GetScheduledTournaments(ctx)
	if err != nil {
		return nil, err
	}

	if len(scheduled) == 0 {
		return nil, nil
	}

	// Find the earliest
	var next *ScheduledTournament
	for i := range scheduled {
		if next == nil || scheduled[i].ScheduledAt.Before(next.ScheduledAt) {
			next = &scheduled[i]
		}
	}

	return next, nil
}

// GetTournamentsStartingSoon gets tournaments starting within the specified duration
func (s *Scheduler) GetTournamentsStartingSoon(ctx context.Context, within time.Duration) ([]ScheduledTournament, error) {
	scheduled, err := s.GetScheduledTournaments(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	until := now.Add(within)

	result := make([]ScheduledTournament, 0)
	for _, st := range scheduled {
		if st.ScheduledAt.After(now) && st.ScheduledAt.Before(until) {
			result = append(result, st)
		}
	}

	return result, nil
}

// RescheduleTournament reschedules a tournament to a new time
func (s *Scheduler) RescheduleTournament(ctx context.Context, tournamentID string, newStartTime time.Time) error {
	// Cancel existing schedule
	if err := s.CancelScheduledTournament(ctx, tournamentID); err != nil {
		// Ignore error if not scheduled
	}

	// Create new schedule
	return s.ScheduleTournament(ctx, tournamentID, newStartTime)
}

// startTournament starts a scheduled tournament
func (s *Scheduler) startTournament(ctx context.Context, tournamentID string) {
	s.mu.Lock()
	delete(s.scheduledJobs, tournamentID)
	s.mu.Unlock()

	// This would typically call the tournament manager to start
	// For now, we just mark it as started in Redis
	key := fmt.Sprintf("scheduled:%s", tournamentID)
	data, err := s.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return
	}

	var scheduled ScheduledTournament
	if err := json.Unmarshal(data, &scheduled); err != nil {
		return
	}

	scheduled.Started = true
	data, _ = json.Marshal(scheduled)

	_ = s.redisClient.Set(ctx, key, data, time.Hour).Err()
}

// StartScheduler starts the background scheduler
func (s *Scheduler) StartScheduler(ctx context.Context, checkInterval time.Duration) {
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.checkAndStartTournaments(ctx)
		}
	}
}

// checkAndStartTournaments checks for tournaments that need to start
func (s *Scheduler) checkAndStartTournaments(ctx context.Context) {
	scheduled, err := s.GetTournamentsStartingSoon(ctx, time.Minute)
	if err != nil {
		return
	}

	for _, st := range scheduled {
		if time.Now().After(st.ScheduledAt) || time.Now().Equal(st.ScheduledAt) {
			s.startTournament(ctx, st.TournamentID)
		}
	}
}

// GetScheduledTime gets the scheduled start time for a tournament
func (s *Scheduler) GetScheduledTime(ctx context.Context, tournamentID string) (*time.Time, error) {
	key := fmt.Sprintf("scheduled:%s", tournamentID)

	data, err := s.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var scheduled ScheduledTournament
	if err := json.Unmarshal(data, &scheduled); err != nil {
		return nil, err
	}

	return &scheduled.ScheduledAt, nil
}

// IsScheduled checks if a tournament is scheduled
func (s *Scheduler) IsScheduled(ctx context.Context, tournamentID string) (bool, error) {
	key := fmt.Sprintf("scheduled:%s", tournamentID)

	exists, err := s.redisClient.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return exists > 0, nil
}
