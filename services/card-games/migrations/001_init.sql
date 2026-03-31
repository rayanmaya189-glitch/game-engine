-- Card Games Service Database Schema

-- Game Sessions table: tracks card game sessions
CREATE TABLE IF NOT EXISTS game_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id VARCHAR(255) UNIQUE NOT NULL,
    game_type VARCHAR(50) NOT NULL, -- POKER, BLACKJACK, BACCARAT, etc.
    player_id VARCHAR(255) NOT NULL,
    table_id VARCHAR(255),
    status VARCHAR(50) DEFAULT 'ACTIVE', -- ACTIVE, COMPLETED, CANCELLED
    stake BIGINT DEFAULT 0,
    currency VARCHAR(10) DEFAULT 'USD',
    deck_state JSONB DEFAULT '{}',
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    ended_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_card_game_sessions_session_id ON game_sessions(session_id);
CREATE INDEX idx_card_game_sessions_player_id ON game_sessions(player_id);
CREATE INDEX idx_card_game_sessions_game_type ON game_sessions(game_type);
CREATE INDEX idx_card_game_sessions_status ON game_sessions(status);
CREATE INDEX idx_card_game_sessions_created_at ON game_sessions(created_at);

-- Game Rounds table: tracks individual rounds within a session
CREATE TABLE IF NOT EXISTS game_rounds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    round_id VARCHAR(255) UNIQUE NOT NULL,
    session_id VARCHAR(255) NOT NULL REFERENCES game_sessions(session_id) ON DELETE CASCADE,
    round_number INTEGER NOT NULL,
    player_hand JSONB DEFAULT '[]',
    dealer_hand JSONB DEFAULT '[]',
    community_cards JSONB DEFAULT '[]',
    status VARCHAR(50) DEFAULT 'PENDING', -- PENDING, IN_PROGRESS, COMPLETED
    result VARCHAR(50), -- WIN, LOSE, PUSH, FOLD
    stake BIGINT DEFAULT 0,
    payout BIGINT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    completed_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_card_game_rounds_round_id ON game_rounds(round_id);
CREATE INDEX idx_card_game_rounds_session_id ON game_rounds(session_id);
CREATE INDEX idx_card_game_rounds_status ON game_rounds(status);
CREATE INDEX idx_card_game_rounds_created_at ON game_rounds(created_at);

-- Game Results table: final results and payouts
CREATE TABLE IF NOT EXISTS game_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    result_id VARCHAR(255) UNIQUE NOT NULL,
    session_id VARCHAR(255) NOT NULL REFERENCES game_sessions(session_id),
    round_id VARCHAR(255) REFERENCES game_rounds(round_id),
    player_id VARCHAR(255) NOT NULL,
    game_type VARCHAR(50) NOT NULL,
    hand_rank VARCHAR(100),
    stake BIGINT NOT NULL DEFAULT 0,
    payout BIGINT NOT NULL DEFAULT 0,
    net_result BIGINT NOT NULL DEFAULT 0,
    is_winner BOOLEAN DEFAULT FALSE,
    rng_seed VARCHAR(255),
    result_data JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_card_game_results_result_id ON game_results(result_id);
CREATE INDEX idx_card_game_results_session_id ON game_results(session_id);
CREATE INDEX idx_card_game_results_round_id ON game_results(round_id);
CREATE INDEX idx_card_game_results_player_id ON game_results(player_id);
CREATE INDEX idx_card_game_results_game_type ON game_results(game_type);
CREATE INDEX idx_card_game_results_created_at ON game_results(created_at);

-- Trigger to update updated_at on game_sessions changes
CREATE TRIGGER update_card_game_sessions_updated_at BEFORE UPDATE ON game_sessions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update updated_at on game_rounds changes
CREATE TRIGGER update_card_game_rounds_updated_at BEFORE UPDATE ON game_rounds
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
