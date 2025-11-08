package services

import (
	"errors"

	"backend/internal/models"
	"backend/internal/repositories"

	"gorm.io/gorm"
)

type KriteriaService struct {
	kriteriaRepo *repositories.KriteriaRepository
	aspekRepo    *repositories.AspekRepository
}

func NewKriteriaService(kriteriaRepo *repositories.KriteriaRepository, aspekRepo *repositories.AspekRepository) *KriteriaService {
	return &KriteriaService{
		kriteriaRepo: kriteriaRepo,
		aspekRepo:    aspekRepo,
	}
}

func (s *KriteriaService) GetAll() ([]models.Kriteria, error) {
	return s.kriteriaRepo.GetAll()
}

func (s *KriteriaService) GetByID(id uint) (*models.Kriteria, error) {
	kriteria, err := s.kriteriaRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("kriteria not found")
		}
		return nil, err
	}
	return kriteria, nil
}

func (s *KriteriaService) Create(kriteria *models.Kriteria) error {
	if kriteria.Nama == "" {
		return errors.New("nama kriteria tidak boleh kosong")
	}
	if kriteria.Kode == "" {
		return errors.New("kode kriteria tidak boleh kosong")
	}

	// Validate aspek exists
	_, err := s.aspekRepo.GetByID(kriteria.AspekID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("aspek not found")
		}
		return err
	}

	return s.kriteriaRepo.Create(kriteria)
}

func (s *KriteriaService) Update(id uint, kriteria *models.Kriteria) error {
	// Check if kriteria exists
	_, err := s.kriteriaRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("kriteria not found")
		}
		return err
	}

	// Validate aspek exists if AspekID is being updated
	if kriteria.AspekID != 0 {
		_, err := s.aspekRepo.GetByID(kriteria.AspekID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("aspek not found")
			}
			return err
		}
	}

	return s.kriteriaRepo.Update(id, kriteria)
}

func (s *KriteriaService) Delete(id uint) error {
	// Check if kriteria exists
	_, err := s.kriteriaRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("kriteria not found")
		}
		return err
	}

	return s.kriteriaRepo.Delete(id)
}
