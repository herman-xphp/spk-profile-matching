package database

import (
	"fmt"
	"testing"

	"backend/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectTestDB connects to the test database
// It uses TEST_DB_NAME environment variable if available, otherwise uses DB_NAME_test
func ConnectTestDB(t *testing.T) (*gorm.DB, error) {
	config := GetTestDBConfig()
	dsn := config.BuildDSN()

	// Use silent logger for tests
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to test database: %v", err)
	}

	// Auto migrate test database
	err = db.AutoMigrate(
		&models.User{},
		&models.Jabatan{},
		&models.Aspek{},
		&models.Kriteria{},
		&models.TargetProfile{},
		&models.TenagaKerja{},
		&models.NilaiTenagaKerja{},
		&models.ProfileMatchResult{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate test database: %v", err)
	}

	return db, nil
}

// CleanTestDB clears all data from test database tables
func CleanTestDB(db *gorm.DB) error {
	tables := []string{
		"profile_match_results",
		"nilai_tenaga_kerjas",
		"target_profiles",
		"tenaga_kerjas",
		"kriterias",
		"aspeks",
		"jabatans",
		"users",
	}

	for _, table := range tables {
		if err := db.Exec(fmt.Sprintf("DELETE FROM %s", table)).Error; err != nil {
			return fmt.Errorf("failed to clean table %s: %v", table, err)
		}
		// Reset AUTO_INCREMENT
		db.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", table))
	}

	return nil
}
