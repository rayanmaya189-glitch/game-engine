-- Betting Service Database Schema

-- Bets table: stores all bet records
CREATE TABLE IF NOT EXISTS bets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bet_id VARCHAR(255) UNIQUE NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    bet_type VARCHAR(50) NOT NULL, -- SINGLE, MULTI, SYSTEM, LIVE
    status VARCHAR(50) DEFAULT 'PENDING', -- PENDING, ACCEPTED, WON, LOST, CANCELLED, CASHED_OUT
    stake BIGINT NOT NULL DEFAULT 0,
    currency VARCHAR(10) DEFAULT 'USD',
    potential_payout BIGINT DEFAULT 0,
    actual_payout BIGINT DEFAULT 0,
    total_odds DECIMAL(10, 4) DEFAULT 1.0,
    selections_count INTEGER DEFAULT 1,
    is_live BOOLEAN DEFAULT FALSE,
    is_free_bet BOOLEAN DEFAULT FALSE,
    bonus_id VARCHAR(255),
    game_id VARCHAR(255),
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    settled_at TIMESTAMP WITH TIME ZONE,
    expires_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_bets_bet_id ON bets(bet_id);
CREATE INDEX idx_bets_player_id ON bets(player_id);
CREATE INDEX idx_bets_bet_type ON bets(bet_type);
CREATE INDEX idx_bets_status ON bets(status);
CREATE INDEX idx_bets_game_id ON bets(game_id);
CREATE INDEX idx_bets_is_live ON bets(is_live);
CREATE INDEX idx_bets_created_at ON bets(created_at);

-- Bet Selections table: stores individual selections within a bet
CREATE TABLE IF NOT EXISTS bet_selections (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    selection_id VARCHAR(255) UNIQUE NOT NULL,
    bet_id VARCHAR(255) NOT NULL REFERENCES bets(bet_id) ON DELETE CASCADE,
    event_id VARCHAR(255) NOT NULL,
    event_name VARCHAR(255),
    market_id VARCHAR(255),
    market_name VARCHAR(255),
    selection_name VARCHAR(255) NOT NULL,
    odds DECIMAL(10, 4) NOT NULL,
    status VARCHAR(50) DEFAULT 'PENDING', -- PENDING, WON, LOST, VOID
    result VARCHAR(255),
    is_live BOOLEAN DEFAULT FALSE,
    sort_order INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    settled_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_bet_selections_selection_id ON bet_selections(selection_id);
CREATE INDEX idx_bet_selections_bet_id ON bet_selections(bet_id);
CREATE INDEX idx_bet_selections_event_id ON bet_selections(event_id);
CREATE INDEX idx_bet_selections_status ON bet_selections(status);
CREATE INDEX idx_bet_selections_created_at ON bet_selections(created_at);

-- Bet Settlements table: stores settlement records for completed bets
CREATE TABLE IF NOT EXISTS bet_settlements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    settlement_id VARCHAR(255) UNIQUE NOT NULL,
    bet_id VARCHAR(255) NOT NULL REFERENCES bets(bet_id),
    player_id VARCHAR(255) NOT NULL,
    settlement_type VARCHAR(50) NOT NULL, -- WIN, LOSS, PARTIAL, CASHOUT, VOID
    stake BIGINT NOT NULL DEFAULT 0,
    payout BIGINT NOT NULL DEFAULT 0,
    net_amount BIGINT NOT NULL DEFAULT 0,
    total_odds DECIMAL(10, 4) DEFAULT 1.0,
    transaction_id VARCHAR(255),
    settled_by VARCHAR(255), -- SYSTEM or user id
    settlement_data JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_bet_settlements_settlement_id ON bet_settlements(settlement_id);
CREATE INDEX idx_bet_settlements_bet_id ON bet_settlements(bet_id);
CREATE INDEX idx_bet_settlements_player_id ON bet_settlements(player_id);
CREATE INDEX idx_bet_settlements_settlement_type ON bet_settlements(settlement_type);
CREATE INDEX idx_bet_settlements_created_at ON bet_settlements(created_at);

-- Trigger to update updated_at on bets changes
CREATE TRIGGER update_bets_updated_at BEFORE UPDATE ON bets
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update updated_at on bet_selections changes
CREATE TRIGGER update_bet_selections_updated_at BEFORE UPDATE ON bet_selections
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
