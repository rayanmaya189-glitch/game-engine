package tournament

import (
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// TournamentType represents the type of tournament
type TournamentType string

const (
	TournamentTypeSitAndGo  TournamentType = "sit_and_go"
	TournamentTypeScheduled TournamentType = "scheduled"
	TournamentTypeKnockout  TournamentType = "knockout"
	TournamentTypeFreeroll  TournamentType = "freeroll"
)

// TournamentStatus represents the status of a tournament
type TournamentStatus string

const (
	TournamentStatusPending     TournamentStatus = "pending"
	TournamentStatusRegistering TournamentStatus = "registering"
	TournamentStatusRunning     TournamentStatus = "running"
	TournamentStatusCompleted   TournamentStatus = "completed"
	TournamentStatusCancelled   TournamentStatus = "cancelled"
)

// Tournament represents a tournament in the system
type Tournament struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	Type            TournamentType         `json:"type"`
	Status          TournamentStatus       `json:"status"`
	GameType        string                 `json:"game_type"`
	MinPlayers      int                    `json:"min_players"`
	MaxPlayers      int                    `json:"max_players"`
	CurrentPlayers  int                    `json:"current_players"`
	EntryFee        int64                  `json:"entry_fee"`
	BuyIn           int64                  `json:"buy_in"`
	PrizePool       int64                  `json:"prize_pool"`
	StartTime       time.Time              `json:"start_time"`
	EndTime         time.Time              `json:"end_time"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	Settings        TournamentSettings     `json:"settings"`
	RegisteredUsers map[string]Participant `json:"registered_users"`
	Results         []Result               `json:"results"`
	Bracket         *Bracket               `json:"bracket,omitempty"`
}

// TournamentSettings contains tournament-specific settings
type TournamentSettings struct {
	AutoStart         bool         `json:"auto_start"`
	RebuyEnabled      bool         `json:"rebuy_enabled"`
	AddonEnabled      bool         `json:"addon_enabled"`
	Reentries         int          `json:"reentries"`
	LateRegistration  int          `json:"late_registration"` // seconds
	StartingChips     int          `json:"starting_chips"`
	BlindLevels       []BlindLevel `json:"blind_levels"`
	PrizeDistribution []int        `json:"prize_distribution"`
}

// BlindLevel represents a blind level in a tournament
type BlindLevel struct {
	Level      int `json:"level"`
	SmallBlind int `json:"small_blind"`
	BigBlind   int `json:"big_blind"`
	Duration   int `json:"duration"` // seconds
}

// Participant represents a player in a tournament
type Participant struct {
	UserID       string    `json:"user_id"`
	Username     string    `json:"username"`
	Chips        int       `json:"chips"`
	Position     int       `json:"position"`
	Status       string    `json:"status"` // active, eliminated, registered
	Score        int       `json:"score"`
	Rank         int       `json:"rank"`
	Knockouts    int       `json:"knockouts"`
	PrizeWon     int64     `json:"prize_won"`
	RegisteredAt time.Time `json:"registered_at"`
}

// Result represents a tournament result
type Result struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Rank      int    `json:"rank"`
	Prize     int64  `json:"prize"`
	Knockouts int    `json:"knockouts"`
	Score     int    `json:"score"`
}

// Bracket represents a knockout tournament bracket
type Bracket struct {
	Rounds []Round `json:"rounds"`
}

// Round represents a single round in a bracket
type Round struct {
	Number   int     `json:"number"`
	Matches  []Match `json:"matches"`
	Complete bool    `json:"complete"`
}

// Match represents a single match in the bracket
type Match struct {
	ID       string  `json:"id"`
	Player1  *string `json:"player1"`
	Player2  *string `json:"player2"`
	Winner   *string `json:"winner"`
	Score1   int     `json:"score1"`
	Score2   int     `json:"score2"`
	Complete bool    `json:"complete"`
}

// Manager handles tournament lifecycle
type Manager struct {
	mu          sync.RWMutex
	tournaments map[string]*Tournament
	RedisClient *redis.Client
	Config      *Config
	Leaderboard *Leaderboard
	PrizePool   *PrizePool
	Scheduler   *Scheduler
}
