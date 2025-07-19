package seeders

import (
	"absensibe/models"
	"log"

	"gorm.io/gorm"
)

func SeedClasses(db *gorm.DB) error {
	log.Println("🎓 Seeding classes...")

	classes := []models.Class{
		{
			ID:                "550e8400-e29b-41d4-a716-446655440030",
			Name:              "XI TBSM 1",
			Grade:             "XI",
			Major:             "TBSM",
			SchoolID:          "550e8400-e29b-41d4-a716-446655440001",
			HomeroomTeacherID: stringPtr("550e8400-e29b-41d4-a716-446655440012"), // MAWARDI,S.Pd sebagai wali kelas
			AcademicYear:      "2024/2025",
			IsActive:          true,
		},
	}

	for _, class := range classes {
		if err := db.FirstOrCreate(&class, models.Class{Name: class.Name, AcademicYear: class.AcademicYear}).Error; err != nil {
			return err
		}
	}

	log.Printf("✅ Successfully seeded %d classes", len(classes))
	return nil
}
