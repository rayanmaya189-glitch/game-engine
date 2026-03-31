-- Slot Games Service Database Schema

-- Spin Sessions table: tracks slot game sessions
CREATE TABLE IF NOT EXISTS spin_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id VARCHAR(255) UNIQUE NOT NULL,
    game_id VARCHAR(255) NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'ACTIVE', -- ACTIVE, COMPLETED, CANCELLED
    stake BIGINT DEFAULT 0,
    currency VARCHAR(10) DEFAULT 'USD',
    total_spins INTEGER DEFAULT 0,
    total_wagered BIGINT DEFAULT 0,
    total_won BIGINT DEFAULT 0,
    balance BIGINT DEFAULT 0,
    config JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    ended_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_spin_sessions_session_id ON spin_sessions(session_id);
CREATE INDEX idx_spin_sessions_game_id ON spin_sessions(game_id);
CREATE INDEX idx_spin_sessions_player_id ON spin_sessions(player_id);
CREATE INDEX idx_spin_sessions_status ON spin_sessions(status);
CREATE INDEX idx_spin_sessions_created_at ON spin_sessions(created_at);

-- Spin Results table: stores individual spin outcomes
CREATE TABLE IF NOT EXISTS spin_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    spin_id VARCHAR(255) UNIQUE NOT NULL,
    session_id VARCHAR(255) NOT NULL REFERENCES spin_sessions(session_id) ON DELETE CASCADE,
    game_id VARCHAR(255) NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    spin_number INTEGER NOT NULL,
    reel_strips JSONB NOT NULL DEFAULT '[]', -- Final reel positions
    symbols JSONB NOT NULL DEFAULT '[]', -- 2D grid of symbols
    stake BIGINT NOT NULL DEFAULT 0,
    win_amount BIGINT DEFAULT 0,
    multiplier DECIMAL(10, 4) DEFAULT 1.0,
    is_bonus_spin BOOLEAN DEFAULT FALSE,
    is_free_spin BOOLEAN DEFAULT FALSE,
    free_spins_remaining INTEGER DEFAULT 0,
    rng_seed VARCHAR(255),
    status VARCHAR(50) DEFAULT 'COMPLETED',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_spin_results_spin_id ON spin_results(spin_id);
CREATE INDEX idx_spin_results_session_id ON spin_results(session_id);
CREATE INDEX idx_spin_results_game_id ON spin_results(game_id);
CREATE INDEX idx_spin_results_player_id ON spin_results(player_id);
CREATE INDEX idx_spin_results_created_at ON spin_results(created_at);

-- Payline Wins table: stores individual payline winning combinations
CREATE TABLE IF NOT EXISTS payline_wins (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    spin_id VARCHAR(255) NOT NULL REFERENCES spin_results(spin_id) ON DELETE CASCADE,
    session_id VARCHAR(255) NOT NULL,
    payline_index INTEGER NOT NULL,
    payline_name VARCHAR(100),
    symbols JSONB NOT NULL DEFAULT '[]',
    match_count INTEGER NOT NULL DEFAULT 0,
    symbol_multiplier DECIMAL(10, 4) DEFAULT 1.0,
    win_amount BIGINT NOT NULL DEFAULT 0,
    is_wild BOOLEAN DEFAULT FALSE,
    is_scatter BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_payline_wins_spin_id ON payline_wins(spin_id);
CREATE INDEX idx_payline_wins_session_id ON payline_wins(session_id);
CREATE INDEX idx_payline_wins_created_at ON payline_wins(created_at);

-- Trigger to update updated_at on spin_sessions changes
CREATE TRIGGER update_spin_sessions_updated_at BEFORE UPDATE ON spin_sessions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
