-- Sports Betting Service Database Schema

-- Sports table
CREATE TABLE IF NOT EXISTS sports (
    id SERIAL PRIMARY KEY,
    sport_id VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    icon VARCHAR(255),
    status VARCHAR(50) DEFAULT 'active',
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_sports_status ON sports(status);

-- Insert default sports
INSERT INTO sports (sport_id, name, icon, sort_order) VALUES
('football', 'Football', 'soccer', 1),
('basketball', 'Basketball', 'basketball', 2),
('tennis', 'Tennis', 'tennis', 3),
('baseball', 'Baseball', 'baseball', 4),
('hockey', 'Hockey', 'hockey', 5),
('ufc', 'UFC/MMA', 'fight', 6),
('boxing', 'Boxing', 'boxing', 7),
('golf', 'Golf', 'golf', 8);

-- Events table
CREATE TABLE IF NOT EXISTS sports_events (
    id SERIAL PRIMARY KEY,
    event_id VARCHAR(255) UNIQUE NOT NULL,
    sport_id VARCHAR(50) NOT NULL,
    league_id VARCHAR(100),
    home_team VARCHAR(255) NOT NULL,
    away_team VARCHAR(255) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    status VARCHAR(50) DEFAULT 'scheduled',
    home_score INT DEFAULT 0,
    away_score INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sport_id) REFERENCES sports(sport_id)
);

CREATE INDEX idx_sports_events_sport_id ON sports_events(sport_id);
CREATE INDEX idx_sports_events_status ON sports_events(status);
CREATE INDEX idx_sports_events_start_time ON sports_events(start_time);

-- Markets table
CREATE TABLE IF NOT EXISTS sports_markets (
    id SERIAL PRIMARY KEY,
    market_id VARCHAR(255) UNIQUE NOT NULL,
    event_id VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    status VARCHAR(50) DEFAULT 'open',
    home_odds DECIMAL(10, 3),
    draw_odds DECIMAL(10, 3),
    away_odds DECIMAL(10, 3),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES sports_events(event_id)
);

CREATE INDEX idx_sports_markets_event_id ON sports_markets(event_id);

-- Bets table
CREATE TABLE IF NOT EXISTS sports_bets (
    id SERIAL PRIMARY KEY,
    bet_id VARCHAR(255) UNIQUE NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    event_id VARCHAR(255) NOT NULL,
    market_id VARCHAR(255) NOT NULL,
    selection VARCHAR(50) NOT NULL,
    stake DECIMAL(15, 2) NOT NULL,
    odds DECIMAL(10, 3) NOT NULL,
    potential_win DECIMAL(15, 2) NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    placed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    settled_at TIMESTAMP,
    FOREIGN KEY (event_id) REFERENCES sports_events(event_id),
    FOREIGN KEY (market_id) REFERENCES sports_markets(market_id)
);

CREATE INDEX idx_sports_bets_user_id ON sports_bets(user_id);
CREATE INDEX idx_sports_bets_status ON sports_bets(status);
CREATE INDEX idx_sports_bets_placed_at ON sports_bets(placed_at DESC);

-- Odds history table
CREATE TABLE IF NOT EXISTS sports_odds_history (
    id SERIAL PRIMARY KEY,
    market_id VARCHAR(255) NOT NULL,
    home_odds DECIMAL(10, 3),
    draw_odds DECIMAL(10, 3),
    away_odds DECIMAL(10, 3),
    recorded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (market_id) REFERENCES sports_markets(market_id)
);

CREATE INDEX idx_odds_history_market_id ON sports_odds_history(market_id);
CREATE INDEX idx_odds_history_recorded_at ON sports_odds_history(recorded_at DESC);
