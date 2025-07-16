package seeders

import (
	"absensibe/models"

	"gorm.io/gorm"
)

type AcademicCalendarSeeder struct {
	BaseSeeder
}

func (s *AcademicCalendarSeeder) Run(db *gorm.DB) error {
	if DataExists(db, &models.AcademicCalendar{}, "academic_year = ?", "2024/2025") {
		return nil
	}

	var schools []models.School
	db.Find(&schools)

	for _, school := range schools {

		ganjil := models.AcademicCalendar{
			SchoolID:     school.ID,
			AcademicYear: "2024/2025",
			Semester:     "ganjil",
			StartDate:    mustParseDate("2024-07-15"),
			EndDate:      mustParseDate("2024-12-20"),
			Description:  stringPtr("Semester Ganjil Tahun Ajaran 2024/2025"),
			IsActive:     true,
		}

		genap := models.AcademicCalendar{
			SchoolID:     school.ID,
			AcademicYear: "2024/2025",
			Semester:     "genap",
			StartDate:    mustParseDate("2025-01-06"),
			EndDate:      mustParseDate("2025-06-20"),
			Description:  stringPtr("Semester Genap Tahun Ajaran 2024/2025"),
			IsActive:     true,
		}

		db.Create(&ganjil)
		db.Create(&genap)
	}

	return nil
}
