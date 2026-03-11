package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/game-engine/loyalty-service/internal/config"
	"github.com/game-engine/loyalty-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type LoyaltyRepository struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

func NewLoyaltyRepository(db *pgxpool.Pool, redis *redis.Client) *LoyaltyRepository {
	return &LoyaltyRepository{db: db, redis: redis}
}

func NewPostgresDB(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connStr := cfg.ConnectionString()
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
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

func (r *LoyaltyRepository) BeginTx(ctx context.Context) (*pgxpool.Tx, error) {
	return r.db.Begin(ctx)
}

func (r *LoyaltyRepository) GetMember(ctx context.Context, userID string) (*model.Member, error) {
	var m model.Member
	err := r.db.QueryRow(ctx, `
		SELECT user_id, username, email, points, lifetime_points, tier, status, joined_at, updated_at
		FROM loyalty_members WHERE user_id = $1
	`, userID).Scan(&m.UserID, &m.Username, &m.Email, &m.Points, &m.LifetimePoints, &m.Tier, &m.Status, &m.JoinedAt, &m.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get member: %w", err)
	}
	return &m, nil
}

func (r *LoyaltyRepository) GetMemberTx(ctx context.Context, tx *pgxpool.Tx, userID string) (*model.Member, error) {
	var m model.Member
	err := tx.QueryRow(ctx, `
		SELECT user_id, username, email, points, lifetime_points, tier, status, joined_at, updated_at
		FROM loyalty_members WHERE user_id = $1
	`, userID).Scan(&m.UserID, &m.Username, &m.Email, &m.Points, &m.LifetimePoints, &m.Tier, &m.Status, &m.JoinedAt, &m.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get member: %w", err)
	}
	return &m, nil
}

func (r *LoyaltyRepository) CreateMember(ctx context.Context, member *model.Member) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO loyalty_members (user_id, username, email, points, lifetime_points, tier, status, joined_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
	`, member.UserID, member.Username, member.Email, member.Points, member.LifetimePoints, member.Tier, member.Status)
	return err
}

func (r *LoyaltyRepository) UpdateMemberPointsTx(ctx context.Context, tx *pgxpool.Tx, userID string, points, lifetimePoints int, tier string) error {
	_, err := tx.Exec(ctx, `
		UPDATE loyalty_members SET points = $1, lifetime_points = $2, tier = $3, updated_at = NOW()
		WHERE user_id = $4
	`, points, lifetimePoints, tier, userID)
	return err
}

func (r *LoyaltyRepository) AddPointsTransactionTx(ctx context.Context, tx *pgxpool.Tx, txRecord *model.PointsTransaction) error {
	_, err := tx.Exec(ctx, `
		INSERT INTO loyalty_points_transactions (transaction_id, user_id, amount, type, source, reference_id, description, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`, txRecord.TransactionID, txRecord.UserID, txRecord.Amount, txRecord.Type, txRecord.Source, txRecord.ReferenceID, txRecord.Description)
	return err
}

func (r *LoyaltyRepository) AddPointsTransaction(ctx context.Context, tx *model.PointsTransaction) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO loyalty_points_transactions (transaction_id, user_id, amount, type, source, reference_id, description, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`, tx.TransactionID, tx.UserID, tx.Amount, tx.Type, tx.Source, tx.ReferenceID, tx.Description)
	return err
}

func (r *LoyaltyRepository) GetPointsHistory(ctx context.Context, userID string, limit, offset int) ([]model.PointsTransaction, int, error) {
	var total int
	if err := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM loyalty_points_transactions WHERE user_id = $1`, userID).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("failed to count transactions: %w", err)
	}

	rows, err := r.db.Query(ctx, `
		SELECT transaction_id, user_id, amount, type, source, reference_id, description, created_at
		FROM loyalty_points_transactions WHERE user_id = $1
		ORDER BY created_at DESC LIMIT $2 OFFSET $3
	`, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var txs []model.PointsTransaction
	for rows.Next() {
		var tx model.PointsTransaction
		if err := rows.Scan(&tx.TransactionID, &tx.UserID, &tx.Amount, &tx.Type, &tx.Source, &tx.ReferenceID, &tx.Description, &tx.CreatedAt); err != nil {
			return nil, 0, err
		}
		txs = append(txs, tx)
	}
	return txs, total, nil
}

func (r *LoyaltyRepository) GetTiers(ctx context.Context) ([]model.Tier, error) {
	rows, err := r.db.Query(ctx, `
		SELECT tier_id, name, min_points, max_points, points_multiplier, cashback_percent, benefits
		FROM loyalty_tiers ORDER BY min_points ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tiers []model.Tier
	for rows.Next() {
		var t model.Tier
		if err := rows.Scan(&t.TierID, &t.Name, &t.MinPoints, &t.MaxPoints, &t.PointsMultiplier, &t.CashbackPercent, &t.Benefits); err != nil {
			return nil, err
		}
		tiers = append(tiers, t)
	}
	return tiers, nil
}

func (r *LoyaltyRepository) GetRewards(ctx context.Context) ([]model.Reward, error) {
	rows, err := r.db.Query(ctx, `
		SELECT reward_id, name, description, points_cost, type, value, status, expires_at, created_at
		FROM loyalty_rewards WHERE status = 'active'
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rewards []model.Reward
	for rows.Next() {
		var rwd model.Reward
		if err := rows.Scan(&rwd.RewardID, &rwd.Name, &rwd.Description, &rwd.PointsCost, &rwd.Type, &rwd.Value, &rwd.Status, &rwd.ExpiresAt, &rwd.CreatedAt); err != nil {
			return nil, err
		}
		rewards = append(rewards, rwd)
	}
	return rewards, nil
}

func (r *LoyaltyRepository) RedeemReward(ctx context.Context, redemption *model.Redemption) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO loyalty_redemptions (redemption_id, user_id, reward_id, points_spent, status, redeemed_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`, redemption.RedemptionID, redemption.UserID, redemption.RewardID, redemption.PointsSpent, redemption.Status)
	return err
}

func (r *LoyaltyRepository) GetTopMembers(ctx context.Context, limit int) ([]model.Member, error) {
	rows, err := r.db.Query(ctx, `
		SELECT user_id, username, email, points, lifetime_points, tier, status, joined_at, updated_at
		FROM loyalty_members ORDER BY lifetime_points DESC LIMIT $1
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []model.Member
	for rows.Next() {
		var m model.Member
		if err := rows.Scan(&m.UserID, &m.Username, &m.Email, &m.Points, &m.LifetimePoints, &m.Tier, &m.Status, &m.JoinedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	return members, nil
}

func (r *LoyaltyRepository) CacheMember(ctx context.Context, member *model.Member) error {
	key := fmt.Sprintf("loyalty:member:%s", member.UserID)
	data, err := json.Marshal(member)
	if err != nil {
		return err
	}
	return r.redis.Set(ctx, key, data, 3600).Err()
}

func (r *LoyaltyRepository) GetCachedMember(ctx context.Context, userID string) (*model.Member, error) {
	key := fmt.Sprintf("loyalty:member:%s", userID)
	data, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var member model.Member
	if err := json.Unmarshal([]byte(data), &member); err != nil {
		return nil, err
	}
	return &member, nil
}
