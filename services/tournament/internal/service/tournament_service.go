package service

import (
	"context"
	"time"

	"github.com/gameengine/tournament/internal/tournament"
)

// TournamentService provides tournament business logic
type TournamentService struct {
	manager *tournament.Manager
}

// NewTournamentService creates a new tournament service
func NewTournamentService(manager *tournament.Manager) (*TournamentService, error) {
	return &TournamentService{
		manager: manager,
	}, nil
}

// CreateTournament creates a new tournament
func (s *TournamentService) CreateTournament(ctx context.Context, name string, tType tournament.TournamentType, gameType string, entryFee int64, buyIn int64, minPlayers, maxPlayers int, startTime time.Time) (*tournament.Tournament, error) {
	settings := tournament.TournamentSettings{
		AutoStart:        true,
		RebuyEnabled:     false,
		AddonEnabled:     false,
		Reentries:        1,
		LateRegistration: 300,
		StartingChips:    1000,
	}

	return s.manager.CreateTournament(ctx, name, tType, gameType, entryFee, buyIn, minPlayers, maxPlayers, startTime, settings)
}

// GetTournament retrieves a tournament by ID
func (s *TournamentService) GetTournament(ctx context.Context, id string) (*tournament.Tournament, error) {
	return s.manager.GetTournament(ctx, id)
}

// RegisterUser registers a user for a tournament
func (s *TournamentService) RegisterUser(ctx context.Context, tournamentID string, userID, username string) error {
	return s.manager.RegisterUser(ctx, tournamentID, userID, username)
}

// UnregisterUser unregisters a user from a tournament
func (s *TournamentService) UnregisterUser(ctx context.Context, tournamentID string, userID string) error {
	return s.manager.UnregisterUser(ctx, tournamentID, userID)
}

// StartTournament starts a tournament
func (s *TournamentService) StartTournament(ctx context.Context, tournamentID string) error {
	return s.manager.StartTournament(ctx, tournamentID)
}

// EndTournament ends a tournament
func (s *TournamentService) EndTournament(ctx context.Context, tournamentID string) error {
	return s.manager.EndTournament(ctx, tournamentID)
}

// UpdatePlayerScore updates a player's score
func (s *TournamentService) UpdatePlayerScore(ctx context.Context, tournamentID string, userID string, scoreDelta int) error {
	return s.manager.UpdatePlayerScore(ctx, tournamentID, userID, scoreDelta)
}

// EliminatePlayer eliminates a player
func (s *TournamentService) EliminatePlayer(ctx context.Context, tournamentID string, userID string, rank int) error {
	return s.manager.EliminatePlayer(ctx, tournamentID, userID, rank)
}

// ListTournaments lists tournaments with filters
func (s *TournamentService) ListTournaments(ctx context.Context, status tournament.TournamentStatus, tType tournament.TournamentType) ([]*tournament.Tournament, error) {
	return s.manager.ListTournaments(ctx, status, tType)
}

// GetLeaderboard gets tournament leaderboard
func (s *TournamentService) GetLeaderboard(ctx context.Context, tournamentID string, limit int) ([]tournament.LeaderboardEntry, error) {
	return s.manager.leaderboard.GetLeaderboard(ctx, tournamentID, limit)
}

// GetPlayerRank gets a player's rank
func (s *TournamentService) GetPlayerRank(ctx context.Context, tournamentID string, userID string) (*tournament.LeaderboardEntry, error) {
	return s.manager.leaderboard.GetPlayerRank(ctx, tournamentID, userID)
}

// GetTopPlayers gets top N players
func (s *TournamentService) GetTopPlayers(ctx context.Context, tournamentID string, n int) ([]tournament.LeaderboardEntry, error) {
	return s.manager.leaderboard.GetTopPlayers(ctx, tournamentID, n)
}

// GetScheduledTournaments gets scheduled tournaments
func (s *TournamentService) GetScheduledTournaments(ctx context.Context) ([]tournament.ScheduledTournament, error) {
	return s.manager.scheduler.GetScheduledTournaments(ctx)
}

// GetPrizeDistribution gets prize distribution
func (s *TournamentService) GetPrizeDistribution(ctx context.Context, tournamentID string) ([]tournament.Result, error) {
	return s.manager.prizePool.GetPrizeDistribution(ctx, tournamentID)
}
