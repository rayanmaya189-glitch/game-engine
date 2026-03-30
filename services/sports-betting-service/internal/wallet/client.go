package wallet

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCWalletClient implements the WalletClient interface using gRPC
type GRPCWalletClient struct {
	conn   *grpc.ClientConn
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
	// The actual gRPC call would use generated stubs:
	//   client := walletv1.NewWalletServiceClient(c.conn)
	//   resp, err := client.GetBalance(ctx, &walletv1.GetBalanceRequest{
	//       UserId: userID,
	//   })
	//   if err != nil {
	//       return 0, err
	//   }
	//   return float64(resp.Balance.Amount) / 100, nil
	return 0, fmt.Errorf("gRPC stubs not yet generated; run protoc to generate wallet client stubs")
}

// DeductBalance deducts the specified amount from the user's wallet
func (c *GRPCWalletClient) DeductBalance(ctx context.Context, userID string, amount float64, betID string) error {
	// client := walletv1.NewWalletServiceClient(c.conn)
	// _, err := client.PlaceBet(ctx, &walletv1.PlaceBetRequest{
	//     UserId:  userID,
	//     BetId:   betID,
	//     Amount:  &commonv1.Money{Amount: int64(amount * 100), Currency: "USD"},
	// })
	// return err
	return fmt.Errorf("gRPC stubs not yet generated; run protoc to generate wallet client stubs")
}

// AddWinnings credits winnings to the user's wallet
func (c *GRPCWalletClient) AddWinnings(ctx context.Context, userID string, amount float64, betID string) error {
	// client := walletv1.NewWalletServiceClient(c.conn)
	// _, err := client.SettleBet(ctx, &walletv1.SettleBetRequest{
	//     BetId: betID,
	//     SettlementType: walletv1.BetSettlementType_BET_SETTLEMENT_TYPE_WON,
	//     WinAmount: &commonv1.Money{Amount: int64(amount * 100), Currency: "USD"},
	// })
	// return err
	return fmt.Errorf("gRPC stubs not yet generated; run protoc to generate wallet client stubs")
}

// RefundStake refunds a bet stake to the user's wallet
func (c *GRPCWalletClient) RefundStake(ctx context.Context, userID string, amount float64, betID string) error {
	// client := walletv1.NewWalletServiceClient(c.conn)
	// _, err := client.CancelBet(ctx, &walletv1.CancelBetRequest{
	//     BetId: betID,
	//     Reason: "bet cancelled",
	// })
	// return err
	return fmt.Errorf("gRPC stubs not yet generated; run protoc to generate wallet client stubs")
}
