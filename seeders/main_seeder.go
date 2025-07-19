package seeders

import (
	"log"

	"gorm.io/gorm"
)

func RunAllSeeders(db *gorm.DB) error {
	log.Println("🌱 Starting database seeding...")

	seeders := []func(*gorm.DB) error{
		SeedSchools,
		SeedTeachers,
		SeedSubjects,
		SeedClasses,
		SeedStudents,
		SeedClassSchedules,
		SeedAttendanceSettings,
		SeedSafeAreas,
		SeedHolidays,
		SeedSampleAttendances,
	}

	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return err
		}
	}

	log.Println("✅ Database seeding completed successfully!")
	return nil
}
