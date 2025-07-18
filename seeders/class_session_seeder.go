package seeders

import (
	"absensibe/models"
	"time"

	"gorm.io/gorm"
)

type ClassSessionSeeder struct{}

func (s *ClassSessionSeeder) GetName() string {
	return "Class Sessions"
}

func (s *ClassSessionSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.ClassSession{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var schedule models.ClassSchedule
	if err := db.First(&schedule).Error; err != nil {
		return err
	}

	var teacher models.Teacher
	if err := db.First(&teacher).Error; err != nil {
		return err
	}

	now := time.Now()
	actualStart := time.Date(now.Year(), now.Month(), now.Day(), 7, 30, 0, 0, time.UTC)
	actualEnd := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, time.UTC)

	sessions := []models.ClassSession{
		{
			ScheduleID:      schedule.ID,
			Date:            time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
			ActualStartTime: &actualStart,
			ActualEndTime:   &actualEnd,
			Topic:           stringPtr("Pengenalan Pemrograman"),
			Material:        stringPtr("Materi dasar tentang algoritma dan pemrograman"),
			Status:          "completed",
			Notes:           stringPtr("Siswa antusias dalam pembelajaran"),
			CreatedBy:       teacher.ID,
		},
	}

	return db.Create(&sessions).Error
}
