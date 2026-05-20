package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/game_engine/game-engine/internal/config"
	"github.com/game_engine/game-engine/internal/handler"
	"github.com/game_engine/game-engine/internal/registry"
	"github.com/game_engine/game-engine/internal/rng"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	rngService, err := rng.NewRNGService(cfg.RNG)
	if err != nil {
		log.Fatalf("Failed to initialize RNG: %v", err)
	}

	gameRegistry := registry.NewGameRegistry()

	gameEngineHandler := handler.NewGameEngineHandler(rngService, gameRegistry, cfg.GameEngine)

	grpcServer := grpc.NewServer()

	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthSrv)
	reflection.Register(grpcServer)

	gameEngineHandler.RegisterServices(grpcServer)

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", addr, err)
	}

	log.Printf("Game Engine gRPC server starting on %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}
