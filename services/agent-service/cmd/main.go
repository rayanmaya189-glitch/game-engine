package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/game_engine/agent-service/internal/config"
	"github.com/game_engine/agent-service/internal/handler"
	"github.com/game_engine/agent-service/internal/repository"
	"github.com/game_engine/agent-service/internal/service"
	agentpb "github.com/game_engine/gen/go/game_engine/agent/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := repository.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Redis
	redisClient, err := repository.NewRedisClient(cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	// Initialize repository
	repo := repository.NewAgentRepository(db, redisClient)

	// Initialize service
	agentService := service.NewAgentService(repo, cfg)

	// Initialize handler
	agentHandler := handler.NewAgentHandler(agentService)

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register service
	agentpb.RegisterAgentServiceServer(grpcServer, agentHandler)

	// Register health check
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("agent_service", grpc_health_v1.HealthCheckResponse_SERVING)

	// Enable reflection for debugging
	reflection.Register(grpcServer)

	// Create listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start server in goroutine
	go func() {
		log.Printf("Agent Service starting on port %d", cfg.GRPC.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down Agent Service...")

	// Graceful shutdown
	grpcServer.GracefulStop()

	// Close connections
	db.Close()
	redisClient.Close()

	log.Println("Agent Service stopped")
}
