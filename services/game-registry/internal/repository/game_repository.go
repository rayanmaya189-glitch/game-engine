package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gameengine/game-registry/internal/enums"
	"github.com/gameengine/game-registry/internal/model"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// GameRepository handles database operations for games
type GameRepository struct {
	db    *sql.DB
	redis *redis.Client
}

// NewGameRepository creates a new GameRepository
func NewGameRepository(db *sql.DB, redis *redis.Client) *GameRepository {
	return &GameRepository{
		db:    db,
		redis: redis,
	}
}

// ListGames retrieves games with pagination and filters
func (r *GameRepository) ListGames(ctx context.Context, filter model.GameListFilter) ([]model.GameSummary, model.PaginationResult, error) {
	// Build query
	baseQuery := `SELECT g.id, g.name, g.provider_id, gp.name as provider_name, 
		g.category_id, gc.name as category_name, g.type, g.status, 
		g.thumbnail_url, g.banner_url, g.rtp, g.volatility, g.min_bet, g.max_bet, 
		g.max_win, g.supported_devices, g.is_featured, g.is_new, g.is_popular, 
		g.is_jackpot, g.launch_url, g.popularity_score
		FROM games g
		LEFT JOIN game_providers gp ON g.provider_id = gp.id
		LEFT JOIN game_categories gc ON g.category_id = gc.id
		WHERE 1=1`

	var args []interface{}
	argIndex := 1

	// Apply filters
	if filter.CategoryID != "" {
		baseQuery += fmt.Sprintf(" AND g.category_id = $%d", argIndex)
		args = append(args, filter.CategoryID)
		argIndex++
	}

	if filter.ProviderID != "" {
		baseQuery += fmt.Sprintf(" AND g.provider_id = $%d", argIndex)
		args = append(args, filter.ProviderID)
		argIndex++
	}

	if len(filter.Categories) > 0 {
		placeholders := make([]string, len(filter.Categories))
		for i, cat := range filter.Categories {
			placeholders[i] = fmt.Sprintf("$%d", argIndex)
			args = append(args, int(cat))
			argIndex++
		}
		baseQuery += fmt.Sprintf(" AND g.type IN (%s)", strings.Join(placeholders, ","))
	}

	if len(filter.Providers) > 0 {
		placeholders := make([]string, len(filter.Providers))
		for i, prov := range filter.Providers {
			placeholders[i] = fmt.Sprintf("$%d", argIndex)
			args = append(args, int(prov))
			argIndex++
		}
		baseQuery += fmt.Sprintf(" AND g.provider_id IN (%s)", strings.Join(placeholders, ","))
	}

	if filter.Status != 0 {
		baseQuery += fmt.Sprintf(" AND g.status = $%d", argIndex)
		args = append(args, int(filter.Status))
		argIndex++
	}

	if filter.IsFeatured {
		baseQuery += " AND g.is_featured = true"
	}

	if filter.IsJackpot {
		baseQuery += " AND g.is_jackpot = true"
	}

	if filter.Query != "" {
		baseQuery += fmt.Sprintf(" AND (g.name ILIKE $%d OR gp.name ILIKE $%d)", argIndex, argIndex)
		args = append(args, "%"+filter.Query+"%")
		argIndex++
	}

	// Get total count
	countQuery := strings.Replace(baseQuery, "SELECT g.id, g.name, gp.name as provider_name, gc.name as category_name, g.type, g.status, g.thumbnail_url, g.banner_url, g.rtp, g.volatility, g.min_bet, g.max_bet, g.max_win, g.supported_devices, g.is_featured, g.is_new, g.is_popular, g.is_jackpot, g.launch_url, g.popularity_score", "SELECT COUNT(*)", 1)

	var totalCount int64
	err := r.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		return nil, model.PaginationResult{}, fmt.Errorf("failed to count games: %w", err)
	}

	// Apply sorting
	sortColumn := "g.sort_order"
	sortOrder := "ASC"
	switch filter.SortBy {
	case "name":
		sortColumn = "g.name"
	case "popularity":
		sortColumn = "g.popularity_score"
	case "newest":
		sortColumn = "g.release_date"
		sortOrder = "DESC"
	case "rtp":
		sortColumn = "g.rtp"
	}
	baseQuery += fmt.Sprintf(" ORDER BY %s %s", sortColumn, sortOrder)

	// Apply pagination
	offset := (filter.Page - 1) * filter.PageSize
	baseQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, filter.PageSize, offset)

	// Execute query
	rows, err := r.db.QueryContext(ctx, baseQuery, args...)
	if err != nil {
		return nil, model.PaginationResult{}, fmt.Errorf("failed to query games: %w", err)
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
			return nil, model.PaginationResult{}, fmt.Errorf("failed to scan game: %w", err)
		}

		// Parse devices JSON
		if devicesJSON != "" {
			var devices []enums.DeviceType
			json.Unmarshal([]byte(devicesJSON), &devices)
			g.SupportedDevices = devices
		}

		games = append(games, g)
	}

	totalPages := int(totalCount) / filter.PageSize
	if int(totalCount)%filter.PageSize > 0 {
		totalPages++
	}

	return games, model.PaginationResult{
		Page:       filter.Page,
		PageSize:   filter.PageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}, nil
}

