package handler

import (
	"context"
	"time"

	authv1 "game_engine/gen/go/auth/v1"

	"github.com/game_engine/auth-service/internal/model"
	"github.com/game_engine/auth-service/internal/service"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
