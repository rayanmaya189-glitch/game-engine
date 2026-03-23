# Phase 2 Completion Plan - Core Game Engine

## Overview
Complete Phase 2 implementation to reach 100% completion. This plan addresses the gaps in Dice Games, Slot Games, Betting Service, and adds the Game Engine Service.

---

## 1. Craps Game Implementation

### Current State
- Basic Game struct with ID, Point, Dice, Phase (42 lines)
- Only supports come_out phase

### Required Implementation
Complete Craps game with all bet types and game phases.

### Bet Types to Implement
| Bet Type | Description | Payout |
|----------|-------------|--------|
| Pass Line | Win on 7/11 on come-out, lose on 2/3/12 | 1:1 |
| Don't Pass | Opposite of Pass Line | 1:1 |
| Come | Bet on next 7/11 after point established | 1:1 |
| Don't Come | Opposite of Come | 1:1 |
| Place Win | Bet on specific number (4,5,6,8,9,10) before 7 | Varies |
| Place Lose | Bet against specific number before 7 | Varies |
| Field | Bet on 2,3,4,9,10,11,12 | 2:1 for 2,12; 1:1 others |
| Big 6/Big 8 | Bet on 6 or 8 before 7 | 1:1 |
| Any 7 | Bet on any combination totaling 7 | 4:1 |
| Any Craps | Bet on 2,3,12 | 7:1 |
| Horn | Bet on 2,3,11,12 simultaneously | Varies |
| Hardway | Bet on double (e.g., 4+4) before 7 | 7:1 or 9:1 |

### Game Phases
1. **Come Out** - Initial roll, establish point
2. **Point** - Roll until point (win) or 7 (lose)

### Files to Modify
- `services/dice-games/internal/games/craps/game.go`

---

## 2. Sic Bo Game Implementation

### Current State
- Basic Game struct with ID, Dice, Bets (54 lines)
- Only supports basic roll and total calculation

### Required Implementation
Complete Sic Bo with all bet types and payouts.

### Bet Types to Implement
| Bet Type | Description | Payout |
|----------|-------------|--------|
| Small | Total 4-10 (excluding triples) | 1:1 |
| Big | Total 11-17 (excluding triples) | 1:1 |
| Specific Triple | All three dice same specific number | 150:1 |
| Any Triple | All three dice same (any number) | 24:1 |
| Specific Double | Two dice show same specific number | 8:1 |
| Four Number | Bet on 4 specific numbers | 7:1 |
| Three Number | Three dice total specific combination | 50:1 |
| Two Number | Two dice specific combination | 5:1 |
| Single | One specific number appears | 1:1, 2:1, 3:1 |

### Files to Modify
- `services/dice-games/internal/games/sicbo/game.go`

---

## 3. Slot Games Enhancement

### Current State
- Classic 3-reel and Video 5-reel games (406 lines)
- Basic paylines, scatter, bonus triggers
- Provably fair support

### Required Implementation
Add advanced slot game types.

### Additional Game Types
1. **Megaways Slots**
   - Dynamic reel sizes (up to 7 symbols per reel)
   - Up to 117,649 ways to win
   - Cascading/Avalanche feature
   - Multiplier increases with consecutive wins

2. **Cluster Pays**
   - Grid-based matching (not paylines)
   - Symbols cluster together for wins
   - Cascading removal of winning symbols
   - New symbols fall from above

3. **Progressive Jackpot**
   - Seed amount (minimum jackpot)
   - Contribution model (percentage of bet)
   - Trigger mechanisms (symbol combination or random)
   - Multi-tier (Mini, Minor, Major, Grand)

### Files to Modify
- `services/slot-games/internal/games/slot.go`
- Add new files for Megaways and Cluster games

---

## 4. Betting Service Implementation

### Current State
- Empty BettingService struct (9 lines)
- Basic limit manager and odds calculator exist

### Required Implementation
Complete betting service with all bet types and lifecycle management.

### Core Features
1. **Bet Types**
   - Single Bet: One outcome
   - Accumulator: Multiple outcomes, all must win
   - System Bet: Multiple combinations (Patent, Yankee, etc.)

2. **Bet Lifecycle**
   - Placed → Accepted → Active → Settled → Paid
   - Void/Cancel support with reason tracking

3. **Odds Formats**
   - Decimal (e.g., 2.50)
   - Fractional (e.g., 3/2)
   - American (e.g., +150)
   - Hong Kong

4. **Bet Validation**
   - Balance check
   - Min/Max limits per game, player, table
   - Game state validation

5. **Settlement Engine**
   - Automatic settlement based on game results
   - Manual void capability
   - Full audit trail

### Files to Modify
- `services/betting/internal/service/betting_service.go`
- `services/betting/internal/handler/betting_handler.go`

---

## 5. Game Engine Service (New)

### Required Implementation
Create a centralized Game Engine Service for RNG, game state management, and provably fair verification.

### Core Features
1. **RNG (Random Number Generation)**
   - Certified PRNG (Fortuna algorithm)
   - Cryptographically secure
   - Provably fair verification

2. **Game State Machine**
   - Init → Betting → Playing → Settling → Complete
   - Full audit trail

3. **Game Registry**
   - Dynamic game configuration
   - RTP configuration per game
   - House edge management

4. **Rake/Commission Engine**
   - Fixed amount
   - Percentage-based
   - Hybrid (percentage with min/max cap)

### New Files to Create
- `services/game-engine/` directory with full Go service structure

---

## Implementation Order

1. **Week 1**: Complete Craps and Sic Bo games
2. **Week 2**: Enhance Slot games with Megaways/Cluster
3. **Week 3**: Complete Betting Service
4. **Week 4**: Create Game Engine Service

---

## Success Criteria

- [ ] Craps: All bet types functional with correct payouts
- [ ] Sic Bo: All bet types functional with correct payouts
- [ ] Slots: Megaways, Cluster Pays, Progressive Jackpot working
- [ ] Betting: Full lifecycle management, all bet types
- [ ] Game Engine: RNG certified, provably fair, state machine working
- [ ] All services compile and pass basic tests
