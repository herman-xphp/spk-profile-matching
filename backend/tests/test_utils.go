package tests

import (
	"testing"

	"backend/internal/models"
	"backend/pkg/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// setupTestDB connects to a MySQL test database using ConnectTestDB
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := database.ConnectTestDB(t)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Clear all tables before test
	if err := database.CleanTestDB(db); err != nil {
		t.Fatalf("Failed to clean test database: %v", err)
	}

	return db
}

func seedTestUser(t *testing.T, db *gorm.DB) {
	// Clear existing users
	db.Exec("DELETE FROM users")

	// Create hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	user := models.User{
		Email:    "admin@kpsggroup.com",
		Password: string(hashedPassword),
		Role:     "admin",
		IsActive: true,
	}

	result := db.Create(&user)
	if result.Error != nil {
		t.Fatalf("Failed to seed test user: %v", result.Error)
	}
}
