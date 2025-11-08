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

func TestTargetProfileController_GetAll(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)

	jabatanService := services.NewJabatanService(jabatanRepo)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	targetProfileService := services.NewTargetProfileService(targetProfileRepo, jabatanRepo, kriteriaRepo)
	targetProfileCtrl := NewTargetProfileController(targetProfileService)

	jabatan := &models.Jabatan{Nama: "Manager"}
	jabatanService.Create(jabatan)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	targetProfile1 := &models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: kriteria.ID, TargetNilai: 4.0}
	targetProfile2 := &models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: kriteria.ID, TargetNilai: 3.0}
	targetProfileService.Create(targetProfile1)
	targetProfileService.Create(targetProfile2)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/target-profiles", targetProfileCtrl.GetAll)

	req := httptest.NewRequest("GET", "/api/target-profiles", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.GreaterOrEqual(t, len(response), 2)
}

func TestTargetProfileController_Create(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)

	jabatanService := services.NewJabatanService(jabatanRepo)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	targetProfileService := services.NewTargetProfileService(targetProfileRepo, jabatanRepo, kriteriaRepo)
	targetProfileCtrl := NewTargetProfileController(targetProfileService)

	jabatan := &models.Jabatan{Nama: "Manager"}
	jabatanService.Create(jabatan)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/target-profiles", targetProfileCtrl.Create)

	payload := map[string]interface{}{
		"jabatan_id":   jabatan.ID,
		"kriteria_id":  kriteria.ID,
		"target_nilai": 4.0,
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/target-profiles", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(jabatan.ID), response["jabatan_id"])
}

func TestTargetProfileController_GetByID(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)

	jabatanService := services.NewJabatanService(jabatanRepo)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	targetProfileService := services.NewTargetProfileService(targetProfileRepo, jabatanRepo, kriteriaRepo)
	targetProfileCtrl := NewTargetProfileController(targetProfileService)

	jabatan := &models.Jabatan{Nama: "Manager"}
	jabatanService.Create(jabatan)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	targetProfile := &models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: kriteria.ID, TargetNilai: 4.0}
	targetProfileService.Create(targetProfile)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/target-profiles/:id", targetProfileCtrl.GetByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/target-profiles/%d", targetProfile.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(targetProfile.ID), response["id"])
}

func TestTargetProfileController_Update(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)

	jabatanService := services.NewJabatanService(jabatanRepo)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	targetProfileService := services.NewTargetProfileService(targetProfileRepo, jabatanRepo, kriteriaRepo)
	targetProfileCtrl := NewTargetProfileController(targetProfileService)

	jabatan := &models.Jabatan{Nama: "Manager"}
	jabatanService.Create(jabatan)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	targetProfile := &models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: kriteria.ID, TargetNilai: 4.0}
	targetProfileService.Create(targetProfile)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.PUT("/api/target-profiles/:id", targetProfileCtrl.Update)

	payload := map[string]interface{}{
		"target_nilai": 5.0,
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("PUT", fmt.Sprintf("/api/target-profiles/%d", targetProfile.ID), bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTargetProfileController_Delete(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	targetProfileRepo := repositories.NewTargetProfileRepository(db)

	jabatanService := services.NewJabatanService(jabatanRepo)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	targetProfileService := services.NewTargetProfileService(targetProfileRepo, jabatanRepo, kriteriaRepo)
	targetProfileCtrl := NewTargetProfileController(targetProfileService)

	jabatan := &models.Jabatan{Nama: "Manager"}
	jabatanService.Create(jabatan)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	targetProfile := &models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: kriteria.ID, TargetNilai: 4.0}
	targetProfileService.Create(targetProfile)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/api/target-profiles/:id", targetProfileCtrl.Delete)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/target-profiles/%d", targetProfile.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

