-- Migration: Live Betting and Cash Out Support
-- Adds support for live events, parlay bets, cash outs, and odds history

-- Add new columns to events table
ALTER TABLE sports_events 
ADD COLUMN IF NOT EXISTS period VARCHAR(50),
ADD COLUMN IF NOT EXISTS minute INTEGER DEFAULT 0;

-- Add market_type column to markets
ALTER TABLE sports_markets 
ADD COLUMN IF NOT EXISTS market_type VARCHAR(50) DEFAULT 'moneyline';

-- Create market selections table
CREATE TABLE IF NOT EXISTS sports_market_selections (
    selection_id VARCHAR(50) PRIMARY KEY,
    market_id VARCHAR(50) NOT NULL REFERENCES sports_markets(market_id),
    selection VARCHAR(100) NOT NULL,
    odds DECIMAL(10, 3) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_market_selections_market_id ON sports_market_selections(market_id);

-- Create odds history table
CREATE TABLE IF NOT EXISTS sports_odds_history (
    id SERIAL PRIMARY KEY,
    event_id VARCHAR(50) NOT NULL,
    market_id VARCHAR(50) NOT NULL,
    selection VARCHAR(100) NOT NULL,
    old_odds DECIMAL(10, 3) NOT NULL,
    new_odds DECIMAL(10, 3) NOT NULL,
    changed_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_odds_history_event_id ON sports_odds_history(event_id);
CREATE INDEX idx_odds_history_changed_at ON sports_odds_history(changed_at);

-- Create parlay bets table
CREATE TABLE IF NOT EXISTS sports_parlay_bets (
    bet_id VARCHAR(50) PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    stake DECIMAL(15, 2) NOT NULL,
    total_odds DECIMAL(10, 3) NOT NULL,
    potential_win DECIMAL(15, 2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    placed_at TIMESTAMP DEFAULT NOW(),
    settled_at TIMESTAMP
);

CREATE INDEX idx_parlay_bets_user_id ON sports_parlay_bets(user_id);
CREATE INDEX idx_parlay_bets_status ON sports_parlay_bets(status);

-- Create parlay selections table
CREATE TABLE IF NOT EXISTS sports_parlay_selections (
    id SERIAL PRIMARY KEY,
    bet_id VARCHAR(50) NOT NULL REFERENCES sports_parlay_bets(bet_id),
    event_id VARCHAR(50) NOT NULL,
    event_name VARCHAR(200),
    market_id VARCHAR(50) NOT NULL,
    selection VARCHAR(100) NOT NULL,
    odds DECIMAL(10, 3) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    result VARCHAR(20)
);

CREATE INDEX idx_parlay_selections_bet_id ON sports_parlay_selections(bet_id);

-- Create cash outs table
CREATE TABLE IF NOT EXISTS sports_cashouts (
    cash_out_id VARCHAR(50) PRIMARY KEY,
    bet_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    original_stake DECIMAL(15, 2) NOT NULL,
    original_odds DECIMAL(10, 3) NOT NULL,
    current_odds DECIMAL(10, 3) NOT NULL,
    cash_out_amount DECIMAL(15, 2) NOT NULL,
    profit DECIMAL(15, 2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    requested_at TIMESTAMP DEFAULT NOW(),
    completed_at TIMESTAMP
);

CREATE INDEX idx_cashouts_user_id ON sports_cashouts(user_id);
CREATE INDEX idx_cashouts_bet_id ON sports_cashouts(bet_id);
CREATE INDEX idx_cashouts_status ON sports_cashouts(status);

-- Create system bets table (for Patent, Yankee, etc.)
CREATE TABLE IF NOT EXISTS sports_system_bets (
    bet_id VARCHAR(50) PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    bet_type VARCHAR(20) NOT NULL,
    system_type VARCHAR(20) NOT NULL,
    stake DECIMAL(15, 2) NOT NULL,
    num_selections INTEGER NOT NULL,
    num_wins INTEGER DEFAULT 0,
    total_odds DECIMAL(10, 3) NOT NULL,
    potential_win DECIMAL(15, 2) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    placed_at TIMESTAMP DEFAULT NOW(),
    settled_at TIMESTAMP
);

CREATE INDEX idx_system_bets_user_id ON sports_system_bets(user_id);
CREATE INDEX idx_system_bets_status ON sports_system_bets(status);

-- Create system bet selections table
CREATE TABLE IF NOT EXISTS sports_system_selections (
    id SERIAL PRIMARY KEY,
    bet_id VARCHAR(50) NOT NULL REFERENCES sports_system_bets(bet_id),
    event_id VARCHAR(50) NOT NULL,
    event_name VARCHAR(200),
    market_id VARCHAR(50) NOT NULL,
    selection VARCHAR(100) NOT NULL,
    odds DECIMAL(10, 3) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending'
);

CREATE INDEX idx_system_selections_bet_id ON sports_system_selections(bet_id);

-- Create live events view
CREATE OR REPLACE VIEW v_live_events AS
SELECT 
    e.event_id,
    e.sport_id,
    e.league_id,
    e.home_team,
    e.away_team,
    e.home_score,
    e.away_score,
    e.status,
    e.period,
    e.minute,
    e.start_time,
    s.name as sport_name
FROM sports_events e
JOIN sports s ON e.sport_id = s.sport_id
WHERE e.status = 'live';

-- Create pending bets view
CREATE OR REPLACE VIEW v_pending_bets AS
SELECT 
    b.bet_id,
    b.user_id,
    b.event_id,
    e.home_team,
    e.away_team,
    b.market_id,
    b.selection,
    b.stake,
    b.odds,
    b.potential_win,
    b.status,
    b.placed_at
FROM sports_bets b
JOIN sports_events e ON b.event_id = e.event_id
WHERE b.status = 'pending';
