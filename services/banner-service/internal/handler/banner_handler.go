package handler

import (
	"context"
	"encoding/json"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/game_engine/banner-service/internal/model"
	"github.com/game_engine/banner-service/internal/service"
)

type BannerHandler struct {
	svc *service.BannerService
}

func NewBannerHandler(svc *service.BannerService) *BannerHandler {
	return &BannerHandler{svc: svc}
}

func (h *BannerHandler) RegisterServices(server *grpc.Server) {}

func (h *BannerHandler) CreateBanner(ctx context.Context, req []byte) (*model.Banner, error) {
	var banner model.Banner
	if err := json.Unmarshal(req, &banner); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}
	return h.svc.CreateBanner(&banner)
}

func (h *BannerHandler) GetBanner(ctx context.Context, id string) (*model.Banner, error) {
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	banner, err := h.svc.GetBanner(id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "banner not found: %v", err)
	}
	return banner, nil
}

func (h *BannerHandler) ListBanners(ctx context.Context, bannerType, statusFilter string, page, pageSize int) (*model.BannerList, error) {
	filter := &model.BannerFilter{
		BannerType: bannerType,
		Status:     statusFilter,
		Page:       page,
		PageSize:   pageSize,
	}
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PageSize < 1 {
		filter.PageSize = 20
	}
	return h.svc.ListBanners(filter)
}

func (h *BannerHandler) UpdateBanner(ctx context.Context, id string, req []byte) (*model.Banner, error) {
	var banner model.Banner
	if err := json.Unmarshal(req, &banner); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}
	return h.svc.UpdateBanner(id, &banner)
}

func (h *BannerHandler) DeleteBanner(ctx context.Context, id string) error {
	if id == "" {
		return status.Error(codes.InvalidArgument, "id is required")
	}
	return h.svc.DeleteBanner(id)
}

func (h *BannerHandler) GetActiveBanners(ctx context.Context, bannerType, country, vipLevel string) ([]*model.Banner, error) {
	return h.svc.GetActiveBanners(bannerType, country, vipLevel)
}

func (h *BannerHandler) RecordImpression(ctx context.Context, id string) error {
	return h.svc.RecordImpression(id)
}

func (h *BannerHandler) RecordClick(ctx context.Context, id string) error {
	return h.svc.RecordClick(id)
}

func (h *BannerHandler) CreateAnnouncement(ctx context.Context, req []byte) (*model.Announcement, error) {
	var announcement model.Announcement
	if err := json.Unmarshal(req, &announcement); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %v", err)
	}
	return h.svc.CreateAnnouncement(&announcement)
}

func (h *BannerHandler) ListActiveAnnouncements(ctx context.Context) ([]*model.Announcement, error) {
	return h.svc.ListActiveAnnouncements()
}

func parsePagination(pageStr, sizeStr string) (int, int) {
	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}
	size, _ := strconv.Atoi(sizeStr)
	if size < 1 || size > 100 {
		size = 20
	}
	return page, size
}
