package handler

import (
	"context"
	"time"

	commonv1 "github.com/game_engine/common-service/proto/gen/go/common/v1"
	walletsv1 "github.com/game_engine/common-service/proto/gen/go/wallet/v1"

	"github.com/game_engine/wallet-service/internal/service"
)

// CreateBonusCredit adds bonus funds
func (h *WalletHandler) CreateBonusCredit(ctx context.Context, req *walletsv1.CreateBonusCreditRequest) (*walletsv1.CreateBonusCreditResponse, error) {
	currency := service.CurrencyUSD
	if req.Amount != nil && req.Amount.Currency != 0 {
		currency = req.Amount.Currency.String()
	}

	var expiresAt *time.Time
	if req.ExpiresAt != nil {
		t := time.Unix(req.ExpiresAt.Seconds, int64(req.ExpiresAt.Nanos))
		expiresAt = &t
	}

	bonus, err := h.walletService.CreateBonusCredit(ctx, req.UserId, currency, req.BonusType.String(), req.BonusCode, req.Amount.Amount, expiresAt)
	if err != nil {
		return nil, err
	}

	return &walletsv1.CreateBonusCreditResponse{
		Success: true,
		Bonus:   bonus.ToBonusProto(),
		Message: "Bonus credited successfully",
	}, nil
}

// ReverseTransaction reverses a transaction (admin operation)
func (h *WalletHandler) ReverseTransaction(ctx context.Context, req *walletsv1.ReverseTransactionRequest) (*walletsv1.ReverseTransactionResponse, error) {
	tx, err := h.walletService.ReverseTransaction(ctx, req.TransactionId, req.Reason)
	if err != nil {
		return nil, err
	}

	return &walletsv1.ReverseTransactionResponse{
		Success:  true,
		Reversal: tx.ToTransactionProto(),
		Message:  "Transaction reversed successfully",
	}, nil
}

// GetPendingBets retrieves pending bets
func (h *WalletHandler) GetPendingBets(ctx context.Context, req *walletsv1.GetPendingBetsRequest) (*walletsv1.GetPendingBetsResponse, error) {
	page := 1
	pageSize := 20

	if req.Pagination != nil {
		page = int(req.Pagination.Page)
		pageSize = int(req.Pagination.PageSize)
	}

	bets, total, err := h.walletService.GetPendingBets(ctx, req.UserId, req.GameId, page, pageSize)
	if err != nil {
		return nil, err
	}

	protoBets := make([]*walletsv1.Bet, len(bets))
	for i, b := range bets {
		protoBets[i] = b.ToBetProto()
	}

	totalPages := (total + pageSize - 1) / pageSize

	return &walletsv1.GetPendingBetsResponse{
		Bets: protoBets,
		Pagination: &commonv1.PaginationResponse{
			Page:        int32(page),
			PageSize:    int32(pageSize),
			TotalItems:  int32(total),
			TotalPages:  int32(totalPages),
			HasNext:     page < totalPages,
			HasPrevious: page > 1,
		},
	}, nil
}
