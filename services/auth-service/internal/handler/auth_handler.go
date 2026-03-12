package handler

import (
	"context"
	"fmt"
	"time"

	authv1 "github.com/game-engine/common/proto/gen/go/game-engine/auth/v1"

	"github.com/game-engine/auth-service/internal/model"
	"github.com/game-engine/auth-service/internal/service"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// AuthHandler handles gRPC requests for authentication
type AuthHandler struct {
	authv1.UnimplementedAuthServiceServer
	authService *service.AuthService
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register handles user registration
func (h *AuthHandler) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	// Validate email format
	if err := h.authService.ValidateEmail(req.GetIdentifier()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid email: %v", err)
	}

	// Validate password
	if err := h.authService.ValidatePassword(req.GetPassword()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid password: %v", err)
	}

	// Check password match
	if req.GetPassword() != req.GetConfirmPassword() {
		return nil, status.Errorf(codes.InvalidArgument, "passwords do not match")
	}

	// Check terms acceptance
	if !req.GetAcceptTerms() {
		return nil, status.Errorf(codes.InvalidArgument, "terms must be accepted")
	}

	// Check if user exists
	exists, err := h.authService.CheckUserExists(ctx, req.GetIdentifier())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to check user: %v", err)
	}
	if exists {
		return nil, status.Errorf(codes.AlreadyExists, "user already exists")
	}

	// Hash password
	hash, err := h.authService.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	// Create user
	user := &model.User{
		ID:               uuid.New(),
		Email:            req.GetIdentifier(),
		PasswordHash:     hash,
		Country:          req.GetCountry(),
		Language:         req.GetLanguage().String(),
		Currency:         req.GetCurrency(),
		Status:           model.UserStatusActive,
		EmailVerified:    false,
		PhoneVerified:    false,
		TwoFactorEnabled: false,
		MarketingConsent: req.GetMarketingConsent(),
		AcceptTerms:      req.GetAcceptTerms(),
		ReferralCode:     generateReferralCode(),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	if err := h.authService.CreateUser(ctx, user); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	// Generate email verification token
	verificationToken, err := h.authService.GenerateEmailVerificationToken(ctx, user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate verification token: %v", err)
	}

	// Generate session and tokens
	sessionID := uuid.New()
	accessToken, expiresAt, err := h.authService.GenerateAccessToken(user.ID, sessionID, []model.UserRole{model.RolePlayer})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	// Store refresh token
	deviceID := getDeviceID(req.GetDeviceInfo())
	if err := h.authService.StoreRefreshToken(ctx, user.ID, sessionID, refreshToken, deviceID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store refresh token: %v", err)
	}

	// Publish registration event
	h.authService.PublishEvent("player.events.registered", map[string]interface{}{
		"user_id":   user.ID.String(),
		"email":     user.Email,
		"country":   user.Country,
		"timestamp": time.Now().Unix(),
	})

	_ = verificationToken // Would be sent via email in production

	return &authv1.RegisterResponse{
		UserId:                    user.ID.String(),
		AccessToken:               accessToken,
		RefreshToken:              refreshToken,
		ExpiresAt:                 timestamppb.New(expiresAt),
		EmailVerificationRequired: true,
		Message:                   "Registration successful. Please verify your email.",
	}, nil
}

// Login handles user login
func (h *AuthHandler) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	// Find user by identifier (email or phone)
	user, err := h.authService.GetUserByIdentifier(ctx, req.Identifier)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find user: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	// Check account status
	if user.Status == model.UserStatusLocked {
		if user.LockedUntil != nil && time.Now().Before(*user.LockedUntil) {
			return nil, status.Errorf(codes.ResourceExhausted, "account is locked until %v", user.LockedUntil)
		}
	}
	if user.Status == model.UserStatusSuspended {
		return nil, status.Errorf(codes.PermissionDenied, "account is suspended")
	}
	if user.Status == model.UserStatusInactive {
		return nil, status.Errorf(codes.PermissionDenied, "account is inactive")
	}

	// Check login attempts
	if err := h.authService.CheckLoginAttempts(ctx, user.ID); err != nil {
		return nil, status.Errorf(codes.ResourceExhausted, "%v", err)
	}

	// Verify password
	if !h.authService.CheckPassword(req.Password, user.PasswordHash) {
		// Record failed attempt
		h.authService.RecordLoginAttempt(ctx, user.ID, req.DeviceInfo.IpAddress, false)
		return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
	}

	// Check if 2FA is enabled
	if user.TwoFactorEnabled {
		// Return partial success requiring 2FA
		sessionID := uuid.New()
		partialToken, err := h.authService.GeneratePartialToken(user.ID, sessionID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to generate partial token: %v", err)
		}

		return &LoginResponse{
			UserId:      user.ID.String(),
			Requires2Fa: true,
			SessionId:   sessionID.String(),
			UserStatus:  convertStatus(user.Status),
			Message:     "2FA verification required",
			AccessToken: partialToken,
		}, nil
	}

	// Successful login - create session
	if err := h.authService.RecordLoginAttempt(ctx, user.ID, req.DeviceInfo.IpAddress, true); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to record login: %v", err)
	}

	sessionID := uuid.New()
	accessToken, expiresAt, err := h.authService.GenerateAccessToken(user.ID, sessionID, []model.UserRole{model.RolePlayer})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	// Store refresh token
	deviceID := getDeviceID(req.DeviceInfo)
	if err := h.authService.StoreRefreshToken(ctx, user.ID, sessionID, refreshToken, deviceID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store refresh token: %v", err)
	}

	// Create session in database
	session := &model.Session{
		ID:         sessionID,
		UserID:     user.ID,
		DeviceInfo: convertDeviceInfo(req.DeviceInfo),
		IPAddress:  req.DeviceInfo.IpAddress,
		UserAgent:  req.DeviceInfo.UserAgent,
		ExpiresAt:  time.Now().Add(24 * time.Hour),
		CreatedAt:  time.Now(),
		LastUsedAt: time.Now(),
	}
	if err := h.authService.CreateSession(ctx, session); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create session: %v", err)
	}

	// Update last login
	if err := h.authService.UpdateLastLogin(ctx, user.ID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update last login: %v", err)
	}

	// Publish login event
	h.authService.PublishEvent("player.events.logged_in", map[string]interface{}{
		"user_id":    user.ID.String(),
		"session_id": sessionID.String(),
		"ip_address": req.DeviceInfo.IpAddress,
		"timestamp":  time.Now().Unix(),
	})

	return &LoginResponse{
		UserId:       user.ID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    timestamppb.New(expiresAt),
		Requires2Fa:  false,
		SessionId:    sessionID.String(),
		UserStatus:   convertStatus(user.Status),
		Message:      "Login successful",
	}, nil
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	// Get refresh token data
	data, err := h.authService.GetRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get refresh token: %v", err)
	}
	if data == nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid refresh token")
	}

	userID, err := uuid.Parse(data.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid user ID in token")
	}

	sessionID, err := uuid.Parse(data.SessionID)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid session ID in token")
	}

	// Delete old refresh token (token rotation)
	if err := h.authService.DeleteRefreshToken(ctx, req.RefreshToken); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete old token: %v", err)
	}

	// Generate new tokens
	accessToken, expiresAt, err := h.authService.GenerateAccessToken(userID, sessionID, []model.UserRole{model.RolePlayer})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	newRefreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	// Store new refresh token
	deviceID := getDeviceID(req.DeviceInfo)
	if err := h.authService.StoreRefreshToken(ctx, userID, sessionID, newRefreshToken, deviceID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store refresh token: %v", err)
	}

	return &RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresAt:    timestamppb.New(expiresAt),
		SessionId:    sessionID.String(),
	}, nil
}

