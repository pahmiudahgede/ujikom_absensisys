// ===== seeders/school_seeder.go =====
package seeders

import (
	"absensibe/models"
	"time"

	"gorm.io/gorm"
)

type SchoolSeeder struct {
	BaseSeeder
}

func (s *SchoolSeeder) Run(db *gorm.DB) error {
	// Skip if data already exists
	if DataExists(db, &models.School{}, "npsn = ?", "12345678") {
		return nil
	}

	standFrom, _ := time.Parse("2006-01-02", "1995-08-17")

	schools := []models.School{
		{
			NPSN:          "12345678",
			Name:          "SMK Negeri 1 Jakarta",
			Email:         "smkn1jakarta@education.go.id",
			Phone:         "021-12345678",
			Fax:           stringPtr("021-12345679"),
			Website:       stringPtr("https://smkn1jakarta.sch.id"),
			StandFrom:     &standFrom,
			Akreditasi:    "A",
			KepalaSekolah: stringPtr("Dr. Ahmad Santoso, S.Pd., M.M."),
			IsActive:      true,
		},
		{
			NPSN:          "87654321",
			Name:          "SMK Swasta Teknologi Nusantara",
			Email:         "admin@smkteknologi.sch.id",
			Phone:         "021-87654321",
			Fax:           stringPtr("021-87654322"),
			Website:       stringPtr("https://smkteknologi.sch.id"),
			StandFrom:     &standFrom,
			Akreditasi:    "B",
			KepalaSekolah: stringPtr("Dra. Siti Rahayu, M.Pd."),
			IsActive:      true,
		},
	}

	for _, school := range schools {
		if err := db.Create(&school).Error; err != nil {
			return err
		}
	}

	return nil
}

// Helper function
func stringPtr(s string) *string {
	return &s
}
