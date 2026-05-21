#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Royal Platform Data Seeder${NC}"
echo -e "${BLUE}========================================${NC}"
echo ""

print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Load environment variables
if [ -f "$ROOT_DIR/.env" ]; then
    export $(cat "$ROOT_DIR/.env" | grep -v '^#' | xargs)
fi

POSTGRES_HOST="${POSTGRES_HOST:-localhost}"
POSTGRES_PORT="${POSTGRES_PORT:-5432}"
POSTGRES_USER="${POSTGRES_USER:-royal_platform}"
POSTGRES_PASSWORD="${POSTGRES_PASSWORD:-royal_platform_password}"
POSTGRES_DB="${POSTGRES_DB:-royal_platform}"

print_info "Seeding databases with test data..."

# Seed auth-service
print_info "Seeding auth-service..."
PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
    -U "$POSTGRES_USER" -d "${POSTGRES_DB}_auth-service" <<EOF
-- Insert test admin user (password: Admin123!)
INSERT INTO users (id, email, username, password_hash, role, email_verified, status, created_at)
VALUES 
    ('00000000-0000-0000-0000-000000000001', 'admin@royalplatform.com', 'admin', 
     '\$2a\$10\$XQxNqQKjQZzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8', 
     'admin', true, 'active', NOW()),
    ('00000000-0000-0000-0000-000000000002', 'player1@test.com', 'player1', 
     '\$2a\$10\$XQxNqQKjQZzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8', 
     'user', true, 'active', NOW()),
    ('00000000-0000-0000-0000-000000000003', 'player2@test.com', 'player2', 
     '\$2a\$10\$XQxNqQKjQZzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8yFzJ8', 
     'user', true, 'active', NOW())
ON CONFLICT (id) DO NOTHING;
EOF
print_success "Auth-service seeded"

# Seed wallet-service
print_info "Seeding wallet-service..."
PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
    -U "$POSTGRES_USER" -d "${POSTGRES_DB}_wallet-service" <<EOF
-- Create wallets for test users
INSERT INTO wallets (id, user_id, currency, wallet_type, balance, locked_balance, status, created_at)
VALUES 
    ('10000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000002', 'USD', 'main', 1000000, 0, 'active', NOW()),
    ('10000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000002', 'USD', 'bonus', 50000, 0, 'active', NOW()),
    ('10000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000003', 'USD', 'main', 500000, 0, 'active', NOW()),
    ('10000000-0000-0000-0000-000000000004', '00000000-0000-0000-0000-000000000003', 'EUR', 'main', 750000, 0, 'active', NOW())
ON CONFLICT (id) DO NOTHING;
EOF
print_success "Wallet-service seeded"

# Seed game-registry-service
print_info "Seeding game-registry-service..."
PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
    -U "$POSTGRES_USER" -d "${POSTGRES_DB}_game-registry-service" <<EOF
-- Insert sample games
INSERT INTO games (id, name, type, provider, min_bet, max_bet, rtp, enabled, created_at)
VALUES 
    ('game-0000-0000-0000-0000-000000000001', 'Blackjack Classic', 'card', 'internal', 100, 100000, 99.5, true, NOW()),
    ('game-0000-0000-0000-0000-000000000002', 'Baccarat Royale', 'card', 'internal', 500, 500000, 98.9, true, NOW()),
    ('game-0000-0000-0000-0000-000000000003', 'Texas Hold''em Poker', 'card', 'internal', 200, 200000, 97.5, true, NOW()),
    ('game-0000-0000-0000-0000-000000000004', 'Andar Bahar', 'card', 'internal', 100, 100000, 97.0, true, NOW()),
    ('game-0000-0000-0000-0000-000000000005', 'Dragon Tiger', 'card', 'internal', 100, 100000, 96.5, true, NOW()),
    ('game-0000-0000-0000-0000-000000000006', 'Craps Master', 'dice', 'internal', 200, 200000, 98.0, true, NOW()),
    ('game-0000-0000-0000-0000-000000000007', 'Sic Bo Deluxe', 'dice', 'internal', 100, 100000, 97.5, true, NOW()),
    ('game-0000-0000-0000-0000-000000000008', 'Royal Slots', 'slot', 'internal', 50, 50000, 96.0, true, NOW()),
    ('game-0000-0000-0000-0000-000000000009', 'Megaways Fortune', 'slot', 'internal', 100, 100000, 96.5, true, NOW()),
    ('game-0000-0000-0000-0000-000000000010', 'Sports Betting', 'sports', 'internal', 100, 1000000, 95.0, true, NOW())
