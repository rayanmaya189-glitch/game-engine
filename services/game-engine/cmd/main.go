package main

import (
	"fmt"
	"log"

	"github.com/game-engine/game-engine/internal/config"
	"github.com/game-engine/game-engine/internal/registry"
	"github.com/game-engine/game-engine/internal/rng"
	"github.com/game-engine/game-engine/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize RNG
	rngService, err := rng.NewRNGService(cfg.RNG)
	if err != nil {
		log.Fatalf("Failed to initialize RNG: %v", err)
	}

	// Initialize game registry
	gameRegistry := registry.NewGameRegistry(cfg.Games)

	// Initialize game engine service
	gameEngine := service.NewGameEngineService(rngService, gameRegistry, cfg.GameEngine)

	// Start server
	fmt.Printf("Game Engine Service starting on %s...\n", cfg.Server.Port)
	if err := gameEngine.Start(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
