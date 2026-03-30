package tournament

import (
	"context"
	"fmt"
	"time"
)

// RegisterUser registers a user for a tournament
func (m *Manager) RegisterUser(ctx context.Context, tournamentID string, userID, username string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	tournament, ok := m.tournaments[tournamentID]
	if !ok {
		return fmt.Errorf("tournament not found: %s", tournamentID)
	}

	if tournament.Status != TournamentStatusPending && tournament.Status != TournamentStatusRegistering {
		return fmt.Errorf("tournament is not accepting registrations")
	}

	if tournament.CurrentPlayers >= tournament.MaxPlayers {
		return fmt.Errorf("tournament is full")
	}

	if _, exists := tournament.RegisteredUsers[userID]; exists {
		return fmt.Errorf("user already registered")
	}

	participant := Participant{
		UserID:       userID,
		Username:     username,
		Chips:        tournament.Settings.StartingChips,
		Status:       "registered",
		RegisteredAt: time.Now(),
	}

	tournament.RegisteredUsers[userID] = participant
	tournament.CurrentPlayers++
	tournament.PrizePool = int64(tournament.CurrentPlayers) * tournament.EntryFee
	tournament.UpdatedAt = time.Now()

	// Update status to registering if we have at least min players
	if tournament.CurrentPlayers >= tournament.MinPlayers && tournament.Status == TournamentStatusPending {
		tournament.Status = TournamentStatusRegistering
	}

	// Save to Redis
	if err := m.saveTournament(ctx, tournament); err != nil {
		return err
	}

	// Update leaderboard
	m.Leaderboard.UpdateRegistration(ctx, tournamentID, participant)

	// Auto-start if enabled
	if tournament.Settings.AutoStart && tournament.Status == TournamentStatusRegistering {
		if tournament.Type == TournamentTypeSitAndGo {
			go m.startTournament(ctx, tournamentID)
		}
	}

	return nil
}

// UnregisterUser unregisters a user from a tournament
func (m *Manager) UnregisterUser(ctx context.Context, tournamentID string, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	tournament, ok := m.tournaments[tournamentID]
	if !ok {
		return fmt.Errorf("tournament not found: %s", tournamentID)
	}

	if tournament.Status == TournamentStatusRunning {
		return fmt.Errorf("cannot unregister while tournament is running")
	}

	participant, exists := tournament.RegisteredUsers[userID]
	if !exists {
		return fmt.Errorf("user not registered")
	}

	delete(tournament.RegisteredUsers, userID)
	tournament.CurrentPlayers--
	tournament.PrizePool = int64(tournament.CurrentPlayers) * tournament.EntryFee
	tournament.UpdatedAt = time.Now()

	// Revert status if below minimum
	if tournament.CurrentPlayers < tournament.MinPlayers && tournament.Status == TournamentStatusRegistering {
		tournament.Status = TournamentStatusPending
	}

	// Save to Redis
	if err := m.saveTournament(ctx, tournament); err != nil {
		return err
	}

	// Remove from leaderboard
	m.Leaderboard.RemoveParticipant(ctx, tournamentID, userID)
	_ = participant // Use participant for refund logic

	return nil
}
