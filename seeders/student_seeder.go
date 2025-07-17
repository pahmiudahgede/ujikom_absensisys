// ===== seeders/student_seeder.go =====
package seeders

import (
	"absensibe/models"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type StudentSeeder struct {
	BaseSeeder
}

func (s *StudentSeeder) Run(db *gorm.DB) error {
	// Skip if data already exists
	if DataExists(db, &models.Student{}, "nisn = ?", "2024000001") {
		return nil
	}

	// Get classes
	var classes []models.Class
	if err := db.Find(&classes).Error; err != nil {
		return err
	}

	if len(classes) == 0 {
		return nil
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("student123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Sample student names
	maleNames := []string{
		"Ahmad Rizki", "Budi Santoso", "Candra Wijaya", "Deni Prasetyo", "Eko Saputra",
		"Fajar Nugroho", "Gilang Ramadhan", "Hendra Kusuma", "Indra Permana", "Joko Susilo",
		"Kevin Ananda", "Lukman Hakim", "Mario Teguh", "Nanda Pratama", "Oscar Lawalata",
	}

	femaleNames := []string{
		"Aisyah Putri", "Bella Safira", "Citra Dewi", "Diva Maharani", "Elsa Kartika",
		"Fiona Anggraini", "Gita Savitri", "Hana Pertiwi", "Indah Sari", "Jessica Mila",
		"Kirana Larasati", "Luna Maya", "Maudy Ayunda", "Nabila Razali", "Olivia Zalianty",
	}

	cities := []string{
		"Jakarta", "Bandung", "Surabaya", "Medan", "Makassar",
		"Palembang", "Semarang", "Denpasar", "Yogyakarta", "Malang",
	}

	studentCounter := 1

	for _, class := range classes {
		// Create 30 students per class
		for i := 1; i <= 30; i++ {
			var name string
			var gender string

			// Alternate between male and female
			if i%2 == 1 {
				name = maleNames[(i-1)/2%len(maleNames)]
				gender = "L"
			} else {
				name = femaleNames[(i-1)/2%len(femaleNames)]
				gender = "P"
			}

			// Generate birth date (age 15-18)
			birthYear := 2024 - 15 - (i % 4) // Age 15-18
			birthMonth := (i % 12) + 1
			birthDay := (i % 28) + 1
			birthDate := time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC)

			// Entry year based on grade
			entryYear := 2024
			if class.Grade == "XI" {
				entryYear = 2023
			} else if class.Grade == "XII" {
				entryYear = 2022
			}

			student := models.Student{
				NISN:         fmt.Sprintf("2024%06d", studentCounter),
				NIS:          fmt.Sprintf("24%04d", studentCounter),
				ClassesID:    class.ID,
				Password:     string(hashedPassword),
				Fullname:     name,
				Gender:       gender,
				BirthDate:    birthDate,
				PlaceOfBirth: cities[i%len(cities)],
				Phone:        fmt.Sprintf("0812345%04d", studentCounter),
				Status:       "aktif",
				EntryYear:    entryYear,
			}

			if err := db.Create(&student).Error; err != nil {
				return err
			}

			studentCounter++
		}
	}

	return nil
}