// GetGameByID retrieves a game by ID
func (r *GameRepository) GetGameByID(ctx context.Context, gameID string) (*model.Game, error) {
	query := `SELECT g.id, g.name, g.description, g.provider_id, gp.name as provider_name,
		g.category_id, gc.name as category_name, g.type, g.status, g.thumbnail_url, 
		g.banner_url, g.rtp, g.volatility, g.min_bet, g.max_bet, g.max_win,
		g.paylines, g.reels, g.features, g.supported_devices, g.supported_languages,
		g.supported_currencies, g.is_featured, g.is_new, g.is_popular, g.is_jackpot,
		g.launch_url, g.release_date, g.popularity_score, g.sort_order, g.created_at, g.updated_at
		FROM games g
		LEFT JOIN game_providers gp ON g.provider_id = gp.id
		LEFT JOIN game_categories gc ON g.category_id = gc.id
		WHERE g.id = $1`

	var g model.Game
	var featuresJSON, devicesJSON, languagesJSON, currenciesJSON string

	err := r.db.QueryRowContext(ctx, query, gameID).Scan(
		&g.ID, &g.Name, &g.Description, &g.ProviderID, &g.ProviderName,
		&g.CategoryID, &g.CategoryName, &g.Type, &g.Status, &g.ThumbnailURL,
		&g.BannerURL, &g.RTP, &g.Volatility, &g.MinBet, &g.MaxBet, &g.MaxWin,
		&g.Paylines, &g.Reels, &featuresJSON, &devicesJSON, &languagesJSON,
		&currenciesJSON, &g.IsFeatured, &g.IsNew, &g.IsPopular, &g.IsJackpot,
		&g.LaunchURL, &g.ReleaseDate, &g.PopularityScore, &g.SortOrder,
		&g.CreatedAt, &g.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get game: %w", err)
	}

	// Parse JSON fields
	if featuresJSON != "" {
		json.Unmarshal([]byte(featuresJSON), &g.Features)
	}
	if devicesJSON != "" {
		json.Unmarshal([]byte(devicesJSON), &g.SupportedDevices)
	}
	if languagesJSON != "" {
		json.Unmarshal([]byte(languagesJSON), &g.SupportedLanguages)
	}
	if currenciesJSON != "" {
		json.Unmarshal([]byte(currenciesJSON), &g.SupportedCurrencies)
	}

	return &g, nil
}

// GetGameConfig retrieves game configuration
func (r *GameRepository) GetGameConfig(ctx context.Context, gameID string) (*model.GameConfig, error) {
	query := `SELECT id, game_id, session_token, game_url, player_id, balance, 
		currency, language, config_json, created_at, expires_at
		FROM game_config WHERE game_id = $1`

	var gc model.GameConfig
	err := r.db.QueryRowContext(ctx, query, gameID).Scan(
		&gc.ID, &gc.GameID, &gc.SessionToken, &gc.GameURL, &gc.PlayerID,
		&gc.Balance, &gc.Currency, &gc.Language, &gc.ConfigJSON,
		&gc.CreatedAt, &gc.ExpiresAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get game config: %w", err)
	}

	return &gc, nil
}

