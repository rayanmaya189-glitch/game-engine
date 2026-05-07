package handler

import (
	"context"

	commonv1 "github.com/game_engine/common-service/proto/gen/go/common/v1"
	walletsv1 "github.com/game_engine/common-service/proto/gen/go/wallet/v1"
	"github.com/game_engine/wallet-service/internal/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// PlaceBet locks funds for a bet
func (h *WalletHandler) PlaceBet(ctx context.Context, req *walletsv1.PlaceBetRequest) (*walletsv1.PlaceBetResponse, error) {
	currency := service.CurrencyUSD
	if req.Amount != nil && req.Amount.Currency != 0 {
		currency = req.Amount.Currency.String()
	}

	bet, wallet, err := h.walletService.PlaceBet(ctx, req.UserId, req.GameId, req.BetType, req.Selection, req.Odds, currency, req.Amount.Amount)
	if err != nil {
		return nil, err
	}

	tx := &walletsv1.Transaction{
		UserId:    req.UserId,
		Type:      commonv1.TransactionType_TRANSACTION_TYPE_BET,
		Status:    commonv1.TransactionStatus_TRANSACTION_STATUS_COMPLETED,
		GameId:    req.GameId,
		BetId:     bet.BetID,
		CreatedAt: timestamppb.New(bet.PlacedAt),
	}

	_ = tx

	return &walletsv1.PlaceBetResponse{
		Bet:          bet.ToTransactionProto(),
		NewBalance:   &commonv1.Money{Amount: wallet.Amount, Currency: commonv1.Currency(commonv1.Currency_value[wallet.Currency])},
		LockedAmount: &commonv1.Money{Amount: wallet.LockedAmount, Currency: commonv1.Currency(commonv1.Currency_value[wallet.Currency])},
		BetId:        bet.BetID,
		Message:      "Bet placed successfully",
	}, nil
}

// SettleBet processes bet result
func (h *WalletHandler) SettleBet(ctx context.Context, req *walletsv1.SettleBetRequest) (*walletsv1.SettleBetResponse, error) {
	bet, err := h.walletService.SettleBet(ctx, req.BetId, req.SettlementType.String(), req.WinAmount.Amount, req.Result)
	if err != nil {
		return nil, err
	}

	wallet, err := h.walletService.GetBalance(ctx, bet.UserID, bet.Currency, service.BalanceTypeReal)
	if err != nil {
		return nil, err
	}

	var winTx *walletsv1.Transaction
	if req.WinAmount.Amount > 0 {
		winTx = &walletsv1.Transaction{
			UserId:      bet.UserID,
			Type:        commonv1.TransactionType_TRANSACTION_TYPE_WIN,
			Status:      commonv1.TransactionStatus_TRANSACTION_STATUS_COMPLETED,
			Amount:      &commonv1.TransactionAmount{Requested: req.WinAmount},
			GameId:      bet.GameID,
			BetId:       bet.BetID,
			Description: req.Result,
		}
	}

	return &walletsv1.SettleBetResponse{
		Success:    true,
		Bet:        bet.ToTransactionProto(),
		Win:        winTx,
		NewBalance: &commonv1.Money{Amount: wallet.Amount, Currency: commonv1.Currency(commonv1.Currency_value[wallet.Currency])},
		Message:    "Bet settled successfully",
	}, nil
}

// CancelBet cancels a pending bet
func (h *WalletHandler) CancelBet(ctx context.Context, req *walletsv1.CancelBetRequest) (*walletsv1.CancelBetResponse, error) {
	bet, err := h.walletService.CancelBet(ctx, req.BetId, req.Reason)
	if err != nil {
		return nil, err
	}

	wallet, err := h.walletService.GetBalance(ctx, bet.UserID, bet.Currency, service.BalanceTypeReal)
	if err != nil {
		return nil, err
	}

	return &walletsv1.CancelBetResponse{
		Success: true,
		Refund: &walletsv1.Transaction{
			UserId:      bet.UserID,
			Type:        commonv1.TransactionType_TRANSACTION_TYPE_REFUND,
			Status:      commonv1.TransactionStatus_TRANSACTION_STATUS_COMPLETED,
			Amount:      &commonv1.TransactionAmount{Requested: &commonv1.Money{Amount: bet.Stake}},
			GameId:      bet.GameID,
			BetId:       bet.BetID,
			Description: req.Reason,
		},
		NewBalance: &commonv1.Money{Amount: wallet.Amount, Currency: commonv1.Currency(commonv1.Currency_value[wallet.Currency])},
		Message:    "Bet cancelled successfully",
	}, nil
}
