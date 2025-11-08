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

func TestKriteriaController_GetAll(t *testing.T) {
	db := setupControllerTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	kriteriaCtrl := NewKriteriaController(kriteriaService)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria1 := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteria2 := &models.Kriteria{AspekID: aspek.ID, Kode: "K2", Nama: "Kriteria 2", IsCore: false, Bobot: 1.0}
	kriteriaService.Create(kriteria1)
	kriteriaService.Create(kriteria2)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/kriteria", kriteriaCtrl.GetAll)

	req := httptest.NewRequest("GET", "/api/kriteria", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.GreaterOrEqual(t, len(response), 2)
}

func TestKriteriaController_Create(t *testing.T) {
	db := setupControllerTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	kriteriaCtrl := NewKriteriaController(kriteriaService)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/kriteria", kriteriaCtrl.Create)

	payload := map[string]interface{}{
		"aspek_id": aspek.ID,
		"kode":     "K1",
		"nama":     "Kriteria 1",
		"is_core":  true,
		"bobot":    1.0,
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/kriteria", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "K1", response["kode"])
}

func TestKriteriaController_GetByID(t *testing.T) {
	db := setupControllerTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	kriteriaCtrl := NewKriteriaController(kriteriaService)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/kriteria/:id", kriteriaCtrl.GetByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/kriteria/%d", kriteria.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "K1", response["kode"])
}

func TestKriteriaController_Update(t *testing.T) {
	db := setupControllerTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	kriteriaCtrl := NewKriteriaController(kriteriaService)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.PUT("/api/kriteria/:id", kriteriaCtrl.Update)

	payload := map[string]interface{}{
		"nama": "Updated Kriteria",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("PUT", fmt.Sprintf("/api/kriteria/%d", kriteria.ID), bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestKriteriaController_Delete(t *testing.T) {
	db := setupControllerTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	kriteriaRepo := repositories.NewKriteriaRepository(db)
	aspekService := services.NewAspekService(aspekRepo)
	kriteriaService := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	kriteriaCtrl := NewKriteriaController(kriteriaService)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	kriteria := &models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "Kriteria 1", IsCore: true, Bobot: 1.0}
	kriteriaService.Create(kriteria)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/api/kriteria/:id", kriteriaCtrl.Delete)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/kriteria/%d", kriteria.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

