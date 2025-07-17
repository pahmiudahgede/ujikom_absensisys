// ===== seeders/teacher_seeder.go =====
package seeders

import (
	"absensibe/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type TeacherSeeder struct {
	BaseSeeder
}

func (s *TeacherSeeder) Run(db *gorm.DB) error {
	// Skip if data already exists
	if DataExists(db, &models.Teacher{}, "nip = ?", "197501012000031001") {
		return nil
	}

	// Get school IDs
	var schools []models.School
	if err := db.Find(&schools).Error; err != nil {
		return err
	}

	if len(schools) == 0 {
		return nil
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	teacherData := []struct {
		NIP          string
		Fullname     string
		Email        string
		Phone        string
		Gender       string
		BirthDate    string
		PlaceOfBirth string
		EmployeeType string
		Status       string
	}{
		{
			NIP:          "197501012000031001",
			Fullname:     "Dr. Ahmad Santoso, S.Pd., M.M.",
			Email:        "ahmad.santoso@smk.edu",
			Phone:        "08123456789",
			Gender:       "L",
			BirthDate:    "1975-01-01",
			PlaceOfBirth: "Jakarta",
			EmployeeType: "PNS",
			Status:       "aktif",
		},
		{
			NIP:          "198203152005012002",
			Fullname:     "Siti Rahayu, S.Kom., M.T.",
			Email:        "siti.rahayu@smk.edu",
			Phone:        "08123456790",
			Gender:       "P",
			BirthDate:    "1982-03-15",
			PlaceOfBirth: "Bandung",
			EmployeeType: "PNS",
			Status:       "aktif",
		},
		{
			NIP:          "198506102010121003",
			Fullname:     "Budi Prasetyo, S.T., M.Kom.",
			Email:        "budi.prasetyo@smk.edu",
			Phone:        "08123456791",
			Gender:       "L",
			BirthDate:    "1985-06-10",
			PlaceOfBirth: "Surabaya",
			EmployeeType: "PNS",
			Status:       "aktif",
		},
		{
			NIP:          "199012252015022004",
			Fullname:     "Dewi Lestari, S.Pd., M.Pd.",
			Email:        "dewi.lestari@smk.edu",
			Phone:        "08123456792",
			Gender:       "P",
			BirthDate:    "1990-12-25",
			PlaceOfBirth: "Yogyakarta",
			EmployeeType: "PPPK",
			Status:       "aktif",
		},
		{
			NIP:          "198808182012031005",
			Fullname:     "Andi Wijaya, S.Kom.",
			Email:        "andi.wijaya@smk.edu",
			Phone:        "08123456793",
			Gender:       "L",
			BirthDate:    "1988-08-18",
			PlaceOfBirth: "Medan",
			EmployeeType: "GTT",
			Status:       "aktif",
		},
		{
			NIP:          "199205052018012006",
			Fullname:     "Maya Sari, S.Sn., M.Ds.",
			Email:        "maya.sari@smk.edu",
			Phone:        "08123456794",
			Gender:       "P",
			BirthDate:    "1992-05-05",
			PlaceOfBirth: "Denpasar",
			EmployeeType: "GTT",
			Status:       "aktif",
		},
	}

	// Create teachers for each school
	for _, school := range schools {
		for _, data := range teacherData {
			birthDate, _ := time.Parse("2006-01-02", data.BirthDate)

			teacher := models.Teacher{
				NIP:          data.NIP,
				SchoolID:     school.ID,
				Fullname:     data.Fullname,
				Email:        data.Email,
				Phone:        data.Phone,
				Password:     string(hashedPassword),
				Gender:       data.Gender,
				BirthDate:    birthDate,
				PlaceOfBirth: data.PlaceOfBirth,
				EmployeeType: data.EmployeeType,
				Status:       data.Status,
			}

			if err := db.Create(&teacher).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
