-- User Service Database Schema

-- Profiles table
CREATE TABLE IF NOT EXISTS profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(50),
    username VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR(255),
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    date_of_birth DATE,
    gender VARCHAR(20),
    avatar_url TEXT,
    country VARCHAR(2),
    language VARCHAR(10) DEFAULT 'en',
    currency VARCHAR(3) DEFAULT 'USD',
    timezone VARCHAR(50) DEFAULT 'UTC',
    status VARCHAR(20) DEFAULT 'STATUS_ACTIVE',
    kyc_level VARCHAR(30) DEFAULT 'KYC_LEVEL_NONE',
    email_verified BOOLEAN DEFAULT FALSE,
    phone_verified BOOLEAN DEFAULT FALSE,
    two_factor_enabled BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    last_login_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_profiles_user_id ON profiles(user_id);
CREATE INDEX idx_profiles_email ON profiles(email);
CREATE INDEX idx_profiles_username ON profiles(username);
CREATE INDEX idx_profiles_status ON profiles(status);
CREATE INDEX idx_profiles_kyc_level ON profiles(kyc_level);
CREATE INDEX idx_profiles_country ON profiles(country);
CREATE INDEX idx_profiles_created_at ON profiles(created_at);

-- Addresses table
CREATE TABLE IF NOT EXISTS addresses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    profile_id UUID NOT NULL REFERENCES profiles(id) ON DELETE CASCADE,
    street VARCHAR(255),
    city VARCHAR(100),
    state VARCHAR(100),
    postal_code VARCHAR(20),
    country VARCHAR(2),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(profile_id)
);

CREATE INDEX idx_addresses_profile_id ON addresses(profile_id);

-- KYC Status table
CREATE TABLE IF NOT EXISTS kyc_status (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(255) NOT NULL,
    status VARCHAR(30) DEFAULT 'VERIFICATION_STATUS_UNSPECIFIED',
    level VARCHAR(30) DEFAULT 'KYC_LEVEL_NONE',
    submitted_at TIMESTAMP WITH TIME ZONE,
    reviewed_at TIMESTAMP WITH TIME ZONE,
    rejection_reason TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_kyc_status_user_id ON kyc_status(user_id);
CREATE INDEX idx_kyc_status_status ON kyc_status(status);
CREATE INDEX idx_kyc_status_level ON kyc_status(level);

-- KYC Documents table
CREATE TABLE IF NOT EXISTS kyc_documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(255) NOT NULL,
    document_type VARCHAR(30) NOT NULL,
    document_number VARCHAR(100),
    document_data TEXT, -- Would be encrypted in production
    status VARCHAR(30) DEFAULT 'VERIFICATION_STATUS_PENDING',
    submitted_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    reviewed_at TIMESTAMP WITH TIME ZONE,
    reviewer_id VARCHAR(255),
    review_comment TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_kyc_documents_user_id ON kyc_documents(user_id);
CREATE INDEX idx_kyc_documents_status ON kyc_documents(status);
CREATE INDEX idx_kyc_documents_document_type ON kyc_documents(document_type);

-- Player Settings table
CREATE TABLE IF NOT EXISTS player_settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(255) UNIQUE NOT NULL,
    email_notifications BOOLEAN DEFAULT TRUE,
    sms_notifications BOOLEAN DEFAULT FALSE,
    push_notifications BOOLEAN DEFAULT TRUE,
    profile_public BOOLEAN DEFAULT FALSE,
    show_online_status BOOLEAN DEFAULT TRUE,
    auto_play BOOLEAN DEFAULT FALSE,
    sound_volume INTEGER DEFAULT 50,
    theme VARCHAR(50) DEFAULT 'default',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_player_settings_user_id ON player_settings(user_id);

-- Player Limits table (Responsible Gaming)
CREATE TABLE IF NOT EXISTS player_limits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(255) UNIQUE NOT NULL,
    daily_limit DECIMAL(15, 2) DEFAULT 10000,
    weekly_limit DECIMAL(15, 2) DEFAULT 50000,
    monthly_limit DECIMAL(15, 2) DEFAULT 100000,
    daily_loss_limit DECIMAL(15, 2),
    self_exclusion BOOLEAN DEFAULT FALSE,
    exclusion_end_date DATE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_player_limits_user_id ON player_limits(user_id);

-- Player Stats table
CREATE TABLE IF NOT EXISTS player_stats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(255) UNIQUE NOT NULL,
    total_deposits DECIMAL(15, 2) DEFAULT 0,
    total_withdrawals DECIMAL(15, 2) DEFAULT 0,
    total_bets DECIMAL(15, 2) DEFAULT 0,
    total_wins DECIMAL(15, 2) DEFAULT 0,
    total_bonuses DECIMAL(15, 2) DEFAULT 0,
    deposit_count INTEGER DEFAULT 0,
    withdrawal_count INTEGER DEFAULT 0,
    bet_count INTEGER DEFAULT 0,
    win_count INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_player_stats_user_id ON player_stats(user_id);

-- Player Status History table (for audit)
CREATE TABLE IF NOT EXISTS player_status_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(255) NOT NULL,
    old_status VARCHAR(20),
    new_status VARCHAR(20) NOT NULL,
    reason TEXT,
    changed_by VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_player_status_history_user_id ON player_status_history(user_id);
CREATE INDEX idx_player_status_history_created_at ON player_status_history(created_at);
