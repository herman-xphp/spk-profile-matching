package services

import (
	"testing"

	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestTenagaKerjaService_Create(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewTenagaKerjaRepository(db)
	service := NewTenagaKerjaService(repo)

	tenagaKerja := &models.TenagaKerja{
		NIK:  "TK001",
		Nama: "John Doe",
	}

	err := service.Create(tenagaKerja)
	assert.NoError(t, err)
	assert.NotZero(t, tenagaKerja.ID)
}

func TestTenagaKerjaService_Create_DuplicateNIK(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewTenagaKerjaRepository(db)
	service := NewTenagaKerjaService(repo)

	tenagaKerja1 := &models.TenagaKerja{
		NIK:  "TK001",
		Nama: "John Doe",
	}
	service.Create(tenagaKerja1)

	tenagaKerja2 := &models.TenagaKerja{
		NIK:  "TK001",
		Nama: "Jane Doe",
	}

	err := service.Create(tenagaKerja2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "NIK sudah terdaftar")
}

func TestTenagaKerjaService_Create_EmptyNIK(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewTenagaKerjaRepository(db)
	service := NewTenagaKerjaService(repo)

	tenagaKerja := &models.TenagaKerja{
		NIK:  "",
		Nama: "John Doe",
	}

	err := service.Create(tenagaKerja)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "NIK tidak boleh kosong")
}

func TestTenagaKerjaService_GetByID(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewTenagaKerjaRepository(db)
	service := NewTenagaKerjaService(repo)

	tenagaKerja := &models.TenagaKerja{
		NIK:  "TK001",
		Nama: "John Doe",
	}
	service.Create(tenagaKerja)

	found, err := service.GetByID(tenagaKerja.ID)
	assert.NoError(t, err)
	assert.Equal(t, tenagaKerja.NIK, found.NIK)
}

func TestTenagaKerjaService_Update(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewTenagaKerjaRepository(db)
	service := NewTenagaKerjaService(repo)

	tenagaKerja := &models.TenagaKerja{
		NIK:  "TK001",
		Nama: "John Doe",
	}
	service.Create(tenagaKerja)

	tenagaKerja.Nama = "Jane Doe"
	err := service.Update(tenagaKerja.ID, tenagaKerja)
	assert.NoError(t, err)

	updated, _ := service.GetByID(tenagaKerja.ID)
	assert.Equal(t, "Jane Doe", updated.Nama)
}

func TestTenagaKerjaService_Delete(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewTenagaKerjaRepository(db)
	service := NewTenagaKerjaService(repo)

	tenagaKerja := &models.TenagaKerja{
		NIK:  "TK001",
		Nama: "John Doe",
	}
	service.Create(tenagaKerja)

	err := service.Delete(tenagaKerja.ID)
	assert.NoError(t, err)

	_, err = service.GetByID(tenagaKerja.ID)
	assert.Error(t, err)
}

