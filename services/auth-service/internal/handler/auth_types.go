package handler

import (
	"fmt"
	"time"

	authv1 "github.com/game_engine/auth-service/pkg/game_engine/auth/v1"
	commonv1 "github.com/game_engine/auth-service/pkg/game_engine/common/v1"

	"github.com/game_engine/auth-service/internal/model"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type RegisterResponse = authv1.RegisterResponse
type LoginRequest = authv1.LoginRequest
type LoginResponse = authv1.LoginResponse
type RefreshTokenRequest = authv1.RefreshTokenRequest
type RefreshTokenResponse = authv1.RefreshTokenResponse
type LogoutRequest = authv1.LogoutRequest
type ValidateTokenRequest = authv1.ValidateTokenRequest
type ValidateTokenResponse = authv1.ValidateTokenResponse
type ResetPasswordRequest = authv1.ResetPasswordRequest
type ConfirmResetPasswordRequest = authv1.ConfirmResetPasswordRequest
type Enable2FARequest = authv1.Enable2FARequest
type Enable2FAResponse = authv1.Enable2FAResponse
type Verify2FARequest = authv1.Verify2FARequest
type Verify2FAResponse = authv1.Verify2FAResponse

type Disable2FARequest struct {
	UserId   string
	Password string
}

type VerifyEmailRequest = authv1.VerifyEmailRequest
type VerifyEmailResponse = authv1.VerifyEmailResponse
type VerifyPhoneRequest = authv1.VerifyPhoneRequest
type VerifyPhoneResponse = authv1.VerifyPhoneResponse
type ChangePasswordRequest = authv1.ChangePasswordRequest

type DeviceInfo = authv1.DeviceInfo
type Status = commonv1.Status
type UserRole = commonv1.UserRole
type Language = commonv1.Language

func generateReferralCode() string {
	return fmt.Sprintf("GE%d", time.Now().UnixNano()%100000)
}

func getDeviceID(info *authv1.DeviceInfo) string {
	if info == nil {
		return ""
	}
	return info.GetDeviceId()
}

func convertStatus(status model.UserStatus) commonv1.Status {
	switch status {
	case model.UserStatusActive:
		return commonv1.Status_STATUS_ACTIVE
	case model.UserStatusSuspended:
		return commonv1.Status_STATUS_SUSPENDED
	case model.UserStatusLocked:
		return commonv1.Status_STATUS_LOCKED
	default:
		return commonv1.Status_STATUS_INACTIVE
	}
}

func convertRoles(roles []model.UserRole) []commonv1.UserRole {
	result := make([]commonv1.UserRole, len(roles))
	for i, r := range roles {
		switch r {
		case model.RoleAdmin:
			result[i] = commonv1.UserRole_USER_ROLE_ADMIN
		case model.RoleSupport:
			result[i] = commonv1.UserRole_USER_ROLE_SUPPORT
		default:
			result[i] = commonv1.UserRole_USER_ROLE_PLAYER
		}
	}
	return result
}

func convertDeviceInfo(info *authv1.DeviceInfo) model.DeviceInfo {
	if info == nil {
		return model.DeviceInfo{}
	}
	return model.DeviceInfo{
		DeviceType:  info.GetDeviceType().String(),
		OSType:      info.GetOsType().String(),
		BrowserType: info.GetBrowserType().String(),
		DeviceID:    info.GetDeviceId(),
		DeviceName:  info.GetDeviceName(),
		IPAddress:   info.GetIpAddress(),
		UserAgent:   info.GetUserAgent(),
		Country:     info.GetCountry(),
		City:        info.GetCity(),
		Timezone:    info.GetTimezone(),
	}
}

type Empty = emptypb.Empty
