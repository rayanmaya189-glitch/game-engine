package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/game_engine/winners-showcase-service/internal/config"
	"github.com/game_engine/winners-showcase-service/internal/handler"
	"github.com/game_engine/winners-showcase-service/internal/repository"
	"github.com/game_engine/winners-showcase-service/internal/service"
	"github.com/gin-gonic/gin"
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

	// Setup router
	router := setupRouter(winnersHandler)

	// Create server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Winners Showcase service starting on port %d", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

func setupRouter(h *handler.WinnersHandler) *gin.Engine {
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Recent winners feed
		api.GET("/winners/recent", h.GetRecentWinners)

		// Big win highlights
		api.GET("/winners/big", h.GetBigWins)

		// Jackpot winners
		api.GET("/winners/jackpot", h.GetJackpotWinners)

		// Record a new win (internal)
		api.POST("/winners/record", h.RecordWin)

		// Get player's privacy settings
		api.GET("/winners/privacy/:userId", h.GetPrivacySettings)

		// Update player's privacy settings
		api.PUT("/winners/privacy/:userId", h.UpdatePrivacySettings)
	}

	return router
}
