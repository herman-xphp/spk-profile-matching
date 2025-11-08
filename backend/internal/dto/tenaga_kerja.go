package dto

import "time"

// TenagaKerjaResponse represents tenaga kerja data in API response
type TenagaKerjaResponse struct {
	ID        uint      `json:"id"`
	NIK       string    `json:"nik"`
	Nama      string    `json:"nama"`
	TglLahir  time.Time `json:"tgl_lahir"`
	Alamat    string    `json:"alamat"`
	Telepon   string    `json:"telepon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TenagaKerjaCreateRequest represents tenaga kerja creation request
type TenagaKerjaCreateRequest struct {
	NIK      string   `json:"nik" binding:"required"`
	Nama     string   `json:"nama" binding:"required"`
	TglLahir DateOnly `json:"tgl_lahir" binding:"required"`
	Alamat   string   `json:"alamat" binding:"required"`
	Telepon  string   `json:"telepon,omitempty"`
}

// TenagaKerjaUpdateRequest represents tenaga kerja update request
type TenagaKerjaUpdateRequest struct {
	NIK      string   `json:"nik,omitempty"`
	Nama     string   `json:"nama,omitempty"`
	TglLahir DateOnly `json:"tgl_lahir,omitempty"`
	Alamat   string   `json:"alamat,omitempty"`
	Telepon  string   `json:"telepon,omitempty"`
}
