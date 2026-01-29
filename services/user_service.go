package services

import (
	"errors"
	"yuedsen-backend/models"
	"yuedsen-backend/repositories"
)

// UserService defines the methods that our service should implement
type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
}

// userService is the concrete implementation of UserService
type userService struct {
	repo repositories.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Save(user)
}
