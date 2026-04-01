CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE bonuses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) UNIQUE,
    bonus_type VARCHAR(30) NOT NULL,
    amount DECIMAL(15, 2),
    percentage DECIMAL(5, 2),
    max_amount DECIMAL(15, 2),
    min_deposit DECIMAL(15, 2) DEFAULT 0.00,
    wagering_multiplier DECIMAL(5, 2) NOT NULL DEFAULT 1.00,
    wagering_deadline_days INTEGER,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    start_date TIMESTAMP WITH TIME ZONE,
    end_date TIMESTAMP WITH TIME ZONE,
    max_claims INTEGER,
    current_claims INTEGER NOT NULL DEFAULT 0,
    terms TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_bonuses_code ON bonuses (code);
CREATE INDEX idx_bonuses_bonus_type ON bonuses (bonus_type);
CREATE INDEX idx_bonuses_status ON bonuses (status);
CREATE INDEX idx_bonuses_start_date ON bonuses (start_date);
CREATE INDEX idx_bonuses_end_date ON bonuses (end_date);

CREATE TABLE bonus_claims (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    bonus_id UUID NOT NULL REFERENCES bonuses (id),
    user_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    bonus_amount DECIMAL(15, 2) NOT NULL,
    wagering_requirement DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    wagering_completed DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    expires_at TIMESTAMP WITH TIME ZONE,
    claimed_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_bonus_claims_bonus_id ON bonus_claims (bonus_id);
CREATE INDEX idx_bonus_claims_user_id ON bonus_claims (user_id);
CREATE INDEX idx_bonus_claims_status ON bonus_claims (status);
CREATE INDEX idx_bonus_claims_expires_at ON bonus_claims (expires_at);

CREATE TABLE bonus_wagering (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    claim_id UUID NOT NULL REFERENCES bonus_claims (id),
    user_id UUID NOT NULL,
    game_id UUID,
    amount_wagered DECIMAL(15, 2) NOT NULL,
    amount_counted DECIMAL(15, 2) NOT NULL,
    wager_type VARCHAR(30) NOT NULL DEFAULT 'standard',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_bonus_wagering_claim_id ON bonus_wagering (claim_id);
CREATE INDEX idx_bonus_wagering_user_id ON bonus_wagering (user_id);
CREATE INDEX idx_bonus_wagering_game_id ON bonus_wagering (game_id);
CREATE INDEX idx_bonus_wagering_created_at ON bonus_wagering (created_at);

CREATE TABLE bonus_credits (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    claim_id UUID NOT NULL REFERENCES bonus_claims (id),
    user_id UUID NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    credit_type VARCHAR(30) NOT NULL DEFAULT 'bonus',
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    expires_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_bonus_credits_claim_id ON bonus_credits (claim_id);
CREATE INDEX idx_bonus_credits_user_id ON bonus_credits (user_id);
CREATE INDEX idx_bonus_credits_status ON bonus_credits (status);
