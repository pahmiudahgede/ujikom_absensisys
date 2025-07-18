package seeders

import (
	"absensibe/models"
	"time"

	"gorm.io/gorm"
)

type AcademicCalendarSeeder struct{}

func (s *AcademicCalendarSeeder) GetName() string {
	return "Academic Calendars"
}

func (s *AcademicCalendarSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.AcademicCalendar{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var school models.School
	if err := db.First(&school).Error; err != nil {
		return err
	}

	calendars := []models.AcademicCalendar{
		{
			SchoolID:     school.ID,
			AcademicYear: "2024/2025",
			Semester:     "ganjil",
			StartDate:    time.Date(2024, 7, 15, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2024, 12, 20, 0, 0, 0, 0, time.UTC),
			Description:  stringPtr("Semester Ganjil Tahun Ajaran 2024/2025"),
			IsActive:     true,
		},
		{
			SchoolID:     school.ID,
			AcademicYear: "2024/2025",
			Semester:     "genap",
			StartDate:    time.Date(2025, 1, 7, 0, 0, 0, 0, time.UTC),
			EndDate:      time.Date(2025, 6, 15, 0, 0, 0, 0, time.UTC),
			Description:  stringPtr("Semester Genap Tahun Ajaran 2024/2025"),
			IsActive:     true,
		},
	}

	return db.Create(&calendars).Error
}
