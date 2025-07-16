// ===== seeders/jurusan_seeder.go =====
package seeders

import (
	"absensibe/models"
	"gorm.io/gorm"
)

type JurusanSeeder struct {
	BaseSeeder
}

func (s *JurusanSeeder) Run(db *gorm.DB) error {
	// Skip if data already exists
	if DataExists(db, &models.Jurusan{}, "code = ?", "RPL") {
		return nil
	}

	// Get school IDs
	var schools []models.School
	if err := db.Find(&schools).Error; err != nil {
		return err
	}

	if len(schools) == 0 {
		return nil // No schools to reference
	}

	jurusanData := []struct {
		Code        string
		Name        string
		Description string
	}{
		{
			Code:        "RPL",
			Name:        "Rekayasa Perangkat Lunak",
			Description: "Program keahlian yang mempelajari pengembangan aplikasi dan sistem perangkat lunak",
		},
		{
			Code:        "TKJ",
			Name:        "Teknik Komputer dan Jaringan",
			Description: "Program keahlian yang mempelajari instalasi, konfigurasi, dan maintenance jaringan komputer",
		},
		{
			Code:        "MM",
			Name:        "Multimedia",
			Description: "Program keahlian yang mempelajari desain grafis, video editing, dan produksi multimedia",
		},
		{
			Code:        "TKRO",
			Name:        "Teknik Kendaraan Ringan Otomotif",
			Description: "Program keahlian yang mempelajari perbaikan dan maintenance kendaraan bermotor",
		},
		{
			Code:        "TEI",
			Name:        "Teknik Elektronika Industri",
			Description: "Program keahlian yang mempelajari sistem elektronika untuk industri",
		},
	}

	// Create jurusan for each school
	for _, school := range schools {
		for _, data := range jurusanData {
			jurusan := models.Jurusan{
				Code:        data.Code,
				Name:        data.Name,
				SchoolID:    school.ID,
				Description: &data.Description,
				IsActive:    true,
			}

			if err := db.Create(&jurusan).Error; err != nil {
				return err
			}
		}
	}

	return nil
}