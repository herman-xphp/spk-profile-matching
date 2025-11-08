package repositories

import (
	"testing"

	"backend/internal/models"
	"backend/pkg/database"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupRepositoryTestDB(t *testing.T) *gorm.DB {
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

func TestUserRepository_Create(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewUserRepository(db)

	user := &models.User{
		Email:    "test@example.com",
		Password: "hashedpassword",
		Nama:     "Test User",
		Role:     "user",
		IsActive: true,
	}

	err := repo.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
}

func TestUserRepository_GetByID(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewUserRepository(db)

	// Create a user
	user := &models.User{
		Email:    "test@example.com",
		Password: "hashedpassword",
		Nama:     "Test User",
	}
	repo.Create(user)

	// Get by ID
	found, err := repo.GetByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Nama, found.Nama)
}

func TestUserRepository_GetByID_NotFound(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewUserRepository(db)

	_, err := repo.GetByID(999)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUserRepository_GetAll(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewUserRepository(db)

	// Create multiple users
	user1 := &models.User{Email: "user1@example.com", Password: "pass", Nama: "User 1"}
	user2 := &models.User{Email: "user2@example.com", Password: "pass", Nama: "User 2"}
	repo.Create(user1)
	repo.Create(user2)

	users, err := repo.GetAll()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(users), 2)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewUserRepository(db)

	email := "test@example.com"
	user := &models.User{
		Email:    email,
		Password: "hashedpassword",
		Nama:     "Test User",
	}
	repo.Create(user)

	found, err := repo.FindByEmail(email)
	assert.NoError(t, err)
	assert.Equal(t, email, found.Email)
}

func TestUserRepository_Update(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewUserRepository(db)

	user := &models.User{
		Email:    "test@example.com",
		Password: "hashedpassword",
		Nama:     "Test User",
	}
	repo.Create(user)

	user.Nama = "Updated Name"
	err := repo.Update(user.ID, user)
	assert.NoError(t, err)

	updated, _ := repo.GetByID(user.ID)
	assert.Equal(t, "Updated Name", updated.Nama)
}

func TestUserRepository_Delete(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewUserRepository(db)

	user := &models.User{
		Email:    "test@example.com",
		Password: "hashedpassword",
		Nama:     "Test User",
	}
	repo.Create(user)

	err := repo.Delete(user.ID)
	assert.NoError(t, err)

	_, err = repo.GetByID(user.ID)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUserRepository_ExistsByEmail(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewUserRepository(db)

	email := "test@example.com"
	exists, err := repo.ExistsByEmail(email)
	assert.NoError(t, err)
	assert.False(t, exists)

	user := &models.User{Email: email, Password: "pass", Nama: "Test"}
	repo.Create(user)

	exists, err = repo.ExistsByEmail(email)
	assert.NoError(t, err)
	assert.True(t, exists)
}
