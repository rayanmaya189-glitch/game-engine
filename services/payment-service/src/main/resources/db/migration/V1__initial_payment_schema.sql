-- Payment Service Database Schema
-- Version: V1

-- Payments table
CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(255) NOT NULL,
    external_id VARCHAR(255) NOT NULL UNIQUE,
    type VARCHAR(50) NOT NULL,
    method VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    amount DECIMAL(19, 4) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    converted_amount DECIMAL(19, 4),
    converted_currency VARCHAR(3),
    fee DECIMAL(19, 4),
    net_amount DECIMAL(19, 4),
    description TEXT,
    metadata TEXT,
    payment_gateway VARCHAR(100),
    gateway_response TEXT,
    failure_reason TEXT,
    retries INTEGER DEFAULT 0,
    processed_at TIMESTAMP,
    completed_at TIMESTAMP,
    expires_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for faster queries
CREATE INDEX idx_payment_user_id ON payments(user_id);
CREATE INDEX idx_payment_status ON payments(status);
CREATE INDEX idx_payment_type ON payments(type);
CREATE INDEX idx_payment_external_id ON payments(external_id);
CREATE INDEX idx_payment_created_at ON payments(created_at);
CREATE INDEX idx_payment_user_id_type_status ON payments(user_id, type, status);

-- Payment methods configuration table
CREATE TABLE IF NOT EXISTS payment_methods_config (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    method VARCHAR(50) NOT NULL UNIQUE,
    enabled BOOLEAN DEFAULT true,
    min_amount DECIMAL(19, 4),
    max_amount DECIMAL(19, 4),
    fee_percentage DECIMAL(5, 2) DEFAULT 0,
    fee_fixed DECIMAL(19, 4) DEFAULT 0,
    processing_time VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Insert default payment methods configuration
INSERT INTO payment_methods_config (method, enabled, min_amount, max_amount, fee_percentage, fee_fixed, processing_time) VALUES
    ('CREDIT_CARD', true, 10, 50000, 2.5, 0, 'Instant'),
    ('DEBIT_CARD', true, 10, 50000, 2.5, 0, 'Instant'),
    ('VIRTUAL_CARD', true, 10, 10000, 2.5, 0, 'Instant'),
    ('PAYPAL', true, 10, 50000, 2.9, 0, 'Instant'),
    ('SKRILL', true, 10, 50000, 2.5, 0, 'Instant'),
    ('NETELLER', true, 10, 50000, 2.5, 0, 'Instant'),
    ('BITCOIN', true, 50, 100000, 1.0, 0, '10-30 minutes'),
    ('ETHEREUM', true, 50, 100000, 1.0, 0, '10-30 minutes'),
    ('BANK_TRANSFER', true, 100, 100000, 0.5, 5, '2-5 business days'),
    ('INSTANT_BANK_TRANSFER', true, 10, 50000, 1.0, 0, 'Instant')
ON CONFLICT (method) DO NOTHING;

-- Currency exchange rates table
CREATE TABLE IF NOT EXISTS exchange_rates (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_currency VARCHAR(3) NOT NULL,
    to_currency VARCHAR(3) NOT NULL,
    rate DECIMAL(20, 8) NOT NULL,
    valid_from TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    valid_to TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(from_currency, to_currency, valid_from)
);

-- Insert default exchange rates
INSERT INTO exchange_rates (from_currency, to_currency, rate) VALUES
    ('USD', 'EUR', 0.92),
    ('USD', 'GBP', 0.79),
    ('EUR', 'USD', 1.09),
    ('EUR', 'GBP', 0.86),
    ('GBP', 'USD', 1.27),
    ('GBP', 'EUR', 1.16)
ON CONFLICT DO NOTHING;

-- Trigger to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_payments_updated_at BEFORE UPDATE ON payments
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_payment_methods_config_updated_at BEFORE UPDATE ON payment_methods_config
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Comments
COMMENT ON TABLE payments IS 'Stores all payment transactions including deposits, withdrawals, refunds';
COMMENT ON TABLE payment_methods_config IS 'Configuration for supported payment methods';
COMMENT ON TABLE exchange_rates IS 'Currency exchange rates';
