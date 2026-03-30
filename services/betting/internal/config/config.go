package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Betting BettingConfig `yaml:"betting"`
	Redis   RedisConfig   `yaml:"redis"`
	NATS    NATSConfig    `yaml:"nats"`
	Logging LoggingConfig `yaml:"logging"`
	DB      DBConfig      `yaml:"database"`
}

type ServerConfig struct {
	GRPCPort int `yaml:"grpc_port"`
	HTTPPort int `yaml:"http_port"`
}

type BettingConfig struct {
	MinBet      int64   `yaml:"min_bet"`
	MaxBet      int64   `yaml:"max_bet"`
	MaxPayout   int64   `yaml:"max_payout"`
	MaxOdds     float64 `yaml:"max_odds"`
	PlatformRake float64 `yaml:"platform_rake"`
	DefaultOdds string  `yaml:"default_odds"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type NATSConfig struct {
	URLs      []string `yaml:"urls"`
	ClusterID string   `yaml:"cluster_id"`
	ClientID  string   `yaml:"client_id"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslmode"`
}

func (d DBConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}

func Load() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/config.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("config file not found at %s: %w", configPath, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.Betting.MinBet == 0 {
		cfg.Betting.MinBet = 1
	}
	if cfg.Betting.MaxBet == 0 {
		cfg.Betting.MaxBet = 100000
	}
	if cfg.Betting.MaxPayout == 0 {
		cfg.Betting.MaxPayout = 1000000
	}
	if cfg.Betting.MaxOdds == 0 {
		cfg.Betting.MaxOdds = 1000.0
	}

	return &cfg, nil
}
