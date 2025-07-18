package seeders

import (
	"absensibe/models"
	"time"

	"gorm.io/gorm"
)

type SubjectAttendanceSeeder struct{}

func (s *SubjectAttendanceSeeder) GetName() string {
	return "Subject Attendances"
}

func (s *SubjectAttendanceSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.SubjectAttendance{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var session models.ClassSession
	if err := db.First(&session).Error; err != nil {
		return err
	}

	var students []models.Student
	if err := db.Limit(5).Find(&students).Error; err != nil {
		return err
	}

	if len(students) == 0 {
		return nil
	}

	now := time.Now()
	checkinTime := time.Date(now.Year(), now.Month(), now.Day(), 7, 45, 0, 0, time.UTC)

	attendances := []models.SubjectAttendance{}
	for i, student := range students {
		status := "hadir"
		if i == 3 {
			status = "alpha"
		}

		attendance := models.SubjectAttendance{
			SessionID:   session.ID,
			StudentID:   student.ID,
			Status:      status,
			CheckInTime: &checkinTime,
			Latitude:    float64Ptr(-6.1751),
			Longitude:   float64Ptr(106.8650),
		}
		attendances = append(attendances, attendance)
	}

	return db.Create(&attendances).Error
}
