package registry

import (
	"context"
	"fmt"
	"time"

	gamesv1 "game_engine/gen/go/game/v1"

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
	client gamesv1.GameRegistryServiceClient
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
		client: gamesv1.NewGameRegistryServiceClient(conn),
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
	resp, err := c.client.ListGames(ctx, &gamesv1.ListGamesRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to list games from registry: %w", err)
	}

	definitions := make([]*GameDefinition, 0, len(resp.Games))
	for _, g := range resp.Games {
		definitions = append(definitions, gameSummaryToDefinition(g))
	}
	return definitions, nil
}

// GetGame fetches a single game from the game-registry service via gRPC
func (c *GRPCGameRegistryClient) GetGame(ctx context.Context, id string) (*GameDefinition, error) {
	resp, err := c.client.GetGame(ctx, &gamesv1.GetGameRequest{
		GameId: id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get game %s from registry: %w", id, err)
	}
	if resp.Game == nil {
		return nil, fmt.Errorf("game %s not found", id)
	}
	return gameToDefinition(resp.Game), nil
}

// GetFeaturedGames fetches featured games from the game-registry service
func (c *GRPCGameRegistryClient) GetFeaturedGames(ctx context.Context, limit int32) ([]*GameDefinition, error) {
	resp, err := c.client.GetFeaturedGames(ctx, &gamesv1.GetFeaturedGamesRequest{
		Limit: limit,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get featured games: %w", err)
	}

	definitions := make([]*GameDefinition, 0, len(resp.Games))
	for _, g := range resp.Games {
		definitions = append(definitions, gameSummaryToDefinition(g))
	}
	return definitions, nil
}

// GetCategories fetches game categories from the game-registry service
func (c *GRPCGameRegistryClient) GetCategories(ctx context.Context) ([]*gamesv1.GameCategoryInfo, error) {
	resp, err := c.client.GetCategories(ctx, &gamesv1.GetCategoriesRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	return resp.Categories, nil
}

// gameSummaryToDefinition converts a proto GameSummary to a GameDefinition
func gameSummaryToDefinition(g *gamesv1.GameSummary) *GameDefinition {
	d := &GameDefinition{
		ID:          g.GameId,
		Name:        g.Name,
		Type:        g.CategoryName,
		Description: "",
		Thumbnail:   g.ThumbnailUrl,
		RTP:         g.Rtp,
		Status:      statusToString(g.Status),
	}
	if g.MinBet != nil {
		d.MinBet = g.MinBet.Amount
	}
	if g.MaxBet != nil {
		d.MaxBet = g.MaxBet.Amount
	}
	return d
}

// gameToDefinition converts a proto Game to a GameDefinition
func gameToDefinition(g *gamesv1.Game) *GameDefinition {
	d := &GameDefinition{
		ID:          g.GameId,
		Name:        g.Name,
		Type:        g.CategoryName,
		Description: g.Description,
		Thumbnail:   g.ThumbnailUrl,
		RTP:         g.Rtp,
		HouseEdge:   100.0 - g.Rtp,
		Status:      statusToString(g.Status),
	}
	if g.MinBet != nil {
		d.MinBet = g.MinBet.Amount
	}
	if g.MaxBet != nil {
		d.MaxBet = g.MaxBet.Amount
	}
	return d
}

// statusToString converts a proto Status to a string
func statusToString(s gamesv1.Status) string {
	switch s {
	case gamesv1.Status_STATUS_ACTIVE:
		return "active"
	case gamesv1.Status_STATUS_INACTIVE:
		return "inactive"
	default:
		return "unknown"
	}
}
