package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/game-engine/sports-betting-service/internal/config"
	"github.com/game-engine/sports-betting-service/internal/repository"
	"github.com/game-engine/sports-betting-service/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to PostgreSQL
	db, err := repository.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Connect to Redis
	redisClient, err := repository.NewRedisClient(cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	// Initialize repository
	repo := repository.NewSportsRepository(db, redisClient)

	// Initialize service
	sportsService := service.NewSportsService(repo, cfg)

	// Initialize handler - uncomment when proto is generated
	// sportsHandler := handler.NewSportsHandler(sportsService)
	_ = sportsService // Keep service to avoid unused error

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register sports betting service (in production, this would use generated protobuf code)
	// sportsv1.RegisterSportsBettingServiceServer(grpcServer, sportsHandler)

	// Register reflection for development
	reflection.Register(grpcServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting Sports Betting Service gRPC server on port %d", cfg.GRPC.Port)

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		log.Println("Shutting down Sports Betting Service...")
		grpcServer.GracefulStop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
