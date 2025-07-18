package seeders

import (
	"log"

	"gorm.io/gorm"
)

type SeederInterface interface {
	Seed(db *gorm.DB) error
	GetName() string
}

func RunAllSeeders(db *gorm.DB) {
	log.Println("🌱 Starting database seeding...")

	seeders := []SeederInterface{
		&SchoolSeeder{},
		&SchoolAddressSeeder{},
		&MajorSeeder{},
		&SubjectSeeder{},
		&TeacherSeeder{},
		&TeacherAddressSeeder{},
		&ClassSeeder{},
		&StudentSeeder{},
		&StudentAddressSeeder{},
		&ParentSeeder{},
		&SafeAreaSeeder{},
		&AttendanceSettingsSeeder{},
		&AcademicCalendarSeeder{},
		&ClassScheduleSeeder{},
		&ClassSessionSeeder{},
		&AttendanceSeeder{},
		&SubjectAttendanceSeeder{},
		&PermitSeeder{},
		&NotificationSeeder{},
	}

	successCount := 0
	for i, seeder := range seeders {
		log.Printf("🌱 [%d/%d] Seeding %s...", i+1, len(seeders), seeder.GetName())

		if err := seeder.Seed(db); err != nil {
			log.Printf("❌ Failed to seed %s: %v", seeder.GetName(), err)
			continue
		}

		log.Printf("✅ [%d/%d] Seeded %s successfully", i+1, len(seeders), seeder.GetName())
		successCount++
	}

	log.Printf("🎉 Database seeding completed! (%d/%d successful)", successCount, len(seeders))
}
