package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration for the card games service
type Config struct {
	Server     ServerConfig     `yaml:"server"`
	Redis      RedisConfig      `yaml:"redis"`
	NATS       NATSConfig       `yaml:"nats"`
	RNGService RNGServiceConfig `yaml:"rng_service"`
	Game       GameConfig       `yaml:"game"`
	Logging    LoggingConfig    `yaml:"logging"`
}

// ServerConfig holds gRPC and HTTP server configuration
type ServerConfig struct {
	GRPCPort int `yaml:"grpc_port"`
	HTTPPort int `yaml:"http_port"`
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// NATSConfig holds NATS configuration
type NATSConfig struct {
	URLs      []string `yaml:"urls"`
	ClusterID string   `yaml:"cluster_id"`
	ClientID  string   `yaml:"client_id"`
}

// RNGServiceConfig holds RNG service connection configuration
type RNGServiceConfig struct {
	Address string `yaml:"address"`
	Timeout string `yaml:"timeout"`
}

// GameConfig holds game-specific configuration
type GameConfig struct {
	DefaultDeckCount int              `yaml:"default_deck_count"`
	MaxDeckCount     int              `yaml:"max_deck_count"`
	Shuffle          ShuffleConfig    `yaml:"shuffle"`
	Blackjack        BlackjackConfig  `yaml:"blackjack"`
	Baccarat         BaccaratConfig   `yaml:"baccarat"`
	Poker            PokerConfig      `yaml:"poker"`
	AndarBahar       AndarBaharConfig `yaml:"andar_bahar"`
	TeenPatti        TeenPattiConfig  `yaml:"teen_patti"`
}

// ShuffleConfig holds shuffle configuration
type ShuffleConfig struct {
	ReshuffleThreshold int `yaml:"reshuffle_threshold"`
}

// BlackjackConfig holds blackjack-specific configuration
type BlackjackConfig struct {
	AllowSurrender       bool    `yaml:"allow_surrender"`
	AllowLateSurrender   bool    `yaml:"allow_late_surrender"`
	DealerStandsOnSoft17 bool    `yaml:"dealer_stands_on_soft_17"`
	MaxSplits            int     `yaml:"max_splits"`
	BlackjackPayout      float64 `yaml:"blackjack_payout"`
}

// BaccaratConfig holds baccarat-specific configuration
type BaccaratConfig struct {
	Commission float64 `yaml:"commission"`
	MaxPlayers int     `yaml:"max_players"`
	ShoeSize   int     `yaml:"shoe_size"`
}

// PokerConfig holds poker-specific configuration
type PokerConfig struct {
	MinPlayers      int    `yaml:"min_players"`
	MaxPlayers      int    `yaml:"max_players"`
	BlindStructure  string `yaml:"blind_structure"`
	TimeBankSeconds int    `yaml:"time_bank_seconds"`
}

// AndarBaharConfig holds Andar Bahar-specific configuration
type AndarBaharConfig struct {
	SideToDealFirst   string `yaml:"side_to_deal_first"`
	MaxCommunityCards int    `yaml:"max_community_cards"`
}

// TeenPattiConfig holds Teen Patti-specific configuration
type TeenPattiConfig struct {
	Variant    string `yaml:"variant"`
	MinPlayers int    `yaml:"min_players"`
	MaxPlayers int    `yaml:"max_players"`
	ShowOrder  string `yaml:"show_order"`
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

// Load loads configuration from a YAML file
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// LoadFromEnv loads configuration with environment variable overrides
func LoadFromEnv(path string) (*Config, error) {
	cfg, err := Load(path)
	if err != nil {
		return nil, err
	}

	// Environment variable overrides
	if v := os.Getenv("REDIS_HOST"); v != "" {
		cfg.Redis.Host = v
	}
	if v := os.Getenv("REDIS_PORT"); v != "" {
		cfg.Redis.Port = 0
	}
	if v := os.Getenv("GRPC_PORT"); v != "" {
		cfg.Server.GRPCPort = 0
	}
	if v := os.Getenv("HTTP_PORT"); v != "" {
		cfg.Server.HTTPPort = 0
	}

	return cfg, nil
}
