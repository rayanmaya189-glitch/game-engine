package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	merchantpb "github.com/game_engine/gen/go/merchant/v1"
	"github.com/game_engine/merchant-service/internal/config"
	"github.com/game_engine/merchant-service/internal/handler"
	"github.com/game_engine/merchant-service/internal/repository"
	"github.com/game_engine/merchant-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := repository.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	redisClient, err := repository.NewRedisClient(cfg.Redis)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisClient.Close()

	repo := repository.NewMerchantRepository(db, redisClient)
	merchantService := service.NewMerchantService(repo, cfg)
	merchantHandler := handler.NewMerchantHandler(merchantService)

	grpcServer := grpc.NewServer()
	merchantpb.RegisterMerchantServiceServer(grpcServer, merchantHandler)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("merchant_service", grpc_health_v1.HealthCheckResponse_SERVING)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	go func() {
		log.Printf("Merchant Service starting on port %d", cfg.GRPC.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down Merchant Service...")
	grpcServer.GracefulStop()
	db.Close()
	redisClient.Close()
	log.Println("Merchant Service stopped")
}
