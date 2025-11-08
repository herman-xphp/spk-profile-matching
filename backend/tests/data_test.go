package tests

import (
	"testing"
	"time"

	"backend/internal/models"

	"gorm.io/gorm"
)

func TestJabatan(t *testing.T, db *gorm.DB) uint {
	db.Exec("DELETE FROM jabatans")

	jabatan := models.Jabatan{
		Nama:      "Test Jabatan",
		Deskripsi: "Test Deskripsi",
	}

	if err := db.Create(&jabatan).Error; err != nil {
		t.Fatalf("Failed to seed test jabatan: %v", err)
	}

	return jabatan.ID
}

func TestAspek(t *testing.T, db *gorm.DB) uint {
	db.Exec("DELETE FROM aspeks")

	aspek := models.Aspek{
		Nama:       "Test Aspek",
		Persentase: 50.0,
	}

	if err := db.Create(&aspek).Error; err != nil {
		t.Fatalf("Failed to seed test aspek: %v", err)
	}

	return aspek.ID
}

func TestKriteria(t *testing.T, db *gorm.DB, aspekID uint) uint {
	db.Exec("DELETE FROM kriterias")

	kriteria := models.Kriteria{
		AspekID: aspekID,
		Kode:    "TK-001",
		Nama:    "Test Kriteria",
		IsCore:  true,
		Bobot:   30.0,
	}

	if err := db.Create(&kriteria).Error; err != nil {
		t.Fatalf("Failed to seed test kriteria: %v", err)
	}

	return kriteria.ID
}

func TestTenagaKerja(t *testing.T, db *gorm.DB) uint {
	db.Exec("DELETE FROM tenaga_kerjas")

	tenagaKerja := models.TenagaKerja{
		NIK:      "123456789",
		Nama:     "Test Tenaga Kerja",
		TglLahir: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Alamat:   "Test Alamat",
		Telepon:  "081234567890",
	}

	if err := db.Create(&tenagaKerja).Error; err != nil {
		t.Fatalf("Failed to seed test tenaga kerja: %v", err)
	}

	return tenagaKerja.ID
}

func TestNilaiTenagaKerja(t *testing.T, db *gorm.DB, tenagaKerjaID, kriteriaID uint, nilai float64) {
	db.Where("tenaga_kerja_id = ? AND kriteria_id = ?", tenagaKerjaID, kriteriaID).Delete(&models.NilaiTenagaKerja{})

	nilaiTenagaKerja := models.NilaiTenagaKerja{
		TenagaKerjaID: tenagaKerjaID,
		KriteriaID:    kriteriaID,
		Nilai:         nilai,
	}

	if err := db.Create(&nilaiTenagaKerja).Error; err != nil {
		t.Fatalf("Failed to seed test nilai tenaga kerja: %v", err)
	}
}

func TestTargetProfile(t *testing.T, db *gorm.DB, jabatanID, kriteriaID uint, targetNilai float64) {
	db.Where("jabatan_id = ? AND kriteria_id = ?", jabatanID, kriteriaID).Delete(&models.TargetProfile{})

	targetProfile := models.TargetProfile{
		JabatanID:   jabatanID,
		KriteriaID:  kriteriaID,
		TargetNilai: targetNilai,
	}

	if err := db.Create(&targetProfile).Error; err != nil {
		t.Fatalf("Failed to seed test target profile: %v", err)
	}
}
