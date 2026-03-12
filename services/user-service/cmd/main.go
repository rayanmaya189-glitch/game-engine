package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	userv1 "game-engine/gen/go/user/v1"

	"github.com/game-engine/user-service/internal/config"
	"github.com/game-engine/user-service/internal/handler"
	"github.com/game-engine/user-service/internal/repository"
	"github.com/game-engine/user-service/internal/service"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load configuration
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/config.yaml"
	}

	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Configure connection pool
	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	defer redisClient.Close()

	// Connect to NATS
	var natsConn *nats.Conn
	if len(cfg.NATS.URLs) > 0 {
		natsConn, err = nats.Connect(cfg.NATS.URLs[0], nats.Name(cfg.NATS.ClientID))
		if err != nil {
			log.Printf("Warning: Failed to connect to NATS: %v", err)
		} else {
			defer natsConn.Close()
		}
	}

	// Initialize repository
	repo := repository.NewUserRepository(db)

	// Initialize service
	userService, err := service.NewUserService(repo, redisClient, natsConn, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize user service: %v", err)
	}

	// Initialize handler
	userHandler := handler.NewUserHandler(userService)

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register user service
	userv1.RegisterUserServiceServer(grpcServer, userHandler)

	// Register reflection for development
	reflection.Register(grpcServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting gRPC server on port %d", cfg.Server.GRPCPort)

	// Health check server for HTTP gateway
	go func() {
		// In production, this would start an HTTP gateway using grpc-gateway
		// For now, we'll just keep it simple
		log.Printf("HTTP gateway would run on port %d", cfg.Server.HTTPPort)
	}()

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		log.Println("Shutting down server...")
		grpcServer.GracefulStop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// Context with timeout helper
func withTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, timeout)
}
