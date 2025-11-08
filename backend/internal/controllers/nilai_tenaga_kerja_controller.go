package controllers

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

type NilaiTenagaKerjaController struct {
	nilaiTenagaKerjaService *services.NilaiTenagaKerjaService
}

func NewNilaiTenagaKerjaController(nilaiTenagaKerjaService *services.NilaiTenagaKerjaService) *NilaiTenagaKerjaController {
	return &NilaiTenagaKerjaController{nilaiTenagaKerjaService: nilaiTenagaKerjaService}
}

func (ntkc *NilaiTenagaKerjaController) GetAll(c *gin.Context) {
	// Check if tenaga_kerja_id query param exists
	tenagaKerjaIDStr := c.Query("tenaga_kerja_id")
	if tenagaKerjaIDStr != "" {
		tenagaKerjaID, err := strconv.ParseUint(tenagaKerjaIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenaga_kerja_id format"})
			return
		}

		nilai, err := ntkc.nilaiTenagaKerjaService.GetByTenagaKerjaID(uint(tenagaKerjaID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch nilai tenaga kerja"})
			return
		}
		c.JSON(http.StatusOK, dto.MapNilaiTenagaKerjasToResponse(nilai))
		return
	}

	// Get all nilai
	nilai, err := ntkc.nilaiTenagaKerjaService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch nilai tenaga kerja"})
		return
	}
	c.JSON(http.StatusOK, dto.MapNilaiTenagaKerjasToResponse(nilai))
}

func (ntkc *NilaiTenagaKerjaController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	nilai, err := ntkc.nilaiTenagaKerjaService.GetByID(uint(id64))
	if err != nil {
		if err.Error() == "nilai tenaga kerja not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch nilai tenaga kerja"})
		return
	}

	c.JSON(http.StatusOK, nilai)
}

func (ntkc *NilaiTenagaKerjaController) Create(c *gin.Context) {
	var req dto.NilaiTenagaKerjaCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nilai := &models.NilaiTenagaKerja{
		TenagaKerjaID: req.TenagaKerjaID,
		KriteriaID:    req.KriteriaID,
		Nilai:         req.Nilai,
	}

	if err := ntkc.nilaiTenagaKerjaService.Create(nilai); err != nil {
		if err.Error() == "tenaga kerja not found" || err.Error() == "kriteria not found" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create nilai tenaga kerja"})
		return
	}

	c.JSON(http.StatusCreated, dto.MapNilaiTenagaKerjaToResponse(nilai))
}

func (ntkc *NilaiTenagaKerjaController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req dto.NilaiTenagaKerjaUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nilai := &models.NilaiTenagaKerja{
		TenagaKerjaID: req.TenagaKerjaID,
		KriteriaID:    req.KriteriaID,
		Nilai:         req.Nilai,
	}

	if err := ntkc.nilaiTenagaKerjaService.Update(uint(id64), nilai); err != nil {
		if err.Error() == "nilai tenaga kerja not found" || err.Error() == "tenaga kerja not found" || err.Error() == "kriteria not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update nilai tenaga kerja"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nilai tenaga kerja updated successfully"})
}

func (ntkc *NilaiTenagaKerjaController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := ntkc.nilaiTenagaKerjaService.Delete(uint(id64)); err != nil {
		if err.Error() == "nilai tenaga kerja not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete nilai tenaga kerja"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Nilai tenaga kerja deleted successfully"})
}
