package registry

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ListByType returns game definitions filtered by type
func (r *GameRegistry) ListByType(gameType string) []*GameDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*GameDefinition
	for _, game := range r.games {
		if game.Type == gameType {
			result = append(result, game)
		}
	}

	return result
}

// GetActiveGames returns all active games
func (r *GameRegistry) GetActiveGames() []*GameDefinition {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*GameDefinition
	for _, game := range r.games {
		if game.Status == "active" {
			result = append(result, game)
		}
	}

	return result
}

// GRPCGameRegistryClient implements GameRegistryClient using gRPC
type GRPCGameRegistryClient struct {
	conn   *grpc.ClientConn
	target string
}

// NewGRPCGameRegistryClient creates a new gRPC client for the game registry service
func NewGRPCGameRegistryClient(target string) (*GRPCGameRegistryClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to game registry at %s: %w", target, err)
	}

	return &GRPCGameRegistryClient{
		conn:   conn,
		target: target,
	}, nil
}

// Close closes the gRPC connection
func (c *GRPCGameRegistryClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// ListGames fetches all games from the game-registry service via gRPC
func (c *GRPCGameRegistryClient) ListGames(ctx context.Context) ([]*GameDefinition, error) {
	// The actual gRPC call would use generated stubs:
	//   client := gamesv1.NewGameRegistryServiceClient(c.conn)
	//   resp, err := client.ListGames(ctx, &gamesv1.ListGamesRequest{...})
	//
	// For now, return an error indicating the generated stubs are needed.
	// Once protoc generates the Go code, this method will call the real service.
	return nil, fmt.Errorf("gRPC stubs not yet generated; run protoc to generate game registry client stubs")
}

// GetGame fetches a single game from the game-registry service via gRPC
func (c *GRPCGameRegistryClient) GetGame(ctx context.Context, id string) (*GameDefinition, error) {
	return nil, fmt.Errorf("gRPC stubs not yet generated; run protoc to generate game registry client stubs")
}
