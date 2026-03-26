-- Leaderboard Service Database Schema

-- Player scores table
CREATE TABLE IF NOT EXISTS player_scores (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    username VARCHAR(100) NOT NULL,
    score DECIMAL(15, 2) DEFAULT 0,
    wins INTEGER DEFAULT 0,
    win_amount DECIMAL(15, 2) DEFAULT 0,
    game_type VARCHAR(50) NOT NULL,
    period VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user_id, game_type, period)
);

CREATE INDEX idx_player_scores_user_id ON player_scores(user_id);
CREATE INDEX idx_player_scores_game_type ON player_scores(game_type);
CREATE INDEX idx_player_scores_period ON player_scores(period);
CREATE INDEX idx_player_scores_score ON player_scores(score DESC);

-- Leaderboard entries view (for querying ranked results)
CREATE OR REPLACE VIEW v_leaderboard_entries AS
SELECT 
    ROW_NUMBER() OVER (PARTITION BY period, game_type ORDER BY score DESC) as rank,
    user_id,
    username,
    score,
    wins,
    win_amount,
    game_type,
    period,
    updated_at
FROM player_scores;

-- Daily leaderboard entries
CREATE TABLE IF NOT EXISTS leaderboard_daily (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    username VARCHAR(100) NOT NULL,
    score DECIMAL(15, 2) DEFAULT 0,
    wins INTEGER DEFAULT 0,
    win_amount DECIMAL(15, 2) DEFAULT 0,
    game_type VARCHAR(50) NOT NULL,
    rank INTEGER,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(game_type, date, user_id)
);

CREATE INDEX idx_leaderboard_daily_date ON leaderboard_daily(date);
CREATE INDEX idx_leaderboard_daily_rank ON leaderboard_daily(rank);

-- Weekly leaderboard entries
CREATE TABLE IF NOT EXISTS leaderboard_weekly (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    username VARCHAR(100) NOT NULL,
    score DECIMAL(15, 2) DEFAULT 0,
    wins INTEGER DEFAULT 0,
    win_amount DECIMAL(15, 2) DEFAULT 0,
    game_type VARCHAR(50) NOT NULL,
    rank INTEGER,
    year INTEGER NOT NULL,
    week INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(game_type, year, week, user_id)
);

CREATE INDEX idx_leaderboard_weekly_year_week ON leaderboard_weekly(year, week);
CREATE INDEX idx_leaderboard_weekly_rank ON leaderboard_weekly(rank);

-- Monthly leaderboard entries
CREATE TABLE IF NOT EXISTS leaderboard_monthly (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    username VARCHAR(100) NOT NULL,
    score DECIMAL(15, 2) DEFAULT 0,
    wins INTEGER DEFAULT 0,
    win_amount DECIMAL(15, 2) DEFAULT 0,
    game_type VARCHAR(50) NOT NULL,
    rank INTEGER,
    year INTEGER NOT NULL,
    month INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(game_type, year, month, user_id)
);

CREATE INDEX idx_leaderboard_monthly_year_month ON leaderboard_monthly(year, month);
CREATE INDEX idx_leaderboard_monthly_rank ON leaderboard_monthly(rank);

-- All-time leaderboard entries
CREATE TABLE IF NOT EXISTS leaderboard_alltime (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    username VARCHAR(100) NOT NULL,
    score DECIMAL(15, 2) DEFAULT 0,
    wins INTEGER DEFAULT 0,
    win_amount DECIMAL(15, 2) DEFAULT 0,
    game_type VARCHAR(50) NOT NULL,
    rank INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(game_type, user_id)
);

CREATE INDEX idx_leaderboard_alltime_rank ON leaderboard_alltime(rank);

-- Function to update daily leaderboard
CREATE OR REPLACE FUNCTION update_daily_leaderboard()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO leaderboard_daily (user_id, username, score, wins, win_amount, game_type, date)
    VALUES (NEW.user_id, NEW.username, NEW.score, NEW.wins, NEW.win_amount, NEW.game_type, CURRENT_DATE)
    ON CONFLICT (game_type, date, user_id) 
    DO UPDATE SET 
        score = leaderboard_daily.score + NEW.score,
        wins = leaderboard_daily.wins + NEW.wins,
        win_amount = leaderboard_daily.win_amount + NEW.win_amount,
        updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to update daily leaderboard
