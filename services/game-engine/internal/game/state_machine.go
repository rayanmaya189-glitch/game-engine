package game

import (
	"errors"
	"time"
)

// Game phases
const (
	PhaseInit     = "init"
	PhaseBetting  = "betting"
	PhasePlaying  = "playing"
	PhaseSettling = "settling"
	PhaseComplete = "complete"
)

// Rake types
const (
	RakeTypeNone       = "none"
	RakeTypeFixed      = "fixed"
	RakeTypePercentage = "percentage"
	RakeTypeHybrid     = "hybrid"
)

// GameState represents the state of a game
type GameState struct {
	GameID      string                 `json:"game_id"`
	GameType    string                 `json:"game_type"`
	Phase       string                 `json:"phase"`
	PlayerID    string                 `json:"player_id"`
	Data        map[string]interface{} `json:"data"`
	TotalBet    int64                  `json:"total_bet"`
	TotalWin    int64                  `json:"total_win"`
	RakeAmount  int64                  `json:"rake_amount"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	CompletedAt *time.Time             `json:"completed_at,omitempty"`
}

// StateMachine manages game state transitions
type StateMachine struct {
	transitions map[string][]string
}

// NewStateMachine creates a new state machine
func NewStateMachine() *StateMachine {
	return &StateMachine{
		transitions: map[string][]string{
			PhaseInit:     {PhaseBetting},
			PhaseBetting:  {PhasePlaying, PhaseInit},
			PhasePlaying:  {PhaseSettling, PhaseBetting},
			PhaseSettling: {PhaseComplete},
			PhaseComplete: {PhaseInit},
		},
	}
}

// CanTransition checks if a state transition is valid
func (sm *StateMachine) CanTransition(from, to string) bool {
	allowed, ok := sm.transitions[from]
	if !ok {
		return false
	}
	for _, state := range allowed {
		if state == to {
			return true
		}
	}
	return false
}

// Transition performs a state transition
func (sm *StateMachine) Transition(state *GameState, to string) error {
	if !sm.CanTransition(state.Phase, to) {
		return errors.New("invalid state transition")
	}

	state.Phase = to
	state.UpdatedAt = time.Now()

	if to == PhaseComplete {
		now := time.Now()
		state.CompletedAt = &now
	}

	return nil
}

// GameConfig represents game configuration
type GameConfig struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Subtype     string  `json:"subtype"`
	RTP         float64 `json:"rtp"`
	HouseEdge   float64 `json:"house_edge"`
	MinBet      int64   `json:"min_bet"`
	MaxBet      int64   `json:"max_bet"`
	MaxPayout   int64   `json:"max_payout"`
	RakeType    string  `json:"rake_type"`
	RakeFixed   int64   `json:"rake_fixed"`
	RakePercent float64 `json:"rake_percent"`
	RakeMinCap  int64   `json:"rake_min_cap"`
	RakeMaxCap  int64   `json:"rake_max_cap"`
	Status      string  `json:"status"`
}

// RakeConfig represents rake configuration
type RakeConfig struct {
	Type    string  `json:"type"`
	Fixed   int64   `json:"fixed"`
	Percent float64 `json:"percent"`
	MinCap  int64   `json:"min_cap"`
	MaxCap  int64   `json:"max_cap"`
}

// CalculateRake calculates the rake for a game round
func CalculateRake(rakeConfig *RakeConfig, netWin int64) int64 {
	if rakeConfig == nil || rakeConfig.Type == RakeTypeNone {
		return 0
	}

	var rake int64

	switch rakeConfig.Type {
	case RakeTypeFixed:
		rake = rakeConfig.Fixed

	case RakeTypePercentage:
		rake = int64(float64(netWin) * rakeConfig.Percent)

	case RakeTypeHybrid:
		// Calculate percentage-based rake with caps
		percentRake := int64(float64(netWin) * rakeConfig.Percent)
		if percentRake < rakeConfig.MinCap {
			rake = rakeConfig.MinCap
		} else if percentRake > rakeConfig.MaxCap {
			rake = rakeConfig.MaxCap
		} else {
			rake = percentRake
		}
	}

	return rake
}

// NewGame creates a new game instance
func NewGame(gameType, playerID string, config *GameConfig) *GameState {
	return &GameState{
		GameID:     generateGameID(),
		GameType:   gameType,
		Phase:      PhaseBetting,
		PlayerID:   playerID,
		Data:       make(map[string]interface{}),
		TotalBet:   0,
		TotalWin:   0,
		RakeAmount: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

// PlaceBet places a bet in the game
func (g *GameState) PlaceBet(amount int64) error {
	if g.Phase != PhaseBetting {
		return errors.New("betting phase is over")
	}
	g.TotalBet += amount
	g.UpdatedAt = time.Now()
	return nil
}

// AddWin adds winnings to the game
func (g *GameState) AddWin(amount int64) {
	g.TotalWin += amount
	g.UpdatedAt = time.Now()
}

// Settle settles the game and calculates rake
func (g *GameState) Settle(rakeConfig *RakeConfig) error {
	if g.Phase != PhasePlaying {
		return errors.New("game is not in playing phase")
	}

	netWin := g.TotalWin - g.TotalBet
	if netWin > 0 {
		g.RakeAmount = CalculateRake(rakeConfig, netWin)
		g.TotalWin -= g.RakeAmount
	}

	g.UpdatedAt = time.Now()
	return nil
}

// Complete completes the game
func (g *GameState) Complete() error {
	if g.Phase != PhaseSettling {
		return errors.New("game is not in settling phase")
	}

	now := time.Now()
	g.CompletedAt = &now
	g.Phase = PhaseComplete
	g.UpdatedAt = now
	return nil
}

// generateGameID generates a unique game ID
func generateGameID() string {
	return time.Now().Format("20060102150405") + "-" + randomID(8)
}

func randomID(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[time.Now().UnixNano()%int64(len(letters))]
		time.Sleep(time.Nanosecond)
	}
	return string(b)
}
