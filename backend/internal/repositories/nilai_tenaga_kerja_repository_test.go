package repositories

import (
	"testing"

	"backend/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestNilaiTenagaKerjaRepository_Create(t *testing.T) {
	db := setupRepositoryTestDB(t)
	tenagaKerjaRepo := NewTenagaKerjaRepository(db)
	aspekRepo := NewAspekRepository(db)
	kriteriaRepo := NewKriteriaRepository(db)
	nilaiRepo := NewNilaiTenagaKerjaRepository(db)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaRepo.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaRepo.Create(kriteria)

	nilai := &models.NilaiTenagaKerja{
		TenagaKerjaID: tenagaKerja.ID,
		KriteriaID:    kriteria.ID,
		Nilai:         4.0,
	}

	err := nilaiRepo.Create(nilai)
	assert.NoError(t, err)
	assert.NotZero(t, nilai.ID)
}

func TestNilaiTenagaKerjaRepository_GetByTenagaKerjaID(t *testing.T) {
	db := setupRepositoryTestDB(t)
	tenagaKerjaRepo := NewTenagaKerjaRepository(db)
	aspekRepo := NewAspekRepository(db)
	kriteriaRepo := NewKriteriaRepository(db)
	nilaiRepo := NewNilaiTenagaKerjaRepository(db)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaRepo.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekRepo.Create(aspek)

	kriteria1 := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteria2 := &models.Kriteria{AspekID: aspek.ID, Kode: "K2", Nama: "Kriteria 2", IsCore: false, Bobot: 1.0}
	kriteriaRepo.Create(kriteria1)
	kriteriaRepo.Create(kriteria2)

	nilai1 := &models.NilaiTenagaKerja{TenagaKerjaID: tenagaKerja.ID, KriteriaID: kriteria1.ID, Nilai: 4.0}
	nilai2 := &models.NilaiTenagaKerja{TenagaKerjaID: tenagaKerja.ID, KriteriaID: kriteria2.ID, Nilai: 3.0}
	nilaiRepo.Create(nilai1)
	nilaiRepo.Create(nilai2)

	nilais, err := nilaiRepo.GetByTenagaKerjaID(tenagaKerja.ID)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(nilais), 2)
}

