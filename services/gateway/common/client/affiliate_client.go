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

	affiliatepb "github.com/game-engine/gen/go/gameengine/affiliate/v1"
)

type AffiliateClient struct {
	client affiliatepb.AffiliateServiceClient
	conn   *grpc.ClientConn
}

type AffiliateClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewAffiliateClient(config *AffiliateClientConfig) (*AffiliateClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		fmt.Fprintf(os.Stderr, "[WARNING] Using insecure gRPC connection to affiliate service - TLS recommended for production\n")
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
		return nil, fmt.Errorf("failed to connect to affiliate service: %w", err)
	}

	return &AffiliateClient{
		client: affiliatepb.NewAffiliateServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *AffiliateClient) TrackClick(ctx context.Context, req *affiliatepb.TrackClickRequest) (*affiliatepb.TrackClickResponse, error) {
	return c.client.TrackClick(ctx, req)
}

func (c *AffiliateClient) GetPerformanceReport(ctx context.Context, req *affiliatepb.GetPerformanceReportRequest) (*affiliatepb.GetPerformanceReportResponse, error) {
	return c.client.GetPerformanceReport(ctx, req)
}

func (c *AffiliateClient) GetClickReports(ctx context.Context, req *affiliatepb.GetClickReportsRequest) (*affiliatepb.GetClickReportsResponse, error) {
	return c.client.GetClickReports(ctx, req)
}

func (c *AffiliateClient) GetConversionReports(ctx context.Context, req *affiliatepb.GetConversionReportsRequest) (*affiliatepb.GetConversionReportsResponse, error) {
	return c.client.GetConversionReports(ctx, req)
}

func (c *AffiliateClient) GetAffiliateLinks(ctx context.Context, req *affiliatepb.GetAffiliateLinksRequest) (*affiliatepb.GetAffiliateLinksResponse, error) {
	return c.client.GetAffiliateLinks(ctx, req)
}

func (c *AffiliateClient) CreateAffiliateLink(ctx context.Context, req *affiliatepb.CreateAffiliateLinkRequest) (*affiliatepb.CreateAffiliateLinkResponse, error) {
	return c.client.CreateAffiliateLink(ctx, req)
}

func (c *AffiliateClient) Close() error {
	return c.conn.Close()
}
