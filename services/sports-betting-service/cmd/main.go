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

	"github.com/game-engine/sports-betting-service/internal/config"
	"github.com/game-engine/sports-betting-service/internal/handler"
	"github.com/game-engine/sports-betting-service/internal/repository"
	"github.com/game-engine/sports-betting-service/internal/service"
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

	repo := repository.NewSportsRepository(db, redisClient)
	sportsService := service.NewSportsService(repo, cfg)
	sportsHandler := handler.NewSportsHandler(sportsService)

	// Set up HTTP server with routes
	mux := http.NewServeMux()
	setupRoutes(mux, sportsHandler)

	server := &http.Server{
		Addr:         ":8082",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Use HTTP port from config if available, otherwise default
	if cfg.HTTP.Port > 0 {
		server.Addr = ":" + strconv.Itoa(cfg.HTTP.Port)
	}

	go func() {
		log.Printf("Sports Betting Service starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down Sports Betting Service...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	db.Close()
	redisClient.Close()
	log.Println("Sports Betting Service stopped")
}

func setupRoutes(mux *http.ServeMux, h *handler.SportsHandler) {
	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Sports endpoints
	mux.HandleFunc("/api/v1/sports", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		resp, err := h.GetSports(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("/api/v1/sports/live", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		resp, err := h.GetLiveEvents(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("/api/v1/sports/upcoming", func(w http.ResponseWriter, r *http.Request) {
		sportID := r.URL.Query().Get("sport_id")
		ctx := context.Background()
		resp, err := h.GetUpcomingEvents(ctx, sportID, 50)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("/api/v1/sports/markets", func(w http.ResponseWriter, r *http.Request) {
		eventID := r.URL.Query().Get("event_id")
		if eventID == "" {
			http.Error(w, "event_id required", http.StatusBadRequest)
			return
		}
		ctx := context.Background()
		resp, err := h.GetMarkets(ctx, eventID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("/api/v1/bets", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Place bet
			userID := r.URL.Query().Get("user_id")
			eventID := r.URL.Query().Get("event_id")
			marketID := r.URL.Query().Get("market_id")
			selection := r.URL.Query().Get("selection")
			stakeStr := r.URL.Query().Get("stake")
			oddsStr := r.URL.Query().Get("odds")

			stake, _ := strconv.ParseFloat(stakeStr, 64)
			odds, _ := strconv.ParseFloat(oddsStr, 64)

			ctx := context.Background()
			resp, err := h.PlaceBet(ctx, userID, eventID, marketID, selection, stake, odds)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(resp)
		} else {
			// Get user bets
			userID := r.URL.Query().Get("user_id")
			ctx := context.Background()
			resp, err := h.GetUserBets(ctx, userID, 1, 50)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(resp)
		}
	})
}
