package services

import (
	"testing"

	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestAspekService_Create(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewAspekRepository(db)
	service := NewAspekService(repo)

	aspek := &models.Aspek{
		Nama:       "Kompetensi",
		Deskripsi:  "Aspek Kompetensi",
		Persentase: 50.0,
	}

	err := service.Create(aspek)
	assert.NoError(t, err)
	assert.NotZero(t, aspek.ID)
}

func TestAspekService_Create_EmptyName(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewAspekRepository(db)
	service := NewAspekService(repo)

	aspek := &models.Aspek{
		Nama:       "",
		Persentase: 50.0,
	}

	err := service.Create(aspek)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nama aspek tidak boleh kosong")
}

func TestAspekService_GetByID(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewAspekRepository(db)
	service := NewAspekService(repo)

	aspek := &models.Aspek{
		Nama:       "Kompetensi",
		Persentase: 50.0,
	}
	service.Create(aspek)

	found, err := service.GetByID(aspek.ID)
	assert.NoError(t, err)
	assert.Equal(t, aspek.Nama, found.Nama)
}

func TestAspekService_Update(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewAspekRepository(db)
	service := NewAspekService(repo)

	aspek := &models.Aspek{
		Nama:       "Kompetensi",
		Persentase: 50.0,
	}
	service.Create(aspek)

	aspek.Nama = "Kompetensi Teknis"
	err := service.Update(aspek.ID, aspek)
	assert.NoError(t, err)

	updated, _ := service.GetByID(aspek.ID)
	assert.Equal(t, "Kompetensi Teknis", updated.Nama)
}

func TestAspekService_Delete(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewAspekRepository(db)
	service := NewAspekService(repo)

	aspek := &models.Aspek{
		Nama:       "Kompetensi",
		Persentase: 50.0,
	}
	service.Create(aspek)

	err := service.Delete(aspek.ID)
	assert.NoError(t, err)

	_, err = service.GetByID(aspek.ID)
	assert.Error(t, err)
}
