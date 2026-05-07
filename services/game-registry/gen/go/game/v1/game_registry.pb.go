// Package gamesv1 contains types for the GameRegistry gRPC service.
// This file is a manual stub. Replace with protoc-generated code when available.
package gamesv1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Enum types

type Status int32

const (
	Status_STATUS_UNSPECIFIED Status = 0
	Status_STATUS_ACTIVE      Status = 2
	Status_STATUS_INACTIVE    Status = 3
)

type GameCategoryEnum int32

const (
	GameCategoryEnum_GAME_CATEGORY_UNSPECIFIED GameCategoryEnum = 0
)

type DeviceType int32

const (
	DeviceType_DEVICE_TYPE_UNSPECIFIED DeviceType = 0
	DeviceType_DEVICE_TYPE_DESKTOP     DeviceType = 1
	DeviceType_DEVICE_TYPE_MOBILE      DeviceType = 2
)

type Language int32

const (
	Language_LANGUAGE_UNSPECIFIED Language = 0
	Language_LANGUAGE_EN          Language = 1
)

type GameProviderEnum int32

const (
	GameProviderEnum_GAME_PROVIDER_UNSPECIFIED GameProviderEnum = 0
)

// GameRegistryServiceServer is the server API for GameRegistryService.
type GameRegistryServiceServer interface {
	ListGames(context.Context, *ListGamesRequest) (*ListGamesResponse, error)
	GetGame(context.Context, *GetGameRequest) (*GetGameResponse, error)
	GetGameConfig(context.Context, *GetGameConfigRequest) (*GetGameConfigResponse, error)
	GetGameURL(context.Context, *GetGameURLRequest) (*GetGameURLResponse, error)
	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)
	GetProviders(context.Context, *GetProvidersRequest) (*GetProvidersResponse, error)
	SearchGames(context.Context, *SearchGamesRequest) (*SearchGamesResponse, error)
	GetFeaturedGames(context.Context, *GetFeaturedGamesRequest) (*GetFeaturedGamesResponse, error)
	GetPopularGames(context.Context, *GetPopularGamesRequest) (*GetPopularGamesResponse, error)
	GetNewGames(context.Context, *GetNewGamesRequest) (*GetNewGamesResponse, error)
}

// UnimplementedGameRegistryServiceServer can be embedded for forward compatibility.
type UnimplementedGameRegistryServiceServer struct{}

func (UnimplementedGameRegistryServiceServer) ListGames(context.Context, *ListGamesRequest) (*ListGamesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ListGames not implemented")
}
func (UnimplementedGameRegistryServiceServer) GetGame(context.Context, *GetGameRequest) (*GetGameResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedGameRegistryServiceServer) GetGameConfig(context.Context, *GetGameConfigRequest) (*GetGameConfigResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetGameConfig not implemented")
}
func (UnimplementedGameRegistryServiceServer) GetGameURL(context.Context, *GetGameURLRequest) (*GetGameURLResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetGameURL not implemented")
}
func (UnimplementedGameRegistryServiceServer) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedGameRegistryServiceServer) GetProviders(context.Context, *GetProvidersRequest) (*GetProvidersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetProviders not implemented")
}
func (UnimplementedGameRegistryServiceServer) SearchGames(context.Context, *SearchGamesRequest) (*SearchGamesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method SearchGames not implemented")
}
func (UnimplementedGameRegistryServiceServer) GetFeaturedGames(context.Context, *GetFeaturedGamesRequest) (*GetFeaturedGamesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetFeaturedGames not implemented")
}
func (UnimplementedGameRegistryServiceServer) GetPopularGames(context.Context, *GetPopularGamesRequest) (*GetPopularGamesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetPopularGames not implemented")
}
func (UnimplementedGameRegistryServiceServer) GetNewGames(context.Context, *GetNewGamesRequest) (*GetNewGamesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetNewGames not implemented")
}

// Request/Response types

