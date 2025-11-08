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

func TestAspekController_GetAll(t *testing.T) {
	db := setupControllerTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	aspekService := services.NewAspekService(aspekRepo)
	aspekCtrl := NewAspekController(aspekService)

	aspek1 := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspek2 := &models.Aspek{Nama: "Kepribadian", Persentase: 50.0}
	aspekService.Create(aspek1)
	aspekService.Create(aspek2)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/aspek", aspekCtrl.GetAll)

	req := httptest.NewRequest("GET", "/api/aspek", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.GreaterOrEqual(t, len(response), 2)
}

func TestAspekController_Create(t *testing.T) {
	db := setupControllerTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	aspekService := services.NewAspekService(aspekRepo)
	aspekCtrl := NewAspekController(aspekService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/aspek", aspekCtrl.Create)

	payload := map[string]interface{}{
		"nama":       "Kompetensi",
		"deskripsi":  "Aspek Kompetensi",
		"persentase": 50.0,
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/aspek", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Kompetensi", response["nama"])
}

func TestAspekController_GetByID(t *testing.T) {
	db := setupControllerTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	aspekService := services.NewAspekService(aspekRepo)
	aspekCtrl := NewAspekController(aspekService)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/aspek/:id", aspekCtrl.GetByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/aspek/%d", aspek.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Kompetensi", response["nama"])
}

func TestAspekController_Update(t *testing.T) {
	db := setupControllerTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	aspekService := services.NewAspekService(aspekRepo)
	aspekCtrl := NewAspekController(aspekService)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.PUT("/api/aspek/:id", aspekCtrl.Update)

	payload := map[string]interface{}{
		"nama": "Kompetensi Teknis",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("PUT", fmt.Sprintf("/api/aspek/%d", aspek.ID), bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAspekController_Delete(t *testing.T) {
	db := setupControllerTestDB(t)
	aspekRepo := repositories.NewAspekRepository(db)
	aspekService := services.NewAspekService(aspekRepo)
	aspekCtrl := NewAspekController(aspekService)

	aspek := &models.Aspek{Nama: "Kompetensi", Persentase: 50.0}
	aspekService.Create(aspek)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/api/aspek/:id", aspekCtrl.Delete)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/aspek/%d", aspek.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

