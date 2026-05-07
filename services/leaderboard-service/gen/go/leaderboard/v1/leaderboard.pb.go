// Package leaderv1 contains types for the Leaderboard gRPC service.
// This file is a manual stub. Replace with protoc-generated code when available.
package leaderv1

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LeaderboardServiceServer is the server API for LeaderboardService.
type LeaderboardServiceServer interface {
	GetDailyLeaderboard(context.Context, *GetLeaderboardRequest) (*GetLeaderboardResponse, error)
	GetWeeklyLeaderboard(context.Context, *GetLeaderboardRequest) (*GetLeaderboardResponse, error)
	GetMonthlyLeaderboard(context.Context, *GetLeaderboardRequest) (*GetLeaderboardResponse, error)
	GetAllTimeLeaderboard(context.Context, *GetLeaderboardRequest) (*GetLeaderboardResponse, error)
	GetPlayerRank(context.Context, *GetPlayerRankRequest) (*GetPlayerRankResponse, error)
	UpdatePlayerScore(context.Context, *UpdatePlayerScoreRequest) (*UpdatePlayerScoreResponse, error)
	DistributePrizes(context.Context, *DistributePrizesRequest) (*DistributePrizesResponse, error)
	SyncLeaderboard(context.Context, *SyncLeaderboardRequest) (*SyncLeaderboardResponse, error)
	ResetLeaderboard(context.Context, *ResetLeaderboardRequest) (*ResetLeaderboardResponse, error)
}

// UnimplementedLeaderboardServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLeaderboardServiceServer struct{}

func (UnimplementedLeaderboardServiceServer) GetDailyLeaderboard(context.Context, *GetLeaderboardRequest) (*GetLeaderboardResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetDailyLeaderboard not implemented")
}
func (UnimplementedLeaderboardServiceServer) GetWeeklyLeaderboard(context.Context, *GetLeaderboardRequest) (*GetLeaderboardResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetWeeklyLeaderboard not implemented")
}
func (UnimplementedLeaderboardServiceServer) GetMonthlyLeaderboard(context.Context, *GetLeaderboardRequest) (*GetLeaderboardResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetMonthlyLeaderboard not implemented")
}
func (UnimplementedLeaderboardServiceServer) GetAllTimeLeaderboard(context.Context, *GetLeaderboardRequest) (*GetLeaderboardResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetAllTimeLeaderboard not implemented")
}
func (UnimplementedLeaderboardServiceServer) GetPlayerRank(context.Context, *GetPlayerRankRequest) (*GetPlayerRankResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetPlayerRank not implemented")
}
func (UnimplementedLeaderboardServiceServer) UpdatePlayerScore(context.Context, *UpdatePlayerScoreRequest) (*UpdatePlayerScoreResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method UpdatePlayerScore not implemented")
}
func (UnimplementedLeaderboardServiceServer) DistributePrizes(context.Context, *DistributePrizesRequest) (*DistributePrizesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method DistributePrizes not implemented")
}
func (UnimplementedLeaderboardServiceServer) SyncLeaderboard(context.Context, *SyncLeaderboardRequest) (*SyncLeaderboardResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method SyncLeaderboard not implemented")
}
func (UnimplementedLeaderboardServiceServer) ResetLeaderboard(context.Context, *ResetLeaderboardRequest) (*ResetLeaderboardResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ResetLeaderboard not implemented")
}

// Message types

type GetLeaderboardRequest struct {
	GameType string `json:"game_type"`
	Limit    int32  `json:"limit"`
}

