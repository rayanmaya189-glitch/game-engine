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

	merchantpb "github.com/game_engine/common-service/proto/gen/go/merchant/v1"
)

type MerchantClient struct {
	client merchantpb.MerchantServiceClient
	conn   *grpc.ClientConn
}

type MerchantClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewMerchantClient(config *MerchantClientConfig) (*MerchantClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		fmt.Fprintf(os.Stderr, "[WARNING] Using insecure gRPC connection to merchant service - TLS recommended for production\n")
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
		return nil, fmt.Errorf("failed to connect to merchant service: %w", err)
	}

	return &MerchantClient{
		client: merchantpb.NewMerchantServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *MerchantClient) ListPlayers(ctx context.Context, req *merchantpb.ListPlayersRequest) (*merchantpb.ListPlayersResponse, error) {
	return c.client.ListPlayers(ctx, req)
}

func (c *MerchantClient) GetPlayer(ctx context.Context, req *merchantpb.GetPlayerRequest) (*merchantpb.GetPlayerResponse, error) {
	return c.client.GetPlayer(ctx, req)
}

func (c *MerchantClient) GetRevenueReport(ctx context.Context, req *merchantpb.GetRevenueReportRequest) (*merchantpb.GetRevenueReportResponse, error) {
	return c.client.GetRevenueReport(ctx, req)
}

func (c *MerchantClient) GetPlayerReport(ctx context.Context, req *merchantpb.GetPlayerReportRequest) (*merchantpb.GetPlayerReportResponse, error) {
	return c.client.GetPlayerReport(ctx, req)
}

func (c *MerchantClient) GetGameReport(ctx context.Context, req *merchantpb.GetGameReportRequest) (*merchantpb.GetGameReportResponse, error) {
	return c.client.GetGameReport(ctx, req)
}

func (c *MerchantClient) GetConfig(ctx context.Context, req *merchantpb.GetConfigRequest) (*merchantpb.GetConfigResponse, error) {
	return c.client.GetConfig(ctx, req)
}

func (c *MerchantClient) UpdateConfig(ctx context.Context, req *merchantpb.UpdateConfigRequest) (*merchantpb.UpdateConfigResponse, error) {
	return c.client.UpdateConfig(ctx, req)
}

func (c *MerchantClient) RegisterWebhook(ctx context.Context, req *merchantpb.RegisterWebhookRequest) (*merchantpb.RegisterWebhookResponse, error) {
	return c.client.RegisterWebhook(ctx, req)
}

func (c *MerchantClient) ListWebhooks(ctx context.Context, req *merchantpb.ListWebhooksRequest) (*merchantpb.ListWebhooksResponse, error) {
	return c.client.ListWebhooks(ctx, req)
}

func (c *MerchantClient) DeleteWebhook(ctx context.Context, req *merchantpb.DeleteWebhookRequest) (*merchantpb.DeleteWebhookResponse, error) {
	return c.client.DeleteWebhook(ctx, req)
}

func (c *MerchantClient) ListAgents(ctx context.Context, req *merchantpb.ListAgentsRequest) (*merchantpb.ListAgentsResponse, error) {
	return c.client.ListAgents(ctx, req)
}

func (c *MerchantClient) GetAgent(ctx context.Context, req *merchantpb.GetAgentRequest) (*merchantpb.GetAgentResponse, error) {
	return c.client.GetAgent(ctx, req)
}

func (c *MerchantClient) CreateAgent(ctx context.Context, req *merchantpb.CreateAgentRequest) (*merchantpb.CreateAgentResponse, error) {
	return c.client.CreateAgent(ctx, req)
}

func (c *MerchantClient) UpdateAgent(ctx context.Context, req *merchantpb.UpdateAgentRequest) (*merchantpb.UpdateAgentResponse, error) {
	return c.client.UpdateAgent(ctx, req)
}

func (c *MerchantClient) UpdateAgentStatus(ctx context.Context, req *merchantpb.UpdateAgentStatusRequest) (*merchantpb.UpdateAgentStatusResponse, error) {
	return c.client.UpdateAgentStatus(ctx, req)
}

func (c *MerchantClient) Close() error {
	return c.conn.Close()
}
