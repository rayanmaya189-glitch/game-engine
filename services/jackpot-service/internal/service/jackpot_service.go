package service

import (
	"context"

	jackpotpb "github.com/game_engine/gen/go/game_engine/jackpot/v1"
	"github.com/game_engine/jackpot-service/internal/config"
	"github.com/game_engine/jackpot-service/internal/repository"
)

type JackpotService struct {
	repo *repository.JackpotRepository
	cfg  *config.Config
}

func NewJackpotService(repo *repository.JackpotRepository, cfg *config.Config) *JackpotService {
	return &JackpotService{repo: repo, cfg: cfg}
}

func (s *JackpotService) ListJackpots(ctx context.Context, req *jackpotpb.ListJackpotsRequest) (*jackpotpb.ListJackpotsResponse, error) {
	jackpots, err := s.repo.ListJackpots(ctx, req.Status)
	if err != nil {
		return nil, err
	}

	pbJackpots := make([]*jackpotpb.Jackpot, len(jackpots))
	for i, j := range jackpots {
		pbJackpots[i] = &jackpotpb.Jackpot{
			JackpotId:     j.JackpotID,
			Name:          j.Name,
			Description:   j.Description,
			CurrentAmount: j.CurrentAmount,
			MinBet:        j.MinBet,
			MaxBet:        j.MaxBet,
			Status:        j.Status,
			StartsAt:      j.StartsAt.Unix(),
			EndsAt:        j.EndsAt.Unix(),
		}
	}
	return &jackpotpb.ListJackpotsResponse{Jackpots: pbJackpots}, nil
}

func (s *JackpotService) GetJackpot(ctx context.Context, req *jackpotpb.GetJackpotRequest) (*jackpotpb.GetJackpotResponse, error) {
	j, err := s.repo.GetJackpot(ctx, req.JackpotId)
	if err != nil {
		return nil, err
	}
	return &jackpotpb.GetJackpotResponse{
		Jackpot: &jackpotpb.Jackpot{
			JackpotId:     j.JackpotID,
			Name:          j.Name,
			Description:   j.Description,
			CurrentAmount: j.CurrentAmount,
			MinBet:        j.MinBet,
			MaxBet:        j.MaxBet,
			Status:        j.Status,
			StartsAt:      j.StartsAt.Unix(),
			EndsAt:        j.EndsAt.Unix(),
		},
	}, nil
}

func (s *JackpotService) GetWinners(ctx context.Context, req *jackpotpb.GetWinnersRequest) (*jackpotpb.GetWinnersResponse, error) {
	winners, err := s.repo.GetWinners(ctx, req.JackpotId, req.Limit)
	if err != nil {
		return nil, err
	}
	pbWinners := make([]*jackpotpb.Winner, len(winners))
	for i, w := range winners {
		pbWinners[i] = &jackpotpb.Winner{
			WinnerId: w.WinnerID,
			Username: w.Username,
			Amount:   w.Amount,
			WonAt:    w.WonAt.Unix(),
		}
	}
	return &jackpotpb.GetWinnersResponse{Winners: pbWinners}, nil
}

func (s *JackpotService) JoinJackpot(ctx context.Context, req *jackpotpb.JoinJackpotRequest) (*jackpotpb.JoinJackpotResponse, error) {
	success, message, err := s.repo.JoinJackpot(ctx, req.JackpotId, req.UserId, req.BetAmount)
	if err != nil {
		return nil, err
	}
	return &jackpotpb.JoinJackpotResponse{Success: success, Message: message}, nil
}

func (s *JackpotService) GetJackpotHistory(ctx context.Context, req *jackpotpb.GetJackpotHistoryRequest) (*jackpotpb.GetJackpotHistoryResponse, error) {
	entries, total, err := s.repo.GetJackpotHistory(ctx, req.UserId, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	pbEntries := make([]*jackpotpb.JackpotHistoryEntry, len(entries))
	for i, e := range entries {
		pbEntries[i] = &jackpotpb.JackpotHistoryEntry{
			JackpotId:   e.JackpotID,
			JackpotName: e.JackpotName,
			Amount:      e.Amount,
			Result:      e.Result,
			PlayedAt:    e.PlayedAt.Unix(),
		}
	}
	return &jackpotpb.GetJackpotHistoryResponse{Entries: pbEntries, Total: int32(total)}, nil
}
