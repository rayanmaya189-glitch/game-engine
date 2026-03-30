package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration for the auth service
type Config struct {
	Server       ServerConfig       `yaml:"server"`
	Database     DatabaseConfig     `yaml:"database"`
	Redis        RedisConfig        `yaml:"redis"`
	NATS         NATSConfig         `yaml:"nats"`
	JWT          JWTConfig          `yaml:"jwt"`
	Password     PasswordConfig     `yaml:"password"`
	RateLimiting RateLimitingConfig `yaml:"rate_limiting"`
	TokenExpiry  TokenExpiryConfig  `yaml:"token_expiry"`
	TOTP         TOTPConfig         `yaml:"totp"`
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

// JWTConfig holds JWT configuration
type JWTConfig struct {
	PrivateKeyPath     string        `yaml:"private_key_path"`
	AccessTokenExpiry  time.Duration `yaml:"access_token_expiry"`
	RefreshTokenExpiry time.Duration `yaml:"refresh_token_expiry"`
	Issuer             string        `yaml:"issuer"`
	Audience           string        `yaml:"audience"`
}

// PasswordConfig holds password validation configuration
type PasswordConfig struct {
	MinLength        int  `yaml:"min_length"`
	RequireUppercase bool `yaml:"require_uppercase"`
	RequireLowercase bool `yaml:"require_lowercase"`
	RequireNumber    bool `yaml:"require_number"`
	RequireSpecial   bool `yaml:"require_special"`
	BcryptCost       int  `yaml:"bcrypt_cost"`
}

// RateLimitingConfig holds rate limiting configuration
type RateLimitingConfig struct {
	MaxLoginAttempts int           `yaml:"max_login_attempts"`
	WindowDuration   time.Duration `yaml:"window_duration"`
	LockoutDuration  time.Duration `yaml:"lockout_duration"`
}

// TokenExpiryConfig holds token expiry configuration
type TokenExpiryConfig struct {
	EmailVerification time.Duration `yaml:"email_verification"`
	PasswordReset     time.Duration `yaml:"password_reset"`
	TOTP              time.Duration `yaml:"totp"`
}

// TOTPConfig holds TOTP configuration
type TOTPConfig struct {
	Issuer string `yaml:"issuer"`
	Digits int    `yaml:"digits"`
	Period int    `yaml:"period"`
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
