package model

import (
	"time"

	userv1 "gen/go/user/v1"
)

// Profile represents a player profile in the database
type Profile struct {
	ID               string     `json:"id"`
	UserID           string     `json:"user_id"`
	Email            string     `json:"email"`
	Phone            string     `json:"phone"`
	Username         string     `json:"username"`
	DisplayName      string     `json:"display_name"`
	FirstName        string     `json:"first_name"`
	LastName         string     `json:"last_name"`
	DateOfBirth      *time.Time `json:"date_of_birth"`
	Gender           string     `json:"gender"`
	AvatarURL        string     `json:"avatar_url"`
	Country          string     `json:"country"`
	Language         string     `json:"language"`
	Currency         string     `json:"currency"`
	Timezone         string     `json:"timezone"`
	Status           string     `json:"status"`
	KYCLevel         string     `json:"kyc_level"`
	EmailVerified    bool       `json:"email_verified"`
	PhoneVerified    bool       `json:"phone_verified"`
	TwoFactorEnabled bool       `json:"two_factor_enabled"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	LastLoginAt      *time.Time `json:"last_login_at"`
}

// Address represents an address in the database
type Address struct {
	ID         string    `json:"id"`
	ProfileID  string    `json:"profile_id"`
	Street     string    `json:"street"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// KYCStatus represents a KYC status record
type KYCStatus struct {
	ID              string     `json:"id"`
	UserID          string     `json:"user_id"`
	Status          string     `json:"status"`
	Level           string     `json:"level"`
	SubmittedAt     *time.Time `json:"submitted_at"`
	ReviewedAt      *time.Time `json:"reviewed_at"`
	RejectionReason string     `json:"rejection_reason"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

// KYCDocument represents a KYC document submission
type KYCDocument struct {
	ID             string     `json:"id"`
	UserID         string     `json:"user_id"`
	DocumentType   string     `json:"document_type"`
	DocumentNumber string     `json:"document_number"`
	DocumentData   string     `json:"document_data"`
	Status         string     `json:"status"`
	SubmittedAt    time.Time  `json:"submitted_at"`
	ReviewedAt     *time.Time `json:"reviewed_at"`
	ReviewerID     string     `json:"reviewer_id"`
	ReviewComment  string     `json:"review_comment"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// PlayerSettings represents player settings
type PlayerSettings struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	EmailNotifications bool      `json:"email_notifications"`
	SMSNotifications   bool      `json:"sms_notifications"`
	PushNotifications  bool      `json:"push_notifications"`
	ProfilePublic      bool      `json:"profile_public"`
	ShowOnlineStatus   bool      `json:"show_online_status"`
	AutoPlay           bool      `json:"auto_play"`
	SoundVolume        int       `json:"sound_volume"`
	Theme              string    `json:"theme"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// PlayerLimits represents responsible gaming limits
type PlayerLimits struct {
	ID               string     `json:"id"`
	UserID           string     `json:"user_id"`
	DailyLimit       int        `json:"daily_limit"`
	WeeklyLimit      int        `json:"weekly_limit"`
	MonthlyLimit     int        `json:"monthly_limit"`
	DailyLossLimit   int        `json:"daily_loss_limit"`
	SelfExclusion    bool       `json:"self_exclusion"`
	ExclusionEndDate *time.Time `json:"exclusion_end_date"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// PlayerStats represents aggregated player statistics
type PlayerStats struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	TotalDeposits    float64   `json:"total_deposits"`
	TotalWithdrawals float64   `json:"total_withdrawals"`
	TotalBets        float64   `json:"total_bets"`
	TotalWins        float64   `json:"total_wins"`
	TotalBonuses     float64   `json:"total_bonuses"`
	DepositCount     int       `json:"deposit_count"`
	WithdrawalCount  int       `json:"withdrawal_count"`
	BetCount         int       `json:"bet_count"`
	WinCount         int       `json:"win_count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ToProto converts Profile to protobuf message
func (p *Profile) ToProto() *userv1.UserProfile {
	proto := &userv1.UserProfile{
		UserId:           p.UserID,
		Email:            p.Email,
		Phone:            p.Phone,
		Username:         p.Username,
		DisplayName:      p.DisplayName,
		FirstName:        p.FirstName,
		LastName:         p.LastName,
		AvatarUrl:        p.AvatarURL,
		Country:          p.Country,
		Language:         userv1.Language(userv1.Language_value[p.Language]),
		Currency:         p.Currency,
		Timezone:         p.Timezone,
		Status:           userv1.Status(userv1.Status_value[p.Status]),
		KycLevel:         userv1.KYCLevel(userv1.KYCLevel_value[p.KYCLevel]),
		EmailVerified:    p.EmailVerified,
		PhoneVerified:    p.PhoneVerified,
		TwoFactorEnabled: p.TwoFactorEnabled,
		CreatedAt:        p.CreatedAt.String(),
		UpdatedAt:        p.UpdatedAt.String(),
	}

	if p.DateOfBirth != nil {
		proto.DateOfBirth = p.DateOfBirth.String()
	}

	if p.LastLoginAt != nil {
		proto.LastLoginAt = p.LastLoginAt.String()
	}

	return proto
}

// ToProto converts Address to protobuf message
func (a *Address) ToProto() *userv1.Address {
	return &userv1.Address{
		Street:     a.Street,
		City:       a.City,
		State:      a.State,
		PostalCode: a.PostalCode,
		Country:    a.Country,
	}
}

// ToProto converts KYCStatus to protobuf message
func (k *KYCStatus) ToProto() *userv1.KYCStatus {
	proto := &userv1.KYCStatus{
		Status:          userv1.VerificationStatus(userv1.VerificationStatus_value[k.Status]),
		Level:           userv1.KYCLevel(userv1.KYCLevel_value[k.Level]),
		RejectionReason: k.RejectionReason,
	}

	if k.SubmittedAt != nil {
		proto.SubmittedAt = k.SubmittedAt.String()
	}

	if k.ReviewedAt != nil {
		proto.ReviewedAt = k.ReviewedAt.String()
	}

	return proto
}

// ToProto converts PlayerSettings to protobuf message
func (ps *PlayerSettings) ToProto() *userv1.PlayerSettings {
	return &userv1.PlayerSettings{
		Notifications: &userv1.NotificationSettings{
			EmailEnabled: ps.EmailNotifications,
			SmsEnabled:   ps.SMSNotifications,
			PushEnabled:  ps.PushNotifications,
		},
		Privacy: &userv1.PrivacySettings{
			ProfilePublic:    ps.ProfilePublic,
			ShowOnlineStatus: ps.ShowOnlineStatus,
		},
		Gaming: &userv1.GamingSettings{
			AutoPlay:    ps.AutoPlay,
			SoundVolume: int32(ps.SoundVolume),
			Theme:       ps.Theme,
		},
	}
}

// ProfileFromProto creates Profile from protobuf message
func ProfileFromProto(proto *userv1.UserProfile) *Profile {
	p := &Profile{
		UserID:           proto.UserId,
		Email:            proto.Email,
		Phone:            proto.Phone,
		Username:         proto.Username,
		DisplayName:      proto.DisplayName,
		FirstName:        proto.FirstName,
		LastName:         proto.LastName,
		AvatarURL:        proto.AvatarUrl,
		Country:          proto.Country,
		Language:         proto.Language.String(),
		Currency:         proto.Currency,
		Timezone:         proto.Timezone,
		Status:           proto.Status.String(),
		KYCLevel:         proto.KycLevel.String(),
		EmailVerified:    proto.EmailVerified,
		PhoneVerified:    proto.PhoneVerified,
		TwoFactorEnabled: proto.TwoFactorEnabled,
	}

	return p
}
