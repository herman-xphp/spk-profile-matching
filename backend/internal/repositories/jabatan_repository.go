package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type JabatanRepository struct {
	db *gorm.DB
}

func NewJabatanRepository(db *gorm.DB) *JabatanRepository {
	return &JabatanRepository{db: db}
}

func (r *JabatanRepository) Create(j *models.Jabatan) error {
	return r.db.Create(j).Error
}

func (r *JabatanRepository) GetAll() ([]models.Jabatan, error) {
	var list []models.Jabatan
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *JabatanRepository) GetByID(id uint) (*models.Jabatan, error) {
	var j models.Jabatan
	if err := r.db.First(&j, id).Error; err != nil {
		return nil, err
	}
	return &j, nil
}

func (r *JabatanRepository) Update(id uint, j *models.Jabatan) error {
	return r.db.Model(&models.Jabatan{}).Where("id = ?", id).Updates(j).Error
}

func (r *JabatanRepository) Delete(id uint) error {
	return r.db.Delete(&models.Jabatan{}, id).Error
}
