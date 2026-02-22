package services

import (
	"yuedsen-backend/models"
	"yuedsen-backend/repositories"
)

// ScoreboardEntry เป็น response ที่ frontend ต้องการ
type ScoreboardEntry struct {
	UserID         uint   `json:"user_id"`
	UserName       string `json:"user_name"`
	UserEmail      string `json:"user_email"`
	CategoryID     uint   `json:"category_id"`
	CategoryName   string `json:"category_name"`
	Progress       int    `json:"progress"`
	TotalScore     int    `json:"total_score"`
	Status         string `json:"status"`
}

type ScoreboardService interface {
	GetScoreboard() ([]ScoreboardEntry, error)
}

type scoreboardService struct {
	repo repositories.ScoreboardRepository
}

func NewScoreboardService(repo repositories.ScoreboardRepository) ScoreboardService {
	return &scoreboardService{repo: repo}
}

func (s *scoreboardService) GetScoreboard() ([]ScoreboardEntry, error) {
	processes, err := s.repo.GetAllUserProcesses()
	if err != nil {
		return nil, err
	}

	entries := make([]ScoreboardEntry, 0, len(processes))
	for _, p := range processes {
		entry := ScoreboardEntry{
			UserID:       p.UserID,
			UserName:     p.User.Name,
			UserEmail:    p.User.Email,
			CategoryID:   p.PoseCategoryID,
			CategoryName: p.PoseCategory.Category,
			Progress:     p.Progress,
			TotalScore:   p.TotalScore,
			Status:       p.Status,
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// UserScoreboardSummary สรุปคะแนนรวมทุก category per user
type UserScoreboardSummary struct {
	UserID    uint              `json:"user_id"`
	UserName  string            `json:"user_name"`
	UserEmail string            `json:"user_email"`
	Total     int               `json:"total_progress"`
	Details   []models.UserProcess `json:"details"`
}
