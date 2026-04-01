package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Referral ReferralConfig `yaml:"referral"`
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

type ReferralConfig struct {
	CodePrefix          string  `yaml:"code_prefix"`
	CodeLength          int     `yaml:"code_length"`
	DefaultReferrerBonus float64 `yaml:"default_referrer_bonus"`
	DefaultRefereeBonus  float64 `yaml:"default_referee_bonus"`
	MinDepositForReward float64 `yaml:"min_deposit_for_reward"`
	MaxReferralDepth    int     `yaml:"max_referral_depth"`
	MultiTierRate       float64 `yaml:"multi_tier_rate"`
	RewardExpiryDays    int     `yaml:"reward_expiry_days"`
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
		Server: ServerConfig{Host: "0.0.0.0", GRPCPort: 9098},
		Database: DatabaseConfig{
			Host: "localhost", Port: 5432, User: "postgres",
			Password: "postgres", Database: "referral_service", MaxConns: 20,
		},
		Redis: RedisConfig{Host: "localhost", Port: 6379, DB: 0},
		Referral: ReferralConfig{
			CodePrefix: "REF", CodeLength: 8,
			DefaultReferrerBonus: 10.0, DefaultRefereeBonus: 5.0,
			MinDepositForReward: 20.0, MaxReferralDepth: 3,
			MultiTierRate: 0.1, RewardExpiryDays: 90,
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
