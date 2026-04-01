CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE commission_configs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    commission_type VARCHAR(30) NOT NULL,
    rate DECIMAL(5, 4) NOT NULL,
    min_amount DECIMAL(15, 2) DEFAULT 0.00,
    max_amount DECIMAL(15, 2),
    tier VARCHAR(20) NOT NULL DEFAULT 'standard',
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    effective_from TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    effective_to TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_commission_configs_commission_type ON commission_configs (commission_type);
CREATE INDEX idx_commission_configs_status ON commission_configs (status);
CREATE INDEX idx_commission_configs_tier ON commission_configs (tier);
CREATE INDEX idx_commission_configs_effective_from ON commission_configs (effective_from);

CREATE TABLE commission_claims (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    config_id UUID REFERENCES commission_configs (id),
    user_id UUID NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    source_type VARCHAR(50) NOT NULL,
    source_id UUID,
    period_start TIMESTAMP WITH TIME ZONE NOT NULL,
    period_end TIMESTAMP WITH TIME ZONE NOT NULL,
    calculated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_commission_claims_config_id ON commission_claims (config_id);
CREATE INDEX idx_commission_claims_user_id ON commission_claims (user_id);
CREATE INDEX idx_commission_claims_status ON commission_claims (status);
CREATE INDEX idx_commission_claims_period_start ON commission_claims (period_start);
CREATE INDEX idx_commission_claims_period_end ON commission_claims (period_end);

CREATE TABLE commission_settlements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    claim_id UUID NOT NULL REFERENCES commission_claims (id),
    user_id UUID NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    payment_method VARCHAR(50),
    payment_reference VARCHAR(255),
    processed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_commission_settlements_claim_id ON commission_settlements (claim_id);
CREATE INDEX idx_commission_settlements_user_id ON commission_settlements (user_id);
CREATE INDEX idx_commission_settlements_status ON commission_settlements (status);
CREATE INDEX idx_commission_settlements_created_at ON commission_settlements (created_at);
