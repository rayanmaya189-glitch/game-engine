package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/redis/go-redis/v9"

	"github.com/game_engine/gateway/common/client"
	"github.com/game_engine/gateway/common/handler"
	"github.com/game_engine/gateway/common/middleware"
)

type Config struct {
	Server struct {
		Port         int    `yaml:"port"`
		Host         string `yaml:"host"`
		ReadTimeout  int    `yaml:"read_timeout"`
		WriteTimeout int    `yaml:"write_timeout"`
	} `yaml:"server"`

	JWT struct {
		Secret          string `yaml:"secret"`
		ExpirationHours int    `yaml:"expiration_hours"`
		RefreshDays     int    `yaml:"refresh_days"`
	} `yaml:"jwt"`

	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	Services struct {
		AuthService   string `yaml:"auth_service"`
		UserService   string `yaml:"user_service"`
		WalletService string `yaml:"wallet_service"`
		GameService   string `yaml:"game_service"`
	} `yaml:"services"`
}

func main() {
	// Load configuration
	cfg := loadConfig()

	// Initialize Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Printf("Warning: Redis connection failed: %v", err)
	}

	// Initialize middleware
	loggerMiddleware := middleware.NewLoggerMiddleware(&middleware.LoggerConfig{
		OutputPath: "",
		LogLevel:   "info",
		Format:     "json",
	})

	authMiddleware := middleware.NewAuthMiddleware(&middleware.AuthConfig{
		JWTSecret:         cfg.JWT.Secret,
		JWTExpiration:     time.Duration(cfg.JWT.ExpirationHours) * time.Hour,
		RefreshExpiration: time.Duration(cfg.JWT.RefreshDays) * 24 * time.Hour,
		RedisClient:       redisClient,
		TokenBlacklistTTL: 24 * time.Hour,
	})

	rateLimiterMiddleware := middleware.NewRateLimiterMiddleware(&middleware.RateLimiterConfig{
		RedisClient:       redisClient,
		RequestsPerMinute: 100,
		KeyPrefix:         "player",
	})

	corsMiddleware := middleware.NewCORSMiddleware(&middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: false,
	})

	validatorMiddleware := middleware.NewValidatorMiddleware(middleware.GetPlayerValidationRules())

	errorHandler := handler.NewErrorHandler()

	// Initialize gRPC clients
	authClient, err := client.NewAuthClient(&client.AuthClientConfig{
		Address: cfg.Services.AuthService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	if err != nil {
		log.Printf("Warning: Auth service connection failed: %v", err)
	}
	defer authClient.Close()

	userClient, err := client.NewUserClient(&client.UserClientConfig{
		Address: cfg.Services.UserService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	if err != nil {
		log.Printf("Warning: User service connection failed: %v", err)
	}
	defer userClient.Close()

	walletClient, err := client.NewWalletClient(&client.WalletClientConfig{
		Address: cfg.Services.WalletService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	if err != nil {
		log.Printf("Warning: Wallet service connection failed: %v", err)
	}
	defer walletClient.Close()

	gameClient, err := client.NewGameClient(&client.GameClientConfig{
		Address: cfg.Services.GameService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	if err != nil {
		log.Printf("Warning: Game service connection failed: %v", err)
	}
	defer gameClient.Close()

	// Create Hertz server
	h := server.Default(
		server.WithHostPorts(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)),
		server.WithReadTimeout(time.Duration(cfg.Server.ReadTimeout)*time.Second),
		server.WithWriteTimeout(time.Duration(cfg.Server.WriteTimeout)*time.Second),
	)

	// Setup routes
	router := NewRouter(&RouterConfig{
		AuthMiddleware:        authMiddleware,
		LoggerMiddleware:      loggerMiddleware,
		RateLimiterMiddleware: rateLimiterMiddleware,
		CORSMiddleware:        corsMiddleware,
		ValidatorMiddleware:   validatorMiddleware,
		ErrorHandler:          errorHandler,
		AuthClient:            authClient,
		UserClient:            userClient,
		WalletClient:          walletClient,
		GameClient:            gameClient,
	})

	h.SetRouter(router)

	// Handle errors
	h.SetNotFoundHandler(errorHandler.NotFoundHandler)
	h.SetMethodNotAllowedHandler(errorHandler.MethodNotAllowedHandler)

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("Shutting down server...")
		h.Shutdown()
	}()

	log.Printf("Player Gateway starting on %s:%d", cfg.Server.Host, cfg.Server.Port)
	if err := h.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func loadConfig() *Config {
	// Default configuration
	cfg := &Config{
		Server: struct {
			Port         int    `yaml:"port"`
			Host         string `yaml:"host"`
			ReadTimeout  int    `yaml:"read_timeout"`
			WriteTimeout int    `yaml:"write_timeout"`
		}{
			Port:         8080,
			Host:         "0.0.0.0",
			ReadTimeout:  30,
			WriteTimeout: 30,
		},
		JWT: struct {
			Secret          string `yaml:"secret"`
			ExpirationHours int    `yaml:"expiration_hours"`
			RefreshDays     int    `yaml:"refresh_days"`
		}{
			Secret:          getEnv("JWT_SECRET", ""),
			ExpirationHours: 24,
			RefreshDays:     7,
		},
		Redis: struct {
			Addr     string `yaml:"addr"`
			Password string `yaml:"password"`
			DB       int    `yaml:"db"`
		}{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: "",
			DB:       0,
		},
		Services: struct {
			AuthService   string `yaml:"auth_service"`
			UserService   string `yaml:"user_service"`
			WalletService string `yaml:"wallet_service"`
			GameService   string `yaml:"game_service"`
		}{
			AuthService:   getEnv("AUTH_SERVICE_ADDR", "auth-service:50051"),
			UserService:   getEnv("USER_SERVICE_ADDR", "user-service:50051"),
			WalletService: getEnv("WALLET_SERVICE_ADDR", "wallet-service:50051"),
			GameService:   getEnv("GAME_SERVICE_ADDR", "game-registry:50051"),
		},
	}

	// Try to load from config file
	// In production, use viper or similar library

	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
