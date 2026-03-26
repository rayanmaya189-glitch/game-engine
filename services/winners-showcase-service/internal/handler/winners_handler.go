package handler

import (
	"net/http"

	"github.com/game_engine/winners-showcase-service/internal/model"
	"github.com/game_engine/winners-showcase-service/internal/service"
	"github.com/gin-gonic/gin"
)

type WinnersHandler struct {
	service *service.WinnersService
}

func NewWinnersHandler(s *service.WinnersService) *WinnersHandler {
	return &WinnersHandler{service: s}
}

func (h *WinnersHandler) GetRecentWinners(c *gin.Context) {
	resp, err := h.service.GetRecentWinners(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *WinnersHandler) GetBigWins(c *gin.Context) {
	resp, err := h.service.GetBigWins(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *WinnersHandler) GetJackpotWinners(c *gin.Context) {
	resp, err := h.service.GetJackpotWinners(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *WinnersHandler) RecordWin(c *gin.Context) {
	var req model.RecordWinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RecordWin(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Win recorded successfully"})
}

func (h *WinnersHandler) GetPrivacySettings(c *gin.Context) {
	userID := c.Param("userId")
	resp, err := h.service.GetPrivacySettings(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *WinnersHandler) UpdatePrivacySettings(c *gin.Context) {
	userID := c.Param("userId")
	var req model.UpdatePrivacyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdatePrivacySettings(c.Request.Context(), userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Privacy settings updated"})
}
