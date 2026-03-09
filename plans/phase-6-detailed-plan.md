# Phase 6: Platform Benefits & Engagement - Detailed Implementation Plan

## Overview

Phase 6 implements player engagement and retention features: Leaderboards, Winners Showcase, Banner & Announcement system, Commission/Revenue Share for agents, Affiliate Program, Loyalty & VIP Program, and Referral Program.

**Prerequisites**: Phase 5 complete (Mobile Apps).

---

## 1. Leaderboard Service (Golang Kratos)

### 1.1 Project Setup
- Kratos project, gRPC (port 9018), PostgreSQL (casino_engagement), Redis

### 1.2 Leaderboard Types

| Type | Period | Update Frequency | Prize |
|------|--------|-----------------|-------|
| Daily Winners | 00:00-23:59 UTC | Real-time | Points/Credits |
| Weekly Winners | Monday 00:00 - Sunday 23:59 UTC | Real-time | Bonus + Trophies |
| Monthly Winners | 1st - Last day of month | Real-time | Bonus + VIP Points |
| All-Time Winners | Since platform launch | Hourly | Status Tiers |
| Biggest Win | Rolling 24 hours | Real-time | - |
| Most Active | Rolling 24 hours | Real-time | - |
| Game-Specific | Per game, configurable period | Real-time | - |
| Tournament | During tournament | Real-time | Tournament Prizes |
| VIP Points | Rolling 30 days | Real-time | VIP Tier Advancement |

### 1.3 Redis Implementation
- Sorted sets per leaderboard: `leaderboard:{type}:{period}` → score = metric value, member = player_id
- Near real-time updates via WebSocket broadcast
- Historical snapshots: Daily dump to PostgreSQL for historical queries

### 1.4 Anti-Gaming Rules
- Minimum bet threshold (e.g., $0.50 per spin/hand) to qualify
- Unique player validation (no bot accounts)
- Automatic flagging of suspicious patterns
- Manual review queue for edge cases

### 1.5 Prize Distribution
- Configurable prize pools per leaderboard
- Auto-credit: Top N players receive prizes at period end
- Prizes: Bonus money, free spins, merchandise (future), VIP points

---

## 2. Winners Showcase Service (NestJS)

### 2.1 Project Setup
- NestJS project, gRPC (port 9019), PostgreSQL, Redis, WebSocket

### 2.2 Features

#### Recent Winners Feed
- Real-time stream of recent wins across all games
- Filtered by: game type, amount threshold, time window
- Data: Player (anonymized), game, amount, timestamp
- Delivered via WebSocket and REST API for lobby ticker

#### Big Win Highlights
- Curated list of notable wins above configurable threshold
- Threshold tiers: $100+, $1,000+, $10,000+
- Includes: Player name (anonymized), game, amount, win multiplier
- Featured on homepage banner

#### Jackpot Winners
- Dedicated celebration for jackpot wins
- Animation trigger on mobile/web
- Announced to all connected players via WebSocket
- Optional: Winner testimonials (with consent)

#### Privacy Controls
- Player preference: Anonymize name (J***n W.) or use username
- Opt-out of winner showcase entirely
- GDPR compliant: Clear consent workflow

### 2.3 Display Formats
- Scrolling ticker in lobby (lobby screen)
- Widget on homepage
- Dedicated "Winners" page
- Mobile push notification (opt-in) for big wins

---

## 3. Banner & Announcement Service (NestJS)

### 3.1 Project Setup
- NestJS project, gRPC (port 9020), PostgreSQL, S3 (assets), Redis

### 3.2 Banner Types

| Type | Placement | Size | Click Action |
|------|-----------|------|--------------|
| Hero Banner | Homepage top | 1920x600 | Deep link to game/promo |
| Sidebar Banner | Desktop sidebar | 300x600 | Deep link |
| Popup Banner | Modal on entry | 800x600 | Action or dismiss |
| In-Game Banner | Between rounds | 728x90 | Context-aware |
| Interstitial | Full screen transition | Full screen | Action or skip |
| Notification Bar | Top of screen | Full width | Click to expand |

