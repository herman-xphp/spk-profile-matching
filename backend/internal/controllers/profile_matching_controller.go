package controllers

import (
	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

type ProfileMatchingController struct {
	profileMatchingService *services.ProfileMatchingService
}

func NewProfileMatchingController(profileMatchingService *services.ProfileMatchingService) *ProfileMatchingController {
	return &ProfileMatchingController{profileMatchingService: profileMatchingService}
}

func (pmc *ProfileMatchingController) Calculate(c *gin.Context) {
	var req dto.CalculationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results, err := pmc.profileMatchingService.Calculate(services.CalculationRequest{
		JabatanID:      req.JabatanID,
		TenagaKerjaIDs: req.TenagaKerjaIDs,
	})
	if err != nil {
		if err.Error() == "jabatan not found" || err.Error() == "no target profiles found for this jabatan" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to DTO response
	response := dto.MapProfileMatchResultsToResponse(results)
	c.JSON(http.StatusOK, response)
}

func (pmc *ProfileMatchingController) GetAllResults(c *gin.Context) {
	// Check if jabatan_id query param exists
	jabatanIDStr := c.Query("jabatan_id")
	if jabatanIDStr != "" {
		jabatanID, err := strconv.ParseUint(jabatanIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid jabatan_id format"})
			return
		}

		results, err := pmc.profileMatchingService.GetResultsByJabatanID(uint(jabatanID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch results"})
			return
		}

		// Convert to RankingResponse with rank for frontend
		response := dto.MapProfileMatchResultsToRankingResponse(results)
		c.JSON(http.StatusOK, response)
		return
	}

	// Get all results
	results, err := pmc.profileMatchingService.GetAllResults()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch results"})
		return
	}

	// Convert to DTO response
	response := dto.MapProfileMatchResultsToResponse(results)
	c.JSON(http.StatusOK, response)
}

func (pmc *ProfileMatchingController) GetResultByID(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	result, err := pmc.profileMatchingService.GetResultByID(uint(id64))
	if err != nil {
		if err.Error() == "result not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch result"})
		return
	}

	// Convert to DTO response
	response := dto.MapProfileMatchResultToResponse(result)
	c.JSON(http.StatusOK, response)
}

