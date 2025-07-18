package seeders

import (
	"absensibe/models"
	"time"

	"gorm.io/gorm"
)

type AttendanceSeeder struct{}

func (s *AttendanceSeeder) GetName() string {
	return "Attendances"
}

func (s *AttendanceSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Attendance{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var students []models.Student
	if err := db.Limit(5).Find(&students).Error; err != nil {
		return err
	}

	if len(students) == 0 {
		return nil
	}

	now := time.Now()
	checkinTime := time.Date(now.Year(), now.Month(), now.Day(), 7, 15, 0, 0, time.UTC)
	checkoutTime := time.Date(now.Year(), now.Month(), now.Day(), 15, 30, 0, 0, time.UTC)

	attendances := []models.Attendance{}
	for i, student := range students {
		status := "hadir"
		if i == 4 {
			status = "terlambat"
		}

		attendance := models.Attendance{
			StudentID:    student.ID,
			Date:         time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
			CheckInTime:  &checkinTime,
			CheckOutTime: &checkoutTime,
			Status:       status,
		}
		attendances = append(attendances, attendance)
	}

	return db.Create(&attendances).Error
}