// CreateGame creates a new game
func (r *GameRepository) CreateGame(ctx context.Context, game *model.Game) error {
	game.ID = uuid.New().String()
	game.CreatedAt = time.Now()
	game.UpdatedAt = time.Now()

	featuresJSON, _ := json.Marshal(game.Features)
	devicesJSON, _ := json.Marshal(game.SupportedDevices)
	languagesJSON, _ := json.Marshal(game.SupportedLanguages)
	currenciesJSON, _ := json.Marshal(game.SupportedCurrencies)

	query := `INSERT INTO games (id, name, description, provider_id, category_id, type, status,
		thumbnail_url, banner_url, rtp, volatility, min_bet, max_bet, max_win, paylines, reels,
		features, supported_devices, supported_languages, supported_currencies,
		is_featured, is_new, is_popular, is_jackpot, launch_url, release_date,
		popularity_score, sort_order, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
		$17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31)`

	_, err := r.db.ExecContext(ctx, query,
		game.ID, game.Name, game.Description, game.ProviderID, game.CategoryID,
		game.Type, game.Status, game.ThumbnailURL, game.BannerURL, game.RTP,
		game.Volatility, game.MinBet, game.MaxBet, game.MaxWin, game.Paylines,
		game.Reels, featuresJSON, devicesJSON, languagesJSON, currenciesJSON,
		game.IsFeatured, game.IsNew, game.IsPopular, game.IsJackpot,
		game.LaunchURL, game.ReleaseDate, game.PopularityScore, game.SortOrder,
		game.CreatedAt, game.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create game: %w", err)
	}

	return nil
}

// UpdateGame updates an existing game
func (r *GameRepository) UpdateGame(ctx context.Context, game *model.Game) error {
	game.UpdatedAt = time.Now()

	featuresJSON, _ := json.Marshal(game.Features)
	devicesJSON, _ := json.Marshal(game.SupportedDevices)
	languagesJSON, _ := json.Marshal(game.SupportedLanguages)
	currenciesJSON, _ := json.Marshal(game.SupportedCurrencies)

	query := `UPDATE games SET name = $1, description = $2, provider_id = $3, category_id = $4,
		type = $5, status = $6, thumbnail_url = $7, banner_url = $8, rtp = $9,
		volatility = $10, min_bet = $11, max_bet = $12, max_win = $13, paylines = $14,
		reels = $15, features = $16, supported_devices = $17, supported_languages = $18,
		supported_currencies = $19, is_featured = $20, is_new = $21, is_popular = $22,
		is_jackpot = $23, launch_url = $24, release_date = $25, popularity_score = $26,
		sort_order = $27, updated_at = $28 WHERE id = $29`

	_, err := r.db.ExecContext(ctx, query,
		game.Name, game.Description, game.ProviderID, game.CategoryID,
		game.Type, game.Status, game.ThumbnailURL, game.BannerURL, game.RTP,
		game.Volatility, game.MinBet, game.MaxBet, game.MaxWin, game.Paylines,
		game.Reels, featuresJSON, devicesJSON, languagesJSON, currenciesJSON,
		game.IsFeatured, game.IsNew, game.IsPopular, game.IsJackpot,
		game.LaunchURL, game.ReleaseDate, game.PopularityScore, game.SortOrder,
		game.UpdatedAt, game.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update game: %w", err)
	}

	return nil
}

// ToggleGame enables/disables a game
func (r *GameRepository) ToggleGame(ctx context.Context, gameID string, status enums.Status) error {
	query := `UPDATE games SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, status, time.Now(), gameID)
	if err != nil {
		return fmt.Errorf("failed to toggle game: %w", err)
	}
	return nil
}

// SetGameOrder sets the sort order for games
func (r *GameRepository) SetGameOrder(ctx context.Context, gameID string, sortOrder int) error {
	query := `UPDATE games SET sort_order = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, sortOrder, time.Now(), gameID)
	if err != nil {
		return fmt.Errorf("failed to set game order: %w", err)
	}
	return nil
}

