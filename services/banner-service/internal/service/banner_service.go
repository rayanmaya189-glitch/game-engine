package service

import (
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/game_engine/banner-service/internal/model"
	"github.com/game_engine/banner-service/internal/repository"
)

type BannerService struct {
	repo *repository.BannerRepository
}

func NewBannerService(repo *repository.BannerRepository) *BannerService {
	return &BannerService{repo: repo}
}

func (s *BannerService) CreateBanner(req *model.Banner) (*model.Banner, error) {
	if req.Title == "" {
		return nil, errors.New("title is required")
	}
	if req.ImageURL == "" {
		return nil, errors.New("image_url is required")
	}

	now := time.Now()
	req.ID = uuid.New().String()
	req.Status = model.BannerStatusActive
	req.CreatedAt = now
	req.UpdatedAt = now

	if req.Width == 0 {
		req.Width = 1920
	}
	if req.Height == 0 {
		req.Height = 600
	}

	if err := s.repo.Create(req); err != nil {
		return nil, err
	}

	return req, nil
}

func (s *BannerService) GetBanner(id string) (*model.Banner, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	return s.repo.GetByID(id)
}

func (s *BannerService) ListBanners(filter *model.BannerFilter) (*model.BannerList, error) {
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PageSize < 1 {
		filter.PageSize = 20
	}
	if filter.PageSize > 100 {
		filter.PageSize = 100
	}
	return s.repo.List(filter)
}

func (s *BannerService) UpdateBanner(id string, req *model.Banner) (*model.Banner, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		existing.Title = req.Title
	}
	if req.Description != "" {
		existing.Description = req.Description
	}
	if req.ImageURL != "" {
		existing.ImageURL = req.ImageURL
	}
	if req.ClickURL != "" {
		existing.ClickURL = req.ClickURL
	}
	if req.BannerType != "" {
		existing.BannerType = req.BannerType
	}
	if req.Status != "" {
		existing.Status = req.Status
	}
	if req.Priority != 0 {
		existing.Priority = req.Priority
	}
	if req.Width != 0 {
		existing.Width = req.Width
	}
	if req.Height != 0 {
		existing.Height = req.Height
	}
	if req.StartDate != nil {
		existing.StartDate = req.StartDate
	}
	if req.EndDate != nil {
		existing.EndDate = req.EndDate
	}
	if len(req.TargetCountries) > 0 {
		existing.TargetCountries = req.TargetCountries
	}
	if len(req.TargetVIPLevels) > 0 {
		existing.TargetVIPLevels = req.TargetVIPLevels
	}
	if len(req.TargetGameTypes) > 0 {
		existing.TargetGameTypes = req.TargetGameTypes
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *BannerService) DeleteBanner(id string) error {
	if id == "" {
		return errors.New("id is required")
	}
	return s.repo.Delete(id)
}

func (s *BannerService) RecordImpression(id string) error {
	return s.repo.IncrementImpression(id)
}

func (s *BannerService) RecordClick(id string) error {
	return s.repo.IncrementClick(id)
}

func (s *BannerService) GetActiveBanners(bannerType, country, vipLevel string) ([]*model.Banner, error) {
	filter := &model.BannerFilter{
		BannerType: bannerType,
		Status:     string(model.BannerStatusActive),
		Country:    country,
		VIPLevel:   vipLevel,
		Page:       1,
		PageSize:   50,
	}

	result, err := s.repo.List(filter)
	if err != nil {
		return nil, err
	}

	var active []*model.Banner
	now := time.Now()
	for _, b := range result.Banners {
		if b.StartDate != nil && b.StartDate.After(now) {
			continue
		}
		if b.EndDate != nil && b.EndDate.Before(now) {
			continue
		}
		active = append(active, b)
	}

	return active, nil
}

func (s *BannerService) CreateAnnouncement(req *model.Announcement) (*model.Announcement, error) {
	if req.Title == "" {
		return nil, errors.New("title is required")
	}
	if req.Content == "" {
		return nil, errors.New("content is required")
	}

	now := time.Now()
	req.ID = uuid.New().String()
	req.Status = model.BannerStatusActive
	req.CreatedAt = now
	req.UpdatedAt = now

	if err := s.repo.CreateAnnouncement(req); err != nil {
		return nil, err
	}

	return req, nil
}

func (s *BannerService) ListActiveAnnouncements() ([]*model.Announcement, error) {
	return s.repo.ListActiveAnnouncements()
}
