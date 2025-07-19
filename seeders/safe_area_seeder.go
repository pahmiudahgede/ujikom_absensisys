package seeders

import (
	"absensibe/models"
	"log"

	"gorm.io/gorm"
)

func SeedSafeAreas(db *gorm.DB) error {
	log.Println("🛡️ Seeding safe areas...")

	safeAreas := []models.SafeArea{
		{
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			Name:        "Area Sekolah Utama",
			Latitude:    -6.9174,
			Longitude:   107.6191,
			Radius:      150.0,
			Description: stringPtr("Area utama sekolah untuk absensi siswa"),
			IsActive:    true,
		},
	}

	for _, safeArea := range safeAreas {
		if err := db.FirstOrCreate(&safeArea, models.SafeArea{
			SchoolID: safeArea.SchoolID,
			Name:     safeArea.Name,
		}).Error; err != nil {
			return err
		}
	}

	log.Printf("✅ Successfully seeded %d safe areas", len(safeAreas))
	return nil
}
