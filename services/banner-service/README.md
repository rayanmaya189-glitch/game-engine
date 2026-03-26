# Banner Service - Phase 6 Platform Benefits

## Overview
The Banner & Announcement Service manages targeted banners and announcements for player engagement.

## Implementation Status: Started

### Created Files
- `cmd/main.go` - Main entry point with routes
- `internal/config/config.go` - Configuration
- `internal/model/banner.go` - Data models
- `internal/repository/banner_repository.go` - Database operations
- `internal/service/banner_service.go` - Business logic
- `internal/handler/banner_handler.go` - HTTP handlers
- `migrations/001_init_schema.sql` - Database schema
- `go.mod` - Dependencies
- `config.yaml` - Configuration file
- `Dockerfile` - Container definition

### API Endpoints
| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/v1/banners` | GET | List active banners |
| `/api/v1/banners/:id` | GET | Get banner details |
| `/api/v1/banners` | POST | Create banner (admin) |
| `/api/v1/banners/:id` | PUT | Update banner (admin) |
| `/api/v1/banners/:id` | DELETE | Delete banner (admin) |
| `/api/v1/announcements` | GET | List active announcements |
| `/api/v1/announcements` | POST | Create announcement (admin) |

### Banner Types
- Hero Banner (1920x600) - Homepage top
- Sidebar Banner (300x600) - Desktop sidebar
- Popup Banner (800x600) - Modal on entry
- In-Game Banner (728x90) - Between rounds
- Interstitial - Full screen transition
- Notification Bar - Top of screen

### Targeting Options
- By country
- By VIP level
- By game type
- By deposit status
- By last activity
- Random rotation

## Next Steps
1. Build and test the Banner Service
2. Create Referral Service
3. Enhance Loyalty Service