package handler

import (
	"fmt"
	"time"

	authv1 "game_engine/gen/go/auth/v1"

	"github.com/game_engine/auth-service/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RegisterResponse struct {
	UserId                    string
	AccessToken               string
	RefreshToken              string
	ExpiresAt                 *timestamppb.Timestamp
	EmailVerificationRequired bool
	PhoneVerificationRequired bool
	Message                   string
}
type LoginRequest struct {
	Identifier string
	Password   string
	DeviceInfo *DeviceInfo
	RememberMe bool
}

type LoginResponse struct {
	UserId       string
	AccessToken  string
	RefreshToken string
	ExpiresAt    *timestamppb.Timestamp
	Requires2Fa  bool
	SessionId    string
	UserStatus   Status
	Message      string
}

type RefreshTokenRequest struct {
	RefreshToken string
	DeviceInfo   *DeviceInfo
}

type RefreshTokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    *timestamppb.Timestamp
	SessionId    string
}

type LogoutRequest struct {
	SessionId   string
	AllSessions bool
}

type ValidateTokenRequest struct {
	Token        string
	ExpectedType string
}

type ValidateTokenResponse struct {
	Valid     bool
	UserId    string
	SessionId string
	TokenType string
	ExpiresAt *timestamppb.Timestamp
	Roles     []UserRole
}

type ResetPasswordRequest struct {
	Identifier string
}

type ConfirmResetPasswordRequest struct {
	Identifier      string
	Token           string
	NewPassword     string
	ConfirmPassword string
}

type Enable2FARequest struct {
	UserId   string
	Password string
}

type Enable2FAResponse struct {
	Secret      string
	QrCodeUrl   string
	BackupCodes []string
}

type Verify2FARequest struct {
	UserId       string
	Code         string
	IsBackupCode bool
}

type Verify2FAResponse struct {
	Success      bool
	Message      string
	AccessToken  string
	RefreshToken string
	ExpiresAt    *timestamppb.Timestamp
}

type Disable2FARequest struct {
	UserId   string
	Password string
}

type VerifyEmailRequest struct {
	UserId string
	Code   string
}

type VerifyEmailResponse struct {
	Success      bool
	Message      string
	AccessToken  string
	RefreshToken string
}

type VerifyPhoneRequest struct {
	UserId string
	Code   string
}

type VerifyPhoneResponse struct {
	Success      bool
	Message      string
	AccessToken  string
	RefreshToken string
}

type ChangePasswordRequest struct {
	UserId          string
	CurrentPassword string
	NewPassword     string
	ConfirmPassword string
}

// Enum types (would normally be generated from proto)
type Status int32
type UserRole int32
type Language int32
type DeviceType int32
type OSType int32
type BrowserType int32

const (
	Status_STATUS_UNSPECIFIED Status = 0
	Status_STATUS_ACTIVE      Status = 1
	Status_STATUS_SUSPENDED   Status = 2
	Status_STATUS_LOCKED      Status = 3
	Status_STATUS_INACTIVE    Status = 4
)

const (
	UserRole_ROLE_UNSPECIFIED UserRole = 0
	UserRole_ROLE_PLAYER      UserRole = 1
	UserRole_ROLE_ADMIN       UserRole = 2
	UserRole_ROLE_SUPPORT     UserRole = 3
)

const (
	Language_LANGUAGE_UNSPECIFIED Language = 0
	Language_LANGUAGE_EN          Language = 1
	Language_LANGUAGE_TH          Language = 2
	Language_LANGUAGE_VI          Language = 3
	Language_LANGUAGE_ID          Language = 4
	Language_LANGUAGE_MS          Language = 5
	Language_LANGUAGE_ZH          Language = 6
)

const (
	DeviceType_DEVICE_TYPE_UNSPECIFIED DeviceType = 0
	DeviceType_DEVICE_TYPE_MOBILE      DeviceType = 1
	DeviceType_DEVICE_TYPE_DESKTOP     DeviceType = 2
	DeviceType_DEVICE_TYPE_TABLET      DeviceType = 3
	DeviceType_DEVICE_TYPE_TV          DeviceType = 4
)

const (
	OSType_OS_TYPE_UNSPECIFIED OSType = 0
	OSType_OS_TYPE_IOS         OSType = 1
	OSType_OS_TYPE_ANDROID     OSType = 2
	OSType_OS_TYPE_WINDOWS     OSType = 3
	OSType_OS_TYPE_MACOS       OSType = 4
	OSType_OS_TYPE_LINUX       OSType = 5
)

const (
	BrowserType_BROWSER_TYPE_UNSPECIFIED BrowserType = 0
	BrowserType_BROWSER_TYPE_CHROME      BrowserType = 1
	BrowserType_BROWSER_TYPE_FIREFOX     BrowserType = 2
	BrowserType_BROWSER_TYPE_SAFARI      BrowserType = 3
	BrowserType_BROWSER_TYPE_EDGE        BrowserType = 4
	BrowserType_BROWSER_TYPE_OPERA       BrowserType = 5
)

type DeviceInfo struct {
	DeviceType  DeviceType
	OSType      OSType
	BrowserType BrowserType
	DeviceId    string
	DeviceName  string
	IpAddress   string
	UserAgent   string
	Country     string
	City        string
	Timezone    string
}

func (d DeviceType) String() string {
	return "DEVICE_TYPE_UNSPECIFIED"
}

func (o OSType) String() string {
	return "OS_TYPE_UNSPECIFIED"
}

func (b BrowserType) String() string {
	return "BROWSER_TYPE_UNSPECIFIED"
}

func (l Language) String() string {
	return "LANGUAGE_UNSPECIFIED"
}

// Helper functions

func generateReferralCode() string {
	return fmt.Sprintf("GE%d", time.Now().UnixNano()%100000)
}

func getDeviceID(info *authv1.DeviceInfo) string {
	if info == nil {
		return ""
	}
	return info.GetDeviceId()
}

func convertStatus(status model.UserStatus) Status {
	switch status {
	case model.UserStatusActive:
		return Status_STATUS_ACTIVE
	case model.UserStatusSuspended:
		return Status_STATUS_SUSPENDED
	case model.UserStatusLocked:
		return Status_STATUS_LOCKED
	default:
		return Status_STATUS_INACTIVE
	}
}

func convertRoles(roles []model.UserRole) []UserRole {
	result := make([]UserRole, len(roles))
	for i, r := range roles {
		switch r {
		case model.RoleAdmin:
			result[i] = UserRole_ROLE_ADMIN
		case model.RoleSupport:
			result[i] = UserRole_ROLE_SUPPORT
		default:
			result[i] = UserRole_ROLE_PLAYER
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
