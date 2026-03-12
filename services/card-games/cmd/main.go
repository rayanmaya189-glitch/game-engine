package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/game-engine/card-games/internal/config"
	"github.com/game-engine/card-games/internal/handler"
	"github.com/game-engine/card-games/internal/service"
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

	// Initialize game service
	gameService, err := service.NewGameService(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize game service: %v", err)
	}

	// Initialize handler
	gameHandler := handler.NewGameHandler(gameService)

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register game service (in production, this would use generated protobuf code)
	_ = gameHandler

	// For now, just log that the handler is ready
	log.Printf("Card Games Service initialized successfully")

	// Register reflection for development
	reflection.Register(grpcServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting Card Games gRPC server on port %d", cfg.Server.GRPCPort)

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		log.Println("Shutting down Card Games server...")
		grpcServer.GracefulStop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
