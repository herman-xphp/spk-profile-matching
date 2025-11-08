package services

import (
	"errors"

	"backend/internal/models"
	"backend/internal/repositories"

	"gorm.io/gorm"
)

type ProfileMatchingService struct {
	targetProfileRepo      *repositories.TargetProfileRepository
	kriteriaRepo           *repositories.KriteriaRepository
	nilaiTenagaKerjaRepo   *repositories.NilaiTenagaKerjaRepository
	tenagaKerjaRepo        *repositories.TenagaKerjaRepository
	profileMatchResultRepo *repositories.ProfileMatchResultRepository
	jabatanRepo            *repositories.JabatanRepository
}

func NewProfileMatchingService(
	targetProfileRepo *repositories.TargetProfileRepository,
	kriteriaRepo *repositories.KriteriaRepository,
	nilaiTenagaKerjaRepo *repositories.NilaiTenagaKerjaRepository,
	tenagaKerjaRepo *repositories.TenagaKerjaRepository,
	profileMatchResultRepo *repositories.ProfileMatchResultRepository,
	jabatanRepo *repositories.JabatanRepository,
) *ProfileMatchingService {
	return &ProfileMatchingService{
		targetProfileRepo:      targetProfileRepo,
		kriteriaRepo:           kriteriaRepo,
		nilaiTenagaKerjaRepo:   nilaiTenagaKerjaRepo,
		tenagaKerjaRepo:        tenagaKerjaRepo,
		profileMatchResultRepo: profileMatchResultRepo,
		jabatanRepo:            jabatanRepo,
	}
}

type CalculationRequest struct {
	JabatanID       uint
	TenagaKerjaIDs  []uint
}

func (s *ProfileMatchingService) Calculate(req CalculationRequest) ([]models.ProfileMatchResult, error) {
	// Validate jabatan exists
	_, err := s.jabatanRepo.GetByID(req.JabatanID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("jabatan not found")
		}
		return nil, err
	}

	// Get target profiles for the position
	targetProfiles, err := s.targetProfileRepo.GetByJabatanID(req.JabatanID)
	if err != nil {
		return nil, errors.New("could not fetch target profiles")
	}

	if len(targetProfiles) == 0 {
		return nil, errors.New("no target profiles found for this jabatan")
	}

	// Get all kriteria
	kriterias, err := s.kriteriaRepo.GetAll()
	if err != nil {
		return nil, errors.New("could not fetch kriteria")
	}

	// Create map of kriteria for easy lookup
	kriteriaMap := make(map[uint]models.Kriteria)
	for _, k := range kriterias {
		kriteriaMap[k.ID] = k
	}

	// Get tenaga kerja IDs to process
	tenagaKerjaIDs := req.TenagaKerjaIDs
	if len(tenagaKerjaIDs) == 0 {
		// If no tenaga_kerja_ids provided, get all tenaga kerja that have nilai
		allNilai, err := s.nilaiTenagaKerjaRepo.GetAll()
		if err != nil {
			return nil, errors.New("could not fetch nilai tenaga kerja")
		}
		
		// Get unique tenaga kerja IDs
		idMap := make(map[uint]bool)
		for _, n := range allNilai {
			idMap[n.TenagaKerjaID] = true
		}
		
		for id := range idMap {
			tenagaKerjaIDs = append(tenagaKerjaIDs, id)
		}
	}

	var results []models.ProfileMatchResult

	for _, tenagaKerjaID := range tenagaKerjaIDs {
		// Validate tenaga kerja exists
		_, err := s.tenagaKerjaRepo.GetByID(tenagaKerjaID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue // Skip invalid tenaga kerja
			}
			return nil, err
		}

		// Get nilai for this tenaga kerja
		nilaiList, err := s.nilaiTenagaKerjaRepo.GetByTenagaKerjaID(tenagaKerjaID)
		if err != nil {
			continue // Skip if no nilai found
		}

		// Create map of nilai for easy lookup
		nilaiMap := make(map[uint]float64)
		for _, n := range nilaiList {
			nilaiMap[n.KriteriaID] = n.Nilai
		}

		// Calculate scores
		var totalCoreGap float64
		var totalSecondaryCap float64
		var countCore int
		var countSecondary int

		for _, target := range targetProfiles {
			kriteria, exists := kriteriaMap[target.KriteriaID]
			if !exists {
				continue
			}

			nilai, exists := nilaiMap[target.KriteriaID]
			if !exists {
				continue // Skip if no nilai for this kriteria
			}

			// Calculate GAP
			gap := nilai - target.TargetNilai

			// Convert GAP to weight based on profile matching rules
			weight := calculateWeight(gap)

			if kriteria.IsCore {
				totalCoreGap += weight
				countCore++
			} else {
				totalSecondaryCap += weight
				countSecondary++
			}
		}

		// Avoid divide by zero
		var coreFactor float64
		var secondaryFactor float64
		if countCore > 0 {
			coreFactor = totalCoreGap / float64(countCore)
		}
		if countSecondary > 0 {
			secondaryFactor = totalSecondaryCap / float64(countSecondary)
		}

		// Final calculation (60% core factor + 40% secondary factor)
		totalScore := (0.6 * coreFactor) + (0.4 * secondaryFactor)

		// Create result
		result := models.ProfileMatchResult{
			TenagaKerjaID:   tenagaKerjaID,
			JabatanID:       req.JabatanID,
			TotalScore:      totalScore,
			CoreFactor:      coreFactor,
			SecondaryFactor: secondaryFactor,
		}

		results = append(results, result)
	}

	// Save all results to database
	if len(results) > 0 {
		if err := s.profileMatchResultRepo.CreateBatch(results); err != nil {
			return nil, errors.New("could not save calculation results")
		}
	}

	return results, nil
}

func (s *ProfileMatchingService) GetAllResults() ([]models.ProfileMatchResult, error) {
	return s.profileMatchResultRepo.GetAllWithRelations()
}

func (s *ProfileMatchingService) GetResultsByJabatanID(jabatanID uint) ([]models.ProfileMatchResult, error) {
	return s.profileMatchResultRepo.GetByJabatanID(jabatanID)
}

func (s *ProfileMatchingService) GetResultByID(id uint) (*models.ProfileMatchResult, error) {
	result, err := s.profileMatchResultRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("result not found")
		}
		return nil, err
	}
	return result, nil
}

// calculateWeight converts GAP to weight according to profile matching rules
func calculateWeight(gap float64) float64 {
	switch gap {
	case 0:
		return 5 // No difference
	case 1:
		return 4.5 // Competency excess 1 level
	case -1:
		return 4 // Competency lack 1 level
	case 2:
		return 3.5 // Competency excess 2 levels
	case -2:
		return 3 // Competency lack 2 levels
	case 3:
		return 2.5 // Competency excess 3 levels
	case -3:
		return 2 // Competency lack 3 levels
	case 4:
		return 1.5 // Competency excess 4 levels
	case -4:
		return 1 // Competency lack 4 levels
	default:
		return 0 // Gap too large
	}
}

