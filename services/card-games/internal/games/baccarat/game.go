package baccarat

import (
	"errors"

	"github.com/game_engine/card-games/internal/games/common"
)

// BetType represents the type of bet in baccarat
type BetType int

const (
	BetPlayer BetType = iota
	BetBanker
	BetTie
	BetPlayerPair
	BetBankerPair
)

// Outcome represents the outcome of the game
type Outcome int

const (
	OutcomePlayer Outcome = iota
	OutcomeBanker
	OutcomeTie
)

// Result represents a bet result
type Result struct {
	BetType  BetType `json:"bet_type"`
	Outcome  Outcome `json:"outcome"`
	Won      bool    `json:"won"`
	Payout   float64 `json:"payout"`
	Amount   int64   `json:"amount"`
	Winnings int64   `json:"winnings"`
}

// Game represents a baccarat game (Punto Banco variant)
type Game struct {
	ID         string
	Shoe       *common.Shoe
	PlayerHand []common.Card
	BankerHand []common.Card
	Community  []common.Card
	Results    []Result
	Config     *Config
}

// Config holds baccarat configuration
type Config struct {
	Commission float64
	MaxPlayers int
	ShoeSize   int
}

// NewGame creates a new baccarat game
func NewGame(id string, shoe *common.Shoe, config *Config) *Game {
	return &Game{
		ID:         id,
		Shoe:       shoe,
		PlayerHand: make([]common.Card, 0),
		BankerHand: make([]common.Card, 0),
		Community:  make([]common.Card, 0),
		Results:    make([]Result, 0),
		Config:     config,
	}
}

// DealInitialCards deals the initial four cards
func (g *Game) DealInitialCards() error {
	// Deal in order: Player, Banker, Player, Banker
	card1, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	g.PlayerHand = append(g.PlayerHand, card1)

	card2, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	g.BankerHand = append(g.BankerHand, card2)

	card3, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	g.PlayerHand = append(g.PlayerHand, card3)

	card4, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	g.BankerHand = append(g.BankerHand, card4)

	return nil
}

// Evaluate evaluates the hands and determines the winner
func (g *Game) Evaluate() (Outcome, error) {
	playerTotal := g.calculateBaccaratTotal(g.PlayerHand)
	bankerTotal := g.calculateBaccaratTotal(g.BankerHand)

	// Check for natural win (8 or 9)
	playerNatural := playerTotal >= 8
	bankerNatural := bankerTotal >= 8

	if playerNatural || bankerNatural {
		// No more cards are dealt
		return g.determineWinner(playerTotal, bankerTotal), nil
	}

	// Player draws third card rule
	playerDraw := false
	if playerTotal <= 5 {
		playerDraw = true
		card, err := g.Shoe.Draw()
		if err != nil {
			return OutcomeTie, err
		}
		g.PlayerHand = append(g.PlayerHand, card)
		playerTotal = g.calculateBaccaratTotal(g.PlayerHand)
	}

	// Banker draws third card rule
	bankerDraw := g.shouldBankerDraw(bankerTotal, playerDraw, g.PlayerHand)
	if bankerDraw {
		card, err := g.Shoe.Draw()
		if err != nil {
			return OutcomeTie, err
		}
		g.BankerHand = append(g.BankerHand, card)
		bankerTotal = g.calculateBaccaratTotal(g.BankerHand)
	}

	return g.determineWinner(playerTotal, bankerTotal), nil
}

// calculateBaccaratTotal calculates the baccarat point total
func (g *Game) calculateBaccaratTotal(cards []common.Card) int {
	total := 0
	for _, card := range cards {
		value := card.Value()
		if value > 9 {
			value = 0
		}
		total += value
	}
	return total % 10
}

