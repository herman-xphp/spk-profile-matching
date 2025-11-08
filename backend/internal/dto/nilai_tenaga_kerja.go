package dto

import "time"

// NilaiTenagaKerjaResponse represents nilai tenaga kerja data in API response
type NilaiTenagaKerjaResponse struct {
	ID           uint                  `json:"id"`
	TenagaKerjaID uint                 `json:"tenaga_kerja_id"`
	KriteriaID   uint                  `json:"kriteria_id"`
	Nilai        float64               `json:"nilai"`
	TenagaKerja  *TenagaKerjaResponse  `json:"tenaga_kerja,omitempty"`
	Kriteria     *KriteriaResponse     `json:"kriteria,omitempty"`
	CreatedAt    time.Time             `json:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at"`
}

// NilaiTenagaKerjaCreateRequest represents nilai tenaga kerja creation request
type NilaiTenagaKerjaCreateRequest struct {
	TenagaKerjaID uint    `json:"tenaga_kerja_id" binding:"required"`
	KriteriaID   uint    `json:"kriteria_id" binding:"required"`
	Nilai        float64 `json:"nilai" binding:"required,min=0"`
}

// NilaiTenagaKerjaUpdateRequest represents nilai tenaga kerja update request
type NilaiTenagaKerjaUpdateRequest struct {
	TenagaKerjaID uint    `json:"tenaga_kerja_id,omitempty"`
	KriteriaID   uint    `json:"kriteria_id,omitempty"`
	Nilai        float64 `json:"nilai,omitempty" binding:"omitempty,min=0"`
}

