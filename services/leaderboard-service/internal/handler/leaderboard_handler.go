package handler

import (
	"net/http"

	"github.com/game_engine/leaderboard-service/internal/model"
	"github.com/game_engine/leaderboard-service/internal/service"
	"github.com/gin-gonic/gin"
)

type LeaderboardHandler struct {
	service *service.LeaderboardService
}

func NewLeaderboardHandler(s *service.LeaderboardService) *LeaderboardHandler {
	return &LeaderboardHandler{service: s}
}

func (h *LeaderboardHandler) GetDailyLeaderboard(c *gin.Context) {
	gameType := c.Query("game_type")
	resp, err := h.service.GetDailyLeaderboard(c.Request.Context(), gameType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetDailyLeaderboardByGame(c *gin.Context) {
	gameType := c.Param("gameType")
	resp, err := h.service.GetDailyLeaderboard(c.Request.Context(), gameType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetWeeklyLeaderboard(c *gin.Context) {
	gameType := c.Query("game_type")
	resp, err := h.service.GetWeeklyLeaderboard(c.Request.Context(), gameType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetWeeklyLeaderboardByGame(c *gin.Context) {
	gameType := c.Param("gameType")
	resp, err := h.service.GetWeeklyLeaderboard(c.Request.Context(), gameType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetMonthlyLeaderboard(c *gin.Context) {
	gameType := c.Query("game_type")
	resp, err := h.service.GetMonthlyLeaderboard(c.Request.Context(), gameType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetMonthlyLeaderboardByGame(c *gin.Context) {
	gameType := c.Param("gameType")
	resp, err := h.service.GetMonthlyLeaderboard(c.Request.Context(), gameType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetAllTimeLeaderboard(c *gin.Context) {
	gameType := c.Query("game_type")
	resp, err := h.service.GetAllTimeLeaderboard(c.Request.Context(), gameType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetAllTimeLeaderboardByGame(c *gin.Context) {
	gameType := c.Param("gameType")
	resp, err := h.service.GetAllTimeLeaderboard(c.Request.Context(), gameType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetPlayerDailyRank(c *gin.Context) {
	userID := c.Param("userId")
	gameType := c.Query("game_type")
	resp, err := h.service.GetPlayerDailyRank(c.Request.Context(), userID, gameType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found on leaderboard"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetPlayerWeeklyRank(c *gin.Context) {
	userID := c.Param("userId")
	gameType := c.Query("game_type")
	resp, err := h.service.GetPlayerWeeklyRank(c.Request.Context(), userID, gameType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found on leaderboard"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetPlayerMonthlyRank(c *gin.Context) {
	userID := c.Param("userId")
	gameType := c.Query("game_type")
	resp, err := h.service.GetPlayerMonthlyRank(c.Request.Context(), userID, gameType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found on leaderboard"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) GetPlayerAllTimeRank(c *gin.Context) {
	userID := c.Param("userId")
	gameType := c.Query("game_type")
	resp, err := h.service.GetPlayerAllTimeRank(c.Request.Context(), userID, gameType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found on leaderboard"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *LeaderboardHandler) UpdatePlayerScore(c *gin.Context) {
	var req model.UpdateScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdatePlayerScore(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Score updated successfully"})
}

func (h *LeaderboardHandler) DistributePrizes(c *gin.Context) {
	var req model.PrizeDistributionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.DistributePrizes(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *LeaderboardHandler) SyncLeaderboard(c *gin.Context) {
	leaderboardType := c.Param("type")
	gameType := c.Query("game_type")

	err := h.service.SyncLeaderboard(c.Request.Context(), model.LeaderboardType(leaderboardType), gameType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leaderboard synced successfully"})
}

func (h *LeaderboardHandler) ResetLeaderboard(c *gin.Context) {
	leaderboardType := c.Param("type")
	gameType := c.Query("game_type")

	err := h.service.ResetLeaderboard(c.Request.Context(), model.LeaderboardType(leaderboardType), gameType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leaderboard reset successfully"})
}
