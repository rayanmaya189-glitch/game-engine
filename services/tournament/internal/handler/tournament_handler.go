package handler

import (
	"context"
	"time"

	"github.com/game_engine/tournament/internal/service"
	"github.com/game_engine/tournament/internal/tournament"
)

// TournamentHandler handles tournament-related requests
type TournamentHandler struct {
	service *service.TournamentService
}

// NewTournamentHandler creates a new tournament handler
func NewTournamentHandler(service *service.TournamentService) *TournamentHandler {
	return &TournamentHandler{
		service: service,
	}
}

// CreateTournamentRequest represents a create tournament request
type CreateTournamentRequest struct {
	Name       string                    `json:"name"`
	Type       tournament.TournamentType `json:"type"`
	GameType   string                    `json:"game_type"`
	EntryFee   int64                     `json:"entry_fee"`
	BuyIn      int64                     `json:"buy_in"`
	MinPlayers int                       `json:"min_players"`
	MaxPlayers int                       `json:"max_players"`
	StartTime  time.Time                 `json:"start_time"`
}

// RegisterUserRequest represents a register user request
type RegisterUserRequest struct {
	TournamentID string `json:"tournament_id"`
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
}

// TournamentResponse represents a tournament response
type TournamentResponse struct {
	ID             string                      `json:"id"`
	Name           string                      `json:"name"`
	Type           tournament.TournamentType   `json:"type"`
	Status         tournament.TournamentStatus `json:"status"`
	GameType       string                      `json:"game_type"`
	MinPlayers     int                         `json:"min_players"`
	MaxPlayers     int                         `json:"max_players"`
	CurrentPlayers int                         `json:"current_players"`
	EntryFee       int64                       `json:"entry_fee"`
	PrizePool      int64                       `json:"prize_pool"`
	StartTime      time.Time                   `json:"start_time"`
}

// CreateTournament creates a new tournament
func (h *TournamentHandler) CreateTournament(ctx context.Context, req *CreateTournamentRequest) (*TournamentResponse, error) {
	t, err := h.service.CreateTournament(ctx, req.Name, req.Type, req.GameType, req.EntryFee, req.BuyIn, req.MinPlayers, req.MaxPlayers, req.StartTime)
	if err != nil {
		return nil, err
	}

	return &TournamentResponse{
		ID:             t.ID,
		Name:           t.Name,
		Type:           t.Type,
		Status:         t.Status,
		GameType:       t.GameType,
		MinPlayers:     t.MinPlayers,
		MaxPlayers:     t.MaxPlayers,
		CurrentPlayers: t.CurrentPlayers,
		EntryFee:       t.EntryFee,
		PrizePool:      t.PrizePool,
		StartTime:      t.StartTime,
	}, nil
}

// GetTournament retrieves a tournament
func (h *TournamentHandler) GetTournament(ctx context.Context, id string) (*TournamentResponse, error) {
	t, err := h.service.GetTournament(ctx, id)
	if err != nil {
		return nil, err
	}

	return &TournamentResponse{
		ID:             t.ID,
		Name:           t.Name,
		Type:           t.Type,
		Status:         t.Status,
		GameType:       t.GameType,
		MinPlayers:     t.MinPlayers,
		MaxPlayers:     t.MaxPlayers,
		CurrentPlayers: t.CurrentPlayers,
		EntryFee:       t.EntryFee,
		PrizePool:      t.PrizePool,
		StartTime:      t.StartTime,
	}, nil
}

// RegisterUser registers a user for a tournament
func (h *TournamentHandler) RegisterUser(ctx context.Context, req *RegisterUserRequest) error {
	return h.service.RegisterUser(ctx, req.TournamentID, req.UserID, req.Username)
}

// UnregisterUser unregisters a user from a tournament
func (h *TournamentHandler) UnregisterUser(ctx context.Context, tournamentID, userID string) error {
	return h.service.UnregisterUser(ctx, tournamentID, userID)
}

// StartTournament starts a tournament
func (h *TournamentHandler) StartTournament(ctx context.Context, tournamentID string) error {
	return h.service.StartTournament(ctx, tournamentID)
}

// EndTournament ends a tournament
func (h *TournamentHandler) EndTournament(ctx context.Context, tournamentID string) error {
	return h.service.EndTournament(ctx, tournamentID)
}

// UpdatePlayerScore updates a player's score
func (h *TournamentHandler) UpdatePlayerScore(ctx context.Context, tournamentID, userID string, scoreDelta int) error {
	return h.service.UpdatePlayerScore(ctx, tournamentID, userID, scoreDelta)
}

// EliminatePlayer eliminates a player
func (h *TournamentHandler) EliminatePlayer(ctx context.Context, tournamentID, userID string, rank int) error {
	return h.service.EliminatePlayer(ctx, tournamentID, userID, rank)
}

// ListTournaments lists tournaments
func (h *TournamentHandler) ListTournaments(ctx context.Context, status tournament.TournamentStatus, tType tournament.TournamentType) ([]*TournamentResponse, error) {
	tournaments, err := h.service.ListTournaments(ctx, status, tType)
	if err != nil {
		return nil, err
	}

	result := make([]*TournamentResponse, 0, len(tournaments))
	for _, t := range tournaments {
		result = append(result, &TournamentResponse{
			ID:             t.ID,
			Name:           t.Name,
			Type:           t.Type,
			Status:         t.Status,
			GameType:       t.GameType,
			MinPlayers:     t.MinPlayers,
			MaxPlayers:     t.MaxPlayers,
			CurrentPlayers: t.CurrentPlayers,
			EntryFee:       t.EntryFee,
			PrizePool:      t.PrizePool,
			StartTime:      t.StartTime,
		})
	}

	return result, nil
}

// GetLeaderboard gets tournament leaderboard
func (h *TournamentHandler) GetLeaderboard(ctx context.Context, tournamentID string, limit int) ([]tournament.LeaderboardEntry, error) {
	return h.service.GetLeaderboard(ctx, tournamentID, limit)
}

// GetPlayerRank gets a player's rank
func (h *TournamentHandler) GetPlayerRank(ctx context.Context, tournamentID, userID string) (*tournament.LeaderboardEntry, error) {
	return h.service.GetPlayerRank(ctx, tournamentID, userID)
}
