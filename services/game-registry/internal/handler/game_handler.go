package handler

import (
	"context"
	"net/http"

	gamesv1 "github.com/gameengine/gen/go/gameengine/game/v1"

	"github.com/gameengine/game-registry/internal/enums"
	"github.com/gameengine/game-registry/internal/model"
	"github.com/gameengine/game-registry/internal/service"
	"github.com/gin-gonic/gin"
)

// GameHandler handles HTTP and gRPC requests for games
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

// ListGames handles GET /games
func (h *GameHandler) ListGames(c *gin.Context) {
	req := service.ListGamesRequest{
		CategoryID:       c.Query("category_id"),
		ProviderID:       c.Query("provider_id"),
		MobileSupported:  c.Query("mobile_supported") == "true",
		DesktopSupported: c.Query("desktop_supported") == "true",
		IsFeatured:       c.Query("is_featured") == "true",
		IsJackpot:        c.Query("is_jackpot") == "true",
		Query:            c.Query("query"),
		SortBy:           c.Query("sort_by"),
	}

	// Parse pagination
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "20")
	req.Pagination.Page = parseInt32(page, 1)
	req.Pagination.PageSize = parseInt32(pageSize, 20)

	// Parse status
	if status := c.Query("status"); status != "" {
		req.Status = parseInt32(status, 0)
	}

	ctx := context.Background()
	resp, err := h.gameService.ListGames(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetGame handles GET /games/:id
func (h *GameHandler) GetGame(c *gin.Context) {
	gameID := c.Param("id")
	if gameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "game id is required"})
		return
	}

	ctx := context.Background()
	game, err := h.gameService.GetGame(ctx, gameID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

// GetGameConfig handles GET /games/:id/config
func (h *GameHandler) GetGameConfig(c *gin.Context) {
	gameID := c.Param("id")
	if gameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "game id is required"})
		return
	}

	userID := c.Query("user_id")
	deviceType := enums.DeviceType(parseInt32(c.Query("device_type"), 1))
	language := enums.GameLanguage(parseInt32(c.Query("language"), 1))
	currency := c.DefaultQuery("currency", "USD")
	sessionID := c.Query("session_id")

	ctx := context.Background()
	config, err := h.gameService.GetGameConfig(ctx, gameID, userID, deviceType, language, currency, sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}

// GetGameURL handles POST /games/:id/url
func (h *GameHandler) GetGameURL(c *gin.Context) {
	gameID := c.Param("id")
	if gameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "game id is required"})
		return
	}

	var req struct {
		UserID     string `json:"user_id"`
		DeviceType int32  `json:"device_type"`
		SessionID  string `json:"session_id"`
		Language   int32  `json:"language"`
		Currency   string `json:"currency"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		req.DeviceType = 1 // Default to desktop
		req.Language = 1   // Default to English
		req.Currency = "USD"
	}

	ctx := context.Background()
	result, err := h.gameService.GetGameURL(ctx, gameID, req.UserID, enums.DeviceType(req.DeviceType), req.SessionID, enums.GameLanguage(req.Language), req.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetCategories handles GET /categories
func (h *GameHandler) GetCategories(c *gin.Context) {
	includeGamesCount := c.Query("include_games_count") == "true"

	ctx := context.Background()
	categories, err := h.gameService.GetCategories(ctx, includeGamesCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// GetProviders handles GET /providers
func (h *GameHandler) GetProviders(c *gin.Context) {
	activeOnly := c.Query("active_only") == "true"

	ctx := context.Background()
	providers, err := h.gameService.GetProviders(ctx, activeOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"providers": providers})
}

// SearchGames handles GET /games/search
func (h *GameHandler) SearchGames(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query is required"})
		return
	}

	limit := parseInt32(c.Query("limit"), 20)
	categoryID := c.Query("category_id")

	ctx := context.Background()
	games, err := h.gameService.SearchGames(ctx, query, int(limit), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"games": games})
}

// GetFeaturedGames handles GET /games/featured
func (h *GameHandler) GetFeaturedGames(c *gin.Context) {
	limit := parseInt32(c.Query("limit"), 10)
	categoryID := c.Query("category_id")

	ctx := context.Background()
	games, err := h.gameService.GetFeaturedGames(ctx, int(limit), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"games": games})
}

// GetPopularGames handles GET /games/popular
func (h *GameHandler) GetPopularGames(c *gin.Context) {
	limit := parseInt32(c.Query("limit"), 10)
	categoryID := c.Query("category_id")

	ctx := context.Background()
	games, err := h.gameService.GetPopularGames(ctx, int(limit), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"games": games})
}

// GetNewGames handles GET /games/new
func (h *GameHandler) GetNewGames(c *gin.Context) {
	limit := parseInt32(c.Query("limit"), 10)
	categoryID := c.Query("category_id")

	ctx := context.Background()
	games, err := h.gameService.GetNewGames(ctx, int(limit), categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"games": games})
}

// CreateGame handles POST /games (admin)
func (h *GameHandler) CreateGame(c *gin.Context) {
	var game model.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	err := h.gameService.CreateGame(ctx, &game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, game)
}

// UpdateGame handles PUT /games/:id (admin)
func (h *GameHandler) UpdateGame(c *gin.Context) {
	gameID := c.Param("id")
	if gameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "game id is required"})
		return
	}

	var game model.Game
	game.ID = gameID
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	err := h.gameService.UpdateGame(ctx, &game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, game)
}

// ToggleGame handles POST /games/:id/toggle (admin)
func (h *GameHandler) ToggleGame(c *gin.Context) {
	gameID := c.Param("id")
	if gameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "game id is required"})
		return
	}

	var req struct {
		Enable bool `json:"enable"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	err := h.gameService.ToggleGame(ctx, gameID, req.Enable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"game_id": gameID, "enabled": req.Enable})
}

// SetGameOrder handles POST /games/:id/order (admin)
func (h *GameHandler) SetGameOrder(c *gin.Context) {
	gameID := c.Param("id")
	if gameID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "game id is required"})
		return
	}

	var req struct {
		SortOrder int `json:"sort_order"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	err := h.gameService.SetGameOrder(ctx, gameID, req.SortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"game_id": gameID, "sort_order": req.SortOrder})
}

// Helper function to parse int32
func parseInt32(s string, defaultValue int32) int32 {
	var result int32
	for _, c := range s {
		if c < '0' || c > '9' {
			return defaultValue
		}
		result = result*10 + int32(c-'0')
	}
	if result == 0 {
		return defaultValue
	}
	return result
}