// shouldBankerDraw determines if the banker should draw a third card
func (g *Game) shouldBankerDraw(bankerTotal int, playerDrew bool, playerHand []common.Card) bool {
	// If player didn't draw, banker draws on 0-5
	if !playerDrew {
		return bankerTotal <= 5
	}

	// Player drew a card
	playerThirdCardValue := playerHand[2].Value()
	if playerThirdCardValue > 9 {
		playerThirdCardValue = 0
	}

	// Banker draws based on player third card
	switch bankerTotal {
	case 0, 1, 2:
		return true
	case 3:
		return playerThirdCardValue != 8
	case 4:
		return playerThirdCardValue >= 2 && playerThirdCardValue <= 7
	case 5:
		return playerThirdCardValue >= 4 && playerThirdCardValue <= 7
	case 6:
		return playerThirdCardValue == 6 || playerThirdCardValue == 7
	case 7, 8, 9:
		return false
	}
	return false
}

// determineWinner determines the winner
func (g *Game) determineWinner(playerTotal, bankerTotal int) Outcome {
	if playerTotal > bankerTotal {
		return OutcomePlayer
	} else if bankerTotal > playerTotal {
		return OutcomeBanker
	}
	return OutcomeTie
}

// SettleBets settles all bets based on the outcome
func (g *Game) SettleBets(bets map[BetType]int64) []Result {
	playerTotal := g.calculateBaccaratTotal(g.PlayerHand)
	bankerTotal := g.calculateBaccaratTotal(g.BankerHand)
	outcome := g.determineWinner(playerTotal, bankerTotal)

	// Check for pairs
	playerPair := len(g.PlayerHand) >= 2 && g.PlayerHand[0].Rank == g.PlayerHand[1].Rank
	bankerPair := len(g.BankerHand) >= 2 && g.BankerHand[0].Rank == g.BankerHand[1].Rank

	results := make([]Result, 0)

	for betType, amount := range bets {
		result := Result{
			BetType: betType,
			Amount:  amount,
		}

		switch betType {
		case BetPlayer:
			result.Outcome = outcome
			result.Won = outcome == OutcomePlayer
			if result.Won {
				result.Payout = 1.0 // 1:1
				result.Winnings = amount
			}

		case BetBanker:
			result.Outcome = outcome
			result.Won = outcome == OutcomeBanker
			if result.Won {
				result.Payout = 1.0 - g.Config.Commission // 1:1 minus commission
				result.Winnings = int64(float64(amount) * result.Payout)
			}

		case BetTie:
			result.Outcome = outcome
			result.Won = outcome == OutcomeTie
			if result.Won {
				result.Payout = 8.0 // 8:1
				result.Winnings = amount * 8
			}

		case BetPlayerPair:
			result.Won = playerPair
			if result.Won {
				result.Payout = 11.0 // 11:1
				result.Winnings = amount * 11
			}

		case BetBankerPair:
			result.Won = bankerPair
			if result.Won {
				result.Payout = 11.0 // 11:1
				result.Winnings = amount * 11
			}
		}

		results = append(results, result)
	}

	g.Results = results
	return results
}

// GetState returns the current game state
func (g *Game) GetState() *GameState {
	playerTotal := g.calculateBaccaratTotal(g.PlayerHand)
	bankerTotal := g.calculateBaccaratTotal(g.BankerHand)

	return &GameState{
		GameID:      g.ID,
		PlayerHand:  g.PlayerHand,
		BankerHand:  g.BankerHand,
		PlayerTotal: playerTotal,
		BankerTotal: bankerTotal,
		Results:     g.Results,
	}
}

// GameState represents the game state for clients
type GameState struct {
	GameID      string        `json:"game_id"`
	PlayerHand  []common.Card `json:"player_hand"`
	BankerHand  []common.Card `json:"banker_hand"`
	PlayerTotal int           `json:"player_total"`
	BankerTotal int           `json:"banker_total"`
	Results     []Result      `json:"results"`
}

// ValidateBet validates a bet
func ValidateBet(betType BetType, amount int64) error {
	if amount <= 0 {
		return errors.New("bet amount must be positive")
	}
	if betType < BetPlayer || betType > BetBankerPair {
		return errors.New("invalid bet type")
	}
	return nil
}
