package services

import (
	"testing"

	"backend/internal/models"
	"backend/internal/repositories"
	"backend/pkg/database"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func setupServiceTestDB(t *testing.T) *gorm.DB {
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

func TestUserService_Create(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewUserRepository(db)
	service := NewUserService(repo)

	user := &models.User{
		Email:    "test@example.com",
		Password: "password123",
		Nama:     "Test User",
		Role:     "user",
	}

	err := service.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.NotEqual(t, "password123", user.Password) // Password should be hashed
}

func TestUserService_Create_DuplicateEmail(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewUserRepository(db)
	service := NewUserService(repo)

	user1 := &models.User{
		Email:    "test@example.com",
		Password: "password123",
		Nama:     "Test User 1",
	}
	service.Create(user1)

	user2 := &models.User{
		Email:    "test@example.com",
		Password: "password123",
		Nama:     "Test User 2",
	}

	err := service.Create(user2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "email already exists")
}

func TestUserService_GetByID(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewUserRepository(db)
	service := NewUserService(repo)

	user := &models.User{
		Email:    "test@example.com",
		Password: "password123",
		Nama:     "Test User",
	}
	service.Create(user)

	found, err := service.GetByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, "", found.Password) // Password should be empty in response
}

func TestUserService_GetAll(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewUserRepository(db)
	service := NewUserService(repo)

	user1 := &models.User{Email: "user1@example.com", Password: "pass", Nama: "User 1"}
	user2 := &models.User{Email: "user2@example.com", Password: "pass", Nama: "User 2"}
	service.Create(user1)
	service.Create(user2)

	users, err := service.GetAll()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(users), 2)

	// Check that all passwords are empty
	for _, u := range users {
		assert.Equal(t, "", u.Password)
	}
}

func TestUserService_Update(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewUserRepository(db)
	service := NewUserService(repo)

	user := &models.User{
		Email:    "test@example.com",
		Password: "password123",
		Nama:     "Test User",
	}
	service.Create(user)

	user.Nama = "Updated Name"
	err := service.Update(user.ID, user)
	assert.NoError(t, err)

	updated, _ := service.GetByID(user.ID)
	assert.Equal(t, "Updated Name", updated.Nama)
}

func TestUserService_Update_NotFound(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewUserRepository(db)
	service := NewUserService(repo)

	user := &models.User{
		Email:    "test@example.com",
		Password: "password123",
		Nama:     "Test User",
	}

	err := service.Update(999, user)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestUserService_Delete(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewUserRepository(db)
	service := NewUserService(repo)

	user := &models.User{
		Email:    "test@example.com",
		Password: "password123",
		Nama:     "Test User",
	}
	service.Create(user)

	err := service.Delete(user.ID)
	assert.NoError(t, err)

	_, err = service.GetByID(user.ID)
	assert.Error(t, err)
}

func TestUserService_Register(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewUserRepository(db)
	service := NewUserService(repo)

	user := &models.User{
		Email:    "test@example.com",
		Password: "password123",
		Nama:     "Test User",
	}

	err := service.Register(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
	assert.True(t, user.IsActive)
	assert.Equal(t, "user", user.Role) // Default role
}

func TestAuthService_Authenticate(t *testing.T) {
	db := setupServiceTestDB(t)
	userRepo := repositories.NewUserRepository(db)
	authService := NewAuthService(userRepo)

	// Create user with hashed password
	password := "password123"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{
		Email:    "test@example.com",
		Password: string(hashedPassword),
		Nama:     "Test User",
	}
	userRepo.Create(user)

	// Test valid credentials
	authenticated, err := authService.Authenticate("test@example.com", password)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, authenticated.Email)
	assert.Equal(t, "", authenticated.Password) // Password should be empty

	// Test invalid password
	_, err = authService.Authenticate("test@example.com", "wrongpassword")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid credentials")

	// Test invalid email
	_, err = authService.Authenticate("nonexistent@example.com", password)
	assert.Error(t, err)
}
