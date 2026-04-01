CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE affiliates (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    affiliate_code VARCHAR(50) NOT NULL UNIQUE,
    commission_rate DECIMAL(5, 4) NOT NULL DEFAULT 0.0000,
    tier VARCHAR(20) NOT NULL DEFAULT 'standard',
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    total_earnings DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    total_referrals INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_affiliates_user_id ON affiliates (user_id);
CREATE INDEX idx_affiliates_affiliate_code ON affiliates (affiliate_code);
CREATE INDEX idx_affiliates_status ON affiliates (status);

CREATE TABLE referrals (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    affiliate_id UUID NOT NULL REFERENCES affiliates (id),
    referred_user_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    source VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_referrals_affiliate_id ON referrals (affiliate_id);
CREATE INDEX idx_referrals_referred_user_id ON referrals (referred_user_id);
CREATE INDEX idx_referrals_status ON referrals (status);

CREATE TABLE affiliate_commissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    affiliate_id UUID NOT NULL REFERENCES affiliates (id),
    referral_id UUID REFERENCES referrals (id),
    amount DECIMAL(15, 2) NOT NULL,
    commission_type VARCHAR(30) NOT NULL,
    source_type VARCHAR(50),
    source_id UUID,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_affiliate_commissions_affiliate_id ON affiliate_commissions (affiliate_id);
CREATE INDEX idx_affiliate_commissions_status ON affiliate_commissions (status);
CREATE INDEX idx_affiliate_commissions_created_at ON affiliate_commissions (created_at);

CREATE TABLE affiliate_payouts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    affiliate_id UUID NOT NULL REFERENCES affiliates (id),
    amount DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    payment_method VARCHAR(50),
    payment_reference VARCHAR(255),
    processed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_affiliate_payouts_affiliate_id ON affiliate_payouts (affiliate_id);
CREATE INDEX idx_affiliate_payouts_status ON affiliate_payouts (status);
CREATE INDEX idx_affiliate_payouts_created_at ON affiliate_payouts (created_at);
