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
