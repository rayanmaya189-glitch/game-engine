-- Jackpot Service Database Schema

-- Jackpots table
CREATE TABLE IF NOT EXISTS jackpots (
    id SERIAL PRIMARY KEY,
    jackpot_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    current_amount DECIMAL(15, 2) DEFAULT 0,
    min_bet DECIMAL(15, 2) DEFAULT 0,
    max_bet DECIMAL(15, 2) DEFAULT 0,
    status VARCHAR(50) DEFAULT 'inactive',
    starts_at TIMESTAMP,
    ends_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_jackpots_status ON jackpots(status);

-- Jackpot winners table
CREATE TABLE IF NOT EXISTS jackpot_winners (
    id SERIAL PRIMARY KEY,
    winner_id VARCHAR(255) UNIQUE NOT NULL,
    jackpot_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    won_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (jackpot_id) REFERENCES jackpots(jackpot_id)
);

CREATE INDEX idx_jackpot_winners_jackpot_id ON jackpot_winners(jackpot_id);

-- Jackpot participants table
CREATE TABLE IF NOT EXISTS jackpot_participants (
    id SERIAL PRIMARY KEY,
    jackpot_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    bet_amount DECIMAL(15, 2) NOT NULL,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(jackpot_id, user_id),
    FOREIGN KEY (jackpot_id) REFERENCES jackpots(jackpot_id)
);

CREATE INDEX idx_jackpot_participants_jackpot_id ON jackpot_participants(jackpot_id);
CREATE INDEX idx_jackpot_participants_user_id ON jackpot_participants(user_id);

-- Jackpot history table
CREATE TABLE IF NOT EXISTS jackpot_history (
    id SERIAL PRIMARY KEY,
    jackpot_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    result VARCHAR(50) NOT NULL,
    played_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (jackpot_id) REFERENCES jackpots(jackpot_id)
);

CREATE INDEX idx_jackpot_history_user_id ON jackpot_history(user_id);
