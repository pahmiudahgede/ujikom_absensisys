package seeders

import (
	"absensibe/models"

	"gorm.io/gorm"
)

type SubjectSeeder struct{}

func (s *SubjectSeeder) GetName() string {
	return "Subjects"
}

func (s *SubjectSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Subject{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var school models.School
	if err := db.First(&school).Error; err != nil {
		return err
	}

	subjects := []models.Subject{
		// Mata Pelajaran Umum
		{Code: "BIND", Name: "Bahasa Indonesia", SchoolID: school.ID, CreditHours: 3, IsActive: true},
		{Code: "BING", Name: "Bahasa Inggris", SchoolID: school.ID, CreditHours: 3, IsActive: true},
		{Code: "MTK", Name: "Matematika", SchoolID: school.ID, CreditHours: 4, IsActive: true},
		{Code: "SEJARAH", Name: "Sejarah Indonesia", SchoolID: school.ID, CreditHours: 2, IsActive: true},
		{Code: "PKN", Name: "Pendidikan Kewarganegaraan", SchoolID: school.ID, CreditHours: 2, IsActive: true},
		{Code: "AGAMA", Name: "Pendidikan Agama dan Budi Pekerti", SchoolID: school.ID, CreditHours: 3, IsActive: true},
		{Code: "PJOK", Name: "Pendidikan Jasmani, Olahraga, dan Kesehatan", SchoolID: school.ID, CreditHours: 2, IsActive: true},

		// Mata Pelajaran Kejuruan - RPL
		{Code: "PEMROG", Name: "Pemrograman Dasar", SchoolID: school.ID, CreditHours: 6, IsActive: true},
		{Code: "BASIS_DATA", Name: "Basis Data", SchoolID: school.ID, CreditHours: 6, IsActive: true},
		{Code: "WEB", Name: "Pemrograman Web", SchoolID: school.ID, CreditHours: 8, IsActive: true},
		{Code: "MOBILE", Name: "Pemrograman Mobile", SchoolID: school.ID, CreditHours: 6, IsActive: true},
		{Code: "RPL", Name: "Rekayasa Perangkat Lunak", SchoolID: school.ID, CreditHours: 6, IsActive: true},

		// Mata Pelajaran Kejuruan - TKJ
		{Code: "JARKOM", Name: "Jaringan Komputer", SchoolID: school.ID, CreditHours: 6, IsActive: true},
		{Code: "SERVER", Name: "Administrasi Server", SchoolID: school.ID, CreditHours: 6, IsActive: true},
		{Code: "KEAMANAN", Name: "Keamanan Jaringan", SchoolID: school.ID, CreditHours: 4, IsActive: true},

		// Mata Pelajaran Kejuruan - MM
		{Code: "GRAFIS", Name: "Desain Grafis", SchoolID: school.ID, CreditHours: 6, IsActive: true},
		{Code: "VIDEO", Name: "Video Editing", SchoolID: school.ID, CreditHours: 6, IsActive: true},
		{Code: "ANIMASI", Name: "Animasi 2D/3D", SchoolID: school.ID, CreditHours: 6, IsActive: true},
	}

	return db.Create(&subjects).Error
}
