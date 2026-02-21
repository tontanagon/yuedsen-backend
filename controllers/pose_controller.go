package controllers

import (
	"net/http"
	
	"yuedsen-backend/models"
	"yuedsen-backend/services"
	
	"github.com/gin-gonic/gin"
)

type PoseController struct {
	service services.PoseService
}

func NewPoseController(service services.PoseService) *PoseController {
	return &PoseController{service: service}
}

type CreatePoseRequest struct {
	PoseName        string  `json:"pose_name" binding:"required"`
	PoseDescription string  `json:"pose_description"`
	PoseCondition   string  `json:"pose_condition"`
	
	// Rule-based checking (Optional)
	PosePoint       string  `json:"pose_point"` // [11, 13, 15]

	Status          string  `json:"status"`
	
	// Landmarks for Similarity Checking (Optional)
	Landmarks       []struct {
		LandmarkIndex int     `json:"landmark_index"`
		X             float64 `json:"x"`
		Y             float64 `json:"y"`
		Z             float64 `json:"z"`
		Visibility    float64 `json:"visibility"`
	} `json:"landmarks"`
}

func (c *PoseController) CreatePose(ctx *gin.Context) {
	var req CreatePoseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Map Request to Pose Model
	pose := models.Pose{
		PoseName:        req.PoseName,
		PoseDescription: req.PoseDescription,
		PoseCondition:   req.PoseCondition,
		PosePoint:       req.PosePoint,
		Status:          req.Status,
	}

	// 2. Map Request Landmarks to PoseLandmark Models
	var poseLandmarks []models.PoseLandmark
	for _, l := range req.Landmarks {
		poseLandmarks = append(poseLandmarks, models.PoseLandmark{
			LandmarkIndex: l.LandmarkIndex,
			X:             l.X,
			Y:             l.Y,
			Z:             l.Z,
			Visibility:    l.Visibility,
		})
	}

	// 3. Save via Service
	if err := c.service.CreatePose(&pose, poseLandmarks); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Pose created successfully",
		"pose_id": pose.ID,
	})
}

func (c *PoseController) GetPoses(ctx *gin.Context) {
	poses, err := c.service.GetAllPoses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, poses)
}