ON CONFLICT (id) DO NOTHING;
EOF
print_success "Game-registry-service seeded"

# Seed tournament-service
print_info "Seeding tournament-service..."
PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
    -U "$POSTGRES_USER" -d "${POSTGRES_DB}_tournament-service" <<EOF
-- Create sample tournaments
INSERT INTO tournaments (id, name, game_type, entry_fee, prize_pool, max_players, start_time, status, created_at)
VALUES 
    ('tour-0000-0000-0000-0000-000000000001', 'Daily Blackjack Championship', 'blackjack', 10000, 100000, 100, NOW() + INTERVAL '1 hour', 'scheduled', NOW()),
    ('tour-0000-0000-0000-0000-000000000002', 'Weekly Poker Tournament', 'poker', 50000, 500000, 500, NOW() + INTERVAL '24 hours', 'scheduled', NOW()),
    ('tour-0000-0000-0000-0000-000000000003', 'Weekend Slots Race', 'slots', 5000, 50000, 1000, NOW() + INTERVAL '48 hours', 'scheduled', NOW())
ON CONFLICT (id) DO NOTHING;
EOF
print_success "Tournament-service seeded"

# Seed leaderboard-service (Redis would be seeded via script, but we'll note it)
print_warning "Leaderboard data should be seeded in Redis separately"

# Seed loyalty-service
print_info "Seeding loyalty-service..."
PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
    -U "$POSTGRES_USER" -d "${POSTGRES_DB}_loyalty-service" <<EOF
-- Create VIP levels
INSERT INTO vip_levels (id, name, min_points, max_points, benefits, cashback_percentage, created_at)
VALUES 
    ('vip-0000-0000-0000-0000-000000000001', 'Bronze', 0, 9999, '{"bonus": 0, "cashback": 0}', 0.0, NOW()),
    ('vip-0000-0000-0000-0000-000000000002', 'Silver', 10000, 49999, '{"bonus": 100, "cashback": 1}', 1.0, NOW()),
    ('vip-0000-0000-0000-0000-000000000003', 'Gold', 50000, 199999, '{"bonus": 500, "cashback": 2}', 2.0, NOW()),
    ('vip-0000-0000-0000-0000-000000000004', 'Platinum', 200000, 999999, '{"bonus": 2000, "cashback": 3}', 3.0, NOW()),
    ('vip-0000-0000-0000-0000-000000000005', 'Diamond', 1000000, 999999999, '{"bonus": 10000, "cashback": 5}', 5.0, NOW())
ON CONFLICT (id) DO NOTHING;
EOF
print_success "Loyalty-service seeded"

# Seed banner-service
print_info "Seeding banner-service..."
PGPASSWORD="$POSTGRES_PASSWORD" psql -h "$POSTGRES_HOST" -p "$POSTGRES_PORT" \
    -U "$POSTGRES_USER" -d "${POSTGRES_DB}_banner-service" <<EOF
-- Create welcome banners
INSERT INTO banners (id, title, image_url, target_url, position, priority, enabled, start_date, end_date, created_at)
VALUES 
    ('banner-0000-0000-0000-0000-000000000001', 'Welcome Bonus 100%', '/images/welcome-bonus.jpg', '/promotions/welcome', 'hero', 1, true, NOW(), NOW() + INTERVAL '30 days', NOW()),
    ('banner-0000-0000-0000-0000-000000000002', 'Daily Blackjack Tournament', '/images/blackjack-tourney.jpg', '/tournaments/blackjack', 'sidebar', 2, true, NOW(), NOW() + INTERVAL '7 days', NOW()),
    ('banner-0000-0000-0000-0000-000000000003', 'New Slots Released!', '/images/new-slots.jpg', '/games/slots', 'footer', 3, true, NOW(), NOW() + INTERVAL '14 days', NOW())
ON CONFLICT (id) DO NOTHING;
EOF
print_success "Banner-service seeded"

print_success ""
print_success "=========================================="
print_success "  Database seeding completed successfully!"
print_success "=========================================="
print_success ""
print_info "Test credentials:"
print_info "  Admin: admin@royalplatform.com / Admin123!"
print_info "  Player 1: player1@test.com / Player123!"
print_info "  Player 2: player2@test.com / Player123!"
print_success ""

exit 0
