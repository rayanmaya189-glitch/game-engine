-- Winners Showcase Service Database Schema

-- Winners table
CREATE TABLE IF NOT EXISTS winners (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    username VARCHAR(100) NOT NULL,
    win_amount DECIMAL(15, 2) NOT NULL,
    currency VARCHAR(10) NOT NULL DEFAULT 'USD',
    game_type VARCHAR(50) NOT NULL,
    game_name VARCHAR(200) NOT NULL,
    win_type VARCHAR(20) NOT NULL DEFAULT 'regular',
    multiplier DECIMAL(10, 2) DEFAULT 1.0,
    timestamp TIMESTAMP DEFAULT NOW(),
    display_on_feed BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_winners_timestamp ON winners(timestamp DESC);
CREATE INDEX idx_winners_win_amount ON winners(win_amount DESC);
CREATE INDEX idx_winners_win_type ON winners(win_type);
CREATE INDEX idx_winners_game_type ON winners(game_type);

-- Winner privacy settings table
CREATE TABLE IF NOT EXISTS winner_privacy_settings (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL UNIQUE,
    anonymize_name BOOLEAN DEFAULT true,
    show_on_leaderboard BOOLEAN DEFAULT true,
    show_on_jackpot_list BOOLEAN DEFAULT true,
    opt_out_of_showcase BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_privacy_user_id ON winner_privacy_settings(user_id);

-- Function to auto-categorize win type based on amount
CREATE OR REPLACE FUNCTION determine_win_type()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.win_amount >= 10000 THEN
        NEW.win_type := 'progressive';
    ELSIF NEW.win_amount >= 1000 THEN
        NEW.win_type := 'jackpot';
    ELSIF NEW.win_amount >= 100 THEN
        NEW.win_type := 'big';
    ELSE
        NEW.win_type := 'regular';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger for auto-categorization
CREATE TRIGGER trigger_determine_win_type
    BEFORE INSERT ON winners
    FOR EACH ROW
    EXECUTE FUNCTION determine_win_type();