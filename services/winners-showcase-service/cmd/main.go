package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	winnersv1 "github.com/game_engine/common-service/proto/gen/go/winners/v1"
	"github.com/game_engine/winners-showcase-service/internal/config"
	"github.com/game_engine/winners-showcase-service/internal/handler"
	"github.com/game_engine/winners-showcase-service/internal/repository"
	"github.com/game_engine/winners-showcase-service/internal/service"
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

	// Initialize database connection
	db, err := repository.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Redis connection
	redisClient, err := repository.NewRedisClient(cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	// Initialize repository
	repo := repository.NewWinnersRepository(db, redisClient)

	// Initialize service
	winnersService := service.NewWinnersService(repo, &cfg.Winners)

	// Initialize handler
	winnersHandler := handler.NewWinnersHandler(winnersService)

	// Start gRPC server
	grpcAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	grpcServer := grpc.NewServer()

	// Register winners service
	winnersv1.RegisterWinnersServiceServer(grpcServer, winnersHandler)

	// Register gRPC health check
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("winners-showcase", grpc_health_v1.HealthCheckResponse_SERVING)

	reflection.Register(grpcServer)

	go func() {
		log.Printf("Winners Showcase gRPC server starting on %s", grpcAddr)
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
	log.Println("Shutting down gRPC server...")

	grpcServer.GracefulStop()

	log.Println("Server exited")
}
