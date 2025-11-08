package services

import (
	"errors"

	"backend/internal/models"
	"backend/internal/repositories"

	"gorm.io/gorm"
)

type NilaiTenagaKerjaService struct {
	nilaiTenagaKerjaRepo *repositories.NilaiTenagaKerjaRepository
	tenagaKerjaRepo      *repositories.TenagaKerjaRepository
	kriteriaRepo         *repositories.KriteriaRepository
}

func NewNilaiTenagaKerjaService(
	nilaiTenagaKerjaRepo *repositories.NilaiTenagaKerjaRepository,
	tenagaKerjaRepo *repositories.TenagaKerjaRepository,
	kriteriaRepo *repositories.KriteriaRepository,
) *NilaiTenagaKerjaService {
	return &NilaiTenagaKerjaService{
		nilaiTenagaKerjaRepo: nilaiTenagaKerjaRepo,
		tenagaKerjaRepo:      tenagaKerjaRepo,
		kriteriaRepo:         kriteriaRepo,
	}
}

func (s *NilaiTenagaKerjaService) GetAll() ([]models.NilaiTenagaKerja, error) {
	return s.nilaiTenagaKerjaRepo.GetAll()
}

func (s *NilaiTenagaKerjaService) GetByID(id uint) (*models.NilaiTenagaKerja, error) {
	nilai, err := s.nilaiTenagaKerjaRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("nilai tenaga kerja not found")
		}
		return nil, err
	}
	return nilai, nil
}

func (s *NilaiTenagaKerjaService) GetByTenagaKerjaID(tenagaKerjaID uint) ([]models.NilaiTenagaKerja, error) {
	return s.nilaiTenagaKerjaRepo.GetByTenagaKerjaID(tenagaKerjaID)
}

func (s *NilaiTenagaKerjaService) Create(nilai *models.NilaiTenagaKerja) error {
	// Validate tenaga kerja exists
	_, err := s.tenagaKerjaRepo.GetByID(nilai.TenagaKerjaID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("tenaga kerja not found")
		}
		return err
	}

	// Validate kriteria exists
	_, err = s.kriteriaRepo.GetByID(nilai.KriteriaID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("kriteria not found")
		}
		return err
	}

	return s.nilaiTenagaKerjaRepo.Create(nilai)
}

func (s *NilaiTenagaKerjaService) Update(id uint, nilai *models.NilaiTenagaKerja) error {
	// Check if nilai exists
	_, err := s.nilaiTenagaKerjaRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("nilai tenaga kerja not found")
		}
		return err
	}

	// Validate tenaga kerja exists if TenagaKerjaID is being updated
	if nilai.TenagaKerjaID != 0 {
		_, err := s.tenagaKerjaRepo.GetByID(nilai.TenagaKerjaID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("tenaga kerja not found")
			}
			return err
		}
	}

	// Validate kriteria exists if KriteriaID is being updated
	if nilai.KriteriaID != 0 {
		_, err := s.kriteriaRepo.GetByID(nilai.KriteriaID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("kriteria not found")
			}
			return err
		}
	}

	return s.nilaiTenagaKerjaRepo.Update(id, nilai)
}

func (s *NilaiTenagaKerjaService) Delete(id uint) error {
	// Check if nilai exists
	_, err := s.nilaiTenagaKerjaRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("nilai tenaga kerja not found")
		}
		return err
	}

	return s.nilaiTenagaKerjaRepo.Delete(id)
}

