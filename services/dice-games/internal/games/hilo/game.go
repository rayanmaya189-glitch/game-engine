package hilo

import (
	"errors"

	"github.com/game_engine/dice-games/internal/games/dice"
)

// Game represents a Hi-Lo game
type Game struct {
	ID           string
	Current      int
	Next         int
	Guess        string
	Correct      bool
	Multiplier   float64
	Payout       float64
	PlayerBet    float64
	IsComplete   bool
	Result       string
	ProvablyFair bool
	ServerSeed   string
	ClientSeed   string
}

// NewGame creates a new Hi-Lo game
func NewGame(id string) *Game {
	return &Game{
		ID:         id,
		Current:    0,
		Next:       0,
		Guess:      "",
		Correct:    false,
		IsComplete: false,
	}
}

// NewGameWithSeeds creates a new Hi-Lo game with provably fair seeds
func NewGameWithSeeds(id, serverSeed, clientSeed string) *Game {
	return &Game{
		ID:           id,
		Current:      0,
		Next:         0,
		Guess:        "",
		Correct:      false,
		IsComplete:   false,
		ProvablyFair: true,
		ServerSeed:   serverSeed,
		ClientSeed:   clientSeed,
	}
}

// SetGuess sets the player's guess (HI, LO, or SEVEN)
func (g *Game) SetGuess(guess string) error {
	guess = normalizeGuess(guess)
	if !isValidGuess(guess) {
		return errors.New("invalid guess: must be HI, LO, or SEVEN")
	}
	g.Guess = guess
	return nil
}

// SetBet sets the player's bet amount
func (g *Game) SetBet(amount float64) error {
	if amount <= 0 {
		return errors.New("bet amount must be positive")
	}
	g.PlayerBet = amount
	return nil
}

// Play plays one round of Hi-Lo
func (g *Game) Play(roller dice.Roller) error {
	if g.Guess == "" {
		return errors.New("guess not set")
	}

	// Roll a single die
	dice, err := roller.Roll(1)
	if err != nil {
		return err
	}
	g.Next = dice[0]

	// Determine result
	g.IsComplete = true
	g.Result = determineResult(g.Next, g.Guess)

	// Check if correct
	g.Correct = (g.Result == "WIN")

	// Calculate multiplier and payout
	g.Multiplier = getMultiplier(g.Guess, g.Next)
	g.Payout = g.PlayerBet * g.Multiplier

	return nil
}

// GetState returns the game state
func (g *Game) GetState() *GameState {
	return &GameState{
		GameID:       g.ID,
		Current:      g.Current,
		Next:         g.Next,
		Guess:        g.Guess,
		Correct:      g.Correct,
		Multiplier:   g.Multiplier,
		Payout:       g.Payout,
		PlayerBet:    g.PlayerBet,
		IsComplete:   g.IsComplete,
		Result:       g.Result,
		ProvablyFair: g.ProvablyFair,
	}
}

// Reset resets the game for next round
func (g *Game) Reset() {
	g.Current = g.Next
	g.Next = 0
	g.Guess = ""
	g.Correct = false
	g.Multiplier = 0
	g.Payout = 0
	g.IsComplete = false
	g.Result = ""
}

// normalizeGuess normalizes the guess string
func normalizeGuess(guess string) string {
	switch guess {
	case "HI", "hi", "Hi", "HIGH", "high", "High":
		return "HI"
	case "LO", "lo", "Lo", "LOW", "low", "Low":
		return "LO"
	case "SEVEN", "seven", "7":
		return "SEVEN"
	default:
		return guess
	}
}

// isValidGuess checks if guess is valid
func isValidGuess(guess string) bool {
	return guess == "HI" || guess == "LO" || guess == "SEVEN"
}

// determineResult determines the result based on dice value and guess
func determineResult(diceValue int, guess string) string {
	switch guess {
	case "HI":
		if diceValue >= 4 {
			return "WIN"
		}
		return "LOSE"
	case "LO":
		if diceValue <= 3 {
			return "WIN"
		}
		return "LOSE"
	case "SEVEN":
		if diceValue == 7 {
			return "WIN"
		}
		return "LOSE"
	default:
		return "ERROR"
	}
}

// getMultiplier returns the payout multiplier
func getMultiplier(guess string, diceValue int) float64 {
	switch guess {
	case "HI":
		if diceValue == 6 {
			return 5.0 // Highest odds
		}
		return 2.0
	case "LO":
		if diceValue == 1 {
			return 5.0 // Highest odds
		}
		return 2.0
	case "SEVEN":
		return 6.0 // Hardest to hit
	default:
		return 0
	}
}

// GameState represents the game state
type GameState struct {
	GameID       string  `json:"game_id"`
	Current      int     `json:"current"`
	Next         int     `json:"next"`
	Guess        string  `json:"guess"`
	Correct      bool    `json:"correct"`
	Multiplier   float64 `json:"multiplier"`
	Payout       float64 `json:"payout"`
	PlayerBet    float64 `json:"player_bet"`
	IsComplete   bool    `json:"is_complete"`
	Result       string  `json:"result"`
	ProvablyFair bool    `json:"provably_fair"`
}
