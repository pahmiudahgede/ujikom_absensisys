package seeders

import (
	"absensibe/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SubjectAttendanceSettingsSeeder struct {
	BaseSeeder
}

func (s *SubjectAttendanceSettingsSeeder) Run(db *gorm.DB) error {
	if DataExists(db, &models.SubjectAttendanceSettings{}, "check_in_tolerance = ?", 15) {
		return nil
	}

	var schedules []struct {
		SubjectID string
		ClassID   string
	}

	err := db.Model(&models.ClassSchedule{}).
		Select("DISTINCT subject_id, class_id").
		Find(&schedules).Error

	if err != nil {
		return err
	}

	for _, schedule := range schedules {

		var existingCount int64
		db.Model(&models.SubjectAttendanceSettings{}).
			Where("subject_id = ? AND class_id = ?", schedule.SubjectID, schedule.ClassID).
			Count(&existingCount)

		if existingCount > 0 {
			continue
		}

		settings := models.SubjectAttendanceSettings{
			SubjectID:             schedule.SubjectID,
			ClassID:               schedule.ClassID,
			CheckInTolerance:      15,
			AutoCheckout:          true,
			RequirePhoto:          false,
			RequireLocation:       false,
			MinAttendanceDuration: 30,
		}

		if err := db.Create(&settings).Error; err != nil {
			return fmt.Errorf("failed to create subject attendance settings for subject %s, class %s: %v",
				schedule.SubjectID, schedule.ClassID, err)
		}
	}

	return nil
}

func float64Ptr(f float64) *float64 {
	return &f
}

func mustParseDate(dateStr string) time.Time {
	t, _ := time.Parse("2006-01-02", dateStr)
	return t
}
