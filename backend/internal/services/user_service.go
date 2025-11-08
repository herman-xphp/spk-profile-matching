package services

import (
	"errors"

	"backend/internal/models"
	"backend/internal/repositories"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// Remove password from all users
	for i := range users {
		users[i].Password = ""
	}

	return users, nil
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func (s *UserService) Create(user *models.User) error {
	// Check if email already exists
	exists, err := s.userRepo.ExistsByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("could not hash password")
	}

	user.Password = string(hashedPassword)
	if user.Role == "" {
		user.Role = "user"
	}
	user.IsActive = true

	return s.userRepo.Create(user)
}

func (s *UserService) Update(id uint, user *models.User) error {
	// Check if user exists
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user not found")
		}
		return err
	}

	// Don't update password if empty
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("could not hash password")
		}
		user.Password = string(hashedPassword)
	} else {
		// Get existing user to preserve password
		existing, _ := s.userRepo.GetByID(id)
		user.Password = existing.Password
	}

	return s.userRepo.Update(id, user)
}

func (s *UserService) Delete(id uint) error {
	// Check if user exists
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("user not found")
		}
		return err
	}

	return s.userRepo.Delete(id)
}

func (s *UserService) Register(user *models.User) error {
	return s.Create(user)
}

