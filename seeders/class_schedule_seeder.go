package seeders

import (
	"absensibe/models"
	"time"

	"gorm.io/gorm"
)

type ClassScheduleSeeder struct{}

func (s *ClassScheduleSeeder) GetName() string {
	return "Class Schedules"
}

func (s *ClassScheduleSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.ClassSchedule{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var class models.Class
	if err := db.First(&class).Error; err != nil {
		return err
	}

	var subjects []models.Subject
	if err := db.Limit(5).Find(&subjects).Error; err != nil {
		return err
	}

	var teachers []models.Teacher
	if err := db.Limit(3).Find(&teachers).Error; err != nil {
		return err
	}

	if len(subjects) == 0 || len(teachers) == 0 {
		return nil
	}

	startTime1, _ := time.Parse("15:04:05", "07:30:00")
	endTime1, _ := time.Parse("15:04:05", "09:00:00")
	startTime2, _ := time.Parse("15:04:05", "09:15:00")
	endTime2, _ := time.Parse("15:04:05", "10:45:00")

	schedules := []models.ClassSchedule{
		{
			ClassID:      class.ID,
			SubjectID:    subjects[0].ID,
			TeacherID:    teachers[0].ID,
			DayOfWeek:    "senin",
			StartTime:    startTime1,
			EndTime:      endTime1,
			Room:         stringPtr("Lab RPL 1"),
			AcademicYear: "2024/2025",
			Semester:     "ganjil",
			IsActive:     true,
		},
		{
			ClassID:      class.ID,
			SubjectID:    subjects[1].ID,
			TeacherID:    teachers[1].ID,
			DayOfWeek:    "senin",
			StartTime:    startTime2,
			EndTime:      endTime2,
			Room:         stringPtr("Ruang 101"),
			AcademicYear: "2024/2025",
			Semester:     "ganjil",
			IsActive:     true,
		},
	}

	return db.Create(&schedules).Error
}