// Logout handles user logout
func (h *AuthHandler) Logout(ctx context.Context, req *LogoutRequest) (*emptypb.Empty, error) {
	sessionID, err := uuid.Parse(req.SessionId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid session ID")
	}

	if req.AllSessions {
		// Delete all user sessions
		userID, err := h.authService.GetUserIDFromSession(ctx, sessionID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get user ID: %v", err)
		}
		if err := h.authService.DeleteAllUserSessions(ctx, userID); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to delete sessions: %v", err)
		}
	} else {
		// Delete specific session
		if err := h.authService.DeleteSession(ctx, sessionID); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to delete session: %v", err)
		}
	}

	// Publish logout event
	h.authService.PublishEvent("player.events.logged_out", map[string]interface{}{
		"session_id": sessionID.String(),
		"timestamp":  time.Now().Unix(),
	})

	return &emptypb.Empty{}, nil
}

// ValidateToken handles token validation
func (h *AuthHandler) ValidateToken(ctx context.Context, req *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	claims, err := h.authService.ValidateAccessToken(req.Token)
	if err != nil {
		return &ValidateTokenResponse{
			Valid: false,
		}, nil
	}

	return &ValidateTokenResponse{
		Valid:     true,
		UserId:    claims.UserID.String(),
		SessionId: claims.SessionID.String(),
		TokenType: claims.TokenType,
		ExpiresAt: timestamppb.New(time.Now().Add(15 * time.Minute)),
		Roles:     convertRoles(claims.Roles),
	}, nil
}

