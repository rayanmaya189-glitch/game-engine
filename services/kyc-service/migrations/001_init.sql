CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE kyc_verifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    verification_level VARCHAR(20) NOT NULL DEFAULT 'basic',
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    provider VARCHAR(50),
    provider_reference VARCHAR(255),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    date_of_birth DATE,
    nationality VARCHAR(2),
    risk_rating VARCHAR(20) DEFAULT 'unverified',
    verified_at TIMESTAMP WITH TIME ZONE,
    expires_at TIMESTAMP WITH TIME ZONE,
    rejection_reason TEXT,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_kyc_verifications_user_id ON kyc_verifications (user_id);
CREATE INDEX idx_kyc_verifications_status ON kyc_verifications (status);
CREATE INDEX idx_kyc_verifications_verification_level ON kyc_verifications (verification_level);
CREATE INDEX idx_kyc_verifications_risk_rating ON kyc_verifications (risk_rating);

CREATE TABLE kyc_documents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    verification_id UUID NOT NULL REFERENCES kyc_verifications (id),
    user_id UUID NOT NULL,
    document_type VARCHAR(50) NOT NULL,
    document_number VARCHAR(100),
    issuing_country VARCHAR(2),
    issuing_authority VARCHAR(100),
    issue_date DATE,
    expiry_date DATE,
    file_path VARCHAR(500),
    file_hash VARCHAR(255),
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    verified_at TIMESTAMP WITH TIME ZONE,
    rejection_reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_kyc_documents_verification_id ON kyc_documents (verification_id);
CREATE INDEX idx_kyc_documents_user_id ON kyc_documents (user_id);
CREATE INDEX idx_kyc_documents_document_type ON kyc_documents (document_type);
CREATE INDEX idx_kyc_documents_status ON kyc_documents (status);

CREATE TABLE kyc_address_verifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    verification_id UUID NOT NULL REFERENCES kyc_verifications (id),
    user_id UUID NOT NULL,
    address_line1 VARCHAR(255) NOT NULL,
    address_line2 VARCHAR(255),
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100),
    postal_code VARCHAR(20) NOT NULL,
    country_code VARCHAR(2) NOT NULL,
    document_type VARCHAR(50),
    file_path VARCHAR(500),
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    verified_at TIMESTAMP WITH TIME ZONE,
    rejection_reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_kyc_address_verifications_verification_id ON kyc_address_verifications (verification_id);
CREATE INDEX idx_kyc_address_verifications_user_id ON kyc_address_verifications (user_id);
CREATE INDEX idx_kyc_address_verifications_status ON kyc_address_verifications (status);
CREATE INDEX idx_kyc_address_verifications_country_code ON kyc_address_verifications (country_code);

CREATE TABLE kyc_audit_log (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    verification_id UUID REFERENCES kyc_verifications (id),
    user_id UUID NOT NULL,
    action VARCHAR(50) NOT NULL,
    performed_by UUID,
    old_status VARCHAR(20),
    new_status VARCHAR(20),
    details JSONB,
    ip_address INET,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_kyc_audit_log_verification_id ON kyc_audit_log (verification_id);
CREATE INDEX idx_kyc_audit_log_user_id ON kyc_audit_log (user_id);
CREATE INDEX idx_kyc_audit_log_action ON kyc_audit_log (action);
CREATE INDEX idx_kyc_audit_log_created_at ON kyc_audit_log (created_at);
