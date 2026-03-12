package blackjack

import (
	"errors"
	"fmt"

	"github.com/game_engine/card-games/internal/games/common"
)

// Action represents a player's action in blackjack
type Action int

const (
	ActionHit Action = iota
	ActionStand
	ActionDouble
	ActionSplit
	ActionSurrender
)

// GameState represents the current state of a blackjack game
type GameState int

const (
	StateWaiting GameState = iota
	StatePlayerTurn
	StateDealerTurn
	StateGameOver
)

// Result represents the outcome of a blackjack hand
type Result int

const (
	ResultPending Result = iota
	ResultWin
	ResultLoss
	ResultPush
	ResultBlackjack
	ResultSurrender
)

// Player represents a player in the game
type Player struct {
	ID            string  `json:"id"`
	Hands         []*Hand `json:"hands"`
	CurrentHand   int     `json:"current_hand"`
	InsuranceBet  int64   `json:"insurance_bet"`
	CurrentAction Action  `json:"current_action"`
	Result        Result  `json:"result"`
}

// Hand represents a player's hand in blackjack
type Hand struct {
	Cards       []common.Card `json:"cards"`
	Bet         int64         `json:"bet"`
	IsDoubled   bool          `json:"is_doubled"`
	IsSplit     bool          `json:"is_split"`
	CanSplit    bool          `json:"can_split"`
	Surrendered bool          `json:"surrendered"`
	Result      Result        `json:"result"`
}

// Dealer represents the dealer
type Dealer struct {
	HoleCard  common.Card   `json:"hole_card"`
	Hand      []common.Card `json:"hand"`
	GameState GameState     `json:"game_state"`
}

// Game represents a blackjack game
type Game struct {
	ID              string
	Shoe            *common.Shoe
	Dealer          *Dealer
	Players         map[string]*Player
	Config          *Config
	GameState       GameState
	CurrentPlayerID string
}

// Config holds blackjack configuration
type Config struct {
	AllowSurrender       bool
	AllowLateSurrender   bool
	DealerStandsOnSoft17 bool
	MaxSplits            int
	BlackjackPayout      float64
}

// NewGame creates a new blackjack game
func NewGame(id string, shoe *common.Shoe, config *Config) *Game {
	return &Game{
		ID:        id,
		Shoe:      shoe,
		Dealer:    &Dealer{},
		Players:   make(map[string]*Player),
		Config:    config,
		GameState: StateWaiting,
	}
}

// AddPlayer adds a player to the game
func (g *Game) AddPlayer(playerID string, bet int64) error {
	if _, exists := g.Players[playerID]; exists {
		return errors.New("player already exists")
	}

	// Draw initial cards
	card1, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	card2, err := g.Shoe.Draw()
	if err != nil {
		return err
	}

	player := &Player{
		ID: playerID,
		Hands: []*Hand{
			{
				Cards:     []common.Card{card1, card2},
				Bet:       bet,
				IsDoubled: false,
				IsSplit:   false,
				CanSplit:  card1.Rank == card2.Rank,
				Result:    ResultPending,
			},
		},
		CurrentHand: 0,
		Result:      ResultPending,
	}

	g.Players[playerID] = player
	return nil
}

// Start starts the game
func (g *Game) Start() error {
	// Dealer's hole card
	holeCard, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	g.Dealer.HoleCard = holeCard

	// Check for dealer blackjack
	dealerTotal := calculateTotal(g.Dealer.Hand)
	if dealerTotal == 21 {
		g.GameState = StateGameOver
		g.settleDealerBlackjack()
		return nil
	}

	// Check for player blackjacks
	g.GameState = StatePlayerTurn
	for _, player := range g.Players {
		for _, hand := range player.Hands {
			if isBlackjack(hand.Cards) {
				hand.Result = ResultBlackjack
			}
		}
	}

	// Set first player
	g.setNextPlayer()

	return nil
}

// PlayerAction performs an action for the current player
func (g *Game) PlayerAction(playerID string, action Action) error {
	if g.GameState != StatePlayerTurn {
		return errors.New("not player's turn")
	}

	if g.CurrentPlayerID != playerID {
		return errors.New("not current player")
	}

	player, exists := g.Players[playerID]
	if !exists {
		return errors.New("player not found")
	}

	hand := player.Hands[player.CurrentHand]
	if hand == nil {
		return errors.New("hand not found")
	}

	switch action {
	case ActionHit:
		return g.hit(player, hand)
	case ActionStand:
		return g.stand(player)
	case ActionDouble:
		return g.double(player, hand)
	case ActionSplit:
		return g.split(player, hand)
	case ActionSurrender:
		return g.surrender(player, hand)
	default:
		return errors.New("invalid action")
	}
}

// Hit draws a card
func (g *Game) hit(player *Player, hand *Hand) error {
	card, err := g.Shoe.Draw()
	if err != nil {
		return err
	}

	hand.Cards = append(hand.Cards, card)

	// Check if busted
	total := calculateTotal(hand.Cards)
	if total > 21 {
		hand.Result = ResultLoss
		return g.nextHand()
	}

	// Check if can still hit
	return nil
}

