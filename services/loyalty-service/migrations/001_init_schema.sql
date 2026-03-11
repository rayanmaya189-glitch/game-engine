-- Loyalty Service Database Schema

-- Loyalty members table
CREATE TABLE IF NOT EXISTS loyalty_members (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    points INTEGER DEFAULT 0,
    lifetime_points INTEGER DEFAULT 0,
    tier VARCHAR(50) DEFAULT 'bronze',
    status VARCHAR(50) DEFAULT 'active',
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_loyalty_members_user_id ON loyalty_members(user_id);
CREATE INDEX idx_loyalty_members_tier ON loyalty_members(tier);
CREATE INDEX idx_loyalty_members_lifetime_points ON loyalty_members(lifetime_points DESC);

-- Loyalty tiers table
CREATE TABLE IF NOT EXISTS loyalty_tiers (
    id SERIAL PRIMARY KEY,
    tier_id VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    min_points INTEGER DEFAULT 0,
    max_points INTEGER,
    points_multiplier DECIMAL(3, 2) DEFAULT 1.00,
    cashback_percent DECIMAL(5, 2) DEFAULT 0,
    benefits TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default tiers
INSERT INTO loyalty_tiers (tier_id, name, min_points, max_points, points_multiplier, cashback_percent, benefits) VALUES
('bronze', 'Bronze', 0, 999, 1.00, 0, '{"weekly_bonus": false, "birthday_bonus": 100, "priority_support": false}'),
('silver', 'Silver', 1000, 4999, 1.50, 1, '{"weekly_bonus": 5, "birthday_bonus": 200, "priority_support": false}'),
('gold', 'Gold', 5000, 19999, 2.00, 2, '{"weekly_bonus": 10, "birthday_bonus": 500, "priority_support": true}'),
('platinum', 'Platinum', 20000, 99999, 3.00, 3, '{"weekly_bonus": 15, "birthday_bonus": 1000, "priority_support": true, "personal_manager": false}'),
('diamond', 'Diamond', 100000, 499999, 5.00, 5, '{"weekly_bonus": 20, "birthday_bonus": 5000, "priority_support": true, "personal_manager": true}'),
('vip', 'VIP', 500000, NULL, 10.00, 10, '{"weekly_bonus": 25, "birthday_bonus": 10000, "priority_support": true, "personal_manager": true, "exclusive_events": true}');

-- Points transactions table
CREATE TABLE IF NOT EXISTS loyalty_points_transactions (
    id SERIAL PRIMARY KEY,
    transaction_id VARCHAR(255) UNIQUE NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    amount INTEGER NOT NULL,
    type VARCHAR(50) NOT NULL,
    source VARCHAR(100),
    reference_id VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES loyalty_members(user_id)
);

CREATE INDEX idx_points_transactions_user_id ON loyalty_points_transactions(user_id);
CREATE INDEX idx_points_transactions_created_at ON loyalty_points_transactions(created_at DESC);

-- Rewards table
CREATE TABLE IF NOT EXISTS loyalty_rewards (
    id SERIAL PRIMARY KEY,
    reward_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    points_cost INTEGER NOT NULL,
    type VARCHAR(50) NOT NULL,
    value DECIMAL(15, 2),
    status VARCHAR(50) DEFAULT 'active',
    expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_loyalty_rewards_status ON loyalty_rewards(status);

-- Insert sample rewards
INSERT INTO loyalty_rewards (reward_id, name, description, points_cost, type, value, status) VALUES
('bonus_10', '$10 Bonus', 'Get $10 bonus credits', 1000, 'bonus', 10.00, 'active'),
('bonus_50', '$50 Bonus', 'Get $50 bonus credits', 4500, 'bonus', 50.00, 'active'),
('bonus_100', '$100 Bonus', 'Get $100 bonus credits', 8500, 'bonus', 100.00, 'active'),
('freespins_10', '10 Free Spins', 'Get 10 free spins on any slot', 500, 'free_spin', 10, 'active'),
('freespins_50', '50 Free Spins', 'Get 50 free spins on any slot', 2000, 'free_spin', 50, 'active'),
('voucher_25', '$25 Voucher', 'Get $25 store voucher', 2500, 'voucher', 25.00, 'active');

-- Redemptions table
CREATE TABLE IF NOT EXISTS loyalty_redemptions (
    id SERIAL PRIMARY KEY,
    redemption_id VARCHAR(255) UNIQUE NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    reward_id VARCHAR(255) NOT NULL,
    points_spent INTEGER NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    redeemed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fulfilled_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES loyalty_members(user_id),
    FOREIGN KEY (reward_id) REFERENCES loyalty_rewards(reward_id)
);

CREATE INDEX idx_redemptions_user_id ON loyalty_redemptions(user_id);
CREATE INDEX idx_redemptions_status ON loyalty_redemptions(status);

-- Promotions table
CREATE TABLE IF NOT EXISTS loyalty_promotions (
    id SERIAL PRIMARY KEY,
    promotion_id VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    type VARCHAR(50) NOT NULL,
    points_bonus INTEGER DEFAULT 0,
    points_percent DECIMAL(5, 2) DEFAULT 0,
    min_deposit DECIMAL(15, 2) DEFAULT 0,
    max_bonus DECIMAL(15, 2),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    status VARCHAR(50) DEFAULT 'active',
    target_tiers TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_promotions_status ON loyalty_promotions(status);
CREATE INDEX idx_promotions_dates ON loyalty_promotions(start_date, end_date);
