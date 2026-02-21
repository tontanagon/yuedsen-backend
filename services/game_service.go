package services

import (
	"errors"
	"time"
	"yuedsen-backend/models"
	"yuedsen-backend/repositories"
)

type GameService interface {
	GetGamePlanForUser(userID uint) ([]models.Plan, *models.UserProcess, bool, error)
	CompleteDayForUser(userID uint) error
}

type gameService struct {
	repo repositories.GameRepository
}

func NewGameService(repo repositories.GameRepository) GameService {
	return &gameService{
		repo: repo,
	}
}

func (s *gameService) GetGamePlanForUser(userID uint) ([]models.Plan, *models.UserProcess, bool, error) {
	process, err := s.repo.GetUserProcess(userID)
	if err != nil {
		return nil, nil, false, err
	}

	// If user has no active process, create one
	if process == nil {
		newProcess := &models.UserProcess{
			UserID:         userID,
			PoseCategoryID: 1, // Defaulting to category 1
			Progress:       1, // Start at day 1
			Status:         "in_progress",
			TotalScore:     0,
		}
		if err := s.repo.CreateUserProcess(newProcess); err != nil {
			return nil, nil, false, errors.New("failed to initialize user progress")
		}
		process = newProcess
	}

	// Check if blocked: User completed a day today and their progress is > 1
	isBlocked := process.UpdatedAt.Format("2006-01-02") == time.Now().Format("2006-01-02") && process.Progress > 1
	
	// Check to ensure we are not blocking if they literally just registered right now.
	// if CreatedAt == UpdatedAt, we shouldn't block, but checking Progress > 1 already handles this!

	// Fetch plans logic based on process.Progress (which represents the Day)
	plans, err := s.repo.GetPlansByDay(process.Progress)
	if err != nil {
		return nil, nil, false, errors.New("failed to fetch game plan")
	}

	return plans, process, isBlocked, nil
}

func (s *gameService) CompleteDayForUser(userID uint) error {
	process, err := s.repo.GetUserProcess(userID)
	if err != nil {
		return err
	}
	if process == nil {
		return errors.New("user process not found")
	}

	// Check if already updated today to prevent double submission
	if process.UpdatedAt.Format("2006-01-02") == time.Now().Format("2006-01-02") && process.Progress > 1 {
		return errors.New("user has already completed a process today")
	}

	process.Progress += 1
	return s.repo.UpdateUserProcess(process)
}
