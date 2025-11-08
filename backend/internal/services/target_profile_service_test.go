package services

import (
	"testing"

	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestTargetProfileService_Create(t *testing.T) {
	db := setupServiceTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)

	jabatanService := NewJabatanService(jabatanRepo)
	aspekService := NewAspekService(aspekRepo)
	kriteriaService := NewKriteriaService(kriteriaRepo, aspekRepo)
	service := NewTargetProfileService(targetProfileRepo, jabatanRepo, kriteriaRepo)

	jabatan := &models.Jabatan{Nama: "Manager"}
	jabatanService.Create(jabatan)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	targetProfile := &models.TargetProfile{
		JabatanID:  jabatan.ID,
		KriteriaID: kriteria.ID,
		TargetNilai: 4.0,
	}

	err := service.Create(targetProfile)
	assert.NoError(t, err)
	assert.NotZero(t, targetProfile.ID)
}

func TestTargetProfileService_Create_InvalidJabatan(t *testing.T) {
	db := setupServiceTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)
	jabatanRepo := repositories.NewJabatanRepository(db)

	aspekService := NewAspekService(aspekRepo)
	kriteriaService := NewKriteriaService(kriteriaRepo, aspekRepo)
	service := NewTargetProfileService(targetProfileRepo, jabatanRepo, kriteriaRepo)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	targetProfile := &models.TargetProfile{
		JabatanID:  999, // Non-existent jabatan
		KriteriaID: kriteria.ID,
		TargetNilai: 4.0,
	}

	err := service.Create(targetProfile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "jabatan not found")
}

func TestTargetProfileService_GetByID(t *testing.T) {
	db := setupServiceTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)

	jabatanService := NewJabatanService(jabatanRepo)
	aspekService := NewAspekService(aspekRepo)
	kriteriaService := NewKriteriaService(kriteriaRepo, aspekRepo)
	service := NewTargetProfileService(targetProfileRepo, jabatanRepo, kriteriaRepo)

	jabatan := &models.Jabatan{Nama: "Manager"}
	jabatanService.Create(jabatan)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	targetProfile := &models.TargetProfile{
		JabatanID:  jabatan.ID,
		KriteriaID: kriteria.ID,
		TargetNilai: 4.0,
	}
	service.Create(targetProfile)

	found, err := service.GetByID(targetProfile.ID)
	assert.NoError(t, err)
	assert.Equal(t, targetProfile.TargetNilai, found.TargetNilai)
}

