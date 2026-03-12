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

	bonuspb "github.com/game_engine/gen/go/game_engine/bonus/v1"
)

type BonusClient struct {
	client bonuspb.BonusServiceClient
	conn   *grpc.ClientConn
}

type BonusClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewBonusClient(config *BonusClientConfig) (*BonusClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		fmt.Fprintf(os.Stderr, "[WARNING] Using insecure gRPC connection to bonus service - TLS recommended for production\n")
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
		return nil, fmt.Errorf("failed to connect to bonus service: %w", err)
	}

	return &BonusClient{
		client: bonuspb.NewBonusServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *BonusClient) ListBonuses(ctx context.Context, req *bonuspb.ListBonusesRequest) (*bonuspb.ListBonusesResponse, error) {
	return c.client.ListBonuses(ctx, req)
}

func (c *BonusClient) GetBonus(ctx context.Context, req *bonuspb.GetBonusRequest) (*bonuspb.GetBonusResponse, error) {
	return c.client.GetBonus(ctx, req)
}

func (c *BonusClient) ClaimBonus(ctx context.Context, req *bonuspb.ClaimBonusRequest) (*bonuspb.ClaimBonusResponse, error) {
	return c.client.ClaimBonus(ctx, req)
}

func (c *BonusClient) GetUserBonuses(ctx context.Context, req *bonuspb.GetUserBonusesRequest) (*bonuspb.GetUserBonusesResponse, error) {
	return c.client.GetUserBonuses(ctx, req)
}

func (c *BonusClient) CreateRebetClaim(ctx context.Context, req *bonuspb.CreateRebetClaimRequest) (*bonuspb.CreateRebetClaimResponse, error) {
	return c.client.CreateRebetClaim(ctx, req)
}

func (c *BonusClient) GetUserRebetClaims(ctx context.Context, req *bonuspb.GetUserRebetClaimsRequest) (*bonuspb.GetUserRebetClaimsResponse, error) {
	return c.client.GetUserRebetClaims(ctx, req)
}

func (c *BonusClient) GetClaimableRebets(ctx context.Context, req *bonuspb.GetClaimableRebetsRequest) (*bonuspb.GetClaimableRebetsResponse, error) {
	return c.client.GetClaimableRebets(ctx, req)
}

func (c *BonusClient) ClaimRebet(ctx context.Context, req *bonuspb.ClaimRebetRequest) (*bonuspb.ClaimRebetResponse, error) {
	return c.client.ClaimRebet(ctx, req)
}

func (c *BonusClient) SubmitInsuranceClaim(ctx context.Context, req *bonuspb.SubmitInsuranceClaimRequest) (*bonuspb.SubmitInsuranceClaimResponse, error) {
	return c.client.SubmitInsuranceClaim(ctx, req)
}

func (c *BonusClient) GetUserInsuranceClaims(ctx context.Context, req *bonuspb.GetUserInsuranceClaimsRequest) (*bonuspb.GetUserInsuranceClaimsResponse, error) {
	return c.client.GetUserInsuranceClaims(ctx, req)
}

func (c *BonusClient) Close() error {
	return c.conn.Close()
}
