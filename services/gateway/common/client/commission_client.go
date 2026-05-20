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

	commissionpb "github.com/game_engine/common-service/proto/gen/go/commission/v1"
)

type CommissionClient struct {
	client      commissionpb.CommissionServiceClient
	claimClient commissionpb.ClaimServiceClient
	conn        *grpc.ClientConn
}

type CommissionClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewCommissionClient(config *CommissionClientConfig) (*CommissionClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		fmt.Fprintf(os.Stderr, "[WARNING] Using insecure gRPC connection to commission service - TLS recommended for production\n")
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
		return nil, fmt.Errorf("failed to connect to commission service: %w", err)
	}

	return &CommissionClient{
		client:      commissionpb.NewCommissionServiceClient(conn),
		claimClient: commissionpb.NewClaimServiceClient(conn),
		conn:        conn,
	}, nil
}

func (c *CommissionClient) SubmitClaim(ctx context.Context, req *commissionpb.SubmitClaimRequest) (*commissionpb.SubmitClaimResponse, error) {
	return c.client.SubmitClaim(ctx, req)
}

func (c *CommissionClient) GetUserClaims(ctx context.Context, req *commissionpb.GetUserClaimsRequest) (*commissionpb.GetUserClaimsResponse, error) {
	return c.client.GetUserClaims(ctx, req)
}

func (c *CommissionClient) GetClaimsByStatus(ctx context.Context, req *commissionpb.GetClaimsByStatusRequest) (*commissionpb.GetClaimsByStatusResponse, error) {
	return c.client.GetClaimsByStatus(ctx, req)
}

func (c *CommissionClient) ClaimCommission(ctx context.Context, req *commissionpb.ClaimCommissionRequest) (*commissionpb.ClaimCommissionResponse, error) {
	return c.client.ClaimCommission(ctx, req)
}

func (c *CommissionClient) GetUserSettlements(ctx context.Context, req *commissionpb.GetUserSettlementsRequest) (*commissionpb.GetUserSettlementsResponse, error) {
	return c.claimClient.GetUserSettlements(ctx, req)
}

func (c *CommissionClient) GetSettlementById(ctx context.Context, req *commissionpb.GetSettlementByIdRequest) (*commissionpb.GetSettlementByIdResponse, error) {
	return c.claimClient.GetSettlementById(ctx, req)
}

func (c *CommissionClient) GetUserTotalPending(ctx context.Context, req *commissionpb.GetUserTotalPendingRequest) (*commissionpb.GetUserTotalPendingResponse, error) {
	return c.claimClient.GetUserTotalPending(ctx, req)
}

func (c *CommissionClient) GetUserTotalSettled(ctx context.Context, req *commissionpb.GetUserTotalSettledRequest) (*commissionpb.GetUserTotalSettledResponse, error) {
	return c.claimClient.GetUserTotalSettled(ctx, req)
}

func (c *CommissionClient) GetAgentCommissions(ctx context.Context, req *commissionpb.GetAgentCommissionsRequest) (*commissionpb.GetAgentCommissionsResponse, error) {
	return c.client.GetAgentCommissions(ctx, req)
}

func (c *CommissionClient) GetPendingCommissions(ctx context.Context, req *commissionpb.GetPendingCommissionsRequest) (*commissionpb.GetPendingCommissionsResponse, error) {
	return c.client.GetPendingCommissions(ctx, req)
}

func (c *CommissionClient) GetCommissionHistory(ctx context.Context, req *commissionpb.GetCommissionHistoryRequest) (*commissionpb.GetCommissionHistoryResponse, error) {
	return c.client.GetCommissionHistory(ctx, req)
}

func (c *CommissionClient) SubmitInsuranceClaim(ctx context.Context, req *commissionpb.SubmitInsuranceClaimRequest) (*commissionpb.SubmitInsuranceClaimResponse, error) {
	return c.claimClient.SubmitInsuranceClaim(ctx, req)
}

func (c *CommissionClient) ApproveInsuranceClaim(ctx context.Context, req *commissionpb.ApproveInsuranceClaimRequest) (*commissionpb.ApproveInsuranceClaimResponse, error) {
	return c.claimClient.ApproveInsuranceClaim(ctx, req)
}

func (c *CommissionClient) RejectInsuranceClaim(ctx context.Context, req *commissionpb.RejectInsuranceClaimRequest) (*commissionpb.RejectInsuranceClaimResponse, error) {
	return c.claimClient.RejectInsuranceClaim(ctx, req)
}

func (c *CommissionClient) PayInsuranceClaim(ctx context.Context, req *commissionpb.PayInsuranceClaimRequest) (*commissionpb.PayInsuranceClaimResponse, error) {
	return c.claimClient.PayInsuranceClaim(ctx, req)
}

func (c *CommissionClient) GetUserInsuranceClaims(ctx context.Context, req *commissionpb.GetUserInsuranceClaimsRequest) (*commissionpb.GetUserInsuranceClaimsResponse, error) {
	return c.claimClient.GetUserInsuranceClaims(ctx, req)
}

func (c *CommissionClient) GetInsuranceClaimsByStatus(ctx context.Context, req *commissionpb.GetInsuranceClaimsByStatusRequest) (*commissionpb.GetInsuranceClaimsByStatusResponse, error) {
	return c.claimClient.GetInsuranceClaimsByStatus(ctx, req)
}

func (c *CommissionClient) Close() error {
	return c.conn.Close()
}
