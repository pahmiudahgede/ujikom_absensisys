// ===== seeders/subject_seeder.go =====
package seeders

import (
	"absensibe/models"
	"gorm.io/gorm"
)

type SubjectSeeder struct {
	BaseSeeder
}

func (s *SubjectSeeder) Run(db *gorm.DB) error {
	// Skip if data already exists
	if DataExists(db, &models.Subject{}, "code = ?", "MTK") {
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

	// Common subjects for all schools
	commonSubjects := []struct {
		Code         string
		Name         string
		CreditHours  int
		Description  string
	}{
		{
			Code:        "MTK",
			Name:        "Matematika",
			CreditHours: 4,
			Description: "Mata pelajaran matematika dasar dan terapan",
		},
		{
			Code:        "BIND",
			Name:        "Bahasa Indonesia",
			CreditHours: 3,
			Description: "Mata pelajaran bahasa Indonesia",
		},
		{
			Code:        "BING",
			Name:        "Bahasa Inggris",
			CreditHours: 3,
			Description: "Mata pelajaran bahasa Inggris",
		},
		{
			Code:        "PJOK",
			Name:        "Pendidikan Jasmani dan Kesehatan",
			CreditHours: 2,
			Description: "Mata pelajaran olahraga dan kesehatan",
		},
		{
			Code:        "PKN",
			Name:        "Pendidikan Kewarganegaraan",
			CreditHours: 2,
			Description: "Mata pelajaran pendidikan kewarganegaraan",
		},
		{
			Code:        "PAI",
			Name:        "Pendidikan Agama Islam",
			CreditHours: 2,
			Description: "Mata pelajaran pendidikan agama Islam",
		},
		{
			Code:        "SEJIND",
			Name:        "Sejarah Indonesia",
			CreditHours: 2,
			Description: "Mata pelajaran sejarah Indonesia",
		},
		{
			Code:        "KIMIA",
			Name:        "Kimia",
			CreditHours: 3,
			Description: "Mata pelajaran kimia",
		},
		{
			Code:        "FISIKA",
			Name:        "Fisika",
			CreditHours: 3,
			Description: "Mata pelajaran fisika",
		},
	}

	// Vocational subjects by major
	vocationalSubjects := map[string][]struct {
		Code         string
		Name         string
		CreditHours  int
		Description  string
	}{
		"RPL": {
			{
				Code:        "PROGDAS",
				Name:        "Pemrograman Dasar",
				CreditHours: 6,
				Description: "Mata pelajaran pemrograman dasar",
			},
			{
				Code:        "PROGWEB",
				Name:        "Pemrograman Web",
				CreditHours: 6,
				Description: "Mata pelajaran pemrograman web",
			},
			{
				Code:        "PROGMOB",
				Name:        "Pemrograman Mobile",
				CreditHours: 6,
				Description: "Mata pelajaran pemrograman mobile",
			},
			{
				Code:        "BASDAT",
				Name:        "Basis Data",
				CreditHours: 4,
				Description: "Mata pelajaran basis data",
			},
		},
		"TKJ": {
			{
				Code:        "JARKOM",
				Name:        "Jaringan Komputer",
				CreditHours: 6,
				Description: "Mata pelajaran jaringan komputer",
			},
			{
				Code:        "SISOP",
				Name:        "Sistem Operasi",
				CreditHours: 4,
				Description: "Mata pelajaran sistem operasi",
			},
			{
				Code:        "ADMSER",
				Name:        "Administrasi Server",
				CreditHours: 6,
				Description: "Mata pelajaran administrasi server",
			},
		},
		"MM": {
			{
				Code:        "DESGRAF",
				Name:        "Desain Grafis",
				CreditHours: 6,
				Description: "Mata pelajaran desain grafis",
			},
			{
				Code:        "VIDED",
				Name:        "Video Editing",
				CreditHours: 4,
				Description: "Mata pelajaran video editing",
			},
			{
				Code:        "ANIMASI",
				Name:        "Animasi",
				CreditHours: 6,
				Description: "Mata pelajaran animasi",
			},
		},
	}

	// Create subjects for each school
	for _, school := range schools {
		// Create common subjects
		for _, subjectData := range commonSubjects {
			subject := models.Subject{
				Code:        subjectData.Code,
				Name:        subjectData.Name,
				SchoolID:    school.ID,
				CreditHours: subjectData.CreditHours,
				Description: &subjectData.Description,
				IsActive:    true,
			}

			if err := db.Create(&subject).Error; err != nil {
				return err
			}
		}

		// Create vocational subjects based on available jurusan
		var jurusans []models.Jurusan
		db.Where("school_id = ?", school.ID).Find(&jurusans)

		for _, jurusan := range jurusans {
			if subjects, exists := vocationalSubjects[jurusan.Code]; exists {
				for _, subjectData := range subjects {
					subject := models.Subject{
						Code:        subjectData.Code,
						Name:        subjectData.Name,
						SchoolID:    school.ID,
						CreditHours: subjectData.CreditHours,
						Description: &subjectData.Description,
						IsActive:    true,
					}

					if err := db.Create(&subject).Error; err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}