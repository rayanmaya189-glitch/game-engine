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

	walletsv1 "game_engine/gen/go/wallet/v1"

	"github.com/game_engine/wallet-service/internal/config"
	"github.com/game_engine/wallet-service/internal/handler"
	"github.com/game_engine/wallet-service/internal/repository"
	"github.com/game_engine/wallet-service/internal/service"

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
	repo := repository.NewWalletRepository(db)

	// Initialize service
	walletService, err := service.NewWalletService(repo, redisClient, natsConn, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize wallet service: %v", err)
	}

	// Initialize handler
	walletHandler := handler.NewWalletHandler(walletService)

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register wallet service
	walletsv1.RegisterWalletServiceServer(grpcServer, walletHandler)

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

	// Start bonus expiry worker
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()
		for range ticker.C {
			ctx := context.Background()
			expired, err := repo.ExpireBonuses(ctx)
			if err != nil {
				log.Printf("Error expiring bonuses: %v", err)
			} else if expired > 0 {
				log.Printf("Expired %d bonuses", expired)
			}
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
