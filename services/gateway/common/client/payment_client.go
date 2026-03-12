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

	paymentpb "github.com/game-engine/gen/go/game-engine/payment/v1"
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

func (c *PaymentClient) CreatePayment(ctx context.Context, req *paymentpb.CreatePaymentRequest) (*paymentpb.CreatePaymentResponse, error) {
	return c.client.CreatePayment(ctx, req)
}

func (c *PaymentClient) GetPayment(ctx context.Context, req *paymentpb.GetPaymentRequest) (*paymentpb.GetPaymentResponse, error) {
	return c.client.GetPayment(ctx, req)
}

func (c *PaymentClient) ApprovePayment(ctx context.Context, req *paymentpb.ApprovePaymentRequest) (*paymentpb.ApprovePaymentResponse, error) {
	return c.client.ApprovePayment(ctx, req)
}

func (c *PaymentClient) RejectPayment(ctx context.Context, req *paymentpb.RejectPaymentRequest) (*paymentpb.RejectPaymentResponse, error) {
	return c.client.RejectPayment(ctx, req)
}

func (c *PaymentClient) ProcessPayment(ctx context.Context, req *paymentpb.ProcessPaymentRequest) (*paymentpb.ProcessPaymentResponse, error) {
	return c.client.ProcessPayment(ctx, req)
}

func (c *PaymentClient) ListPayments(ctx context.Context, req *paymentpb.ListPaymentsRequest) (*paymentpb.ListPaymentsResponse, error) {
	return c.client.ListPayments(ctx, req)
}

func (c *PaymentClient) GetPaymentMethods(ctx context.Context, req *paymentpb.GetPaymentMethodsRequest) (*paymentpb.GetPaymentMethodsResponse, error) {
	return c.client.GetPaymentMethods(ctx, req)
}

func (c *PaymentClient) Close() error {
	return c.conn.Close()
}
