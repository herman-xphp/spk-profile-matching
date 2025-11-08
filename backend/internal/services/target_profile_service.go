package services

import (
	"errors"

	"backend/internal/models"
	"backend/internal/repositories"

	"gorm.io/gorm"
)

type TargetProfileService struct {
	targetProfileRepo *repositories.TargetProfileRepository
	jabatanRepo       *repositories.JabatanRepository
	kriteriaRepo      *repositories.KriteriaRepository
}

func NewTargetProfileService(
	targetProfileRepo *repositories.TargetProfileRepository,
	jabatanRepo *repositories.JabatanRepository,
	kriteriaRepo *repositories.KriteriaRepository,
) *TargetProfileService {
	return &TargetProfileService{
		targetProfileRepo: targetProfileRepo,
		jabatanRepo:       jabatanRepo,
		kriteriaRepo:      kriteriaRepo,
	}
}

func (s *TargetProfileService) GetAll() ([]models.TargetProfile, error) {
	return s.targetProfileRepo.GetAll()
}

func (s *TargetProfileService) GetByID(id uint) (*models.TargetProfile, error) {
	profile, err := s.targetProfileRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("target profile not found")
		}
		return nil, err
	}
	return profile, nil
}

func (s *TargetProfileService) GetByJabatanID(jabatanID uint) ([]models.TargetProfile, error) {
	return s.targetProfileRepo.GetByJabatanID(jabatanID)
}

func (s *TargetProfileService) Create(profile *models.TargetProfile) error {
	// Validate jabatan exists
	_, err := s.jabatanRepo.GetByID(profile.JabatanID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("jabatan not found")
		}
		return err
	}

	// Validate kriteria exists
	_, err = s.kriteriaRepo.GetByID(profile.KriteriaID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("kriteria not found")
		}
		return err
	}

	return s.targetProfileRepo.Create(profile)
}

func (s *TargetProfileService) Update(id uint, profile *models.TargetProfile) error {
	// Check if profile exists
	_, err := s.targetProfileRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("target profile not found")
		}
		return err
	}

	// Validate jabatan exists if JabatanID is being updated
	if profile.JabatanID != 0 {
		_, err := s.jabatanRepo.GetByID(profile.JabatanID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("jabatan not found")
			}
			return err
		}
	}

	// Validate kriteria exists if KriteriaID is being updated
	if profile.KriteriaID != 0 {
		_, err := s.kriteriaRepo.GetByID(profile.KriteriaID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return errors.New("kriteria not found")
			}
			return err
		}
	}

	return s.targetProfileRepo.Update(id, profile)
}

func (s *TargetProfileService) Delete(id uint) error {
	// Check if profile exists
	_, err := s.targetProfileRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("target profile not found")
		}
		return err
	}

	return s.targetProfileRepo.Delete(id)
}

