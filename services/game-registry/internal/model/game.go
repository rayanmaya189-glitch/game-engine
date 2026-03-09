package model

import (
	"time"

	"github.com/gameengine/game-registry/internal/enums"
)

// Game represents a game in the catalog
type Game struct {
	ID                  string               `json:"id" db:"id"`
	Name                string               `json:"name" db:"name"`
	Description         string               `json:"description" db:"description"`
	ProviderID          string               `json:"provider_id" db:"provider_id"`
	ProviderName        string               `json:"provider_name" db:"provider_name"`
	CategoryID          string               `json:"category_id" db:"category_id"`
	CategoryName        string               `json:"category_name" db:"category_name"`
	Type                enums.GameCategory   `json:"type" db:"type"`
	Status              enums.Status         `json:"status" db:"status"`
	ThumbnailURL        string               `json:"thumbnail_url" db:"thumbnail_url"`
	BannerURL           string               `json:"banner_url" db:"banner_url"`
	RTP                 float64              `json:"rtp" db:"rtp"`
	Volatility          string               `json:"volatility" db:"volatility"`
	MinBet              int64                `json:"min_bet" db:"min_bet"`
	MaxBet              int64                `json:"max_bet" db:"max_bet"`
	MaxWin              string               `json:"max_win" db:"max_win"`
	Paylines            int                  `json:"paylines" db:"paylines"`
	Reels               int                  `json:"reels" db:"reels"`
	Features            []string             `json:"features" db:"features"`
	SupportedDevices    []enums.DeviceType   `json:"supported_devices" db:"supported_devices"`
	SupportedLanguages  []enums.GameLanguage `json:"supported_languages" db:"supported_languages"`
	SupportedCurrencies []string             `json:"supported_currencies" db:"supported_currencies"`
	IsFeatured          bool                 `json:"is_featured" db:"is_featured"`
	IsNew               bool                 `json:"is_new" db:"is_new"`
	IsPopular           bool                 `json:"is_popular" db:"is_popular"`
	IsJackpot           bool                 `json:"is_jackpot" db:"is_jackpot"`
	LaunchURL           string               `json:"launch_url" db:"launch_url"`
	ReleaseDate         *time.Time           `json:"release_date" db:"release_date"`
	PopularityScore     int                  `json:"popularity_score" db:"popularity_score"`
	SortOrder           int                  `json:"sort_order" db:"sort_order"`
	CreatedAt           time.Time            `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at" db:"updated_at"`
}

// GameSummary represents a summary of a game for listing
type GameSummary struct {
	GameID           string             `json:"game_id" db:"game_id"`
	Name             string             `json:"name" db:"name"`
	ProviderID       string             `json:"provider_id" db:"provider_id"`
	ProviderName     string             `json:"provider_name" db:"provider_name"`
	CategoryID       string             `json:"category_id" db:"category_id"`
	CategoryName     string             `json:"category_name" db:"category_name"`
	Type             enums.GameCategory `json:"type" db:"type"`
	Status           enums.Status       `json:"status" db:"status"`
	ThumbnailURL     string             `json:"thumbnail_url" db:"thumbnail_url"`
	BannerURL        string             `json:"banner_url" db:"banner_url"`
	RTP              float64            `json:"rtp" db:"rtp"`
	Volatility       string             `json:"volatility" db:"volatility"`
	MinBet           int64              `json:"min_bet" db:"min_bet"`
	MaxBet           int64              `json:"max_bet" db:"max_bet"`
	MaxWin           string             `json:"max_win" db:"max_win"`
	SupportedDevices []enums.DeviceType `json:"supported_devices" db:"supported_devices"`
	IsFeatured       bool               `json:"is_featured" db:"is_featured"`
	IsNew            bool               `json:"is_new" db:"is_new"`
	IsPopular        bool               `json:"is_popular" db:"is_popular"`
	IsJackpot        bool               `json:"is_jackpot" db:"is_jackpot"`
	LaunchURL        string             `json:"launch_url" db:"launch_url"`
	PopularityScore  int                `json:"popularity_score" db:"popularity_score"`
}

// GameConfig represents game configuration
type GameConfig struct {
	ID           string    `json:"id" db:"id"`
	GameID       string    `json:"game_id" db:"game_id"`
	SessionToken string    `json:"session_token" db:"session_token"`
	GameURL      string    `json:"game_url" db:"game_url"`
	PlayerID     string    `json:"player_id" db:"player_id"`
	Balance      int64     `json:"balance" db:"balance"`
	Currency     string    `json:"currency" db:"currency"`
	Language     string    `json:"language" db:"language"`
	ConfigJSON   string    `json:"config_json" db:"config_json"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	ExpiresAt    time.Time `json:"expires_at" db:"expires_at"`
}

// GameCategory represents a game category
type GameCategory struct {
	ID          string       `json:"id" db:"id"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	IconURL     string       `json:"icon_url" db:"icon_url"`
	BannerURL   string       `json:"banner_url" db:"banner_url"`
	ParentID    string       `json:"parent_id" db:"parent_id"`
	SortOrder   int          `json:"sort_order" db:"sort_order"`
	Status      enums.Status `json:"status" db:"status"`
	IsFeatured  bool         `json:"is_featured" db:"is_featured"`
	GamesCount  int          `json:"games_count" db:"games_count"`
	Slug        string       `json:"slug" db:"slug"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

// GameProvider represents a game provider
type GameProvider struct {
	ID          string       `json:"id" db:"id"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	LogoURL     string       `json:"logo_url" db:"logo_url"`
	WebsiteURL  string       `json:"website_url" db:"website_url"`
	Status      enums.Status `json:"status" db:"status"`
	GamesCount  int          `json:"games_count" db:"games_count"`
	License     string       `json:"license" db:"license"`
	Established int          `json:"established" db:"established"`
	IsFeatured  bool         `json:"is_featured" db:"is_featured"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

// GameTag represents a searchable tag for games
type GameTag struct {
	ID        string    `json:"id" db:"id"`
	GameID    string    `json:"game_id" db:"game_id"`
	Tag       string    `json:"tag" db:"tag"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// GameListFilter represents filters for listing games
type GameListFilter struct {
	CategoryID       string
	ProviderID       string
	Categories       []enums.GameCategory
	Providers        []enums.GameProvider
	Status           enums.Status
	MobileSupported  bool
	DesktopSupported bool
	IsFeatured       bool
	IsJackpot        bool
	Query            string
	SortBy           string
	Page             int
	PageSize         int
}

// PaginationResult represents pagination result
type PaginationResult struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalCount int64 `json:"total_count"`
	TotalPages int   `json:"total_pages"`
}
