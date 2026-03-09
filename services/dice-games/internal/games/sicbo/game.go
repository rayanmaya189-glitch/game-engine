package sicbo

// Game represents a Sic Bo game
type Game struct {
	ID     string
	Dice   []int
	Bets   map[string]int64
	Config *Config
}

// Config holds Sic Bo configuration
type Config struct {
	MaxDice int
}

// NewGame creates a new Sic Bo game
func NewGame(id string, diceCount int, config *Config) *Game {
	return &Game{
		ID:     id,
		Dice:   make([]int, diceCount),
		Bets:   make(map[string]int64),
		Config: config,
	}
}

// Roll rolls the dice
func (g *Game) Roll(dice []int) {
	g.Dice = dice
}

// GetTotal returns the sum of dice
func (g *Game) GetTotal() int {
	total := 0
	for _, d := range g.Dice {
		total += d
	}
	return total
}

// GetState returns the game state
func (g *Game) GetState() *GameState {
	return &GameState{
		GameID: g.ID,
		Dice:   g.Dice,
		Total:  g.GetTotal(),
	}
}

// GameState represents the game state
type GameState struct {
	GameID string `json:"game_id"`
	Dice   []int  `json:"dice"`
	Total  int    `json:"total"`
}
