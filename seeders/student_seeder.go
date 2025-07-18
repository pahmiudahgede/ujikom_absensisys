package seeders

import (
	"absensibe/models"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type StudentSeeder struct{}

func (s *StudentSeeder) GetName() string {
	return "Students"
}

func (s *StudentSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Student{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var classes []models.Class
	if err := db.Find(&classes).Error; err != nil {
		return err
	}

	if len(classes) == 0 {
		return nil
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("student123"), bcrypt.DefaultCost)

	students := []models.Student{}

	// Create 10 students for first class
	for i := 1; i <= 10; i++ {
		student := models.Student{
			NISN:         fmt.Sprintf("300120240%02d", i),
			NIS:          fmt.Sprintf("240%02d", i),
			ClassID:      classes[0].ID,
			Password:     string(hashedPassword),
			Fullname:     fmt.Sprintf("Siswa %d", i),
			Gender:       "L",
			BirthDate:    time.Date(2008, time.Month(i%12+1), i+10, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Jakarta",
			Phone:        fmt.Sprintf("08123456789%d", i),
			Status:       "aktif",
			EntryYear:    2024,
		}

		if i%2 == 0 {
			student.Gender = "P"
			student.Fullname = fmt.Sprintf("Siswi %d", i)
		}

		students = append(students, student)
	}

	return db.Create(&students).Error
}
