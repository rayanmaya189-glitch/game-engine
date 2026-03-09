# Phase 5: Native Mobile Apps - Detailed Implementation Plan

## Overview

Phase 5 delivers native Android (Kotlin) and iOS (Swift) apps for players, covering authentication, game lobby, all game types (card/dice/slot), wallet, tournaments, profile, and push notifications with full mobile security.

**Prerequisites**: Phase 4 complete (Payments, KYC, AML, Fraud, Bonuses).

---

## 1. Android App (Kotlin + Jetpack Compose)

### 1.1 Project Setup
- Android Studio project with Kotlin DSL (build.gradle.kts)
- Min SDK: 26 (Android 8.0), Target SDK: 35
- Architecture: MVVM + Clean Architecture + Multi-module
- Dependencies: Hilt (DI), Retrofit + OkHttp (networking), Room (local DB), Jetpack Compose (UI), Compose Navigation, Coil (images), DataStore (preferences)

### 1.2 Module Structure
```
app/                    # Main app module
├── core/
│   ├── network/        # Retrofit, OkHttp, WebSocket, interceptors
│   ├── database/       # Room DB, DAOs
│   ├── security/       # Certificate pinning, root detection, encryption
│   ├── analytics/      # Firebase Analytics
│   └── common/         # Shared utilities, extensions
├── feature/
│   ├── auth/           # Login, Register, 2FA, Forgot Password
│   ├── lobby/          # Game catalog, search, filters
│   ├── card-games/     # Blackjack, Baccarat, Poker UIs
│   ├── dice-games/     # Hi-Lo, Sic Bo, Craps UIs
│   ├── slot-games/     # Slot machine animations, spin UI
│   ├── wallet/         # Balance, Deposit, Withdraw, History
│   ├── tournament/     # Browse, Register, Play, Leaderboard
│   ├── profile/        # Settings, KYC, Limits, Self-exclusion
│   ├── notifications/  # Notification center
│   └── support/        # Chat, FAQ, Contact
└── design-system/      # Shared UI components, theme, typography
```

### 1.3 Core Networking Layer
- **Retrofit** for REST API calls with JWT interceptor
- **OkHttp WebSocket** for real-time game communication
- **Certificate pinning**: Pin API server certificates
- **Token management**: Auto-refresh expired JWT, retry failed requests
- **Offline handling**: Queue actions when offline, sync when online
- **Request/Response logging**: Debug builds only

### 1.4 Feature: Authentication
- Login screen (email/phone + password)
- Registration screen (email, phone, password, country)
- 2FA screen (TOTP code entry)
- Forgot password flow
- Biometric login (fingerprint/face) using AndroidX Biometric
- Social login (Google Sign-In, Apple Sign-In)
- Session management: Auto-logout on token expiry

### 1.5 Feature: Game Lobby
- Game catalog with categories (Cards, Dice, Slots, Live, Sports)
- Search and filter (by name, type, popularity, RTP)
- Game cards with thumbnail, name, RTP, min bet
- Recently played games
- Favorite games
- Featured/promoted games banner carousel

### 1.6 Feature: Card Games UI
- **Blackjack**: Animated card dealing, chip placement, action buttons (Hit/Stand/Double/Split), dealer reveal animation
- **Baccarat**: Card dealing animation, roadmap display, bet areas (Player/Banker/Tie)
- **Poker**: Table view with player avatars, hole cards, community cards, pot display, action buttons, timer
- **Common**: Card flip animations, chip stack visualization, win/loss animations

### 1.7 Feature: Dice Games UI
- **Hi-Lo**: Slider for target number, roll animation, result display
- **Sic Bo**: Dice roll animation, bet board with all positions, payout display
- **Craps**: Table layout, dice throw animation, bet placement areas

### 1.8 Feature: Slot Games UI
- **Reel animation**: Smooth spinning reels with easing
- **Win celebration**: Particle effects, coin shower, big win screen
- **Payline visualization**: Highlight winning paylines
- **Free spins**: Special UI mode with counter and multiplier display
- **Gamble feature**: Card flip or coin toss animation
- **Auto-spin**: Configuration dialog, progress indicator

### 1.9 Feature: Wallet
- Balance display (real, bonus, free bet)
- Deposit flow: Select method → Enter amount → Gateway redirect/in-app
- Withdrawal flow: Select method → Enter amount → Confirmation
- Transaction history with filters
- Pending withdrawals with cancel option

### 1.10 Feature: Tournaments
- Tournament lobby: Browse upcoming, running, completed
- Tournament detail: Info, prize structure, blind structure, registered players
- Registration/unregistration
- In-tournament: Table view, leaderboard, blind timer
- Tournament results and prize history

### 1.11 Feature: Profile & Settings
- Profile editing (name, avatar, preferences)
- KYC document upload (camera capture, gallery select)
- Responsible gaming: Set deposit/bet/loss limits, self-exclusion
- Notification preferences
- Language selection
- App settings (sound, vibration, auto-play defaults)

### 1.12 Push Notifications
- Firebase Cloud Messaging (FCM) integration
- Token registration on login, refresh on token change
- Notification channels: Game alerts, Financial, Promotions, System
- Deep linking: Tap notification → navigate to relevant screen
- In-app notification center with read/unread status

