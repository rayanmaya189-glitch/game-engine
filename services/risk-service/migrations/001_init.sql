CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE risk_profiles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    overall_risk_level VARCHAR(20) NOT NULL DEFAULT 'low',
    risk_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    kyc_status VARCHAR(20) DEFAULT 'unverified',
    country_risk VARCHAR(20) DEFAULT 'low',
    device_risk VARCHAR(20) DEFAULT 'low',
    behavioral_risk VARCHAR(20) DEFAULT 'low',
    transaction_risk VARCHAR(20) DEFAULT 'low',
    is_blacklisted BOOLEAN NOT NULL DEFAULT FALSE,
    blacklist_reason TEXT,
    notes TEXT,
    last_assessed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_risk_profiles_user_id ON risk_profiles (user_id);
CREATE INDEX idx_risk_profiles_overall_risk_level ON risk_profiles (overall_risk_level);
CREATE INDEX idx_risk_profiles_risk_score ON risk_profiles (risk_score);
CREATE INDEX idx_risk_profiles_is_blacklisted ON risk_profiles (is_blacklisted);

CREATE TABLE risk_assessments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    profile_id UUID REFERENCES risk_profiles (id),
    assessment_type VARCHAR(50) NOT NULL,
    trigger_event VARCHAR(50) NOT NULL,
    risk_level VARCHAR(20) NOT NULL,
    risk_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    factors JSONB,
    recommendations TEXT[],
    assessed_by VARCHAR(30) DEFAULT 'system',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_risk_assessments_user_id ON risk_assessments (user_id);
CREATE INDEX idx_risk_assessments_profile_id ON risk_assessments (profile_id);
CREATE INDEX idx_risk_assessments_assessment_type ON risk_assessments (assessment_type);
CREATE INDEX idx_risk_assessments_risk_level ON risk_assessments (risk_level);
CREATE INDEX idx_risk_assessments_created_at ON risk_assessments (created_at);

CREATE TABLE risk_limits (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    profile_id UUID REFERENCES risk_profiles (id),
    limit_type VARCHAR(30) NOT NULL,
    limit_value DECIMAL(15, 2) NOT NULL,
    period VARCHAR(20) NOT NULL DEFAULT 'daily',
    current_usage DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    effective_from TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    effective_to TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_risk_limits_user_id ON risk_limits (user_id);
CREATE INDEX idx_risk_limits_profile_id ON risk_limits (profile_id);
CREATE INDEX idx_risk_limits_limit_type ON risk_limits (limit_type);
CREATE INDEX idx_risk_limits_status ON risk_limits (status);
CREATE INDEX idx_risk_limits_effective_from ON risk_limits (effective_from);
