package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	jackpotpb "github.com/game-engine/gen/go/game-engine/jackpot/v1"
)

type JackpotClient struct {
	client jackpotpb.JackpotServiceClient
	conn   *grpc.ClientConn
}

type JackpotClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewJackpotClient(config *JackpotClientConfig) (*JackpotClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		fmt.Fprintf(os.Stderr, "[WARNING] Using insecure gRPC connection to jackpot service - TLS recommended for production\n")
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	opts = append(opts, grpc.WithUnaryInterceptor(timeoutInterceptor(config.Timeout)))

	// Keep-alive for connection health
	opts = append(opts, grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:    10 * time.Second,
		Timeout: 5 * time.Second,
	}))

	conn, err := grpc.Dial(config.Address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to jackpot service: %w", err)
	}

	return &JackpotClient{
		client: jackpotpb.NewJackpotServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *JackpotClient) ListJackpots(ctx context.Context, req *jackpotpb.ListJackpotsRequest) (*jackpotpb.ListJackpotsResponse, error) {
	return c.client.ListJackpots(ctx, req)
}

func (c *JackpotClient) GetJackpot(ctx context.Context, req *jackpotpb.GetJackpotRequest) (*jackpotpb.GetJackpotResponse, error) {
	return c.client.GetJackpot(ctx, req)
}

func (c *JackpotClient) GetWinners(ctx context.Context, req *jackpotpb.GetWinnersRequest) (*jackpotpb.GetWinnersResponse, error) {
	return c.client.GetWinners(ctx, req)
}

func (c *JackpotClient) JoinJackpot(ctx context.Context, req *jackpotpb.JoinJackpotRequest) (*jackpotpb.JoinJackpotResponse, error) {
	return c.client.JoinJackpot(ctx, req)
}

func (c *JackpotClient) GetJackpotHistory(ctx context.Context, req *jackpotpb.GetJackpotHistoryRequest) (*jackpotpb.GetJackpotHistoryResponse, error) {
	return c.client.GetJackpotHistory(ctx, req)
}

func (c *JackpotClient) Close() error {
	return c.conn.Close()
}
