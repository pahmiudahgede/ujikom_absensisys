package seeders

import (
	"absensibe/models"
	"fmt"

	"gorm.io/gorm"
)

type AttendanceSettingsSeeder struct {
	BaseSeeder
}

func (s *AttendanceSettingsSeeder) Run(db *gorm.DB) error {
	if DataExists(db, &models.AttendanceSettings{}, "school_id != ?", "") {
		return nil
	}

	var schools []models.School
	db.Find(&schools)

	for _, school := range schools {
		settings := models.AttendanceSettings{
			SchoolID:        school.ID,
			CheckinStart:    models.NewTimeOnly(6, 0, 0),
			CheckinEnd:      models.NewTimeOnly(7, 30, 0),
			CheckoutStart:   models.NewTimeOnly(15, 0, 0),
			CheckoutEnd:     models.NewTimeOnly(17, 0, 0),
			LateTolerance:   15,
			RequirePhoto:    true,
			RequireLocation: true,
			MaxDistance:     100,
		}

		if err := db.Create(&settings).Error; err != nil {
			return fmt.Errorf("failed to create attendance settings for school %s: %v", school.Name, err)
		}
	}

	return nil
}
