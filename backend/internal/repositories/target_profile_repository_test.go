package repositories

import (
	"testing"

	"backend/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestTargetProfileRepository_Create(t *testing.T) {
	db := setupRepositoryTestDB(t)
	jabatanRepo := NewJabatanRepository(db)
	aspekRepo := NewAspekRepository(db)
	kriteriaRepo := NewKriteriaRepository(db)
	targetProfileRepo := NewTargetProfileRepository(db)

	jabatan := &models.Jabatan{Nama: "Manager"}
	jabatanRepo.Create(jabatan)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaRepo.Create(kriteria)

	targetProfile := &models.TargetProfile{
		JabatanID:  jabatan.ID,
		KriteriaID: kriteria.ID,
		TargetNilai: 4.0,
	}

	err := targetProfileRepo.Create(targetProfile)
	assert.NoError(t, err)
	assert.NotZero(t, targetProfile.ID)
}

func TestTargetProfileRepository_GetByJabatanID(t *testing.T) {
	db := setupRepositoryTestDB(t)
	jabatanRepo := NewJabatanRepository(db)
	aspekRepo := NewAspekRepository(db)
	kriteriaRepo := NewKriteriaRepository(db)
	targetProfileRepo := NewTargetProfileRepository(db)

	jabatan := &models.Jabatan{Nama: "Manager"}
	jabatanRepo.Create(jabatan)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaRepo.Create(kriteria)

	targetProfile1 := &models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: kriteria.ID, TargetNilai: 4.0}
	targetProfile2 := &models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: kriteria.ID, TargetNilai: 3.0}
	targetProfileRepo.Create(targetProfile1)
	targetProfileRepo.Create(targetProfile2)

	profiles, err := targetProfileRepo.GetByJabatanID(jabatan.ID)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(profiles), 2)
}