type PaginationRequest struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`
}

func (m *PaginationRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}
func (m *PaginationRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type PaginationResponse struct {
	Page       int32 `json:"page"`
	PageSize   int32 `json:"page_size"`
	TotalCount int64 `json:"total_count"`
	TotalPages int32 `json:"total_pages"`
}

type ListGamesRequest struct {
	CategoryId       string              `json:"category_id"`
	ProviderId       string              `json:"provider_id"`
	Categories       []GameCategoryEnum  `json:"categories"`
	Providers        []GameProviderEnum  `json:"providers"`
	Status           Status              `json:"status"`
	MobileSupported  bool              `json:"mobile_supported"`
	DesktopSupported bool              `json:"desktop_supported"`
	IsFeatured       bool              `json:"is_featured"`
	IsJackpot        bool              `json:"is_jackpot"`
	Pagination       *PaginationRequest `json:"pagination"`
	SortBy           string            `json:"sort_by"`
	Query            string            `json:"query"`
}

func (m *ListGamesRequest) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}
func (m *ListGamesRequest) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}
func (m *ListGamesRequest) GetCategories() []GameCategoryEnum {
	if m != nil {
		return m.Categories
	}
	return nil
}
func (m *ListGamesRequest) GetProviders() []GameProviderEnum {
	if m != nil {
		return m.Providers
	}
	return nil
}
func (m *ListGamesRequest) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_UNSPECIFIED
}
func (m *ListGamesRequest) GetMobileSupported() bool {
	if m != nil {
		return m.MobileSupported
	}
	return false
}
func (m *ListGamesRequest) GetDesktopSupported() bool {
	if m != nil {
		return m.DesktopSupported
	}
	return false
}
func (m *ListGamesRequest) GetIsFeatured() bool {
	if m != nil {
		return m.IsFeatured
	}
	return false
}
func (m *ListGamesRequest) GetIsJackpot() bool {
	if m != nil {
		return m.IsJackpot
	}
	return false
}
func (m *ListGamesRequest) GetPagination() *PaginationRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}
func (m *ListGamesRequest) GetSortBy() string {
	if m != nil {
		return m.SortBy
	}
	return ""
}
func (m *ListGamesRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type ListGamesResponse struct {
	Games      []*GameSummary       `json:"games"`
	Pagination *PaginationResponse  `json:"pagination"`
}

type GetGameRequest struct {
	GameId string `json:"game_id"`
}

func (m *GetGameRequest) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type GetGameResponse struct {
	Game *Game `json:"game"`
}

type GetGameConfigRequest struct {
	GameId     string     `json:"game_id"`
	UserId     string     `json:"user_id"`
	DeviceType DeviceType `json:"device_type"`
	Language   Language   `json:"language"`
	Currency   string     `json:"currency"`
	SessionId  string     `json:"session_id"`
}

func (m *GetGameConfigRequest) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}
func (m *GetGameConfigRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}
func (m *GetGameConfigRequest) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_DEVICE_TYPE_UNSPECIFIED
}
func (m *GetGameConfigRequest) GetLanguage() Language {
	if m != nil {
		return m.Language
	}
	return Language_LANGUAGE_UNSPECIFIED
}
func (m *GetGameConfigRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}
func (m *GetGameConfigRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type GetGameConfigResponse struct {
	Config        *GameConfig `json:"config"`
	GameUrl       string      `json:"game_url"`
	SessionToken  string      `json:"session_token"`
}

type GetGameURLRequest struct {
	GameId     string     `json:"game_id"`
	UserId     string     `json:"user_id"`
	DeviceType DeviceType `json:"device_type"`
	SessionId  string     `json:"session_id"`
	Language   Language   `json:"language"`
	Currency   string     `json:"currency"`
}

func (m *GetGameURLRequest) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}
func (m *GetGameURLRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}
func (m *GetGameURLRequest) GetDeviceType() DeviceType {
	if m != nil {
		return m.DeviceType
	}
	return DeviceType_DEVICE_TYPE_UNSPECIFIED
}
func (m *GetGameURLRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}
func (m *GetGameURLRequest) GetLanguage() Language {
	if m != nil {
		return m.Language
	}
	return Language_LANGUAGE_UNSPECIFIED
}
func (m *GetGameURLRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type GetGameURLResponse struct {
	GameUrl      string       `json:"game_url"`
	SessionToken string       `json:"session_token"`
	Game         *GameSummary `json:"game"`
}

type GetCategoriesRequest struct {
	IncludeGamesCount bool `json:"include_games_count"`
}

func (m *GetCategoriesRequest) GetIncludeGamesCount() bool {
	if m != nil {
		return m.IncludeGamesCount
	}
	return false
}

type GetCategoriesResponse struct {
	Categories []*GameCategory `json:"categories"`
}

type GetProvidersRequest struct {
	ActiveOnly bool `json:"active_only"`
}

func (m *GetProvidersRequest) GetActiveOnly() bool {
	if m != nil {
		return m.ActiveOnly
	}
	return false
}

type GetProvidersResponse struct {
	Providers []*GameProvider `json:"providers"`
}

type SearchGamesRequest struct {
	Query      string `json:"query"`
	Limit      int32  `json:"limit"`
	CategoryId string `json:"category_id"`
}

func (m *SearchGamesRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}
func (m *SearchGamesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}
func (m *SearchGamesRequest) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

type SearchGamesResponse struct {
	Games      []*GameSummary `json:"games"`
	TotalCount int32          `json:"total_count"`
}

type GetFeaturedGamesRequest struct {
	Limit      int32  `json:"limit"`
	CategoryId string `json:"category_id"`
}

func (m *GetFeaturedGamesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}
func (m *GetFeaturedGamesRequest) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

type GetFeaturedGamesResponse struct {
	Games []*GameSummary `json:"games"`
}

type GetPopularGamesRequest struct {
	Limit      int32  `json:"limit"`
	CategoryId string `json:"category_id"`
}

func (m *GetPopularGamesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}
func (m *GetPopularGamesRequest) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

type GetPopularGamesResponse struct {
	Games []*GameSummary `json:"games"`
}

type GetNewGamesRequest struct {
	Limit      int32  `json:"limit"`
	CategoryId string `json:"category_id"`
}

func (m *GetNewGamesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}
func (m *GetNewGamesRequest) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

type GetNewGamesResponse struct {
	Games []*GameSummary `json:"games"`
}

// Data model types

type GameSummary struct {
	GameId           string          `json:"game_id"`
	Name             string          `json:"name"`
	ProviderId       string          `json:"provider_id"`
	ProviderName     string          `json:"provider_name"`
	CategoryId       string          `json:"category_id"`
	CategoryName     string          `json:"category_name"`
	Type             GameCategoryEnum `json:"type"`
	Status           Status          `json:"status"`
	ThumbnailUrl     string       `json:"thumbnail_url"`
	BannerUrl        string       `json:"banner_url"`
	Rtp              float64      `json:"rtp"`
	Volatility       string       `json:"volatility"`
	MaxWin           string       `json:"max_win"`
	IsFeatured       bool         `json:"is_featured"`
	IsNew            bool         `json:"is_new"`
	IsPopular        bool         `json:"is_popular"`
	IsJackpot        bool         `json:"is_jackpot"`
	LaunchUrl        string       `json:"launch_url"`
	PopularityScore  int32        `json:"popularity_score"`
}

type Game struct {
	GameId       string          `json:"game_id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	ProviderId   string          `json:"provider_id"`
	ProviderName string          `json:"provider_name"`
	CategoryId   string          `json:"category_id"`
	CategoryName string          `json:"category_name"`
	Type         GameCategoryEnum `json:"type"`
	Status       Status          `json:"status"`
	ThumbnailUrl string       `json:"thumbnail_url"`
	BannerUrl    string       `json:"banner_url"`
	Rtp          float64      `json:"rtp"`
	Volatility   string       `json:"volatility"`
	MaxWin       string       `json:"max_win"`
	IsFeatured   bool         `json:"is_featured"`
	IsNew        bool         `json:"is_new"`
	IsPopular    bool         `json:"is_popular"`
	IsJackpot    bool         `json:"is_jackpot"`
	LaunchUrl    string       `json:"launch_url"`
}

