package config

import (
	"fmt"
	"os"
)

// Config holds all configuration for the service
type Config struct {
	Server      ServerConfig
	Database    DatabaseConfig
	Redis       RedisConfig
	Leaderboard LeaderboardConfig
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port int
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// LeaderboardConfig holds leaderboard-specific configuration
type LeaderboardConfig struct {
	TopPlayersCount    int
	CacheTTLSeconds    int
	ResetIntervalHours int
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
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
		},
	}, nil
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

// ConnectionString returns PostgreSQL connection string
func (d *DatabaseConfig) ConnectionString() string {
	return "host=" + d.Host +
		" port=" + string(rune(d.Port)) +
		" user=" + d.User +
		" password=" + d.Password +
		" dbname=" + d.DBName +
		" sslmode=" + d.SSLMode
}
