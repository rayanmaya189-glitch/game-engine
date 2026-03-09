package games

// Symbol represents a slot symbol
type Symbol struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Value     int    `json:"value"`
	IsWild    bool   `json:"is_wild"`
	IsScatter bool   `json:"is_scatter"`
}

// Payline represents a payline
type Payline struct {
	ID        int     `json:"id"`
	Positions [][]int `json:"positions"`
	Payout    int64   `json:"payout"`
}

// Game represents a slot game
type Game struct {
	ID        string
	Reels     int
	Symbols   []Symbol
	Paylines  []Payline
	Positions []int
	Bet       int64
	Win       int64
}

// NewGame creates a new slot game
func NewGame(id string, reels int) *Game {
	return &Game{
		ID:        id,
		Reels:     reels,
		Symbols:   make([]Symbol, 0),
		Paylines:  make([]Payline, 0),
		Positions: make([]int, reels),
	}
}

// Spin performs a spin
func (g *Game) Spin(positions []int) {
	g.Positions = positions
}

// Evaluate evaluates the spin result
func (g *Game) Evaluate() int64 {
	// Check paylines for wins
	g.Win = 0
	return g.Win
}

// GetState returns the game state
func (g *Game) GetState() *GameState {
	return &GameState{
		GameID:    g.ID,
		Reels:     g.Reels,
		Positions: g.Positions,
		Bet:       g.Bet,
		Win:       g.Win,
	}
}

// GameState represents the game state
type GameState struct {
	GameID    string `json:"game_id"`
	Reels     int    `json:"reels"`
	Positions []int  `json:"positions"`
	Bet       int64  `json:"bet"`
	Win       int64  `json:"win"`
}
