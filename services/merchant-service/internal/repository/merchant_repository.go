package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/game_engine/merchant-service/internal/config"
	"github.com/game_engine/merchant-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type MerchantRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewMerchantRepository(db *pgxpool.Pool, redis *redis.Client) *MerchantRepository {
	return &MerchantRepository{db: db, redis: redis}
}

func NewPostgresDB(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := cfg.ConnectionString()
	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return pool, nil
}

func NewRedisClient(cfg config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	return client, nil
}

// Player operations
func (r *MerchantRepository) ListPlayers(ctx context.Context, merchantID string, page, limit int, search string) ([]model.MerchantPlayer, int, error) {
	offset := (page - 1) * limit

	var total int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM merchant_players WHERE merchant_id = $1`, merchantID).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count players: %w", err)
	}

	rows, err := r.db.Query(ctx, `SELECT player_id, username, email, status FROM merchant_players WHERE merchant_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, merchantID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list players: %w", err)
	}
	defer rows.Close()

	var players []model.MerchantPlayer
	for rows.Next() {
		var p model.MerchantPlayer
		if err := rows.Scan(&p.PlayerID, &p.Username, &p.Email, &p.Status); err != nil {
			return nil, 0, fmt.Errorf("failed to scan player: %w", err)
		}
		players = append(players, p)
	}
	return players, total, nil
}

func (r *MerchantRepository) GetPlayer(ctx context.Context, merchantID, playerID string) (*model.MerchantPlayer, error) {
	var p model.MerchantPlayer
	err := r.db.QueryRow(ctx, `SELECT player_id, username, email, status FROM merchant_players WHERE merchant_id = $1 AND player_id = $2`, merchantID, playerID).Scan(&p.PlayerID, &p.Username, &p.Email, &p.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to get player: %w", err)
	}
	return &p, nil
}

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

// Config
func (r *MerchantRepository) GetConfig(ctx context.Context, merchantID string) (map[string]string, error) {
	rows, err := r.db.Query(ctx, `SELECT key, value FROM merchant_config WHERE merchant_id = $1`, merchantID)
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %w", err)
	}
	defer rows.Close()

	config := make(map[string]string)
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, fmt.Errorf("failed to scan config: %w", err)
		}
		config[key] = value
	}
	return config, nil
}

func (r *MerchantRepository) UpdateConfig(ctx context.Context, merchantID string, config map[string]string) error {
	for key, value := range config {
		_, err := r.db.Exec(ctx, `INSERT INTO merchant_config (merchant_id, key, value) VALUES ($1, $2, $3) ON CONFLICT (merchant_id, key) DO UPDATE SET value = $3`, merchantID, key, value)
		if err != nil {
			return fmt.Errorf("failed to update config: %w", err)
		}
	}
	return nil
}

// Webhooks
func (r *MerchantRepository) RegisterWebhook(ctx context.Context, merchantID, url, events string) (string, error) {
	var webhookID string
	err := r.db.QueryRow(ctx, `INSERT INTO merchant_webhooks (merchant_id, url, events) VALUES ($1, $2, $3) RETURNING webhook_id`, merchantID, url, events).Scan(&webhookID)
	if err != nil {
		return "", fmt.Errorf("failed to register webhook: %w", err)
	}
	return webhookID, nil
}

func (r *MerchantRepository) ListWebhooks(ctx context.Context, merchantID string) ([]model.Webhook, error) {
	rows, err := r.db.Query(ctx, `SELECT webhook_id, url, events, status FROM merchant_webhooks WHERE merchant_id = $1`, merchantID)
	if err != nil {
		return nil, fmt.Errorf("failed to list webhooks: %w", err)
	}
	defer rows.Close()

	var webhooks []model.Webhook
	for rows.Next() {
		var w model.Webhook
		if err := rows.Scan(&w.WebhookID, &w.URL, &w.Events, &w.Status); err != nil {
			return nil, fmt.Errorf("failed to scan webhook: %w", err)
		}
		webhooks = append(webhooks, w)
	}
	return webhooks, nil
}

