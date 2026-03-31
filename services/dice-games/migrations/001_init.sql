-- Dice Games Service Database Schema

-- Game Sessions table: tracks dice game sessions
CREATE TABLE IF NOT EXISTS game_sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id VARCHAR(255) UNIQUE NOT NULL,
    game_type VARCHAR(50) NOT NULL, -- CRAPS, SIC_BO, DICE_ROLL, etc.
    player_id VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'ACTIVE', -- ACTIVE, COMPLETED, CANCELLED
    currency VARCHAR(10) DEFAULT 'USD',
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    ended_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_dice_game_sessions_session_id ON game_sessions(session_id);
CREATE INDEX idx_dice_game_sessions_player_id ON game_sessions(player_id);
CREATE INDEX idx_dice_game_sessions_game_type ON game_sessions(game_type);
CREATE INDEX idx_dice_game_sessions_status ON game_sessions(status);
CREATE INDEX idx_dice_game_sessions_created_at ON game_sessions(created_at);

-- Game Rounds table: tracks individual dice rolls
CREATE TABLE IF NOT EXISTS game_rounds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    round_id VARCHAR(255) UNIQUE NOT NULL,
    session_id VARCHAR(255) NOT NULL REFERENCES game_sessions(session_id) ON DELETE CASCADE,
    round_number INTEGER NOT NULL,
    dice_values JSONB NOT NULL DEFAULT '[]', -- Array of dice results
    target_value INTEGER,
    bet_type VARCHAR(50), -- OVER, UNDER, EXACT, RANGE
    status VARCHAR(50) DEFAULT 'PENDING', -- PENDING, ROLLED, SETTLED
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    completed_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_dice_game_rounds_round_id ON game_rounds(round_id);
CREATE INDEX idx_dice_game_rounds_session_id ON game_rounds(session_id);
CREATE INDEX idx_dice_game_rounds_status ON game_rounds(status);
CREATE INDEX idx_dice_game_rounds_created_at ON game_rounds(created_at);

-- Bet Results table: stores individual bet outcomes
CREATE TABLE IF NOT EXISTS bet_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    result_id VARCHAR(255) UNIQUE NOT NULL,
    session_id VARCHAR(255) NOT NULL REFERENCES game_sessions(session_id),
    round_id VARCHAR(255) NOT NULL REFERENCES game_rounds(round_id),
    player_id VARCHAR(255) NOT NULL,
    bet_type VARCHAR(50) NOT NULL,
    bet_value VARCHAR(100),
    stake BIGINT NOT NULL DEFAULT 0,
    multiplier DECIMAL(10, 4) DEFAULT 1.0,
    payout BIGINT NOT NULL DEFAULT 0,
    net_result BIGINT NOT NULL DEFAULT 0,
    is_winner BOOLEAN DEFAULT FALSE,
    rng_seed VARCHAR(255),
    result_data JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_dice_bet_results_result_id ON bet_results(result_id);
CREATE INDEX idx_dice_bet_results_session_id ON bet_results(session_id);
CREATE INDEX idx_dice_bet_results_round_id ON bet_results(round_id);
CREATE INDEX idx_dice_bet_results_player_id ON bet_results(player_id);
CREATE INDEX idx_dice_bet_results_created_at ON bet_results(created_at);

-- Trigger to update updated_at on game_sessions changes
CREATE TRIGGER update_dice_game_sessions_updated_at BEFORE UPDATE ON game_sessions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update updated_at on game_rounds changes
CREATE TRIGGER update_dice_game_rounds_updated_at BEFORE UPDATE ON game_rounds
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
