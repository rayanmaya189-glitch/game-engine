-- Merchant Service Database Schema

-- Merchants table
CREATE TABLE IF NOT EXISTS merchants (
    id SERIAL PRIMARY KEY,
    merchant_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    commission_rate DECIMAL(5, 2) DEFAULT 10.00,
    api_key VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_merchants_merchant_id ON merchants(merchant_id);

-- Merchant players table
CREATE TABLE IF NOT EXISTS merchant_players (
    id SERIAL PRIMARY KEY,
    merchant_id VARCHAR(255) NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(merchant_id, player_id),
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id)
);

CREATE INDEX idx_merchant_players_merchant_id ON merchant_players(merchant_id);

-- Merchant agents table
CREATE TABLE IF NOT EXISTS merchant_agents (
    id SERIAL PRIMARY KEY,
    merchant_id VARCHAR(255) NOT NULL,
    agent_id VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(merchant_id, agent_id),
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id)
);

CREATE INDEX idx_merchant_agents_merchant_id ON merchant_agents(merchant_id);

-- Merchant agent invitations table
CREATE TABLE IF NOT EXISTS merchant_agent_invitations (
    id SERIAL PRIMARY KEY,
    merchant_id VARCHAR(255) NOT NULL,
    agent_id VARCHAR(255) NOT NULL,
    invitation_id VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id),
    FOREIGN KEY (agent_id) REFERENCES merchant_agents(agent_id)
);

CREATE INDEX idx_agent_invitations_agent_id ON merchant_agent_invitations(agent_id);
CREATE INDEX idx_agent_invitations_status ON merchant_agent_invitations(status);

-- Merchant config table
CREATE TABLE IF NOT EXISTS merchant_config (
    id SERIAL PRIMARY KEY,
    merchant_id VARCHAR(255) NOT NULL,
    key VARCHAR(255) NOT NULL,
    value TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(merchant_id, key),
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id)
);

-- Merchant webhooks table
CREATE TABLE IF NOT EXISTS merchant_webhooks (
    id SERIAL PRIMARY KEY,
    merchant_id VARCHAR(255) NOT NULL,
    webhook_id VARCHAR(255) UNIQUE NOT NULL,
    url TEXT NOT NULL,
    events TEXT,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id)
);

CREATE INDEX idx_merchant_webhooks_merchant_id ON merchant_webhooks(merchant_id);

-- Merchant reports tables
CREATE TABLE IF NOT EXISTS merchant_reports (
    id SERIAL PRIMARY KEY,
    merchant_id VARCHAR(255) NOT NULL,
    player_id VARCHAR(255),
    total_revenue DECIMAL(15, 2) DEFAULT 0,
    total_deposits DECIMAL(15, 2) DEFAULT 0,
    total_withdrawals DECIMAL(15, 2) DEFAULT 0,
    report_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id)
);

CREATE TABLE IF NOT EXISTS player_reports (
    id SERIAL PRIMARY KEY,
    merchant_id VARCHAR(255) NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    total_bets DECIMAL(15, 2) DEFAULT 0,
    total_wins DECIMAL(15, 2) DEFAULT 0,
    net_revenue DECIMAL(15, 2) DEFAULT 0,
    games_played INTEGER DEFAULT 0,
    report_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id)
);

CREATE TABLE IF NOT EXISTS game_reports (
    id SERIAL PRIMARY KEY,
    merchant_id VARCHAR(255) NOT NULL,
    game_id VARCHAR(255) NOT NULL,
    total_bets DECIMAL(15, 2) DEFAULT 0,
    total_wins DECIMAL(15, 2) DEFAULT 0,
    total_players INTEGER DEFAULT 0,
    plays INTEGER DEFAULT 0,
    report_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (merchant_id) REFERENCES merchants(merchant_id)
);
