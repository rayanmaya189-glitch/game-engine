package handler

import (
	"context"

	gamesv1 "github.com/game_engine/game-registry/gen/go/game/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetCategories handles the GetCategories gRPC call
func (h *GameHandler) GetCategories(ctx context.Context, req *gamesv1.GetCategoriesRequest) (*gamesv1.GetCategoriesResponse, error) {
	categories, err := h.gameService.GetCategories(ctx, req.GetIncludeGamesCount())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get categories: %v", err)
	}

	cats := make([]*gamesv1.GameCategory, len(categories))
	for i, c := range categories {
		cats[i] = &gamesv1.GameCategory{
			CategoryId:  c.ID,
			Name:        c.Name,
			Description: c.Description,
			IconUrl:     c.IconURL,
			BannerUrl:   c.BannerURL,
			ParentId:    c.ParentID,
			SortOrder:   int32(c.SortOrder),
			Status:      gamesv1.Status(c.Status),
			IsFeatured:  c.IsFeatured,
			GamesCount:  int32(c.GamesCount),
			Slug:        c.Slug,
		}
	}

	return &gamesv1.GetCategoriesResponse{
		Categories: cats,
	}, nil
}

// GetProviders handles the GetProviders gRPC call
func (h *GameHandler) GetProviders(ctx context.Context, req *gamesv1.GetProvidersRequest) (*gamesv1.GetProvidersResponse, error) {
	providers, err := h.gameService.GetProviders(ctx, req.GetActiveOnly())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get providers: %v", err)
	}

	provs := make([]*gamesv1.GameProvider, len(providers))
	for i, p := range providers {
		provs[i] = &gamesv1.GameProvider{
			ProviderId:  p.ID,
			Name:        p.Name,
			Description: p.Description,
			LogoUrl:     p.LogoURL,
			WebsiteUrl:  p.WebsiteURL,
			Status:      gamesv1.Status(p.Status),
			GamesCount:  int32(p.GamesCount),
			License:     p.License,
			Established: int32(p.Established),
			IsFeatured:  p.IsFeatured,
		}
	}

	return &gamesv1.GetProvidersResponse{
		Providers: provs,
	}, nil
}
