package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/game_engine/game-registry/internal/enums"
	"github.com/game_engine/game-registry/internal/model"
)

// GetFeaturedGames retrieves featured games
func (r *GameRepository) GetFeaturedGames(ctx context.Context, limit int, categoryID string) ([]model.GameSummary, error) {
	// Try to get from cache first
	cacheKey := fmt.Sprintf("games:featured:%d:%s", limit, categoryID)
	cached, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var games []model.GameSummary
		json.Unmarshal([]byte(cached), &games)
		return games, nil
	}

	query := `SELECT g.id, g.name, g.provider_id, gp.name as provider_name,
		g.category_id, gc.name as category_name, g.type, g.status,
		g.thumbnail_url, g.banner_url, g.rtp, g.volatility, g.min_bet, g.max_bet,
		g.max_win, g.supported_devices, g.is_featured, g.is_new, g.is_popular,
		g.is_jackpot, g.launch_url, g.popularity_score
		FROM games g
		LEFT JOIN game_providers gp ON g.provider_id = gp.id
		LEFT JOIN game_categories gc ON g.category_id = gc.id
		WHERE g.status = 2 AND g.is_featured = true`

	args := []interface{}{}

	if categoryID != "" {
		query += fmt.Sprintf(" AND g.category_id = $%d", len(args)+1)
		args = append(args, categoryID)
	}

	query += fmt.Sprintf(" ORDER BY g.sort_order LIMIT $%d", len(args)+1)
	args = append(args, limit)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get featured games: %w", err)
	}
	defer rows.Close()

	var games []model.GameSummary
	for rows.Next() {
		var g model.GameSummary
		var devicesJSON string
		err := rows.Scan(
			&g.GameID, &g.Name, &g.ProviderID, &g.ProviderName,
			&g.CategoryID, &g.CategoryName, &g.Type, &g.Status,
			&g.ThumbnailURL, &g.BannerURL, &g.RTP, &g.Volatility,
			&g.MinBet, &g.MaxBet, &g.MaxWin, &devicesJSON,
			&g.IsFeatured, &g.IsNew, &g.IsPopular, &g.IsJackpot,
			&g.LaunchURL, &g.PopularityScore,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan game: %w", err)
		}

		if devicesJSON != "" {
			var devices []enums.DeviceType
			json.Unmarshal([]byte(devicesJSON), &devices)
			g.SupportedDevices = devices
		}

		games = append(games, g)
	}

	// Cache the result
	if len(games) > 0 {
		gamesJSON, _ := json.Marshal(games)
		r.redis.Set(ctx, cacheKey, gamesJSON, 10*time.Minute)
	}

	return games, nil
}

// GetPopularGames retrieves popular games
func (r *GameRepository) GetPopularGames(ctx context.Context, limit int, categoryID string) ([]model.GameSummary, error) {
	// Try to get from cache first
	cacheKey := fmt.Sprintf("games:popular:%d:%s", limit, categoryID)
	cached, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var games []model.GameSummary
		json.Unmarshal([]byte(cached), &games)
		return games, nil
	}

	query := `SELECT g.id, g.name, g.provider_id, gp.name as provider_name,
		g.category_id, gc.name as category_name, g.type, g.status,
		g.thumbnail_url, g.banner_url, g.rtp, g.volatility, g.min_bet, g.max_bet,
		g.max_win, g.supported_devices, g.is_featured, g.is_new, g.is_popular,
		g.is_jackpot, g.launch_url, g.popularity_score
		FROM games g
		LEFT JOIN game_providers gp ON g.provider_id = gp.id
		LEFT JOIN game_categories gc ON g.category_id = gc.id
		WHERE g.status = 2`

	args := []interface{}{}

	if categoryID != "" {
		query += fmt.Sprintf(" AND g.category_id = $%d", len(args)+1)
		args = append(args, categoryID)
	}

	query += fmt.Sprintf(" ORDER BY g.popularity_score DESC LIMIT $%d", len(args)+1)
	args = append(args, limit)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get popular games: %w", err)
	}
	defer rows.Close()

	var games []model.GameSummary
	for rows.Next() {
		var g model.GameSummary
		var devicesJSON string
		err := rows.Scan(
			&g.GameID, &g.Name, &g.ProviderID, &g.ProviderName,
			&g.CategoryID, &g.CategoryName, &g.Type, &g.Status,
			&g.ThumbnailURL, &g.BannerURL, &g.RTP, &g.Volatility,
			&g.MinBet, &g.MaxBet, &g.MaxWin, &devicesJSON,
			&g.IsFeatured, &g.IsNew, &g.IsPopular, &g.IsJackpot,
			&g.LaunchURL, &g.PopularityScore,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan game: %w", err)
		}

		if devicesJSON != "" {
			var devices []enums.DeviceType
			json.Unmarshal([]byte(devicesJSON), &devices)
			g.SupportedDevices = devices
		}

		games = append(games, g)
	}

	// Cache the result
	if len(games) > 0 {
		gamesJSON, _ := json.Marshal(games)
		r.redis.Set(ctx, cacheKey, gamesJSON, 10*time.Minute)
	}

	return games, nil
}

// GetNewGames retrieves recently added games
func (r *GameRepository) GetNewGames(ctx context.Context, limit int, categoryID string) ([]model.GameSummary, error) {
	query := `SELECT g.id, g.name, g.provider_id, gp.name as provider_name,
		g.category_id, gc.name as category_name, g.type, g.status,
		g.thumbnail_url, g.banner_url, g.rtp, g.volatility, g.min_bet, g.max_bet,
		g.max_win, g.supported_devices, g.is_featured, g.is_new, g.is_popular,
		g.is_jackpot, g.launch_url, g.popularity_score
		FROM games g
		LEFT JOIN game_providers gp ON g.provider_id = gp.id
		LEFT JOIN game_categories gc ON g.category_id = gc.id
		WHERE g.status = 2 AND g.is_new = true`

	args := []interface{}{}

	if categoryID != "" {
		query += fmt.Sprintf(" AND g.category_id = $%d", len(args)+1)
		args = append(args, categoryID)
	}

	query += fmt.Sprintf(" ORDER BY g.created_at DESC LIMIT $%d", len(args)+1)
	args = append(args, limit)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get new games: %w", err)
	}
	defer rows.Close()

	var games []model.GameSummary
	for rows.Next() {
		var g model.GameSummary
		var devicesJSON string
		err := rows.Scan(
			&g.GameID, &g.Name, &g.ProviderID, &g.ProviderName,
			&g.CategoryID, &g.CategoryName, &g.Type, &g.Status,
			&g.ThumbnailURL, &g.BannerURL, &g.RTP, &g.Volatility,
			&g.MinBet, &g.MaxBet, &g.MaxWin, &devicesJSON,
			&g.IsFeatured, &g.IsNew, &g.IsPopular, &g.IsJackpot,
			&g.LaunchURL, &g.PopularityScore,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan game: %w", err)
		}

		if devicesJSON != "" {
			var devices []enums.DeviceType
			json.Unmarshal([]byte(devicesJSON), &devices)
			g.SupportedDevices = devices
		}

		games = append(games, g)
	}

	return games, nil
}
