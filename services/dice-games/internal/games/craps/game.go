package craps

// Game represents a Craps game
type Game struct {
	ID    string
	Point int
	Dice  []int
	Phase string
}

// NewGame creates a new Craps game
func NewGame(id string) *Game {
	return &Game{
		ID:    id,
		Point: 0,
		Dice:  make([]int, 2),
		Phase: "come_out",
	}
}

// Roll rolls the dice
func (g *Game) Roll(dice []int) {
	g.Dice = dice
}

// GetState returns the game state
func (g *Game) GetState() *GameState {
	return &GameState{
		GameID: g.ID,
		Dice:   g.Dice,
		Point:  g.Point,
		Phase:  g.Phase,
	}
}

// GameState represents the game state
type GameState struct {
	GameID string `json:"game_id"`
	Dice   []int  `json:"dice"`
	Point  int    `json:"point"`
	Phase  string `json:"phase"`
}
