# Referral Service - Phase 6 Platform Benefits

## Overview
The Referral Service enables player-to-player referral programs where existing players can refer new players and earn rewards.

## Implementation Status: Not Started

### Required Features
1. **Referral Links** - Generate unique referral codes/links
2. **Referral Tracking** - Track referrals from signup to first deposit
3. **Reward System** - Configurable rewards for referrer and referred
4. **Multi-tier Referrals** - Support for sub-affiliates
5. **Referral Dashboard** - Player view of referral history

### Database Schema (planned)
```sql
-- Referrals table
CREATE TABLE referrals (
    id SERIAL PRIMARY KEY,
    referrer_id VARCHAR(50) NOT NULL,
    referee_id VARCHAR(50),
    referral_code VARCHAR(50) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    reward_amount DECIMAL(15, 2),
    reward_claimed BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW(),
    claimed_at TIMESTAMP
);

-- Referral rewards configuration
CREATE TABLE referral_rewards (
    id SERIAL PRIMARY KEY,
    reward_type VARCHAR(20) NOT NULL,
    referrer_bonus DECIMAL(15, 2),
    referee_bonus DECIMAL(15, 2),
    min_deposit DECIMAL(15, 2),
    min_bet DECIMAL(15, 2),
    active BOOLEAN DEFAULT true
);
```

### API Endpoints (planned)
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/v1/referrals/code` | GET | Get player's referral code |
| `/api/v1/referrals/code` | POST | Generate new referral code |
| `/api/v1/referrals/track` | POST | Track a referral (internal) |
| `/api/v1/referrals/history` | GET | Get player's referral history |
| `/api/v1/referrals/stats` | GET | Get referral statistics |
| `/api/v1/referrals/rewards` | GET | Get available rewards |
| `/api/v1/referrals/claim` | POST | Claim referral reward |

### Reward Types
- Deposit Bonus - Percentage of first deposit
- Free Spins - Slot game free spins
- VIP Points - Loyalty points
- Cash - Direct money credit

### Implementation Notes
- Use unique referral codes (e.g., REF123456)
- Track referral from registration through to qualifying action
- Support multi-tier referrals (referrer gets cut of sub-referrals)
- Integrate with Bonus Service for reward distribution
- GDPR compliant - clear consent for tracking