package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/controllers"
	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	db := setupTestDB(t)
	seedTestUser(t, db)

	// Setup controller
	userRepo := repositories.NewUserRepository(db)
	authSvc := services.NewAuthService(userRepo)
	authCtrl := controllers.NewAuthController(authSvc)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/auth/login", authCtrl.Login)

	tests := []struct {
		name       string
		payload    map[string]interface{}
		wantStatus int
		wantError  bool
	}{
		{
			name: "Valid Login",
			payload: map[string]interface{}{
				"email":    "admin@kpsggroup.com",
				"password": "admin123",
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name: "Invalid Password",
			payload: map[string]interface{}{
				"email":    "admin@kpsggroup.com",
				"password": "wrongpassword",
			},
			wantStatus: http.StatusUnauthorized,
			wantError:  true,
		},
		{
			name: "Invalid Email Format",
			payload: map[string]interface{}{
				"email":    "notanemail",
				"password": "password123",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payloadBytes, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(payloadBytes))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)

			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)

			if tt.wantError {
				assert.Contains(t, response, "error")
			} else {
				assert.Contains(t, response, "token")
				assert.Contains(t, response, "user")
			}
		})
	}
}