// Stand moves to the next hand
func (g *Game) stand(player *Player) error {
	return g.nextHand()
}

// Double doubles the bet and draws one card
func (g *Game) double(player *Player, hand *Hand) error {
	if hand.IsDoubled {
		return errors.New("already doubled")
	}

	hand.Bet *= 2
	hand.IsDoubled = true

	card, err := g.Shoe.Draw()
	if err != nil {
		return err
	}

	hand.Cards = append(hand.Cards, card)

	// Check if busted
	total := calculateTotal(hand.Cards)
	if total > 21 {
		hand.Result = ResultLoss
	}

	return g.nextHand()
}

// Split splits the hand into two
func (g *Game) split(player *Player, hand *Hand) error {
	if !hand.CanSplit {
		return errors.New("cannot split")
	}

	if len(player.Hands) >= g.Config.MaxSplits {
		return errors.New("max splits reached")
	}

	// Get the two cards
	card1 := hand.Cards[0]
	card2 := hand.Cards[1]

	// Create two new hands
	newHand1 := &Hand{
		Cards:    []common.Card{card1},
		Bet:      hand.Bet,
		IsSplit:  true,
		CanSplit: card1.Rank == card1.Rank, // Can split Aces only once
		Result:   ResultPending,
	}

	newHand2 := &Hand{
		Cards:    []common.Card{card2},
		Bet:      hand.Bet,
		IsSplit:  true,
		CanSplit: card2.Rank == card2.Rank,
		Result:   ResultPending,
	}

	// Draw one card for each new hand
	draw1, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	newHand1.Cards = append(newHand1.Cards, draw1)

	draw2, err := g.Shoe.Draw()
	if err != nil {
		return err
	}
	newHand2.Cards = append(newHand2.Cards, draw2)

	// Replace the original hand
	player.Hands[player.CurrentHand] = newHand1
	player.Hands = append(player.Hands, newHand2)

	return nil
}

// Surrender forfeits half the bet
func (g *Game) surrender(player *Player, hand *Hand) error {
	if !g.Config.AllowSurrender {
		return errors.New("surrender not allowed")
	}

	hand.Surrendered = true
	hand.Result = ResultSurrender
	hand.Bet /= 2 // Return half

	return g.nextHand()
}

// DealerPlay plays the dealer's hand
func (g *Game) DealerPlay() error {
	if g.GameState != StateDealerTurn {
		return errors.New("not dealer's turn")
	}

	// Reveal hole card
	g.Dealer.Hand = append(g.Dealer.Hand, g.Dealer.HoleCard)

	// Dealer draws according to rules
	for {
		total := calculateTotal(g.Dealer.Hand)
		soft := isSoft(g.Dealer.Hand)

		// Dealer rules
		if total > 21 {
			break // Busted
		}
		if total > 17 {
			break // Stand
		}
		if total == 17 && soft && !g.Config.DealerStandsOnSoft17 {
			// Hit on soft 17
		}
		if total < 17 {
			// Hit
		} else {
			break
		}

		card, err := g.Shoe.Draw()
		if err != nil {
			return err
		}
		g.Dealer.Hand = append(g.Dealer.Hand, card)
	}

	g.GameState = StateGameOver
	g.settle()

	return nil
}

// calculateTotal calculates the total of a hand
func calculateTotal(cards []common.Card) int {
	total := 0
	aces := 0

	for _, card := range cards {
		if card.IsAce() {
			aces++
			total += 11
		} else {
			total += card.Value()
		}
	}

	for aces > 0 && total > 21 {
		total -= 10
		aces--
	}

	return total
}

// isSoft returns true if the hand is soft (Ace counted as 11)
func isSoft(cards []common.Card) bool {
	total := 0
	aces := 0

	for _, card := range cards {
		if card.IsAce() {
			aces++
			total += 11
		} else {
			total += card.Value()
		}
	}

	return aces > 0 && total <= 21
}

// isBlackjack returns true if the hand is a natural blackjack
func isBlackjack(cards []common.Card) bool {
	if len(cards) != 2 {
		return false
	}

	card1, card2 := cards[0], cards[1]
	hasAce := card1.IsAce() || card2.IsAce()
	hasTen := card1.Value() == 10 || card2.Value() == 10

	return hasAce && hasTen
}

// nextHand moves to the next hand
func (g *Game) nextHand() error {
	player := g.Players[g.CurrentPlayerID]

	// Move to next hand
	player.CurrentHand++

	if player.CurrentHand >= len(player.Hands) {
		// All hands done for this player
		if !g.setNextPlayer() {
			// All players done, dealer's turn
			g.GameState = StateDealerTurn
		}
	}

	return nil
}

