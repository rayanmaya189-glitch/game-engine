package repository

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/game_engine/referral-service/internal/config"
	"github.com/game_engine/referral-service/internal/model"
)

type ReferralRepository struct {
	db *sql.DB
}

func NewPostgresDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxConns)
	db.SetMaxIdleConns(cfg.MaxConns / 2)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewReferralRepository(db *sql.DB) *ReferralRepository {
	return &ReferralRepository{db: db}
}

func (r *ReferralRepository) CreateReferral(ref *model.Referral) error {
	query := `INSERT INTO referrals (id, referrer_id, referee_id, referral_code, status, reward_amount,
		reward_type, reward_claimed, source, campaign_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := r.db.Exec(query,
		ref.ID, ref.ReferrerID, ref.RefereeID, ref.ReferralCode, ref.Status,
		ref.RewardAmount, ref.RewardType, ref.RewardClaimed, ref.Source, ref.CampaignID, ref.CreatedAt,
	)
	return err
}

func (r *ReferralRepository) GetReferralByID(id string) (*model.Referral, error) {
	query := `SELECT id, referrer_id, referee_id, referral_code, status, reward_amount, reward_type,
		reward_claimed, source, campaign_id, created_at, qualified_at, rewarded_at, claimed_at
		FROM referrals WHERE id = $1`

	var ref model.Referral
	var qualifiedAt, rewardedAt, claimedAt sql.NullTime

	err := r.db.QueryRow(query, id).Scan(
		&ref.ID, &ref.ReferrerID, &ref.RefereeID, &ref.ReferralCode, &ref.Status,
		&ref.RewardAmount, &ref.RewardType, &ref.RewardClaimed, &ref.Source, &ref.CampaignID,
		&ref.CreatedAt, &qualifiedAt, &rewardedAt, &claimedAt,
	)
	if err != nil {
		return nil, err
	}

	if qualifiedAt.Valid {
		ref.QualifiedAt = &qualifiedAt.Time
	}
	if rewardedAt.Valid {
		ref.RewardedAt = &rewardedAt.Time
	}
	if claimedAt.Valid {
		ref.ClaimedAt = &claimedAt.Time
	}

	return &ref, nil
}

func (r *ReferralRepository) GetReferralByCode(code string) (*model.Referral, error) {
	query := `SELECT id, referrer_id, referee_id, referral_code, status, reward_amount, reward_type,
		reward_claimed, source, campaign_id, created_at, qualified_at, rewarded_at, claimed_at
		FROM referrals WHERE referral_code = $1`

	var ref model.Referral
	var qualifiedAt, rewardedAt, claimedAt sql.NullTime

	err := r.db.QueryRow(query, code).Scan(
		&ref.ID, &ref.ReferrerID, &ref.RefereeID, &ref.ReferralCode, &ref.Status,
		&ref.RewardAmount, &ref.RewardType, &ref.RewardClaimed, &ref.Source, &ref.CampaignID,
		&ref.CreatedAt, &qualifiedAt, &rewardedAt, &claimedAt,
	)
	if err != nil {
		return nil, err
	}

	if qualifiedAt.Valid {
		ref.QualifiedAt = &qualifiedAt.Time
	}
	if rewardedAt.Valid {
		ref.RewardedAt = &rewardedAt.Time
	}
	if claimedAt.Valid {
		ref.ClaimedAt = &claimedAt.Time
	}

	return &ref, nil
}

func (r *ReferralRepository) GetReferralsByReferrer(referrerID string, filter *model.ReferralFilter) (*model.ReferralList, error) {
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PageSize < 1 {
		filter.PageSize = 20
	}
	offset := (filter.Page - 1) * filter.PageSize

	var total int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM referrals WHERE referrer_id = $1", referrerID).Scan(&total)
	if err != nil {
		return nil, err
	}

	query := `SELECT id, referrer_id, referee_id, referral_code, status, reward_amount, reward_type,
		reward_claimed, source, campaign_id, created_at, qualified_at, rewarded_at, claimed_at
		FROM referrals WHERE referrer_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`

	rows, err := r.db.Query(query, referrerID, filter.PageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var referrals []*model.Referral
	for rows.Next() {
		var ref model.Referral
		var qualifiedAt, rewardedAt, claimedAt sql.NullTime

		if err := rows.Scan(
			&ref.ID, &ref.ReferrerID, &ref.RefereeID, &ref.ReferralCode, &ref.Status,
			&ref.RewardAmount, &ref.RewardType, &ref.RewardClaimed, &ref.Source, &ref.CampaignID,
			&ref.CreatedAt, &qualifiedAt, &rewardedAt, &claimedAt,
		); err != nil {
			return nil, err
		}

		if qualifiedAt.Valid {
			ref.QualifiedAt = &qualifiedAt.Time
		}
		if rewardedAt.Valid {
			ref.RewardedAt = &rewardedAt.Time
		}
		if claimedAt.Valid {
			ref.ClaimedAt = &claimedAt.Time
		}

		referrals = append(referrals, &ref)
	}

	return &model.ReferralList{
		Referrals: referrals,
		Total:     total,
		Page:      filter.Page,
		PageSize:  filter.PageSize,
	}, nil
}

func (r *ReferralRepository) UpdateReferralStatus(id string, status model.ReferralStatus) error {
	_, err := r.db.Exec("UPDATE referrals SET status = $1 WHERE id = $2", status, id)
	return err
}

func (r *ReferralRepository) MarkQualified(id string) error {
	now := time.Now()
	_, err := r.db.Exec("UPDATE referrals SET status = $1, qualified_at = $2 WHERE id = $3",
		model.ReferralStatusQualified, now, id)
	return err
}

func (r *ReferralRepository) MarkRewarded(id string, amount float64, rewardType model.RewardType) error {
	now := time.Now()
	_, err := r.db.Exec("UPDATE referrals SET status = $1, reward_amount = $2, reward_type = $3, rewarded_at = $4 WHERE id = $5",
		model.ReferralStatusRewarded, amount, rewardType, now, id)
	return err
}

func (r *ReferralRepository) ClaimReward(id string) error {
	now := time.Now()
	_, err := r.db.Exec("UPDATE referrals SET reward_claimed = true, claimed_at = $1 WHERE id = $2", now, id)
	return err
}

func (r *ReferralRepository) GetReferralStats(referrerID string) (*model.ReferralStats, error) {
	stats := &model.ReferralStats{}

	err := r.db.QueryRow(`SELECT COUNT(*) FROM referrals WHERE referrer_id = $1`, referrerID).Scan(&stats.TotalReferrals)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`SELECT COUNT(*) FROM referrals WHERE referrer_id = $1 AND status = 'ACTIVE'`, referrerID).Scan(&stats.ActiveReferrals)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`SELECT COUNT(*) FROM referrals WHERE referrer_id = $1 AND status IN ('QUALIFIED', 'REWARDED')`, referrerID).Scan(&stats.QualifiedReferrals)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`SELECT COALESCE(SUM(reward_amount), 0) FROM referrals WHERE referrer_id = $1 AND status = 'REWARDED'`, referrerID).Scan(&stats.TotalRewardsEarned)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`SELECT COALESCE(SUM(reward_amount), 0) FROM referrals WHERE referrer_id = $1 AND reward_claimed = true`, referrerID).Scan(&stats.TotalRewardsPaid)
	if err != nil {
		return nil, err
	}

	stats.PendingRewards = stats.TotalRewardsEarned - stats.TotalRewardsPaid

	return stats, nil
}

func (r *ReferralRepository) CreateReferralReward(reward *model.ReferralReward) error {
	query := `INSERT INTO referral_rewards (reward_type, referrer_bonus, referee_bonus, min_deposit, min_bet, active)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`
	return r.db.QueryRow(query, reward.RewardType, reward.ReferrerBonus, reward.RefereeBonus,
		reward.MinDeposit, reward.MinBet, reward.Active).Scan(&reward.ID, &reward.CreatedAt)
}

func (r *ReferralRepository) GetActiveRewards() ([]*model.ReferralReward, error) {
	query := `SELECT id, reward_type, referrer_bonus, referee_bonus, min_deposit, min_bet, active, created_at
		FROM referral_rewards WHERE active = true ORDER BY referrer_bonus DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rewards []*model.ReferralReward
	for rows.Next() {
		var rw model.ReferralReward
		if err := rows.Scan(&rw.ID, &rw.RewardType, &rw.ReferrerBonus, &rw.RefereeBonus,
			&rw.MinDeposit, &rw.MinBet, &rw.Active, &rw.CreatedAt); err != nil {
			return nil, err
		}
		rewards = append(rewards, &rw)
	}
	return rewards, nil
}
