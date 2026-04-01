package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/game_engine/banner-service/internal/config"
	"github.com/game_engine/banner-service/internal/handler"
	"github.com/game_engine/banner-service/internal/repository"
	"github.com/game_engine/banner-service/internal/service"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := repository.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := repository.NewBannerRepository(db)
	bannerSvc := service.NewBannerService(repo)
	bannerHandler := handler.NewBannerHandler(bannerSvc)

	grpcServer := grpc.NewServer()

	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthSrv)
	reflection.Register(grpcServer)

	bannerHandler.RegisterServices(grpcServer)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.GRPCPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	go func() {
		log.Printf("Banner Service gRPC server starting on %s", addr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down Banner Service...")
	grpcServer.GracefulStop()

	_ = context.Background()

	log.Println("Banner Service stopped")
}
