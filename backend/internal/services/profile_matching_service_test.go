package services

import (
	"testing"

	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestProfileMatchingService_Calculate(t *testing.T) {
	db := setupServiceTestDB(t)

	// Setup repositories
	jabatanRepo := repositories.NewJabatanRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	nilaiTenagaKerjaRepo := repositories.NewNilaiTenagaKerjaRepository(db)
	profileMatchResultRepo := repositories.NewProfileMatchResultRepository(db)

	// Setup service
	service := NewProfileMatchingService(
		targetProfileRepo,
		kriteriaRepo,
		nilaiTenagaKerjaRepo,
		tenagaKerjaRepo,
		profileMatchResultRepo,
		jabatanRepo,
	)

	// Create test data
	jabatan := &models.Jabatan{Nama: "Manager", Deskripsi: "Manager Position"}
	jabatanRepo.Create(jabatan)

	aspek := &models.Aspek{Nama: "Test Aspek", Deskripsi: "Test", Persentase: 100.0}
	aspekRepo.Create(aspek)

	kriteria1 := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteria2 := &models.Kriteria{AspekID: aspek.ID, Kode: "K2", Nama: "Kriteria 2", IsCore: false, Bobot: 1.0}
	kriteriaRepo.Create(kriteria1)
	kriteriaRepo.Create(kriteria2)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "Test TK"}
	tenagaKerjaRepo.Create(tenagaKerja)

	targetProfile1 := &models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: kriteria1.ID, TargetNilai: 4.0}
	targetProfile2 := &models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: kriteria2.ID, TargetNilai: 3.0}
	targetProfileRepo.Create(targetProfile1)
	targetProfileRepo.Create(targetProfile2)

	nilai1 := &models.NilaiTenagaKerja{TenagaKerjaID: tenagaKerja.ID, KriteriaID: kriteria1.ID, Nilai: 4.0}
	nilai2 := &models.NilaiTenagaKerja{TenagaKerjaID: tenagaKerja.ID, KriteriaID: kriteria2.ID, Nilai: 3.0}
	nilaiTenagaKerjaRepo.Create(nilai1)
	nilaiTenagaKerjaRepo.Create(nilai2)

	// Test calculation
	req := CalculationRequest{
		JabatanID:      jabatan.ID,
		TenagaKerjaIDs: []uint{tenagaKerja.ID},
	}

	results, err := service.Calculate(req)
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Greater(t, results[0].TotalScore, 0.0)
}

func TestProfileMatchingService_Calculate_NoTenagaKerjaIDs(t *testing.T) {
	db := setupServiceTestDB(t)

	// Setup repositories
	jabatanRepo := repositories.NewJabatanRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	nilaiTenagaKerjaRepo := repositories.NewNilaiTenagaKerjaRepository(db)
	profileMatchResultRepo := repositories.NewProfileMatchResultRepository(db)

	// Setup service
	service := NewProfileMatchingService(
		targetProfileRepo,
		kriteriaRepo,
		nilaiTenagaKerjaRepo,
		tenagaKerjaRepo,
		profileMatchResultRepo,
		jabatanRepo,
	)

	// Create test data
	jabatan := &models.Jabatan{Nama: "Manager", Deskripsi: "Manager Position"}
	jabatanRepo.Create(jabatan)

	aspek := &models.Aspek{Nama: "Test Aspek", Deskripsi: "Test", Persentase: 100.0}
	aspekRepo.Create(aspek)

	kriteria1 := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaRepo.Create(kriteria1)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "Test TK"}
	tenagaKerjaRepo.Create(tenagaKerja)

	targetProfile1 := &models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: kriteria1.ID, TargetNilai: 4.0}
	targetProfileRepo.Create(targetProfile1)

	nilai1 := &models.NilaiTenagaKerja{TenagaKerjaID: tenagaKerja.ID, KriteriaID: kriteria1.ID, Nilai: 4.0}
	nilaiTenagaKerjaRepo.Create(nilai1)

	// Test calculation with empty TenagaKerjaIDs
	req := CalculationRequest{
		JabatanID:      jabatan.ID,
		TenagaKerjaIDs: []uint{}, // Empty, should get all
	}

	results, err := service.Calculate(req)
	assert.NoError(t, err)
	assert.Greater(t, len(results), 0)
}

func TestProfileMatchingService_Calculate_InvalidJabatan(t *testing.T) {
	db := setupServiceTestDB(t)

	// Setup repositories
	jabatanRepo := repositories.NewJabatanRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	nilaiTenagaKerjaRepo := repositories.NewNilaiTenagaKerjaRepository(db)
	profileMatchResultRepo := repositories.NewProfileMatchResultRepository(db)

	// Setup service
	service := NewProfileMatchingService(
		targetProfileRepo,
		kriteriaRepo,
		nilaiTenagaKerjaRepo,
		tenagaKerjaRepo,
		profileMatchResultRepo,
		jabatanRepo,
	)

	req := CalculationRequest{
		JabatanID:      999, // Non-existent jabatan
		TenagaKerjaIDs: []uint{1},
	}

	_, err := service.Calculate(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "jabatan not found")
}

func TestProfileMatchingService_GetAllResults(t *testing.T) {
	db := setupServiceTestDB(t)

	// Setup repositories
	jabatanRepo := repositories.NewJabatanRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	nilaiTenagaKerjaRepo := repositories.NewNilaiTenagaKerjaRepository(db)
	profileMatchResultRepo := repositories.NewProfileMatchResultRepository(db)

	// Setup service
	service := NewProfileMatchingService(
		targetProfileRepo,
		kriteriaRepo,
		nilaiTenagaKerjaRepo,
		tenagaKerjaRepo,
		profileMatchResultRepo,
		jabatanRepo,
	)

	results, err := service.GetAllResults()
	assert.NoError(t, err)
	assert.NotNil(t, results)
}

func TestProfileMatchingService_GetResultByID(t *testing.T) {
	db := setupServiceTestDB(t)

	// Setup repositories
	jabatanRepo := repositories.NewJabatanRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	nilaiTenagaKerjaRepo := repositories.NewNilaiTenagaKerjaRepository(db)
	profileMatchResultRepo := repositories.NewProfileMatchResultRepository(db)

	// Setup service
	service := NewProfileMatchingService(
		targetProfileRepo,
		kriteriaRepo,
		nilaiTenagaKerjaRepo,
		tenagaKerjaRepo,
		profileMatchResultRepo,
		jabatanRepo,
	)

	// Create a result
	jabatan := &models.Jabatan{Nama: "Manager"}
	jabatanRepo.Create(jabatan)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "Test TK"}
	tenagaKerjaRepo.Create(tenagaKerja)

	result := &models.ProfileMatchResult{
		TenagaKerjaID:   tenagaKerja.ID,
		JabatanID:       jabatan.ID,
		TotalScore:      85.5,
		CoreFactor:      90.0,
		SecondaryFactor: 80.0,
	}
	profileMatchResultRepo.Create(result)

	found, err := service.GetResultByID(result.ID)
	assert.NoError(t, err)
	assert.Equal(t, result.TotalScore, found.TotalScore)
}

