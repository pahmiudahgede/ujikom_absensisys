package seeders

import (
	"log"

	"gorm.io/gorm"
)

type SeederInterface interface {
	Run(db *gorm.DB) error
	GetName() string
}

type BaseSeeder struct {
	Name string
}

func (s BaseSeeder) GetName() string {
	return s.Name
}

func RunAllSeeders(db *gorm.DB) {
	log.Println("🌱 Starting database seeding...")

	seeders := []SeederInterface{
		&SchoolSeeder{BaseSeeder: BaseSeeder{Name: "SchoolSeeder"}},
		&JurusanSeeder{BaseSeeder: BaseSeeder{Name: "JurusanSeeder"}},
		&SubjectSeeder{BaseSeeder: BaseSeeder{Name: "SubjectSeeder"}},
		&TeacherSeeder{BaseSeeder: BaseSeeder{Name: "TeacherSeeder"}},
		&ClassSeeder{BaseSeeder: BaseSeeder{Name: "ClassSeeder"}},
		&StudentSeeder{BaseSeeder: BaseSeeder{Name: "StudentSeeder"}},
		&ParentSeeder{BaseSeeder: BaseSeeder{Name: "ParentSeeder"}},
		&AddressSeeder{BaseSeeder: BaseSeeder{Name: "AddressSeeder"}},
		&SafeAreaSeeder{BaseSeeder: BaseSeeder{Name: "SafeAreaSeeder"}},
		&AttendanceSettingsSeeder{BaseSeeder: BaseSeeder{Name: "AttendanceSettingsSeeder"}},
		&AcademicCalendarSeeder{BaseSeeder: BaseSeeder{Name: "AcademicCalendarSeeder"}},
		&HolidaySeeder{BaseSeeder: BaseSeeder{Name: "HolidaySeeder"}},
		&ClassScheduleSeeder{BaseSeeder: BaseSeeder{Name: "ClassScheduleSeeder"}},
		&SubjectAttendanceSettingsSeeder{BaseSeeder: BaseSeeder{Name: "SubjectAttendanceSettingsSeeder"}},
	}

	for i, seeder := range seeders {
		log.Printf("🌱 [%d/%d] Running %s...", i+1, len(seeders), seeder.GetName())

		if err := seeder.Run(db); err != nil {
			log.Printf("❌ Failed to run %s: %v", seeder.GetName(), err)
			continue
		}

		log.Printf("✅ [%d/%d] %s completed successfully", i+1, len(seeders), seeder.GetName())
	}

	log.Println("🎉 Database seeding completed successfully!")
}

func DataExists(db *gorm.DB, model interface{}, condition string, args ...interface{}) bool {
	var count int64
	db.Model(model).Where(condition, args...).Count(&count)
	return count > 0
}
