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
		AuthService       string `yaml:"auth_service"`
		UserService       string `yaml:"user_service"`
		WalletService     string `yaml:"wallet_service"`
		GameService       string `yaml:"game_service"`
		CommissionService string `yaml:"commission_service"`
		BonusService      string `yaml:"bonus_service"`
		TournamentService string `yaml:"tournament_service"`
		JackpotService    string `yaml:"jackpot_service"`
	} `yaml:"services"`

	Admin struct {
		AllowedIPs []string `yaml:"allowed_ips"`
	} `yaml:"admin"`
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
		JWTSecret:         cfg.JWT.Secret,
		JWTExpiration:     time.Duration(cfg.JWT.ExpirationHours) * time.Hour,
		RefreshExpiration: time.Duration(cfg.JWT.RefreshDays) * 24 * time.Hour,
		RedisClient:       redisClient,
		TokenBlacklistTTL: 24 * time.Hour,
	})

	rateLimiterMiddleware := middleware.NewRateLimiterMiddleware(&middleware.RateLimiterConfig{
		RedisClient:       redisClient,
		RequestsPerMinute: 1000,
		KeyPrefix:         "admin",
	})

	corsMiddleware := middleware.NewCORSMiddleware(&middleware.CORSConfig{
		AllowOrigins:     cfg.Admin.AllowedIPs,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	})

	validatorMiddleware := middleware.NewValidatorMiddleware(middleware.GetAdminValidationRules())
	errorHandler := handler.NewErrorHandler()

	authClient, _ := client.NewAuthClient(&client.AuthClientConfig{
		Address: cfg.Services.AuthService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	defer authClient.Close()

	userClient, _ := client.NewUserClient(&client.UserClientConfig{
		Address: cfg.Services.UserService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	defer userClient.Close()

	walletClient, _ := client.NewWalletClient(&client.WalletClientConfig{
		Address: cfg.Services.WalletService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	defer walletClient.Close()

	gameClient, _ := client.NewGameClient(&client.GameClientConfig{
		Address: cfg.Services.GameService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	defer gameClient.Close()

	commissionClient, _ := client.NewCommissionClient(&client.CommissionClientConfig{
		Address: cfg.Services.CommissionService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	defer commissionClient.Close()

	bonusClient, _ := client.NewBonusClient(&client.BonusClientConfig{
		Address: cfg.Services.BonusService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	defer bonusClient.Close()

	tournamentClient, _ := client.NewTournamentClient(&client.TournamentClientConfig{
		Address: cfg.Services.TournamentService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	defer tournamentClient.Close()

	jackpotClient, _ := client.NewJackpotClient(&client.JackpotClientConfig{
		Address: cfg.Services.JackpotService,
		Timeout: 5 * time.Second,
		UseTLS:  false,
	})
	defer jackpotClient.Close()

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
		AuthClient:            authClient,
		UserClient:            userClient,
		WalletClient:          walletClient,
		GameClient:            gameClient,
		CommissionClient:      commissionClient,
		BonusClient:           bonusClient,
		TournamentClient:      tournamentClient,
		JackpotClient:         jackpotClient,
		AllowedIPs:            cfg.Admin.AllowedIPs,
	})
	h.NoRoute(errorHandler.NotFoundHandler)
	h.NoMethod(errorHandler.MethodNotAllowedHandler)

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("Shutting down server...")
		h.Shutdown(context.Background())
	}()

	log.Printf("Admin Gateway starting on %s:%d", cfg.Server.Host, cfg.Server.Port)
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
			Port: 8081, Host: "0.0.0.0", ReadTimeout: 30, WriteTimeout: 30,
		},
		JWT: struct {
			Secret          string `yaml:"secret"`
			ExpirationHours int    `yaml:"expiration_hours"`
			RefreshDays     int    `yaml:"refresh_days"`
		}{
			Secret: getEnv("JWT_SECRET", ""), ExpirationHours: 24, RefreshDays: 7,
		},
		Redis: struct {
			Addr     string `yaml:"addr"`
			Password string `yaml:"password"`
			DB       int    `yaml:"db"`
		}{
			Addr: getEnv("REDIS_ADDR", "redis:6379"), Password: "", DB: 0,
		},
		Services: struct {
			AuthService       string `yaml:"auth_service"`
			UserService       string `yaml:"user_service"`
			WalletService     string `yaml:"wallet_service"`
			GameService       string `yaml:"game_service"`
			CommissionService string `yaml:"commission_service"`
			BonusService      string `yaml:"bonus_service"`
			TournamentService string `yaml:"tournament_service"`
			JackpotService    string `yaml:"jackpot_service"`
		}{
			AuthService: getEnv("AUTH_SERVICE_ADDR", "auth-service:50051"), UserService: getEnv("USER_SERVICE_ADDR", "user-service:50051"),
			WalletService: getEnv("WALLET_SERVICE_ADDR", "wallet-service:50051"), GameService: getEnv("GAME_SERVICE_ADDR", "game-registry:50051"),
			CommissionService: getEnv("COMMISSION_SERVICE_ADDR", "commission-service:50051"), BonusService: getEnv("BONUS_SERVICE_ADDR", "bonus-service:50051"),
			TournamentService: getEnv("TOURNAMENT_SERVICE_ADDR", "tournament-service:50051"), JackpotService: getEnv("JACKPOT_SERVICE_ADDR", "jackpot-service:50051"),
		},
		Admin: struct {
			AllowedIPs []string `yaml:"allowed_ips"`
		}{
			AllowedIPs: []string{"admin.example.com"},
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
