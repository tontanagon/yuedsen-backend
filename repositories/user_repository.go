package repositories

import (
	"errors"
	"yuedsen-backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) error
	SaveUserToken(token *models.UserToken) error
	FindAllWithProcesses() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepository) FindAllWithProcesses() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Save(user *models.User) error {
	// FirstOrCreate: สร้างใหม่ถ้าไม่มี, ไม่ทำอะไรถ้ามีอยู่แล้ว (ป้องกัน duplicate email error)
	result := r.db.Where(models.User{Email: user.Email}).FirstOrCreate(user)
	return result.Error
}


func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // user not found → return nil (not an error)
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) SaveUserToken(token *models.UserToken) error {
	result := r.db.Create(token)
	return result.Error
}
