package seeders

import (
	"absensibe/models"
	"log"

	"gorm.io/gorm"
)

func SeedSchools(db *gorm.DB) error {
	log.Println("🏫 Seeding schools...")

	schools := []models.School{
		{
			ID:        "550e8400-e29b-41d4-a716-446655440001",
			NPSN:      "20123456",
			Name:      "SMK Negeri 1 Bandung",
			Email:     "info@smkn1bandung.sch.id",
			Phone:     "022-7234567",
			Website:   stringPtr("https://smkn1bandung.sch.id"),
			Address:   "Jl. Wastukancana No.3, Bandung Wetan, Kec. Bandung Wetan, Kota Bandung, Jawa Barat 40115",
			Principal: "Dr. Ahmad Susanto, S.Pd., M.Pd.",
			Latitude:  -6.9174,
			Longitude: 107.6191,
			IsActive:  true,
		},
	}

	for _, school := range schools {
		if err := db.FirstOrCreate(&school, models.School{NPSN: school.NPSN}).Error; err != nil {
			return err
		}
	}

	log.Printf("✅ Successfully seeded %d schools", len(schools))
	return nil
}
