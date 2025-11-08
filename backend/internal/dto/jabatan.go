package dto

import "time"

// JabatanResponse represents jabatan data in API response
type JabatanResponse struct {
	ID        uint      `json:"id"`
	Nama      string    `json:"nama"`
	Deskripsi string    `json:"deskripsi"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// JabatanCreateRequest represents jabatan creation request
type JabatanCreateRequest struct {
	Nama      string `json:"nama" binding:"required"`
	Deskripsi string `json:"deskripsi,omitempty"`
}

// JabatanUpdateRequest represents jabatan update request
type JabatanUpdateRequest struct {
	Nama      string `json:"nama,omitempty"`
	Deskripsi string `json:"deskripsi,omitempty"`
}

