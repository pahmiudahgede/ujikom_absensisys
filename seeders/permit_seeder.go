package seeders

import (
	"absensibe/models"
	"time"

	"gorm.io/gorm"
)

type PermitSeeder struct{}

func (s *PermitSeeder) GetName() string {
	return "Permits"
}

func (s *PermitSeeder) Seed(db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.Permit{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var students []models.Student
	if err := db.Limit(2).Find(&students).Error; err != nil {
		return err
	}

	if len(students) == 0 {
		return nil
	}

	now := time.Now()

	permits := []models.Permit{
		{
			StudentID:      students[0].ID,
			Type:           "sakit",
			Reason:         "Demam dan batuk",
			StartDate:      time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
			DocumentProof:  stringPtr("surat_dokter.pdf"),
			ApprovalStatus: "approved",
		},
	}

	if len(students) > 1 {
		permits = append(permits, models.Permit{
			StudentID:      students[1].ID,
			Type:           "izin",
			Reason:         "Acara keluarga",
			StartDate:      time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC),
			EndDate:        time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC),
			ApprovalStatus: "pending",
		})
	}

	return db.Create(&permits).Error
}
