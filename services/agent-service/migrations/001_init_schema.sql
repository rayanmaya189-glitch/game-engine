-- Agent Service Database Schema

-- Agent players table (downline players managed by agent)
CREATE TABLE IF NOT EXISTS agent_players (
    id SERIAL PRIMARY KEY,
    player_id VARCHAR(255) UNIQUE NOT NULL,
    agent_id VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    status VARCHAR(50) DEFAULT 'active',
    total_deposits DECIMAL(15, 2) DEFAULT 0,
    total_bets DECIMAL(15, 2) DEFAULT 0,
    balance DECIMAL(15, 2) DEFAULT 0,
    deposit_limit DECIMAL(15, 2),
    bet_limit DECIMAL(15, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_agent_players_agent_id ON agent_players(agent_id);
CREATE INDEX idx_agent_players_status ON agent_players(status);

-- Agent dashboard summary table
CREATE TABLE IF NOT EXISTS agent_dashboard (
    id SERIAL PRIMARY KEY,
    agent_id VARCHAR(255) UNIQUE NOT NULL,
    total_players INTEGER DEFAULT 0,
    active_players INTEGER DEFAULT 0,
    total_commission DECIMAL(15, 2) DEFAULT 0,
    pending_commission DECIMAL(15, 2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_agent_dashboard_agent_id ON agent_dashboard(agent_id);

-- Agent commissions table
CREATE TABLE IF NOT EXISTS agent_commissions (
    id SERIAL PRIMARY KEY,
    commission_id VARCHAR(255) UNIQUE NOT NULL,
    agent_id VARCHAR(255) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    type VARCHAR(50) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    transaction_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    claimed_at TIMESTAMP
);

CREATE INDEX idx_agent_commissions_agent_id ON agent_commissions(agent_id);
CREATE INDEX idx_agent_commissions_status ON agent_commissions(status);

-- Agent table
CREATE TABLE IF NOT EXISTS agents (
    id SERIAL PRIMARY KEY,
    agent_id VARCHAR(255) UNIQUE NOT NULL,
    merchant_id VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'active',
    tier VARCHAR(50) DEFAULT 'bronze',
    commission_rate DECIMAL(5, 2) DEFAULT 10.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_agents_agent_id ON agents(agent_id);
CREATE INDEX idx_agents_merchant_id ON agents(merchant_id);
