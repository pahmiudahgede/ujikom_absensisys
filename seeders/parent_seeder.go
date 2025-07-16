package seeders

import (
	"absensibe/models"
	"fmt"

	"gorm.io/gorm"
)

type ParentSeeder struct {
	BaseSeeder
}

func (s *ParentSeeder) Run(db *gorm.DB) error {
	if DataExists(db, &models.Parent{}, "sebagai = ?", "ayah") {
		return nil
	}

	var students []models.Student
	if err := db.Limit(100).Find(&students).Error; err != nil {
		return err
	}

	jobs := []string{"PNS", "Swasta", "Wiraswasta", "Petani", "Guru", "Dokter", "Pedagang", "Buruh"}

	for i, student := range students {

		father := models.Parent{
			StudentID: student.ID,
			Fullname:  fmt.Sprintf("Ayah %s", student.Fullname),
			Phone:     fmt.Sprintf("0813%07d", i+1000000),
			Job:       jobs[i%len(jobs)],
			Sebagai:   "ayah",
		}

		mother := models.Parent{
			StudentID: student.ID,
			Fullname:  fmt.Sprintf("Ibu %s", student.Fullname),
			Phone:     fmt.Sprintf("0814%07d", i+1000000),
			Job:       jobs[(i+1)%len(jobs)],
			Sebagai:   "ibu",
		}

		if err := db.Create(&father).Error; err != nil {
			return err
		}
		if err := db.Create(&mother).Error; err != nil {
			return err
		}
	}

	return nil
}
