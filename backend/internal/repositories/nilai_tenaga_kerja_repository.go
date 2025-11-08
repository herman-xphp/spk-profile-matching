package repositories

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type NilaiTenagaKerjaRepository struct {
	db *gorm.DB
}

func NewNilaiTenagaKerjaRepository(db *gorm.DB) *NilaiTenagaKerjaRepository {
	return &NilaiTenagaKerjaRepository{db: db}
}

func (r *NilaiTenagaKerjaRepository) Create(ntk *models.NilaiTenagaKerja) error {
	return r.db.Create(ntk).Error
}

func (r *NilaiTenagaKerjaRepository) GetAll() ([]models.NilaiTenagaKerja, error) {
	var list []models.NilaiTenagaKerja
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *NilaiTenagaKerjaRepository) GetByID(id uint) (*models.NilaiTenagaKerja, error) {
	var ntk models.NilaiTenagaKerja
	if err := r.db.First(&ntk, id).Error; err != nil {
		return nil, err
	}
	return &ntk, nil
}

func (r *NilaiTenagaKerjaRepository) GetByTenagaKerjaID(tenagaKerjaID uint) ([]models.NilaiTenagaKerja, error) {
	var list []models.NilaiTenagaKerja
	if err := r.db.Where("tenaga_kerja_id = ?", tenagaKerjaID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *NilaiTenagaKerjaRepository) GetByKriteriaID(kriteriaID uint) ([]models.NilaiTenagaKerja, error) {
	var list []models.NilaiTenagaKerja
	if err := r.db.Where("kriteria_id = ?", kriteriaID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *NilaiTenagaKerjaRepository) Update(id uint, ntk *models.NilaiTenagaKerja) error {
	return r.db.Model(&models.NilaiTenagaKerja{}).Where("id = ?", id).Updates(ntk).Error
}

func (r *NilaiTenagaKerjaRepository) Delete(id uint) error {
	return r.db.Delete(&models.NilaiTenagaKerja{}, id).Error
}

