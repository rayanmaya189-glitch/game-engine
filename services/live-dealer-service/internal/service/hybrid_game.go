package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

// GameMode represents the current game execution mode
type GameMode string

const (
	GameModeLive GameMode = "live" // Real dealer, live video
	GameModeRNG  GameMode = "rng"  // RNG fallback, no live video
)

// GameResult represents the outcome of a game round
type GameResult struct {
	TableID   string      `json:"table_id"`
	RoundID   string      `json:"round_id"`
	GameMode  GameMode    `json:"game_mode"`
	DealerID  string      `json:"dealer_id,omitempty"`
	Cards     []string    `json:"cards,omitempty"`
	DiceValue []int       `json:"dice_value,omitempty"`
	Roulette  int         `json:"roulette,omitempty"`
	Outcome   string      `json:"outcome"`
	Payouts   map[string]float64 `json:"payouts"`
	Proof     string      `json:"proof,omitempty"` // Provably fair proof for RNG mode
	CreatedAt time.Time   `json:"created_at"`
}

// HybridGameService manages game execution with live/RNG fallback
type HybridGameService struct {
	liveDealer *LiveDealerService
	rngService *RNGService
}

// NewHybridGameService creates a new hybrid game service
func NewHybridGameService(liveDealer *LiveDealerService, rngService *RNGService) *HybridGameService {
	return &HybridGameService{
		liveDealer: liveDealer,
		rngService: rngService,
	}
}

// DetermineGameMode checks if live dealer is available, falls back to RNG
func (s *HybridGameService) DetermineGameMode(ctx context.Context, tableID string) (GameMode, string, error) {
	table, err := s.liveDealer.GetSession(ctx, tableID)
	if err != nil {
		// No live session, use RNG
		return GameModeRNG, "", nil
	}

	// Check if dealer is assigned and available
	if table.DealerID == "" {
		return GameModeRNG, "", nil
	}

	dealer, err := s.liveDealer.repo.GetDealer(ctx, table.DealerID)
	if err != nil {
		// Dealer not found, fall back to RNG
		return GameModeRNG, "", nil
	}

	if dealer.Status != "available" && dealer.Status != "busy" {
		// Dealer not available, fall back to RNG
		return GameModeRNG, "", nil
	}

	// Check if video stream is active
	stream, err := s.liveDealer.repo.GetStreamInfo(ctx, tableID)
	if err != nil || stream == nil || !stream.IsLive {
		// No active stream, fall back to RNG
		return GameModeRNG, "", nil
	}

	return GameModeLive, table.DealerID, nil
}

// PlayBlackjackRound plays a blackjack round (live or RNG)
func (s *HybridGameService) PlayBlackjackRound(ctx context.Context, tableID string, players []string) (*GameResult, error) {
	mode, dealerID, err := s.DetermineGameMode(ctx, tableID)
	if err != nil {
		return nil, err
	}

	roundID := generateRoundID()

	if mode == GameModeLive {
		// Live mode: dealer controls the game
		return s.playBlackjackLive(ctx, tableID, roundID, dealerID, players)
	}

	// RNG mode: provably fair card generation
	return s.playBlackjackRNG(ctx, tableID, roundID, players)
}

func (s *HybridGameService) playBlackjackLive(ctx context.Context, tableID, roundID, dealerID string, players []string) (*GameResult, error) {
	// In live mode, cards are dealt by the real dealer
	// This function waits for dealer actions via WebSocket/NATS
	return &GameResult{
		TableID:  tableID,
		RoundID:  roundID,
		GameMode: GameModeLive,
		DealerID: dealerID,
		Outcome:  "awaiting_dealer",
	}, nil
}

func (s *HybridGameService) playBlackjackRNG(ctx context.Context, tableID, roundID string, players []string) (*GameResult, error) {
	deck := s.generateDeck()
	shuffleDeck(deck)

	// Deal 2 cards to each player + dealer
	cards := make([]string, 0)
	dealerCards := []string{deck[0], deck[1]}
	cards = append(cards, dealerCards...)

	deckIndex := 2
	for range players {
		playerCards := []string{deck[deckIndex], deck[deckIndex+1]}
		cards = append(cards, playerCards...)
		deckIndex += 2
	}

	// Generate provably fair proof
	proof := s.generateProof(roundID, cards)

	return &GameResult{
		TableID:   tableID,
		RoundID:   roundID,
		GameMode:  GameModeRNG,
		Cards:     cards,
		Outcome:   "dealt",
		Proof:     proof,
		CreatedAt: time.Now(),
	}, nil
}

