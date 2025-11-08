package services

import (
	"errors"

	"backend/internal/models"
	"backend/internal/repositories"

	"gorm.io/gorm"
)

type AspekService struct {
	aspekRepo *repositories.AspekRepository
}

func NewAspekService(aspekRepo *repositories.AspekRepository) *AspekService {
	return &AspekService{aspekRepo: aspekRepo}
}

func (s *AspekService) GetAll() ([]models.Aspek, error) {
	return s.aspekRepo.GetAll()
}

func (s *AspekService) GetByID(id uint) (*models.Aspek, error) {
	aspek, err := s.aspekRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("aspek not found")
		}
		return nil, err
	}
	return aspek, nil
}

func (s *AspekService) Create(aspek *models.Aspek) error {
	if aspek.Nama == "" {
		return errors.New("nama aspek tidak boleh kosong")
	}
	return s.aspekRepo.Create(aspek)
}

func (s *AspekService) Update(id uint, aspek *models.Aspek) error {
	// Check if aspek exists
	_, err := s.aspekRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("aspek not found")
		}
		return err
	}

	return s.aspekRepo.Update(id, aspek)
}

func (s *AspekService) Delete(id uint) error {
	// Check if aspek exists
	_, err := s.aspekRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("aspek not found")
		}
		return err
	}

	return s.aspekRepo.Delete(id)
}

