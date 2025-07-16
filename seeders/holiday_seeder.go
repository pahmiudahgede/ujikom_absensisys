package seeders

import (
	"absensibe/models"

	"gorm.io/gorm"
)

type HolidaySeeder struct {
	BaseSeeder
}

func (s *HolidaySeeder) Run(db *gorm.DB) error {
	if DataExists(db, &models.Holiday{}, "name = ?", "Hari Kemerdekaan RI") {
		return nil
	}

	var schools []models.School
	db.Find(&schools)

	holidays := []struct {
		Name        string
		Date        string
		Type        string
		Description string
	}{
		{"Hari Kemerdekaan RI", "2024-08-17", "nasional", "Hari Kemerdekaan Republik Indonesia"},
		{"Hari Raya Idul Fitri", "2024-04-10", "agama", "Hari Raya Idul Fitri 1445 H"},
		{"Hari Raya Idul Adha", "2024-06-17", "agama", "Hari Raya Idul Adha 1445 H"},
		{"Hari Natal", "2024-12-25", "agama", "Hari Raya Natal"},
		{"Tahun Baru", "2025-01-01", "nasional", "Tahun Baru Masehi"},
	}

	for _, school := range schools {
		for _, h := range holidays {
			holiday := models.Holiday{
				SchoolID:    school.ID,
				Name:        h.Name,
				Date:        mustParseDate(h.Date),
				Type:        h.Type,
				Description: stringPtr(h.Description),
			}
			db.Create(&holiday)
		}
	}

	return nil
}
