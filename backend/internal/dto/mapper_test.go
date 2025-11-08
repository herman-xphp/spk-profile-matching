package dto

import (
	"testing"

	"backend/internal/models"
	"backend/pkg/database"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupMapperTestDB(t *testing.T) *gorm.DB {
	db, err := database.ConnectTestDB(t)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Clear all tables
	if err := database.CleanTestDB(db); err != nil {
		t.Fatalf("Failed to clean test database: %v", err)
	}

	return db
}

func TestMapUserToResponse(t *testing.T) {
	user := &models.User{
		Email:    "test@example.com",
		Nama:     "Test User",
		Role:     "user",
		IsActive: true,
	}

	response := MapUserToResponse(user)
	assert.Equal(t, user.Email, response.Email)
	assert.Equal(t, user.Nama, response.Nama)
	assert.Equal(t, user.Role, response.Role)
	assert.Equal(t, user.IsActive, response.IsActive)
}

func TestMapUsersToResponse(t *testing.T) {
	users := []models.User{
		{Email: "user1@example.com", Nama: "User 1"},
		{Email: "user2@example.com", Nama: "User 2"},
	}

	responses := MapUsersToResponse(users)
	assert.Len(t, responses, 2)
	assert.Equal(t, users[0].Email, responses[0].Email)
	assert.Equal(t, users[1].Email, responses[1].Email)
}

func TestMapJabatanToResponse(t *testing.T) {
	jabatan := &models.Jabatan{
		Nama:      "Manager",
		Deskripsi: "Manager Position",
	}

	response := MapJabatanToResponse(jabatan)
	assert.Equal(t, jabatan.Nama, response.Nama)
	assert.Equal(t, jabatan.Deskripsi, response.Deskripsi)
}

func TestMapJabatansToResponse(t *testing.T) {
	jabatans := []models.Jabatan{
		{Nama: "Manager", Deskripsi: "Manager Position"},
		{Nama: "Staff", Deskripsi: "Staff Position"},
	}

	responses := MapJabatansToResponse(jabatans)
	assert.Len(t, responses, 2)
	assert.Equal(t, jabatans[0].Nama, responses[0].Nama)
}

func TestMapAspekToResponse(t *testing.T) {
	aspek := &models.Aspek{
		Nama:       "Kompetensi",
		Deskripsi:  "Aspek Kompetensi",
		Persentase: 50.0,
	}

	response := MapAspekToResponse(aspek)
	assert.Equal(t, aspek.Nama, response.Nama)
	assert.Equal(t, aspek.Persentase, response.Persentase)
}

func TestMapKriteriaToResponse(t *testing.T) {
	aspek := models.Aspek{
		Nama: "Kompetensi",
	}

	kriteria := &models.Kriteria{
		AspekID: 1,
		Kode:    "K1",
		Nama:    "Kriteria 1",
		IsCore:  true,
		Bobot:   1.0,
		Aspek:   aspek,
	}

	response := MapKriteriaToResponse(kriteria)
	assert.Equal(t, kriteria.Kode, response.Kode)
	assert.Equal(t, kriteria.Nama, response.Nama)
	assert.Equal(t, kriteria.IsCore, response.IsCore)
	if kriteria.Aspek.ID != 0 {
		assert.NotNil(t, response.Aspek)
	}
}

func TestMapProfileMatchResultToResponse(t *testing.T) {
	tenagaKerja := models.TenagaKerja{
		NIK:  "TK001",
		Nama: "John Doe",
	}

	jabatan := models.Jabatan{
		Nama: "Manager",
	}

	result := &models.ProfileMatchResult{
		TenagaKerjaID:   1,
		JabatanID:       1,
		TotalScore:      85.5,
		CoreFactor:      90.0,
		SecondaryFactor: 80.0,
		TenagaKerja:     tenagaKerja,
		Jabatan:         jabatan,
	}

	response := MapProfileMatchResultToResponse(result)
	assert.Equal(t, result.TotalScore, response.TotalScore)
	assert.Equal(t, result.TotalScore, response.ScoreTotal) // Alias
	assert.Equal(t, result.CoreFactor, response.CoreFactor)
	assert.Equal(t, result.SecondaryFactor, response.SecondaryFactor)
	if result.TenagaKerja.ID != 0 {
		assert.NotNil(t, response.TenagaKerja)
	}
	if result.Jabatan.ID != 0 {
		assert.NotNil(t, response.Jabatan)
	}
}

func TestMapProfileMatchResultToRankingResponse(t *testing.T) {
	tenagaKerja := models.TenagaKerja{
		NIK:  "TK001",
		Nama: "John Doe",
	}

	jabatan := models.Jabatan{
		Nama: "Manager",
	}

	result := &models.ProfileMatchResult{
		TenagaKerjaID:   1,
		JabatanID:       1,
		TotalScore:      85.5,
		CoreFactor:      90.0,
		SecondaryFactor: 80.0,
		TenagaKerja:     tenagaKerja,
		Jabatan:         jabatan,
	}

	rank := 1
	response := MapProfileMatchResultToRankingResponse(result, rank)
	assert.Equal(t, rank, response.Rank)
	assert.Equal(t, result.TotalScore, response.ScoreTotal)
	assert.Equal(t, result.CoreFactor, response.CoreFactor)
	if result.TenagaKerja.ID != 0 {
		assert.NotNil(t, response.TenagaKerja)
	}
}

func TestMapProfileMatchResultsToRankingResponse(t *testing.T) {
	db := setupMapperTestDB(t)

	// Create test data
	tenagaKerja1 := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerja2 := &models.TenagaKerja{NIK: "TK002", Nama: "Jane Doe"}
	db.Create(tenagaKerja1)
	db.Create(tenagaKerja2)

	jabatan := &models.Jabatan{Nama: "Manager"}
	db.Create(jabatan)

	result1 := &models.ProfileMatchResult{
		TenagaKerjaID:   tenagaKerja1.ID,
		JabatanID:       jabatan.ID,
		TotalScore:      90.0,
		CoreFactor:      95.0,
		SecondaryFactor: 85.0,
	}
	result2 := &models.ProfileMatchResult{
		TenagaKerjaID:   tenagaKerja2.ID,
		JabatanID:       jabatan.ID,
		TotalScore:      85.0,
		CoreFactor:      90.0,
		SecondaryFactor: 80.0,
	}
	db.Create(result1)
	db.Create(result2)

	// Reload with relations
	db.Preload("TenagaKerja").Preload("Jabatan").Find(result1, result1.ID)
	db.Preload("TenagaKerja").Preload("Jabatan").Find(result2, result2.ID)

	results := []models.ProfileMatchResult{*result1, *result2}
	responses := MapProfileMatchResultsToRankingResponse(results)

	assert.Len(t, responses, 2)
	assert.Equal(t, 1, responses[0].Rank)
	assert.Equal(t, 2, responses[1].Rank)
	assert.Equal(t, result1.TotalScore, responses[0].ScoreTotal)
	assert.Equal(t, result2.TotalScore, responses[1].ScoreTotal)
}
