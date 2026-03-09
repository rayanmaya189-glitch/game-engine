package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gameengine/rng-service/internal/config"
	"github.com/gameengine/rng-service/internal/handler"
	"github.com/gameengine/rng-service/internal/service"
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

	// Initialize RNG service
	rngService, err := service.NewRNGService(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize RNG service: %v", err)
	}

	// Initialize handler
	rngHandler := handler.NewRNGHandler(rngService)

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register RNG service (in production, this would use generated protobuf code)
	// rngev1.RegisterRNGServiceServer(grpcServer, rngHandler)

	// For now, just log that the handler is ready
	log.Printf("RNG Service initialized successfully")

	// Register reflection for development
	reflection.Register(grpcServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting RNG gRPC server on port %d", cfg.Server.GRPCPort)

	// Health check loop
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if err := rngService.HealthCheck(); err != nil {
				log.Printf("Health check failed: %v", err)
			}
		}
	}()

	// Seed cleanup loop
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			cleaned := rngService.Cleanup()
			if cleaned > 0 {
				log.Printf("Cleaned up %d old seed pairs", cleaned)
			}
		}
	}()

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		log.Println("Shutting down RNG server...")
		grpcServer.GracefulStop()
	}()

	// Create a simple server for demonstration
	// In production, this would register the protobuf-generated gRPC service
	type RNGServer interface {
		Ping(ctx interface{}) (interface{}, error)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
