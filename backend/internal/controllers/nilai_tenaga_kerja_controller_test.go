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

func TestNilaiTenagaKerjaController_GetAll(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	nilaiRepo := repositories.NewNilaiTenagaKerjaRepository(db)

	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	nilaiService := services.NewNilaiTenagaKerjaService(nilaiRepo, tenagaKerjaRepo, kriteriaRepo)
	nilaiCtrl := NewNilaiTenagaKerjaController(nilaiService)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria1 := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteria2 := &models.Kriteria{AspekID: aspek.ID, Kode: "K2", Nama: "Kriteria 2", IsCore: false, Bobot: 1.0}
	kriteriaService.Create(kriteria1)
	kriteriaService.Create(kriteria2)

	nilai1 := &models.NilaiTenagaKerja{TenagaKerjaID: tenagaKerja.ID, KriteriaID: kriteria1.ID, Nilai: 4.0}
	nilai2 := &models.NilaiTenagaKerja{TenagaKerjaID: tenagaKerja.ID, KriteriaID: kriteria2.ID, Nilai: 3.0}
	nilaiService.Create(nilai1)
	nilaiService.Create(nilai2)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/nilai-tenaga-kerja", nilaiCtrl.GetAll)

	req := httptest.NewRequest("GET", "/api/nilai-tenaga-kerja", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.GreaterOrEqual(t, len(response), 2)
}

func TestNilaiTenagaKerjaController_Create(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	nilaiRepo := repositories.NewNilaiTenagaKerjaRepository(db)

	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	nilaiService := services.NewNilaiTenagaKerjaService(nilaiRepo, tenagaKerjaRepo, kriteriaRepo)
	nilaiCtrl := NewNilaiTenagaKerjaController(nilaiService)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/nilai-tenaga-kerja", nilaiCtrl.Create)

	payload := map[string]interface{}{
		"tenaga_kerja_id": tenagaKerja.ID,
		"kriteria_id":     kriteria.ID,
		"nilai":           4.0,
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/nilai-tenaga-kerja", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 4.0, response["nilai"])
}

func TestNilaiTenagaKerjaController_GetByID(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	nilaiRepo := repositories.NewNilaiTenagaKerjaRepository(db)

	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	nilaiService := services.NewNilaiTenagaKerjaService(nilaiRepo, tenagaKerjaRepo, kriteriaRepo)
	nilaiCtrl := NewNilaiTenagaKerjaController(nilaiService)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	nilai := &models.NilaiTenagaKerja{TenagaKerjaID: tenagaKerja.ID, KriteriaID: kriteria.ID, Nilai: 4.0}
	nilaiService.Create(nilai)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/nilai-tenaga-kerja/:id", nilaiCtrl.GetByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/nilai-tenaga-kerja/%d", nilai.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, 4.0, response["nilai"])
}

func TestNilaiTenagaKerjaController_Update(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	nilaiRepo := repositories.NewNilaiTenagaKerjaRepository(db)

	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	nilaiService := services.NewNilaiTenagaKerjaService(nilaiRepo, tenagaKerjaRepo, kriteriaRepo)
	nilaiCtrl := NewNilaiTenagaKerjaController(nilaiService)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	nilai := &models.NilaiTenagaKerja{TenagaKerjaID: tenagaKerja.ID, KriteriaID: kriteria.ID, Nilai: 4.0}
	nilaiService.Create(nilai)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.PUT("/api/nilai-tenaga-kerja/:id", nilaiCtrl.Update)

	payload := map[string]interface{}{
		"nilai": 5.0,
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("PUT", fmt.Sprintf("/api/nilai-tenaga-kerja/%d", nilai.ID), bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNilaiTenagaKerjaController_Delete(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	nilaiRepo := repositories.NewNilaiTenagaKerjaRepository(db)

	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	nilaiService := services.NewNilaiTenagaKerjaService(nilaiRepo, tenagaKerjaRepo, kriteriaRepo)
	nilaiCtrl := NewNilaiTenagaKerjaController(nilaiService)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	nilai := &models.NilaiTenagaKerja{TenagaKerjaID: tenagaKerja.ID, KriteriaID: kriteria.ID, Nilai: 4.0}
	nilaiService.Create(nilai)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/api/nilai-tenaga-kerja/:id", nilaiCtrl.Delete)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/nilai-tenaga-kerja/%d", nilai.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

