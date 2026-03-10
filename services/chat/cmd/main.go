package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"

	"github.com/gameengine/chat/internal/handler"
	"github.com/gameengine/chat/internal/room"
	"github.com/gameengine/chat/internal/service"
)

func main() {
	log.Printf("Chat Service starting...")

	// Load configuration
	config := &room.Config{}
	configData, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		log.Printf("Warning: Could not read config file, using defaults: %v", err)
	} else {
		if err := yaml.Unmarshal(configData, config); err != nil {
			log.Fatalf("Failed to parse config: %v", err)
		}
	}

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		DB:   config.Redis.DB,
	})

	// Initialize room manager
	manager, err := room.NewManager(config, redisClient)
	if err != nil {
		log.Fatalf("Failed to create room manager: %v", err)
	}

	// Initialize service
	chatService, err := service.NewChatService(manager)
	if err != nil {
		log.Fatalf("Failed to create chat service: %v", err)
	}

	// Initialize handler
	chatHandler := handler.NewChatHandler(chatService)
	_ = chatHandler

	// Start gRPC server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Get port from config or use default
	port := config.Server.GRPCPort
	if port == 0 {
		port = 9022
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Printf("Starting Chat gRPC server on port %d\n", port)

	// Handle graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down Chat server...")
		grpcServer.GracefulStop()
		redisClient.Close()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
