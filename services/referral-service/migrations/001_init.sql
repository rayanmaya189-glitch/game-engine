CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS referrals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    referrer_id VARCHAR(255) NOT NULL,
    referee_id VARCHAR(255),
    referral_code VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    reward_amount DECIMAL(15, 2) NOT NULL DEFAULT 0,
    reward_type VARCHAR(50) NOT NULL DEFAULT 'CASH',
    reward_claimed BOOLEAN NOT NULL DEFAULT false,
    source VARCHAR(100),
    campaign_id VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    qualified_at TIMESTAMPTZ,
    rewarded_at TIMESTAMPTZ,
    claimed_at TIMESTAMPTZ
);

CREATE INDEX idx_referrals_referrer ON referrals(referrer_id);
CREATE INDEX idx_referrals_referee ON referrals(referee_id);
CREATE INDEX idx_referrals_code ON referrals(referral_code);
CREATE INDEX idx_referrals_status ON referrals(status);
CREATE INDEX idx_referrals_created ON referrals(created_at DESC);

CREATE TABLE IF NOT EXISTS referral_codes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id VARCHAR(255) NOT NULL UNIQUE,
    code VARCHAR(50) NOT NULL UNIQUE,
    referral_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    click_count BIGINT NOT NULL DEFAULT 0,
    signup_count BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX idx_referral_codes_player ON referral_codes(player_id);
CREATE INDEX idx_referral_codes_code ON referral_codes(code);

CREATE TABLE IF NOT EXISTS referral_rewards (
    id SERIAL PRIMARY KEY,
    reward_type VARCHAR(50) NOT NULL,
    referrer_bonus DECIMAL(15, 2) NOT NULL DEFAULT 0,
    referee_bonus DECIMAL(15, 2) NOT NULL DEFAULT 0,
    min_deposit DECIMAL(15, 2) NOT NULL DEFAULT 0,
    min_bet DECIMAL(15, 2) NOT NULL DEFAULT 0,
    active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_referral_rewards_active ON referral_rewards(active);
CREATE INDEX idx_referral_rewards_type ON referral_rewards(reward_type);
