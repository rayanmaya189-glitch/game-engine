package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration for the RNG service
type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Redis   RedisConfig   `yaml:"redis"`
	NATS    NATSConfig    `yaml:"nats"`
	RNG     RNGConfig     `yaml:"rng"`
	Logging LoggingConfig `yaml:"logging"`
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

// RNGConfig holds RNG-specific configuration
type RNGConfig struct {
	HardwareRNG    HardwareRNGConfig    `yaml:"hardware_rng"`
	SoftwareRNG    SoftwareRNGConfig    `yaml:"software_rng"`
	SeedGeneration SeedGenerationConfig `yaml:"seed_generation"`
	Certification  CertificationConfig  `yaml:"certification"`
	Limits         LimitsConfig         `yaml:"limits"`
}

// HardwareRNGConfig holds hardware RNG configuration
type HardwareRNGConfig struct {
	Enabled            bool `yaml:"enabled"`
	AWSNitro           bool `yaml:"aws_nitro"`
	FallbackToSoftware bool `yaml:"fallback_to_software"`
}

// SoftwareRNGConfig holds software RNG configuration
type SoftwareRNGConfig struct {
	SeedSource string `yaml:"seed_source"`
}

// SeedGenerationConfig holds seed generation configuration
type SeedGenerationConfig struct {
	ServerSeedLength int  `yaml:"server_seed_length"`
	ClientSeedLength int  `yaml:"client_seed_length"`
	NonceIncrement   bool `yaml:"nonce_increment"`
}

// CertificationConfig holds randomness certification configuration
type CertificationConfig struct {
	Enabled              bool          `yaml:"enabled"`
	ValidityPeriod       time.Duration `yaml:"validity_period"`
	VerificationInterval int64         `yaml:"verification_interval"`
}

// LimitsConfig holds RNG limits configuration
type LimitsConfig struct {
	MaxInt       int64   `yaml:"max_int"`
	MaxFloat     float64 `yaml:"max_float"`
	MaxDeckSize  int     `yaml:"max_deck_size"`
	MaxDiceCount int     `yaml:"max_dice_count"`
	MaxSlotReels int     `yaml:"max_slot_reels"`
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
	if v := os.Getenv("REDIS_PASSWORD"); v != "" {
		cfg.Redis.Password = v
	}

	if v := os.Getenv("GRPC_PORT"); v != "" {
		cfg.Server.GRPCPort = 0
	}
	if v := os.Getenv("HTTP_PORT"); v != "" {
		cfg.Server.HTTPPort = 0
	}

	return cfg, nil
}
