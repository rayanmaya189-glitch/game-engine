package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/game-engine/loyalty-service/internal/config"
	"github.com/game-engine/loyalty-service/internal/handler"
	"github.com/game-engine/loyalty-service/internal/repository"
	"github.com/game-engine/loyalty-service/internal/service"
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

	repo := repository.NewLoyaltyRepository(db, redisClient)
	loyaltyService := service.NewLoyaltyService(repo, cfg)
	loyaltyHandler := handler.NewLoyaltyHandler(loyaltyService)

	// Set up HTTP server with routes
	mux := http.NewServeMux()
	setupRoutes(mux, loyaltyHandler)

	server := &http.Server{
		Addr:         ":8081",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Use HTTP port from config if available, otherwise default
	if cfg.HTTP.Port > 0 {
		server.Addr = ":" + strconv.Itoa(cfg.HTTP.Port)
	}

	go func() {
		log.Printf("Loyalty Service starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down Loyalty Service...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	db.Close()
	redisClient.Close()
	log.Println("Loyalty Service stopped")
}

func setupRoutes(mux *http.ServeMux, h *handler.LoyaltyHandler) {
	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Loyalty endpoints
	mux.HandleFunc("/api/v1/loyalty/member/", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "user_id required", http.StatusBadRequest)
			return
		}
		ctx := context.Background()
		resp, err := h.GetMember(ctx, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("/api/v1/loyalty/points/history", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "user_id required", http.StatusBadRequest)
			return
		}
		ctx := context.Background()
		resp, err := h.GetPointsHistory(ctx, userID, 50, 0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("/api/v1/loyalty/tiers", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		resp, err := h.GetTiers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("/api/v1/loyalty/rewards", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		resp, err := h.GetRewards(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("/api/v1/loyalty/leaderboard", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		resp, err := h.GetLeaderboard(ctx, 100)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})
}
