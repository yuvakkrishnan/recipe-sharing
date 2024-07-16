package services

import (
	"github.com/yuvakkrishnan/user-service/internal/models"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	Login(username, password string) (string, error)
	GetProfile(userID int64) (*models.User, error)
	ForgotPassword(email string) error
	ResetPassword(token, newPassword string) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) Register(user *models.User) error {
	// Check if the user already exists
	existingUser, err := s.GetUserByEmail(user.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	// If user does not exist, create a new record; otherwise, update the existing one
	if existingUser.ID == 0 {
		if err := s.db.Create(user).Error; err != nil {
			return err
		}
	} else {
		if err := s.db.Model(existingUser).Updates(user).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (s *userService) Login(username, password string) (string, error) {
	// Implement login logic
	return "", nil
}

func (s *userService) GetProfile(userID int64) (*models.User, error) {
	// Implement get profile logic
	return nil, nil
}

func (s *userService) ForgotPassword(email string) error {
	// Implement forgot password logic
	return nil
}

func (s *userService) ResetPassword(token, newPassword string) error {
	// Implement reset password logic
	return nil
}
