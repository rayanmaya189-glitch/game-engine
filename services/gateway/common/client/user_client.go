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

func (c *UserClient) GetUserByID(ctx context.Context, req *userpb.GetUserByIDRequest) (*userpb.GetUserByIDResponse, error) {
	return c.client.GetUserByID(ctx, req)
}

func (c *UserClient) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	return c.client.ListUsers(ctx, req)
}

func (c *UserClient) UpdateUserStatus(ctx context.Context, req *userpb.UpdateUserStatusRequest) (*userpb.UpdateUserStatusResponse, error) {
	return c.client.UpdateUserStatus(ctx, req)
}

func (c *UserClient) GetUserStats(ctx context.Context, req *userpb.GetUserStatsRequest) (*userpb.GetUserStatsResponse, error) {
	return c.client.GetUserStats(ctx, req)
}

func (c *UserClient) GetKYCList(ctx context.Context, req *userpb.GetKYCListRequest) (*userpb.GetKYCListResponse, error) {
	return c.client.GetKYCList(ctx, req)
}

func (c *UserClient) ApproveKYC(ctx context.Context, req *userpb.ApproveKYCRequest) (*userpb.ApproveKYCResponse, error) {
	return c.client.ApproveKYC(ctx, req)
}

func (c *UserClient) RejectKYC(ctx context.Context, req *userpb.RejectKYCRequest) (*userpb.RejectKYCResponse, error) {
	return c.client.RejectKYC(ctx, req)
}

func (c *UserClient) Close() error {
	return c.conn.Close()
}