### 3.3 Targeting Rules
- **Player segment**: VIP level, deposit history, game preference, country
- **Device**: Android, iOS, Web, Desktop
- **Geography**: Country, language
- **Time**: Start/end datetime, timezone
- **Frequency**: Impression cap per user per period
- **A/B testing**: Multiple variants with traffic split

### 3.4 Content Management
- WYSIWYG banner editor
- Image upload to S3 with CDN distribution
- Template system for quick creation
- Preview mode (desktop/mobile)

### 3.5 Analytics
- Impression count, unique impressions
- Click-through rate (CTR)
- Conversion rate (deposit/bonus claimed after click)
- A/B test winner determination

---

## 4. Commission & Revenue Share Service (Java Spring Boot)

### 4.1 Project Setup
- Spring Boot project, gRPC (port 9021), PostgreSQL (casino_commissions), Redis

### 4.2 Commission Models

#### Revenue Share (RevShare)
- Percentage of Net Gaming Revenue (NGR) from referred players
- NGR = (Bets - Wins - Bonuses - Fees)
- Tiered rates: 20-40% based on player volume
- Negative carryover: Option to carry negative balance to next month

#### CPA (Cost Per Acquisition)
- Fixed payment per qualified new player
- Qualification: Minimum deposit + minimum wager
- Payment: One-time per player

#### Hybrid
- Combination of CPA + RevShare
- Lower RevShare percentage + CPA for each new player

#### Sub-Affiliate Commission
- Multi-tier: 5-10% of sub-affiliate's commission
- Configurable depth: 2-5 levels

### 4.3 Configuration
- Global default commission plan
- Per-agent/affiliate commission overrides
- Per-game commission rates (poker vs slots vs sports)
- Currency-specific rates

### 4.4 Settlement
- Calculation periods: Weekly, bi-weekly, monthly
- Minimum payout threshold: $50
- Payment methods: Bank transfer, crypto, e-wallet
- Auto-generated invoices

### 4.5 Reporting
- Real-time commission dashboard
- Breakdown by: player, game, period
- Historical comparison
- Export to CSV/PDF

---

## 5. Affiliate Program Service (Java Spring Boot)

### 5.1 Project Setup
- Spring Boot project, gRPC (port 9022), PostgreSQL (casino_affiliates), Redis

### 5.2 Affiliate Features

#### Self-Service Portal
- Registration form with business info
- Approval workflow (pending → approved/rejected)
- Dashboard: Clicks, registrations, deposits, revenue
- Marketing tools: Links, banners, QR codes

#### Tracking System
- Unique tracking links with UTM parameters
- Click tracking with attribution window (30/60/90 days)
- Registration, FTD (first deposit), LTV tracking
- Multi-touch attribution (first click, last click)

#### Sub-Affiliate Program
- Affiliate can refer other affiliates
- Multi-level referral (configurable depth)
- Tier-specific commission percentages

#### Compliance
- Affiliate agreement acceptance
- Prohibited marketing method enforcement
- Traffic quality monitoring
- Geo-restriction enforcement

### 5.3 Affiliate Tiers
| Tier | Monthly NGR Required | Commission Boost |
|------|---------------------|------------------|
| Bronze | $0 | 0% |
| Silver | $5,000 | +5% |
| Gold | $20,000 | +10% |
| Platinum | $50,000 | +15% |

---

## 6. Loyalty & VIP Program Service (Golang Kratos)

### 6.1 Project Setup
- Kratos project, gRPC (port 9023), PostgreSQL (casino_engagement), Redis

### 6.2 Loyalty Points System

#### Earning Points
- Points per wager: 1 point per $1 wagered (slots)
- Game weights: Slots 100%, Blackjack 10%, Roulette 20%, Poker 5%
- Bonus point events: Double points weekends, birthday
- VIP tier multipliers

#### Redeeming Points
- Convert to bonus money (100 points = $1)
- Free spins (specific games)
- Merchandise (future)
- Tournament entries
- Exclusive experiences

#### Expiry
- Points expire after 12 months of inactivity
- Grace period for expiring points notification

### 6.3 VIP Tiers

