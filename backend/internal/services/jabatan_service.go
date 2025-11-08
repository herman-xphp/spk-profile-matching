package services

import (
	"errors"

	"backend/internal/models"
	"backend/internal/repositories"

	"gorm.io/gorm"
)

type JabatanService struct {
	jabatanRepo *repositories.JabatanRepository
}

func NewJabatanService(jabatanRepo *repositories.JabatanRepository) *JabatanService {
	return &JabatanService{jabatanRepo: jabatanRepo}
}

func (s *JabatanService) GetAll() ([]models.Jabatan, error) {
	return s.jabatanRepo.GetAll()
}

func (s *JabatanService) GetByID(id uint) (*models.Jabatan, error) {
	jabatan, err := s.jabatanRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("jabatan not found")
		}
		return nil, err
	}
	return jabatan, nil
}

func (s *JabatanService) Create(jabatan *models.Jabatan) error {
	if jabatan.Nama == "" {
		return errors.New("nama jabatan tidak boleh kosong")
	}
	return s.jabatanRepo.Create(jabatan)
}

func (s *JabatanService) Update(id uint, jabatan *models.Jabatan) error {
	// Check if jabatan exists
	_, err := s.jabatanRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("jabatan not found")
		}
		return err
	}

	return s.jabatanRepo.Update(id, jabatan)
}

func (s *JabatanService) Delete(id uint) error {
	// Check if jabatan exists
	_, err := s.jabatanRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("jabatan not found")
		}
		return err
	}

	return s.jabatanRepo.Delete(id)
}

