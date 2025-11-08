package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthController_Login(t *testing.T) {
	db := setupControllerTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	authSvc := services.NewAuthService(userRepo)
	authCtrl := NewAuthController(authSvc)

	// Set SECRET_KEY for JWT
	os.Setenv("SECRET_KEY", "test-secret-key")

	// Create test user
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{
		Email:    "test@example.com",
		Password: string(hashedPassword),
		Nama:     "Test User",
		Role:     "user",
		IsActive: true,
	}
	userRepo.Create(user)

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
				"email":    "test@example.com",
				"password": "password123",
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name: "Invalid Password",
			payload: map[string]interface{}{
				"email":    "test@example.com",
				"password": "wrongpassword",
			},
			wantStatus: http.StatusUnauthorized,
			wantError:  true,
		},
		{
			name: "Invalid Email",
			payload: map[string]interface{}{
				"email":    "nonexistent@example.com",
				"password": "password123",
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

