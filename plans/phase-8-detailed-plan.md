# Phase 8: Live Dealer - Detailed Implementation Plan

## Overview

Phase 8 adds live dealer casino games with real-time video streaming, table management, dealer interfaces, and integration with third-party live dealer providers.

**Prerequisites**: Phase 7 complete (Sports Betting).

---

## 1. Live Dealer Service (Golang Kratos)

### 1.1 Project Setup
- Kratos project, gRPC (port 9027), PostgreSQL (casino_live_dealer), Redis, WebSocket

### 1.2 Supported Games
- Live Blackjack (7-seat tables)
- Live Baccarat (squeeze, speed, no-commission variants)
- Live Roulette (European, American, Lightning)
- Live Casino Hold'em
- Live Three Card Poker
- Live Caribbean Stud
- Live Sic Bo
- Live Dragon Tiger
- Live Game Shows (Dream Catcher style)

### 1.3 Video Streaming Infrastructure
- **WebRTC**: Primary streaming protocol (< 1s latency)
- **Adaptive bitrate**: 360p, 720p, 1080p based on bandwidth
- **CDN**: CloudFront for global distribution
- **HLS/DASH**: Fallback for compatibility

### 1.4 Table Management
- Min/max bet configuration per table
- Game variant selection
- Dealer assignment
- VIP/Private tables (invitation-only, higher limits)
- Table scheduling (open/close times)

### 1.5 Table-Level Rake
- Configurable fixed or percentage rake per table
- Override game-level rake configuration

### 1.6 Dealer Interface
- Tablet application for game control
- Card scanning (OCR/RFID) for result entry
- Manual result entry with verification
- Performance tracking

### 1.7 Game State Synchronization
- Real-time broadcast to all connected players
- Betting window management
- Result announcement with animations

---

## 2. Third-Party Integration

### 2.1 Provider Support
- Evolution Gaming (Ezugi)
- Pragmatic Play Live
- Custom adapter pattern for additional providers

### 2.2 Integration Points
- Game feed (video + results)
- Bet placement API
- Player management
- Settlement reconciliation

---

## Phase 8 Completion Criteria

- [ ] Live Dealer Service with table management
- [ ] Video streaming via WebRTC with adaptive bitrate
- [ ] Dealer interface for game control
- [ ] Table-level rake configuration
- [ ] Integration with at least one third-party provider
- [ ] Live games playable on web and mobile
