package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"backend/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	config := GetDBConfig()
	dsn := config.BuildDSN()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}

	// Auto migrate
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
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	DB = db
	return db, nil
}
