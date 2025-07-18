package seeders

import (
	"absensibe/models"

	"gorm.io/gorm"
)

type SafeAreaSeeder struct{}

func (s *SafeAreaSeeder) GetName() string {
	return "Safe Areas"
}

func (s *SafeAreaSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.SafeArea{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var school models.School
	if err := db.First(&school).Error; err != nil {
		return err
	}

	safeAreas := []models.SafeArea{
		{
			SchoolID:    school.ID,
			Name:        "Area Sekolah Utama",
			Latitude:    -6.1751,
			Longitude:   106.8650,
			Radius:      100.0,
			Description: stringPtr("Area utama sekolah dengan radius 100 meter"),
			IsActive:    true,
		},
		{
			SchoolID:    school.ID,
			Name:        "Area Parkir",
			Latitude:    -6.1755,
			Longitude:   106.8655,
			Radius:      50.0,
			Description: stringPtr("Area parkir sekolah"),
			IsActive:    true,
		},
	}

	return db.Create(&safeAreas).Error
}
