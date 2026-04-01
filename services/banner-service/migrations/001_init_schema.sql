CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS banners (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image_url TEXT NOT NULL,
    click_url TEXT,
    banner_type VARCHAR(50) NOT NULL DEFAULT 'HERO',
    status VARCHAR(50) NOT NULL DEFAULT 'ACTIVE',
    priority INTEGER NOT NULL DEFAULT 0,
    width INTEGER NOT NULL DEFAULT 1920,
    height INTEGER NOT NULL DEFAULT 600,
    start_date TIMESTAMPTZ,
    end_date TIMESTAMPTZ,
    target_countries TEXT,
    target_vip_levels TEXT,
    target_game_types TEXT,
    click_count BIGINT NOT NULL DEFAULT 0,
    impression_count BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_banners_status ON banners(status);
CREATE INDEX idx_banners_type ON banners(banner_type);
CREATE INDEX idx_banners_priority ON banners(priority DESC);
CREATE INDEX idx_banners_dates ON banners(start_date, end_date);
CREATE INDEX idx_banners_created ON banners(created_at DESC);

CREATE TABLE IF NOT EXISTS announcements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    type VARCHAR(50) NOT NULL DEFAULT 'INFO',
    status VARCHAR(50) NOT NULL DEFAULT 'ACTIVE',
    priority INTEGER NOT NULL DEFAULT 0,
    start_date TIMESTAMPTZ,
    end_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_announcements_status ON announcements(status);
CREATE INDEX idx_announcements_priority ON announcements(priority DESC);
CREATE INDEX idx_announcements_dates ON announcements(start_date, end_date);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_banners_updated_at BEFORE UPDATE ON banners
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_announcements_updated_at BEFORE UPDATE ON announcements
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
