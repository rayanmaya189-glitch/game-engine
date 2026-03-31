package wallet

import (
	"context"
	"fmt"
	"time"

	commonv1 "game_engine/gen/go/common/v1"
	walletv1 "game_engine/gen/go/wallet/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCWalletClient implements the WalletClient interface using gRPC
type GRPCWalletClient struct {
	conn   *grpc.ClientConn
	client walletv1.WalletServiceClient
	target string
}

// NewGRPCWalletClient creates a new gRPC wallet client
func NewGRPCWalletClient(target string) (*GRPCWalletClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to wallet service at %s: %w", target, err)
	}

	return &GRPCWalletClient{
		conn:   conn,
		client: walletv1.NewWalletServiceClient(conn),
		target: target,
	}, nil
}

// Close closes the gRPC connection
func (c *GRPCWalletClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// GetBalance retrieves the user's wallet balance via gRPC
func (c *GRPCWalletClient) GetBalance(ctx context.Context, userID string) (float64, error) {
	resp, err := c.client.GetBalance(ctx, &walletv1.GetBalanceRequest{
		UserId: userID,
	})
	if err != nil {
		return 0, fmt.Errorf("failed to get balance for user %s: %w", userID, err)
	}
	if resp.Balance == nil {
		return 0, nil
	}
	return float64(resp.Balance.Amount) / 100, nil
}

// DeductBalance deducts the specified amount from the user's wallet via PlaceBet
func (c *GRPCWalletClient) DeductBalance(ctx context.Context, userID string, amount float64, betID string) error {
	_, err := c.client.PlaceBet(ctx, &walletv1.PlaceBetRequest{
		UserId: userID,
		GameId: betID,
		Amount: &commonv1.Money{
			Amount: int64(amount * 100),
		},
	})
	if err != nil {
		return fmt.Errorf("failed to deduct balance for user %s: %w", userID, err)
	}
	return nil
}

// AddWinnings credits winnings to the user's wallet via SettleBet
func (c *GRPCWalletClient) AddWinnings(ctx context.Context, userID string, amount float64, betID string) error {
	_, err := c.client.SettleBet(ctx, &walletv1.SettleBetRequest{
		BetId:          betID,
		SettlementType: walletv1.BetSettlementType_BET_SETTLEMENT_TYPE_WON,
		WinAmount: &commonv1.Money{
			Amount: int64(amount * 100),
		},
	})
	if err != nil {
		return fmt.Errorf("failed to add winnings for user %s: %w", userID, err)
	}
	return nil
}

// RefundStake refunds a bet stake to the user's wallet via CancelBet
func (c *GRPCWalletClient) RefundStake(ctx context.Context, userID string, amount float64, betID string) error {
	_, err := c.client.CancelBet(ctx, &walletv1.CancelBetRequest{
		BetId:  betID,
		Reason: "bet cancelled",
	})
	if err != nil {
		return fmt.Errorf("failed to refund stake for user %s: %w", userID, err)
	}
	return nil
}
