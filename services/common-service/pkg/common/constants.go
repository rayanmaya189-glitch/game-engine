// Package common provides shared utilities, error codes, and constants
// used across all services in the casino game engine.
package common

import "time"

// ===========================================
// Service Constants
// ===========================================

// Service names
const (
	ServiceGateway      = "gateway"
	ServiceAuth         = "auth-service"
	ServiceUser         = "user-service"
	ServiceWallet       = "wallet-service"
	ServiceGameRegistry = "game-registry"
)

// Default ports for services
const (
	GatewayPort       = 8080
	GatewayGRPCPort   = 8081
	AuthServicePort   = 4433
	AuthServiceGRPC   = 4434
	UserServicePort   = 4435
	UserServiceGRPC   = 4436
	WalletServicePort = 4437
	WalletServiceGRPC = 4438
	GameRegistryPort  = 4439
	GameRegistryGRPC  = 4440
)

// ===========================================
// Database Constants
// ===========================================

const (
	// PostgreSQL
	PostgresDefaultPort     = 5432
	PostgresTimescalePort   = 5433
	PostgresMaxConnections  = 100
	PostgresMaxIdleConns    = 10
	PostgresConnMaxLifetime = 5 * time.Minute
	PostgresConnMaxIdleTime = 1 * time.Minute

	// Redis
	RedisDefaultPort  = 6379
	RedisDefaultDB    = 0
	RedisMaxRetries   = 3
	RedisPoolSize     = 100
	RedisMinIdleConns = 10

	// NATS
	NATSDefaultPort    = 4222
	NATSMgmtPort       = 8222
	NATSDefaultCluster = "game_engine"
	NATSDefaultStream  = "GAME_EVENTS"
)

// ===========================================
// JWT & Authentication Constants
// ===========================================

const (
	// JWT settings
	JWTDefaultAlgorithm  = "RS256"
	JWTDefaultExpiration = 24 * time.Hour
	JWTRefreshExpiration = 7 * 24 * time.Hour

	// Session settings
	SessionDefaultTimeout = 30 * time.Minute
	SessionCookieName     = "game_session"

	// Rate limiting
	RateLimitDefaultRequests = 100
	RateLimitDefaultWindow   = 1 * time.Minute
)

// ===========================================
// Game Engine Constants
// ===========================================

const (
	// Game settings
	MaxPlayersPerTable = 6
	MinBetAmount       = 1.0
	MaxBetAmount       = 100000.0

	// Transaction limits
	MaxDepositAmount    = 1000000.0
	MaxWithdrawalAmount = 100000.0
	MinWithdrawalAmount = 10.0

	// Wallet settings
	WalletDecimalPlaces = 2
)

// DefaultChipDenominations are the default chip denominations for games
var DefaultChipDenominations = []int64{1, 5, 10, 25, 100, 500, 1000, 5000, 10000}

// ===========================================
// API Constants
// ===========================================

const (
	// API versions
	APIVersion1 = "v1"
	APIVersion  = APIVersion1

	// API paths
	APIPathAuth   = "/api/v1/auth"
	APIPathUsers  = "/api/v1/users"
	APIPathWallet = "/api/v1/wallet"
	APIPathGames  = "/api/v1/games"
	APIPathHealth = "/health"
	APIPathReady  = "/ready"

	// Headers
	HeaderRequestID      = "X-Request-ID"
	HeaderAuthorization  = "Authorization"
	HeaderContentType    = "Content-Type"
	HeaderAcceptLanguage = "Accept-Language"

	// Content types
	ContentTypeJSON     = "application/json"
	ContentTypeProtoBuf = "application/protobuf"
	ContentTypeForm     = "application/x-www-form-urlencoded"
)

// ===========================================
// Time Constants
// ===========================================

const (
	// Timeouts
	DefaultHTTPTimeout  = 30 * time.Second
	DefaultGRPCTimeout  = 10 * time.Second
	DefaultDBTimeout    = 5 * time.Second
	DefaultCacheTimeout = 1 * time.Second

	// Intervals
	HealthCheckInterval = 30 * time.Second
	MetricsInterval     = 10 * time.Second

	// Durations
	DefaultRetryDelay = 100 * time.Millisecond
	MaxRetryDelay     = 5 * time.Second
)

// ===========================================
// Pagination Constants
// ===========================================

const (
	DefaultPage     = 1
	DefaultPageSize = 20
	MaxPageSize     = 100
)

// ===========================================
// Error Codes
// ===========================================

// Error codes for the game engine
const (
	// General errors (1xxx)
	ErrCodeInternalError      = "E1001"
	ErrCodeInvalidRequest     = "E1002"
	ErrCodeNotFound           = "E1003"
	ErrCodeUnauthorized       = "E1004"
	ErrCodeForbidden          = "E1005"
	ErrCodeConflict           = "E1006"
	ErrCodeRateLimitExceeded  = "E1007"
	ErrCodeServiceUnavailable = "E1008"
	ErrCodeTimeout            = "E1009"
	ErrCodeBadRequest         = "E1010"
	ErrCodeValidationError    = "E1011"

	// Authentication errors (2xxx)
	ErrCodeInvalidCredentials = "E2001"
	ErrCodeTokenExpired       = "E2002"
	ErrCodeTokenInvalid       = "E2003"
	ErrCodeAccountLocked      = "E2004"
	ErrCodeAccountDisabled    = "E2005"
	ErrCodeSessionExpired     = "E2006"
	ErrCodeMFARequired        = "E2007"
	ErrCodeMFAInvalid         = "E2008"

	// User errors (3xxx)
	ErrCodeUserNotFound      = "E3001"
	ErrCodeUserAlreadyExists = "E3002"
	ErrCodeInvalidEmail      = "E3003"
	ErrCodeInvalidUsername   = "E3004"
	ErrCodeWeakPassword      = "E3005"
	ErrCodeEmailNotVerified  = "E3006"

	// Wallet errors (4xxx)
	ErrCodeInsufficientFunds = "E4001"
	ErrCodeInvalidAmount     = "E4002"
	ErrCodeTransactionFailed = "E4003"
	ErrCodeWithdrawalPending = "E4004"
	ErrCodeDepositFailed     = "E4005"
	ErrCodeWalletLocked      = "E4006"

	// Game errors (5xxx)
	ErrCodeGameNotFound      = "E5001"
	ErrCodeGameClosed        = "E5002"
	ErrCodeTableFull         = "E5003"
	ErrCodeNotYourTurn       = "E5004"
	ErrCodeBetTooLow         = "E5005"
	ErrCodeBetTooHigh        = "E5006"
	ErrCodeInsufficientChips = "E5007"
	ErrCodeGameUnavailable   = "E5008"
)

// ===========================================
// Currency & Locale Constants
// ===========================================

const (
	// Default currency
	DefaultCurrency = "USD"

	// Supported currencies
	CurrencyUSD = "USD"
	CurrencyEUR = "EUR"
	CurrencyGBP = "GBP"
	CurrencyJPY = "JPY"

	// Supported locales
	LocaleEnUS = "en-US"
	LocaleJaJP = "ja-JP"
	LocaleZhCN = "zh-CN"

	// Default locale
	DefaultLocale = LocaleEnUS
)