DROP TRIGGER IF EXISTS trigger_update_daily_leaderboard ON player_scores;
CREATE TRIGGER trigger_update_daily_leaderboard
AFTER INSERT OR UPDATE ON player_scores
FOR EACH ROW
EXECUTE FUNCTION update_daily_leaderboard();

-- Function to update weekly leaderboard
CREATE OR REPLACE FUNCTION update_weekly_leaderboard()
RETURNS TRIGGER AS $$
DECLARE
    v_year INTEGER;
    v_week INTEGER;
BEGIN
    v_year := EXTRACT(YEAR FROM CURRENT_DATE);
    v_week := EXTRACT(WEEK FROM CURRENT_DATE);
    
    INSERT INTO leaderboard_weekly (user_id, username, score, wins, win_amount, game_type, year, week)
    VALUES (NEW.user_id, NEW.username, NEW.score, NEW.wins, NEW.win_amount, NEW.game_type, v_year, v_week)
    ON CONFLICT (game_type, year, week, user_id) 
    DO UPDATE SET 
        score = leaderboard_weekly.score + NEW.score,
        wins = leaderboard_weekly.wins + NEW.wins,
        win_amount = leaderboard_weekly.win_amount + NEW.win_amount,
        updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to update weekly leaderboard
DROP TRIGGER IF EXISTS trigger_update_weekly_leaderboard ON player_scores;
CREATE TRIGGER trigger_update_weekly_leaderboard
AFTER INSERT OR UPDATE ON player_scores
FOR EACH ROW
EXECUTE FUNCTION update_weekly_leaderboard();

-- Function to update monthly leaderboard
CREATE OR REPLACE FUNCTION update_monthly_leaderboard()
RETURNS TRIGGER AS $$
DECLARE
    v_year INTEGER;
    v_month INTEGER;
BEGIN
    v_year := EXTRACT(YEAR FROM CURRENT_DATE);
    v_month := EXTRACT(MONTH FROM CURRENT_DATE);
    
    INSERT INTO leaderboard_monthly (user_id, username, score, wins, win_amount, game_type, year, month)
    VALUES (NEW.user_id, NEW.username, NEW.score, NEW.wins, NEW.win_amount, NEW.game_type, v_year, v_month)
    ON CONFLICT (game_type, year, month, user_id) 
    DO UPDATE SET 
        score = leaderboard_monthly.score + NEW.score,
        wins = leaderboard_monthly.wins + NEW.wins,
        win_amount = leaderboard_monthly.win_amount + NEW.win_amount,
        updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to update monthly leaderboard
DROP TRIGGER IF EXISTS trigger_update_monthly_leaderboard ON player_scores;
CREATE TRIGGER trigger_update_monthly_leaderboard
AFTER INSERT OR UPDATE ON player_scores
FOR EACH ROW
EXECUTE FUNCTION update_monthly_leaderboard();

-- Function to update all-time leaderboard
CREATE OR REPLACE FUNCTION update_alltime_leaderboard()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO leaderboard_alltime (user_id, username, score, wins, win_amount, game_type)
    VALUES (NEW.user_id, NEW.username, NEW.score, NEW.wins, NEW.win_amount, NEW.game_type)
    ON CONFLICT (game_type, user_id) 
    DO UPDATE SET 
        score = leaderboard_alltime.score + NEW.score,
        wins = leaderboard_alltime.wins + NEW.wins,
        win_amount = leaderboard_alltime.win_amount + NEW.win_amount,
        updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to update all-time leaderboard
DROP TRIGGER IF EXISTS trigger_update_alltime_leaderboard ON player_scores;
CREATE TRIGGER trigger_update_alltime_leaderboard
AFTER INSERT OR UPDATE ON player_scores
FOR EACH ROW
EXECUTE FUNCTION update_alltime_leaderboard();

-- Leaderboard Prize Configuration (Database-managed, not YAML)
-- Prize configuration is now stored in database for tournament-based management

