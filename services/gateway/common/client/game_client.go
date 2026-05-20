package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	gamepb "github.com/game_engine/common-service/proto/gen/go/game/v1"
)

type GameClient struct {
	client gamepb.GameRegistryServiceClient
	conn   *grpc.ClientConn
}

type GameClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewGameClient(config *GameClientConfig) (*GameClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	opts = append(opts, grpc.WithUnaryInterceptor(timeoutInterceptor(config.Timeout)))

	conn, err := grpc.Dial(config.Address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to game registry service: %w", err)
	}

	return &GameClient{
		client: gamepb.NewGameRegistryServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *GameClient) ListGames(ctx context.Context, req *gamepb.ListGamesRequest) (*gamepb.ListGamesResponse, error) {
	return c.client.ListGames(ctx, req)
}

func (c *GameClient) GetGame(ctx context.Context, req *gamepb.GetGameRequest) (*gamepb.GetGameResponse, error) {
	return c.client.GetGame(ctx, req)
}

func (c *GameClient) GetGameURL(ctx context.Context, req *gamepb.GetGameURLRequest) (*gamepb.GetGameURLResponse, error) {
	return c.client.GetGameURL(ctx, req)
}

func (c *GameClient) GetCategories(ctx context.Context, req *gamepb.GetCategoriesRequest) (*gamepb.GetCategoriesResponse, error) {
	return c.client.GetCategories(ctx, req)
}

func (c *GameClient) GetFeaturedGames(ctx context.Context, req *gamepb.GetFeaturedGamesRequest) (*gamepb.GetFeaturedGamesResponse, error) {
	return c.client.GetFeaturedGames(ctx, req)
}

func (c *GameClient) GetPopularGames(ctx context.Context, req *gamepb.GetPopularGamesRequest) (*gamepb.GetPopularGamesResponse, error) {
	return c.client.GetPopularGames(ctx, req)
}


func (c *GameClient) Close() error {
	return c.conn.Close()
}
