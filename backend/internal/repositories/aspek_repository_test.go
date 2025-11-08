package repositories

import (
	"testing"

	"backend/internal/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAspekRepository_Create(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewAspekRepository(db)

	aspek := &models.Aspek{
		Nama:       "Kompetensi",
		Deskripsi:  "Aspek Kompetensi",
		Persentase: 50.0,
	}

	err := repo.Create(aspek)
	assert.NoError(t, err)
	assert.NotZero(t, aspek.ID)
}

func TestAspekRepository_GetByID(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewAspekRepository(db)

	aspek := &models.Aspek{
		Nama:       "Kompetensi",
		Deskripsi:  "Aspek Kompetensi",
		Persentase: 50.0,
	}
	repo.Create(aspek)

	found, err := repo.GetByID(aspek.ID)
	assert.NoError(t, err)
	assert.Equal(t, aspek.Nama, found.Nama)
}

func TestAspekRepository_GetAll(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewAspekRepository(db)

	aspek1 := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspek2 := &models.Aspek{Nama: "Kepribadian", Persentase: 50.0}
	repo.Create(aspek1)
	repo.Create(aspek2)

	aspeks, err := repo.GetAll()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(aspeks), 2)
}

func TestAspekRepository_Update(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewAspekRepository(db)

	aspek := &models.Aspek{
		Nama:       "Kompetensi",
		Persentase: 50.0,
	}
	repo.Create(aspek)

	aspek.Nama = "Kompetensi Teknis"
	err := repo.Update(aspek.ID, aspek)
	assert.NoError(t, err)

	updated, _ := repo.GetByID(aspek.ID)
	assert.Equal(t, "Kompetensi Teknis", updated.Nama)
}

func TestAspekRepository_Delete(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewAspekRepository(db)

	aspek := &models.Aspek{
		Nama:       "Kompetensi",
		Persentase: 50.0,
	}
	repo.Create(aspek)

	err := repo.Delete(aspek.ID)
	assert.NoError(t, err)

	_, err = repo.GetByID(aspek.ID)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

