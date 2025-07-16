// ===== seeders/class_seeder.go =====
package seeders

import (
	"absensibe/models"
	"fmt"
	"gorm.io/gorm"
)

type ClassSeeder struct {
	BaseSeeder
}

func (s *ClassSeeder) Run(db *gorm.DB) error {
	// Skip if data already exists
	if DataExists(db, &models.Class{}, "name = ?", "X-RPL-1") {
		return nil
	}

	// Get schools with their jurusan
	var schools []models.School
	if err := db.Preload("Jurusan").Find(&schools).Error; err != nil {
		return err
	}

	if len(schools) == 0 {
		return nil
	}

	grades := []string{"X", "XI", "XII"}
	academicYear := "2024/2025"

	for _, school := range schools {
		// Get teachers for homeroom assignment
		var teachers []models.Teacher
		db.Where("school_id = ?", school.ID).Find(&teachers)
		
		teacherIndex := 0

		for _, jurusan := range school.Jurusan {
			// Create 2 classes per grade per jurusan
			for _, grade := range grades {
				for classNum := 1; classNum <= 2; classNum++ {
					className := fmt.Sprintf("%s-%s-%d", grade, jurusan.Code, classNum)
					
					class := models.Class{
						Name:          className,
						Grade:         grade,
						JurusanID:     jurusan.ID,
						SchoolID:      school.ID,
						MaxStudents:   36,
						AcademicYear:  academicYear,
						IsActive:      true,
					}

					// Assign homeroom teacher if available
					if teacherIndex < len(teachers) {
						class.HomeroomTeacherID = &teachers[teacherIndex].ID
						teacherIndex = (teacherIndex + 1) % len(teachers)
					}

					if err := db.Create(&class).Error; err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}