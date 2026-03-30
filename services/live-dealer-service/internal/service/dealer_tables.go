package service

import (
	"errors"
	"time"

	"github.com/game_engine/live-dealer-service/internal/model"
)

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
