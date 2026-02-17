package repositories

import (
	"yuedsen-backend/models"

	"gorm.io/gorm"
)

type PoseRepository interface {
	Create(pose *models.Pose, landmarks []models.PoseLandmark) error
	FindAll() ([]models.Pose, error)
	FindByID(id uint) (*models.Pose, error)
	// Add other methods as needed: FindByCategory, etc.
}

type poseRepository struct {
	db *gorm.DB
}

func NewPoseRepository(db *gorm.DB) PoseRepository {
	return &poseRepository{db: db}
}

func (r *poseRepository) Create(pose *models.Pose, landmarks []models.PoseLandmark) error {
	// Transaction to ensure both Pose and Landmarks are saved
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(pose).Error; err != nil {
			return err
		}

		// Assign PoseID to landmarks
		for i := range landmarks {
			landmarks[i].PoseID = pose.ID
		}

		if len(landmarks) > 0 {
			if err := tx.Create(&landmarks).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *poseRepository) FindAll() ([]models.Pose, error) {
	var poses []models.Pose
	// Preload PoseCategory if needed
	result := r.db.Preload("PoseCategory").Find(&poses)
	return poses, result.Error
}

func (r *poseRepository) FindByID(id uint) (*models.Pose, error) {
	var pose models.Pose
	result := r.db.Preload("PoseCategory").First(&pose, id)
	return &pose, result.Error
}
