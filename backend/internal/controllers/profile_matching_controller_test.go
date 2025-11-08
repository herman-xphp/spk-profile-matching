package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProfileMatchingController_Calculate(t *testing.T) {
	db := setupControllerTestDB(t)

	// Setup repositories
	jabatanRepo := repositories.NewJabatanRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	nilaiTenagaKerjaRepo := repositories.NewNilaiTenagaKerjaRepository(db)
	profileMatchResultRepo := repositories.NewProfileMatchResultRepository(db)

	// Setup services
	profileMatchingSvc := services.NewProfileMatchingService(
		targetProfileRepo,
		kriteriaRepo,
		nilaiTenagaKerjaRepo,
		tenagaKerjaRepo,
		profileMatchResultRepo,
		jabatanRepo,
	)

	// Setup controller
	profileMatchingCtrl := NewProfileMatchingController(profileMatchingSvc)

	// Create test data
	jabatan := &models.Jabatan{Nama: "Manager"}
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

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/profile-matching/calculate", profileMatchingCtrl.Calculate)

	payload := map[string]interface{}{
		"jabatan_id":       jabatan.ID,
		"tenaga_kerja_ids": []uint{tenagaKerja.ID},
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/profile-matching/calculate", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Greater(t, len(response), 0)
	assert.Contains(t, response[0], "total_score")
}

func TestProfileMatchingController_GetAllResults(t *testing.T) {
	db := setupControllerTestDB(t)

	// Setup repositories
	jabatanRepo := repositories.NewJabatanRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	nilaiTenagaKerjaRepo := repositories.NewNilaiTenagaKerjaRepository(db)
	profileMatchResultRepo := repositories.NewProfileMatchResultRepository(db)

	// Setup services
	profileMatchingSvc := services.NewProfileMatchingService(
		targetProfileRepo,
		kriteriaRepo,
		nilaiTenagaKerjaRepo,
		tenagaKerjaRepo,
		profileMatchResultRepo,
		jabatanRepo,
	)

	// Setup controller
	profileMatchingCtrl := NewProfileMatchingController(profileMatchingSvc)

	// Create test data
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

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/profile-matching/results", profileMatchingCtrl.GetAllResults)

	req := httptest.NewRequest("GET", "/api/profile-matching/results", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Greater(t, len(response), 0)
}

func TestProfileMatchingController_GetResultByID(t *testing.T) {
	db := setupControllerTestDB(t)

	// Setup repositories
	jabatanRepo := repositories.NewJabatanRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	nilaiTenagaKerjaRepo := repositories.NewNilaiTenagaKerjaRepository(db)
	profileMatchResultRepo := repositories.NewProfileMatchResultRepository(db)

	// Setup services
	profileMatchingSvc := services.NewProfileMatchingService(
		targetProfileRepo,
		kriteriaRepo,
		nilaiTenagaKerjaRepo,
		tenagaKerjaRepo,
		profileMatchResultRepo,
		jabatanRepo,
	)

	// Setup controller
	profileMatchingCtrl := NewProfileMatchingController(profileMatchingSvc)

	// Create test data
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

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/profile-matching/results/:id", profileMatchingCtrl.GetResultByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/profile-matching/results/%d", result.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(result.ID), response["id"])
	assert.Contains(t, response, "total_score")
}

