package seeders

import (
	"absensibe/models"
	"time"

	"gorm.io/gorm"
)

type SchoolSeeder struct{}

func (s *SchoolSeeder) GetName() string {
	return "Schools"
}

func (s *SchoolSeeder) Seed(db *gorm.DB) error {
	// Check if schools already exist
	var count int64
	if err := db.Model(&models.School{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil // Skip if data exists
	}

	standFrom := time.Date(1985, 7, 15, 0, 0, 0, 0, time.UTC)

	schools := []models.School{
		{
			NPSN:          "20401234",
			Name:          "SMK Negeri 1 Jakarta",
			Email:         "smkn1jakarta@education.go.id",
			Phone:         "021-1234567",
			Fax:           stringPtr("021-1234568"),
			Website:       stringPtr("https://smkn1jakarta.sch.id"),
			StandFrom:     &standFrom,
			Akreditasi:    "A",
			KepalaSekolah: stringPtr("Dr. Ahmad Susanto, M.Pd"),
			IsActive:      true,
		},
		{
			NPSN:          "20401235",
			Name:          "SMK Negeri 2 Jakarta",
			Email:         "smkn2jakarta@education.go.id",
			Phone:         "021-2345678",
			Fax:           stringPtr("021-2345679"),
			Website:       stringPtr("https://smkn2jakarta.sch.id"),
			StandFrom:     &standFrom,
			Akreditasi:    "A",
			KepalaSekolah: stringPtr("Dra. Siti Nurhaliza, M.M"),
			IsActive:      true,
		},
	}

	return db.Create(&schools).Error
}
