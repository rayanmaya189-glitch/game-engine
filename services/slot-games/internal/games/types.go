package games

// MegawaysConfig represents Megaways slot game configuration
type MegawaysConfig struct {
	Reels                int     `json:"reels"`
	MinRows              int     `json:"min_rows"` // Minimum symbols per reel
	MaxRows              int     `json:"max_rows"` // Maximum symbols per reel
	MinBet               int64   `json:"min_bet"`
	MaxBet               int64   `json:"max_bet"`
	MinLineBet           int64   `json:"min_line_bet"`
	MaxLineBet           int64   `json:"max_line_bet"`
	RTP                  float64 `json:"rtp"`                    // Return to player
	Volatility           string  `json:"volatility"`             // low, medium, high
	MaxMegaways          int     `json:"max_megaways"`           // Maximum ways to win (default 117649)
	CascadeEnabled       bool    `json:"cascade_enabled"`        // Avalanche/cascade feature
	CascadeMultiplier    float64 `json:"cascade_multiplier"`     // Multiplier increase per cascade
	MaxCascadeMultiplier float64 `json:"max_cascade_multiplier"` // Maximum cascade multiplier
}

// MegawaysGame represents a Megaways slot game
type MegawaysGame struct {
	ID                string
	Config            *MegawaysConfig
	Symbols           []Symbol
	ReelSymbols       [][]int // [reel][row] - variable height
	RowCounts         []int   // Number of rows for each reel
	Bet               int64
	LineBet           int64
	Ways              int // Current number of megaways
	Win               int64
	CascadeLevel      int
	CascadeMultiplier float64
	WinLines          []WinLine
	TotalWin          int64
	IsComplete        bool
	ProvablyFair      bool
	ServerSeed        string
	ClientSeed        string
	Nonce             int
	CascadeWins       []int64 // Wins from each cascade level
}

// MegawaysGameState represents the game state
type MegawaysGameState struct {
	GameID            string     `json:"game_id"`
	Reels             int        `json:"reels"`
	RowCounts         []int      `json:"row_counts"`
	Symbols           [][]string `json:"symbols"`
	Bet               int64      `json:"bet"`
	LineBet           int64      `json:"line_bet"`
	Ways              int        `json:"ways"`
	Win               int64      `json:"win"`
	TotalWin          int64      `json:"total_win"`
	WinLines          []WinLine  `json:"win_lines"`
	CascadeLevel      int        `json:"cascade_level"`
	CascadeMultiplier float64    `json:"cascade_multiplier"`
	CascadeWins       []int64    `json:"cascade_wins"`
	IsComplete        bool       `json:"is_complete"`
	ProvablyFair      bool       `json:"provably_fair"`
}

// ClusterGame represents a Cluster Pays slot game
type ClusterGame struct {
	ID           string
	Config       *ClusterConfig
	Symbols      []Symbol
	Grid         [][]int // [row][col]
	Bet          int64
	Win          int64
	WinClusters  []Cluster
	TotalWin     int64
	IsComplete   bool
	ProvablyFair bool
	ServerSeed   string
	ClientSeed   string
	Nonce        int
}

// ClusterConfig represents Cluster game configuration
type ClusterConfig struct {
	Rows       int     `json:"rows"`
	Cols       int     `json:"cols"`
	MinBet     int64   `json:"min_bet"`
	MaxBet     int64   `json:"max_bet"`
	RTP        float64 `json:"rtp"`
	Volatility string  `json:"volatility"`
	MinCluster int     `json:"min_cluster"` // Minimum symbols for a cluster win
	Cascade    bool    `json:"cascade"`
}

// Cluster represents a winning cluster
type Cluster struct {
	Symbol    string     `json:"symbol"`
	Count     int        `json:"count"`
	Payout    int64      `json:"payout"`
	Positions []Position `json:"positions"`
}

// Position represents a grid position
type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// ClusterGameState represents the game state
type ClusterGameState struct {
	GameID     string     `json:"game_id"`
	Rows       int        `json:"rows"`
	Cols       int        `json:"cols"`
	Symbols    [][]string `json:"symbols"`
	Bet        int64      `json:"bet"`
	Win        int64      `json:"win"`
	TotalWin   int64      `json:"total_win"`
	Clusters   []Cluster  `json:"clusters"`
	IsComplete bool       `json:"is_complete"`
}

// ProgressiveGame represents a Progressive Jackpot slot game
type ProgressiveGame struct {
	ID               string
	GameType         string
	Config           *ProgressiveConfig
	Symbols          []Symbol
	Positions        [][]int
	Bet              int64
	LineBet          int64
	Win              int64
	TotalWin         int64
	JackpotTriggered bool
	JackpotTier      string
	JackpotWin       int64
	IsComplete       bool
}

// ProgressiveConfig represents Progressive game configuration
type ProgressiveConfig struct {
	Reels      int                           `json:"reels"`
	Rows       int                           `json:"rows"`
	MinBet     int64                         `json:"min_bet"`
	MaxBet     int64                         `json:"max_bet"`
	RTP        float64                       `json:"rtp"`
	Volatility string                        `json:"volatility"`
	Jackpots   map[string]ProgressiveJackpot `json:"jackpots"`
}

// ProgressiveJackpot represents a progressive jackpot tier
type ProgressiveJackpot struct {
	Name             string   `json:"name"`
	SeedAmount       int64    `json:"seed_amount"`
	ContributionRate float64  `json:"contribution_rate"` // Percentage of bet
	MinBet           int64    `json:"min_bet"`
	Trigger          string   `json:"trigger"` // "symbol", "random", "feature"
	Symbols          []string `json:"symbols,omitempty"`
	Odds             float64  `json:"odds"` // Probability of winning
}

// ProgressiveGameState represents the game state
type ProgressiveGameState struct {
	GameID           string     `json:"game_id"`
	GameType         string     `json:"game_type"`
	Reels            int        `json:"reels"`
	Rows             int        `json:"rows"`
	Symbols          [][]string `json:"symbols"`
	Bet              int64      `json:"bet"`
	Win              int64      `json:"win"`
	TotalWin         int64      `json:"total_win"`
	JackpotTriggered bool       `json:"jackpot_triggered"`
	JackpotTier      string     `json:"jackpot_tier,omitempty"`
	JackpotWin       int64      `json:"jackpot_win,omitempty"`
	IsComplete       bool       `json:"is_complete"`
}
