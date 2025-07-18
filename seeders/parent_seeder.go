package seeders

import (
	"absensibe/models"
	"fmt"

	"gorm.io/gorm"
)

type ParentSeeder struct{}

func (s *ParentSeeder) GetName() string {
	return "Parents"
}

func (s *ParentSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Parent{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var students []models.Student
	if err := db.Limit(5).Find(&students).Error; err != nil {
		return err
	}

	parents := []models.Parent{}
	for i, student := range students {
		// Father
		father := models.Parent{
			StudentID: student.ID,
			Fullname:  fmt.Sprintf("Bapak Siswa %d", i+1),
			Phone:     fmt.Sprintf("08111111%03d", i+1),
			Job:       "Pegawai Swasta",
			Relation:  "ayah",
		}

		// Mother
		mother := models.Parent{
			StudentID: student.ID,
			Fullname:  fmt.Sprintf("Ibu Siswa %d", i+1),
			Phone:     fmt.Sprintf("08222222%03d", i+1),
			Job:       "Ibu Rumah Tangga",
			Relation:  "ibu",
		}

		parents = append(parents, father, mother)
	}

	return db.Create(&parents).Error
}
