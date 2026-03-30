package games

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// Symbol represents a slot symbol
type Symbol struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Value     int    `json:"value"`
	IsWild    bool   `json:"is_wild"`
	IsScatter bool   `json:"is_scatter"`
	IsBonus   bool   `json:"is_bonus"`
	Weight    int    `json:"weight"`
}

// Payline represents a payline
type Payline struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Positions [][]int `json:"positions"` // [reel][row]
	Payout    int64   `json:"payout"`
	MinSymbol int     `json:"min_symbol"`
}

// GameConfig represents slot game configuration
type GameConfig struct {
	Reels      int     `json:"reels"`
	Rows       int     `json:"rows"`
	MinBet     int64   `json:"min_bet"`
	MaxBet     int64   `json:"max_bet"`
	MinLineBet int64   `json:"min_line_bet"`
	MaxLineBet int64   `json:"max_line_bet"`
	MaxLines   int     `json:"max_lines"`
	RTP        float64 `json:"rtp"`
	Volatility string  `json:"volatility"`
}

// Game represents a slot game
type Game struct {
	ID            string
	Config        *GameConfig
	Symbols       []Symbol
	Paylines      []Payline
	Positions     [][]int // [reel][row]
	Bet           int64
	LineBet       int64
	Lines         int
	Win           int64
	IsComplete    bool
	WinLines      []WinLine
	ScatterWin    int64
	BonusWin      int64
	TotalWin      int64
	FreeSpins     int
	FreeSpinsLeft int
	ProvablyFair  bool
	ServerSeed    string
	ClientSeed    string
	Nonce         int
}

// WinLine represents a winning payline
type WinLine struct {
	PaylineID int    `json:"payline_id"`
	Symbol    string `json:"symbol"`
	Count     int    `json:"count"`
	Payout    int64  `json:"payout"`
	Positions []int  `json:"positions"`
}

// GameState represents the game state
type GameState struct {
	GameID        string     `json:"game_id"`
	Reels         int        `json:"reels"`
	Rows          int        `json:"rows"`
	Symbols       [][]string `json:"symbols"`
	Bet           int64      `json:"bet"`
	LineBet       int64      `json:"line_bet"`
	Lines         int        `json:"lines"`
	Win           int64      `json:"win"`
	TotalWin      int64      `json:"total_win"`
	WinLines      []WinLine  `json:"win_lines"`
	ScatterWin    int64      `json:"scatter_win"`
	BonusWin      int64      `json:"bonus_win"`
	FreeSpins     int        `json:"free_spins"`
	FreeSpinsLeft int        `json:"free_spins_left"`
	IsComplete    bool       `json:"is_complete"`
	ProvablyFair  bool       `json:"provably_fair"`
}

// NewClassicSlotGame creates a classic 3-reel slot game
func NewClassicSlotGame(id string) *Game {
	config := &GameConfig{
		Reels:      3,
		Rows:       3,
		MinBet:     1,
		MaxBet:     100,
		MinLineBet: 1,
		MaxLineBet: 10,
		MaxLines:   1,
		RTP:        95.5,
		Volatility: "medium",
	}

	symbols := []Symbol{
		{ID: "7", Name: "Seven", Value: 7, Weight: 10},
		{ID: "BB", Name: "DoubleBar", Value: 6, Weight: 20},
		{ID: "B", Name: "Bar", Value: 5, Weight: 30},
		{ID: "C", Name: "Cherry", Value: 4, Weight: 40},
		{ID: "LR", Name: "Lemon", Value: 3, Weight: 50},
	}

	paylines := []Payline{
		{ID: 1, Name: "Middle Row", Positions: [][]int{{0, 1}, {1, 1}, {2, 1}}, Payout: 100, MinSymbol: 3},
	}

	return &Game{
		ID:        id,
		Config:    config,
		Symbols:   symbols,
		Paylines:  paylines,
		Positions: make([][]int, config.Reels),
	}
}

