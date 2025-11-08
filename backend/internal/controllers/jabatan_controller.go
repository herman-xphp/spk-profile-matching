package controllers

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

type JabatanController struct {
	jabatanService *services.JabatanService
}

func NewJabatanController(jabatanService *services.JabatanService) *JabatanController {
	return &JabatanController{jabatanService: jabatanService}
}

func (jc *JabatanController) GetAll(c *gin.Context) {
	jabatan, err := jc.jabatanService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch jabatan"})
		return
	}
	c.JSON(http.StatusOK, dto.MapJabatansToResponse(jabatan))
}

func (jc *JabatanController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	jabatan, err := jc.jabatanService.GetByID(uint(id64))
	if err != nil {
		if err.Error() == "jabatan not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch jabatan"})
		return
	}

	c.JSON(http.StatusOK, dto.MapJabatanToResponse(jabatan))
}

func (jc *JabatanController) Create(c *gin.Context) {
	var req dto.JabatanCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jabatan := &models.Jabatan{
		Nama:      req.Nama,
		Deskripsi: req.Deskripsi,
	}

	if err := jc.jabatanService.Create(jabatan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.MapJabatanToResponse(jabatan))
}

func (jc *JabatanController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req dto.JabatanUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jabatan := &models.Jabatan{
		Nama:      req.Nama,
		Deskripsi: req.Deskripsi,
	}

	if err := jc.jabatanService.Update(uint(id64), jabatan); err != nil {
		if err.Error() == "jabatan not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update jabatan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jabatan updated successfully"})
}

func (jc *JabatanController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := jc.jabatanService.Delete(uint(id64)); err != nil {
		if err.Error() == "jabatan not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete jabatan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jabatan deleted successfully"})
}
