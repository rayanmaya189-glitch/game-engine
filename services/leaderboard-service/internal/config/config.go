package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds all configuration for the service
type Config struct {
	Server      ServerConfig      `mapstructure:"server"`
	Database    DatabaseConfig    `mapstructure:"database"`
	Redis       RedisConfig       `mapstructure:"redis"`
	Leaderboard LeaderboardConfig `mapstructure:"leaderboard"`
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// LeaderboardConfig holds leaderboard-specific configuration
type LeaderboardConfig struct {
	TopPlayersCount    int                      `mapstructure:"top_players_count"`
	CacheTTLSeconds    int                      `mapstructure:"cache_ttl_seconds"`
	ResetIntervalHours int                      `mapstructure:"reset_interval_hours"`
	MinBetThreshold    float64                  `mapstructure:"min_bet_threshold"`
	Prizes             map[string][]PrizeConfig `mapstructure:"prizes"`
	PrizeAutoCredit    bool                     `mapstructure:"prize_auto_credit"`
}

// PrizeConfig holds prize configuration for a leaderboard
type PrizeConfig struct {
	FromRank int     `mapstructure:"from_rank"`
	ToRank   int     `mapstructure:"to_rank"`
	Type     string  `mapstructure:"type"` // bonus, freespins, vip_points
	Value    float64 `mapstructure:"value"`
}

// Load loads configuration from config.yaml and environment variables
func Load() (*Config, error) {
	cfg := &Config{}

	// First, try to load from config.yaml
	configFile := getEnv("CONFIG_FILE", "config.yaml")
	if data, err := os.ReadFile(configFile); err == nil {
		if err := yaml.Unmarshal(data, cfg); err == nil {
			// Override with environment variables
			applyEnvOverrides(cfg)
			return cfg, nil
		}
	}

	// Fallback to environment variables
	return loadFromEnv(), nil
}

func loadFromEnv() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnvAsInt("SERVER_PORT", 8080),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "leaderboard"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		Leaderboard: LeaderboardConfig{
			TopPlayersCount:    getEnvAsInt("LEADERBOARD_TOP_COUNT", 100),
			CacheTTLSeconds:    getEnvAsInt("LEADERBOARD_CACHE_TTL", 300),
			ResetIntervalHours: getEnvAsInt("LEADERBOARD_RESET_HOURS", 24),
			MinBetThreshold:    getEnvAsFloat("LEADERBOARD_MIN_BET", 0.50),
			PrizeAutoCredit:    getEnvAsBool("LEADERBOARD_PRIZE_AUTO_CREDIT", false),
		},
	}
}

func applyEnvOverrides(cfg *Config) {
	if port := getEnvAsInt("SERVER_PORT", 0); port > 0 {
		cfg.Server.Port = port
	}
	if host := getEnv("DB_HOST", ""); host != "" {
		cfg.Database.Host = host
	}
	if port := getEnvAsInt("DB_PORT", 0); port > 0 {
		cfg.Database.Port = port
	}
	if user := getEnv("DB_USER", ""); user != "" {
		cfg.Database.User = user
	}
	if password := getEnv("DB_PASSWORD", ""); password != "" {
		cfg.Database.Password = password
	}
	if dbname := getEnv("DB_NAME", ""); dbname != "" {
		cfg.Database.DBName = dbname
	}
	if sslmode := getEnv("DB_SSLMODE", ""); sslmode != "" {
		cfg.Database.SSLMode = sslmode
	}
	if host := getEnv("REDIS_HOST", ""); host != "" {
		cfg.Redis.Host = host
	}
	if port := getEnvAsInt("REDIS_PORT", 0); port > 0 {
		cfg.Redis.Port = port
	}
	if password := getEnv("REDIS_PASSWORD", ""); password != "" {
		cfg.Redis.Password = password
	}
	if db := getEnvAsInt("REDIS_DB", -1); db >= 0 {
		cfg.Redis.DB = db
	}
	if count := getEnvAsInt("LEADERBOARD_TOP_COUNT", 0); count > 0 {
		cfg.Leaderboard.TopPlayersCount = count
	}
	if ttl := getEnvAsInt("LEADERBOARD_CACHE_TTL", 0); ttl > 0 {
		cfg.Leaderboard.CacheTTLSeconds = ttl
	}
	if reset := getEnvAsInt("LEADERBOARD_RESET_HOURS", 0); reset > 0 {
		cfg.Leaderboard.ResetIntervalHours = reset
	}
	if minBet := getEnvAsFloat("LEADERBOARD_MIN_BET", 0); minBet > 0 {
		cfg.Leaderboard.MinBetThreshold = minBet
	}
	if autoCredit := getEnv("LEADERBOARD_PRIZE_AUTO_CREDIT", ""); autoCredit != "" {
		cfg.Leaderboard.PrizeAutoCredit = autoCredit == "true"
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := parseInt(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func parseInt(s string) (int, error) {
	var result int
	_, err := fmt.Sscanf(s, "%d", &result)
	return result, err
}

func parseFloat(s string) (float64, error) {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	return result, err
}

func getEnvAsFloat(key string, defaultValue float64) float64 {
	if value, exists := os.LookupEnv(key); exists {
		if floatValue, err := parseFloat(value); err == nil {
			return floatValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return value == "true" || value == "1" || value == "yes"
	}
	return defaultValue
}

// ConnectionString returns PostgreSQL connection string
func (d *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.DBName, d.SSLMode)
}
