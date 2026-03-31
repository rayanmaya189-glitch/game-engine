package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GRPC     GRPCConfig     `yaml:"grpc"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	NATS     NATSConfig     `yaml:"nats"`
	Dealer   DealerConfig   `yaml:"dealer"`
}

type GRPCConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	Name            string `yaml:"name"`
	SSLMode         string `yaml:"sslmode"`
	MaxConnections  int    `yaml:"max_connections"`
}

func (d DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSLMode)
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func (r RedisConfig) Address() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

type NATSConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (n NATSConfig) Address() string {
	return fmt.Sprintf("nats://%s:%d", n.Host, n.Port)
}

type DealerConfig struct {
	MaxTablesPerDealer     int `yaml:"max_tables_per_dealer"`
	SessionTimeoutMinutes int `yaml:"session_timeout_minutes"`
	VideoStreamBitrate    int `yaml:"video_stream_bitrate"`
	MaxPlayersPerTable    int `yaml:"max_players_per_table"`
}

func Load() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/config.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("config file not found at %s: please provide CONFIG_PATH or config file", configPath)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.Database.SSLMode == "" {
		cfg.Database.SSLMode = "disable"
	}
	if cfg.Database.MaxConnections == 0 {
		cfg.Database.MaxConnections = 20
	}

	return &cfg, nil
}
