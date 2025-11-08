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

func TestJabatanController_GetAll(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	jabatanService := services.NewJabatanService(jabatanRepo)
	jabatanCtrl := NewJabatanController(jabatanService)

	// Create test jabatan
	jabatan1 := &models.Jabatan{Nama: "Manager", Deskripsi: "Manager Position"}
	jabatan2 := &models.Jabatan{Nama: "Staff", Deskripsi: "Staff Position"}
	jabatanService.Create(jabatan1)
	jabatanService.Create(jabatan2)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/jabatan", jabatanCtrl.GetAll)

	req := httptest.NewRequest("GET", "/api/jabatan", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.GreaterOrEqual(t, len(response), 2)
}

func TestJabatanController_Create(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	jabatanService := services.NewJabatanService(jabatanRepo)
	jabatanCtrl := NewJabatanController(jabatanService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/jabatan", jabatanCtrl.Create)

	payload := map[string]interface{}{
		"nama":      "Manager",
		"deskripsi": "Manager Position",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/jabatan", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Manager", response["nama"])
}

func TestJabatanController_GetByID(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	jabatanService := services.NewJabatanService(jabatanRepo)
	jabatanCtrl := NewJabatanController(jabatanService)

	jabatan := &models.Jabatan{Nama: "Manager", Deskripsi: "Manager Position"}
	jabatanService.Create(jabatan)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/jabatan/:id", jabatanCtrl.GetByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/jabatan/%d", jabatan.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Manager", response["nama"])
}

func TestJabatanController_Update(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	jabatanService := services.NewJabatanService(jabatanRepo)
	jabatanCtrl := NewJabatanController(jabatanService)

	jabatan := &models.Jabatan{Nama: "Manager", Deskripsi: "Manager Position"}
	jabatanService.Create(jabatan)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.PUT("/api/jabatan/:id", jabatanCtrl.Update)

	payload := map[string]interface{}{
		"nama": "Senior Manager",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("PUT", fmt.Sprintf("/api/jabatan/%d", jabatan.ID), bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestJabatanController_Delete(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	jabatanService := services.NewJabatanService(jabatanRepo)
	jabatanCtrl := NewJabatanController(jabatanService)

	jabatan := &models.Jabatan{Nama: "Manager", Deskripsi: "Manager Position"}
	jabatanService.Create(jabatan)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/api/jabatan/:id", jabatanCtrl.Delete)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/jabatan/%d", jabatan.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestJabatanController_Create_EmptyName(t *testing.T) {
	db := setupControllerTestDB(t)
	jabatanRepo := repositories.NewJabatanRepository(db)
	jabatanService := services.NewJabatanService(jabatanRepo)
	jabatanCtrl := NewJabatanController(jabatanService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/jabatan", jabatanCtrl.Create)

	payload := map[string]interface{}{
		"nama":      "",
		"deskripsi": "Manager Position",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/jabatan", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

