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

	paymentpb "github.com/game_engine/common-service/proto/gen/go/payment/v1"
)

type PaymentClient struct {
	client paymentpb.PaymentServiceClient
	conn   *grpc.ClientConn
}

type PaymentClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewPaymentClient(config *PaymentClientConfig) (*PaymentClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		fmt.Fprintf(os.Stderr, "[WARNING] Using insecure gRPC connection to payment service - TLS recommended for production\n")
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
		return nil, fmt.Errorf("failed to connect to payment service: %w", err)
	}

	return &PaymentClient{
		client: paymentpb.NewPaymentServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *PaymentClient) CreateDeposit(ctx context.Context, req *paymentpb.CreateDepositRequest) (*paymentpb.CreateDepositResponse, error) {
	return c.client.CreateDeposit(ctx, req)
}

func (c *PaymentClient) ProcessDeposit(ctx context.Context, req *paymentpb.ProcessDepositRequest) (*paymentpb.ProcessDepositResponse, error) {
	return c.client.ProcessDeposit(ctx, req)
}

func (c *PaymentClient) CreateWithdrawal(ctx context.Context, req *paymentpb.CreateWithdrawalRequest) (*paymentpb.CreateWithdrawalResponse, error) {
	return c.client.CreateWithdrawal(ctx, req)
}

func (c *PaymentClient) ProcessWithdrawal(ctx context.Context, req *paymentpb.ProcessWithdrawalRequest) (*paymentpb.ProcessWithdrawalResponse, error) {
	return c.client.ProcessWithdrawal(ctx, req)
}

func (c *PaymentClient) RefundPayment(ctx context.Context, req *paymentpb.RefundPaymentRequest) (*paymentpb.RefundPaymentResponse, error) {
	return c.client.RefundPayment(ctx, req)
}

func (c *PaymentClient) GetPayment(ctx context.Context, req *paymentpb.GetPaymentRequest) (*paymentpb.GetPaymentResponse, error) {
	return c.client.GetPayment(ctx, req)
}

func (c *PaymentClient) GetPaymentByExternalId(ctx context.Context, req *paymentpb.GetPaymentByExternalIdRequest) (*paymentpb.GetPaymentByExternalIdResponse, error) {
	return c.client.GetPaymentByExternalId(ctx, req)
}

func (c *PaymentClient) GetUserPayments(ctx context.Context, req *paymentpb.GetUserPaymentsRequest) (*paymentpb.GetUserPaymentsResponse, error) {
	return c.client.GetUserPayments(ctx, req)
}

func (c *PaymentClient) GetUserPaymentSummary(ctx context.Context, req *paymentpb.GetUserPaymentSummaryRequest) (*paymentpb.GetUserPaymentSummaryResponse, error) {
	return c.client.GetUserPaymentSummary(ctx, req)
}

func (c *PaymentClient) GetSupportedMethods(ctx context.Context, req *paymentpb.GetSupportedMethodsRequest) (*paymentpb.GetSupportedMethodsResponse, error) {
	return c.client.GetSupportedMethods(ctx, req)
}

func (c *PaymentClient) GetSupportedCurrencies(ctx context.Context, req *paymentpb.GetSupportedCurrenciesRequest) (*paymentpb.GetSupportedCurrenciesResponse, error) {
	return c.client.GetSupportedCurrencies(ctx, req)
}

func (c *PaymentClient) CancelPayment(ctx context.Context, req *paymentpb.CancelPaymentRequest) (*paymentpb.CancelPaymentResponse, error) {
	return c.client.CancelPayment(ctx, req)
}

func (c *PaymentClient) Close() error {
	return c.conn.Close()
}