| Tier | Points Required | Monthly Requirements | Benefits |
|------|----------------|---------------------|----------|
| Bronze | 0 | - | Basic access, welcome offers |
| Silver | 1,000 | $500 | 5% cashback, priority support |
| Gold | 10,000 | $5,000 | 10% cashback, personal manager |
| Platinum | 50,000 | $25,000 | 15% cashback, exclusive games, higher limits |
| Diamond | 200,000 | $100,000 | 20% cashback, custom bonuses, VIP events |

### 6.4 VIP Benefits
- Enhanced deposit/withdrawal limits
- Faster withdrawal processing (24h → 4h)
- Exclusive game access
- Birthday bonuses
- Personal account manager (Platinum+)
- Exclusive tournament invitations
- Custom bonus offers
- Luxury gifts (seasonal)

### 6.5 Automated Management
- Auto-upgrade when points/activity threshold met
- Grace period before downgrade (3 months)
- Manual override by admin for special cases
- VIP event invitations (seasonal tournaments, parties)

---

## 7. Referral Program Service (NestJS)

### 7.1 Project Setup
- NestJS project, gRPC (port 9024), PostgreSQL (casino_referrals), Redis

### 7.2 Referral Mechanics

#### Referral Code
- Each player gets unique referral code
- Shareable via: Social media, messaging, email
- Deep link: `myapp://referral/{code}`
- Web link: `https://casino.com/ref/{code}`

#### Reward Structure
- **Referrer reward**: $50 bonus when referee makes first deposit of $50+
- **Referee reward**: 100% first deposit match up to $100
- **Milestone rewards**:
  - Refer 5 players: $100 bonus
  - Refer 10 players: $250 bonus
  - Refer 25 players: $750 bonus
- **Ongoing**: 5% of referee's wagering (for 90 days)

### 7.3 Tracking
- Funnel tracking: Click → Register → Verify → Deposit → Active
- Attribution: First-touch and last-touch
- Conversion rate analytics

### 7.4 Anti-Abuse
- Self-referral detection (same device/IP)
- Multi-account detection
- Minimum activity requirements before reward
- Referral code sharing restrictions

---

## 8. Admin Panel Extensions

### 8.1 Leaderboard Management
- Create/configure leaderboard: Type, period, prize pool, min bet
- Monitor active leaderboards
- Manual prize adjustment
- Historical leaderboard queries

### 8.2 Winners Management
- Configure win thresholds for showcase
- Approve/modify winner displays
- Winner testimonial management
- Privacy settings configuration

### 8.3 Banner & Announcement Management
- Banner creation workflow
- Targeting rule builder
- A/B test configuration
- Analytics dashboard

### 8.4 Commission Management
- Commission plan configuration
- Tier rate setup
- Settlement approval workflow
- Commission reports

### 8.5 Affiliate Management
- Affiliate approval queue
- Tracking link generator
- Traffic quality dashboard
- Sub-affiliate tree view

### 8.6 Loyalty & VIP Management
- Points/earnings rate configuration
- VIP tier threshold setup
- VIP benefit configuration
- VIP event scheduling

### 8.7 Referral Management
- Reward configuration
- Conversion funnel analytics
- Abuse monitoring dashboard

---

## Phase 6 Completion Criteria

- [ ] Leaderboard Service running with daily/weekly/monthly/all-time/game-specific leaderboards
- [ ] Real-time leaderboard updates via WebSocket
- [ ] Anti-gaming rules preventing manipulation
- [ ] Winners Showcase displaying recent wins, big wins, jackpot winners
- [ ] Privacy controls for player name anonymization
- [ ] Banner Service with all banner types and targeting rules
- [ ] A/B testing capability for banners
- [ ] Banner analytics (impressions, CTR, conversions)
- [ ] Commission Service with RevShare, CPA, Hybrid models
- [ ] Automated commission calculation and settlement
- [ ] Affiliate Portal with tracking links, dashboards, sub-affiliates
- [ ] Loyalty Points system with earn/redeem functionality
- [ ] VIP Program with 5 tiers and auto-upgrade/downgrade
- [ ] Referral Program with rewards for referrer and referee
- [ ] Admin panel: all platform benefits management modules
- [ ] Player-facing: Leaderboards, winners, banners visible in lobby and mobile apps