func (r *MerchantRepository) DeleteWebhook(ctx context.Context, merchantID, webhookID string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM merchant_webhooks WHERE merchant_id = $1 AND webhook_id = $2`, merchantID, webhookID)
	return err
}

// Agents
func (r *MerchantRepository) ListAgents(ctx context.Context, merchantID string, page, limit int) ([]model.Agent, int, error) {
	offset := (page - 1) * limit

	var total int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM merchant_agents WHERE merchant_id = $1`, merchantID).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count agents: %w", err)
	}

	rows, err := r.db.Query(ctx, `SELECT agent_id, username, email, status FROM merchant_agents WHERE merchant_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, merchantID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list agents: %w", err)
	}
	defer rows.Close()

	var agents []model.Agent
	for rows.Next() {
		var a model.Agent
		if err := rows.Scan(&a.AgentID, &a.Username, &a.Email, &a.Status); err != nil {
			return nil, 0, fmt.Errorf("failed to scan agent: %w", err)
		}
		agents = append(agents, a)
	}
	return agents, total, nil
}

func (r *MerchantRepository) GetAgent(ctx context.Context, merchantID, agentID string) (*model.Agent, error) {
	var a model.Agent
	err := r.db.QueryRow(ctx, `SELECT agent_id, username, email, status FROM merchant_agents WHERE merchant_id = $1 AND agent_id = $2`, merchantID, agentID).Scan(&a.AgentID, &a.Username, &a.Email, &a.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to get agent: %w", err)
	}
	return &a, nil
}

func (r *MerchantRepository) CreateAgent(ctx context.Context, merchantID, username, email string, sendInvitation bool) (string, error) {
	var agentID string
	status := "pending"
	// If not sending invitation, agent is active immediately
	if !sendInvitation {
		status = "active"
	}
	err := r.db.QueryRow(ctx, `INSERT INTO merchant_agents (merchant_id, username, email, status) VALUES ($1, $2, $3, $4) RETURNING agent_id`, merchantID, username, email, status).Scan(&agentID)
	if err != nil {
		return "", fmt.Errorf("failed to create agent: %w", err)
	}

	// If sendInvitation is true, queue the invitation for processing
	if sendInvitation {
		if err := r.queueAgentInvitation(ctx, merchantID, agentID, username, email); err != nil {
			// Log the error but don't fail the agent creation
			// The agent can be invited later manually
			log.Printf("failed to queue invitation for agent %s: %v", agentID, err)
		}
	}

	return agentID, nil
}

// queueAgentInvitation creates an invitation record that can be processed by a notification worker
func (r *MerchantRepository) queueAgentInvitation(ctx context.Context, merchantID, agentID, username, email string) error {
	// Create invitation record with pending status
	var invitationID string
	err := r.db.QueryRow(ctx, `
		INSERT INTO merchant_agent_invitations (merchant_id, agent_id, email, username, status, created_at)
		VALUES ($1, $2, $3, $4, 'pending', NOW())
		RETURNING invitation_id
	`, merchantID, agentID, email, username).Scan(&invitationID)
	if err != nil {
		return fmt.Errorf("failed to create invitation record: %w", err)
	}

	// Publish to Redis queue for async processing
	invitationData := fmt.Sprintf("{\"invitation_id\":\"%s\",\"agent_id\":\"%s\",\"email\":\"%s\",\"username\":\"%s\"}",
		invitationID, agentID, email, username)
	if err := r.redis.Publish(ctx, "agent-invitations", invitationData).Err(); err != nil {
		return fmt.Errorf("failed to publish invitation to queue: %w", err)
	}

	// Set TTL for invitation (24 hours)
	if err := r.redis.Set(ctx, fmt.Sprintf("invitation:%s", invitationID),
		"pending", 24*time.Hour).Err(); err != nil {
		log.Printf("warning: failed to set invitation TTL: %v", err)
	}

	return nil
}

func (r *MerchantRepository) UpdateAgent(ctx context.Context, merchantID, agentID, username, email string) error {
	_, err := r.db.Exec(ctx, `UPDATE merchant_agents SET username = $1, email = $2, updated_at = NOW() WHERE merchant_id = $3 AND agent_id = $4`, username, email, merchantID, agentID)
	if err != nil {
		return fmt.Errorf("failed to update agent: %w", err)
	}
	return nil
}

func (r *MerchantRepository) UpdateAgentStatus(ctx context.Context, merchantID, agentID, status string) error {
	_, err := r.db.Exec(ctx, `UPDATE merchant_agents SET status = $1, updated_at = NOW() WHERE merchant_id = $2 AND agent_id = $3`, status, merchantID, agentID)
	if err != nil {
		return fmt.Errorf("failed to update agent status: %w", err)
	}
	return nil
}
