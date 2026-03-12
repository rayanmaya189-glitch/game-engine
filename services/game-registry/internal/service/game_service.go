package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/game-engine/game-registry/internal/config"
	"github.com/game-engine/game-registry/internal/enums"
	"github.com/game-engine/game-registry/internal/model"
	"github.com/game-engine/game-registry/internal/repository"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

// GameService handles game business logic
type GameService struct {
	repo   *repository.GameRepository
	config *config.Config
	nc     *nats.Conn
}

// NewGameService creates a new GameService
func NewGameService(repo *repository.GameRepository, cfg *config.Config, nc *nats.Conn) *GameService {
	return &GameService{
		repo:   repo,
		config: cfg,
		nc:     nc,
	}
}

// ListGames lists games with pagination and filters
func (s *GameService) ListGames(ctx context.Context, req *ListGamesRequest) (*ListGamesResponse, error) {
	filter := model.GameListFilter{
		CategoryID:       req.CategoryID,
		ProviderID:       req.ProviderID,
		Status:           enums.Status(req.Status),
		MobileSupported:  req.MobileSupported,
		DesktopSupported: req.DesktopSupported,
		IsFeatured:       req.IsFeatured,
		IsJackpot:        req.IsJackpot,
		Query:            req.Query,
		SortBy:           req.SortBy,
		Page:             int(req.Pagination.Page),
		PageSize:         int(req.Pagination.PageSize),
	}

	// Convert category and provider enums
	for _, cat := range req.Categories {
		filter.Categories = append(filter.Categories, enums.GameCategory(cat))
	}
	for _, prov := range req.Providers {
		filter.Providers = append(filter.Providers, enums.GameProvider(prov))
	}

	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PageSize < 1 || filter.PageSize > 100 {
		filter.PageSize = 20
	}

	games, pagination, err := s.repo.ListGames(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list games: %w", err)
	}

	return &ListGamesResponse{
		Games:      s.toGameSummaries(games),
		Pagination: toPaginationResponse(pagination),
	}, nil
}

// GetGame retrieves a game by ID
func (s *GameService) GetGame(ctx context.Context, gameID string) (*model.Game, error) {
	game, err := s.repo.GetGameByID(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to get game: %w", err)
	}
	return game, nil
}

// GetGameConfig retrieves game configuration
func (s *GameService) GetGameConfig(ctx context.Context, gameID, userID string, deviceType enums.DeviceType, language enums.GameLanguage, currency string, sessionID string) (*model.GameConfig, error) {
	game, err := s.repo.GetGameByID(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to get game: %w", err)
	}
	if game == nil {
		return nil, fmt.Errorf("game not found")
	}

	// Generate session token
	sessionToken := uuid.New().String()
	sessionUUID := uuid.New().String()

	// Build game URL from template
	gameURL := strings.ReplaceAll(s.config.Game.LaunchURLTemplate, "{game_id}", gameID)
	gameURL = strings.ReplaceAll(gameURL, "{token}", sessionToken)
	gameURL = strings.ReplaceAll(gameURL, "{session}", sessionUUID)

	config := &model.GameConfig{
		GameID:       gameID,
		SessionToken: sessionToken,
		GameURL:      gameURL,
		PlayerID:     userID,
		Currency:     currency,
		Language:     language.String(),
		ExpiresAt:    time.Now().Add(s.config.Game.GetSessionTTL()),
	}

	// Save config to database
	err = s.repo.CreateGameConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create game config: %w", err)
	}

	return config, nil
}

// GetGameURL generates game launch URL with session token
func (s *GameService) GetGameURL(ctx context.Context, gameID, userID string, deviceType enums.DeviceType, sessionID string, language enums.GameLanguage, currency string) (*GameURLResult, error) {
	game, err := s.repo.GetGameByID(ctx, gameID)
	if err != nil {
		return nil, fmt.Errorf("failed to get game: %w", err)
	}
	if game == nil {
		return nil, fmt.Errorf("game not found")
	}

	// Generate session token
	sessionToken := uuid.New().String()
	sessionUUID := sessionID
	if sessionUUID == "" {
		sessionUUID = uuid.New().String()
	}

	// Build game URL from template
	gameURL := strings.ReplaceAll(s.config.Game.LaunchURLTemplate, "{game_id}", gameID)
	gameURL = strings.ReplaceAll(gameURL, "{token}", sessionToken)
	gameURL = strings.ReplaceAll(gameURL, "{session}", sessionUUID)

	return &GameURLResult{
		GameURL:      gameURL,
		SessionToken: sessionToken,
		Game:         s.toGameSummary(*game),
		ExpiresAt:    time.Now().Add(s.config.Game.GetSessionTTL()),
	}, nil
}

// GetCategories retrieves all game categories
func (s *GameService) GetCategories(ctx context.Context, includeGamesCount bool) ([]model.GameCategory, error) {
	return s.repo.GetCategories(ctx, includeGamesCount)
}

// GetProviders retrieves all game providers
func (s *GameService) GetProviders(ctx context.Context, activeOnly bool) ([]model.GameProvider, error) {
	return s.repo.GetProviders(ctx, activeOnly)
}

