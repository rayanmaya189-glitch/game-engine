package repository

import (
	"github.com/game_engine/referral-service/internal/model"
)

func (r *ReferralRepository) TrackReferralClick(code string) error {
	_, err := r.db.Exec("UPDATE referral_codes SET click_count = click_count + 1 WHERE code = $1", code)
	return err
}

func (r *ReferralRepository) TrackReferralSignup(code string) error {
	_, err := r.db.Exec("UPDATE referral_codes SET signup_count = signup_count + 1 WHERE code = $1", code)
	return err
}

func (r *ReferralRepository) CreateReferralCode(rc *model.ReferralCode) error {
	query := `INSERT INTO referral_codes (player_id, code, referral_url, created_at, click_count, signup_count)
		VALUES ($1, $2, $3, $4, 0, 0)`
	_, err := r.db.Exec(query, rc.PlayerID, rc.Code, rc.ReferralURL, rc.CreatedAt)
	return err
}

func (r *ReferralRepository) GetReferralCodeByPlayer(playerID string) (*model.ReferralCode, error) {
	query := `SELECT player_id, code, referral_url, created_at, click_count, signup_count
		FROM referral_codes WHERE player_id = $1`

	var rc model.ReferralCode
	err := r.db.QueryRow(query, playerID).Scan(
		&rc.PlayerID, &rc.Code, &rc.ReferralURL, &rc.CreatedAt, &rc.ClickCount, &rc.SignupCount,
	)
	if err != nil {
		return nil, err
	}
	return &rc, nil
}

func (r *ReferralRepository) GetReferrerByCode(code string) (string, error) {
	var referrerID string
	err := r.db.QueryRow("SELECT player_id FROM referral_codes WHERE code = $1", code).Scan(&referrerID)
	return referrerID, err
}

func (r *ReferralRepository) GetChildReferrals(referrerID string) ([]*model.Referral, error) {
	query := `SELECT id, referrer_id, referee_id, referral_code, status, reward_amount, reward_type,
		reward_claimed, source, created_at FROM referrals WHERE referrer_id = $1 AND status = 'REWARDED'`

	rows, err := r.db.Query(query, referrerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var referrals []*model.Referral
	for rows.Next() {
		var ref model.Referral
		if err := rows.Scan(
			&ref.ID, &ref.ReferrerID, &ref.RefereeID, &ref.ReferralCode, &ref.Status,
			&ref.RewardAmount, &ref.RewardType, &ref.RewardClaimed, &ref.Source, &ref.CreatedAt,
		); err != nil {
			return nil, err
		}
		referrals = append(referrals, &ref)
	}
	return referrals, nil
}
