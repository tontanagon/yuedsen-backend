package controllers

import (
	"net/http"
	"yuedsen-backend/services"

	"github.com/gin-gonic/gin"
)

type ScoreboardController struct {
	service services.ScoreboardService
}

func NewScoreboardController(service services.ScoreboardService) *ScoreboardController {
	return &ScoreboardController{service: service}
}

// GetScoreboard ดึงข้อมูล scoreboard ทั้งหมด
func (c *ScoreboardController) GetScoreboard(ctx *gin.Context) {
	entries, err := c.service.GetScoreboard()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"scoreboard": entries})
}
