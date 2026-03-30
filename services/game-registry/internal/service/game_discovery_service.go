package service

import (
	"context"
	"fmt"
	"time"

	"github.com/game_engine/game-registry/internal/model"
)

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
