package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"yuedsen-backend/models"
	"yuedsen-backend/repositories"

	"github.com/dgrijalva/jwt-go"
)

// UserService defines the methods that our service should implement
type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	GoogleCallback(accessToken string) (string, error)
	GetTokenByEmail(email string) (string, error) // สร้าง JWT จาก email
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

func (s *userService) GoogleCallback(accessToken string) (string, error) {
	// 1. Get User Info from Google
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return "", errors.New("failed to get user info from google")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("google api returned status: %d", resp.StatusCode)
	}

	var googleUser models.GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return "", errors.New("failed to decode google user info")
	}

	// 2. Check if user exists
	user, err := s.repo.FindByEmail(googleUser.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		// Create new user
		newUser := &models.User{
			Name:     googleUser.Name,
			Email:    googleUser.Email,
			Password: "", // No password for OAuth users
		}
		if err := s.repo.Save(newUser); err != nil {
			return "", err
		}
		user = newUser
	}

	// 3. Save User Token (Optional, as requested)
	userToken := &models.UserToken{
		UserID:    user.ID,
		Token:     accessToken,
		Provider:  "google",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(1 * time.Hour), // Example expiry
	}
	if err := s.repo.SaveUserToken(userToken); err != nil {
		// Log error but maybe don't fail the login? Or fail it.
		// For now, let's return error
		return "", err
	}

	// 4. Generate JWT Token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	// TODO: Move secret to env
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "fallback_secret_for_dev"
	}
	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("failed to generate jwt token")
	}

	return tokenString, nil
}

func (s *userService) GetTokenByEmail(email string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	// Generate JWT with user_id, email, name
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"name":    user.Name,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "fallback_secret_for_dev"
	}
	tokenString, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("failed to generate jwt token")
	}
	return tokenString, nil
}
