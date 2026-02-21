package controllers

import (
	"net/http"
	"yuedsen-backend/services"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	service services.GameService
}

func NewGameController(service services.GameService) *GameController {
	return &GameController{
		service: service,
	}
}

func (c *GameController) GetCurrentGamePlan(ctx *gin.Context) {
	// Extract user_id from context (injected by JWTMiddleware)
	userIDFloat, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: user_id missing"})
		return
	}

	userID := uint(userIDFloat.(float64)) // JSON parses numbers as float64

	plans, process, isBlocked, err := c.service.GetGamePlanForUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"process": process,
		"plans":   plans,
		"is_blocked": isBlocked,
	})
}

func (c *GameController) CompleteDay(ctx *gin.Context) {
	userIDFloat, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: user_id missing"})
		return
	}

	userID := uint(userIDFloat.(float64))

	err := c.service.CompleteDayForUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Day completed successfully"})
}
