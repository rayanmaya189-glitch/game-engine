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
	} `yaml:"jwt"`

	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	Services struct {
		UserService string `yaml:"user_service"`
	} `yaml:"services"`
}

func main() {
	cfg := loadConfig()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Printf("Warning: Redis connection failed: %v", err)
	}

	loggerMiddleware := middleware.NewLoggerMiddleware(&middleware.LoggerConfig{
		OutputPath: "",
		LogLevel:   "info",
		Format:     "json",
	})

	authMiddleware := middleware.NewAuthMiddleware(&middleware.AuthConfig{
		JWTSecret:     cfg.JWT.Secret,
		JWTExpiration: time.Duration(cfg.JWT.ExpirationHours) * time.Hour,
		RedisClient:   redisClient,
	})

	rateLimiterMiddleware := middleware.NewRateLimiterMiddleware(&middleware.RateLimiterConfig{
		RedisClient:       redisClient,
		RequestsPerMinute: 300,
		KeyPrefix:         "agent",
	})

	corsMiddleware := middleware.NewCORSMiddleware(&middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-API-Key"},
		AllowCredentials: false,
	})

	validatorMiddleware := middleware.NewValidatorMiddleware(middleware.GetAgentValidationRules())
	errorHandler := handler.NewErrorHandler()

	userClient, _ := client.NewUserClient(&client.UserClientConfig{
		Address: cfg.Services.UserService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	defer userClient.Close()

	h := server.Default(
		server.WithHostPorts(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)),
		server.WithReadTimeout(time.Duration(cfg.Server.ReadTimeout)*time.Second),
		server.WithWriteTimeout(time.Duration(cfg.Server.WriteTimeout)*time.Second),
	)

	SetupRoutes(h.Engine, &RouterConfig{
		AuthMiddleware:        authMiddleware,
		LoggerMiddleware:      loggerMiddleware,
		RateLimiterMiddleware: rateLimiterMiddleware,
		CORSMiddleware:        corsMiddleware,
		ValidatorMiddleware:   validatorMiddleware,
		ErrorHandler:          errorHandler,
		UserClient:            userClient,
	})

	h.NoRoute(errorHandler.NotFoundHandler)

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("Shutting down server...")
		h.Shutdown(context.Background())
	}()

	log.Printf("Agent Gateway starting on %s:%d", cfg.Server.Host, cfg.Server.Port)
	if err := h.Run(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func loadConfig() *Config {
	return &Config{
		Server: struct {
			Port         int    `yaml:"port"`
			Host         string `yaml:"host"`
			ReadTimeout  int    `yaml:"read_timeout"`
			WriteTimeout int    `yaml:"write_timeout"`
		}{
			Port: 8083, Host: "0.0.0.0", ReadTimeout: 30, WriteTimeout: 30,
		},
		JWT: struct {
			Secret          string `yaml:"secret"`
			ExpirationHours int    `yaml:"expiration_hours"`
		}{
			Secret: getEnv("JWT_SECRET", ""), ExpirationHours: 24,
		},
		Redis: struct {
			Addr     string `yaml:"addr"`
			Password string `yaml:"password"`
			DB       int    `yaml:"db"`
		}{
			Addr: getEnv("REDIS_ADDR", "redis:6379"), Password: "", DB: 0,
		},
		Services: struct {
			UserService string `yaml:"user_service"`
		}{
			UserService: getEnv("USER_SERVICE_ADDR", "user-service:50051"),
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