// PlayRouletteRound plays a roulette round (live or RNG)
func (s *HybridGameService) PlayRouletteRound(ctx context.Context, tableID string, bets map[string]float64) (*GameResult, error) {
	mode, dealerID, err := s.DetermineGameMode(ctx, tableID)
	if err != nil {
		return nil, err
	}

	roundID := generateRoundID()

	if mode == GameModeLive {
		return &GameResult{
			TableID:  tableID,
			RoundID:  roundID,
			GameMode: GameModeLive,
			DealerID: dealerID,
			Outcome:  "awaiting_spin",
		}, nil
	}

	// RNG mode: random number 0-36
	num, err := rand.Int(rand.Reader, big.NewInt(37))
	if err != nil {
		return nil, fmt.Errorf("failed to generate random number: %w", err)
	}

	result := int(num.Int64())
	payouts := s.calculateRoulettePayouts(result, bets)
	proof := s.generateProof(roundID, []string{fmt.Sprintf("%d", result)})

	return &GameResult{
		TableID:   tableID,
		RoundID:   roundID,
		GameMode:  GameModeRNG,
		Roulette:  result,
		Outcome:   fmt.Sprintf("number_%d", result),
		Payouts:   payouts,
		Proof:     proof,
		CreatedAt: time.Now(),
	}, nil
}

// PlayBaccaratRound plays a baccarat round (live or RNG)
func (s *HybridGameService) PlayBaccaratRound(ctx context.Context, tableID string, bets map[string]float64) (*GameResult, error) {
	mode, dealerID, err := s.DetermineGameMode(ctx, tableID)
	if err != nil {
		return nil, err
	}

	roundID := generateRoundID()

	if mode == GameModeLive {
		return &GameResult{
			TableID:  tableID,
			RoundID:  roundID,
			GameMode: GameModeLive,
			DealerID: dealerID,
			Outcome:  "awaiting_deal",
		}, nil
	}

	deck := s.generateDeck()
	shuffleDeck(deck)

	playerCards := []string{deck[0], deck[1], deck[4]}
	bankerCards := []string{deck[2], deck[3], deck[5]}

	playerScore := s.calculateBaccaratScore(playerCards[:2])
	bankerScore := s.calculateBaccaratScore(bankerCards[:2])

	var outcome string
	if playerScore > bankerScore {
		outcome = "player_wins"
	} else if bankerScore > playerScore {
		outcome = "banker_wins"
	} else {
		outcome = "tie"
	}

	payouts := s.calculateBaccaratPayouts(outcome, bets)
	cards := append(playerCards, bankerCards...)
	proof := s.generateProof(roundID, cards)

	return &GameResult{
		TableID:   tableID,
		RoundID:   roundID,
		GameMode:  GameModeRNG,
		Cards:     cards,
		Outcome:   outcome,
		Payouts:   payouts,
		Proof:     proof,
		CreatedAt: time.Now(),
	}, nil
}

// PlayDiceRound plays a dice round (Sic Bo style, live or RNG)
func (s *HybridGameService) PlayDiceRound(ctx context.Context, tableID string, bets map[string]float64) (*GameResult, error) {
	mode, dealerID, err := s.DetermineGameMode(ctx, tableID)
	if err != nil {
		return nil, err
	}

	roundID := generateRoundID()

	if mode == GameModeLive {
		return &GameResult{
			TableID:  tableID,
			RoundID:  roundID,
			GameMode: GameModeLive,
			DealerID: dealerID,
			Outcome:  "awaiting_roll",
		}, nil
	}

	dice := make([]int, 3)
	for i := range dice {
		num, _ := rand.Int(rand.Reader, big.NewInt(6))
		dice[i] = int(num.Int64()) + 1
	}

	total := dice[0] + dice[1] + dice[2]
	var outcome string
	if total <= 10 {
		outcome = "small"
	} else {
		outcome = "big"
	}

	payouts := s.calculateDicePayouts(dice, bets)
	proof := s.generateProof(roundID, []string{
		fmt.Sprintf("%d-%d-%d", dice[0], dice[1], dice[2]),
	})

	return &GameResult{
		TableID:   tableID,
		RoundID:   roundID,
		GameMode:  GameModeRNG,
		DiceValue: dice,
		Outcome:   outcome,
		Payouts:   payouts,
		Proof:     proof,
		CreatedAt: time.Now(),
	}, nil
}
