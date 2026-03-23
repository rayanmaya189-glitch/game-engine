package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the game engine configuration
type Config struct {
	Server     ServerConfig     `yaml:"server"`
	RNG        RNGConfig        `yaml:"rng"`
	Games      GamesConfig      `yaml:"games"`
	GameEngine GameEngineConfig `yaml:"game_engine"`
	Database   DatabaseConfig   `yaml:"database"`
	Redis      RedisConfig      `yaml:"redis"`
	NATS       NATSConfig       `yaml:"nats"`
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

// RNGConfig represents RNG configuration
type RNGConfig struct {
	UseHardware bool `yaml:"use_hardware"`
	AwsNitro    bool `yaml:"aws_nitro"`
	SeedLength  int  `yaml:"seed_length"`
}

// GamesConfig represents games configuration
type GamesConfig struct {
	Directory        string  `yaml:"directory"`
	AutoDiscover     bool    `yaml:"auto_discover"`
	DefaultRTP       float64 `yaml:"default_rtp"`
	DefaultHouseEdge float64 `yaml:"default_house_edge"`
}

// GameEngineConfig represents game engine configuration
type GameEngineConfig struct {
	MaxConcurrentGames int  `yaml:"max_concurrent_games"`
	SessionTimeout     int  `yaml:"session_timeout"` // seconds
	EnableProvablyFair bool `yaml:"enable_provably_fair"`
	VerifyRTP          bool `yaml:"verify_rtp"`
}

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"ssl_mode"`
}

// RedisConfig represents Redis configuration
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// NATSConfig represents NATS configuration
type NATSConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Load loads configuration from a YAML file
func Load(path string) (*Config, error) {
	if path == "" {
		path = "configs/config.yaml"
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Set defaults
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	if cfg.Server.Host == "" {
		cfg.Server.Host = "0.0.0.0"
	}
	if cfg.RNG.SeedLength == 0 {
		cfg.RNG.SeedLength = 32
	}
	if cfg.GameEngine.MaxConcurrentGames == 0 {
		cfg.GameEngine.MaxConcurrentGames = 10000
	}
	if cfg.GameEngine.SessionTimeout == 0 {
		cfg.GameEngine.SessionTimeout = 3600
	}

	return &cfg, nil
}

// GetDSN returns the PostgreSQL connection string
func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Database,
		c.Database.SSLMode,
	)
}

// GetRedisAddr returns the Redis address
func (c *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port)
}

// GetNATSAddr returns the NATS address
func (c *Config) GetNATSAddr() string {
	return fmt.Sprintf("nats://%s:%s@%s:%d",
		c.NATS.Username,
		c.NATS.Password,
		c.NATS.Host,
		c.NATS.Port,
	)
}
