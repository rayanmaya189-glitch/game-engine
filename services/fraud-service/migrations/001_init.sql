CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE fraud_device_fingerprints (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID,
    fingerprint_hash VARCHAR(255) NOT NULL,
    device_type VARCHAR(30),
    os_name VARCHAR(50),
    os_version VARCHAR(50),
    browser_name VARCHAR(50),
    browser_version VARCHAR(50),
    ip_address INET,
    country_code VARCHAR(2),
    metadata JSONB,
    first_seen_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    last_seen_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_fraud_device_fingerprints_user_id ON fraud_device_fingerprints (user_id);
CREATE INDEX idx_fraud_device_fingerprints_fingerprint_hash ON fraud_device_fingerprints (fingerprint_hash);
CREATE INDEX idx_fraud_device_fingerprints_ip_address ON fraud_device_fingerprints (ip_address);
CREATE INDEX idx_fraud_device_fingerprints_last_seen_at ON fraud_device_fingerprints (last_seen_at);

CREATE TABLE fraud_scores (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    overall_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    velocity_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    device_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    location_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    behavioral_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    risk_level VARCHAR(20) NOT NULL DEFAULT 'low',
    context_type VARCHAR(50),
    context_id UUID,
    details JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_fraud_scores_user_id ON fraud_scores (user_id);
CREATE INDEX idx_fraud_scores_risk_level ON fraud_scores (risk_level);
CREATE INDEX idx_fraud_scores_overall_score ON fraud_scores (overall_score);
CREATE INDEX idx_fraud_scores_created_at ON fraud_scores (created_at);

CREATE TABLE fraud_collusion_events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    event_type VARCHAR(50) NOT NULL,
    session_id UUID,
    game_id UUID,
    user_ids UUID[] NOT NULL,
    confidence_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    status VARCHAR(20) NOT NULL DEFAULT 'detected',
    evidence JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_fraud_collusion_events_event_type ON fraud_collusion_events (event_type);
CREATE INDEX idx_fraud_collusion_events_session_id ON fraud_collusion_events (session_id);
CREATE INDEX idx_fraud_collusion_events_status ON fraud_collusion_events (status);
CREATE INDEX idx_fraud_collusion_events_created_at ON fraud_collusion_events (created_at);

CREATE TABLE fraud_alerts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    alert_type VARCHAR(50) NOT NULL,
    severity VARCHAR(20) NOT NULL DEFAULT 'medium',
    status VARCHAR(20) NOT NULL DEFAULT 'open',
    description TEXT,
    score_id UUID REFERENCES fraud_scores (id),
    details JSONB,
    assigned_to UUID,
    resolved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_fraud_alerts_user_id ON fraud_alerts (user_id);
CREATE INDEX idx_fraud_alerts_alert_type ON fraud_alerts (alert_type);
CREATE INDEX idx_fraud_alerts_severity ON fraud_alerts (severity);
CREATE INDEX idx_fraud_alerts_status ON fraud_alerts (status);
CREATE INDEX idx_fraud_alerts_created_at ON fraud_alerts (created_at);
