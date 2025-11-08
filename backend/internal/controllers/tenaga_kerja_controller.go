package controllers

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

type TenagaKerjaController struct {
	tenagaKerjaService *services.TenagaKerjaService
}

func NewTenagaKerjaController(tenagaKerjaService *services.TenagaKerjaService) *TenagaKerjaController {
	return &TenagaKerjaController{tenagaKerjaService: tenagaKerjaService}
}

func (tkc *TenagaKerjaController) GetAll(c *gin.Context) {
	tenagaKerja, err := tkc.tenagaKerjaService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tenaga kerja"})
		return
	}
	c.JSON(http.StatusOK, dto.MapTenagaKerjasToResponse(tenagaKerja))
}

func (tkc *TenagaKerjaController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	tenagaKerja, err := tkc.tenagaKerjaService.GetByID(uint(id64))
	if err != nil {
		if err.Error() == "tenaga kerja not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tenaga kerja"})
		return
	}

	c.JSON(http.StatusOK, dto.MapTenagaKerjaToResponse(tenagaKerja))
}

func (tkc *TenagaKerjaController) Create(c *gin.Context) {
	var req dto.TenagaKerjaCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenagaKerja := &models.TenagaKerja{
		NIK:      req.NIK,
		Nama:     req.Nama,
		TglLahir: req.TglLahir.Time(),
		Alamat:   req.Alamat,
		Telepon:  req.Telepon,
	}

	if err := tkc.tenagaKerjaService.Create(tenagaKerja); err != nil {
		if err.Error() == "NIK sudah terdaftar" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.MapTenagaKerjaToResponse(tenagaKerja))
}

func (tkc *TenagaKerjaController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req dto.TenagaKerjaUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenagaKerja := &models.TenagaKerja{
		NIK:      req.NIK,
		Nama:     req.Nama,
		TglLahir: req.TglLahir.Time(),
		Alamat:   req.Alamat,
		Telepon:  req.Telepon,
	}

	if err := tkc.tenagaKerjaService.Update(uint(id64), tenagaKerja); err != nil {
		if err.Error() == "tenaga kerja not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "NIK sudah terdaftar" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update tenaga kerja"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tenaga kerja updated successfully"})
}

func (tkc *TenagaKerjaController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := tkc.tenagaKerjaService.Delete(uint(id64)); err != nil {
		if err.Error() == "tenaga kerja not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete tenaga kerja"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tenaga kerja deleted successfully"})
}
