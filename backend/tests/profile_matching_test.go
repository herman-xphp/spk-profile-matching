package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/controllers"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProfileMatching(t *testing.T) {
	db := setupTestDB(t)

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
	profileMatchingCtrl := controllers.NewProfileMatchingController(profileMatchingSvc)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/calculate", profileMatchingCtrl.Calculate)

	// Setup test data using GORM
	jabatan := models.Jabatan{Nama: "Test Jabatan"}
	if err := db.Create(&jabatan).Error; err != nil {
		t.Fatalf("failed to create jabatan: %v", err)
	}

	// Create aspek first
	aspek := models.Aspek{Nama: "Test Aspek", Deskripsi: "Test", Persentase: 100.0}
	if err := db.Create(&aspek).Error; err != nil {
		t.Fatalf("failed to create aspek: %v", err)
	}

	k1 := models.Kriteria{AspekID: aspek.ID, Kode: "K1", Nama: "K1", IsCore: true, Bobot: 1}
	k2 := models.Kriteria{AspekID: aspek.ID, Kode: "K2", Nama: "K2", IsCore: false, Bobot: 1}
	if err := db.Create(&k1).Error; err != nil {
		t.Fatalf("failed to create k1: %v", err)
	}
	if err := db.Create(&k2).Error; err != nil {
		t.Fatalf("failed to create k2: %v", err)
	}

	tenaga := models.TenagaKerja{NIK: "TK001", Nama: "Test TK"}
	if err := db.Create(&tenaga).Error; err != nil {
		t.Fatalf("failed to create tenaga: %v", err)
	}

	// Insert test target profile
	tp1 := models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: k1.ID, TargetNilai: 4.0}
	tp2 := models.TargetProfile{JabatanID: jabatan.ID, KriteriaID: k2.ID, TargetNilai: 3.0}
	if err := db.Create(&tp1).Error; err != nil {
		t.Fatalf("failed to create tp1: %v", err)
	}
	if err := db.Create(&tp2).Error; err != nil {
		t.Fatalf("failed to create tp2: %v", err)
	}

	// Insert test nilai tenaga kerja
	nt1 := models.NilaiTenagaKerja{TenagaKerjaID: tenaga.ID, KriteriaID: k1.ID, Nilai: 3.5}
	nt2 := models.NilaiTenagaKerja{TenagaKerjaID: tenaga.ID, KriteriaID: k2.ID, Nilai: 4.0}
	if err := db.Create(&nt1).Error; err != nil {
		t.Fatalf("failed to create nt1: %v", err)
	}
	if err := db.Create(&nt2).Error; err != nil {
		t.Fatalf("failed to create nt2: %v", err)
	}

	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
		wantError  bool
	}{
		{
			name: "Calculate Profile Matching",
			payload: map[string]interface{}{
				"jabatan_id":       jabatan.ID,
				"tenaga_kerja_ids": []uint{tenaga.ID},
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name: "Invalid Jabatan ID",
			payload: map[string]interface{}{
				"jabatan_id": "invalid_id",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payloadBytes, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/calculate", bytes.NewBuffer(payloadBytes))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)

			if tt.wantError {
				var response map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &response)
				assert.Contains(t, response, "error")
			} else {
				// expecting an array of results
				var results []interface{}
				json.Unmarshal(w.Body.Bytes(), &results)
				assert.NotEmpty(t, results)
			}
		})
	}
}
