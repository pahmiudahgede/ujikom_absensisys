package seeders

import (
	"absensibe/models"
	"time"

	"gorm.io/gorm"
)

type NotificationSeeder struct{}

func (s *NotificationSeeder) GetName() string {
	return "Notifications"
}

func (s *NotificationSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Notification{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var students []models.Student
	if err := db.Limit(3).Find(&students).Error; err != nil {
		return err
	}

	if len(students) == 0 {
		return nil
	}

	notifications := []models.Notification{
		{
			RecipientID:   students[0].ID,
			RecipientType: "student",
			Title:         "Absensi Hari Ini",
			Message:       "Jangan lupa untuk melakukan absensi masuk hari ini sebelum jam 07:30",
			Type:          "reminder",
			IsRead:        false,
			SentAt:        time.Now().UTC(),
		},
		{
			RecipientID:   students[1].ID,
			RecipientType: "student",
			Title:         "Izin Disetujui",
			Message:       "Permohonan izin Anda telah disetujui oleh wali kelas",
			Type:          "permit",
			IsRead:        false,
			SentAt:        time.Now().UTC(),
		},
	}

	return db.Create(&notifications).Error
}