// GetCategories retrieves all game categories
func (r *GameRepository) GetCategories(ctx context.Context, includeGamesCount bool) ([]model.GameCategory, error) {
	query := `SELECT id, name, description, icon_url, banner_url, parent_id, 
		sort_order, status, is_featured, slug, created_at, updated_at`

	if includeGamesCount {
		query += `, (SELECT COUNT(*) FROM games WHERE category_id = game_categories.id) as games_count`
	}

	query += ` FROM game_categories ORDER BY sort_order`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer rows.Close()

	var categories []model.GameCategory
	for rows.Next() {
		var c model.GameCategory
		err := rows.Scan(
			&c.ID, &c.Name, &c.Description, &c.IconURL, &c.BannerURL,
			&c.ParentID, &c.SortOrder, &c.Status, &c.IsFeatured,
			&c.Slug, &c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, c)
	}

	return categories, nil
}

// GetProviders retrieves all game providers
func (r *GameRepository) GetProviders(ctx context.Context, activeOnly bool) ([]model.GameProvider, error) {
	query := `SELECT id, name, description, logo_url, website_url, status, 
		games_count, license, established, is_featured, created_at, updated_at
		FROM game_providers`

	if activeOnly {
		query += ` WHERE status = 2` // STATUS_ACTIVE
	}

	query += ` ORDER BY name`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get providers: %w", err)
	}
	defer rows.Close()

	var providers []model.GameProvider
	for rows.Next() {
		var p model.GameProvider
		err := rows.Scan(
			&p.ID, &p.Name, &p.Description, &p.LogoURL, &p.WebsiteURL,
			&p.Status, &p.GamesCount, &p.License, &p.Established,
			&p.IsFeatured, &p.CreatedAt, &p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan provider: %w", err)
		}
		providers = append(providers, p)
	}

	return providers, nil
}

// SearchGames searches for games by name, provider, or tags
func (r *GameRepository) SearchGames(ctx context.Context, query string, limit int, categoryID string) ([]model.GameSummary, error) {
	sqlQuery := `SELECT g.id, g.name, g.provider_id, gp.name as provider_name,
		g.category_id, gc.name as category_name, g.type, g.status,
		g.thumbnail_url, g.banner_url, g.rtp, g.volatility, g.min_bet, g.max_bet,
		g.max_win, g.supported_devices, g.is_featured, g.is_new, g.is_popular,
		g.is_jackpot, g.launch_url, g.popularity_score
		FROM games g
		LEFT JOIN game_providers gp ON g.provider_id = gp.id
		LEFT JOIN game_categories gc ON g.category_id = gc.id
		LEFT JOIN game_tags gt ON g.id = gt.game_id
		WHERE g.status = 2 AND (g.name ILIKE $1 OR gp.name ILIKE $1 OR gt.tag ILIKE $1)`

	args := []interface{}{"%" + query + "%"}

	if categoryID != "" {
		sqlQuery += fmt.Sprintf(" AND g.category_id = $%d", len(args)+1)
		args = append(args, categoryID)
	}

	sqlQuery += fmt.Sprintf(" LIMIT $%d", len(args)+1)
	args = append(args, limit)

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search games: %w", err)
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

// CreateGameConfig creates game configuration
func (r *GameRepository) CreateGameConfig(ctx context.Context, config *model.GameConfig) error {
	config.ID = uuid.New().String()
	config.CreatedAt = time.Now()

	query := `INSERT INTO game_config (id, game_id, session_token, game_url, player_id,
		balance, currency, language, config_json, created_at, expires_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.ExecContext(ctx, query,
		config.ID, config.GameID, config.SessionToken, config.GameURL,
		config.PlayerID, config.Balance, config.Currency, config.Language,
		config.ConfigJSON, config.CreatedAt, config.ExpiresAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create game config: %w", err)
	}

	return nil
}
