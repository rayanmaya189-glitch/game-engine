CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE dealer_tables (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    table_name VARCHAR(100) NOT NULL,
    game_type VARCHAR(50) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    min_bet DECIMAL(15, 2) NOT NULL DEFAULT 1.00,
    max_bet DECIMAL(15, 2) NOT NULL DEFAULT 10000.00,
    max_seats INTEGER NOT NULL DEFAULT 7,
    current_seats INTEGER NOT NULL DEFAULT 0,
    stream_url VARCHAR(500),
    dealer_id UUID,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_dealer_tables_game_type ON dealer_tables (game_type);
CREATE INDEX idx_dealer_tables_status ON dealer_tables (status);
CREATE INDEX idx_dealer_tables_dealer_id ON dealer_tables (dealer_id);

CREATE TABLE dealer_sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    table_id UUID NOT NULL REFERENCES dealer_tables (id),
    dealer_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    started_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    ended_at TIMESTAMP WITH TIME ZONE,
    total_rounds INTEGER NOT NULL DEFAULT 0,
    total_wagered DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    total_won DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    metadata JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_dealer_sessions_table_id ON dealer_sessions (table_id);
CREATE INDEX idx_dealer_sessions_dealer_id ON dealer_sessions (dealer_id);
CREATE INDEX idx_dealer_sessions_status ON dealer_sessions (status);
CREATE INDEX idx_dealer_sessions_started_at ON dealer_sessions (started_at);

CREATE TABLE dealer_players (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    session_id UUID NOT NULL REFERENCES dealer_sessions (id),
    user_id UUID NOT NULL,
    seat_number INTEGER NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    total_wagered DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    total_won DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    joined_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    left_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_dealer_players_session_id ON dealer_players (session_id);
CREATE INDEX idx_dealer_players_user_id ON dealer_players (user_id);
CREATE INDEX idx_dealer_players_status ON dealer_players (status);

CREATE TABLE dealer_rounds (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    session_id UUID NOT NULL REFERENCES dealer_sessions (id),
    table_id UUID NOT NULL REFERENCES dealer_tables (id),
    round_number INTEGER NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    result JSONB,
    started_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    ended_at TIMESTAMP WITH TIME ZONE,
    total_wagered DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    total_won DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_dealer_rounds_session_id ON dealer_rounds (session_id);
CREATE INDEX idx_dealer_rounds_table_id ON dealer_rounds (table_id);
CREATE INDEX idx_dealer_rounds_status ON dealer_rounds (status);
CREATE INDEX idx_dealer_rounds_started_at ON dealer_rounds (started_at);
