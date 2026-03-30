package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/game_engine/game-registry/internal/enums"
	"github.com/game_engine/game-registry/internal/model"
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
