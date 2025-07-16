package seeders

import (
	"absensibe/models"

	"gorm.io/gorm"
)

type SafeAreaSeeder struct {
	BaseSeeder
}

func (s *SafeAreaSeeder) Run(db *gorm.DB) error {
	if DataExists(db, &models.SafeArea{}, "name = ?", "Area Sekolah Utama") {
		return nil
	}

	var schools []models.School
	db.Find(&schools)

	for _, school := range schools {
		safeArea := models.SafeArea{
			SchoolID:    school.ID,
			Name:        "Area Sekolah Utama",
			Latitude:    -6.200000,
			Longitude:   106.816666,
			Radius:      100.0,
			Description: stringPtr("Area utama sekolah untuk absensi"),
			IsActive:    true,
		}
		db.Create(&safeArea)
	}

	return nil
}