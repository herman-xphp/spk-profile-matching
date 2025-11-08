package main

import (
	"backend/internal/controllers"
	"backend/internal/middleware"
	"backend/internal/repositories"
	"backend/internal/services"
	"backend/pkg/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Gin mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	// Initialize database connection
	_, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// Initialize router
	router := gin.Default()

	// Add CORS middleware
	router.Use(middleware.CORS())

	// Initialize repositories
	userRepo := repositories.NewUserRepository(database.DB)
	jabatanRepo := repositories.NewJabatanRepository(database.DB)
	aspekRepo := repositories.NewAspekRepository(database.DB)
	kriteriaRepo := repositories.NewKriteriaRepository(database.DB)
	targetProfileRepo := repositories.NewTargetProfileRepository(database.DB)
	tenagaKerjaRepo := repositories.NewTenagaKerjaRepository(database.DB)
	nilaiTenagaKerjaRepo := repositories.NewNilaiTenagaKerjaRepository(database.DB)
	profileMatchResultRepo := repositories.NewProfileMatchResultRepository(database.DB)

	// Initialize services
	authSvc := services.NewAuthService(userRepo)
	userSvc := services.NewUserService(userRepo)
	jabatanSvc := services.NewJabatanService(jabatanRepo)
	aspekSvc := services.NewAspekService(aspekRepo)
	kriteriaSvc := services.NewKriteriaService(kriteriaRepo, aspekRepo)
	targetProfileSvc := services.NewTargetProfileService(targetProfileRepo, jabatanRepo, kriteriaRepo)
	tenagaKerjaSvc := services.NewTenagaKerjaService(tenagaKerjaRepo)
	nilaiTenagaKerjaSvc := services.NewNilaiTenagaKerjaService(nilaiTenagaKerjaRepo, tenagaKerjaRepo, kriteriaRepo)
	profileMatchingSvc := services.NewProfileMatchingService(
		targetProfileRepo,
		kriteriaRepo,
		nilaiTenagaKerjaRepo,
		tenagaKerjaRepo,
		profileMatchResultRepo,
		jabatanRepo,
	)

	// Initialize controllers
	authCtrl := controllers.NewAuthController(authSvc)
	userCtrl := controllers.NewUserController(userSvc)
	jabatanCtrl := controllers.NewJabatanController(jabatanSvc)
	aspekCtrl := controllers.NewAspekController(aspekSvc)
	kriteriaCtrl := controllers.NewKriteriaController(kriteriaSvc)
	targetProfileCtrl := controllers.NewTargetProfileController(targetProfileSvc)
	tenagaKerjaCtrl := controllers.NewTenagaKerjaController(tenagaKerjaSvc)
	nilaiTenagaKerjaCtrl := controllers.NewNilaiTenagaKerjaController(nilaiTenagaKerjaSvc)
	profileMatchingCtrl := controllers.NewProfileMatchingController(profileMatchingSvc)

	// Public routes
	router.POST("/api/auth/login", authCtrl.Login)
	router.POST("/api/auth/register", userCtrl.Register)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Users
		protected.GET("/users", userCtrl.GetAll)
		protected.GET("/users/:id", userCtrl.GetByID)
		protected.PUT("/users/:id", userCtrl.Update)
		protected.DELETE("/users/:id", userCtrl.Delete)

		// Jabatan
		protected.GET("/jabatan", jabatanCtrl.GetAll)
		protected.POST("/jabatan", jabatanCtrl.Create)
		protected.GET("/jabatan/:id", jabatanCtrl.GetByID)
		protected.PUT("/jabatan/:id", jabatanCtrl.Update)
		protected.DELETE("/jabatan/:id", jabatanCtrl.Delete)

		// Aspek
		protected.GET("/aspek", aspekCtrl.GetAll)
		protected.POST("/aspek", aspekCtrl.Create)
		protected.GET("/aspek/:id", aspekCtrl.GetByID)
		protected.PUT("/aspek/:id", aspekCtrl.Update)
		protected.DELETE("/aspek/:id", aspekCtrl.Delete)

		// Kriteria
		protected.GET("/kriteria", kriteriaCtrl.GetAll)
		protected.POST("/kriteria", kriteriaCtrl.Create)
		protected.GET("/kriteria/:id", kriteriaCtrl.GetByID)
		protected.PUT("/kriteria/:id", kriteriaCtrl.Update)
		protected.DELETE("/kriteria/:id", kriteriaCtrl.Delete)

		// Target Profile
		protected.GET("/target-profiles", targetProfileCtrl.GetAll)
		protected.POST("/target-profiles", targetProfileCtrl.Create)
		protected.GET("/target-profiles/:id", targetProfileCtrl.GetByID)
		protected.PUT("/target-profiles/:id", targetProfileCtrl.Update)
		protected.DELETE("/target-profiles/:id", targetProfileCtrl.Delete)

		// Tenaga Kerja
		protected.GET("/tenaga-kerja", tenagaKerjaCtrl.GetAll)
		protected.POST("/tenaga-kerja", tenagaKerjaCtrl.Create)
		protected.GET("/tenaga-kerja/:id", tenagaKerjaCtrl.GetByID)
		protected.PUT("/tenaga-kerja/:id", tenagaKerjaCtrl.Update)
		protected.DELETE("/tenaga-kerja/:id", tenagaKerjaCtrl.Delete)

		// Nilai Tenaga Kerja
		protected.GET("/nilai-tenaga-kerja", nilaiTenagaKerjaCtrl.GetAll)
		protected.POST("/nilai-tenaga-kerja", nilaiTenagaKerjaCtrl.Create)
		protected.GET("/nilai-tenaga-kerja/:id", nilaiTenagaKerjaCtrl.GetByID)
		protected.PUT("/nilai-tenaga-kerja/:id", nilaiTenagaKerjaCtrl.Update)
		protected.DELETE("/nilai-tenaga-kerja/:id", nilaiTenagaKerjaCtrl.Delete)

		// Profile Matching Calculation
		protected.POST("/profile-matching/calculate", profileMatchingCtrl.Calculate)
		protected.GET("/profile-matching/results", profileMatchingCtrl.GetAllResults)
		protected.GET("/profile-matching/results/:id", profileMatchingCtrl.GetResultByID)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
