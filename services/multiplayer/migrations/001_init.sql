-- Multiplayer Service Database Schema

-- Rooms table: stores game rooms
CREATE TABLE IF NOT EXISTS rooms (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    game_type VARCHAR(50) NOT NULL, -- POKER, BLACKJACK, etc.
    room_type VARCHAR(50) DEFAULT 'PUBLIC', -- PUBLIC, PRIVATE, TOURNAMENT
    status VARCHAR(50) DEFAULT 'WAITING', -- WAITING, ACTIVE, FULL, CLOSED
    host_id VARCHAR(255) NOT NULL,
    max_players INTEGER DEFAULT 4,
    current_players INTEGER DEFAULT 0,
    min_bet BIGINT DEFAULT 0,
    max_bet BIGINT DEFAULT 0,
    currency VARCHAR(10) DEFAULT 'USD',
    password_hash VARCHAR(255),
    config JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    started_at TIMESTAMP WITH TIME ZONE,
    ended_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_rooms_room_id ON rooms(room_id);
CREATE INDEX idx_rooms_game_type ON rooms(game_type);
CREATE INDEX idx_rooms_room_type ON rooms(room_type);
CREATE INDEX idx_rooms_status ON rooms(status);
CREATE INDEX idx_rooms_host_id ON rooms(host_id);
CREATE INDEX idx_rooms_created_at ON rooms(created_at);

-- Room Players table: tracks players in rooms
CREATE TABLE IF NOT EXISTS room_players (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_id VARCHAR(255) NOT NULL REFERENCES rooms(room_id) ON DELETE CASCADE,
    player_id VARCHAR(255) NOT NULL,
    seat_number INTEGER,
    status VARCHAR(50) DEFAULT 'JOINED', -- JOINED, READY, PLAYING, LEFT, KICKED
    balance BIGINT DEFAULT 0,
    score BIGINT DEFAULT 0,
    is_host BOOLEAN DEFAULT FALSE,
    is_spectator BOOLEAN DEFAULT FALSE,
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    left_at TIMESTAMP WITH TIME ZONE,
    UNIQUE(room_id, player_id)
);

CREATE INDEX idx_room_players_room_id ON room_players(room_id);
CREATE INDEX idx_room_players_player_id ON room_players(player_id);
CREATE INDEX idx_room_players_status ON room_players(status);
CREATE INDEX idx_room_players_joined_at ON room_players(joined_at);

-- Matchmaking Queue table: stores players waiting for matchmaking
CREATE TABLE IF NOT EXISTS matchmaking_queue (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id VARCHAR(255) NOT NULL,
    game_type VARCHAR(50) NOT NULL,
    skill_level VARCHAR(50),
    min_bet BIGINT DEFAULT 0,
    max_bet BIGINT DEFAULT 0,
    currency VARCHAR(10) DEFAULT 'USD',
    status VARCHAR(50) DEFAULT 'QUEUED', -- QUEUED, MATCHED, CANCELLED, EXPIRED
    priority INTEGER DEFAULT 0,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    matched_at TIMESTAMP WITH TIME ZONE,
    expires_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_matchmaking_queue_player_id ON matchmaking_queue(player_id);
CREATE INDEX idx_matchmaking_queue_game_type ON matchmaking_queue(game_type);
CREATE INDEX idx_matchmaking_queue_status ON matchmaking_queue(status);
CREATE INDEX idx_matchmaking_queue_priority ON matchmaking_queue(priority DESC);
CREATE INDEX idx_matchmaking_queue_created_at ON matchmaking_queue(created_at);

-- Trigger to update updated_at on rooms changes
CREATE TRIGGER update_rooms_updated_at BEFORE UPDATE ON rooms
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update updated_at on room_players changes
CREATE TRIGGER update_room_players_updated_at BEFORE UPDATE ON room_players
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to update updated_at on matchmaking_queue changes
CREATE TRIGGER update_matchmaking_queue_updated_at BEFORE UPDATE ON matchmaking_queue
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
