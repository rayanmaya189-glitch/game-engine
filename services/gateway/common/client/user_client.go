package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	userpb "github.com/game_engine/common-service/proto/gen/go/user/v1"
)

type UserClient struct {
	client userpb.UserServiceClient
	conn   *grpc.ClientConn
}

type UserClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewUserClient(config *UserClientConfig) (*UserClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	opts = append(opts, grpc.WithUnaryInterceptor(timeoutInterceptor(config.Timeout)))

	conn, err := grpc.Dial(config.Address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %w", err)
	}

	return &UserClient{
		client: userpb.NewUserServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *UserClient) GetProfile(ctx context.Context, req *userpb.GetProfileRequest) (*userpb.GetProfileResponse, error) {
	return c.client.GetProfile(ctx, req)
}

func (c *UserClient) UpdateProfile(ctx context.Context, req *userpb.UpdateProfileRequest) (*userpb.UpdateProfileResponse, error) {
	return c.client.UpdateProfile(ctx, req)
}

func (c *UserClient) GetKYCStatus(ctx context.Context, req *userpb.GetKYCStatusRequest) (*userpb.GetKYCStatusResponse, error) {
	return c.client.GetKYCStatus(ctx, req)
}

func (c *UserClient) SubmitKYC(ctx context.Context, req *userpb.SubmitKYCRequest) (*userpb.SubmitKYCResponse, error) {
	return c.client.SubmitKYC(ctx, req)
}

func (c *UserClient) GetPlayerSettings(ctx context.Context, req *userpb.GetPlayerSettingsRequest) (*userpb.GetPlayerSettingsResponse, error) {
	return c.client.GetPlayerSettings(ctx, req)
}

func (c *UserClient) UpdatePlayerSettings(ctx context.Context, req *userpb.UpdatePlayerSettingsRequest) (*userpb.UpdatePlayerSettingsResponse, error) {
	return c.client.UpdatePlayerSettings(ctx, req)
}

func (c *UserClient) GetPlayerByAdmin(ctx context.Context, req *userpb.GetPlayerByAdminRequest) (*userpb.GetPlayerByAdminResponse, error) {
	return c.client.GetPlayerByAdmin(ctx, req)
}

func (c *UserClient) ListPlayers(ctx context.Context, req *userpb.ListPlayersRequest) (*userpb.ListPlayersResponse, error) {
	return c.client.ListPlayers(ctx, req)
}

func (c *UserClient) UpdatePlayerStatus(ctx context.Context, req *userpb.UpdatePlayerStatusRequest) (*userpb.UpdatePlayerStatusResponse, error) {
	return c.client.UpdatePlayerStatus(ctx, req)
}

func (c *UserClient) SetDepositLimit(ctx context.Context, req *userpb.SetDepositLimitRequest) (*userpb.SetDepositLimitResponse, error) {
	return c.client.SetDepositLimit(ctx, req)
}

func (c *UserClient) SetBetLimit(ctx context.Context, req *userpb.SetBetLimitRequest) (*userpb.SetBetLimitResponse, error) {
	return c.client.SetBetLimit(ctx, req)
}

func (c *UserClient) SetLossLimit(ctx context.Context, req *userpb.SetLossLimitRequest) (*userpb.SetLossLimitResponse, error) {
	return c.client.SetLossLimit(ctx, req)
}

func (c *UserClient) GetPlayerLimits(ctx context.Context, req *userpb.GetPlayerLimitsRequest) (*userpb.GetPlayerLimitsResponse, error) {
	return c.client.GetPlayerLimits(ctx, req)
}

func (c *UserClient) SelfExclude(ctx context.Context, req *userpb.SelfExcludeRequest) (*userpb.SelfExcludeResponse, error) {
	return c.client.SelfExclude(ctx, req)
}

func (c *UserClient) Close() error {
	return c.conn.Close()
}
