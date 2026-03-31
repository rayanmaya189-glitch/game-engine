package handler

import (
	"context"

	gamesv1 "github.com/game_engine/game-registry/gen/go/game/v1"

	"github.com/game_engine/game-registry/internal/enums"
	"github.com/game_engine/game-registry/internal/model"
	"github.com/game_engine/game-registry/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GameHandler handles gRPC requests for games
type GameHandler struct {
	gamesv1.UnimplementedGameRegistryServiceServer
	gameService *service.GameService
}

// NewGameHandler creates a new GameHandler
func NewGameHandler(gameService *service.GameService) *GameHandler {
	return &GameHandler{
		gameService: gameService,
	}
}

// ListGames handles the ListGames gRPC call
func (h *GameHandler) ListGames(ctx context.Context, req *gamesv1.ListGamesRequest) (*gamesv1.ListGamesResponse, error) {
	svcReq := &service.ListGamesRequest{
		CategoryID:       req.GetCategoryId(),
		ProviderID:       req.GetProviderId(),
		MobileSupported:  req.GetMobileSupported(),
		DesktopSupported: req.GetDesktopSupported(),
		IsFeatured:       req.GetIsFeatured(),
		IsJackpot:        req.GetIsJackpot(),
		Query:            req.GetQuery(),
		SortBy:           req.GetSortBy(),
		Status:           int32(req.GetStatus()),
	}

	if req.Pagination != nil {
		svcReq.Pagination.Page = req.Pagination.GetPage()
		svcReq.Pagination.PageSize = req.Pagination.GetPageSize()
	}

	for _, cat := range req.GetCategories() {
		svcReq.Categories = append(svcReq.Categories, int32(cat))
	}
	for _, prov := range req.GetProviders() {
		svcReq.Providers = append(svcReq.Providers, int32(prov))
	}

	resp, err := h.gameService.ListGames(ctx, svcReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list games: %v", err)
	}

	games := make([]*gamesv1.GameSummary, len(resp.Games))
	for i, g := range resp.Games {
		games[i] = gameSummaryToProto(&g)
	}

	var pagination *gamesv1.PaginationResponse
	if resp.Pagination != nil {
		pagination = &gamesv1.PaginationResponse{
			Page:       resp.Pagination.Page,
			PageSize:   resp.Pagination.PageSize,
			TotalCount: resp.Pagination.TotalCount,
			TotalPages: resp.Pagination.TotalPages,
		}
	}

	return &gamesv1.ListGamesResponse{
		Games:      games,
		Pagination: pagination,
	}, nil
}

// GetGame handles the GetGame gRPC call
func (h *GameHandler) GetGame(ctx context.Context, req *gamesv1.GetGameRequest) (*gamesv1.GetGameResponse, error) {
	gameID := req.GetGameId()
	if gameID == "" {
		return nil, status.Error(codes.InvalidArgument, "game_id is required")
	}

	game, err := h.gameService.GetGame(ctx, gameID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "game not found: %v", err)
	}

	return &gamesv1.GetGameResponse{
		Game: gameModelToProto(game),
	}, nil
}