// SearchGames searches for games
func (s *GameService) SearchGames(ctx context.Context, query string, limit int, categoryID string) ([]model.GameSummary, error) {
	if limit < 1 || limit > 50 {
		limit = 20
	}
	games, err := s.repo.SearchGames(ctx, query, limit, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to search games: %w", err)
	}
	return games, nil
}

// GetFeaturedGames retrieves featured games
func (s *GameService) GetFeaturedGames(ctx context.Context, limit int, categoryID string) ([]model.GameSummary, error) {
	if limit < 1 || limit > 50 {
		limit = 10
	}
	games, err := s.repo.GetFeaturedGames(ctx, limit, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get featured games: %w", err)
	}
	return games, nil
}

// GetPopularGames retrieves popular games
func (s *GameService) GetPopularGames(ctx context.Context, limit int, categoryID string) ([]model.GameSummary, error) {
	if limit < 1 || limit > 50 {
		limit = 10
	}
	games, err := s.repo.GetPopularGames(ctx, limit, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get popular games: %w", err)
	}
	return games, nil
}

// GetNewGames retrieves recently added games
func (s *GameService) GetNewGames(ctx context.Context, limit int, categoryID string) ([]model.GameSummary, error) {
	if limit < 1 || limit > 50 {
		limit = 10
	}
	games, err := s.repo.GetNewGames(ctx, limit, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get new games: %w", err)
	}
	return games, nil
}

// CreateGame creates a new game (admin operation)
func (s *GameService) CreateGame(ctx context.Context, game *model.Game) error {
	err := s.repo.CreateGame(ctx, game)
	if err != nil {
		return fmt.Errorf("failed to create game: %w", err)
	}

	// Publish event
	s.publishEvent("game.events.created", map[string]interface{}{
		"game_id": game.ID,
		"name":    game.Name,
	})

	return nil
}

// UpdateGame updates an existing game (admin operation)
func (s *GameService) UpdateGame(ctx context.Context, game *model.Game) error {
	err := s.repo.UpdateGame(ctx, game)
	if err != nil {
		return fmt.Errorf("failed to update game: %w", err)
	}

	// Publish event
	s.publishEvent("game.events.updated", map[string]interface{}{
		"game_id": game.ID,
		"name":    game.Name,
	})

	return nil
}

// ToggleGame enables/disables a game (admin operation)
func (s *GameService) ToggleGame(ctx context.Context, gameID string, enable bool) error {
	status := enums.StatusInactive
	if enable {
		status = enums.StatusActive
	}

	err := s.repo.ToggleGame(ctx, gameID, status)
	if err != nil {
		return fmt.Errorf("failed to toggle game: %w", err)
	}

	// Publish event
	s.publishEvent("game.events.toggled", map[string]interface{}{
		"game_id": gameID,
		"enabled": enable,
	})

	return nil
}

// SetGameOrder sets the sort order for games (admin operation)
func (s *GameService) SetGameOrder(ctx context.Context, gameID string, sortOrder int) error {
	return s.repo.SetGameOrder(ctx, gameID, sortOrder)
}

// Helper functions

func (s *GameService) toGameSummaries(games []model.GameSummary) []model.GameSummary {
	return games
}

func (s *GameService) toGameSummary(game model.Game) model.GameSummary {
	return model.GameSummary{
		GameID:           game.ID,
		Name:             game.Name,
		ProviderID:       game.ProviderID,
		ProviderName:     game.ProviderName,
		CategoryID:       game.CategoryID,
		CategoryName:     game.CategoryName,
		Type:             game.Type,
		Status:           game.Status,
		ThumbnailURL:     game.ThumbnailURL,
		BannerURL:        game.BannerURL,
		RTP:              game.RTP,
		Volatility:       game.Volatility,
		MinBet:           game.MinBet,
		MaxBet:           game.MaxBet,
		MaxWin:           game.MaxWin,
		SupportedDevices: game.SupportedDevices,
		IsFeatured:       game.IsFeatured,
		IsNew:            game.IsNew,
		IsPopular:        game.IsPopular,
		IsJackpot:        game.IsJackpot,
		LaunchURL:        game.LaunchURL,
		PopularityScore:  game.PopularityScore,
	}
}

func toPaginationResponse(p model.PaginationResult) *PaginationResponse {
	return &PaginationResponse{
		Page:       int32(p.Page),
		PageSize:   int32(p.PageSize),
		TotalCount: p.TotalCount,
		TotalPages: int32(p.TotalPages),
	}
}

func (s *GameService) publishEvent(subject string, data map[string]interface{}) {
	if s.nc != nil {
		s.nc.Publish(subject, nil)
	}
}

// Request/Response DTOs

type ListGamesRequest struct {
	CategoryID       string
	ProviderID       string
	Categories       []int32
	Providers        []int32
	Status           int32
	MobileSupported  bool
	DesktopSupported bool
	IsFeatured       bool
	IsJackpot        bool
	Query            string
	SortBy           string
	Pagination       struct {
		Page     int32
		PageSize int32
	}
}

type ListGamesResponse struct {
	Games      []model.GameSummary
	Pagination *PaginationResponse
}

type PaginationResponse struct {
	Page       int32
	PageSize   int32
	TotalCount int64
	TotalPages int32
}

type GameURLResult struct {
	GameURL      string
	SessionToken string
	Game         model.GameSummary
	ExpiresAt    time.Time
}
