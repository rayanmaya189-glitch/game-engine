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
		TableID:            tableID,
		Bitrate:            s.cfg.Dealer.VideoStreamBitrate,
		ViewerCount:        len(onlinePlayers),
		StreamURL:          fmt.Sprintf("rtmp://stream.internal/%s/live", tableID),
		IsLive:             table.Status == "open",
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