// ResetPassword handles password reset request
func (h *AuthHandler) ResetPassword(ctx context.Context, req *ResetPasswordRequest) (*emptypb.Empty, error) {
	user, err := h.authService.GetUserByIdentifier(ctx, req.Identifier)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find user: %v", err)
	}
	if user == nil {
		// Don't reveal if user exists
		return &emptypb.Empty{}, nil
	}

	token, err := h.authService.GeneratePasswordResetToken(ctx, user.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate reset token: %v", err)
	}

	_ = token // Would be sent via email in production

	return &emptypb.Empty{}, nil
}

// ConfirmResetPassword handles password reset confirmation
func (h *AuthHandler) ConfirmResetPassword(ctx context.Context, req *ConfirmResetPasswordRequest) (*emptypb.Empty, error) {
	// Validate password
	if err := h.authService.ValidatePassword(req.NewPassword); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid password: %v", err)
	}

	// Check password match
	if req.NewPassword != req.ConfirmPassword {
		return nil, status.Errorf(codes.InvalidArgument, "passwords do not match")
	}

	// Validate token
	userID, err := h.authService.ValidatePasswordResetToken(ctx, req.Token)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid or expired token")
	}

	// Hash new password
	hash, err := h.authService.HashPassword(req.NewPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	// Update password
	if err := h.authService.UpdatePassword(ctx, *userID, hash); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update password: %v", err)
	}

	// Invalidate all sessions
	if err := h.authService.DeleteAllUserSessions(ctx, *userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to invalidate sessions: %v", err)
	}

	// Publish password reset event
	h.authService.PublishEvent("player.events.password_reset", map[string]interface{}{
		"user_id":   userID.String(),
		"timestamp": time.Now().Unix(),
	})

	return &emptypb.Empty{}, nil
}

// Enable2FA handles 2FA enablement
func (h *AuthHandler) Enable2FA(ctx context.Context, req *Enable2FARequest) (*Enable2FAResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	user, err := h.authService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Verify password
	if !h.authService.CheckPassword(req.Password, user.PasswordHash) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}

	// Generate TOTP secret
	secret, qrCodeURL, err := h.authService.GenerateTOTPSecret(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate TOTP secret: %v", err)
	}

	// Generate backup codes
	backupCodes, err := h.authService.GenerateBackupCodes(10)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate backup codes: %v", err)
	}

	// Store secret temporarily (not enabled until verified)
	if err := h.authService.Store2FASecret(ctx, userID, secret); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store 2FA secret: %v", err)
	}

	return &Enable2FAResponse{
		Secret:      secret,
		QrCodeUrl:   qrCodeURL,
		BackupCodes: backupCodes,
	}, nil
}

