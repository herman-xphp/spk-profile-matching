package services

import (
	"errors"

	"backend/internal/models"
	"backend/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	users *repositories.UserRepository
}

func NewAuthService(u *repositories.UserRepository) *AuthService {
	return &AuthService{users: u}
}

// Authenticate checks email/password and returns user (without password) or error
func (s *AuthService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.users.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	// hide password
	user.Password = ""
	return user, nil
}
