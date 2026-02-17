package services

import (
	"yuedsen-backend/models"
	"yuedsen-backend/repositories"
)

type PoseService interface {
	CreatePose(pose *models.Pose, landmarks []models.PoseLandmark) error
	GetAllPoses() ([]models.Pose, error)
	GetPoseByID(id uint) (*models.Pose, error)
}

type poseService struct {
	repo repositories.PoseRepository
}

func NewPoseService(repo repositories.PoseRepository) PoseService {
	return &poseService{repo: repo}
}

func (s *poseService) CreatePose(pose *models.Pose, landmarks []models.PoseLandmark) error {
	return s.repo.Create(pose, landmarks)
}

func (s *poseService) GetAllPoses() ([]models.Pose, error) {
	return s.repo.FindAll()
}

func (s *poseService) GetPoseByID(id uint) (*models.Pose, error) {
	return s.repo.FindByID(id)
}
