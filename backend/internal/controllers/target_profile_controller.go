package controllers

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/models"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

type TargetProfileController struct {
	targetProfileService *services.TargetProfileService
}

func NewTargetProfileController(targetProfileService *services.TargetProfileService) *TargetProfileController {
	return &TargetProfileController{targetProfileService: targetProfileService}
}

func (tpc *TargetProfileController) GetAll(c *gin.Context) {
	// Check if jabatan_id query param exists
	jabatanIDStr := c.Query("jabatan_id")
	if jabatanIDStr != "" {
		jabatanID, err := strconv.ParseUint(jabatanIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid jabatan_id format"})
			return
		}

		profiles, err := tpc.targetProfileService.GetByJabatanID(uint(jabatanID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch target profiles"})
			return
		}
		c.JSON(http.StatusOK, dto.MapTargetProfilesToResponse(profiles))
		return
	}

	// Get all profiles
	profiles, err := tpc.targetProfileService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch target profiles"})
		return
	}
	c.JSON(http.StatusOK, dto.MapTargetProfilesToResponse(profiles))
}

func (tpc *TargetProfileController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	profile, err := tpc.targetProfileService.GetByID(uint(id64))
	if err != nil {
		if err.Error() == "target profile not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch target profile"})
		return
	}

	c.JSON(http.StatusOK, dto.MapTargetProfileToResponse(profile))
}

func (tpc *TargetProfileController) Create(c *gin.Context) {
	var req dto.TargetProfileCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile := &models.TargetProfile{
		JabatanID:   req.JabatanID,
		KriteriaID:  req.KriteriaID,
		TargetNilai: req.TargetNilai,
	}

	if err := tpc.targetProfileService.Create(profile); err != nil {
		if err.Error() == "jabatan not found" || err.Error() == "kriteria not found" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create target profile"})
		return
	}

	c.JSON(http.StatusCreated, dto.MapTargetProfileToResponse(profile))
}

func (tpc *TargetProfileController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req dto.TargetProfileUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile := &models.TargetProfile{
		JabatanID:   req.JabatanID,
		KriteriaID:  req.KriteriaID,
		TargetNilai: req.TargetNilai,
	}

	if err := tpc.targetProfileService.Update(uint(id64), profile); err != nil {
		if err.Error() == "target profile not found" || err.Error() == "jabatan not found" || err.Error() == "kriteria not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update target profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Target profile updated successfully"})
}

func (tpc *TargetProfileController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := tpc.targetProfileService.Delete(uint(id64)); err != nil {
		if err.Error() == "target profile not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete target profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Target profile deleted successfully"})
}
