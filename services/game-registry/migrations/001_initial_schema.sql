-- Game Registry Service Database Schema
-- Version: 1.0.0

-- Create enum types
DO $$ BEGIN
    CREATE TYPE game_category AS ENUM (
        'UNSPECIFIED',
        'SLOTS',
        'TABLE_GAMES',
        'LIVE_CASINO',
        'SPORTS_BETTING',
        'ESPORTS',
        'LOTTERY',
        'BINGO',
        'POKER',
        'VIRTUAL_SPORTS',
        'INSTANT_WIN',
        'SCRATCH_CARDS',
        'KENO',
        'BLACKJACK',
        'ROULETTE',
        'BACCARAT',
        'CRAPS',
        'SIC_BO',
        'DRAGON_TIGER',
        'GAME_SHOWS',
        'OTHER'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE game_provider AS ENUM (
        'UNSPECIFIED',
        'INTERNAL',
        'PRAGMATIC_PLAY',
        'NETENT',
        'MICROGAMING',
        'EVOLUTION',
        'PLAYTECH',
        'BETSOFT',
        'BETGAMING',
        'AMATIC',
        'BELATRA',
        'EGT',
        'IGROSOFT',
        'PLAY_N_GO',
        'YGGDRASIL',
        'QUICKSPIN',
        'RELAX_GAMING',
        'THUNDERKICK',
        'ELK_STUDIOS',
        'NOLIMIT_CITY',
        'REDBEARD_GAMING'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE status AS ENUM (
        'UNSPECIFIED',
        'PENDING',
        'ACTIVE',
        'INACTIVE',
        'SUSPENDED',
        'DELETED',
        'BLOCKED',
        'ARCHIVED',
        'BANNED',
        'LOCKED',
        'EXPIRED',
        'CANCELLED',
        'COMPLETED',
        'FAILED'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE device_type AS ENUM (
        'UNSPECIFIED',
        'DESKTOP',
        'MOBILE',
        'TABLET',
        'TV',
        'WATCH',
        'VR',
        'CONSOLE'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
    CREATE TYPE game_language AS ENUM (
        'UNSPECIFIED',
        'EN',
        'TH',
        'VI',
        'ID',
        'MS',
        'ZH',
        'ZH_TW',
        'JA',
        'KO',
        'ES',
        'PT',
        'RU',
        'AR',
        'FR',
        'DE',
        'IT',
        'NL',
        'PL',
        'TR',
        'HI'
    );
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Game Categories table
CREATE TABLE IF NOT EXISTS game_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    icon_url VARCHAR(512),
    banner_url VARCHAR(512),
    parent_id UUID REFERENCES game_categories(id),
    sort_order INT DEFAULT 0,
    status status DEFAULT 'ACTIVE',
    is_featured BOOLEAN DEFAULT false,
    slug VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Game Providers table
CREATE TABLE IF NOT EXISTS game_providers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    logo_url VARCHAR(512),
    website_url VARCHAR(512),
    status status DEFAULT 'ACTIVE',
    games_count INT DEFAULT 0,
    license VARCHAR(255),
    established INT,
    is_featured BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Games table
CREATE TABLE IF NOT EXISTS games (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    provider_id UUID NOT NULL REFERENCES game_providers(id),
    category_id UUID NOT NULL REFERENCES game_categories(id),
    type game_category NOT NULL,
    status status DEFAULT 'ACTIVE',
    thumbnail_url VARCHAR(512),
    banner_url VARCHAR(512),
    rtp DECIMAL(5,2),
    volatility VARCHAR(20) DEFAULT 'MEDIUM',
    min_bet BIGINT DEFAULT 0,
    max_bet BIGINT DEFAULT 0,
    max_win VARCHAR(50),
    paylines INT DEFAULT 0,
    reels INT DEFAULT 0,
    features JSONB DEFAULT '[]',
    supported_devices JSONB DEFAULT '["DESKTOP", "MOBILE"]',
    supported_languages JSONB DEFAULT '["EN"]',
    supported_currencies JSONB DEFAULT '["USD"]',
    is_featured BOOLEAN DEFAULT false,
    is_new BOOLEAN DEFAULT false,
    is_popular BOOLEAN DEFAULT false,
    is_jackpot BOOLEAN DEFAULT false,
    launch_url VARCHAR(512),
    release_date DATE,
    popularity_score INT DEFAULT 0,
    sort_order INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Game Config table (for session management)
CREATE TABLE IF NOT EXISTS game_config (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    game_id UUID NOT NULL REFERENCES games(id),
    session_token VARCHAR(255) NOT NULL UNIQUE,
    game_url VARCHAR(512) NOT NULL,
    player_id VARCHAR(255) NOT NULL,
    balance BIGINT DEFAULT 0,
    currency VARCHAR(10) DEFAULT 'USD',
    language game_language DEFAULT 'EN',
    config_json JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Game Tags table (for search)
CREATE TABLE IF NOT EXISTS game_tags (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    game_id UUID NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    tag VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(game_id, tag)
);

-- Indexes for games
CREATE INDEX IF NOT EXISTS idx_games_provider_id ON games(provider_id);
CREATE INDEX IF NOT EXISTS idx_games_category_id ON games(category_id);
CREATE INDEX IF NOT EXISTS idx_games_type ON games(type);
CREATE INDEX IF NOT EXISTS idx_games_status ON games(status);
CREATE INDEX IF NOT EXISTS idx_games_is_featured ON games(is_featured);
CREATE INDEX IF NOT EXISTS idx_games_is_new ON games(is_new);
CREATE INDEX IF NOT EXISTS idx_games_is_popular ON games(is_popular);
CREATE INDEX IF NOT EXISTS idx_games_is_jackpot ON games(is_jackpot);
CREATE INDEX IF NOT EXISTS idx_games_popularity_score ON games(popularity_score DESC);
CREATE INDEX IF NOT EXISTS idx_games_sort_order ON games(sort_order);
CREATE INDEX IF NOT EXISTS idx_games_name_search ON games USING gin(name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_games_release_date ON games(release_date DESC);

-- Indexes for game_tags
CREATE INDEX IF NOT EXISTS idx_game_tags_game_id ON game_tags(game_id);
CREATE INDEX IF NOT EXISTS idx_game_tags_tag ON game_tags(tag);
CREATE INDEX IF NOT EXISTS idx_game_tags_tag_search ON game_tags USING gin(tag gin_trgm_ops);

-- Indexes for game_config
CREATE INDEX IF NOT EXISTS idx_game_config_game_id ON game_config(game_id);
CREATE INDEX IF NOT EXISTS idx_game_config_session_token ON game_config(session_token);
CREATE INDEX IF NOT EXISTS idx_game_config_expires_at ON game_config(expires_at);

-- Insert default categories
INSERT INTO game_categories (id, name, description, slug, sort_order, status, is_featured) VALUES
    (gen_random_uuid(), 'Slots', 'Slot machine games', 'slots', 1, 'ACTIVE', true),
    (gen_random_uuid(), 'Table Games', 'Classic table games', 'table-games', 2, 'ACTIVE', true),
    (gen_random_uuid(), 'Live Casino', 'Live dealer games', 'live-casino', 3, 'ACTIVE', true),
    (gen_random_uuid(), 'Blackjack', 'Blackjack variations', 'blackjack', 4, 'ACTIVE', false),
    (gen_random_uuid(), 'Roulette', 'Roulette variations', 'roulette', 5, 'ACTIVE', false),
    (gen_random_uuid(), 'Baccarat', 'Baccarat variations', 'baccarat', 6, 'ACTIVE', false),
    (gen_random_uuid(), 'Poker', 'Poker games', 'poker', 7, 'ACTIVE', false),
    (gen_random_uuid(), 'Dice', 'Dice games', 'dice', 8, 'ACTIVE', false),
    (gen_random_uuid(), 'Lottery', 'Lottery games', 'lottery', 9, 'ACTIVE', false),
    (gen_random_uuid(), 'Scratch Cards', 'Scratch card games', 'scratch-cards', 10, 'ACTIVE', false)
ON CONFLICT (slug) DO NOTHING;

-- Insert default providers
INSERT INTO game_providers (id, name, description, status, is_featured, established) VALUES
    (gen_random_uuid(), 'Internal', 'Internal games', 'ACTIVE', true, 2024),
    (gen_random_uuid(), 'Pragmatic Play', 'Leading slot provider', 'ACTIVE', true, 2015),
    (gen_random_uuid(), 'NetEnt', 'Premium casino games', 'ACTIVE', true, 1996),
    (gen_random_uuid(), 'Microgaming', 'Jackpot specialists', 'ACTIVE', true, 1994),
    (gen_random_uuid(), 'Evolution', 'Live casino leader', 'ACTIVE', true, 2006),
    (gen_random_uuid(), 'Playtech', 'Omni-channel gaming', 'ACTIVE', true, 1999),
    (gen_random_uuid(), 'Yggdrasil', 'Innovative slots', 'ACTIVE', true, 2013),
    (gen_random_uuid(), 'Thunderkick', 'Unique slot games', 'ACTIVE', false, 2012),
    (gen_random_uuid(), 'NoLimit City', 'High volatility slots', 'ACTIVE', false, 2014),
    (gen_random_uuid(), 'Red Tiger', 'Daily jackpots', 'ACTIVE', true, 2014)
ON CONFLICT DO NOTHING;

-- Create trigger for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_game_categories_updated_at BEFORE UPDATE ON game_categories
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_game_providers_updated_at BEFORE UPDATE ON game_providers
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_games_updated_at BEFORE UPDATE ON games
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
