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
	"backend/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupControllerTestDB(t *testing.T) *gorm.DB {
	db, err := database.ConnectTestDB(t)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Clear all tables
	if err := database.CleanTestDB(db); err != nil {
		t.Fatalf("Failed to clean test database: %v", err)
	}

	return db
}

func TestUserController_GetAll(t *testing.T) {
	db := setupControllerTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userCtrl := NewUserController(userService)

	// Create test users
	user1 := &models.User{Email: "user1@example.com", Password: "pass", Nama: "User 1"}
	user2 := &models.User{Email: "user2@example.com", Password: "pass", Nama: "User 2"}
	userService.Create(user1)
	userService.Create(user2)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/users", userCtrl.GetAll)

	req := httptest.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.GreaterOrEqual(t, len(response), 2)
}

func TestUserController_GetByID(t *testing.T) {
	db := setupControllerTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userCtrl := NewUserController(userService)

	user := &models.User{Email: "test@example.com", Password: "pass", Nama: "Test User"}
	userService.Create(user)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/users/:id", userCtrl.GetByID)

	req := httptest.NewRequest("GET", "/api/users/"+fmt.Sprintf("%d", user.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, float64(user.ID), response["id"])
	assert.Equal(t, "test@example.com", response["email"])
}

func TestUserController_Create(t *testing.T) {
	db := setupControllerTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userCtrl := NewUserController(userService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/users", userCtrl.Create)

	payload := map[string]interface{}{
		"email":    "newuser@example.com",
		"password": "password123",
		"nama":     "New User",
		"role":     "user",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "newuser@example.com", response["email"])
	assert.NotContains(t, response, "password")
}

func TestUserController_Update(t *testing.T) {
	db := setupControllerTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userCtrl := NewUserController(userService)

	user := &models.User{Email: "test@example.com", Password: "pass", Nama: "Test User"}
	userService.Create(user)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.PUT("/api/users/:id", userCtrl.Update)

	payload := map[string]interface{}{
		"nama": "Updated Name",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("PUT", "/api/users/"+fmt.Sprintf("%d", user.ID), bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_Delete(t *testing.T) {
	db := setupControllerTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userCtrl := NewUserController(userService)

	user := &models.User{Email: "test@example.com", Password: "pass", Nama: "Test User"}
	userService.Create(user)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/api/users/:id", userCtrl.Delete)

	req := httptest.NewRequest("DELETE", "/api/users/"+fmt.Sprintf("%d", user.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_Register(t *testing.T) {
	db := setupControllerTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userCtrl := NewUserController(userService)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/auth/register", userCtrl.Register)

	payload := map[string]interface{}{
		"email":    "newuser@example.com",
		"password": "password123",
		"nama":     "New User",
	}

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "newuser@example.com", response["email"])
}

