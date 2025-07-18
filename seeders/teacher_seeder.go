package seeders

import (
	"absensibe/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type TeacherSeeder struct{}

func (s *TeacherSeeder) GetName() string {
	return "Teachers"
}

func (s *TeacherSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Teacher{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var school models.School
	if err := db.First(&school).Error; err != nil {
		return err
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("teacher123"), bcrypt.DefaultCost)

	teachers := []models.Teacher{
		{
			NIP:          "196501011990031001",
			SchoolID:     school.ID,
			Fullname:     "Budi Santoso, S.Kom",
			Email:        "budi.santoso@smkn1.edu",
			Phone:        "081234567890",
			Password:     string(hashedPassword),
			Gender:       "L",
			BirthDate:    time.Date(1965, 1, 1, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Jakarta",
			EmployeeType: "PNS",
			Status:       "aktif",
		},
		{
			NIP:          "197203151995122001",
			SchoolID:     school.ID,
			Fullname:     "Siti Rahayu, S.Pd",
			Email:        "siti.rahayu@smkn1.edu",
			Phone:        "081234567891",
			Password:     string(hashedPassword),
			Gender:       "P",
			BirthDate:    time.Date(1972, 3, 15, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Bandung",
			EmployeeType: "PNS",
			Status:       "aktif",
		},
		{
			NIP:          "198505102009021002",
			SchoolID:     school.ID,
			Fullname:     "Ahmad Fauzi, S.T",
			Email:        "ahmad.fauzi@smkn1.edu",
			Phone:        "081234567892",
			Password:     string(hashedPassword),
			Gender:       "L",
			BirthDate:    time.Date(1985, 5, 10, 0, 0, 0, 0, time.UTC),
			PlaceOfBirth: "Surabaya",
			EmployeeType: "PPPK",
			Status:       "aktif",
		},
	}

	return db.Create(&teachers).Error
}