type GameConfig struct {
	GameId       string `json:"game_id"`
	SessionToken string `json:"session_token"`
	GameUrl      string `json:"game_url"`
	PlayerId     string `json:"player_id"`
	Currency     string `json:"currency"`
	Language     string `json:"language"`
}

type GameProvider struct {
	ProviderId  string `json:"provider_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LogoUrl     string `json:"logo_url"`
	WebsiteUrl  string `json:"website_url"`
	Status      Status `json:"status"`
	GamesCount  int32  `json:"games_count"`
	License     string `json:"license"`
	Established int32  `json:"established"`
	IsFeatured  bool   `json:"is_featured"`
}

type GameCategory struct {
	CategoryId  string `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IconUrl     string `json:"icon_url"`
	BannerUrl   string `json:"banner_url"`
	ParentId    string `json:"parent_id"`
	SortOrder   int32  `json:"sort_order"`
	Status      Status `json:"status"`
	IsFeatured  bool   `json:"is_featured"`
	GamesCount  int32  `json:"games_count"`
	Slug        string `json:"slug"`
}

// Service descriptor and registration

var GameRegistryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game_engine.game.v1.GameRegistryService",
	HandlerType: (*GameRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "ListGames"},
		{MethodName: "GetGame"},
		{MethodName: "GetGameConfig"},
		{MethodName: "GetGameURL"},
		{MethodName: "GetCategories"},
		{MethodName: "GetProviders"},
		{MethodName: "SearchGames"},
		{MethodName: "GetFeaturedGames"},
		{MethodName: "GetPopularGames"},
		{MethodName: "GetNewGames"},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game_engine/game/v1/game_registry.proto",
}

func RegisterGameRegistryServiceServer(s *grpc.Server, srv GameRegistryServiceServer) {
	s.RegisterService(&GameRegistryService_ServiceDesc, srv)
}
