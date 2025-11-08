package services

import (
	"errors"

	"backend/internal/models"
	"backend/internal/repositories"

	"gorm.io/gorm"
)

type TenagaKerjaService struct {
	tenagaKerjaRepo *repositories.TenagaKerjaRepository
}

func NewTenagaKerjaService(tenagaKerjaRepo *repositories.TenagaKerjaRepository) *TenagaKerjaService {
	return &TenagaKerjaService{tenagaKerjaRepo: tenagaKerjaRepo}
}

func (s *TenagaKerjaService) GetAll() ([]models.TenagaKerja, error) {
	return s.tenagaKerjaRepo.GetAll()
}

func (s *TenagaKerjaService) GetByID(id uint) (*models.TenagaKerja, error) {
	tenagaKerja, err := s.tenagaKerjaRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("tenaga kerja not found")
		}
		return nil, err
	}
	return tenagaKerja, nil
}

func (s *TenagaKerjaService) Create(tenagaKerja *models.TenagaKerja) error {
	if tenagaKerja.NIK == "" {
		return errors.New("NIK tidak boleh kosong")
	}
	if tenagaKerja.Nama == "" {
		return errors.New("nama tidak boleh kosong")
	}

	// Check if NIK already exists
	exists, err := s.tenagaKerjaRepo.ExistsByNIK(tenagaKerja.NIK)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("NIK sudah terdaftar")
	}

	return s.tenagaKerjaRepo.Create(tenagaKerja)
}

func (s *TenagaKerjaService) Update(id uint, tenagaKerja *models.TenagaKerja) error {
	// Check if tenaga kerja exists
	existing, err := s.tenagaKerjaRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("tenaga kerja not found")
		}
		return err
	}

	// Check if NIK is being changed and if new NIK already exists
	if tenagaKerja.NIK != "" && tenagaKerja.NIK != existing.NIK {
		exists, err := s.tenagaKerjaRepo.ExistsByNIK(tenagaKerja.NIK)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("NIK sudah terdaftar")
		}
	}

	return s.tenagaKerjaRepo.Update(id, tenagaKerja)
}

func (s *TenagaKerjaService) Delete(id uint) error {
	// Check if tenaga kerja exists
	_, err := s.tenagaKerjaRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("tenaga kerja not found")
		}
		return err
	}

	return s.tenagaKerjaRepo.Delete(id)
}

