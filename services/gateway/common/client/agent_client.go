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

	agentpb "github.com/game_engine/gen/go/game_engine/agent/v1"
)

type AgentClient struct {
	client agentpb.AgentServiceClient
	conn   *grpc.ClientConn
}

type AgentClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewAgentClient(config *AgentClientConfig) (*AgentClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		fmt.Fprintf(os.Stderr, "[WARNING] Using insecure gRPC connection to agent service - TLS recommended for production\n")
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
		return nil, fmt.Errorf("failed to connect to agent service: %w", err)
	}

	return &AgentClient{
		client: agentpb.NewAgentServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *AgentClient) ListPlayers(ctx context.Context, req *agentpb.ListPlayersRequest) (*agentpb.ListPlayersResponse, error) {
	return c.client.ListPlayers(ctx, req)
}

func (c *AgentClient) GetPlayer(ctx context.Context, req *agentpb.GetPlayerRequest) (*agentpb.GetPlayerResponse, error) {
	return c.client.GetPlayer(ctx, req)
}

func (c *AgentClient) UpdatePlayerLimit(ctx context.Context, req *agentpb.UpdatePlayerLimitRequest) (*agentpb.UpdatePlayerLimitResponse, error) {
	return c.client.UpdatePlayerLimit(ctx, req)
}

func (c *AgentClient) GetDashboard(ctx context.Context, req *agentpb.GetDashboardRequest) (*agentpb.GetDashboardResponse, error) {
	return c.client.GetDashboard(ctx, req)
}

func (c *AgentClient) Close() error {
	return c.conn.Close()
}
