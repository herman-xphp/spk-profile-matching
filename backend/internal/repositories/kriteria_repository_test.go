package repositories

import (
	"testing"

	"backend/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestKriteriaRepository_Create(t *testing.T) {
	db := setupRepositoryTestDB(t)
	aspekRepo := NewAspekRepository(db)
	kriteriaRepo := NewKriteriaRepository(db)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria := &models.Kriteria{
		AspekID: aspek.ID,
		Kode:    "K1",
		Nama:    "Kriteria 1",
		IsCore:  true,
		Bobot:   1.0,
	}

	err := kriteriaRepo.Create(kriteria)
	assert.NoError(t, err)
	assert.NotZero(t, kriteria.ID)
}

func TestKriteriaRepository_GetByAspekID(t *testing.T) {
	db := setupRepositoryTestDB(t)
	aspekRepo := NewAspekRepository(db)
	kriteriaRepo := NewKriteriaRepository(db)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria1 := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteria2 := &models.Kriteria{AspekID: aspek.ID, Kode: "K2", Nama: "Kriteria 2", IsCore: false, Bobot: 1.0}
	kriteriaRepo.Create(kriteria1)
	kriteriaRepo.Create(kriteria2)

	kriterias, err := kriteriaRepo.GetByAspekID(aspek.ID)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(kriterias), 2)
}

func TestKriteriaRepository_Update(t *testing.T) {
	db := setupRepositoryTestDB(t)
	aspekRepo := NewAspekRepository(db)
	kriteriaRepo := NewKriteriaRepository(db)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria := &models.Kriteria{
		AspekID: aspek.ID,
		Kode:    "K1",
		Nama:    "Kriteria 1",
		IsCore:  true,
		Bobot:   1.0,
	}
	kriteriaRepo.Create(kriteria)

	kriteria.Nama = "Updated Kriteria"
	err := kriteriaRepo.Update(kriteria.ID, kriteria)
	assert.NoError(t, err)

	updated, _ := kriteriaRepo.GetByID(kriteria.ID)
	assert.Equal(t, "Updated Kriteria", updated.Nama)
}

