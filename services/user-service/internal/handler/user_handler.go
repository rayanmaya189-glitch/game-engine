package handler

import (
	"context"
	"errors"

	commonv1 "github.com/game_engine/common-service/proto/gen/go/common/v1"
	userv1 "github.com/game_engine/common-service/proto/gen/go/user/v1"

	"github.com/game_engine/user-service/internal/model"
	"github.com/game_engine/user-service/internal/service"
)

var ErrUnauthorized = errors.New("unauthorized")

// UserHandler handles gRPC requests for user service
type UserHandler struct {
	userv1.UnimplementedUserServiceServer
	userService *service.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetProfile retrieves a player profile
func (h *UserHandler) GetProfile(ctx context.Context, req *userv1.GetProfileRequest) (*userv1.GetProfileResponse, error) {
	profile, err := h.userService.GetProfile(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &userv1.GetProfileResponse{
		Profile: profile.ToProto(),
	}, nil
}

// UpdateProfile updates a player's profile
func (h *UserHandler) UpdateProfile(ctx context.Context, req *userv1.UpdateProfileRequest) (*userv1.UpdateProfileResponse, error) {
	// Get user ID from context (set by auth middleware)
	userID := ctx.Value("user_id")
	if userID == nil {
		return nil, ErrUnauthorized
	}

	updates := &model.Profile{
		DisplayName: req.DisplayName,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Gender:      req.Gender.String(),
		Language:    req.Language.String(),
		Timezone:    req.Timezone,
	}

	if req.Address != nil {
		// Handle address update
	}

	profile, err := h.userService.UpdateProfile(ctx, userID.(string), updates)
	if err != nil {
		return nil, err
	}

	return &userv1.UpdateProfileResponse{
		Profile: profile.ToProto(),
		Message: "Profile updated successfully",
	}, nil
}

// GetKYCStatus retrieves KYC status
func (h *UserHandler) GetKYCStatus(ctx context.Context, req *userv1.GetKYCStatusRequest) (*userv1.GetKYCStatusResponse, error) {
	kycStatus, err := h.userService.GetKYCStatus(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &userv1.GetKYCStatusResponse{
		Status:          kycStatus.ToProto(),
		RejectionReason: kycStatus.RejectionReason,
	}, nil
}

// SubmitKYC submits KYC documents
func (h *UserHandler) SubmitKYC(ctx context.Context, req *userv1.SubmitKYCRequest) (*userv1.SubmitKYCResponse, error) {
	// Get user ID from context
	userID := ctx.Value("user_id")
	if userID == nil {
		return nil, ErrUnauthorized
	}

	err := h.userService.SubmitKYC(ctx, userID.(string), req.DocumentType.String(), req.DocumentNumber, req.DocumentData)
	if err != nil {
		return nil, err
	}

	return &userv1.SubmitKYCResponse{
		Success: true,
		Message: "KYC documents submitted successfully",
	}, nil
}

// GetPlayerSettings retrieves player settings
func (h *UserHandler) GetPlayerSettings(ctx context.Context, req *userv1.GetPlayerSettingsRequest) (*userv1.GetPlayerSettingsResponse, error) {
	settings, err := h.userService.GetPlayerSettings(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &userv1.GetPlayerSettingsResponse{
		Settings: settings.ToProto(),
	}, nil
}

// UpdatePlayerSettings updates player settings
func (h *UserHandler) UpdatePlayerSettings(ctx context.Context, req *userv1.UpdatePlayerSettingsRequest) (*userv1.UpdatePlayerSettingsResponse, error) {
	// Get user ID from context
	userID := ctx.Value("user_id")
	if userID == nil {
		return nil, ErrUnauthorized
	}

	settings, err := h.userService.UpdatePlayerSettings(ctx, userID.(string), model.PlayerSettingsFromProto(req.Settings))
	if err != nil {
		return nil, err
	}

	return &userv1.UpdatePlayerSettingsResponse{
		Settings: settings.ToProto(),
		Message:  "Settings updated successfully",
	}, nil
}

// GetPlayerByAdmin retrieves a player by identifier (admin)
func (h *UserHandler) GetPlayerByAdmin(ctx context.Context, req *userv1.GetPlayerByAdminRequest) (*userv1.GetPlayerByAdminResponse, error) {
	profile, kycStatus, err := h.userService.GetPlayerByAdmin(ctx, req.Identifier)
	if err != nil {
		return nil, err
	}

	return &userv1.GetPlayerByAdminResponse{
		Profile: profile.ToProto(),
		AccountStatus: &userv1.AccountStatus{
			Status:        commonv1.Status(commonv1.Status_value[profile.Status]),
			EmailVerified: profile.EmailVerified,
			PhoneVerified: profile.PhoneVerified,
			KycVerified:   kycStatus.Status == "VERIFICATION_STATUS_VERIFIED",
		},
		KycStatus: kycStatus.ToProto(),
	}, nil
}

// ListPlayers lists players with filters and pagination (admin)
func (h *UserHandler) ListPlayers(ctx context.Context, req *userv1.ListPlayersRequest) (*userv1.ListPlayersResponse, error) {
	status := ""
	kycLevel := ""
	country := ""
	search := ""

	if req.Filters != nil {
		status = req.Filters.Status.String()
		kycLevel = req.Filters.KycLevel.String()
		country = req.Filters.Country
		search = req.Filters.Search
	}

	page := 1
	pageSize := 20

	if req.Pagination != nil {
		page = int(req.Pagination.Page)
		pageSize = int(req.Pagination.PageSize)
	}

	profiles, total, err := h.userService.ListPlayers(ctx, status, kycLevel, country, search, page, pageSize)
	if err != nil {
		return nil, err
	}

	protoProfiles := make([]*userv1.UserProfile, len(profiles))
	for i, p := range profiles {
		protoProfiles[i] = p.ToProto()
	}

	return &userv1.ListPlayersResponse{
		Players: protoProfiles,
		Pagination: &commonv1.PaginationResponse{
			Page:       int32(page),
			PageSize:   int32(pageSize),
			TotalItems: int32(total),
			TotalPages: int32((total + pageSize - 1) / pageSize),
		},
	}, nil
}

// UpdatePlayerStatus updates player status (admin)
func (h *UserHandler) UpdatePlayerStatus(ctx context.Context, req *userv1.UpdatePlayerStatusRequest) (*userv1.UpdatePlayerStatusResponse, error) {
	err := h.userService.UpdatePlayerStatus(ctx, req.UserId, req.Status.String(), req.Reason)
	if err != nil {
		return nil, err
	}

	return &userv1.UpdatePlayerStatusResponse{
		Success: true,
		Message: "Player status updated successfully",
		Status:  req.Status,
	}, nil
}

// Import userv1 for pagination
var _ = userv1.ListPlayersResponse{}
