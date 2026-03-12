package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"

	"github.com/game-engine/tournament/internal/handler"
	"github.com/game-engine/tournament/internal/service"
	"github.com/game-engine/tournament/internal/tournament"
)

func main() {
	log.Printf("Tournament Service starting...")

	// Load configuration
	config := &tournament.Config{}
	configData, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		log.Printf("Warning: Could not read config file, using defaults: %v", err)
	} else {
		if err := yaml.Unmarshal(configData, config); err != nil {
			log.Fatalf("Failed to parse config: %v", err)
		}
	}

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		DB:   config.Redis.DB,
	})

	// Initialize tournament manager
	manager, err := tournament.NewManager(config, redisClient)
	if err != nil {
		log.Fatalf("Failed to create tournament manager: %v", err)
	}

	// Initialize service
	tournamentService, err := service.NewTournamentService(manager)
	if err != nil {
		log.Fatalf("Failed to create tournament service: %v", err)
	}

	// Initialize handler
	tournamentHandler := handler.NewTournamentHandler(tournamentService)
	_ = tournamentHandler // Use handler for HTTP/gRPC endpoints

	// Start gRPC server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Get port from config or use default
	port := config.Server.GRPCPort
	if port == 0 {
		port = 9020
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("Starting Tournament gRPC server on port %d\n", port)

	// Handle graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down Tournament server...")
		grpcServer.GracefulStop()
		redisClient.Close()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
