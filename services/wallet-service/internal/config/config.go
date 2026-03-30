package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration for the wallet service
type Config struct {
	Server       ServerConfig       `yaml:"server"`
	Database     DatabaseConfig     `yaml:"database"`
	Redis        RedisConfig        `yaml:"redis"`
	NATS         NATSConfig         `yaml:"nats"`
	Cache        CacheConfig        `yaml:"cache"`
	RateLimiting RateLimitingConfig `yaml:"rate_limiting"`
	Betting      BettingConfig      `yaml:"betting"`
	Withdrawal   WithdrawalConfig   `yaml:"withdrawal"`
	Bonus        BonusConfig        `yaml:"bonus"`
	Deposit      DepositConfig      `yaml:"deposit"`
	Payment      PaymentConfig      `yaml:"payment"`
}

// DepositConfig holds deposit configuration
type DepositConfig struct {
	MinDepositAmount int64 `yaml:"min_deposit_amount"`
	MaxDepositAmount int64 `yaml:"max_deposit_amount"`
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
	BalanceTTL     int `yaml:"balance_ttl"`
	TransactionTTL int `yaml:"transaction_ttl"`
	BetLockTTL     int `yaml:"bet_lock_ttl"`
	BonusTTL       int `yaml:"bonus_ttl"`
}

// RateLimitingConfig holds rate limiting configuration
type RateLimitingConfig struct {
	TransactionWindow time.Duration `yaml:"transaction_window"`
	TransactionMax    int           `yaml:"transaction_max"`
	BetWindow         time.Duration `yaml:"bet_window"`
	BetMax            int           `yaml:"bet_max"`
}

// BettingConfig holds betting configuration
type BettingConfig struct {
	MaxBetAmount    int64 `yaml:"max_bet_amount"`
	MinBetAmount    int64 `yaml:"min_bet_amount"`
	MaxWinAmount    int64 `yaml:"max_win_amount"`
	BetLockTimeout  int   `yaml:"bet_lock_timeout"`
	AutoSettleDelay int   `yaml:"auto_settle_delay"`
}

// WithdrawalConfig holds withdrawal configuration
type WithdrawalConfig struct {
	MinWithdrawalAmount    int64 `yaml:"min_withdrawal_amount"`
	MaxWithdrawalAmount    int64 `yaml:"max_withdrawal_amount"`
	DailyWithdrawalLimit   int64 `yaml:"daily_withdrawal_limit"`
	WeeklyWithdrawalLimit  int64 `yaml:"weekly_withdrawal_limit"`
	MonthlyWithdrawalLimit int64 `yaml:"monthly_withdrawal_limit"`
	ApprovalRequired       bool  `yaml:"approval_required"`
}

// BonusConfig holds bonus configuration
type BonusConfig struct {
	MaxBonusAmount            int64 `yaml:"max_bonus_amount"`
	MinDepositForBonus        int64 `yaml:"min_deposit_for_bonus"`
	DefaultWageringMultiplier int   `yaml:"default_wagering_multiplier"`
	BonusExpiryDays           int   `yaml:"bonus_expiry_days"`
}

// PaymentConfig holds payment gateway configuration
type PaymentConfig struct {
	ProviderURL       string `yaml:"provider_url"`
	ProviderAPIKey    string `yaml:"provider_api_key"`
	DepositTimeout    int    `yaml:"deposit_timeout"`
	WithdrawalTimeout int    `yaml:"withdrawal_timeout"`
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

	// Database environment overrides
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

	// Redis environment overrides
	if v := os.Getenv("REDIS_HOST"); v != "" {
		cfg.Redis.Host = v
	}
	if v := os.Getenv("REDIS_PORT"); v != "" {
		cfg.Redis.Port = 0
	}
	if v := os.Getenv("REDIS_PASSWORD"); v != "" {
		cfg.Redis.Password = v
	}

	// Server environment overrides
	if v := os.Getenv("GRPC_PORT"); v != "" {
		cfg.Server.GRPCPort = 0
	}
	if v := os.Getenv("HTTP_PORT"); v != "" {
		cfg.Server.HTTPPort = 0
	}

	// NATS environment overrides
	if v := os.Getenv("NATS_URL"); v != "" {
		cfg.NATS.URLs = []string{v}
	}

	return cfg, nil
}