// NewVideoSlotGame creates a 5-reel video slot game
func NewVideoSlotGame(id string) *Game {
	config := &GameConfig{
		Reels:      5,
		Rows:       3,
		MinBet:     1,
		MaxBet:     500,
		MinLineBet: 1,
		MaxLineBet: 10,
		MaxLines:   20,
		RTP:        96.5,
		Volatility: "high",
	}

	symbols := []Symbol{
		{ID: "W", Name: "Wild", Value: 10, IsWild: true, Weight: 5},
		{ID: "S", Name: "Scatter", Value: 9, IsScatter: true, Weight: 10},
		{ID: "7", Name: "Seven", Value: 8, Weight: 15},
		{ID: "D", Name: "Diamond", Value: 7, Weight: 20},
		{ID: "H", Name: "Heart", Value: 6, Weight: 25},
		{ID: "C", Name: "Club", Value: 5, Weight: 30},
		{ID: "SP", Name: "Spade", Value: 4, Weight: 35},
		{ID: "A", Name: "Ace", Value: 3, Weight: 40},
		{ID: "K", Name: "King", Value: 2, Weight: 45},
		{ID: "Q", Name: "Queen", Value: 1, Weight: 50},
	}

	paylines := createPaylines(5, 3)

	return &Game{
		ID:        id,
		Config:    config,
		Symbols:   symbols,
		Paylines:  paylines,
		Positions: make([][]int, config.Reels),
	}
}

// SetBet sets the bet amount
func (g *Game) SetBet(lineBet int64, lines int) error {
	if lineBet < g.Config.MinLineBet || lineBet > g.Config.MaxLineBet {
		return errors.New("invalid line bet amount")
	}
	if lines < 1 || lines > g.Config.MaxLines {
		return errors.New("invalid number of lines")
	}

	g.LineBet = lineBet
	g.Lines = lines
	g.Bet = lineBet * int64(lines)

	return nil
}

// Spin performs a spin with random symbols
func (g *Game) Spin() error {
	if g.Bet < g.Config.MinBet {
		return errors.New("bet not set")
	}

	for reel := 0; reel < g.Config.Reels; reel++ {
		g.Positions[reel] = make([]int, g.Config.Rows)
		for row := 0; row < g.Config.Rows; row++ {
			symbolIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(g.Symbols))))
			if err != nil {
				return err
			}
			g.Positions[reel][row] = int(symbolIndex.Int64())
		}
	}

	g.evaluate()

	g.IsComplete = true
	return nil
}

// ProvablyFairSpin performs a spin with provably fair RNG
func (g *Game) ProvablyFairSpin(serverSeed, clientSeed string, nonce int) error {
	if g.Bet < g.Config.MinBet {
		return errors.New("bet not set")
	}

	g.ServerSeed = serverSeed
	g.ClientSeed = clientSeed
	g.Nonce = nonce
	g.ProvablyFair = true

	for reel := 0; reel < g.Config.Reels; reel++ {
		g.Positions[reel] = make([]int, g.Config.Rows)
		for row := 0; row < g.Config.Rows; row++ {
			g.Positions[reel][row] = g.provablyFairRandom(reel, row)
		}
	}

	g.evaluate()
	g.IsComplete = true
	return nil
}

func (g *Game) provablyFairRandom(reel, row int) int {
	seed := g.ServerSeed + g.ClientSeed + string(rune(g.Nonce)) + string(rune(reel)) + string(rune(row))
	hash := 0
	for i, c := range seed {
		hash = hash*31 + int(c) + i
	}
	return hash % len(g.Symbols)
}

// GetState returns the game state
func (g *Game) GetState() *GameState {
	symbols := make([][]string, len(g.Positions))
	for i, reel := range g.Positions {
		symbols[i] = make([]string, len(reel))
		for j, s := range reel {
			symbols[i][j] = g.Symbols[s].ID
		}
	}

	return &GameState{
		GameID:        g.ID,
		Reels:         g.Config.Reels,
		Rows:          g.Config.Rows,
		Symbols:       symbols,
		Bet:           g.Bet,
		LineBet:       g.LineBet,
		Lines:         g.Lines,
		Win:           g.Win,
		TotalWin:      g.TotalWin,
		WinLines:      g.WinLines,
		ScatterWin:    g.ScatterWin,
		BonusWin:      g.BonusWin,
		FreeSpins:     g.FreeSpins,
		FreeSpinsLeft: g.FreeSpinsLeft,
		IsComplete:    g.IsComplete,
		ProvablyFair:  g.ProvablyFair,
	}
}
