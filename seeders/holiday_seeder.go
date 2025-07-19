package seeders

import (
	"absensibe/models"
	"log"
	"time"

	"gorm.io/gorm"
)

func SeedHolidays(db *gorm.DB) error {
	log.Println("🎉 Seeding holidays...")

	holidays := []models.Holiday{
		{
			ID:          "550e8400-e29b-41d4-a716-446655440080",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			Title:       "Tahun Baru 2025",
			Description: stringPtr("Libur nasional tahun baru"),
			StartDate:   time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			Type:        "libur_nasional",
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440081",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			Title:       "Libur Semester Ganjil",
			Description: stringPtr("Libur akhir semester ganjil 2024/2025"),
			StartDate:   time.Date(2024, 12, 23, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 1, 7, 0, 0, 0, 0, time.UTC),
			Type:        "libur_sekolah",
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440082",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			Title:       "Ujian Tengah Semester",
			Description: stringPtr("Pelaksanaan ujian tengah semester genap"),
			StartDate:   time.Date(2025, 3, 17, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 3, 28, 0, 0, 0, 0, time.UTC),
			Type:        "ujian",
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440083",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			Title:       "Hari Kemerdekaan RI",
			Description: stringPtr("Peringatan hari kemerdekaan Indonesia"),
			StartDate:   time.Date(2025, 8, 17, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 8, 17, 0, 0, 0, 0, time.UTC),
			Type:        "libur_nasional",
			IsActive:    true,
		},
		{
			ID:          "550e8400-e29b-41d4-a716-446655440084",
			SchoolID:    "550e8400-e29b-41d4-a716-446655440001",
			Title:       "Workshop Otomotif Fair",
			Description: stringPtr("Event pameran otomotif sekolah"),
			StartDate:   time.Date(2025, 4, 15, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 4, 17, 0, 0, 0, 0, time.UTC),
			Type:        "event",
			IsActive:    true,
		},
	}

	for _, holiday := range holidays {
		if err := db.FirstOrCreate(&holiday, models.Holiday{
			SchoolID:  holiday.SchoolID,
			Title:     holiday.Title,
			StartDate: holiday.StartDate,
		}).Error; err != nil {
			return err
		}
	}

	log.Printf("✅ Successfully seeded %d holidays", len(holidays))
	return nil
}
