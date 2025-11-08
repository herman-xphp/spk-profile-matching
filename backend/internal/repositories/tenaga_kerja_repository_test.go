package repositories

import (
	"testing"

	"backend/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestTenagaKerjaRepository_Create(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewTenagaKerjaRepository(db)

	tenagaKerja := &models.TenagaKerja{
		NIK:  "TK001",
		Nama: "John Doe",
	}

	err := repo.Create(tenagaKerja)
	assert.NoError(t, err)
	assert.NotZero(t, tenagaKerja.ID)
}

func TestTenagaKerjaRepository_GetByID(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewTenagaKerjaRepository(db)

	tenagaKerja := &models.TenagaKerja{
		NIK:  "TK001",
		Nama: "John Doe",
	}
	repo.Create(tenagaKerja)

	found, err := repo.GetByID(tenagaKerja.ID)
	assert.NoError(t, err)
	assert.Equal(t, tenagaKerja.NIK, found.NIK)
}

func TestTenagaKerjaRepository_ExistsByNIK(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewTenagaKerjaRepository(db)

	nik := "TK001"
	exists, err := repo.ExistsByNIK(nik)
	assert.NoError(t, err)
	assert.False(t, exists)

	tenagaKerja := &models.TenagaKerja{NIK: nik, Nama: "John Doe"}
	repo.Create(tenagaKerja)

	exists, err = repo.ExistsByNIK(nik)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestTenagaKerjaRepository_GetByIDs(t *testing.T) {
	db := setupRepositoryTestDB(t)
	repo := NewTenagaKerjaRepository(db)

	tenagaKerja1 := &models.TenagaKerja{NIK: "TK001", Nama: "John Doe"}
	tenagaKerja2 := &models.TenagaKerja{NIK: "TK002", Nama: "Jane Doe"}
	repo.Create(tenagaKerja1)
	repo.Create(tenagaKerja2)

	tenagaKerjas, err := repo.GetByIDs([]uint{tenagaKerja1.ID, tenagaKerja2.ID})
	assert.NoError(t, err)
	assert.Len(t, tenagaKerjas, 2)
}

