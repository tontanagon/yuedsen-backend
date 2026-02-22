package repositories

import (
	"yuedsen-backend/models"

	"gorm.io/gorm"
)

type ScoreboardRepository interface {
	GetAllUserProcesses() ([]models.UserProcess, error)
}

type scoreboardRepository struct {
	db *gorm.DB
}

func NewScoreboardRepository(db *gorm.DB) ScoreboardRepository {
	return &scoreboardRepository{db: db}
}

// GetAllUserProcesses ดึง UserProcess ทั้งหมด พร้อม preload User และ PoseCategory
func (r *scoreboardRepository) GetAllUserProcesses() ([]models.UserProcess, error) {
	var processes []models.UserProcess
	result := r.db.
		Preload("User").
		Preload("PoseCategory").
		Order("total_score desc, progress desc").
		Find(&processes)
	if result.Error != nil {
		return nil, result.Error
	}
	return processes, nil
}
