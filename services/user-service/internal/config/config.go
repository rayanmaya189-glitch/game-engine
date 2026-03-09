package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration for the user service
type Config struct {
	Server            ServerConfig            `yaml:"server"`
	Database          DatabaseConfig          `yaml:"database"`
	Redis             RedisConfig             `yaml:"redis"`
	NATS              NATSConfig              `yaml:"nats"`
	Cache             CacheConfig             `yaml:"cache"`
	RateLimiting      RateLimitingConfig      `yaml:"rate_limiting"`
	KYC               KYCConfig               `yaml:"kyc"`
	ResponsibleGaming ResponsibleGamingConfig `yaml:"responsible_gaming"`
}

// ServerConfig holds gRPC and HTTP server configuration
type ServerConfig struct {
	GRPCPort int `yaml:"grpc_port"`
	HTTPPort int `yaml:"http_port"`
}

// DatabaseConfig holds PostgreSQL configuration
type DatabaseConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	DBName       string `yaml:"name"`
	SSLMode      string `yaml:"ssl_mode"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
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

// CacheConfig holds cache configuration
type CacheConfig struct {
	ProfileTTL   int `yaml:"profile_ttl"`
	KYCStatusTTL int `yaml:"kyc_status_ttl"`
}

// RateLimitingConfig holds rate limiting configuration
type RateLimitingConfig struct {
	ProfileUpdateWindow time.Duration `yaml:"profile_update_window"`
	ProfileUpdateMax    int           `yaml:"profile_update_max"`
}

// KYCConfig holds KYC configuration
type KYCConfig struct {
	ProviderURL      string `yaml:"provider_url"`
	ProviderAPIKey   string `yaml:"provider_api_key"`
	AutoApproveBasic bool   `yaml:"auto_approve_basic"`
}

// ResponsibleGamingConfig holds responsible gaming configuration
type ResponsibleGamingConfig struct {
	DefaultDailyLimit    int `yaml:"default_daily_limit"`
	DefaultWeeklyLimit   int `yaml:"default_weekly_limit"`
	DefaultMonthlyLimit  int `yaml:"default_monthly_limit"`
	SelfExclusionMinDays int `yaml:"self_exclusion_min_days"`
	SelfExclusionMaxDays int `yaml:"self_exclusion_max_days"`
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
	if v := os.Getenv("DB_HOST"); v != "" {
		cfg.Database.Host = v
	}
	if v := os.Getenv("DB_PORT"); v != "" {
		cfg.Database.Port = 0
	}
	if v := os.Getenv("DB_USER"); v != "" {
		cfg.Database.Username = v
	}
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		cfg.Database.Password = v
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		cfg.Database.DBName = v
	}

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
