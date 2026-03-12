package service

import (
	"errors"
	"time"

	"github.com/game-engine/live-dealer-service/internal/model"
)

type DealerService struct {
	tables     map[string]*model.Table
	players    map[string]*model.Player
	gameStates map[string]*model.GameState
	dealers    map[string]*model.Dealer
	bets       map[string]*model.Bet
}

func NewDealerService() *DealerService {
	return &DealerService{
		tables:     make(map[string]*model.Table),
		players:    make(map[string]*model.Player),
		gameStates: make(map[string]*model.GameState),
		dealers:    make(map[string]*model.Dealer),
		bets:       make(map[string]*model.Bet),
	}
}

// Table Management

func (s *DealerService) CreateTable(gameType, dealerID string, minBet, maxBet float64, maxSeats int) (*model.Table, error) {
	if gameType == "" {
		return nil, errors.New("game type is required")
	}
	if minBet <= 0 || maxBet <= 0 {
		return nil, errors.New("invalid bet limits")
	}

	table := &model.Table{
		TableID:     generateID(),
		GameType:    gameType,
		DealerID:    dealerID,
		Status:      "open",
		MinBet:      minBet,
		MaxBet:      maxBet,
		MaxSeats:    maxSeats,
		CurrentSeat: 0,
		DeckCount:   8,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.tables[table.TableID] = table
	return table, nil
}

func (s *DealerService) GetTable(tableID string) (*model.Table, error) {
	table, ok := s.tables[tableID]
	if !ok {
		return nil, errors.New("table not found")
	}
	return table, nil
}

func (s *DealerService) ListTables(gameType, status string) []*model.Table {
	var result []*model.Table
	for _, table := range s.tables {
		if gameType != "" && table.GameType != gameType {
			continue
		}
		if status != "" && table.Status != status {
			continue
		}
		result = append(result, table)
	}
	return result
}

func (s *DealerService) UpdateTableStatus(tableID, status string) error {
	table, ok := s.tables[tableID]
	if !ok {
		return errors.New("table not found")
	}
	if status != "open" && status != "closed" && status != "maintenance" {
		return errors.New("invalid status")
	}
	table.Status = status
	table.UpdatedAt = time.Now()
	return nil
}

// Player Management

func (s *DealerService) JoinTable(tableID, playerID string, chips float64) (*model.Player, error) {
	table, ok := s.tables[tableID]
	if !ok {
		return nil, errors.New("table not found")
	}

	if table.Status != "open" {
		return nil, errors.New("table is not open")
	}

	if table.CurrentSeat >= table.MaxSeats {
		return nil, errors.New("table is full")
	}

	// Check if player already at table
	for _, p := range s.players {
		if p.PlayerID == playerID && p.TableID == tableID {
			return nil, errors.New("player already at table")
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

	s.players[playerID] = player
	table.CurrentSeat++

	return player, nil
}

func (s *DealerService) LeaveTable(playerID string) error {
	player, ok := s.players[playerID]
	if !ok {
		return errors.New("player not found at table")
	}

	table, ok := s.tables[player.TableID]
	if ok {
		table.CurrentSeat--
	}

	delete(s.players, playerID)
	return nil
}

func (s *DealerService) GetTablePlayers(tableID string) []*model.Player {
	var result []*model.Player
	for _, p := range s.players {
		if p.TableID == tableID {
			result = append(result, p)
		}
	}
	return result
}

// Game State Management

func (s *DealerService) StartRound(tableID string) (*model.GameState, error) {
	table, ok := s.tables[tableID]
	if !ok {
		return nil, errors.New("table not found")
	}

	if table.Status != "open" {
		return nil, errors.New("table is not open")
	}

	// Check for existing active round
	for _, gs := range s.gameStates {
		if gs.TableID == tableID && gs.Phase != "finished" {
			return nil, errors.New("round already in progress")
		}
	}

	gameState := &model.GameState{
		TableID:   tableID,
		RoundID:   generateID(),
		Phase:     "betting",
		Cards:     []string{},
		Pot:       0,
		StartTime: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.gameStates[gameState.RoundID] = gameState
	return gameState, nil
}

func (s *DealerService) PlaceBet(playerID, roundID, betType string, amount float64) (*model.Bet, error) {
	player, ok := s.players[playerID]
	if !ok {
		return nil, errors.New("player not found")
	}

	gameState, ok := s.gameStates[roundID]
	if !ok {
		return nil, errors.New("round not found")
	}

	if gameState.Phase != "betting" {
		return nil, errors.New("betting phase is over")
	}

	table, ok := s.tables[gameState.TableID]
	if !ok {
		return nil, errors.New("table not found")
	}

	if amount < table.MinBet || amount > table.MaxBet {
		return nil, errors.New("bet amount outside table limits")
	}

	if amount > player.Chips {
		return nil, errors.New("insufficient chips")
	}

	bet := &model.Bet{
		BetID:     generateID(),
		PlayerID:  playerID,
		TableID:   player.TableID,
		RoundID:   roundID,
		BetType:   betType,
		BetAmount: amount,
		Odds:      1.0,
		Potential: amount,
		Result:    "pending",
		PlacedAt:  time.Now(),
	}

	s.bets[bet.BetID] = bet
	player.Chips -= amount
	player.CurrentBet += amount
	gameState.Pot += amount

	return bet, nil
}

func (s *DealerService) EndBetting(roundID string) error {
	gameState, ok := s.gameStates[roundID]
	if !ok {
		return errors.New("round not found")
	}

	if gameState.Phase != "betting" {
		return errors.New("not in betting phase")
	}

	gameState.Phase = "playing"
	gameState.UpdatedAt = time.Now()
	return nil
}

func (s *DealerService) ResolveRound(roundID string, winner string) error {
	gameState, ok := s.gameStates[roundID]
	if !ok {
		return errors.New("round not found")
	}

	if gameState.Phase == "finished" {
		return errors.New("round already finished")
	}

	gameState.Winner = winner
	gameState.Phase = "resolving"
	gameState.EndTime = time.Now()
	gameState.UpdatedAt = time.Now()

	// Process all bets
	var totalPayout float64
	for _, bet := range s.bets {
		if bet.RoundID != roundID {
			continue
		}

		if winner == "player" {
			bet.Result = "won"
			bet.Payout = bet.BetAmount * 2
			totalPayout += bet.Payout

			// Update player chips
			if player, ok := s.players[bet.PlayerID]; ok {
				player.Chips += bet.Payout
			}
		} else if winner == "push" {
			bet.Result = "void"
			bet.Payout = bet.BetAmount

			// Return original bet
			if player, ok := s.players[bet.PlayerID]; ok {
				player.Chips += bet.Payout
			}
		} else {
			bet.Result = "lost"
		}

		bet.ResultedAt = time.Now()
	}

	gameState.Payout = totalPayout
	gameState.Phase = "finished"

	return nil
}

func (s *DealerService) GetRound(roundID string) (*model.GameState, error) {
	gameState, ok := s.gameStates[roundID]
	if !ok {
		return nil, errors.New("round not found")
	}
	return gameState, nil
}

func (s *DealerService) GetPlayerBets(playerID string) []*model.Bet {
	var result []*model.Bet
	for _, bet := range s.bets {
		if bet.PlayerID == playerID {
			result = append(result, bet)
		}
	}
	return result
}

// Dealer Management

func (s *DealerService) RegisterDealer(name, language string) (*model.Dealer, error) {
	dealer := &model.Dealer{
		DealerID:   generateID(),
		Name:       name,
		Language:   language,
		Status:     "available",
		ShiftStart: time.Now(),
	}

	s.dealers[dealer.DealerID] = dealer
	return dealer, nil
}

func (s *DealerService) AssignDealerToTable(dealerID, tableID string) error {
	dealer, ok := s.dealers[dealerID]
	if !ok {
		return errors.New("dealer not found")
	}

	table, ok := s.tables[tableID]
	if !ok {
		return errors.New("table not found")
	}

	dealer.Status = "busy"
	dealer.TableID = tableID

	table.DealerID = dealerID
	table.UpdatedAt = time.Now()

	return nil
}

func (s *DealerService) GetDealer(dealerID string) (*model.Dealer, error) {
	dealer, ok := s.dealers[dealerID]
	if !ok {
		return nil, errors.New("dealer not found")
	}
	return dealer, nil
}

func (s *DealerService) ListDealers(status string) []*model.Dealer {
	var result []*model.Dealer
	for _, d := range s.dealers {
		if status != "" && d.Status != status {
			continue
		}
		result = append(result, d)
	}
	return result
}

func generateID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[i%len(letters)]
	}
	return string(b)
}
