CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    gateway_id UUID,
    reference_id VARCHAR(255) NOT NULL UNIQUE,
    external_id VARCHAR(255),
    amount DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'USD',
    payment_type VARCHAR(30) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    provider VARCHAR(50) NOT NULL,
    provider_transaction_id VARCHAR(255),
    provider_response JSONB,
    failure_reason TEXT,
    metadata JSONB,
    processed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_payments_user_id ON payments (user_id);
CREATE INDEX idx_payments_reference_id ON payments (reference_id);
CREATE INDEX idx_payments_status ON payments (status);
CREATE INDEX idx_payments_payment_type ON payments (payment_type);
CREATE INDEX idx_payments_provider ON payments (provider);
CREATE INDEX idx_payments_created_at ON payments (created_at);

CREATE TABLE payment_callbacks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    payment_id UUID REFERENCES payments (id),
    provider VARCHAR(50) NOT NULL,
    callback_type VARCHAR(30) NOT NULL,
    payload JSONB NOT NULL,
    signature VARCHAR(500),
    processed BOOLEAN NOT NULL DEFAULT FALSE,
    processed_at TIMESTAMP WITH TIME ZONE,
    ip_address INET,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_payment_callbacks_payment_id ON payment_callbacks (payment_id);
CREATE INDEX idx_payment_callbacks_provider ON payment_callbacks (provider);
CREATE INDEX idx_payment_callbacks_processed ON payment_callbacks (processed);
CREATE INDEX idx_payment_callbacks_created_at ON payment_callbacks (created_at);

CREATE TABLE payment_gateways (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL UNIQUE,
    provider VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    supported_currencies VARCHAR(3)[] NOT NULL DEFAULT ARRAY['USD'],
    supported_types VARCHAR(30)[] NOT NULL DEFAULT ARRAY['deposit', 'withdrawal'],
    min_amount DECIMAL(15, 2) DEFAULT 1.00,
    max_amount DECIMAL(15, 2) DEFAULT 50000.00,
    fee_percentage DECIMAL(5, 4) DEFAULT 0.0000,
    fee_fixed DECIMAL(15, 2) DEFAULT 0.00,
    credentials JSONB,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_payment_gateways_provider ON payment_gateways (provider);
CREATE INDEX idx_payment_gateways_status ON payment_gateways (status);
