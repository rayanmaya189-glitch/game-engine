package handler

import (
	"context"
	"time"

	commonv1 "github.com/game_engine/common-service/proto/gen/go/common/v1"
	walletsv1 "github.com/game_engine/common-service/proto/gen/go/wallet/v1"

	"github.com/game_engine/wallet-service/internal/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// WalletHandler handles gRPC requests for wallet service
type WalletHandler struct {
	walletsv1.UnimplementedWalletServiceServer
	walletService *service.WalletService
}

// NewWalletHandler creates a new wallet handler
func NewWalletHandler(walletService *service.WalletService) *WalletHandler {
	return &WalletHandler{walletService: walletService}
}

// GetBalance retrieves player balance
func (h *WalletHandler) GetBalance(ctx context.Context, req *walletsv1.GetBalanceRequest) (*walletsv1.GetBalanceResponse, error) {
	balanceType := req.BalanceType.String()
	if balanceType == "" {
		balanceType = service.BalanceTypeReal
	}

	wallet, err := h.walletService.GetBalance(ctx, req.UserId, service.CurrencyUSD, balanceType)
	if err != nil {
		return nil, err
	}

	available := wallet.Amount - wallet.LockedAmount

	return &walletsv1.GetBalanceResponse{
		Balance:         wallet.ToProto(),
		LockedAmount:    &commonv1.Money{Amount: wallet.LockedAmount, Currency: commonv1.Currency(commonv1.Currency_value[wallet.Currency])},
		AvailableAmount: &commonv1.Money{Amount: available, Currency: commonv1.Currency(commonv1.Currency_value[wallet.Currency])},
	}, nil
}

// GetAllBalances retrieves all currency balances for a user
func (h *WalletHandler) GetAllBalances(ctx context.Context, req *walletsv1.GetAllBalancesRequest) (*walletsv1.GetAllBalancesResponse, error) {
	wallets, err := h.walletService.GetAllBalances(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	balances := make([]*commonv1.BalanceEntry, len(wallets))
	for i, w := range wallets {
		balances[i] = w.ToBalanceProto()
	}

	return &walletsv1.GetAllBalancesResponse{
		Balances: balances,
	}, nil
}

// GetTransactionHistory retrieves transaction history
func (h *WalletHandler) GetTransactionHistory(ctx context.Context, req *walletsv1.GetTransactionHistoryRequest) (*walletsv1.GetTransactionHistoryResponse, error) {
	var txTypes []string
	var statuses []string
	var startDate, endDate time.Time

	if req.Types != nil {
		for _, t := range req.Types {
			txTypes = append(txTypes, t.String())
		}
	}

	if req.Statuses != nil {
		for _, s := range req.Statuses {
			statuses = append(statuses, s.String())
		}
	}

	if req.StartDate != nil {
		startDate = time.Unix(req.StartDate.Seconds, int64(req.StartDate.Nanos))
	}

	if req.EndDate != nil {
		endDate = time.Unix(req.EndDate.Seconds, int64(req.EndDate.Nanos))
	}

	page := 1
	pageSize := 20

	if req.Pagination != nil {
		page = int(req.Pagination.Page)
		pageSize = int(req.Pagination.PageSize)
	}

	transactions, total, err := h.walletService.GetTransactionHistory(ctx, req.UserId, txTypes, statuses, startDate, endDate, page, pageSize)
	if err != nil {
		return nil, err
	}

	protoTransactions := make([]*walletsv1.Transaction, len(transactions))
	for i, t := range transactions {
		protoTransactions[i] = t.ToTransactionProto()
	}

	totalPages := (total + pageSize - 1) / pageSize

	return &walletsv1.GetTransactionHistoryResponse{
		Transactions: protoTransactions,
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

// CreateDeposit initiates a deposit
func (h *WalletHandler) CreateDeposit(ctx context.Context, req *walletsv1.CreateDepositRequest) (*walletsv1.CreateDepositResponse, error) {
	currency := service.CurrencyUSD
	if req.Amount != nil && req.Amount.Currency != 0 {
		currency = req.Amount.Currency.String()
	}

	tx, err := h.walletService.CreateDeposit(ctx, req.UserId, currency, req.PaymentMethod.String(), req.PaymentProvider, req.BonusCode, req.Amount.Amount)
	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(30 * time.Minute)

	return &walletsv1.CreateDepositResponse{
		Deposit:          tx.ToTransactionProto(),
		PaymentUrl:       "",
		PaymentReference: tx.PaymentReference,
		ExpiresAt:        timestamppb.New(expiresAt),
	}, nil
}

// ConfirmDeposit confirms a deposit completion
func (h *WalletHandler) ConfirmDeposit(ctx context.Context, req *walletsv1.ConfirmDepositRequest) (*walletsv1.ConfirmDepositResponse, error) {
	tx, err := h.walletService.ConfirmDeposit(ctx, req.TransactionId, req.ProviderStatus)
	if err != nil {
		return &walletsv1.ConfirmDepositResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &walletsv1.ConfirmDepositResponse{
		Success:     true,
		Transaction: tx.ToTransactionProto(),
		Message:     "Deposit confirmed successfully",
	}, nil
}

// CreateWithdrawal initiates a withdrawal request
func (h *WalletHandler) CreateWithdrawal(ctx context.Context, req *walletsv1.CreateWithdrawalRequest) (*walletsv1.CreateWithdrawalResponse, error) {
	currency := service.CurrencyUSD
	if req.Amount != nil && req.Amount.Currency != 0 {
		currency = req.Amount.Currency.String()
	}

	tx, err := h.walletService.CreateWithdrawal(ctx, req.UserId, currency, req.WithdrawalMethodId, req.Amount.Amount)
	if err != nil {
		return nil, err
	}

	return &walletsv1.CreateWithdrawalResponse{
		Withdrawal:       tx.ToTransactionProto(),
		ApprovalRequired: false,
		Message:          "Withdrawal initiated",
	}, nil
}

// ConfirmWithdrawal confirms a withdrawal
func (h *WalletHandler) ConfirmWithdrawal(ctx context.Context, req *walletsv1.ConfirmWithdrawalRequest) (*walletsv1.ConfirmWithdrawalResponse, error) {
	tx, err := h.walletService.ConfirmWithdrawal(ctx, req.TransactionId, req.ProviderReference, req.Status.String())
	if err != nil {
		return &walletsv1.ConfirmWithdrawalResponse{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	return &walletsv1.ConfirmWithdrawalResponse{
		Success:     true,
		Transaction: tx.ToTransactionProto(),
		Message:     "Withdrawal confirmed",
	}, nil
}
