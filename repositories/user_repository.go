package repositories

import (
	"errors"
	"yuedsen-backend/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	Save(user *models.User) error
}

type userRepository struct {
	users []models.User // Mock database
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: []models.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
		},
	}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	return r.users, nil
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *userRepository) Save(user *models.User) error {
	user.ID = uint(len(r.users) + 1)
	r.users = append(r.users, *user)
	return nil
}
