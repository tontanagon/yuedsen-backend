package repositories

import (
	"errors"
	"yuedsen-backend/models"

	"gorm.io/gorm"
)

type GameRepository interface {
	GetUserProcess(userID uint, categoryID uint) (*models.UserProcess, error)
	CreateUserProcess(process *models.UserProcess) error
	UpdateUserProcess(process *models.UserProcess) error
	GetPlansByDay(day int, categoryID uint) ([]models.Plan, error)
}

type gameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) GameRepository {
	return &gameRepository{db: db}
}

func (r *gameRepository) GetUserProcess(userID uint, categoryID uint) (*models.UserProcess, error) {
	var process models.UserProcess
	// Get the active process for this specific category
	result := r.db.Where("user_id = ? AND pose_category_id = ?", userID, categoryID).Order("updated_at desc").First(&process)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // not found
		}
		return nil, result.Error
	}
	return &process, nil
}

func (r *gameRepository) CreateUserProcess(process *models.UserProcess) error {
	result := r.db.Create(process)
	return result.Error
}

func (r *gameRepository) UpdateUserProcess(process *models.UserProcess) error {
	result := r.db.Save(process)
	return result.Error
}

func (r *gameRepository) GetPlansByDay(day int, categoryID uint) ([]models.Plan, error) {
	var plans []models.Plan
	// Preload "Pose", and you might also want to Preload "Pose.PoseCategory" if needed
	result := r.db.Preload("Pose").Where("day = ? AND pose_category_id = ?", day, categoryID).Find(&plans)
	if result.Error != nil {
		return nil, result.Error
	}
	return plans, nil
}