-- Prize templates table (reusable prize configurations)
CREATE TABLE IF NOT EXISTS prize_templates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    prize_type VARCHAR(20) NOT NULL, -- bonus, freespins, vip_points, merchandise
    value DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(10) DEFAULT 'USD',
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Tournament prize configurations (tournament-specific)
CREATE TABLE IF NOT EXISTS tournament_prize_configs (
    id SERIAL PRIMARY KEY,
    tournament_id VARCHAR(50) NOT NULL,
    leaderboard_type VARCHAR(20) NOT NULL, -- daily, weekly, monthly, tournament
    from_rank INTEGER NOT NULL,
    to_rank INTEGER NOT NULL,
    prize_template_id INTEGER REFERENCES prize_templates(id),
    prize_type VARCHAR(20) NOT NULL,
    value DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(10) DEFAULT 'USD',
    is_percentage BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(tournament_id, leaderboard_type, from_rank, to_rank)
);

-- Leaderboard prize history (track distributed prizes)
CREATE TABLE IF NOT EXISTS leaderboard_prize_distributions (
    id SERIAL PRIMARY KEY,
    leaderboard_type VARCHAR(20) NOT NULL,
    game_type VARCHAR(50),
    period VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    rank INTEGER NOT NULL,
    prize_type VARCHAR(20) NOT NULL,
    value DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(10) DEFAULT 'USD',
    distributed_at TIMESTAMP DEFAULT NOW(),
    status VARCHAR(20) DEFAULT 'pending', -- pending, distributed, failed
    transaction_id VARCHAR(100)
);

-- Indexes
CREATE INDEX idx_tournament_prize_tournament ON tournament_prize_configs(tournament_id);
CREATE INDEX idx_leaderboard_prize_period ON leaderboard_prize_distributions(leaderboard_type, period);
CREATE INDEX idx_leaderboard_prize_user ON leaderboard_prize_distributions(user_id);
CREATE INDEX idx_prize_templates_active ON prize_templates(active);

-- Default prize templates
INSERT INTO prize_templates (name, description, prize_type, value, currency, active) VALUES
('Daily 1st Place', 'First place prize for daily leaderboard', 'bonus', 100.00, 'USD', true),
('Daily 2nd Place', 'Second place prize for daily leaderboard', 'bonus', 50.00, 'USD', true),
('Daily 3rd Place', 'Third place prize for daily leaderboard', 'bonus', 25.00, 'USD', true),
('Daily Top 10', 'VIP points for ranks 4-10', 'vip_points', 100.00, 'USD', true),
('Weekly 1st Place', 'First place prize for weekly leaderboard', 'bonus', 500.00, 'USD', true),
('Weekly 2nd Place', 'Second place prize for weekly leaderboard', 'bonus', 250.00, 'USD', true),
('Weekly 3rd Place', 'Third place prize for weekly leaderboard', 'bonus', 100.00, 'USD', true),
('Weekly Top 10', 'Bonus for ranks 4-10', 'bonus', 50.00, 'USD', true),
('Monthly 1st Place', 'First place prize for monthly leaderboard', 'bonus', 2000.00, 'USD', true),
('Monthly 2nd Place', 'Second place prize for monthly leaderboard', 'bonus', 1000.00, 'USD', true),
('Monthly 3rd Place', 'Third place prize for monthly leaderboard', 'bonus', 500.00, 'USD', true),
('Monthly Top 10', 'Bonus for ranks 4-10', 'bonus', 200.00, 'USD', true);

-- Example: Tournament-specific prize configuration
INSERT INTO tournament_prize_configs (tournament_id, leaderboard_type, from_rank, to_rank, prize_type, value, currency) VALUES
('T001', 'tournament', 1, 1, 'bonus', 1000.00, 'USD'),
('T001', 'tournament', 2, 2, 'bonus', 500.00, 'USD'),
('T001', 'tournament', 3, 3, 'bonus', 250.00, 'USD'),
('T001', 'tournament', 4, 10, 'bonus', 100.00, 'USD'),
('T001', 'tournament', 11, 50, 'vip_points', 50.00, 'USD');
