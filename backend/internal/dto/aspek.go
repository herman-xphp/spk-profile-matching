package dto

import "time"

// AspekResponse represents aspek data in API response
type AspekResponse struct {
	ID         uint      `json:"id"`
	Nama       string    `json:"nama"`
	Deskripsi  string    `json:"deskripsi"`
	Persentase float64   `json:"persentase"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// AspekCreateRequest represents aspek creation request
type AspekCreateRequest struct {
	Nama       string  `json:"nama" binding:"required"`
	Deskripsi  string  `json:"deskripsi,omitempty"`
	Persentase float64 `json:"persentase" binding:"required,min=0,max=100"`
}

// AspekUpdateRequest represents aspek update request
type AspekUpdateRequest struct {
	Nama       string  `json:"nama,omitempty"`
	Deskripsi  string  `json:"deskripsi,omitempty"`
	Persentase float64 `json:"persentase,omitempty" binding:"omitempty,min=0,max=100"`
}
