package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	NATS     NATSConfig     `mapstructure:"nats"`
	Game     GameConfig     `mapstructure:"game"`
	Logging  LoggingConfig  `mapstructure:"logging"`
}

// AppConfig holds application-level configuration
type AppConfig struct {
	Name     string `mapstructure:"name"`
	Env      string `mapstructure:"env"`
	Port     int    `mapstructure:"port"`
	GRPCPort int    `mapstructure:"grpc_port"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Name         string `mapstructure:"name"`
	SSLMode      string `mapstructure:"ssl_mode"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// NATSConfig holds NATS configuration
type NATSConfig struct {
	URL       string           `mapstructure:"url"`
	ClusterID string           `mapstructure:"cluster_id"`
	ClientID  string           `mapstructure:"client_id"`
	Events    NATSEventsConfig `mapstructure:"events"`
}

// NATSEventsConfig holds NATS events configuration
type NATSEventsConfig struct {
	Prefix string `mapstructure:"prefix"`
}

// GameConfig holds game-specific configuration
type GameConfig struct {
	LaunchURLTemplate string `mapstructure:"launch_url_template"`
	SessionTTL        int    `mapstructure:"session_ttl"`
	CacheTTL          int    `mapstructure:"cache_ttl"`
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// New creates a new Config instance
func New(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// Set defaults
	viper.SetDefault("app.port", 8004)
	viper.SetDefault("app.grpc_port", 9004)
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.ssl_mode", "disable")
	viper.SetDefault("database.max_open_conns", 25)
	viper.SetDefault("database.max_idle_conns", 5)
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.pool_size", 10)
	viper.SetDefault("game.session_ttl", 3600)
	viper.SetDefault("game.cache_ttl", 600)
	viper.SetDefault("logging.level", "debug")
	viper.SetDefault("logging.format", "json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}

// GetDatabaseDSN returns the database DSN string
func (c *DatabaseConfig) GetDatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode,
	)
}

// GetRedisAddr returns the Redis address
func (c *RedisConfig) GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// GetSessionTTL returns the session TTL as duration
func (c *GameConfig) GetSessionTTL() time.Duration {
	return time.Duration(c.SessionTTL) * time.Second
}

// GetCacheTTL returns the cache TTL as duration
func (c *GameConfig) GetCacheTTL() time.Duration {
	return time.Duration(c.CacheTTL) * time.Second
}
