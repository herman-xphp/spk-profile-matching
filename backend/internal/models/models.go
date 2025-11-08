package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`
	Nama     string `gorm:"type:varchar(100)" json:"nama"`
	Role     string `gorm:"type:enum('admin','user');default:'user'" json:"role"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
}

type Jabatan struct {
	gorm.Model
	Nama      string `gorm:"type:varchar(100);not null" json:"nama"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}

type Aspek struct {
	gorm.Model
	Nama       string  `gorm:"type:varchar(100);not null" json:"nama"`
	Deskripsi  string  `gorm:"type:text" json:"deskripsi"`
	Persentase float64 `gorm:"type:decimal(5,2);not null" json:"persentase"`
}

type Kriteria struct {
	gorm.Model
	AspekID uint    `gorm:"not null" json:"aspek_id"`
	Kode    string  `gorm:"type:varchar(20);not null" json:"kode"`
	Nama    string  `gorm:"type:varchar(100);not null" json:"nama"`
	IsCore  bool    `gorm:"default:false" json:"is_core"`
	Bobot   float64 `gorm:"type:decimal(5,2);not null" json:"bobot"`
	Aspek   Aspek   `gorm:"foreignKey:AspekID" json:"aspek,omitempty"`
}

type TargetProfile struct {
	gorm.Model
	JabatanID   uint     `gorm:"not null" json:"jabatan_id"`
	KriteriaID  uint     `gorm:"not null" json:"kriteria_id"`
	TargetNilai float64  `gorm:"type:decimal(5,2);not null" json:"target_nilai"`
	Jabatan     Jabatan  `gorm:"foreignKey:JabatanID" json:"jabatan,omitempty"`
	Kriteria    Kriteria `gorm:"foreignKey:KriteriaID" json:"kriteria,omitempty"`
}

type TenagaKerja struct {
	gorm.Model
	NIK      string    `gorm:"type:varchar(20);unique;not null" json:"nik"`
	Nama     string    `gorm:"type:varchar(100);not null" json:"nama"`
	TglLahir time.Time `gorm:"type:date" json:"tgl_lahir"`
	Alamat   string    `gorm:"type:text" json:"alamat"`
	Telepon  string    `gorm:"type:varchar(20)" json:"telepon"`
}

type NilaiTenagaKerja struct {
	gorm.Model
	TenagaKerjaID uint        `gorm:"not null" json:"tenaga_kerja_id"`
	KriteriaID    uint        `gorm:"not null" json:"kriteria_id"`
	Nilai         float64     `gorm:"type:decimal(5,2);not null" json:"nilai"`
	TenagaKerja   TenagaKerja `gorm:"foreignKey:TenagaKerjaID" json:"tenaga_kerja,omitempty"`
	Kriteria      Kriteria    `gorm:"foreignKey:KriteriaID" json:"kriteria,omitempty"`
}

type ProfileMatchResult struct {
	gorm.Model
	TenagaKerjaID   uint        `gorm:"not null" json:"tenaga_kerja_id"`
	JabatanID       uint        `gorm:"not null" json:"jabatan_id"`
	TotalScore      float64     `gorm:"type:decimal(5,2);not null" json:"total_score"`
	CoreFactor      float64     `gorm:"type:decimal(5,2);not null" json:"core_factor"`
	SecondaryFactor float64     `gorm:"type:decimal(5,2);not null" json:"secondary_factor"`
	TenagaKerja     TenagaKerja `gorm:"foreignKey:TenagaKerjaID" json:"tenaga_kerja,omitempty"`
	Jabatan         Jabatan     `gorm:"foreignKey:JabatanID" json:"jabatan,omitempty"`
}
