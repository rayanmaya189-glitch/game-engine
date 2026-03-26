package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds all configuration for the service
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Winners  WinnersConfig  `mapstructure:"winners"`
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

// WinnersConfig holds winners showcase specific configuration
type WinnersConfig struct {
	RecentWinnersLimit  int     `mapstructure:"recent_winners_limit"`
	BigWinThreshold     float64 `mapstructure:"big_win_threshold"`
	JackpotThreshold    float64 `mapstructure:"jackpot_threshold"`
	FeedTTLSeconds      int     `mapstructure:"feed_ttl_seconds"`
	EnableJackpotAlerts bool    `mapstructure:"enable_jackpot_alerts"`
	AnonymizeNames      bool    `mapstructure:"anonymize_names"`
}

// Load loads configuration from config.yaml and environment variables
func Load() (*Config, error) {
	cfg := &Config{}

	// First, try to load from config.yaml
	configFile := getEnv("CONFIG_FILE", "config.yaml")
	if data, err := os.ReadFile(configFile); err == nil {
		if err := yaml.Unmarshal(data, cfg); err == nil {
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
			Port: getEnvAsInt("SERVER_PORT", 8082),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "winners"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvAsInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		Winners: WinnersConfig{
			RecentWinnersLimit:  getEnvAsInt("RECENT_WINNERS_LIMIT", 50),
			BigWinThreshold:     getEnvAsFloat("BIG_WIN_THRESHOLD", 100.0),
			JackpotThreshold:    getEnvAsFloat("JACKPOT_THRESHOLD", 1000.0),
			FeedTTLSeconds:      getEnvAsInt("FEED_TTL_SECONDS", 60),
			EnableJackpotAlerts: getEnvAsBool("ENABLE_JACKPOT_ALERTS", true),
			AnonymizeNames:      getEnvAsBool("ANONYMIZE_NAMES", true),
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
	if limit := getEnvAsInt("RECENT_WINNERS_LIMIT", 0); limit > 0 {
		cfg.Winners.RecentWinnersLimit = limit
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

func getEnvAsFloat(key string, defaultValue float64) float64 {
	if value, exists := os.LookupEnv(key); exists {
		if floatValue, err := parseFloat(value); err == nil {
			return floatValue
		}
	}
	return defaultValue
}

func parseFloat(s string) (float64, error) {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	return result, err
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
