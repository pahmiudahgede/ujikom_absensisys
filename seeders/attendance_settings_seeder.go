package seeders

import (
	"absensibe/models"
	"time"

	"gorm.io/gorm"
)

type AttendanceSettingsSeeder struct{}

func (s *AttendanceSettingsSeeder) GetName() string {
	return "Attendance Settings"
}

func (s *AttendanceSettingsSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.AttendanceSettings{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var school models.School
	if err := db.First(&school).Error; err != nil {
		return err
	}

	checkinStart, _ := time.Parse("15:04:05", "06:00:00")
	checkinEnd, _ := time.Parse("15:04:05", "07:30:00")
	checkoutStart, _ := time.Parse("15:04:05", "15:00:00")
	checkoutEnd, _ := time.Parse("15:04:05", "17:00:00")

	settings := models.AttendanceSettings{
		SchoolID:        school.ID,
		CheckinStart:    checkinStart,
		CheckinEnd:      checkinEnd,
		CheckoutStart:   checkoutStart,
		CheckoutEnd:     checkoutEnd,
		LateTolerance:   15,
		RequirePhoto:    true,
		RequireLocation: true,
		MaxDistance:     100,
	}

	return db.Create(&settings).Error
}