---

## 2. iOS App (Swift + SwiftUI)

### 2.1 Project Setup
- Xcode project with Swift Package Manager
- Min iOS: 16.0, Target: 18.0
- Architecture: MVVM + Clean Architecture
- Dependencies: Alamofire (networking), Starscream (WebSocket), SwiftUI, Combine, KeychainAccess, Kingfisher (images)

### 2.2 Module Structure
```
CasinoApp/
├── Core/
│   ├── Network/        # API client, WebSocket, interceptors
│   ├── Storage/        # Core Data, UserDefaults, Keychain
│   ├── Security/       # Certificate pinning, jailbreak detection
│   ├── Analytics/      # Firebase Analytics
│   └── Common/         # Extensions, utilities
├── Features/
│   ├── Auth/           # Login, Register, 2FA
│   ├── Lobby/          # Game catalog
│   ├── CardGames/      # Blackjack, Baccarat, Poker
│   ├── DiceGames/      # Hi-Lo, Sic Bo, Craps
│   ├── SlotGames/      # Slot animations
│   ├── Wallet/         # Balance, Deposit, Withdraw
│   ├── Tournament/     # Browse, Play, Leaderboard
│   ├── Profile/        # Settings, KYC
│   ├── Notifications/  # Notification center
│   └── Support/        # Chat, FAQ
└── DesignSystem/       # Shared components, theme
```

### 2.3 iOS-Specific Implementations
- **Face ID / Touch ID**: LocalAuthentication framework for biometric login
- **Apple Sign-In**: AuthenticationServices framework (required for App Store)
- **Keychain**: Secure storage for tokens and sensitive data
- **Core Data**: Local caching for offline access
- **APNs**: Push notifications via Firebase (FCM wraps APNs)
- **App Clips**: Quick game preview without full app install (optional)
- **Widgets**: WidgetKit for balance display and upcoming tournaments

### 2.4 Feature Parity with Android
- All features from Android section (1.4 - 1.12) implemented with SwiftUI equivalents
- Platform-specific UI patterns (tab bar vs bottom nav, sheets vs dialogs)
- Haptic feedback using UIFeedbackGenerator

---

## 3. Shared Mobile Concerns

### 3.1 Mobile Security Implementation
| Security Measure | Android | iOS |
|-----------------|---------|-----|
| Certificate Pinning | OkHttp CertificatePinner | URLSession delegate |
| Root/Jailbreak Detection | SafetyNet/Play Integrity | IOSSecuritySuite |
| Code Obfuscation | R8/ProGuard | Swift compilation + bitcode |
| Secure Storage | EncryptedSharedPreferences + Keystore | Keychain |
| App Integrity | Play Integrity API | App Attest |
| Anti-Tampering | Runtime checks | Runtime checks |
| Screenshot Prevention | FLAG_SECURE on sensitive screens | UIScreen.isCaptured |
| Clipboard Protection | Clear clipboard on app background | Clear UIPasteboard |
| SSL/TLS | TLS 1.3 minimum | ATS (App Transport Security) |
| Binary Protection | NDK for critical logic | - |

### 3.2 WebSocket Management
- Persistent WebSocket connection while in game
- Auto-reconnect with exponential backoff (1s, 2s, 4s, 8s, max 30s)
- Heartbeat every 30 seconds
- Message queuing during reconnection
- Connection state indicator in UI

### 3.3 Offline Handling
- Cache game catalog, user profile, transaction history
- Show cached data with "last updated" indicator
- Queue deposit/withdrawal requests (process when online)
- Graceful degradation: Show lobby, disable real-money play

### 3.4 Performance Optimization
- Lazy loading for game lists and images
- Image caching with disk and memory cache
- Prefetch next page of paginated lists
- Minimize WebSocket message size (binary protocol for game state)
- Background task management (limit battery usage)

### 3.5 App Store Compliance
- **Google Play**: Gambling app policy compliance, restricted countries, age verification
- **Apple App Store**: Guideline 5.3.4 (gambling apps), age rating 17+, geo-restrictions
- **Both**: Privacy policy, terms of service, responsible gaming information

---

## Phase 5 Completion Criteria

- [ ] Android app installable and functional on Android 8.0+
- [ ] iOS app installable and functional on iOS 16.0+
- [ ] Authentication flow (register, login, 2FA, biometric) working on both platforms
- [ ] Game lobby displaying all available games with search/filter
- [ ] Card games (Blackjack, Baccarat, Poker) playable with animations
- [ ] Dice games (Hi-Lo, Sic Bo, Craps) playable with animations
- [ ] Slot games playable with reel animations, free spins, gamble feature
- [ ] Wallet: deposit, withdraw, transaction history functional
- [ ] Tournament browsing, registration, and in-tournament play working
- [ ] Push notifications delivered and deep-linking to correct screens
- [ ] Certificate pinning and root/jailbreak detection active
- [ ] WebSocket real-time gameplay with reconnection handling
- [ ] KYC document upload via camera/gallery working
- [ ] App Store / Play Store submission requirements met
