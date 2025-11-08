package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type TenagaKerjaRepository struct {
	db *gorm.DB
}

func NewTenagaKerjaRepository(db *gorm.DB) *TenagaKerjaRepository {
	return &TenagaKerjaRepository{db: db}
}

func (r *TenagaKerjaRepository) Create(tk *models.TenagaKerja) error {
	return r.db.Create(tk).Error
}

func (r *TenagaKerjaRepository) GetAll() ([]models.TenagaKerja, error) {
	var list []models.TenagaKerja
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *TenagaKerjaRepository) GetByID(id uint) (*models.TenagaKerja, error) {
	var tk models.TenagaKerja
	if err := r.db.First(&tk, id).Error; err != nil {
		return nil, err
	}
	return &tk, nil
}

func (r *TenagaKerjaRepository) GetByIDs(ids []uint) ([]models.TenagaKerja, error) {
	var list []models.TenagaKerja
	if err := r.db.Where("id IN ?", ids).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *TenagaKerjaRepository) Update(id uint, tk *models.TenagaKerja) error {
	return r.db.Model(&models.TenagaKerja{}).Where("id = ?", id).Updates(tk).Error
}

func (r *TenagaKerjaRepository) Delete(id uint) error {
	return r.db.Delete(&models.TenagaKerja{}, id).Error
}

func (r *TenagaKerjaRepository) ExistsByNIK(nik string) (bool, error) {
	var count int64
	err := r.db.Model(&models.TenagaKerja{}).Where("nik = ?", nik).Count(&count).Error
	return count > 0, err
}

