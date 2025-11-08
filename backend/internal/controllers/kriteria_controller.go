package controllers

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

type KriteriaController struct {
	kriteriaService *services.KriteriaService
}

func NewKriteriaController(kriteriaService *services.KriteriaService) *KriteriaController {
	return &KriteriaController{kriteriaService: kriteriaService}
}

func (kc *KriteriaController) GetAll(c *gin.Context) {
	kriteria, err := kc.kriteriaService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch kriteria"})
		return
	}
	c.JSON(http.StatusOK, dto.MapKriteriasToResponse(kriteria))
}

func (kc *KriteriaController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	kriteria, err := kc.kriteriaService.GetByID(uint(id64))
	if err != nil {
		if err.Error() == "kriteria not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch kriteria"})
		return
	}

	c.JSON(http.StatusOK, dto.MapKriteriaToResponse(kriteria))
}

func (kc *KriteriaController) Create(c *gin.Context) {
	var req dto.KriteriaCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kriteria := &models.Kriteria{
		AspekID: uint(req.AspekID),
		Kode:    req.Kode,
		Nama:    req.Nama,
		IsCore:  req.IsCore,
		Bobot:   req.Bobot,
	}

	if err := kc.kriteriaService.Create(kriteria); err != nil {
		if err.Error() == "aspek not found" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.MapKriteriaToResponse(kriteria))
}

func (kc *KriteriaController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req dto.KriteriaUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	kriteria := &models.Kriteria{
		AspekID: req.AspekID,
		Kode:    req.Kode,
		Nama:    req.Nama,
		IsCore:  *req.IsCore,
		Bobot:   req.Bobot,
	}

	if err := kc.kriteriaService.Update(uint(id64), kriteria); err != nil {
		if err.Error() == "kriteria not found" || err.Error() == "aspek not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update kriteria"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kriteria updated successfully"})
}

func (kc *KriteriaController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := kc.kriteriaService.Delete(uint(id64)); err != nil {
		if err.Error() == "kriteria not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete kriteria"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kriteria deleted successfully"})
}
