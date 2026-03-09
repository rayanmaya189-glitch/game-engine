package hilo

// Game represents a Hi-Lo game
type Game struct {
	ID      string
	Current int
	Next    int
	Guess   string
	Correct bool
}

// NewGame creates a new Hi-Lo game
func NewGame(id string) *Game {
	return &Game{
		ID:      id,
		Current: 0,
		Next:    0,
		Guess:   "",
	}
}

// SetGuess sets the player's guess
func (g *Game) SetGuess(guess string) {
	g.Guess = guess
}

// Roll sets the next dice value
func (g *Game) Roll(value int) {
	g.Next = value
}

// GetState returns the game state
func (g *Game) GetState() *GameState {
	return &GameState{
		GameID:  g.ID,
		Current: g.Current,
		Next:    g.Next,
		Guess:   g.Guess,
	}
}

// GameState represents the game state
type GameState struct {
	GameID  string `json:"game_id"`
	Current int    `json:"current"`
	Next    int    `json:"next"`
	Guess   string `json:"guess"`
}
