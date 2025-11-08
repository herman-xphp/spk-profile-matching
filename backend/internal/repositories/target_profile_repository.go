package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type TargetProfileRepository struct {
	db *gorm.DB
}

func NewTargetProfileRepository(db *gorm.DB) *TargetProfileRepository {
	return &TargetProfileRepository{db: db}
}

func (r *TargetProfileRepository) Create(tp *models.TargetProfile) error {
	return r.db.Create(tp).Error
}

func (r *TargetProfileRepository) GetAll() ([]models.TargetProfile, error) {
	var list []models.TargetProfile
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *TargetProfileRepository) GetByID(id uint) (*models.TargetProfile, error) {
	var tp models.TargetProfile
	if err := r.db.First(&tp, id).Error; err != nil {
		return nil, err
	}
	return &tp, nil
}

func (r *TargetProfileRepository) GetByJabatanID(jabatanID uint) ([]models.TargetProfile, error) {
	var list []models.TargetProfile
	if err := r.db.Where("jabatan_id = ?", jabatanID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *TargetProfileRepository) Update(id uint, tp *models.TargetProfile) error {
	return r.db.Model(&models.TargetProfile{}).Where("id = ?", id).Updates(tp).Error
}

func (r *TargetProfileRepository) Delete(id uint) error {
	return r.db.Delete(&models.TargetProfile{}, id).Error
}

