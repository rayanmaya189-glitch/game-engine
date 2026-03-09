# Phase 7: Sports Betting - Detailed Implementation Plan

## Overview

Phase 7 adds comprehensive sports betting capabilities including pre-match and live/in-play betting, odds management, multiple bet types, cash out, and sports-specific risk management.

**Prerequisites**: Phase 6 complete (Platform Benefits).

---

## 1. Sports Data Feed Service (Golang Kratos)

### 1.1 Project Setup
- Kratos project, gRPC (port 9025), Redis, NATS
- Adapter pattern for multiple data providers

### 1.2 Provider Integration
- **Sportradar**: Primary provider for major sports
- **BetConstruct**: Secondary provider
- **LSports**: Backup for niche sports

### 1.3 Data Types
- Fixtures (scheduled matches)
- Live scores and statistics
- Odds (pre-match and in-play)
- Results (settlement)

### 1.4 Real-Time Processing
- WebSocket feed for live data
- Sub-second latency target
- Data normalization layer (unified model)
- Provider failover logic

---

## 2. Sports Betting Service (Java Spring Boot)

### 2.1 Project Setup
- Spring Boot, gRPC (port 9026), PostgreSQL (casino_sports), Redis, NATS

### 2.2 Supported Sports
- Football/Soccer
- Basketball (NBA, EuroLeague, etc.)
- Tennis
- Cricket
- Baseball
- Hockey
- MMA/Boxing
- Esports (CS2, Dota 2, LoL, Valorant)

### 2.3 Market Types
| Type | Description |
|------|-------------|
| Match Winner | 1X2 (Home/Draw/Away) |
| Over/Under | Total goals/points over/under line |
| Handicap | Asian handicap, European handicap |
| Correct Score | Exact final score |
| Half Time/Full Time | HT/FT combination |
| Both Teams to Score | Yes/No |
| First Goal Scorer | Any/Specific player |
| Outright | League/tournament winner |
| Props | Player-specific bets |

### 2.4 Bet Types
- **Single**: One selection
- **Accumulator/Parlay**: Multiple selections, all must win
- **System**: Trixie, Yankee, Lucky 15, etc.
- **Bet Builder**: Combine selections from same event
- **Cash Out**: Full or partial early settlement
- **Edit Bet**: Add/remove selections

### 2.5 Odds Management
- Decimal, Fractional, American, Hong Kong, Malay, Indonesian formats
- Margin/overround configuration per sport/market
- Auto-suspend on significant odds movement
- Manual trader adjustment

### 2.6 Live Betting
- Real-time odds updates via WebSocket
- Bet acceptance with odds change tolerance
- In-play statistics display
- Match timeline

### 2.7 Risk Management
- Liability monitoring per event/market
- Max payout limits
- Trader alerts for unusual betting patterns
- Auto-market suspension triggers

### 2.8 Settlement
- Auto-settlement from results feed
- Manual override capability
- Void rules (abandoned events)
- Dead heat rules
- Rule 4 deductions (horse racing)

---

## 3. Admin Panel Extensions

### 3.1 Sports Betting Management
- Sport/competition configuration
- Market enable/disable
- Odds margin configuration
- Trader dashboard
- Liability monitoring
- Settlement interface

---

## Phase 7 Completion Criteria

- [ ] Sports Data Feed Service connected to provider(s)
- [ ] Sports Betting Service processing pre-match markets
- [ ] Live/in-play betting functional with WebSocket updates
- [ ] All bet types (single, accumulator, system, cash out) working
- [ ] Odds management with trader adjustments
- [ ] Risk management with liability monitoring
- [ ] Auto-settlement from feed
- [ ] Admin panel: sports configuration and settlement
