package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/game_engine/live-dealer-service/internal/config"
	"github.com/game_engine/live-dealer-service/internal/model"
	"github.com/game_engine/live-dealer-service/internal/repository"
)

type LiveDealerService struct {
	repo *repository.LiveDealerRepository
	cfg  *config.Config
}

func NewLiveDealerService(repo *repository.LiveDealerRepository, cfg *config.Config) *LiveDealerService {
	return &LiveDealerService{repo: repo, cfg: cfg}
}

// --- Session (Table) Management ---

func (s *LiveDealerService) CreateSession(ctx context.Context, gameType, dealerID string, minBet, maxBet float64) (*model.Table, error) {
	if gameType == "" {
		return nil, fmt.Errorf("game type is required")
	}
	if minBet <= 0 || maxBet <= 0 {
		return nil, fmt.Errorf("invalid bet limits")
	}
	if minBet > maxBet {
		return nil, fmt.Errorf("min bet cannot exceed max bet")
	}

	table := &model.Table{
		TableID:     generateID(),
		GameType:    gameType,
		DealerID:    dealerID,
		Status:      "open",
		MinBet:      minBet,
		MaxBet:      maxBet,
		MaxSeats:    s.cfg.Dealer.MaxPlayersPerTable,
		CurrentSeat: 0,
		DeckCount:   8,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.CreateTable(ctx, table); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	if err := s.repo.CacheSessionState(ctx, table.TableID, table); err != nil {
		return nil, fmt.Errorf("failed to cache session state: %w", err)
	}

	return table, nil
}

func (s *LiveDealerService) GetSession(ctx context.Context, tableID string) (*model.Table, error) {
	var table model.Table
	if err := s.repo.GetCachedSessionState(ctx, tableID, &table); err == nil {
		return &table, nil
	}

	t, err := s.repo.GetTable(ctx, tableID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}
	return t, nil
}

func (s *LiveDealerService) ListSessions(ctx context.Context, gameType, status string) ([]*model.Table, error) {
	return s.repo.ListTables(ctx, gameType, status)
}

func (s *LiveDealerService) EndSession(ctx context.Context, tableID string) error {
	table, err := s.repo.GetTable(ctx, tableID)
	if err != nil {
		return fmt.Errorf("session not found: %w", err)
	}

	table.Status = "closed"
	table.UpdatedAt = time.Now()

	if err := s.repo.UpdateTable(ctx, table); err != nil {
		return fmt.Errorf("failed to update table: %w", err)
	}

	if err := s.repo.DeleteSessionState(ctx, tableID); err != nil {
		return fmt.Errorf("failed to clear session cache: %w", err)
	}

	return nil
}

// --- Player Management ---

func (s *LiveDealerService) JoinSession(ctx context.Context, tableID, playerID string, chips float64) (*model.Player, error) {
	table, err := s.repo.GetTable(ctx, tableID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if table.Status != "open" {
		return nil, fmt.Errorf("session is not open")
	}

	if table.CurrentSeat >= table.MaxSeats {
		return nil, fmt.Errorf("session is full")
	}

	existingPlayers, err := s.repo.GetPlayersByTable(ctx, tableID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing players: %w", err)
	}
	for _, p := range existingPlayers {
		if p.PlayerID == playerID {
			return nil, fmt.Errorf("player already in session")
		}
	}

	player := &model.Player{
		PlayerID:   playerID,
		TableID:    tableID,
		SeatNumber: table.CurrentSeat + 1,
		Chips:      chips,
		JoinedAt:   time.Now(),
		IsFinished: false,
	}

	if err := s.repo.CreatePlayer(ctx, player); err != nil {
		return nil, fmt.Errorf("failed to add player: %w", err)
	}

	table.CurrentSeat++
	table.UpdatedAt = time.Now()
	if err := s.repo.UpdateTable(ctx, table); err != nil {
		return nil, fmt.Errorf("failed to update table: %w", err)
	}

	if err := s.repo.SetPlayerOnline(ctx, tableID, playerID); err != nil {
		return nil, fmt.Errorf("failed to mark player online: %w", err)
	}

	if err := s.repo.CacheSessionState(ctx, tableID, table); err != nil {
		return nil, fmt.Errorf("failed to update session cache: %w", err)
	}

	return player, nil
}

func (s *LiveDealerService) LeaveSession(ctx context.Context, tableID, playerID string) error {
	table, err := s.repo.GetTable(ctx, tableID)
	if err != nil {
		return fmt.Errorf("session not found: %w", err)
	}

	if err := s.repo.DeletePlayer(ctx, playerID); err != nil {
		return fmt.Errorf("failed to remove player: %w", err)
	}

	if table.CurrentSeat > 0 {
		table.CurrentSeat--
		table.UpdatedAt = time.Now()
		if err := s.repo.UpdateTable(ctx, table); err != nil {
			return fmt.Errorf("failed to update table: %w", err)
		}
	}

	if err := s.repo.RemovePlayerOnline(ctx, tableID, playerID); err != nil {
		return fmt.Errorf("failed to mark player offline: %w", err)
	}

	if err := s.repo.CacheSessionState(ctx, tableID, table); err != nil {
		return fmt.Errorf("failed to update session cache: %w", err)
	}

	return nil
}

func (s *LiveDealerService) GetSessionPlayers(ctx context.Context, tableID string) ([]*model.Player, error) {
	return s.repo.GetPlayersByTable(ctx, tableID)
}

// --- Game Round Management ---

func (s *LiveDealerService) StartRound(ctx context.Context, tableID string) (*model.GameState, error) {
	table, err := s.repo.GetTable(ctx, tableID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if table.Status != "open" {
		return nil, fmt.Errorf("session is not open")
	}

	existing, _ := s.repo.GetActiveRound(ctx, tableID)
	if existing != nil {
		return nil, fmt.Errorf("round already in progress")
	}

	gameState := &model.GameState{
		TableID:   tableID,
		RoundID:   generateID(),
		Phase:     "betting",
		Cards:     []string{},
		DealerCards: []string{},
		Pot:       0,
		StartTime: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.CreateGameState(ctx, gameState); err != nil {
		return nil, fmt.Errorf("failed to create round: %w", err)
	}

	return gameState, nil
}

func (s *LiveDealerService) PlaceBet(ctx context.Context, playerID, roundID, betType string, amount float64) (*model.Bet, error) {
	gameState, err := s.repo.GetGameState(ctx, roundID)
	if err != nil {
		return nil, fmt.Errorf("round not found: %w", err)
	}

	if gameState.Phase != "betting" {
		return nil, fmt.Errorf("betting phase is over")
	}

	table, err := s.repo.GetTable(ctx, gameState.TableID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	if amount < table.MinBet || amount > table.MaxBet {
		return nil, fmt.Errorf("bet amount outside table limits")
	}

	players, err := s.repo.GetPlayersByTable(ctx, gameState.TableID)
	if err != nil {
		return nil, fmt.Errorf("failed to get players: %w", err)
	}

	var player *model.Player
	for _, p := range players {
		if p.PlayerID == playerID {
			player = p
			break
		}
	}
	if player == nil {
		return nil, fmt.Errorf("player not in session")
	}

	if amount > player.Chips {
		return nil, fmt.Errorf("insufficient chips")
	}

	bet := &model.Bet{
		BetID:     generateID(),
		PlayerID:  playerID,
		TableID:   gameState.TableID,
		RoundID:   roundID,
		BetType:   betType,
		BetAmount: amount,
		Odds:      1.0,
		Potential: amount,
		Result:    "pending",
		PlacedAt:  time.Now(),
	}

	if err := s.repo.CreateBet(ctx, bet); err != nil {
		return nil, fmt.Errorf("failed to place bet: %w", err)
	}

	player.Chips -= amount
	player.CurrentBet += amount
	if err := s.repo.UpdatePlayer(ctx, player); err != nil {
		return nil, fmt.Errorf("failed to update player: %w", err)
	}

	gameState.Pot += amount
	if err := s.repo.UpdateGameState(ctx, gameState); err != nil {
		return nil, fmt.Errorf("failed to update game state: %w", err)
	}

	return bet, nil
}

func (s *LiveDealerService) EndBetting(ctx context.Context, roundID string) error {
	gameState, err := s.repo.GetGameState(ctx, roundID)
	if err != nil {
		return fmt.Errorf("round not found: %w", err)
	}

	if gameState.Phase != "betting" {
		return fmt.Errorf("not in betting phase")
	}

	gameState.Phase = "playing"
	gameState.UpdatedAt = time.Now()

	return s.repo.UpdateGameState(ctx, gameState)
}

func (s *LiveDealerService) ResolveRound(ctx context.Context, roundID, winner string) error {
	gameState, err := s.repo.GetGameState(ctx, roundID)
	if err != nil {
		return fmt.Errorf("round not found: %w", err)
	}

	if gameState.Phase == "finished" {
		return fmt.Errorf("round already finished")
	}

	gameState.Winner = winner
	gameState.Phase = "resolving"
	gameState.EndTime = time.Now()
	gameState.UpdatedAt = time.Now()

	bets, err := s.repo.GetBetsByRound(ctx, roundID)
	if err != nil {
		return fmt.Errorf("failed to get bets: %w", err)
	}

	var totalPayout float64
	for _, bet := range bets {
		switch winner {
		case "player":
			bet.Result = "won"
			bet.Payout = bet.BetAmount * 2
			totalPayout += bet.Payout
		case "push":
			bet.Result = "void"
			bet.Payout = bet.BetAmount
		default:
			bet.Result = "lost"
			bet.Payout = 0
		}
		bet.ResultedAt = time.Now()

		if err := s.repo.UpdateBet(ctx, bet); err != nil {
			return fmt.Errorf("failed to update bet: %w", err)
		}

		if winner == "player" || winner == "push" {
			players, _ := s.repo.GetPlayersByTable(ctx, gameState.TableID)
			for _, p := range players {
				if p.PlayerID == bet.PlayerID {
					p.Chips += bet.Payout
					s.repo.UpdatePlayer(ctx, p)
					break
				}
			}
		}
	}

	gameState.Payout = totalPayout
	gameState.Phase = "finished"

	return s.repo.UpdateGameState(ctx, gameState)
}

func (s *LiveDealerService) GetRound(ctx context.Context, roundID string) (*model.GameState, error) {
	return s.repo.GetGameState(ctx, roundID)
}

// --- Dealer Management ---

func (s *LiveDealerService) RegisterDealer(ctx context.Context, name, language string) (*model.Dealer, error) {
	dealer := &model.Dealer{
		DealerID:   generateID(),
		Name:       name,
		Language:   language,
		Status:     "available",
		ShiftStart: time.Now(),
	}

	if err := s.repo.CreateDealer(ctx, dealer); err != nil {
		return nil, fmt.Errorf("failed to register dealer: %w", err)
	}
	return dealer, nil
}

func (s *LiveDealerService) AssignDealer(ctx context.Context, dealerID, tableID string) error {
	dealer, err := s.repo.GetDealer(ctx, dealerID)
	if err != nil {
		return fmt.Errorf("dealer not found: %w", err)
	}

	table, err := s.repo.GetTable(ctx, tableID)
	if err != nil {
		return fmt.Errorf("session not found: %w", err)
	}

	dealer.Status = "busy"
	dealer.TableID = tableID
	if err := s.repo.UpdateDealer(ctx, dealer); err != nil {
		return fmt.Errorf("failed to update dealer: %w", err)
	}

	table.DealerID = dealerID
	table.DealerName = dealer.Name
	table.UpdatedAt = time.Now()
	if err := s.repo.UpdateTable(ctx, table); err != nil {
		return fmt.Errorf("failed to update table: %w", err)
	}

	return nil
}

func (s *LiveDealerService) GetDealerByID(ctx context.Context, dealerID string) (*model.Dealer, error) {
	return s.repo.GetDealer(ctx, dealerID)
}

func (s *LiveDealerService) ListDealers(ctx context.Context, status string) ([]*model.Dealer, error) {
	return s.repo.ListDealers(ctx, status)
}

// --- Video Streaming Coordination ---

func (s *LiveDealerService) GetStreamInfo(ctx context.Context, tableID string) (*StreamInfo, error) {
	table, err := s.repo.GetTable(ctx, tableID)
	if err != nil {
		return nil, fmt.Errorf("session not found: %w", err)
	}

	onlinePlayers, err := s.repo.GetOnlinePlayers(ctx, tableID)
	if err != nil {
		return nil, fmt.Errorf("failed to get online players: %w", err)
	}

	return &StreamInfo{
		TableID:         tableID,
		Bitrate:         s.cfg.Dealer.VideoStreamBitrate,
		ViewerCount:     len(onlinePlayers),
		StreamURL:       fmt.Sprintf("rtmp://stream.internal/%s/live", tableID),
		IsLive:          table.Status == "open",
		MaxPlayersPerTable: s.cfg.Dealer.MaxPlayersPerTable,
	}, nil
}

type StreamInfo struct {
	TableID            string `json:"table_id"`
	Bitrate            int    `json:"bitrate"`
	ViewerCount        int    `json:"viewer_count"`
	StreamURL          string `json:"stream_url"`
	IsLive             bool   `json:"is_live"`
	MaxPlayersPerTable int    `json:"max_players_per_table"`
}

func generateID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return hex.EncodeToString(b)
}
