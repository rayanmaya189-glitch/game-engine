package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/game-engine/agent-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type AgentRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewAgentRepository(db *pgxpool.Pool, redis *redis.Client) *AgentRepository {
	return &AgentRepository{
		db:    db,
		redis: redis,
	}
}

// Player operations

func (r *AgentRepository) ListPlayers(ctx context.Context, agentID string, page, limit int, search, status string) ([]model.Player, int, error) {
	offset := (page - 1) * limit

	// Build query
	query := `
		SELECT player_id, username, email, status, total_deposits, total_bets, balance, created_at
		FROM agent_players
		WHERE agent_id = $1
	`
	countQuery := `SELECT COUNT(*) FROM agent_players WHERE agent_id = $1`
	args := []interface{}{agentID}
	argNum := 2

	if search != "" {
		query += fmt.Sprintf(" AND (username ILIKE $%d OR email ILIKE $%d)", argNum, argNum+1)
		countQuery += fmt.Sprintf(" AND (username ILIKE $%d OR email ILIKE $%d)", argNum, argNum+1)
		args = append(args, "%"+search+"%", "%"+search+"%")
		argNum += 2
	}

	if status != "" {
		query += fmt.Sprintf(" AND status = $%d", argNum)
		countQuery += fmt.Sprintf(" AND status = $%d", argNum)
		args = append(args, status)
		argNum++
	}

	// Get total count
	var total int
	err := r.db.QueryRow(ctx, countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count players: %w", err)
	}

	// Get players
	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", argNum, argNum+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list players: %w", err)
	}
	defer rows.Close()

	var players []model.Player
	for rows.Next() {
		var p model.Player
		if err := rows.Scan(
			&p.PlayerID,
			&p.Username,
			&p.Email,
			&p.Status,
			&p.TotalDeposits,
			&p.TotalBets,
			&p.Balance,
			&p.CreatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan player: %w", err)
		}
		players = append(players, p)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating players: %w", err)
	}

	return players, total, nil
}

func (r *AgentRepository) GetPlayer(ctx context.Context, agentID, playerID string) (*model.Player, error) {
	// Try cache first
	cacheKey := fmt.Sprintf("agent:%s:player:%s", agentID, playerID)
	cached, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var player model.Player
		if json.Unmarshal([]byte(cached), &player) == nil {
			return &player, nil
		}
	}

	// Get from database
	var p model.Player
	err = r.db.QueryRow(ctx, `
		SELECT player_id, username, email, status, total_deposits, total_bets, balance, created_at
		FROM agent_players
		WHERE agent_id = $1 AND player_id = $2
	`, agentID, playerID).Scan(
		&p.PlayerID,
		&p.Username,
		&p.Email,
		&p.Status,
		&p.TotalDeposits,
		&p.TotalBets,
		&p.Balance,
		&p.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get player: %w", err)
	}

	// Cache for 5 minutes
	if data, err := json.Marshal(p); err == nil {
		if err := r.redis.Set(ctx, cacheKey, data, 5*time.Minute).Err(); err != nil {
			// Log error - cache failures shouldn't break the flow
			log.Printf("failed to cache player: %v", err)
		}
	}

	return &p, nil
}

func (r *AgentRepository) UpdatePlayerLimit(ctx context.Context, agentID, playerID string, depositLimit, betLimit float64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE agent_players
		SET deposit_limit = $1, bet_limit = $2, updated_at = NOW()
		WHERE agent_id = $3 AND player_id = $4
	`, depositLimit, betLimit, agentID, playerID)
	if err != nil {
		return fmt.Errorf("failed to update player limit: %w", err)
	}

	// Invalidate cache
	cacheKey := fmt.Sprintf("agent:%s:player:%s", agentID, playerID)
	r.redis.Del(ctx, cacheKey)

	return nil
}

// Dashboard

func (r *AgentRepository) GetDashboard(ctx context.Context, agentID string) (*model.Dashboard, error) {
	var dashboard model.Dashboard

	err := r.db.QueryRow(ctx, `
		SELECT 
			COUNT(*) as total_players,
			COUNT(*) FILTER (WHERE status = 'active') as active_players,
			COALESCE(SUM(total_commission), 0) as total_commission,
			COALESCE(SUM(pending_commission), 0) as pending_commission
		FROM agent_dashboard
		WHERE agent_id = $1
	`, agentID).Scan(
		&dashboard.TotalPlayers,
		&dashboard.ActivePlayers,
		&dashboard.TotalCommission,
		&dashboard.PendingCommission,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get dashboard: %w", err)
	}

	return &dashboard, nil
}

// Commission operations

func (r *AgentRepository) GetCommissions(ctx context.Context, agentID, startDate, endDate string) ([]model.Commission, float64, error) {
	query := `
		SELECT commission_id, amount, type, status, created_at
		FROM agent_commissions
		WHERE agent_id = $1
	`
	args := []interface{}{agentID}
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

	query += " ORDER BY created_at DESC"

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get commissions: %w", err)
	}
	defer rows.Close()

	var commissions []model.Commission
	var total float64
	for rows.Next() {
		var c model.Commission
		if err := rows.Scan(&c.CommissionID, &c.Amount, &c.Type, &c.Status, &c.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("failed to scan commission: %w", err)
		}
		commissions = append(commissions, c)
		total += c.Amount
	}

	return commissions, total, nil
}

func (r *AgentRepository) GetPendingCommissions(ctx context.Context, agentID string) (float64, error) {
	var pending float64
	err := r.db.QueryRow(ctx, `
		SELECT COALESCE(SUM(amount), 0)
		FROM agent_commissions
		WHERE agent_id = $1 AND status = 'pending'
	`, agentID).Scan(&pending)
	if err != nil {
		return 0, fmt.Errorf("failed to get pending commissions: %w", err)
	}
	return pending, nil
}

func (r *AgentRepository) GetCommissionHistory(ctx context.Context, agentID string, page, limit int) ([]model.Commission, int, error) {
	offset := (page - 1) * limit

	var total int
	err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM agent_commissions WHERE agent_id = $1`, agentID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count commissions: %w", err)
	}

	rows, err := r.db.Query(ctx, `
		SELECT commission_id, amount, type, status, created_at
		FROM agent_commissions
		WHERE agent_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`, agentID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get commission history: %w", err)
	}
	defer rows.Close()

	var commissions []model.Commission
	for rows.Next() {
		var c model.Commission
		if err := rows.Scan(&c.CommissionID, &c.Amount, &c.Type, &c.Status, &c.CreatedAt); err != nil {
			return nil, 0, fmt.Errorf("failed to scan commission: %w", err)
		}
		commissions = append(commissions, c)
	}

	return commissions, total, nil
}

func (r *AgentRepository) ClaimCommission(ctx context.Context, agentID, commissionID string) (string, error) {
	var transactionID string
	err := r.db.QueryRow(ctx, `
		UPDATE agent_commissions
		SET status = 'claimed', claimed_at = NOW()
		WHERE agent_id = $1 AND commission_id = $2 AND status = 'pending'
		RETURNING transaction_id
	`, agentID, commissionID).Scan(&transactionID)
	if err != nil {
		return "", fmt.Errorf("failed to claim commission: %w", err)
	}
	return transactionID, nil
}
