package service

import (
	"errors"
	"time"

	"github.com/game_engine/live-dealer-service/internal/model"
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
