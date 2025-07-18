package seeders

import (
	"absensibe/models"

	"gorm.io/gorm"
)

type MajorSeeder struct{}

func (s *MajorSeeder) GetName() string {
	return "Majors"
}

func (s *MajorSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Major{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	// Get first school
	var school models.School
	if err := db.First(&school).Error; err != nil {
		return err
	}

	majors := []models.Major{
		{
			Code:        "RPL",
			Name:        "Rekayasa Perangkat Lunak",
			SchoolID:    school.ID,
			Description: stringPtr("Jurusan yang mempelajari pengembangan aplikasi dan sistem informasi"),
			IsActive:    true,
		},
		{
			Code:        "TKJ",
			Name:        "Teknik Komputer dan Jaringan",
			SchoolID:    school.ID,
			Description: stringPtr("Jurusan yang mempelajari instalasi, konfigurasi, dan maintenance jaringan komputer"),
			IsActive:    true,
		},
		{
			Code:        "MM",
			Name:        "Multimedia",
			SchoolID:    school.ID,
			Description: stringPtr("Jurusan yang mempelajari desain grafis, video editing, dan produksi multimedia"),
			IsActive:    true,
		},
		{
			Code:        "TBSM",
			Name:        "Teknik dan Bisnis Sepeda Motor",
			SchoolID:    school.ID,
			Description: stringPtr("Jurusan yang mempelajari maintenance dan perbaikan sepeda motor"),
			IsActive:    true,
		},
	}

	return db.Create(&majors).Error
}
