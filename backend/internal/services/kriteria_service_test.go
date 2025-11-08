package services

import (
	"testing"

	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestKriteriaService_Create(t *testing.T) {
	db := setupServiceTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	service := NewKriteriaService(kriteriaRepo, aspekRepo)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria := &models.Kriteria{
		AspekID: aspek.ID,
		Kode:    "K1",
		Nama:    "Kriteria 1",
		IsCore:  true,
		Bobot:   1.0,
	}

	err := service.Create(kriteria)
	assert.NoError(t, err)
	assert.NotZero(t, kriteria.ID)
}

func TestKriteriaService_Create_InvalidAspek(t *testing.T) {
	db := setupServiceTestDB(t)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	service := NewKriteriaService(kriteriaRepo, aspekRepo)

	kriteria := &models.Kriteria{
		AspekID: 999, // Non-existent aspek
		Kode:    "K1",
		Nama:    "Kriteria 1",
		IsCore:  true,
		Bobot:   1.0,
	}

	err := service.Create(kriteria)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "aspek not found")
}

func TestKriteriaService_Create_EmptyName(t *testing.T) {
	db := setupServiceTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	service := NewKriteriaService(kriteriaRepo, aspekRepo)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria := &models.Kriteria{
		AspekID: aspek.ID,
		Kode:    "K1",
		Nama:    "",
		IsCore:  true,
		Bobot:   1.0,
	}

	err := service.Create(kriteria)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nama kriteria tidak boleh kosong")
}

func TestKriteriaService_GetByID(t *testing.T) {
	db := setupServiceTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	service := NewKriteriaService(kriteriaRepo, aspekRepo)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria := &models.Kriteria{
		AspekID: aspek.ID,
		Kode:    "K1",
		Nama:    "Kriteria 1",
		IsCore:  true,
		Bobot:   1.0,
	}
	service.Create(kriteria)

	found, err := service.GetByID(kriteria.ID)
	assert.NoError(t, err)
	assert.Equal(t, kriteria.Nama, found.Nama)
}

func TestKriteriaService_Update(t *testing.T) {
	db := setupServiceTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	service := NewKriteriaService(kriteriaRepo, aspekRepo)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria := &models.Kriteria{
		AspekID: aspek.ID,
		Kode:    "K1",
		Nama:    "Kriteria 1",
		IsCore:  true,
		Bobot:   1.0,
	}
	service.Create(kriteria)

	kriteria.Nama = "Updated Kriteria"
	err := service.Update(kriteria.ID, kriteria)
	assert.NoError(t, err)

	updated, _ := service.GetByID(kriteria.ID)
	assert.Equal(t, "Updated Kriteria", updated.Nama)
}

func TestKriteriaService_Delete(t *testing.T) {
	db := setupServiceTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	service := NewKriteriaService(kriteriaRepo, aspekRepo)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria := &models.Kriteria{
		AspekID: aspek.ID,
		Kode:    "K1",
		Nama:    "Kriteria 1",
		IsCore:  true,
		Bobot:   1.0,
	}
	service.Create(kriteria)

	err := service.Delete(kriteria.ID)
	assert.NoError(t, err)

	_, err = service.GetByID(kriteria.ID)
	assert.Error(t, err)
}