// setNextPlayer sets the next player
func (g *Game) setNextPlayer() bool {
	// Simple implementation - find next player with pending hands
	// In production, would maintain player order
	for _, player := range g.Players {
		if player.CurrentHand < len(player.Hands) {
			hand := player.Hands[player.CurrentHand]
			if hand.Result == ResultPending && !hand.Surrendered {
				g.CurrentPlayerID = player.ID
				return true
			}
		}
	}
	return false
}

// settle settles all bets
func (g *Game) settle() {
	dealerTotal := calculateTotal(g.Dealer.Hand)
	dealerBusted := dealerTotal > 21

	for _, player := range g.Players {
		for _, hand := range player.Hands {
			if hand.Result != ResultPending && hand.Result != ResultBlackjack {
				continue
			}

			playerTotal := calculateTotal(hand.Cards)

			if dealerBusted {
				if playerTotal <= 21 {
					hand.Result = ResultWin
				} else {
					hand.Result = ResultLoss
				}
			} else {
				if playerTotal > 21 {
					hand.Result = ResultLoss
				} else if playerTotal > dealerTotal {
					hand.Result = ResultWin
				} else if playerTotal < dealerTotal {
					hand.Result = ResultLoss
				} else {
					hand.Result = ResultPush
				}
			}
		}
	}
}

// settleDealerBlackjack settles the game when dealer has blackjack
func (g *Game) settleDealerBlackjack() {
	for _, player := range g.Players {
		for _, hand := range player.Hands {
			if hand.Result == ResultBlackjack {
				hand.Result = ResultPush
			} else if hand.Result == ResultPending {
				hand.Result = ResultLoss
			}
		}
	}
}

// GetState returns the current game state
func (g *Game) GetState() *GameStateResponse {
	response := &GameStateResponse{
		GameID:          g.ID,
		GameState:       string(g.GameState),
		DealerHand:      g.Dealer.Hand,
		DealerTotal:     calculateTotal(g.Dealer.Hand),
		Players:         make([]PlayerState, 0),
		CurrentPlayerID: g.CurrentPlayerID,
	}

	for _, player := range g.Players {
		playerState := PlayerState{
			ID:          player.ID,
			Hands:       make([]HandState, len(player.Hands)),
			CurrentHand: player.CurrentHand,
		}

		for i, hand := range player.Hands {
			playerState.Hands[i] = HandState{
				Cards:       hand.Cards,
				Bet:         hand.Bet,
				Total:       calculateTotal(hand.Cards),
				IsDoubled:   hand.IsDoubled,
				IsSplit:     hand.IsSplit,
				Surrendered: hand.Surrendered,
				Result:      string(hand.Result),
			}
		}

		response.Players = append(response.Players, playerState)
	}

	return response
}

// GameStateResponse represents the game state for clients
type GameStateResponse struct {
	GameID          string        `json:"game_id"`
	GameState       string        `json:"game_state"`
	DealerHand      []common.Card `json:"dealer_hand"`
	DealerTotal     int           `json:"dealer_total"`
	Players         []PlayerState `json:"players"`
	CurrentPlayerID string        `json:"current_player_id"`
}

// PlayerState represents a player's state
type PlayerState struct {
	ID          string      `json:"id"`
	Hands       []HandState `json:"hands"`
	CurrentHand int         `json:"current_hand"`
}

// HandState represents a hand's state
type HandState struct {
	Cards       []common.Card `json:"cards"`
	Bet         int64         `json:"bet"`
	Total       int           `json:"total"`
	IsDoubled   bool          `json:"is_doubled"`
	IsSplit     bool          `json:"is_split"`
	Surrendered bool          `json:"surrendered"`
	Result      string        `json:"result"`
}

// String returns a string representation of Action
func (a Action) String() string {
	switch a {
	case ActionHit:
		return "hit"
	case ActionStand:
		return "stand"
	case ActionDouble:
		return "double"
	case ActionSplit:
		return "split"
	case ActionSurrender:
		return "surrender"
	default:
		return "unknown"
	}
}

// String returns a string representation of Result
func (r Result) String() string {
	switch r {
	case ResultPending:
		return "pending"
	case ResultWin:
		return "win"
	case ResultLoss:
		return "loss"
	case ResultPush:
		return "push"
	case ResultBlackjack:
		return "blackjack"
	case ResultSurrender:
		return "surrender"
	default:
		return "unknown"
	}
}

// FormatHand formats a hand for display
func FormatHand(cards []common.Card) string {
	result := ""
	for i, card := range cards {
		if i > 0 {
			result += ", "
		}
		result += card.String()
	}
	return result
}

// GetWinnings calculates the winnings for a hand
func GetWinnings(hand *Hand, payout float64) int64 {
	if hand.Result != ResultWin && hand.Result != ResultBlackjack {
		return -hand.Bet
	}

	if hand.Result == ResultBlackjack {
		return int64(float64(hand.Bet) * (1 + payout))
	}

	return hand.Bet
}

func init() {
	// Register string converters
	_ = fmt.Sprintf("%s", ActionHit)
	_ = fmt.Sprintf("%s", ResultWin)
}
