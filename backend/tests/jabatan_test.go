package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"backend/internal/controllers"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateJabatan(t *testing.T) {
	db := setupTestDB(t)

	// Setup controller
	jabatanRepo := repositories.NewJabatanRepository(db)
	jabatanSvc := services.NewJabatanService(jabatanRepo)
	jabatanCtrl := controllers.NewJabatanController(jabatanSvc)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/jabatan", jabatanCtrl.Create)

	tests := []struct {
		name       string
		payload    models.Jabatan
		wantStatus int
		wantError  bool
	}{
		{
			name: "Valid Jabatan",
			payload: models.Jabatan{
				Nama:      "Test Jabatan",
				Deskripsi: "Test Deskripsi Jabatan",
			},
			wantStatus: http.StatusCreated,
			wantError:  false,
		},
		{
			name: "Invalid Jabatan - Empty Name",
			payload: models.Jabatan{
				Nama:      "",
				Deskripsi: "Test Deskripsi",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payloadBytes, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/jabatan", bytes.NewBuffer(payloadBytes))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)

			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)

			if tt.wantError {
				assert.Contains(t, response, "error")
			} else {
				assert.Contains(t, response, "id")
				assert.Equal(t, tt.payload.Nama, response["nama"])
			}
		})
	}
}

func TestGetJabatan(t *testing.T) {
	db := setupTestDB(t)

	// Setup controller
	jabatanRepo := repositories.NewJabatanRepository(db)
	jabatanSvc := services.NewJabatanService(jabatanRepo)
	jabatanCtrl := controllers.NewJabatanController(jabatanSvc)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/jabatan", jabatanCtrl.GetAll)
	router.GET("/api/jabatan/:id", jabatanCtrl.GetByID)

	// Create test jabatan
	testJabatan := models.Jabatan{
		Nama:      "Test Jabatan",
		Deskripsi: "Test Deskripsi",
	}
	if err := db.Create(&testJabatan).Error; err != nil {
		t.Fatalf("failed to seed jabatan: %v", err)
	}

	t.Run("Get All Jabatan", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/jabatan", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Jabatan
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.NotEmpty(t, response)
	})

	t.Run("Get Jabatan By ID", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/jabatan/"+strconv.FormatUint(uint64(testJabatan.ID), 10), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Jabatan
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, testJabatan.Nama, response.Nama)
	})

	t.Run("Get Jabatan - Invalid ID", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/jabatan/invalidid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
