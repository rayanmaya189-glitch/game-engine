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

	tournamentpb "github.com/game-engine/gen/go/gameengine/tournament/v1"
)

type TournamentClient struct {
	client tournamentpb.TournamentServiceClient
	conn   *grpc.ClientConn
}

type TournamentClientConfig struct {
	Address string
	Timeout time.Duration
	UseTLS  bool
}

func NewTournamentClient(config *TournamentClientConfig) (*TournamentClient, error) {
	var opts []grpc.DialOption

	if config.UseTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		fmt.Fprintf(os.Stderr, "[WARNING] Using insecure gRPC connection to tournament service - TLS recommended for production\n")
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
		return nil, fmt.Errorf("failed to connect to tournament service: %w", err)
	}

	return &TournamentClient{
		client: tournamentpb.NewTournamentServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *TournamentClient) ListTournaments(ctx context.Context, req *tournamentpb.ListTournamentsRequest) (*tournamentpb.ListTournamentsResponse, error) {
	return c.client.ListTournaments(ctx, req)
}

func (c *TournamentClient) GetTournament(ctx context.Context, req *tournamentpb.GetTournamentRequest) (*tournamentpb.GetTournamentResponse, error) {
	return c.client.GetTournament(ctx, req)
}

func (c *TournamentClient) JoinTournament(ctx context.Context, req *tournamentpb.JoinTournamentRequest) (*tournamentpb.JoinTournamentResponse, error) {
	return c.client.JoinTournament(ctx, req)
}

func (c *TournamentClient) LeaveTournament(ctx context.Context, req *tournamentpb.LeaveTournamentRequest) (*tournamentpb.LeaveTournamentResponse, error) {
	return c.client.LeaveTournament(ctx, req)
}

func (c *TournamentClient) GetLeaderboard(ctx context.Context, req *tournamentpb.GetLeaderboardRequest) (*tournamentpb.GetLeaderboardResponse, error) {
	return c.client.GetLeaderboard(ctx, req)
}

func (c *TournamentClient) UpdateScore(ctx context.Context, req *tournamentpb.UpdateScoreRequest) (*tournamentpb.UpdateScoreResponse, error) {
	return c.client.UpdateScore(ctx, req)
}

func (c *TournamentClient) GetMyTournaments(ctx context.Context, req *tournamentpb.GetMyTournamentsRequest) (*tournamentpb.GetMyTournamentsResponse, error) {
	return c.client.GetMyTournaments(ctx, req)
}

func (c *TournamentClient) Close() error {
	return c.conn.Close()
}
