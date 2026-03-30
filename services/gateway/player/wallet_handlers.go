package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	walletpb "github.com/game_engine/gen/go/game_engine/wallet/v1"

	"common/handler"
)

// GetBalance handles getting wallet balance
func (cfg *RouterConfig) GetBalance(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.WalletClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Wallet service unavailable", nil)
		return
	}

	resp, err := cfg.WalletClient.GetBalance(ctx, &walletpb.GetBalanceRequest{
		UserId: userID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"balance": map[string]interface{}{
			"main":     resp.MainBalance,
			"bonus":    resp.BonusBalance,
			"currency": resp.Currency,
		},
	})
}

// GetTransactions handles getting wallet transactions
func (cfg *RouterConfig) GetTransactions(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	if cfg.WalletClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Wallet service unavailable", nil)
		return
	}

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")
	txnType := c.Query("type")
	status := c.Query("status")

	resp, err := cfg.WalletClient.GetTransactions(ctx, &walletpb.GetTransactionsRequest{
		UserId: userID,
		Page:   page,
		Limit:  limit,
		Type:   txnType,
		Status: status,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	transactions := make([]map[string]interface{}, len(resp.Transactions))
	for i, txn := range resp.Transactions {
		transactions[i] = map[string]interface{}{
			"id":          txn.Id,
			"type":        txn.Type,
			"amount":      txn.Amount,
			"status":      txn.Status,
			"description": txn.Description,
			"created_at":  txn.CreatedAt,
		}
	}

	handler.SendSuccess(c, map[string]interface{}{
		"transactions": transactions,
		"total":        resp.Total,
		"page":         resp.Page,
		"limit":        resp.Limit,
	})
}

// Deposit handles deposit request
func (cfg *RouterConfig) Deposit(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	var req struct {
		Amount    float64 `json:"amount"`
		Method    string  `json:"method"`
		PaymentID string  `json:"paymentId"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.WalletClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Wallet service unavailable", nil)
		return
	}

	resp, err := cfg.WalletClient.Deposit(ctx, &walletpb.DepositRequest{
		UserId:    userID,
		Amount:    req.Amount,
		Method:    req.Method,
		PaymentId: req.PaymentID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Deposit successful",
		"transaction_id": resp.TransactionId,
		"new_balance":    resp.NewBalance,
	})
}

// Withdraw handles withdraw request
func (cfg *RouterConfig) Withdraw(ctx context.Context, c *app.RequestContext) {
	userID := c.GetString("user_id")
	if userID == "" {
		handler.SendErrorResponse(c, 401, handler.ErrCodeUnauthorized, "Not authenticated", nil)
		return
	}

	var req struct {
		Amount float64 `json:"amount"`
		Method string  `json:"method"`
		BankID string  `json:"bankId"`
	}

	if err := c.Bind(&req); err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, "Invalid request body", nil)
		return
	}

	if cfg.WalletClient == nil {
		handler.SendErrorResponse(c, 503, handler.ErrCodeServiceUnavailable, "Wallet service unavailable", nil)
		return
	}

	resp, err := cfg.WalletClient.Withdraw(ctx, &walletpb.WithdrawRequest{
		UserId: userID,
		Amount: req.Amount,
		Method: req.Method,
		BankId: req.BankID,
	})

	if err != nil {
		handler.SendErrorResponse(c, 400, handler.ErrCodeBadRequest, err.Error(), nil)
		return
	}

	handler.SendSuccess(c, map[string]interface{}{
		"message":        "Withdrawal request submitted",
		"transaction_id": resp.TransactionId,
	})
}
