package services

import (
	"testing"

	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestNilaiTenagaKerjaService_Create(t *testing.T) {
	db := setupServiceTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	nilaiRepo := repositories.NewNilaiTenagaKerjaRepository(db)

	tenagaKerjaService := NewTenagaKerjaService(tenagaKerjaRepo)
	aspekService := NewAspekService(aspekRepo)
	kriteriaService := NewKriteriaService(kriteriaRepo, aspekRepo)
	service := NewNilaiTenagaKerjaService(nilaiRepo, tenagaKerjaRepo, kriteriaRepo)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	nilai := &models.NilaiTenagaKerja{
		TenagaKerjaID: tenagaKerja.ID,
		KriteriaID:    kriteria.ID,
		Nilai:         4.0,
	}

	err := service.Create(nilai)
	assert.NoError(t, err)
	assert.NotZero(t, nilai.ID)
}

func TestNilaiTenagaKerjaService_Create_InvalidTenagaKerja(t *testing.T) {
	db := setupServiceTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	nilaiRepo := repositories.NewNilaiTenagaKerjaRepository(db)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)

	aspekService := NewAspekService(aspekRepo)
	kriteriaService := NewKriteriaService(kriteriaRepo, aspekRepo)
	service := NewNilaiTenagaKerjaService(nilaiRepo, tenagaKerjaRepo, kriteriaRepo)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	nilai := &models.NilaiTenagaKerja{
		TenagaKerjaID: 999, // Non-existent tenaga kerja
		KriteriaID:    kriteria.ID,
		Nilai:         4.0,
	}

	err := service.Create(nilai)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "tenaga kerja not found")
}

func TestNilaiTenagaKerjaService_GetByID(t *testing.T) {
	db := setupServiceTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	nilaiRepo := repositories.NewNilaiTenagaKerjaRepository(db)

	tenagaKerjaService := NewTenagaKerjaService(tenagaKerjaRepo)
	aspekService := NewAspekService(aspekRepo)
	kriteriaService := NewKriteriaService(kriteriaRepo, aspekRepo)
	service := NewNilaiTenagaKerjaService(nilaiRepo, tenagaKerjaRepo, kriteriaRepo)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	nilai := &models.NilaiTenagaKerja{
		TenagaKerjaID: tenagaKerja.ID,
		KriteriaID:    kriteria.ID,
		Nilai:         4.0,
	}
	service.Create(nilai)

	found, err := service.GetByID(nilai.ID)
	assert.NoError(t, err)
	assert.Equal(t, nilai.Nilai, found.Nilai)
}

func TestNilaiTenagaKerjaService_Update(t *testing.T) {
	db := setupServiceTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	nilaiRepo := repositories.NewNilaiTenagaKerjaRepository(db)

	tenagaKerjaService := NewTenagaKerjaService(tenagaKerjaRepo)
	aspekService := NewAspekService(aspekRepo)
	kriteriaService := NewKriteriaService(kriteriaRepo, aspekRepo)
	service := NewNilaiTenagaKerjaService(nilaiRepo, tenagaKerjaRepo, kriteriaRepo)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	nilai := &models.NilaiTenagaKerja{
		TenagaKerjaID: tenagaKerja.ID,
		KriteriaID:    kriteria.ID,
		Nilai:         4.0,
	}
	service.Create(nilai)

	nilai.Nilai = 5.0
	err := service.Update(nilai.ID, nilai)
	assert.NoError(t, err)

	updated, _ := service.GetByID(nilai.ID)
	assert.Equal(t, 5.0, updated.Nilai)
}

func TestNilaiTenagaKerjaService_Delete(t *testing.T) {
	db := setupServiceTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	nilaiRepo := repositories.NewNilaiTenagaKerjaRepository(db)

	tenagaKerjaService := NewTenagaKerjaService(tenagaKerjaRepo)
	aspekService := NewAspekService(aspekRepo)
	kriteriaService := NewKriteriaService(kriteriaRepo, aspekRepo)
	service := NewNilaiTenagaKerjaService(nilaiRepo, tenagaKerjaRepo, kriteriaRepo)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	nilai := &models.NilaiTenagaKerja{
		TenagaKerjaID: tenagaKerja.ID,
		KriteriaID:    kriteria.ID,
		Nilai:         4.0,
	}
	service.Create(nilai)

	err := service.Delete(nilai.ID)
	assert.NoError(t, err)

	_, err = service.GetByID(nilai.ID)
	assert.Error(t, err)
}