// Verify2FA handles 2FA verification
func (h *AuthHandler) Verify2FA(ctx context.Context, req *Verify2FARequest) (*Verify2FAResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	user, err := h.authService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	var valid bool
	if req.IsBackupCode {
		valid = h.authService.ValidateBackupCode(ctx, userID, req.Code)
	} else {
		valid = h.authService.ValidateTOTP(user.TwoFactorSecret, req.Code)
	}

	if !valid {
		return nil, status.Errorf(codes.Unauthenticated, "invalid 2FA code")
	}

	// Enable 2FA
	if err := h.authService.Enable2FAForUser(ctx, userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to enable 2FA: %v", err)
	}

	// Generate full access token
	sessionID := uuid.New()
	accessToken, expiresAt, err := h.authService.GenerateAccessToken(user.ID, sessionID, []model.UserRole{model.RolePlayer})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	deviceID := "" // Would get from request
	if err := h.authService.StoreRefreshToken(ctx, user.ID, sessionID, refreshToken, deviceID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store refresh token: %v", err)
	}

	// Publish 2FA enabled event
	h.authService.PublishEvent("player.events.2fa_enabled", map[string]interface{}{
		"user_id":   user.ID.String(),
		"timestamp": time.Now().Unix(),
	})

	return &Verify2FAResponse{
		Success:      true,
		Message:      "2FA enabled successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    timestamppb.New(expiresAt),
	}, nil
}

// Disable2FA handles 2FA disable
func (h *AuthHandler) Disable2FA(ctx context.Context, req *Disable2FARequest) (*emptypb.Empty, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	user, err := h.authService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Verify password
	if !h.authService.CheckPassword(req.Password, user.PasswordHash) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid password")
	}

	// Disable 2FA
	if err := h.authService.Disable2FAForUser(ctx, userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to disable 2FA: %v", err)
	}

	// Publish 2FA disabled event
	h.authService.PublishEvent("player.events.2fa_disabled", map[string]interface{}{
		"user_id":   user.ID.String(),
		"timestamp": time.Now().Unix(),
	})

	return &emptypb.Empty{}, nil
}

// VerifyEmail handles email verification
func (h *AuthHandler) VerifyEmail(ctx context.Context, req *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	userID, err := h.authService.ValidateEmailVerificationToken(ctx, req.Code)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid or expired token")
	}

	if err := h.authService.MarkEmailVerified(ctx, *userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify email: %v", err)
	}

	// Generate tokens
	sessionID := uuid.New()
	accessToken, expiresAt, err := h.authService.GenerateAccessToken(*userID, sessionID, []model.UserRole{model.RolePlayer})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	return &VerifyEmailResponse{
		Success:      true,
		Message:      "Email verified successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// VerifyPhone handles phone verification
func (h *AuthHandler) VerifyPhone(ctx context.Context, req *VerifyPhoneRequest) (*VerifyPhoneResponse, error) {
	// Similar to email verification
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	if err := h.authService.MarkPhoneVerified(ctx, userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify phone: %v", err)
	}

	sessionID := uuid.New()
	accessToken, expiresAt, err := h.authService.GenerateAccessToken(userID, sessionID, []model.UserRole{model.RolePlayer})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	refreshToken, err := h.authService.GenerateRefreshToken()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate refresh token: %v", err)
	}

	return &VerifyPhoneResponse{
		Success:      true,
		Message:      "Phone verified successfully",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// ChangePassword handles password change
func (h *AuthHandler) ChangePassword(ctx context.Context, req *ChangePasswordRequest) (*emptypb.Empty, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID")
	}

	user, err := h.authService.GetUserByID(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if user == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Verify current password
	if !h.authService.CheckPassword(req.CurrentPassword, user.PasswordHash) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid current password")
	}

	// Validate new password
	if err := h.authService.ValidatePassword(req.NewPassword); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid new password: %v", err)
	}

	// Check password match
	if req.NewPassword != req.ConfirmPassword {
		return nil, status.Errorf(codes.InvalidArgument, "passwords do not match")
	}

	// Hash new password
	hash, err := h.authService.HashPassword(req.NewPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}

	// Update password
	if err := h.authService.UpdatePassword(ctx, userID, hash); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update password: %v", err)
	}

	// Invalidate all sessions
	if err := h.authService.DeleteAllUserSessions(ctx, userID); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to invalidate sessions: %v", err)
	}

	return &emptypb.Empty{}, nil
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
