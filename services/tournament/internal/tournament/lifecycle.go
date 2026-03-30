package tournament

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

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
