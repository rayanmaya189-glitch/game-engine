package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	walletpb "github.com/game-engine/gen/go/gameengine/wallet/v1"
)

type WalletClient struct {
	client walletpb.WalletServiceClient
	conn   *grpc.ClientConn
}

type WalletClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewWalletClient(config *WalletClientConfig) (*WalletClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	opts = append(opts, grpc.WithUnaryInterceptor(timeoutInterceptor(config.Timeout)))

	conn, err := grpc.Dial(config.Address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to wallet service: %w", err)
	}

	return &WalletClient{
		client: walletpb.NewWalletServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *WalletClient) GetBalance(ctx context.Context, req *walletpb.GetBalanceRequest) (*walletpb.GetBalanceResponse, error) {
	return c.client.GetBalance(ctx, req)
}

func (c *WalletClient) GetTransactions(ctx context.Context, req *walletpb.GetTransactionsRequest) (*walletpb.GetTransactionsResponse, error) {
	return c.client.GetTransactions(ctx, req)
}

func (c *WalletClient) Deposit(ctx context.Context, req *walletpb.DepositRequest) (*walletpb.DepositResponse, error) {
	return c.client.Deposit(ctx, req)
}

func (c *WalletClient) Withdraw(ctx context.Context, req *walletpb.WithdrawRequest) (*walletpb.WithdrawResponse, error) {
	return c.client.Withdraw(ctx, req)
}

func (c *WalletClient) AdjustBalance(ctx context.Context, req *walletpb.AdjustBalanceRequest) (*walletpb.AdjustBalanceResponse, error) {
	return c.client.AdjustBalance(ctx, req)
}

func (c *WalletClient) GetAllTransactions(ctx context.Context, req *walletpb.GetAllTransactionsRequest) (*walletpb.GetAllTransactionsResponse, error) {
	return c.client.GetAllTransactions(ctx, req)
}

func (c *WalletClient) Close() error {
	return c.conn.Close()
}
