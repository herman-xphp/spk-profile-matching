package controllers

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

type AspekController struct {
	aspekService *services.AspekService
}

func NewAspekController(aspekService *services.AspekService) *AspekController {
	return &AspekController{aspekService: aspekService}
}

func (ac *AspekController) GetAll(c *gin.Context) {
	aspek, err := ac.aspekService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch aspek"})
		return
	}
	c.JSON(http.StatusOK, dto.MapAspeksToResponse(aspek))
}

func (ac *AspekController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	aspek, err := ac.aspekService.GetByID(uint(id64))
	if err != nil {
		if err.Error() == "aspek not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch aspek"})
		return
	}

	c.JSON(http.StatusOK, dto.MapAspekToResponse(aspek))
}

func (ac *AspekController) Create(c *gin.Context) {
	// var aspek models.Aspek
	var req dto.AspekCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aspek := &models.Aspek{
		Nama:       req.Nama,
		Deskripsi:  req.Deskripsi,
		Persentase: req.Persentase,
	}

	if err := ac.aspekService.Create(aspek); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.MapAspekToResponse(aspek))
}

func (ac *AspekController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req dto.AspekUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aspek := &models.Aspek{
		Nama:       req.Nama,
		Deskripsi:  req.Deskripsi,
		Persentase: req.Persentase,
	}

	if err := ac.aspekService.Update(uint(id64), aspek); err != nil {
		if err.Error() == "aspek not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update aspek"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aspek updated successfully"})
}

func (ac *AspekController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := ac.aspekService.Delete(uint(id64)); err != nil {
		if err.Error() == "aspek not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete aspek"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aspek deleted successfully"})
}
