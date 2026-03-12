package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	authpb "github.com/game_engine/gen/go/game_engine/auth/v1"
)

type AuthClient struct {
	client authpb.AuthServiceClient
	conn   *grpc.ClientConn
}

type AuthClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewAuthClient(config *AuthClientConfig) (*AuthClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	opts = append(opts, grpc.WithUnaryInterceptor(timeoutInterceptor(config.Timeout)))

	conn, err := grpc.Dial(config.Address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}

	return &AuthClient{
		client: authpb.NewAuthServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *AuthClient) Register(ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error) {
	return c.client.Register(ctx, req)
}

func (c *AuthClient) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	return c.client.Login(ctx, req)
}

func (c *AuthClient) Logout(ctx context.Context, req *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {
	return c.client.Logout(ctx, req)
}

func (c *AuthClient) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {
	return c.client.RefreshToken(ctx, req)
}

func (c *AuthClient) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	return c.client.ValidateToken(ctx, req)
}

func (c *AuthClient) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.ChangePasswordResponse, error) {
	return c.client.ChangePassword(ctx, req)
}

func (c *AuthClient) Enable2FA(ctx context.Context, req *authpb.Enable2FARequest) (*authpb.Enable2FAResponse, error) {
	return c.client.Enable2FA(ctx, req)
}

func (c *AuthClient) Verify2FA(ctx context.Context, req *authpb.Verify2FARequest) (*authpb.Verify2FAResponse, error) {
	return c.client.Verify2FA(ctx, req)
}

func (c *AuthClient) VerifyEmail(ctx context.Context, req *authpb.VerifyEmailRequest) (*authpb.VerifyEmailResponse, error) {
	return c.client.VerifyEmail(ctx, req)
}

func (c *AuthClient) VerifyPhone(ctx context.Context, req *authpb.VerifyPhoneRequest) (*authpb.VerifyPhoneResponse, error) {
	return c.client.VerifyPhone(ctx, req)
}

func (c *AuthClient) ResetPassword(ctx context.Context, req *authpb.ResetPasswordRequest) (*authpb.ResetPasswordResponse, error) {
	return c.client.ResetPassword(ctx, req)
}

func (c *AuthClient) Close() error {
	return c.conn.Close()
}

func timeoutInterceptor(timeout time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func metadataInterceptor(ctx context.Context, md metadata.MD) context.Context {
	return metadata.NewOutgoingContext(ctx, md)
}