// GetGameConfig handles the GetGameConfig gRPC call
func (h *GameHandler) GetGameConfig(ctx context.Context, req *gamesv1.GetGameConfigRequest) (*gamesv1.GetGameConfigResponse, error) {
	gameID := req.GetGameId()
	if gameID == "" {
		return nil, status.Error(codes.InvalidArgument, "game_id is required")
	}

	config, err := h.gameService.GetGameConfig(
		ctx,
		gameID,
		req.GetUserId(),
		enums.DeviceType(req.GetDeviceType()),
		enums.GameLanguage(req.GetLanguage()),
		req.GetCurrency(),
		req.GetSessionId(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get game config: %v", err)
	}

	return &gamesv1.GetGameConfigResponse{
		Config: &gamesv1.GameConfig{
			GameId:       config.GameID,
			SessionToken: config.SessionToken,
			GameUrl:      config.GameURL,
			PlayerId:     config.PlayerID,
			Currency:     config.Currency,
			Language:     config.Language,
		},
		GameUrl:      config.GameURL,
		SessionToken: config.SessionToken,
	}, nil
}

// GetGameURL handles the GetGameURL gRPC call
func (h *GameHandler) GetGameURL(ctx context.Context, req *gamesv1.GetGameURLRequest) (*gamesv1.GetGameURLResponse, error) {
	gameID := req.GetGameId()
	if gameID == "" {
		return nil, status.Error(codes.InvalidArgument, "game_id is required")
	}

	result, err := h.gameService.GetGameURL(
		ctx,
		gameID,
		req.GetUserId(),
		enums.DeviceType(req.GetDeviceType()),
		req.GetSessionId(),
		enums.GameLanguage(req.GetLanguage()),
		req.GetCurrency(),
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get game URL: %v", err)
	}

	return &gamesv1.GetGameURLResponse{
		GameUrl:      result.GameURL,
		SessionToken: result.SessionToken,
		Game:         gameSummaryToProto(&result.Game),
	}, nil
}

// SearchGames handles the SearchGames gRPC call
func (h *GameHandler) SearchGames(ctx context.Context, req *gamesv1.SearchGamesRequest) (*gamesv1.SearchGamesResponse, error) {
	games, err := h.gameService.SearchGames(ctx, req.GetQuery(), int(req.GetLimit()), req.GetCategoryId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to search games: %v", err)
	}

	result := make([]*gamesv1.GameSummary, len(games))
	for i, g := range games {
		result[i] = gameSummaryToProto(&g)
	}

	return &gamesv1.SearchGamesResponse{
		Games:      result,
		TotalCount: int32(len(games)),
	}, nil
}

// GetFeaturedGames handles the GetFeaturedGames gRPC call
func (h *GameHandler) GetFeaturedGames(ctx context.Context, req *gamesv1.GetFeaturedGamesRequest) (*gamesv1.GetFeaturedGamesResponse, error) {
	games, err := h.gameService.GetFeaturedGames(ctx, int(req.GetLimit()), req.GetCategoryId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get featured games: %v", err)
	}

	result := make([]*gamesv1.GameSummary, len(games))
	for i, g := range games {
		result[i] = gameSummaryToProto(&g)
	}

	return &gamesv1.GetFeaturedGamesResponse{
		Games: result,
	}, nil
}

// GetPopularGames handles the GetPopularGames gRPC call
func (h *GameHandler) GetPopularGames(ctx context.Context, req *gamesv1.GetPopularGamesRequest) (*gamesv1.GetPopularGamesResponse, error) {
	games, err := h.gameService.GetPopularGames(ctx, int(req.GetLimit()), req.GetCategoryId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get popular games: %v", err)
	}

	result := make([]*gamesv1.GameSummary, len(games))
	for i, g := range games {
		result[i] = gameSummaryToProto(&g)
	}

	return &gamesv1.GetPopularGamesResponse{
		Games: result,
	}, nil
}

// GetNewGames handles the GetNewGames gRPC call
func (h *GameHandler) GetNewGames(ctx context.Context, req *gamesv1.GetNewGamesRequest) (*gamesv1.GetNewGamesResponse, error) {
	games, err := h.gameService.GetNewGames(ctx, int(req.GetLimit()), req.GetCategoryId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get new games: %v", err)
	}

	result := make([]*gamesv1.GameSummary, len(games))
	for i, g := range games {
		result[i] = gameSummaryToProto(&g)
	}

	return &gamesv1.GetNewGamesResponse{
		Games: result,
	}, nil
}

// Helper to convert model.GameSummary to proto GameSummary
func gameSummaryToProto(g *model.GameSummary) *gamesv1.GameSummary {
	return &gamesv1.GameSummary{
		GameId:          g.GameID,
		Name:            g.Name,
		ProviderId:      g.ProviderID,
		ProviderName:    g.ProviderName,
		CategoryId:      g.CategoryID,
		CategoryName:    g.CategoryName,
		Type:            gamesv1.GameCategoryEnum(g.Type),
		Status:          gamesv1.Status(g.Status),
		ThumbnailUrl:    g.ThumbnailURL,
		BannerUrl:       g.BannerURL,
		Rtp:             g.RTP,
		Volatility:      g.Volatility,
		MaxWin:          g.MaxWin,
		IsFeatured:      g.IsFeatured,
		IsNew:           g.IsNew,
		IsPopular:       g.IsPopular,
		IsJackpot:       g.IsJackpot,
		LaunchUrl:       g.LaunchURL,
		PopularityScore: int32(g.PopularityScore),
	}
}

// Helper to convert model.Game to proto Game
func gameModelToProto(g *model.Game) *gamesv1.Game {
	return &gamesv1.Game{
		GameId:       g.ID,
		Name:         g.Name,
		Description:  g.Description,
		ProviderId:   g.ProviderID,
		ProviderName: g.ProviderName,
		CategoryId:   g.CategoryID,
		CategoryName: g.CategoryName,
		Type:         gamesv1.GameCategoryEnum(g.Type),
		Status:       gamesv1.Status(g.Status),
		ThumbnailUrl: g.ThumbnailURL,
		BannerUrl:    g.BannerURL,
		Rtp:          g.RTP,
		Volatility:   g.Volatility,
		MaxWin:       g.MaxWin,
		IsFeatured:   g.IsFeatured,
		IsNew:        g.IsNew,
		IsPopular:    g.IsPopular,
		IsJackpot:    g.IsJackpot,
		LaunchUrl:    g.LaunchURL,
	}
}
