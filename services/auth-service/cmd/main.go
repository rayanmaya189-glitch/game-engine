package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	authv1 "github.com/game-engine/common/proto/gen/go/gameengine/auth/v1"

	"github.com/gameengine/auth-service/internal/config"
	"github.com/gameengine/auth-service/internal/handler"
	"github.com/gameengine/auth-service/internal/middleware"
	"github.com/gameengine/auth-service/internal/repository"
	"github.com/gameengine/auth-service/internal/service"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
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

	// Connect to PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Configure connection pool
	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)

	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	defer redisClient.Close()

	// Initialize repository
	repo := repository.NewAuthRepository(db)

	// Initialize service
	authService, err := service.NewAuthService(repo, redisClient, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize auth service: %v", err)
	}

	// Initialize handler
	authHandler := handler.NewAuthHandler(authService)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Create gRPC server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authMiddleware.AuthInterceptor()),
	)

	// Register auth service
	authv1.RegisterAuthServiceServer(grpcServer, authHandler)

	// Register reflection for development
	reflection.Register(grpcServer)

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting gRPC server on port %d", cfg.Server.GRPCPort)

	// Start HTTP gateway (in production, this would use grpc-gateway)
	// go func() {
	//  mux := runtime.NewServeMux()
	//  opts := []grpc.DialOption{grpc.WithInsecure()}
	//  err := gw.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", cfg.Server.GRPCPort), opts)
	//  if err != nil {
	//    log.Fatalf("Failed to start gateway: %v", err)
	//  }
	//  http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.HTTPPort), mux)
	// }()

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh

		log.Println("Shutting down server...")
		grpcServer.GracefulStop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
