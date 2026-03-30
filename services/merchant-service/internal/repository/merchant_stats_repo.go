package repository

import (
	"context"
	"fmt"

	"github.com/game_engine/merchant-service/internal/model"
)

// Reports
func (r *MerchantRepository) GetRevenueReport(ctx context.Context, merchantID, startDate, endDate string) (*model.RevenueReport, error) {
	var report model.RevenueReport

	query := `SELECT COALESCE(SUM(total_revenue), 0), COALESCE(SUM(total_deposits), 0), COALESCE(SUM(total_withdrawals), 0), COUNT(DISTINCT player_id) FROM merchant_reports WHERE merchant_id = $1`
	args := []interface{}{merchantID}
	argNum := 2

	if startDate != "" {
		query += fmt.Sprintf(" AND created_at >= $%d", argNum)
		args = append(args, startDate)
		argNum++
	}
	if endDate != "" {
		query += fmt.Sprintf(" AND created_at <= $%d", argNum)
		args = append(args, endDate)
	}

	err := r.db.QueryRow(ctx, query, args...).Scan(&report.TotalRevenue, &report.TotalDeposits, &report.TotalWithdrawals, &report.TotalPlayers)
	if err != nil {
		return nil, fmt.Errorf("failed to get revenue report: %w", err)
	}
	return &report, nil
}

func (r *MerchantRepository) GetPlayerReport(ctx context.Context, merchantID, playerID, startDate, endDate string) (*model.PlayerReport, error) {
	var report model.PlayerReport

	query := `SELECT COALESCE(SUM(total_bets), 0), COALESCE(SUM(total_wins), 0), COALESCE(SUM(net_revenue), 0), COUNT(*) FROM player_reports WHERE merchant_id = $1 AND player_id = $2`
	args := []interface{}{merchantID, playerID}
	argNum := 3

	if startDate != "" {
		query += fmt.Sprintf(" AND created_at >= $%d", argNum)
		args = append(args, startDate)
		argNum++
	}
	if endDate != "" {
		query += fmt.Sprintf(" AND created_at <= $%d", argNum)
		args = append(args, endDate)
	}

	err := r.db.QueryRow(ctx, query, args...).Scan(&report.TotalBets, &report.TotalWins, &report.NetRevenue, &report.GamesPlayed)
	if err != nil {
		return nil, fmt.Errorf("failed to get player report: %w", err)
	}
	return &report, nil
}

func (r *MerchantRepository) GetGameReport(ctx context.Context, merchantID, gameID, startDate, endDate string) (*model.GameReport, error) {
	var report model.GameReport

	query := `SELECT COALESCE(SUM(total_bets), 0), COALESCE(SUM(total_wins), 0), COUNT(DISTINCT player_id), COUNT(*) FROM game_reports WHERE merchant_id = $1 AND game_id = $2`
	args := []interface{}{merchantID, gameID}
	argNum := 3

	if startDate != "" {
		query += fmt.Sprintf(" AND created_at >= $%d", argNum)
		args = append(args, startDate)
		argNum++
	}
	if endDate != "" {
		query += fmt.Sprintf(" AND created_at <= $%d", argNum)
		args = append(args, endDate)
	}

	err := r.db.QueryRow(ctx, query, args...).Scan(&report.TotalBets, &report.TotalWins, &report.TotalPlayers, &report.Plays)
	if err != nil {
		return nil, fmt.Errorf("failed to get game report: %w", err)
	}
	return &report, nil
}
