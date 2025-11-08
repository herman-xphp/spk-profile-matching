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

func TestTenagaKerjaController_GetAll(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	tenagaKerjaCtrl := NewTenagaKerjaController(tenagaKerjaService)

	tenagaKerja1 := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerja2 := &models.TenagaKerja{NIK: "TK002", Nama: "Jane Doe"}
	tenagaKerjaService.Create(tenagaKerja1)
	tenagaKerjaService.Create(tenagaKerja2)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/tenaga-kerja", tenagaKerjaCtrl.GetAll)

	req := httptest.NewRequest("GET", "/api/tenaga-kerja", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.GreaterOrEqual(t, len(response), 2)
}

func TestTenagaKerjaController_Create(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	tenagaKerjaCtrl := NewTenagaKerjaController(tenagaKerjaService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/tenaga-kerja", tenagaKerjaCtrl.Create)

	payload := map[string]interface{}{
		"nik":  "TK001",
		"nama": "John Doe",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/tenaga-kerja", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "TK001", response["nik"])
}

func TestTenagaKerjaController_GetByID(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	tenagaKerjaCtrl := NewTenagaKerjaController(tenagaKerjaService)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/tenaga-kerja/:id", tenagaKerjaCtrl.GetByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/tenaga-kerja/%d", tenagaKerja.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "TK001", response["nik"])
}

func TestTenagaKerjaController_Update(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	tenagaKerjaCtrl := NewTenagaKerjaController(tenagaKerjaService)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.PUT("/api/tenaga-kerja/:id", tenagaKerjaCtrl.Update)

	payload := map[string]interface{}{
		"nama": "Jane Doe",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("PUT", fmt.Sprintf("/api/tenaga-kerja/%d", tenagaKerja.ID), bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTenagaKerjaController_Delete(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	tenagaKerjaCtrl := NewTenagaKerjaController(tenagaKerjaService)

	tenagaKerja := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/api/tenaga-kerja/:id", tenagaKerjaCtrl.Delete)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/tenaga-kerja/%d", tenagaKerja.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTenagaKerjaController_Create_DuplicateNIK(t *testing.T) {
	db := setupControllerTestDB(t)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(db)
	tenagaKerjaService := services.NewTenagaKerjaService(tenagaKerjaRepo)
	tenagaKerjaCtrl := NewTenagaKerjaController(tenagaKerjaService)

	tenagaKerja1 := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerjaService.Create(tenagaKerja1)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/tenaga-kerja", tenagaKerjaCtrl.Create)

	payload := map[string]interface{}{
		"nik":  "TK001",
		"nama": "Jane Doe",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/tenaga-kerja", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

