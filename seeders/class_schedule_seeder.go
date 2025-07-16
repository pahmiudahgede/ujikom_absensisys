package seeders

import (
	"absensibe/models"
	"fmt"

	"gorm.io/gorm"
)

type ClassScheduleSeeder struct {
	BaseSeeder
}

func (s *ClassScheduleSeeder) Run(db *gorm.DB) error {
	if DataExists(db, &models.ClassSchedule{}, "day_of_week = ?", "senin") {
		return nil
	}

	var classes []models.Class
	db.Preload("School").Find(&classes)

	days := []string{"senin", "selasa", "rabu", "kamis", "jumat"}
	timeSlots := []struct {
		start string
		end   string
	}{
		{"07:30:00", "08:15:00"},
		{"08:15:00", "09:00:00"},
		{"09:15:00", "10:00:00"},
		{"10:00:00", "10:45:00"},
		{"11:00:00", "11:45:00"},
		{"12:30:00", "13:15:00"},
		{"13:15:00", "14:00:00"},
		{"14:00:00", "14:45:00"},
	}

	for _, class := range classes {
		var subjects []models.Subject
		db.Where("school_id = ?", class.School.ID).Find(&subjects)

		var teachers []models.Teacher
		db.Where("school_id = ?", class.School.ID).Find(&teachers)

		if len(subjects) == 0 || len(teachers) == 0 {
			continue
		}

		subjectIndex := 0
		teacherIndex := 0

		for _, day := range days {
			for i, slot := range timeSlots {
				if i >= 6 { // Max 6 slots per day
					break
				}

				// Create TimeOnly from string
				startTime, err := models.NewTimeOnlyFromString(slot.start)
				if err != nil {
					continue
				}

				endTime, err := models.NewTimeOnlyFromString(slot.end)
				if err != nil {
					continue
				}

				schedule := models.ClassSchedule{
					ClassID:      class.ID,
					SubjectID:    subjects[subjectIndex%len(subjects)].ID,
					TeacherID:    teachers[teacherIndex%len(teachers)].ID,
					DayOfWeek:    day,
					StartTime:    startTime,
					EndTime:      endTime,
					Room:         stringPtr(fmt.Sprintf("R-%s-%d", class.Name, i+1)),
					AcademicYear: "2024/2025",
					Semester:     "ganjil",
					IsActive:     true,
				}

				if err := db.Create(&schedule).Error; err != nil {
					return fmt.Errorf("failed to create schedule for class %s: %v", class.Name, err)
				}

				subjectIndex++
				teacherIndex++
			}
		}
	}

	return nil
}
