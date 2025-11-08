package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type AspekRepository struct {
	db *gorm.DB
}

func NewAspekRepository(db *gorm.DB) *AspekRepository {
	return &AspekRepository{db: db}
}

func (r *AspekRepository) Create(a *models.Aspek) error {
	return r.db.Create(a).Error
}

func (r *AspekRepository) GetAll() ([]models.Aspek, error) {
	var list []models.Aspek
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *AspekRepository) GetByID(id uint) (*models.Aspek, error) {
	var a models.Aspek
	if err := r.db.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AspekRepository) Update(id uint, a *models.Aspek) error {
	return r.db.Model(&models.Aspek{}).Where("id = ?", id).Updates(a).Error
}

func (r *AspekRepository) Delete(id uint) error {
	return r.db.Delete(&models.Aspek{}, id).Error
}

