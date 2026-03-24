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

	"github.com/game_engine/leaderboard-service/internal/config"
	"github.com/game_engine/leaderboard-service/internal/handler"
	"github.com/game_engine/leaderboard-service/internal/repository"
	"github.com/game_engine/leaderboard-service/internal/service"
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
	repo := repository.NewLeaderboardRepository(db, redisClient)

	// Initialize service
	leaderboardService := service.NewLeaderboardService(repo, cfg.Leaderboard)

	// Initialize handler
	leaderboardHandler := handler.NewLeaderboardHandler(leaderboardService)

	// Setup router
	router := setupRouter(leaderboardHandler)

	// Create server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Leaderboard service starting on port %d", cfg.Server.Port)
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

func setupRouter(h *handler.LeaderboardHandler) *gin.Engine {
	router := gin.Default()

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// API routes
	api := router.Group("/api/v1")
	{
		// Daily leaderboard
		api.GET("/leaderboard/daily", h.GetDailyLeaderboard)
		api.GET("/leaderboard/daily/:gameType", h.GetDailyLeaderboardByGame)

		// Weekly leaderboard
		api.GET("/leaderboard/weekly", h.GetWeeklyLeaderboard)
		api.GET("/leaderboard/weekly/:gameType", h.GetWeeklyLeaderboardByGame)

		// Monthly leaderboard
		api.GET("/leaderboard/monthly", h.GetMonthlyLeaderboard)
		api.GET("/leaderboard/monthly/:gameType", h.GetMonthlyLeaderboardByGame)

		// All-time leaderboard
		api.GET("/leaderboard/alltime", h.GetAllTimeLeaderboard)
		api.GET("/leaderboard/alltime/:gameType", h.GetAllTimeLeaderboardByGame)

		// Player rank
		api.GET("/leaderboard/rank/daily/:userId", h.GetPlayerDailyRank)
		api.GET("/leaderboard/rank/weekly/:userId", h.GetPlayerWeeklyRank)
		api.GET("/leaderboard/rank/monthly/:userId", h.GetPlayerMonthlyRank)
		api.GET("/leaderboard/rank/alltime/:userId", h.GetPlayerAllTimeRank)

		// Update score (internal use)
		api.POST("/leaderboard/update", h.UpdatePlayerScore)
	}

	return router
}
