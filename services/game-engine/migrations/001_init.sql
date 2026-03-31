-- Game Engine Service Database Schema

-- Engine Sessions table: tracks game engine sessions
CREATE TABLE IF NOT EXISTS engine_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id VARCHAR(255) UNIQUE NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    game_id VARCHAR(255) NOT NULL,
    game_type VARCHAR(50) NOT NULL,
    provider VARCHAR(100),
    status VARCHAR(50) DEFAULT 'ACTIVE', -- ACTIVE, PAUSED, COMPLETED, CANCELLED
    balance BIGINT DEFAULT 0,
    currency VARCHAR(10) DEFAULT 'USD',
    language VARCHAR(10) DEFAULT 'EN',
    device_type VARCHAR(50),
    ip_address VARCHAR(45),
    session_token VARCHAR(512),
    config JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    last_activity_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    ended_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_engine_sessions_session_id ON engine_sessions(session_id);
CREATE INDEX idx_engine_sessions_player_id ON engine_sessions(player_id);
CREATE INDEX idx_engine_sessions_game_id ON engine_sessions(game_id);
CREATE INDEX idx_engine_sessions_game_type ON engine_sessions(game_type);
CREATE INDEX idx_engine_sessions_status ON engine_sessions(status);
CREATE INDEX idx_engine_sessions_created_at ON engine_sessions(created_at);

-- Engine Results table: stores game round results from the engine
CREATE TABLE IF NOT EXISTS engine_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    result_id VARCHAR(255) UNIQUE NOT NULL,
    session_id VARCHAR(255) NOT NULL REFERENCES engine_sessions(session_id) ON DELETE CASCADE,
    player_id VARCHAR(255) NOT NULL,
    game_id VARCHAR(255) NOT NULL,
    round_id VARCHAR(255) NOT NULL,
    action VARCHAR(50) NOT NULL, -- SPIN, DEAL, ROLL, HIT, STAND, etc.
    stake BIGINT DEFAULT 0,
    win_amount BIGINT DEFAULT 0,
    balance_before BIGINT DEFAULT 0,
    balance_after BIGINT DEFAULT 0,
    multiplier DECIMAL(10, 4) DEFAULT 1.0,
    rng_seed VARCHAR(255),
    result_data JSONB DEFAULT '{}',
    status VARCHAR(50) DEFAULT 'COMPLETED', -- PENDING, COMPLETED, FAILED
    error_message TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_engine_results_result_id ON engine_results(result_id);
CREATE INDEX idx_engine_results_session_id ON engine_results(session_id);
CREATE INDEX idx_engine_results_player_id ON engine_results(player_id);
CREATE INDEX idx_engine_results_game_id ON engine_results(game_id);
CREATE INDEX idx_engine_results_round_id ON engine_results(round_id);
CREATE INDEX idx_engine_results_action ON engine_results(action);
CREATE INDEX idx_engine_results_created_at ON engine_results(created_at);

-- Trigger to update updated_at on engine_sessions changes
CREATE TRIGGER update_engine_sessions_updated_at BEFORE UPDATE ON engine_sessions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
