package repositories

import (
	"testing"

	"backend/internal/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestJabatanRepository_Create(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewJabatanRepository(db)

	jabatan := &models.Jabatan{
		Nama:      "Manager",
		Deskripsi: "Manager Position",
	}

	err := repo.Create(jabatan)
	assert.NoError(t, err)
	assert.NotZero(t, jabatan.ID)
}

func TestJabatanRepository_GetByID(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewJabatanRepository(db)

	jabatan := &models.Jabatan{
		Nama:      "Manager",
		Deskripsi: "Manager Position",
	}
	repo.Create(jabatan)

	found, err := repo.GetByID(jabatan.ID)
	assert.NoError(t, err)
	assert.Equal(t, jabatan.Nama, found.Nama)
}

func TestJabatanRepository_GetByID_NotFound(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewJabatanRepository(db)

	_, err := repo.GetByID(999)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestJabatanRepository_GetAll(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewJabatanRepository(db)

	jabatan1 := &models.Jabatan{Nama: "Manager", Deskripsi: "Manager Position"}
	jabatan2 := &models.Jabatan{Nama: "Staff", Deskripsi: "Staff Position"}
	repo.Create(jabatan1)
	repo.Create(jabatan2)

	jabatans, err := repo.GetAll()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(jabatans), 2)
}

func TestJabatanRepository_Update(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewJabatanRepository(db)

	jabatan := &models.Jabatan{
		Nama:      "Manager",
		Deskripsi: "Manager Position",
	}
	repo.Create(jabatan)

	jabatan.Nama = "Senior Manager"
	err := repo.Update(jabatan.ID, jabatan)
	assert.NoError(t, err)

	updated, _ := repo.GetByID(jabatan.ID)
	assert.Equal(t, "Senior Manager", updated.Nama)
}

func TestJabatanRepository_Delete(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewJabatanRepository(db)

	jabatan := &models.Jabatan{
		Nama:      "Manager",
		Deskripsi: "Manager Position",
	}
	repo.Create(jabatan)

	err := repo.Delete(jabatan.ID)
	assert.NoError(t, err)

	_, err = repo.GetByID(jabatan.ID)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

