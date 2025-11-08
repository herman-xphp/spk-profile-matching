package dto

import "time"

// TargetProfileResponse represents target profile data in API response
type TargetProfileResponse struct {
	ID         uint                `json:"id"`
	JabatanID  uint                `json:"jabatan_id"`
	KriteriaID uint                `json:"kriteria_id"`
	TargetNilai float64            `json:"target_nilai"`
	Jabatan    *JabatanResponse    `json:"jabatan,omitempty"`
	Kriteria   *KriteriaResponse   `json:"kriteria,omitempty"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}

// TargetProfileCreateRequest represents target profile creation request
type TargetProfileCreateRequest struct {
	JabatanID  uint    `json:"jabatan_id" binding:"required"`
	KriteriaID uint    `json:"kriteria_id" binding:"required"`
	TargetNilai float64 `json:"target_nilai" binding:"required,min=0"`
}

// TargetProfileUpdateRequest represents target profile update request
type TargetProfileUpdateRequest struct {
	JabatanID  uint    `json:"jabatan_id,omitempty"`
	KriteriaID uint    `json:"kriteria_id,omitempty"`
	TargetNilai float64 `json:"target_nilai,omitempty" binding:"omitempty,min=0"`
}

