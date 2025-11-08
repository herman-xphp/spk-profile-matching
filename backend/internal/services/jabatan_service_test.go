package services

import (
	"testing"

	"backend/internal/models"
	"backend/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestJabatanService_Create(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewJabatanRepository(db)
	service := NewJabatanService(repo)

	jabatan := &models.Jabatan{
		Nama:      "Manager",
		Deskripsi: "Manager Position",
	}

	err := service.Create(jabatan)
	assert.NoError(t, err)
	assert.NotZero(t, jabatan.ID)
}

func TestJabatanService_Create_EmptyName(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewJabatanRepository(db)
	service := NewJabatanService(repo)

	jabatan := &models.Jabatan{
		Nama:      "",
		Deskripsi: "Manager Position",
	}

	err := service.Create(jabatan)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nama jabatan tidak boleh kosong")
}

func TestJabatanService_GetByID(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewJabatanRepository(db)
	service := NewJabatanService(repo)

	jabatan := &models.Jabatan{
		Nama:      "Manager",
		Deskripsi: "Manager Position",
	}
	service.Create(jabatan)

	found, err := service.GetByID(jabatan.ID)
	assert.NoError(t, err)
	assert.Equal(t, jabatan.Nama, found.Nama)
}

func TestJabatanService_GetByID_NotFound(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewJabatanRepository(db)
	service := NewJabatanService(repo)

	_, err := service.GetByID(999)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestJabatanService_Update(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewJabatanRepository(db)
	service := NewJabatanService(repo)

	jabatan := &models.Jabatan{
		Nama:      "Manager",
		Deskripsi: "Manager Position",
	}
	service.Create(jabatan)

	jabatan.Nama = "Senior Manager"
	err := service.Update(jabatan.ID, jabatan)
	assert.NoError(t, err)

	updated, _ := service.GetByID(jabatan.ID)
	assert.Equal(t, "Senior Manager", updated.Nama)
}

func TestJabatanService_Delete(t *testing.T) {
	db := setupServiceTestDB(t)
	repo := repositories.NewJabatanRepository(db)
	service := NewJabatanService(repo)

	jabatan := &models.Jabatan{
		Nama:      "Manager",
		Deskripsi: "Manager Position",
	}
	service.Create(jabatan)

	err := service.Delete(jabatan.ID)
	assert.NoError(t, err)

	_, err = service.GetByID(jabatan.ID)
	assert.Error(t, err)
}

