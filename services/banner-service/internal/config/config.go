package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Banner   BannerConfig   `yaml:"banner"`
}

type ServerConfig struct {
	Host     string `yaml:"host"`
	GRPCPort int    `yaml:"grpc_port"`
}

type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Database     string `yaml:"name"`
	MaxConns     int    `yaml:"max_connections"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   int    `yaml:"db"`
}

type BannerConfig struct {
	DefaultPageSize  int  `yaml:"default_page_size"`
	MaxPageSize      int  `yaml:"max_page_size"`
	CacheTTLSec      int  `yaml:"cache_ttl_seconds"`
	MaxImageSizeMB   int  `yaml:"max_image_size_mb"`
	AutoRotate       bool `yaml:"auto_rotate"`
	RotationInterval int  `yaml:"rotation_interval_seconds"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return defaultConfig(), nil
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	applyEnvOverrides(&cfg)

	return &cfg, nil
}

func defaultConfig() *Config {
	return &Config{
		Server: ServerConfig{Host: "0.0.0.0", GRPCPort: 9097},
		Database: DatabaseConfig{
			Host: "localhost", Port: 5432, User: "postgres",
			Password: "postgres", Database: "banner_service", MaxConns: 20,
		},
		Redis: RedisConfig{Host: "localhost", Port: 6379, DB: 0},
		Banner: BannerConfig{
			DefaultPageSize: 20, MaxPageSize: 100, CacheTTLSec: 300,
			MaxImageSizeMB: 5, AutoRotate: true, RotationInterval: 30,
		},
	}
}

func applyEnvOverrides(cfg *Config) {
	if v := os.Getenv("DB_HOST"); v != "" {
		cfg.Database.Host = v
	}
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		cfg.Database.Password = v
	}
	if v := os.Getenv("REDIS_HOST"); v != "" {
		cfg.Redis.Host = v
	}
}
