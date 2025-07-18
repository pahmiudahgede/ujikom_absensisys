package seeders

import (
	"absensibe/models"

	"gorm.io/gorm"
)

type ClassSeeder struct{}

func (s *ClassSeeder) GetName() string {
	return "Classes"
}

func (s *ClassSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Class{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var school models.School
	if err := db.First(&school).Error; err != nil {
		return err
	}

	var majors []models.Major
	if err := db.Find(&majors).Error; err != nil {
		return err
	}

	var teachers []models.Teacher
	if err := db.Find(&teachers).Error; err != nil {
		return err
	}

	if len(majors) == 0 || len(teachers) == 0 {
		return nil
	}

	classes := []models.Class{
		{
			Name:              "X-RPL-1",
			Grade:             "X",
			MajorID:           majors[0].ID, // RPL
			SchoolID:          school.ID,
			HomeroomTeacherID: &teachers[0].ID,
			MaxStudents:       36,
			AcademicYear:      "2024/2025",
			IsActive:          true,
		},
		{
			Name:              "X-RPL-2",
			Grade:             "X",
			MajorID:           majors[0].ID, // RPL
			SchoolID:          school.ID,
			HomeroomTeacherID: &teachers[1].ID,
			MaxStudents:       36,
			AcademicYear:      "2024/2025",
			IsActive:          true,
		},
		{
			Name:              "XI-RPL-1",
			Grade:             "XI",
			MajorID:           majors[0].ID, // RPL
			SchoolID:          school.ID,
			HomeroomTeacherID: &teachers[2].ID,
			MaxStudents:       36,
			AcademicYear:      "2024/2025",
			IsActive:          true,
		},
	}

	if len(majors) > 1 {
		classes = append(classes, models.Class{
			Name:         "X-TKJ-1",
			Grade:        "X",
			MajorID:      majors[1].ID, // TKJ
			SchoolID:     school.ID,
			MaxStudents:  36,
			AcademicYear: "2024/2025",
			IsActive:     true,
		})
	}

	return db.Create(&classes).Error
}
