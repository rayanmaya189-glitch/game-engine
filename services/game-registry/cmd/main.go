package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gamesv1 "github.com/game_engine/gen/go/game_engine/game/v1"

	"github.com/game_engine/game-registry/internal/config"
	"github.com/game_engine/game-registry/internal/handler"
	"github.com/game_engine/game-registry/internal/repository"
	"github.com/game_engine/game-registry/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load configuration
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./configs/config.yaml"
	}

	cfg, err := config.New(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := initDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize Redis
	rdb, err := initRedis(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	defer rdb.Close()

	// Initialize NATS
	nc, err := initNATS(cfg)
	if err != nil {
		log.Printf("Warning: Failed to initialize NATS: %v", err)
		// Continue without NATS
	}
	if nc != nil {
		defer nc.Close()
	}

	// Initialize repository
	gameRepo := repository.NewGameRepository(db, rdb)

	// Initialize service
	gameService := service.NewGameService(gameRepo, cfg, nc)

	// Initialize handler
	gameHandler := handler.NewGameHandler(gameService)

	// Setup router
	router := setupRouter(gameHandler)

	// Start HTTP server
	httpAddr := fmt.Sprintf(":%d", cfg.App.Port)
	httpServer := &http.Server{
		Addr:    httpAddr,
		Handler: router,
	}

	// Start gRPC server
	grpcAddr := fmt.Sprintf(":%d", cfg.App.GRPCPort)
	grpcServer := grpc.NewServer()

	// Register gRPC service
	gamesv1.RegisterGameRegistryServiceServer(grpcServer, gameHandler)

	go func() {
		log.Printf("Starting HTTP server on %s", httpAddr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	go func() {
		log.Printf("Starting gRPC server on %s", grpcAddr)
		reflection.Register(grpcServer)
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("HTTP server forced to shutdown: %v", err)
	}

	grpcServer.GracefulStop()

	log.Println("Servers exited")
}

func initDatabase(cfg *config.Config) (*sql.DB, error) {
	dsn := cfg.Database.GetDatabaseDSN()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}

func initRedis(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.GetRedisAddr(),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
		PoolSize: cfg.Redis.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Println("Redis connection established")
	return rdb, nil
}

func initNATS(cfg *config.Config) (*nats.Conn, error) {
	nc, err := nats.Connect(cfg.NATS.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	log.Println("NATS connection established")
	return nc, nil
}

func setupRouter(gameHandler *handler.GameHandler) *gin.Engine {
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API v1
	v1 := router.Group("/api/v1")
	{
		// Game endpoints
		games := v1.Group("/games")
		{
			games.GET("", gameHandler.ListGames)
			games.GET("/search", gameHandler.SearchGames)
			games.GET("/featured", gameHandler.GetFeaturedGames)
			games.GET("/popular", gameHandler.GetPopularGames)
			games.GET("/new", gameHandler.GetNewGames)
			games.GET("/:id", gameHandler.GetGame)
			games.GET("/:id/config", gameHandler.GetGameConfig)
			games.POST("/:id/url", gameHandler.GetGameURL)

			// Admin endpoints
			games.POST("", gameHandler.CreateGame)
			games.PUT("/:id", gameHandler.UpdateGame)
			games.POST("/:id/toggle", gameHandler.ToggleGame)
			games.POST("/:id/order", gameHandler.SetGameOrder)
		}

		// Category endpoints
		v1.GET("/categories", gameHandler.GetCategories)

		// Provider endpoints
		v1.GET("/providers", gameHandler.GetProviders)
	}

	return router
}
