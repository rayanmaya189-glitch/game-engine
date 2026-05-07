// Package winnersv1 contains types for the Winners gRPC service.
// This file is a manual stub. Replace with protoc-generated code when available.
package winnersv1

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// WinnersServiceServer is the server API for WinnersService.
type WinnersServiceServer interface {
	GetRecentWinners(context.Context, *GetRecentWinnersRequest) (*GetRecentWinnersResponse, error)
	GetBigWins(context.Context, *GetBigWinsRequest) (*GetBigWinsResponse, error)
	GetJackpotWinners(context.Context, *GetJackpotWinnersRequest) (*GetJackpotWinnersResponse, error)
	RecordWin(context.Context, *RecordWinRequest) (*RecordWinResponse, error)
	GetPrivacySettings(context.Context, *GetPrivacySettingsRequest) (*GetPrivacySettingsResponse, error)
	UpdatePrivacySettings(context.Context, *UpdatePrivacySettingsRequest) (*UpdatePrivacySettingsResponse, error)
}

// UnimplementedWinnersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWinnersServiceServer struct{}

func (UnimplementedWinnersServiceServer) GetRecentWinners(context.Context, *GetRecentWinnersRequest) (*GetRecentWinnersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetRecentWinners not implemented")
}
func (UnimplementedWinnersServiceServer) GetBigWins(context.Context, *GetBigWinsRequest) (*GetBigWinsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetBigWins not implemented")
}
func (UnimplementedWinnersServiceServer) GetJackpotWinners(context.Context, *GetJackpotWinnersRequest) (*GetJackpotWinnersResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetJackpotWinners not implemented")
}
func (UnimplementedWinnersServiceServer) RecordWin(context.Context, *RecordWinRequest) (*RecordWinResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method RecordWin not implemented")
}
func (UnimplementedWinnersServiceServer) GetPrivacySettings(context.Context, *GetPrivacySettingsRequest) (*GetPrivacySettingsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetPrivacySettings not implemented")
}
func (UnimplementedWinnersServiceServer) UpdatePrivacySettings(context.Context, *UpdatePrivacySettingsRequest) (*UpdatePrivacySettingsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method UpdatePrivacySettings not implemented")
}

// Message types

type GetRecentWinnersRequest struct {
	Limit int32 `json:"limit"`
}

func (m *GetRecentWinnersRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetRecentWinnersResponse struct {
	Winners   []*Winner  `json:"winners"`
	Total     int32      `json:"total"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type GetBigWinsRequest struct {
	Threshold float64 `json:"threshold"`
	Limit     int32   `json:"limit"`
}

func (m *GetBigWinsRequest) GetThreshold() float64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}
func (m *GetBigWinsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetBigWinsResponse struct {
	Wins      []*Winner  `json:"wins"`
	Threshold float64    `json:"threshold"`
	Total     int32      `json:"total"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type GetJackpotWinnersRequest struct {
	Threshold float64 `json:"threshold"`
	Limit     int32   `json:"limit"`
}

func (m *GetJackpotWinnersRequest) GetThreshold() float64 {
	if m != nil {
		return m.Threshold
	}
	return 0
}
func (m *GetJackpotWinnersRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetJackpotWinnersResponse struct {
	Winners   []*Winner  `json:"winners"`
	Threshold float64    `json:"threshold"`
	Total     int32      `json:"total"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Winner struct {
	Id          int64      `json:"id"`
	UserId      string     `json:"user_id"`
	Username    string     `json:"username"`
	DisplayName string     `json:"display_name"`
	WinAmount   float64    `json:"win_amount"`
	Currency    string     `json:"currency"`
	GameType    string     `json:"game_type"`
	GameName    string     `json:"game_name"`
	WinType     string     `json:"win_type"`
	Multiplier  float64    `json:"multiplier"`
	Timestamp   *time.Time `json:"timestamp"`
}

type RecordWinRequest struct {
	UserId     string  `json:"user_id"`
	Username   string  `json:"username"`
	WinAmount  float64 `json:"win_amount"`
	Currency   string  `json:"currency"`
	GameType   string  `json:"game_type"`
	GameName   string  `json:"game_name"`
	Multiplier float64 `json:"multiplier"`
}

func (m *RecordWinRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}
func (m *RecordWinRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}
func (m *RecordWinRequest) GetWinAmount() float64 {
	if m != nil {
		return m.WinAmount
	}
	return 0
}
func (m *RecordWinRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}
func (m *RecordWinRequest) GetGameType() string {
	if m != nil {
		return m.GameType
	}
	return ""
}
func (m *RecordWinRequest) GetGameName() string {
	if m != nil {
		return m.GameName
	}
	return ""
}
func (m *RecordWinRequest) GetMultiplier() float64 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

type RecordWinResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetPrivacySettingsRequest struct {
	UserId string `json:"user_id"`
}

func (m *GetPrivacySettingsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetPrivacySettingsResponse struct {
	UserId            string `json:"user_id"`
	AnonymizeName     bool   `json:"anonymize_name"`
	ShowOnLeaderboard bool   `json:"show_on_leaderboard"`
	ShowOnJackpotList bool   `json:"show_on_jackpot_list"`
	OptOutOfShowcase  bool   `json:"opt_out_of_showcase"`
}

type UpdatePrivacySettingsRequest struct {
	UserId            string `json:"user_id"`
	AnonymizeName     bool   `json:"anonymize_name"`
	ShowOnLeaderboard bool   `json:"show_on_leaderboard"`
	ShowOnJackpotList bool   `json:"show_on_jackpot_list"`
	OptOutOfShowcase  bool   `json:"opt_out_of_showcase"`
}

func (m *UpdatePrivacySettingsRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}
func (m *UpdatePrivacySettingsRequest) GetAnonymizeName() bool {
	if m != nil {
		return m.AnonymizeName
	}
	return false
}
func (m *UpdatePrivacySettingsRequest) GetShowOnLeaderboard() bool {
	if m != nil {
		return m.ShowOnLeaderboard
	}
	return false
}
func (m *UpdatePrivacySettingsRequest) GetShowOnJackpotList() bool {
	if m != nil {
		return m.ShowOnJackpotList
	}
	return false
}
func (m *UpdatePrivacySettingsRequest) GetOptOutOfShowcase() bool {
	if m != nil {
		return m.OptOutOfShowcase
	}
	return false
}

type UpdatePrivacySettingsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Service descriptor and registration

var WinnersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game_engine.winners.v1.WinnersService",
	HandlerType: (*WinnersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "GetRecentWinners"},
		{MethodName: "GetBigWins"},
		{MethodName: "GetJackpotWinners"},
		{MethodName: "RecordWin"},
		{MethodName: "GetPrivacySettings"},
		{MethodName: "UpdatePrivacySettings"},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game_engine/winners/v1/winners_service.proto",
}

func RegisterWinnersServiceServer(s *grpc.Server, srv WinnersServiceServer) {
	s.RegisterService(&WinnersService_ServiceDesc, srv)
}
