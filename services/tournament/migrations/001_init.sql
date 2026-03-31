-- Tournament Service Database Schema

-- Tournaments table: stores tournament configurations
CREATE TABLE IF NOT EXISTS tournaments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tournament_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    game_id VARCHAR(255),
    game_type VARCHAR(50) NOT NULL, -- SLOTS, POKER, BLACKJACK, etc.
    tournament_type VARCHAR(50) NOT NULL, -- LEADERBOARD, BRACKET, SIT_AND_GO
    status VARCHAR(50) DEFAULT 'DRAFT', -- DRAFT, SCHEDULED, ACTIVE, COMPLETED, CANCELLED
    entry_fee BIGINT DEFAULT 0,
    prize_pool BIGINT DEFAULT 0,
    currency VARCHAR(10) DEFAULT 'USD',
    min_participants INTEGER DEFAULT 2,
    max_participants INTEGER,
    current_participants INTEGER DEFAULT 0,
    scoring_type VARCHAR(50) DEFAULT 'TOTAL_WIN', -- TOTAL_WIN, HIGHEST_WIN, MOST_SPINS, etc.
    is_rebuy_allowed BOOLEAN DEFAULT FALSE,
    is_addon_allowed BOOLEAN DEFAULT FALSE,
    config JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    starts_at TIMESTAMP WITH TIME ZONE,
    ends_at TIMESTAMP WITH TIME ZONE,
    registration_deadline TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_tournaments_tournament_id ON tournaments(tournament_id);
CREATE INDEX idx_tournaments_game_id ON tournaments(game_id);
CREATE INDEX idx_tournaments_game_type ON tournaments(game_type);
CREATE INDEX idx_tournaments_tournament_type ON tournaments(tournament_type);
CREATE INDEX idx_tournaments_status ON tournaments(status);
CREATE INDEX idx_tournaments_starts_at ON tournaments(starts_at);
CREATE INDEX idx_tournaments_created_at ON tournaments(created_at);

-- Tournament Participants table: tracks players in tournaments
CREATE TABLE IF NOT EXISTS tournament_participants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tournament_id VARCHAR(255) NOT NULL REFERENCES tournaments(tournament_id) ON DELETE CASCADE,
    player_id VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'REGISTERED', -- REGISTERED, ACTIVE, ELIMINATED, WINNER
    score BIGINT DEFAULT 0,
    rank INTEGER,
    spins_count INTEGER DEFAULT 0,
    total_wagered BIGINT DEFAULT 0,
    total_won BIGINT DEFAULT 0,
    entry_paid BOOLEAN DEFAULT FALSE,
    rebuy_count INTEGER DEFAULT 0,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    eliminated_at TIMESTAMP WITH TIME ZONE,
    UNIQUE(tournament_id, player_id)
);

CREATE INDEX idx_tournament_participants_tournament_id ON tournament_participants(tournament_id);
CREATE INDEX idx_tournament_participants_player_id ON tournament_participants(player_id);
CREATE INDEX idx_tournament_participants_status ON tournament_participants(status);
CREATE INDEX idx_tournament_participants_score ON tournament_participants(score DESC);
CREATE INDEX idx_tournament_participants_rank ON tournament_participants(rank);

-- Tournament Prizes table: defines prize tiers and payouts
CREATE TABLE IF NOT EXISTS tournament_prizes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tournament_id VARCHAR(255) NOT NULL REFERENCES tournaments(tournament_id) ON DELETE CASCADE,
    place_from INTEGER NOT NULL,
    place_to INTEGER NOT NULL,
    prize_type VARCHAR(50) NOT NULL, -- FIXED, PERCENTAGE, TICKET
    amount BIGINT DEFAULT 0,
    percentage DECIMAL(5, 2),
    currency VARCHAR(10) DEFAULT 'USD',
    description VARCHAR(255),
    winner_id VARCHAR(255),
    claimed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    claimed_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_tournament_prizes_tournament_id ON tournament_prizes(tournament_id);
CREATE INDEX idx_tournament_prizes_winner_id ON tournament_prizes(winner_id);

-- Tournament Brackets table: for bracket-style tournaments
CREATE TABLE IF NOT EXISTS tournament_brackets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tournament_id VARCHAR(255) NOT NULL REFERENCES tournaments(tournament_id) ON DELETE CASCADE,
    round_number INTEGER NOT NULL,
    match_number INTEGER NOT NULL,
    player1_id VARCHAR(255),
    player2_id VARCHAR(255),
    player1_score BIGINT DEFAULT 0,
    player2_score BIGINT DEFAULT 0,
    winner_id VARCHAR(255),
    status VARCHAR(50) DEFAULT 'PENDING', -- PENDING, IN_PROGRESS, COMPLETED, BYE
    started_at TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_tournament_brackets_tournament_id ON tournament_brackets(tournament_id);
CREATE INDEX idx_tournament_brackets_round_number ON tournament_brackets(round_number);
CREATE INDEX idx_tournament_brackets_status ON tournament_brackets(status);

-- Trigger to update updated_at on tournaments changes
CREATE TRIGGER update_tournaments_updated_at BEFORE UPDATE ON tournaments
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update updated_at on tournament_participants changes
CREATE TRIGGER update_tournament_participants_updated_at BEFORE UPDATE ON tournament_participants
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update updated_at on tournament_prizes changes
CREATE TRIGGER update_tournament_prizes_updated_at BEFORE UPDATE ON tournament_prizes
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update updated_at on tournament_brackets changes
CREATE TRIGGER update_tournament_brackets_updated_at BEFORE UPDATE ON tournament_brackets
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
