CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE aml_transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    transaction_id UUID NOT NULL,
    transaction_type VARCHAR(30) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    risk_score DECIMAL(5, 2) DEFAULT 0.00,
    flags TEXT[],
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_aml_transactions_user_id ON aml_transactions (user_id);
CREATE INDEX idx_aml_transactions_transaction_id ON aml_transactions (transaction_id);
CREATE INDEX idx_aml_transactions_status ON aml_transactions (status);
CREATE INDEX idx_aml_transactions_created_at ON aml_transactions (created_at);
CREATE INDEX idx_aml_transactions_risk_score ON aml_transactions (risk_score);

CREATE TABLE aml_alerts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    transaction_id UUID,
    alert_type VARCHAR(50) NOT NULL,
    severity VARCHAR(20) NOT NULL DEFAULT 'medium',
    status VARCHAR(20) NOT NULL DEFAULT 'open',
    description TEXT,
    details JSONB,
    assigned_to UUID,
    resolved_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_aml_alerts_user_id ON aml_alerts (user_id);
CREATE INDEX idx_aml_alerts_alert_type ON aml_alerts (alert_type);
CREATE INDEX idx_aml_alerts_severity ON aml_alerts (severity);
CREATE INDEX idx_aml_alerts_status ON aml_alerts (status);
CREATE INDEX idx_aml_alerts_created_at ON aml_alerts (created_at);

CREATE TABLE aml_risk_scores (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    overall_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    kyc_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    transaction_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    behavioral_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    geographic_score DECIMAL(5, 2) NOT NULL DEFAULT 0.00,
    risk_level VARCHAR(20) NOT NULL DEFAULT 'low',
    last_evaluated_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_aml_risk_scores_user_id ON aml_risk_scores (user_id);
CREATE INDEX idx_aml_risk_scores_risk_level ON aml_risk_scores (risk_level);
CREATE INDEX idx_aml_risk_scores_overall_score ON aml_risk_scores (overall_score);

CREATE TABLE aml_sar_reports (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    alert_id UUID REFERENCES aml_alerts (id),
    report_type VARCHAR(30) NOT NULL DEFAULT 'sar',
    status VARCHAR(20) NOT NULL DEFAULT 'draft',
    narrative TEXT NOT NULL,
    filing_date TIMESTAMP WITH TIME ZONE,
    regulatory_reference VARCHAR(100),
    attachments JSONB,
    created_by UUID,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_aml_sar_reports_user_id ON aml_sar_reports (user_id);
CREATE INDEX idx_aml_sar_reports_status ON aml_sar_reports (status);
CREATE INDEX idx_aml_sar_reports_report_type ON aml_sar_reports (report_type);
CREATE INDEX idx_aml_sar_reports_created_at ON aml_sar_reports (created_at);
