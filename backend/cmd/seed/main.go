package main

import (
	"log"
	"time"

	"backend/internal/models"
	"backend/pkg/database"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Could not hash password:", err)
	}
	return string(hashedPassword)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file loaded, continuing with environment variables")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// Drop tables to ensure fresh seed (be careful in production)
	db.Migrator().DropTable(&models.ProfileMatchResult{}, &models.NilaiTenagaKerja{}, &models.TenagaKerja{}, &models.TargetProfile{}, &models.Kriteria{}, &models.Aspek{}, &models.Jabatan{}, &models.User{})

	// AutoMigrate again
	if err := db.AutoMigrate(&models.User{}, &models.Jabatan{}, &models.Aspek{}, &models.Kriteria{}, &models.TargetProfile{}, &models.TenagaKerja{}, &models.NilaiTenagaKerja{}, &models.ProfileMatchResult{}); err != nil {
		log.Fatal("Could not migrate database:", err)
	}

	// Create admin user
	adminUser := models.User{
		Email:    "admin@kpsggroup.com",
		Password: hashPassword("admin123"),
		Nama:     "Administrator",
		Role:     "admin",
		IsActive: true,
	}
	if err := db.Create(&adminUser).Error; err != nil {
		log.Fatal("Could not create admin user:", err)
	}
	log.Println("✅ Admin user created: admin@kpsggroup.com / admin123")

	// Create regular user
	regularUser := models.User{
		Email:    "user@kpsggroup.com",
		Password: hashPassword("user123"),
		Nama:     "User Biasa",
		Role:     "user",
		IsActive: true,
	}
	if err := db.Create(&regularUser).Error; err != nil {
		log.Fatal("Could not create regular user:", err)
	}
	log.Println("✅ Regular user created: user@kpsggroup.com / user123")

	// Create Jabatan
	jabatanData := []models.Jabatan{
		{Nama: "Operator Produksi", Deskripsi: "Mengoperasikan mesin produksi gula"},
		{Nama: "Supervisor Quality Control", Deskripsi: "Mengawasi kualitas produk gula"},
		{Nama: "Teknisi Maintenance", Deskripsi: "Perawatan dan perbaikan mesin"},
	}
	for i := range jabatanData {
		if err := db.Create(&jabatanData[i]).Error; err != nil {
			log.Fatal("Could not create jabatan:", err)
		}
	}
	log.Printf("✅ %d jabatan created\n", len(jabatanData))

	// Create Aspek
	aspekData := []models.Aspek{
		{Nama: "Kompetensi Teknis", Persentase: 40.0},
		{Nama: "Sikap Kerja", Persentase: 35.0},
		{Nama: "Pengalaman", Persentase: 25.0},
	}
	for i := range aspekData {
		if err := db.Create(&aspekData[i]).Error; err != nil {
			log.Fatal("Could not create aspek:", err)
		}
	}
	log.Printf("✅ %d aspek created\n", len(aspekData))

	// Create Kriteria
	var kriteriaData []models.Kriteria
	kriteriaData = append(kriteriaData, models.Kriteria{AspekID: aspekData[0].ID, Kode: "K1", Nama: "Pengetahuan Mesin", IsCore: true, Bobot: 1.0})
	kriteriaData = append(kriteriaData, models.Kriteria{AspekID: aspekData[0].ID, Kode: "K2", Nama: "Kemampuan Teknis", IsCore: true, Bobot: 1.0})
	kriteriaData = append(kriteriaData, models.Kriteria{AspekID: aspekData[0].ID, Kode: "K3", Nama: "Pemahaman SOP", IsCore: false, Bobot: 1.0})
	kriteriaData = append(kriteriaData, models.Kriteria{AspekID: aspekData[1].ID, Kode: "S1", Nama: "Disiplin", IsCore: true, Bobot: 1.0})
	kriteriaData = append(kriteriaData, models.Kriteria{AspekID: aspekData[1].ID, Kode: "S2", Nama: "Kerjasama Tim", IsCore: true, Bobot: 1.0})
	kriteriaData = append(kriteriaData, models.Kriteria{AspekID: aspekData[1].ID, Kode: "S3", Nama: "Inisiatif", IsCore: false, Bobot: 1.0})
	kriteriaData = append(kriteriaData, models.Kriteria{AspekID: aspekData[2].ID, Kode: "P1", Nama: "Lama Kerja", IsCore: true, Bobot: 1.0})
	kriteriaData = append(kriteriaData, models.Kriteria{AspekID: aspekData[2].ID, Kode: "P2", Nama: "Pengalaman Serupa", IsCore: false, Bobot: 1.0})

	for i := range kriteriaData {
		if err := db.Create(&kriteriaData[i]).Error; err != nil {
			log.Fatal("Could not create kriteria:", err)
		}
	}
	log.Printf("✅ %d kriteria created\n", len(kriteriaData))

	// Create Target Profiles for first jabatan
	var targetProfiles []models.TargetProfile
	if len(jabatanData) > 0 {
		jabID := jabatanData[0].ID
		// use first 6 criteria as sample
		for i := 0; i < len(kriteriaData) && i < 8; i++ {
			tp := models.TargetProfile{JabatanID: jabID, KriteriaID: kriteriaData[i].ID, TargetNilai: 3.0 + float64(i%3)}
			targetProfiles = append(targetProfiles, tp)
			if err := db.Create(&tp).Error; err != nil {
				log.Fatal("Could not create target profile:", err)
			}
		}
	}
	log.Printf("✅ %d target profiles created\n", len(targetProfiles))

	// Create Tenaga Kerja
	tenagaKerjaData := []models.TenagaKerja{
		{NIK: "TK001", Nama: "Budi Santoso", TglLahir: time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC), Alamat: "Jl. Mawar No. 10", Telepon: "081234567890"},
		{NIK: "TK002", Nama: "Siti Nurhaliza", TglLahir: time.Date(1992, 8, 20, 0, 0, 0, 0, time.UTC), Alamat: "Jl. Melati No. 5", Telepon: "081234567891"},
		{NIK: "TK003", Nama: "Ahmad Yani", TglLahir: time.Date(1988, 3, 10, 0, 0, 0, 0, time.UTC), Alamat: "Jl. Anggrek No. 15", Telepon: "081234567892"},
	}
	for i := range tenagaKerjaData {
		if err := db.Create(&tenagaKerjaData[i]).Error; err != nil {
			log.Fatal("Could not create tenaga kerja:", err)
		}
	}
	log.Printf("✅ %d tenaga kerja created\n", len(tenagaKerjaData))

	// Create Nilai Tenaga Kerja for combinations
	for _, tk := range tenagaKerjaData {
		for _, k := range kriteriaData {
			nilai := models.NilaiTenagaKerja{TenagaKerjaID: tk.ID, KriteriaID: k.ID, Nilai: 2 + float64((time.Now().UnixNano() % 4))}
			if err := db.Create(&nilai).Error; err != nil {
				log.Fatal("Could not create nilai tenaga kerja:", err)
			}
		}
	}
	log.Println("✅ Nilai tenaga kerja created untuk semua kombinasi")

	log.Println("\n🎉 Seeding completed successfully!")
	log.Println("\n📝 Login credentials:")
	log.Println("   Admin: admin@kpsggroup.com / admin123")
	log.Println("   User:  user@kpsggroup.com / user123")
}