func (m *GetLeaderboardRequest) GetGameType() string {
	if m != nil {
		return m.GameType
	}
	return ""
}
func (m *GetLeaderboardRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetLeaderboardResponse struct {
	Type      string             `json:"type"`
	Period    string             `json:"period"`
	Entries   []*LeaderboardEntry `json:"entries"`
	Total     int32              `json:"total"`
	UpdatedAt *time.Time         `json:"updated_at"`
}

type LeaderboardEntry struct {
	Rank        int32      `json:"rank"`
	UserId      string     `json:"user_id"`
	Username    string     `json:"username"`
	Score       float64    `json:"score"`
	Wins        int32      `json:"wins"`
	WinAmount   float64    `json:"win_amount"`
	GameType    string     `json:"game_type"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type GetPlayerRankRequest struct {
	UserId          string `json:"user_id"`
	GameType        string `json:"game_type"`
	LeaderboardType string `json:"leaderboard_type"`
}

func (m *GetPlayerRankRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}
func (m *GetPlayerRankRequest) GetGameType() string {
	if m != nil {
		return m.GameType
	}
	return ""
}
func (m *GetPlayerRankRequest) GetLeaderboardType() string {
	if m != nil {
		return m.LeaderboardType
	}
	return ""
}

type GetPlayerRankResponse struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Rank     int32  `json:"rank"`
	Score    float64 `json:"score"`
	Type     string `json:"type"`
	Period   string `json:"period"`
}

type UpdatePlayerScoreRequest struct {
	UserId    string  `json:"user_id"`
	Username  string  `json:"username"`
	Score     float64 `json:"score"`
	GameType  string  `json:"game_type"`
	IsWin     bool    `json:"is_win"`
	WinAmount float64 `json:"win_amount"`
	BetAmount float64 `json:"bet_amount"`
}

func (m *UpdatePlayerScoreRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}
func (m *UpdatePlayerScoreRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}
func (m *UpdatePlayerScoreRequest) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}
func (m *UpdatePlayerScoreRequest) GetGameType() string {
	if m != nil {
		return m.GameType
	}
	return ""
}
func (m *UpdatePlayerScoreRequest) GetIsWin() bool {
	if m != nil {
		return m.IsWin
	}
	return false
}
func (m *UpdatePlayerScoreRequest) GetWinAmount() float64 {
	if m != nil {
		return m.WinAmount
	}
	return 0
}
func (m *UpdatePlayerScoreRequest) GetBetAmount() float64 {
	if m != nil {
		return m.BetAmount
	}
	return 0
}

type UpdatePlayerScoreResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type DistributePrizesRequest struct {
	LeaderboardType string `json:"leaderboard_type"`
	GameType        string `json:"game_type"`
	TournamentId    string `json:"tournament_id"`
	DryRun          bool   `json:"dry_run"`
}

func (m *DistributePrizesRequest) GetLeaderboardType() string {
	if m != nil {
		return m.LeaderboardType
	}
	return ""
}
func (m *DistributePrizesRequest) GetGameType() string {
	if m != nil {
		return m.GameType
	}
	return ""
}
func (m *DistributePrizesRequest) GetTournamentId() string {
	if m != nil {
		return m.TournamentId
	}
	return ""
}
func (m *DistributePrizesRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

type DistributePrizesResponse struct {
	LeaderboardType string   `json:"leaderboard_type"`
	GameType        string   `json:"game_type"`
	Period          string   `json:"period"`
	Prizes          []*Prize `json:"prizes"`
	TotalValue      float64  `json:"total_value"`
	DistributedAt   *time.Time `json:"distributed_at"`
}

type Prize struct {
	Rank     int32   `json:"rank"`
	Type     string  `json:"type"`
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

type SyncLeaderboardRequest struct {
	LeaderboardType string `json:"leaderboard_type"`
	GameType        string `json:"game_type"`
}

func (m *SyncLeaderboardRequest) GetLeaderboardType() string {
	if m != nil {
		return m.LeaderboardType
	}
	return ""
}
func (m *SyncLeaderboardRequest) GetGameType() string {
	if m != nil {
		return m.GameType
	}
	return ""
}

type SyncLeaderboardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResetLeaderboardRequest struct {
	LeaderboardType string `json:"leaderboard_type"`
	GameType        string `json:"game_type"`
}

func (m *ResetLeaderboardRequest) GetLeaderboardType() string {
	if m != nil {
		return m.LeaderboardType
	}
	return ""
}
func (m *ResetLeaderboardRequest) GetGameType() string {
	if m != nil {
		return m.GameType
	}
	return ""
}

type ResetLeaderboardResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Service descriptor and registration

var LeaderboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game_engine.leaderboard.v1.LeaderboardService",
	HandlerType: (*LeaderboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "GetDailyLeaderboard"},
		{MethodName: "GetWeeklyLeaderboard"},
		{MethodName: "GetMonthlyLeaderboard"},
		{MethodName: "GetAllTimeLeaderboard"},
		{MethodName: "GetPlayerRank"},
		{MethodName: "UpdatePlayerScore"},
		{MethodName: "DistributePrizes"},
		{MethodName: "SyncLeaderboard"},
		{MethodName: "ResetLeaderboard"},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game_engine/leaderboard/v1/leaderboard_service.proto",
}

func RegisterLeaderboardServiceServer(s *grpc.Server, srv LeaderboardServiceServer) {
	s.RegisterService(&LeaderboardService_ServiceDesc, srv)
}
