-- Wallet Service Database Schema

-- Wallets table: stores player wallets (main, bonus per currency)
CREATE TABLE IF NOT EXISTS wallets (
    id UUID PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    currency VARCHAR(10) NOT NULL,
    balance_type VARCHAR(50) NOT NULL, -- REAL, BONUS, PROMOTIONAL
    amount BIGINT DEFAULT 0,
    locked_amount BIGINT DEFAULT 0,
    version INTEGER DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, currency, balance_type)
);

CREATE INDEX idx_wallets_user_id ON wallets(user_id);
CREATE INDEX idx_wallets_currency ON wallets(currency);
CREATE INDEX idx_wallets_balance_type ON wallets(balance_type);

-- Transactions table: stores all financial transactions
CREATE TABLE IF NOT EXISTS transactions (
    transaction_id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL, -- DEPOSIT, WITHDRAWAL, BET, WIN, BONUS, etc.
    status VARCHAR(50) NOT NULL, -- PENDING, COMPLETED, FAILED, etc.
    currency VARCHAR(10) NOT NULL,
    amount BIGINT NOT NULL,
    bonus_amount BIGINT DEFAULT 0,
    fee BIGINT DEFAULT 0,
    net_amount BIGINT NOT NULL,
    payment_method VARCHAR(50),
    payment_provider VARCHAR(100),
    payment_reference VARCHAR(255),
    game_id VARCHAR(255),
    bet_id VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    processed_at TIMESTAMP
);

CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_type ON transactions(type);
CREATE INDEX idx_transactions_status ON transactions(status);
CREATE INDEX idx_transactions_created_at ON transactions(created_at);
CREATE INDEX idx_transactions_game_id ON transactions(game_id);
CREATE INDEX idx_transactions_bet_id ON transactions(bet_id);

-- Bets table: stores bet records with settlement status TABLE IF NOT EXISTS bets (
    bet_id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    currency VARCHAR(10) NOT NULL,
    game_id VARCHAR(255) NOT NULL,
    bet_type VARCHAR(50),
    selection TEXT,
    odds VARCHAR(50),
    stake BIGINT NOT NULL,
    potential_win BIGINT DEFAULT 0,
    actual_win BIGINT DEFAULT 0,
    settlement_type VARCHAR(50) DEFAULT 'BET_SETTLEMENT_TYPE_PENDING', -- PENDING, WON, LOST, CANCELLED
    status VARCHAR(50) DEFAULT 'TRANSACTION_STATUS_PENDING', -- PENDING, COMPLETED, CANCELLED
    placed_at TIMESTAMP NOT NULL DEFAULT NOW(),
    settled_at TIMESTAMP
);

CREATE INDEX idx_bets_user_id ON bets(user_id);
CREATE INDEX idx_bets_game_id ON bets(game_id);
CREATE INDEX idx_bets_status ON bets(status);
CREATE INDEX idx_bets_settlement_type ON bets(settlement_type);
CREATE INDEX idx_bets_placed_at ON bets(placed_at);

-- Bonus transactions table: stores bonus credits with wagering requirements
CREATE TABLE IF NOT EXISTS bonus_transactions (
    id UUID PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    transaction_id VARCHAR(255),
    bonus_type VARCHAR(50) NOT NULL, -- WELCOME, DEPOSIT, NO_DEPOSIT, etc.
    currency VARCHAR(10) NOT NULL,
    amount BIGINT NOT NULL,
    wagering_multiplier INTEGER DEFAULT 30,
    wagering_required BIGINT DEFAULT 0,
    wagering_met BIGINT DEFAULT 0,
    bonus_code VARCHAR(100),
    status VARCHAR(50) DEFAULT 'ACTIVE', -- ACTIVE, USED, EXPIRED
    expires_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    used_at TIMESTAMP
);

CREATE INDEX idx_bonus_transactions_user_id ON bonus_transactions(user_id);
CREATE INDEX idx_bonus_transactions_status ON bonus_transactions(status);
CREATE INDEX idx_bonus_transactions_expires_at ON bonus_transactions(expires_at);
CREATE INDEX idx_bonus_transactions_bonus_code ON bonus_transactions(bonus_code);

-- Function to update updated_at timestamp automatically
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Trigger to update updated_at on wallet changes
CREATE TRIGGER update_wallets_updated_at BEFORE UPDATE ON wallets
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- View for player balance summary
CREATE OR REPLACE VIEW v_player_balance_summary AS
SELECT
    user_id,
    currency,
    SUM(CASE WHEN balance_type = 'BALANCE_TYPE_REAL' THEN amount ELSE 0 END) as real_balance,
    SUM(CASE WHEN balance_type = 'BALANCE_TYPE_REAL' THEN locked_amount ELSE 0 END) as real_locked,
    SUM(CASE WHEN balance_type = 'BALANCE_TYPE_BONUS' THEN amount ELSE 0 END) as bonus_balance,
    SUM(CASE WHEN balance_type = 'BALANCE_TYPE_BONUS' THEN locked_amount ELSE 0 END) as bonus_locked
FROM wallets
GROUP BY user_id, currency;

-- View for transaction summary
CREATE OR REPLACE VIEW v_transaction_summary AS
SELECT
    user_id,
    currency,
    type,
    COUNT(*) as transaction_count,
    SUM(amount) as total_amount,
    SUM(CASE WHEN status = 'TRANSACTION_STATUS_COMPLETED' THEN net_amount ELSE 0 END) as completed_amount
FROM transactions
GROUP BY user_id, currency, type;
